// Licensed to Alexandre VILAIN under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Alexandre VILAIN licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package controllers

import (
	"context"
	"errors"
	"time"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"github.com/alexandrevilain/temporal-operator/api/v1beta1"
	certmanagerv1 "github.com/cert-manager/cert-manager/pkg/apis/certmanager/v1"
	istionetworkingv1beta1 "istio.io/client-go/pkg/apis/networking/v1beta1"
	istiosecurityv1beta1 "istio.io/client-go/pkg/apis/security/v1beta1"
	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/alexandrevilain/temporal-operator/pkg/reconciler"
	"github.com/alexandrevilain/temporal-operator/pkg/resourceset"
	"github.com/alexandrevilain/temporal-operator/pkg/status"
	"github.com/alexandrevilain/temporal-operator/pkg/version"
)

const (
	ownerKey  = ".metadata.controller"
	ownerKind = "Cluster"
)

// TemporalClusterReconciler reconciles a Cluster object
type TemporalClusterReconciler struct {
	reconciler.Base

	CertManagerAvailable bool
	IstioAvailable       bool
}

//+kubebuilder:rbac:groups="",resources=services,verbs=get;list;watch;create;update;delete
//+kubebuilder:rbac:groups="",resources=configmaps,verbs=get;list;watch;create;update
//+kubebuilder:rbac:groups="",resources=secrets,verbs=get;list;watch
//+kubebuilder:rbac:groups="",resources=events,verbs=get;create;patch
//+kubebuilder:rbac:groups="",resources=serviceaccounts,verbs=get;list;watch;create;update
//+kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;delete
//+kubebuilder:rbac:groups=batch,resources=jobs,verbs=get;list;watch;create;update;delete
//+kubebuilder:rbac:groups="networking.k8s.io",resources=ingresses,verbs=get;list;watch;create;update;delete
//+kubebuilder:rbac:groups="cert-manager.io",resources=certificates;issuers,verbs=get;list;watch;create;update;delete
//+kubebuilder:rbac:groups="security.istio.io",resources=peerauthentications,verbs=get;list;watch;create;update;delete
//+kubebuilder:rbac:groups="networking.istio.io",resources=destinationrules,verbs=get;list;watch;create;update;delete
//+kubebuilder:rbac:groups=temporal.io,resources=temporalclusters,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=temporal.io,resources=temporalclusters/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=temporal.io,resources=temporalclusters/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *TemporalClusterReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	logger.Info("Starting reconciliation")

	cluster := &v1beta1.TemporalCluster{}
	err := r.Get(ctx, req.NamespacedName, cluster)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}
		return reconcile.Result{}, err
	}

	// Check if the resource has been marked for deletion
	if !cluster.ObjectMeta.DeletionTimestamp.IsZero() {
		logger.Info("Deleting temporal cluster", "name", cluster.Name)
		return reconcile.Result{}, nil
	}

	// Set defaults on unfiled fields.
	updated := r.reconcileDefaults(ctx, cluster)
	if updated {
		err := r.Client.Update(ctx, cluster)
		if err != nil {
			logger.Error(err, "Can't set cluster defaults")
			return r.handleError(ctx, cluster, "", err)
		}
		// As we updated the instance, another reconcile will be triggered.
		return reconcile.Result{}, nil
	}

	// If mTLS is enabled using cert-manager, but cert-manager support is disabled on the controller
	// it can't process the request, return the error.
	if cluster.MTLSWithCertManagerEnabled() && !r.CertManagerAvailable {
		err := errors.New("cert-manager is not available in the cluster")
		logger.Error(err, "Can't process cluster with mTLS enabled using cert-manager")
		return r.handleError(ctx, cluster, v1beta1.TemporalClusterValidationFailedReason, err)
	}

	// Validate that the cluster version is a supported one.
	err = cluster.Spec.Version.Validate()
	if err != nil {
		logger.Error(err, "Can't validate temporal version")
		return r.handleError(ctx, cluster, v1beta1.TemporalClusterValidationFailedReason, err)
	}

	// Ensure ElasticSearch v6 is not used with cluster >= 1.18.0
	// TODO(alexandrevilain): this should be moved to a validating webhook.
	if cluster.Spec.Version.GreaterOrEqual(version.V1_18_0) &&
		cluster.Spec.Persistence.AdvancedVisibilityStore != nil &&
		cluster.Spec.Persistence.AdvancedVisibilityStore.Elasticsearch != nil &&
		cluster.Spec.Persistence.AdvancedVisibilityStore.Elasticsearch.Version == "v6" {
		err := errors.New("temporal cluster version >= 1.18.0 doesn't support ElasticSearch v6")
		return r.handleError(ctx, cluster, v1beta1.TemporalClusterValidationFailedReason, err)
	}

	logger.Info("Retrieved desired cluster version", "version", cluster.Spec.Version.String())

	// Check the ready condition
	cond, exists := v1beta1.GetTemporalClusterReadyCondition(cluster)
	if !exists || cond.ObservedGeneration != cluster.GetGeneration() {
		v1beta1.SetTemporalClusterReady(cluster, metav1.ConditionUnknown, v1beta1.ProgressingReason, "")
		err := r.updateClusterStatus(ctx, cluster)
		if err != nil {
			return r.handleError(ctx, cluster, "", err)
		}
	}

	if requeueAfter, err := r.reconcilePersistence(ctx, cluster); err != nil || requeueAfter > 0 {
		if err != nil {
			logger.Error(err, "Can't reconcile persistence")
			if requeueAfter == 0 {
				requeueAfter = 2 * time.Second
			}
			return r.handleErrorWithRequeue(ctx, cluster, v1beta1.PersistenceReconciliationFailedReason, err, requeueAfter)
		}
		if requeueAfter > 0 {
			err := r.updateClusterStatus(ctx, cluster)
			if err != nil {
				return r.handleError(ctx, cluster, "", err)
			}
			return reconcile.Result{RequeueAfter: requeueAfter}, nil
		}
	}

	if err := r.reconcileResources(ctx, cluster); err != nil {
		logger.Error(err, "Can't reconcile resources")
		return r.handleErrorWithRequeue(ctx, cluster, v1beta1.ResourcesReconciliationFailedReason, err, 2*time.Second)
	}

	return r.handleSuccess(ctx, cluster)
}

