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

package certmanager

import (
	"fmt"

	"github.com/alexandrevilain/temporal-operator/api/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

type AdminToolsFrontendClientCertificateBuilder struct {
	GenericFrontendClientCertificateBuilder
}

func NewAdminToolsFrontendClientCertificateBuilder(instance *v1beta1.TemporalCluster, scheme *runtime.Scheme) *AdminToolsFrontendClientCertificateBuilder {
	return &AdminToolsFrontendClientCertificateBuilder{
		GenericFrontendClientCertificateBuilder{
			instance:   instance,
			scheme:     scheme,
			name:       AdmintoolsFrontendClientCertificate,
			secretName: AdmintoolsFrontendClientCertificate,
			commonName: "Admintools client certificate",
			dnsName:    fmt.Sprintf("admintools.%s", instance.ServerName()),
		},
	}
}

func (b *AdminToolsFrontendClientCertificateBuilder) Update(object client.Object) error {
	err := b.GenericFrontendClientCertificateBuilder.Update(object)
	if err != nil {
		return err
	}

	return controllerutil.SetControllerReference(b.instance, object, b.scheme)
}