func (r *TemporalClusterReconciler) reconcileResources(ctx context.Context, temporalCluster *v1beta1.TemporalCluster) error {
	clusterBuilder := &resourceset.ClusterBuilder{
		Instance: temporalCluster,
		Scheme:   r.Scheme,
	}

	statuses, err := r.ReconcileResources(ctx, temporalCluster, clusterBuilder)
	if err != nil {
		return err
	}

	for _, status := range statuses {
		temporalCluster.Status.AddServiceStatus(status)
	}

	if status.ObservedVersionMatchesDesiredVersion(temporalCluster) {
		temporalCluster.Status.Version = temporalCluster.Spec.Version.String()
	}

	if status.IsClusterReady(temporalCluster) {
		v1beta1.SetTemporalClusterReady(temporalCluster, metav1.ConditionTrue, v1beta1.ServicesReadyReason, "")
	} else {
		v1beta1.SetTemporalClusterReady(temporalCluster, metav1.ConditionFalse, v1beta1.ServicesNotReadyReason, "")
	}

	return r.updateClusterStatus(ctx, temporalCluster)
}

func (r *TemporalClusterReconciler) handleSuccess(ctx context.Context, cluster *v1beta1.TemporalCluster) (ctrl.Result, error) {
	return r.handleSuccessWithRequeue(ctx, cluster, 0)
}

func (r *TemporalClusterReconciler) handleError(ctx context.Context, cluster *v1beta1.TemporalCluster, reason string, err error) (ctrl.Result, error) {
	return r.handleErrorWithRequeue(ctx, cluster, reason, err, 0)
}

func (r *TemporalClusterReconciler) handleSuccessWithRequeue(ctx context.Context, cluster *v1beta1.TemporalCluster, requeueAfter time.Duration) (ctrl.Result, error) {
	v1beta1.SetTemporalClusterReconcileSuccess(cluster, metav1.ConditionTrue, v1beta1.ReconcileSuccessReason, "")
	err := r.updateClusterStatus(ctx, cluster)
	return reconcile.Result{RequeueAfter: requeueAfter}, err
}

func (r *TemporalClusterReconciler) handleErrorWithRequeue(ctx context.Context, cluster *v1beta1.TemporalCluster, reason string, err error, requeueAfter time.Duration) (ctrl.Result, error) {
	r.Recorder.Event(cluster, corev1.EventTypeWarning, "ProcessingError", err.Error())
	if reason == "" {
		reason = v1beta1.ReconcileErrorReason
	}
	v1beta1.SetTemporalClusterReconcileError(cluster, metav1.ConditionTrue, reason, err.Error())
	err = r.updateClusterStatus(ctx, cluster)
	return reconcile.Result{RequeueAfter: requeueAfter}, err
}

func (r *TemporalClusterReconciler) updateClusterStatus(ctx context.Context, cluster *v1beta1.TemporalCluster) error {
	err := r.Status().Update(ctx, cluster)
	if err != nil {
		return err
	}
	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *TemporalClusterReconciler) SetupWithManager(mgr ctrl.Manager) error {
	for _, resource := range []client.Object{&appsv1.Deployment{}, &corev1.ConfigMap{}, &corev1.Service{}, &corev1.ServiceAccount{}, &networkingv1.Ingress{}, &batchv1.Job{}} {
		if err := mgr.GetFieldIndexer().IndexField(context.Background(), resource, ownerKey, addResourceToIndex); err != nil {
			return err
		}
	}

	controller := ctrl.NewControllerManagedBy(mgr).
		For(&v1beta1.TemporalCluster{}, builder.WithPredicates(predicate.Or(
			predicate.GenerationChangedPredicate{},
			predicate.LabelChangedPredicate{},
			predicate.AnnotationChangedPredicate{},
		))).
		Owns(&appsv1.Deployment{}).
		Owns(&corev1.ConfigMap{}).
		Owns(&corev1.Service{}).
		Owns(&corev1.ServiceAccount{}).
		Owns(&networkingv1.Ingress{}).
		Owns(&batchv1.Job{})

	if r.CertManagerAvailable {
		controller = controller.
			Owns(&certmanagerv1.Issuer{}).
			Owns(&certmanagerv1.Certificate{})

		for _, resource := range []client.Object{&certmanagerv1.Issuer{}, &certmanagerv1.Certificate{}} {
			if err := mgr.GetFieldIndexer().IndexField(context.Background(), resource, ownerKey, addCertManagerResourceToIndex); err != nil {
				return err
			}
		}
	}

	if r.IstioAvailable {
		controller = controller.
			Owns(&istiosecurityv1beta1.PeerAuthentication{}).
			Owns(&istionetworkingv1beta1.DestinationRule{})

		for _, resource := range []client.Object{&istiosecurityv1beta1.PeerAuthentication{}, &istionetworkingv1beta1.DestinationRule{}} {
			if err := mgr.GetFieldIndexer().IndexField(context.Background(), resource, ownerKey, addIstioResourceToIndex); err != nil {
				return err
			}
		}
	}

	return controller.Complete(r)
}

func addResourceToIndex(rawObj client.Object) []string {
	switch resourceObject := rawObj.(type) {
	case *appsv1.Deployment,
		*corev1.ConfigMap,
		*corev1.Service,
		*corev1.ServiceAccount,
		*networkingv1.Ingress,
		*batchv1.Job:
		owner := metav1.GetControllerOf(resourceObject)
		return validateAndGetOwner(owner)
	default:
		return nil
	}
}

func addCertManagerResourceToIndex(rawObj client.Object) []string {
	switch resourceObject := rawObj.(type) {
	case *istiosecurityv1beta1.PeerAuthentication,
		*istionetworkingv1beta1.DestinationRule:
		owner := metav1.GetControllerOf(resourceObject)
		return validateAndGetOwner(owner)
	default:
		return nil
	}
}

func addIstioResourceToIndex(rawObj client.Object) []string {
	switch resourceObject := rawObj.(type) {
	case *certmanagerv1.Issuer,
		*certmanagerv1.Certificate:
		owner := metav1.GetControllerOf(resourceObject)
		return validateAndGetOwner(owner)
	default:
		return nil
	}
}

func validateAndGetOwner(owner *metav1.OwnerReference) []string {
	if owner == nil {
		return nil
	}
	if owner.APIVersion != v1beta1.GroupVersion.String() || owner.Kind != ownerKind {
		return nil
	}
	return []string{owner.Name}
}
