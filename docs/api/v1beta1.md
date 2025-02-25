<h1>API reference</h1>
<p>Package v1beta1 contains API Schema definitions for the v1beta1 API group</p>
Resource Types:
<ul class="simple"><li>
<a href="#temporal.io/v1beta1.TemporalCluster">TemporalCluster</a>
</li></ul>
<h3 id="temporal.io/v1beta1.TemporalCluster">TemporalCluster
</h3>
<p>TemporalCluster defines a temporal cluster deployment.</p>
<div class="md-typeset__scrollwrap">
<div class="md-typeset__table">
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br>
string</td>
<td>
<code>temporal.io/v1beta1</code>
</td>
</tr>
<tr>
<td>
<code>kind</code><br>
string
</td>
<td>
<code>TemporalCluster</code>
</td>
</tr>
<tr>
<td>
<code>metadata</code><br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br>
<em>
<a href="#temporal.io/v1beta1.TemporalClusterSpec">
TemporalClusterSpec
</a>
</em>
</td>
<td>
<p>Specification of the desired behavior of the Temporal cluster.</p>
<br/>
<br/>
<table>
<tr>
<td>
<code>image</code><br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Image defines the temporal server docker image the cluster should use for each services.</p>
</td>
</tr>
<tr>
<td>
<code>version</code><br>
<em>
github.com/alexandrevilain/temporal-operator/pkg/version.Version
</em>
</td>
<td>
<em>(Optional)</em>
<p>Version defines the temporal version the cluster to be deployed.
This version impacts the underlying persistence schemas versions.</p>
</td>
</tr>
<tr>
<td>
<code>jobTtlSecondsAfterFinished</code><br>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>JobTtlSecondsAfterFinished is amount of time to keep job pods after jobs are completed.
Defaults to 300 seconds.</p>
</td>
</tr>
<tr>
<td>
<code>numHistoryShards</code><br>
<em>
int32
</em>
</td>
<td>
<p>NumHistoryShards is the desired number of history shards.
This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>services</code><br>
<em>
<a href="#temporal.io/v1beta1.ServicesSpec">
ServicesSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Services allows customizations for for each temporal services deployment.</p>
</td>
</tr>
<tr>
<td>
<code>persistence</code><br>
<em>
<a href="#temporal.io/v1beta1.TemporalPersistenceSpec">
TemporalPersistenceSpec
</a>
</em>
</td>
<td>
<p>Persistence defines temporal persistence configuration.</p>
</td>
</tr>
<tr>
<td>
<code>imagePullSecrets</code><br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#localobjectreference-v1-core">
[]Kubernetes core/v1.LocalObjectReference
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>An optional list of references to secrets in the same namespace
to use for pulling temporal images from registries.</p>
</td>
</tr>
<tr>
<td>
<code>ui</code><br>
<em>
<a href="#temporal.io/v1beta1.TemporalUISpec">
TemporalUISpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>UI allows configuration of the optional temporal web ui deployed alongside the cluster.</p>
</td>
</tr>
<tr>
<td>
<code>admintools</code><br>
<em>
<a href="#temporal.io/v1beta1.TemporalAdminToolsSpec">
TemporalAdminToolsSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>AdminTools allows configuration of the optional admin tool pod deployed alongside the cluster.</p>
</td>
</tr>
<tr>
<td>
<code>mTLS</code><br>
<em>
<a href="#temporal.io/v1beta1.MTLSSpec">
MTLSSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>MTLS allows configuration of the network traffic encryption for the cluster.</p>
</td>
</tr>
<tr>
<td>
<code>metrics</code><br>
<em>
<a href="#temporal.io/v1beta1.MetricsSpec">
MetricsSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Metrics allows configuration of scraping endpoints for stats. prometheus or m3.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br>
<em>
<a href="#temporal.io/v1beta1.TemporalClusterStatus">
TemporalClusterStatus
</a>
</em>
</td>
<td>
<p>Most recent observed status of the Temporal cluster.</p>
</td>
</tr>
</tbody>
</table>
</div>
</div>
<h3 id="temporal.io/v1beta1.CassandraConsistencySpec">CassandraConsistencySpec
</h3>
<p>
(<em>Appears on:</em>
<a href="#temporal.io/v1beta1.CassandraSpec">CassandraSpec</a>)
</p>
<p>CassandraConsistencySpec sets the consistency level for regular &amp; serial queries to Cassandra.</p>
<div class="md-typeset__scrollwrap">
<div class="md-typeset__table">
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>consistency</code><br>
<em>
<a href="https://pkg.go.dev/github.com/gocql/gocql#Consistency">
github.com/gocql/gocql.Consistency
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Consistency sets the default consistency level.
Values identical to gocql Consistency values. (defaults to LOCAL_QUORUM if not set).</p>
</td>
</tr>
<tr>
<td>
<code>serialConsistency</code><br>
<em>
<a href="https://pkg.go.dev/github.com/gocql/gocql#SerialConsistencyy">
github.com/gocql/gocql.SerialConsistency
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>SerialConsistency sets the consistency for the serial prtion of queries. Values identical to gocql SerialConsistency values.
(defaults to LOCAL_SERIAL if not set)</p>
</td>
</tr>
</tbody>
</table>
</div>
</div>
<h3 id="temporal.io/v1beta1.CassandraSpec">CassandraSpec
</h3>
<p>
(<em>Appears on:</em>
<a href="#temporal.io/v1beta1.DatastoreSpec">DatastoreSpec</a>)
</p>
<p>CassandraSpec contains cassandra datastore connections specifications.</p>
<div class="md-typeset__scrollwrap">
<div class="md-typeset__table">
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>hosts</code><br>
<em>
[]string
</em>
</td>
<td>
<p>Hosts is a list of cassandra endpoints.</p>
</td>
</tr>
<tr>
<td>
<code>port</code><br>
<em>
int
</em>
</td>
<td>
<p>Port is the cassandra port used for connection by gocql client.</p>
</td>
</tr>
<tr>
<td>
<code>user</code><br>
<em>
string
</em>
</td>
<td>
<p>User is the cassandra user used for authentication by gocql client.</p>
</td>
</tr>
<tr>
<td>
<code>keyspace</code><br>
<em>
string
</em>
</td>
<td>
<p>Keyspace is the cassandra keyspace.</p>
</td>
</tr>
<tr>
<td>
<code>datacenter</code><br>
<em>
string
</em>
</td>
<td>
<p>Datacenter is the data center filter arg for cassandra.</p>
</td>
</tr>
<tr>
<td>
<code>maxConns</code><br>
<em>
int
</em>
</td>
<td>
<em>(Optional)</em>
<p>MaxConns is the max number of connections to this datastore for a single keyspace.</p>
</td>
</tr>
<tr>
<td>
<code>connectTimeout</code><br>
<em>
<a href="https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#Duration">
Kubernetes meta/v1.Duration
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>ConnectTimeout is a timeout for initial dial to cassandra server.</p>
</td>
</tr>
<tr>
<td>
<code>consistency</code><br>
<em>
<a href="#temporal.io/v1beta1.CassandraConsistencySpec">
CassandraConsistencySpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Consistency configuration.</p>
</td>
</tr>
<tr>
<td>
<code>disableInitialHostLookup</code><br>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>DisableInitialHostLookup instructs the gocql client to connect only using the supplied hosts.</p>
</td>
</tr>
</tbody>
</table>
</div>
</div>
<h3 id="temporal.io/v1beta1.CertificatesDurationSpec">CertificatesDurationSpec
</h3>
<p>
(<em>Appears on:</em>
<a href="#temporal.io/v1beta1.MTLSSpec">MTLSSpec</a>)
</p>
<p>CertificatesDurationSpec defines parameters for the temporal mTLS certificates duration.</p>
<div class="md-typeset__scrollwrap">
<div class="md-typeset__table">
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>rootCACertificate</code><br>
<em>
<a href="https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#Duration">
Kubernetes meta/v1.Duration
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>RootCACertificate is the &lsquo;duration&rsquo; (i.e. lifetime) of the Root CA Certificate.
It defaults to 10 years.</p>
</td>
</tr>
<tr>
<td>
<code>intermediateCAsCertificates</code><br>
<em>
<a href="https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#Duration">
Kubernetes meta/v1.Duration
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>IntermediateCACertificates is the &lsquo;duration&rsquo; (i.e. lifetime) of the intermediate CAs Certificates.
It defaults to 5 years.</p>
</td>
</tr>
<tr>
<td>
<code>clientCertificates</code><br>
<em>
<a href="https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#Duration">
Kubernetes meta/v1.Duration
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>ClientCertificates is the &lsquo;duration&rsquo; (i.e. lifetime) of the client certificates.
It defaults to 1 year.</p>
</td>
</tr>
<tr>
<td>
<code>frontendCertificate</code><br>
<em>
<a href="https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#Duration">
Kubernetes meta/v1.Duration
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>FrontendCertificate is the &lsquo;duration&rsquo; (i.e. lifetime) of the frontend certificate.
It defaults to 1 year.</p>
</td>
</tr>
<tr>
<td>
<code>internodeCertificate</code><br>
<em>
<a href="https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#Duration">
Kubernetes meta/v1.Duration
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>InternodeCertificate is the &lsquo;duration&rsquo; (i.e. lifetime) of the internode certificate.
It defaults to 1 year.</p>
</td>
</tr>
</tbody>
</table>
</div>
</div>
<h3 id="temporal.io/v1beta1.DatastoreSpec">DatastoreSpec
</h3>
<p>
(<em>Appears on:</em>
<a href="#temporal.io/v1beta1.TemporalPersistenceSpec">TemporalPersistenceSpec</a>)
</p>
<p>DatastoreSpec contains temporal datastore specifications.</p>
<div class="md-typeset__scrollwrap">
<div class="md-typeset__table">
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Name is the name of the datatstore.
It should be unique and will be referenced within the persitence spec.
Defaults to &ldquo;default&rdquo; for default sore, &ldquo;visibility&rdquo; for visibility store and
&ldquo;advancedVisibility&rdquo; for advanced visibility store.</p>
</td>
</tr>
<tr>
<td>
<code>sql</code><br>
<em>
<a href="#temporal.io/v1beta1.SQLSpec">
SQLSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>SQL holds all connection parameters for SQL datastores.</p>
</td>
</tr>
<tr>
<td>
<code>elasticsearch</code><br>
<em>
<a href="#temporal.io/v1beta1.ElasticsearchSpec">
ElasticsearchSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Elasticsearch holds all connection parameters for Elasticsearch datastores.</p>
</td>
</tr>
<tr>
<td>
<code>cassandra</code><br>
<em>
<a href="#temporal.io/v1beta1.CassandraSpec">
CassandraSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Cassandra holds all connection parameters for Cassandra datastore.</p>
</td>
</tr>
<tr>
<td>
<code>passwordSecretRef</code><br>
<em>
<a href="#temporal.io/v1beta1.SecretKeyReference">
SecretKeyReference
</a>
</em>
</td>
<td>
<p>PasswordSecret is the reference to the secret holding the password.</p>
</td>
</tr>
<tr>
<td>
<code>tls</code><br>
<em>
<a href="#temporal.io/v1beta1.DatastoreTLSSpec">
DatastoreTLSSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>TLS is an optional option to connect to the datastore using TLS.</p>
</td>
</tr>
</tbody>
</table>
</div>
</div>
<h3 id="temporal.io/v1beta1.DatastoreStatus">DatastoreStatus
</h3>
<p>
(<em>Appears on:</em>
<a href="#temporal.io/v1beta1.TemporalPersistenceStatus">TemporalPersistenceStatus</a>)
</p>
<p>DatastoreStatus contains the current status of a datastore.</p>
<div class="md-typeset__scrollwrap">
<div class="md-typeset__table">
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>created</code><br>
<em>
bool
</em>
</td>
<td>
<p>Created indicates if the database or keyspace has been created.</p>
</td>
</tr>
<tr>
<td>
<code>setup</code><br>
<em>
bool
</em>
</td>
<td>
<p>Setup indicates if tables have been set up.</p>
</td>
</tr>
<tr>
<td>
<code>schemaVersion</code><br>
<em>
github.com/alexandrevilain/temporal-operator/pkg/version.Version
</em>
</td>
<td>
<em>(Optional)</em>
<p>SchemaVersion report the current schema version.</p>
</td>
</tr>
</tbody>
</table>
</div>
</div>
<h3 id="temporal.io/v1beta1.DatastoreTLSSpec">DatastoreTLSSpec
</h3>
<p>
(<em>Appears on:</em>
<a href="#temporal.io/v1beta1.DatastoreSpec">DatastoreSpec</a>)
</p>
<p>DatastoreTLSSpec contains datastore TLS connections specifications.</p>
<div class="md-typeset__scrollwrap">
<div class="md-typeset__table">
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br>
<em>
bool
</em>
</td>
<td>
<p>Enabled defines if the cluster should use a TLS connection to connect to the datastore.</p>
</td>
</tr>
<tr>
<td>
<code>certFileRef</code><br>
<em>
<a href="#temporal.io/v1beta1.SecretKeyReference">
SecretKeyReference
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>CertFileRef is a reference to a secret containing the cert file.</p>
</td>
</tr>
<tr>
<td>
<code>keyFileRef</code><br>
<em>
<a href="#temporal.io/v1beta1.SecretKeyReference">
SecretKeyReference
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>KeyFileRef is a reference to a secret containing the key file.</p>
</td>
</tr>
<tr>
<td>
<code>caFileRef</code><br>
<em>
<a href="#temporal.io/v1beta1.SecretKeyReference">
SecretKeyReference
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>CaFileRef is a reference to a secret containing the ca file.</p>
</td>
</tr>
<tr>
<td>
<code>enableHostVerification</code><br>
<em>
bool
</em>
</td>
<td>
<p>EnableHostVerification defines if the hostname should be verified when connecting to the datastore.</p>
</td>
</tr>
<tr>
<td>
<code>serverName</code><br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>ServerName the datastore should present.</p>
</td>
</tr>
</tbody>
</table>
</div>
</div>
<h3 id="temporal.io/v1beta1.DatastoreType">DatastoreType
(<code>string</code> alias)</h3>
<h3 id="temporal.io/v1beta1.ElasticsearchIndices">ElasticsearchIndices
</h3>
<p>
(<em>Appears on:</em>
<a href="#temporal.io/v1beta1.ElasticsearchSpec">ElasticsearchSpec</a>)
</p>
<p>ElasticsearchIndices holds index names.</p>
<div class="md-typeset__scrollwrap">
<div class="md-typeset__table">
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>visibility</code><br>
<em>
string
</em>
</td>
<td>
<p>Visibility defines visibility&rsquo;s index name.</p>
</td>
</tr>
<tr>
<td>
<code>secondaryVisibility</code><br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>SecondaryVisibility defines secondary visibility&rsquo;s index name.</p>
</td>
</tr>
</tbody>
</table>
</div>
</div>
<h3 id="temporal.io/v1beta1.ElasticsearchSpec">ElasticsearchSpec
</h3>
<p>
(<em>Appears on:</em>
<a href="#temporal.io/v1beta1.DatastoreSpec">DatastoreSpec</a>)
</p>
<p>ElasticsearchSpec contains Elasticsearch datastore connections specifications.</p>
<div class="md-typeset__scrollwrap">
<div class="md-typeset__table">
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>version</code><br>
<em>
string
</em>
</td>
<td>
<p>Version defines the elasticsearch version.</p>
</td>
</tr>
<tr>
<td>
<code>url</code><br>
<em>
string
</em>
</td>
<td>
<p>URL is the connection url to connect to the instance.</p>
</td>
</tr>
<tr>
<td>
<code>username</code><br>
<em>
string
</em>
</td>
<td>
<p>Username is the username to be used for the connection.</p>
</td>
</tr>
<tr>
<td>
<code>indices</code><br>
<em>
<a href="#temporal.io/v1beta1.ElasticsearchIndices">
ElasticsearchIndices
</a>
</em>
</td>
<td>
<p>Indices holds visibility index names.</p>
</td>
</tr>
<tr>
<td>
<code>logLevel</code><br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>LogLevel defines the temporal cluster&rsquo;s es client logger level.</p>
</td>
</tr>
<tr>
<td>
<code>closeIdleConnectionsInterval</code><br>
<em>
<a href="https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#Duration">
Kubernetes meta/v1.Duration
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>CloseIdleConnectionsInterval is the max duration a connection stay open while idle.</p>
</td>
</tr>
<tr>
<td>
<code>enableSniff</code><br>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>EnableSniff enables or disables sniffer on the temporal cluster&rsquo;s es client.</p>
</td>
</tr>
<tr>
<td>
<code>enableHealthcheck</code><br>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>EnableHealthcheck enables or disables healthcheck on the temporal cluster&rsquo;s es client.</p>
</td>
</tr>
</tbody>
</table>
</div>
</div>
<h3 id="temporal.io/v1beta1.FrontendMTLSSpec">FrontendMTLSSpec
</h3>
<p>
(<em>Appears on:</em>
<a href="#temporal.io/v1beta1.MTLSSpec">MTLSSpec</a>)
</p>
<p>InternodeMTLSSpec defines parameters for the temporal encryption in transit with mTLS.</p>
<div class="md-typeset__scrollwrap">
<div class="md-typeset__table">
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Enabled defines if the operator should enable mTLS for cluster&rsquo;s public endpoints.</p>
</td>
</tr>
</tbody>
</table>
</div>
</div>
<h3 id="temporal.io/v1beta1.InternodeMTLSSpec">InternodeMTLSSpec
</h3>
<p>
(<em>Appears on:</em>
<a href="#temporal.io/v1beta1.MTLSSpec">MTLSSpec</a>)
</p>
<p>InternodeMTLSSpec defines parameters for the temporal encryption in transit with mTLS.</p>
<div class="md-typeset__scrollwrap">
<div class="md-typeset__table">
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Enabled defines if the operator should enable mTLS for network between cluster nodes.</p>
</td>
</tr>
</tbody>
</table>
</div>
</div>
<h3 id="temporal.io/v1beta1.MTLSProvider">MTLSProvider
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em>
<a href="#temporal.io/v1beta1.MTLSSpec">MTLSSpec</a>)
</p>
<p>MTLSProvider is the enum for support mTLS provider.</p>
<h3 id="temporal.io/v1beta1.MTLSSpec">MTLSSpec
</h3>
<p>
(<em>Appears on:</em>
<a href="#temporal.io/v1beta1.TemporalClusterSpec">TemporalClusterSpec</a>)
</p>
<p>MTLSSpec defines parameters for the temporal encryption in transit with mTLS.</p>
<div class="md-typeset__scrollwrap">
<div class="md-typeset__table">
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>provider</code><br>
<em>
<a href="#temporal.io/v1beta1.MTLSProvider">
MTLSProvider
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Provider defines the tool used to manage mTLS certificates.</p>
</td>
</tr>
<tr>
<td>
<code>internode</code><br>
<em>
<a href="#temporal.io/v1beta1.InternodeMTLSSpec">
InternodeMTLSSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Internode allows configuration of the internode traffic encryption.
Useless if mTLS provider is not cert-manager.</p>
</td>
</tr>
<tr>
<td>
<code>frontend</code><br>
<em>
<a href="#temporal.io/v1beta1.FrontendMTLSSpec">
FrontendMTLSSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Frontend allows configuration of the frontend&rsquo;s public endpoint traffic encryption.
Useless if mTLS provider is not cert-manager.</p>
</td>
</tr>
<tr>
<td>
<code>certificatesDuration</code><br>
<em>
<a href="#temporal.io/v1beta1.CertificatesDurationSpec">
CertificatesDurationSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>CertificatesDuration allows configuration of maximum certificates lifetime.
Useless if mTLS provider is not cert-manager.</p>
</td>
</tr>
<tr>
<td>
<code>refreshInterval</code><br>
<em>
<a href="https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#Duration">
Kubernetes meta/v1.Duration
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>RefreshInterval defines interval between refreshes of certificates in the cluster components.
Defaults to 1 hour.
Useless if mTLS provider is not cert-manager.</p>
</td>
</tr>
</tbody>
</table>
</div>
</div>
<h3 id="temporal.io/v1beta1.MetricsSpec">MetricsSpec
</h3>
<p>
(<em>Appears on:</em>
<a href="#temporal.io/v1beta1.TemporalClusterSpec">TemporalClusterSpec</a>)
</p>
<p>MetricsSpec determines parameters for configuring metrics endpoints</p>
<div class="md-typeset__scrollwrap">
<div class="md-typeset__table">
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br>
<em>
bool
</em>
</td>
<td>
<p>Enabled defines if the operator should configure metrics</p>
</td>
</tr>
<tr>
<td>
<code>prometheus</code><br>
<em>
<a href="#temporal.io/v1beta1.PrometheusSpec">
PrometheusSpec
</a>
</em>
</td>
<td>
<p>Prometheus Configuration</p>
</td>
</tr>
</tbody>
</table>
</div>
</div>
<h3 id="temporal.io/v1beta1.PrometheusSpec">PrometheusSpec
</h3>
<p>
(<em>Appears on:</em>
<a href="#temporal.io/v1beta1.MetricsSpec">MetricsSpec</a>)
</p>
<p>Prometheus is the configuration for prometheus reporter</p>
<div class="md-typeset__scrollwrap">
<div class="md-typeset__table">
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>listenAddress</code><br>
<em>
string
</em>
</td>
<td>
<p>Address for prometheus to serve metrics from.</p>
</td>
</tr>
</tbody>
</table>
</div>
</div>
<h3 id="temporal.io/v1beta1.SQLSpec">SQLSpec
</h3>
<p>
(<em>Appears on:</em>
<a href="#temporal.io/v1beta1.DatastoreSpec">DatastoreSpec</a>)
</p>
<p>SQLSpec contains SQL datastore connections specifications.</p>
<div class="md-typeset__scrollwrap">
<div class="md-typeset__table">
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>user</code><br>
<em>
string
</em>
</td>
<td>
<p>User is the username to be used for the connection.</p>
</td>
</tr>
<tr>
<td>
<code>pluginName</code><br>
<em>
string
</em>
</td>
<td>
<p>PluginName is the name of SQL plugin.</p>
</td>
</tr>
<tr>
<td>
<code>databaseName</code><br>
<em>
string
</em>
</td>
<td>
<p>DatabaseName is the name of SQL database to connect to.</p>
</td>
</tr>
<tr>
<td>
<code>connectAddr</code><br>
<em>
string
</em>
</td>
<td>
<p>ConnectAddr is the remote addr of the database.</p>
</td>
</tr>
<tr>
<td>
<code>connectProtocol</code><br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>ConnectProtocol is the protocol that goes with the ConnectAddr.</p>
</td>
</tr>
<tr>
<td>
<code>connectAttributes</code><br>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>ConnectAttributes is a set of key-value attributes to be sent as part of connect data_source_name url</p>
</td>
</tr>
<tr>
<td>
<code>maxConns</code><br>
<em>
int
</em>
</td>
<td>
<em>(Optional)</em>
<p>MaxConns the max number of connections to this datastore.</p>
</td>
</tr>
<tr>
<td>
<code>maxIdleConns</code><br>
<em>
int
</em>
</td>
<td>
<em>(Optional)</em>
<p>MaxIdleConns is the max number of idle connections to this datastore.</p>
</td>
</tr>
<tr>
<td>
<code>maxConnLifetime</code><br>
<em>
<a href="https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#Duration">
Kubernetes meta/v1.Duration
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>MaxConnLifetime is the maximum time a connection can be alive</p>
</td>
</tr>
<tr>
<td>
<code>taskScanPartitions</code><br>
<em>
int
</em>
</td>
<td>
<em>(Optional)</em>
<p>TaskScanPartitions is the number of partitions to sequentially scan during ListTaskQueue operations.</p>
</td>
</tr>
</tbody>
</table>
</div>
</div>
<h3 id="temporal.io/v1beta1.SecretKeyReference">SecretKeyReference
</h3>
<p>
(<em>Appears on:</em>
<a href="#temporal.io/v1beta1.DatastoreSpec">DatastoreSpec</a>, 
<a href="#temporal.io/v1beta1.DatastoreTLSSpec">DatastoreTLSSpec</a>)
</p>
<p>SecretKeyReference contains enough information to locate the referenced Kubernetes Secret object in the same
namespace.</p>
<div class="md-typeset__scrollwrap">
<div class="md-typeset__table">
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br>
<em>
string
</em>
</td>
<td>
<p>Name of the Secret.</p>
</td>
</tr>
<tr>
<td>
<code>key</code><br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Key in the Secret.</p>
</td>
</tr>
</tbody>
</table>
</div>
</div>
<h3 id="temporal.io/v1beta1.ServiceSpec">ServiceSpec
</h3>
<p>
(<em>Appears on:</em>
<a href="#temporal.io/v1beta1.ServicesSpec">ServicesSpec</a>)
</p>
<p>ServiceSpec contains a temporal service specifications.</p>
<div class="md-typeset__scrollwrap">
<div class="md-typeset__table">
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>port</code><br>
<em>
int
</em>
</td>
<td>
<em>(Optional)</em>
<p>Port defines a custom gRPC port for the service.
Default values are:
7233 for Frontend service
7234 for History service
7235 for Matching service
7239 for Worker service</p>
</td>
</tr>
<tr>
<td>
<code>membershipPort</code><br>
<em>
int
</em>
</td>
<td>
<em>(Optional)</em>
<p>Port defines a custom membership port for the service.
Default values are:
6933 for Frontend service
6934 for History service
6935 for Matching service
6939 for Worker service</p>
</td>
</tr>
<tr>
<td>
<code>replicas</code><br>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Number of desired replicas for the service. Default to 1.</p>
</td>
</tr>
</tbody>
</table>
</div>
</div>
<h3 id="temporal.io/v1beta1.ServiceStatus">ServiceStatus
</h3>
<p>
(<em>Appears on:</em>
<a href="#temporal.io/v1beta1.TemporalClusterStatus">TemporalClusterStatus</a>)
</p>
<p>ServiceStatus reports a service status.</p>
<div class="md-typeset__scrollwrap">
<div class="md-typeset__table">
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br>
<em>
string
</em>
</td>
<td>
<p>Name of the temporal service.</p>
</td>
</tr>
<tr>
<td>
<code>version</code><br>
<em>
string
</em>
</td>
<td>
<p>Current observed version of the service.</p>
</td>
</tr>
<tr>
<td>
<code>ready</code><br>
<em>
bool
</em>
</td>
<td>
<p>Ready defines if the service is ready.</p>
</td>
</tr>
</tbody>
</table>
</div>
</div>
<h3 id="temporal.io/v1beta1.ServicesSpec">ServicesSpec
</h3>
<p>
(<em>Appears on:</em>
<a href="#temporal.io/v1beta1.TemporalClusterSpec">TemporalClusterSpec</a>)
</p>
<p>ServicesSpec contains all temporal services specifications.</p>
<div class="md-typeset__scrollwrap">
<div class="md-typeset__table">
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>frontend</code><br>
<em>
<a href="#temporal.io/v1beta1.ServiceSpec">
ServiceSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Frontend service custom specifications.</p>
</td>
</tr>
<tr>
<td>
<code>history</code><br>
<em>
<a href="#temporal.io/v1beta1.ServiceSpec">
ServiceSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>History service custom specifications.</p>
</td>
</tr>
<tr>
<td>
<code>matching</code><br>
<em>
<a href="#temporal.io/v1beta1.ServiceSpec">
ServiceSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Matching service custom specifications.</p>
</td>
</tr>
<tr>
<td>
<code>worker</code><br>
<em>
<a href="#temporal.io/v1beta1.ServiceSpec">
ServiceSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Worker service custom specifications.</p>
</td>
</tr>
</tbody>
</table>
</div>
</div>
<h3 id="temporal.io/v1beta1.TemporalAdminToolsSpec">TemporalAdminToolsSpec
</h3>
<p>
(<em>Appears on:</em>
<a href="#temporal.io/v1beta1.TemporalClusterSpec">TemporalClusterSpec</a>)
</p>
<p>TemporalUISpec defines parameters for the temporal admin tools within a Temporal cluster deployment.
Note that deployed admin tools version is the same as the cluster&rsquo;s version.</p>
<div class="md-typeset__scrollwrap">
<div class="md-typeset__table">
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Enabled defines if the operator should deploy the admin tools alongside the cluster.</p>
</td>
</tr>
<tr>
<td>
<code>image</code><br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Image defines the temporal admin tools docker image the instance should run.</p>
</td>
</tr>
</tbody>
</table>
</div>
</div>
<h3 id="temporal.io/v1beta1.TemporalWorkerProcess">TemporalWorkerProcess
</h3>
<p>TemporalWorkerProcess is the Schema for the temporalworkerprocesses API</p>
<div class="md-typeset__scrollwrap">
<div class="md-typeset__table">
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metadata</code><br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br>
<em>
<a href="#temporal.io/v1beta1.TemporalWorkerProcessSpec">
TemporalWorkerProcessSpec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>foo</code><br>
<em>
string
</em>
</td>
<td>
<p>Foo is an example field of TemporalWorkerProcess. Edit temporalworkerprocess_types.go to remove/update</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br>
<em>
<a href="#temporal.io/v1beta1.TemporalWorkerProcessStatus">
TemporalWorkerProcessStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
</div>
</div>
<h3 id="temporal.io/v1beta1.TemporalWorkerProcessSpec">TemporalWorkerProcessSpec
</h3>
<p>
(<em>Appears on:</em>
<a href="#temporal.io/v1beta1.TemporalWorkerProcess">TemporalWorkerProcess</a>)
</p>
<p>TemporalWorkerProcessSpec defines the desired state of TemporalWorkerProcess</p>
<div class="md-typeset__scrollwrap">
<div class="md-typeset__table">
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>foo</code><br>
<em>
string
</em>
</td>
<td>
<p>Foo is an example field of TemporalWorkerProcess. Edit temporalworkerprocess_types.go to remove/update</p>
</td>
</tr>
</tbody>
</table>
</div>
</div>
<h3 id="temporal.io/v1beta1.TemporalWorkerProcessStatus">TemporalWorkerProcessStatus
</h3>
<p>
(<em>Appears on:</em>
<a href="#temporal.io/v1beta1.TemporalWorkerProcess">TemporalWorkerProcess</a>)
</p>
<p>TemporalWorkerProcessStatus defines the observed state of TemporalWorkerProcess</p>
<h3 id="temporal.io/v1beta1.TemporalClusterClient">TemporalClusterClient
</h3>
<p>A TemporalClusterClient creates a new mTLS client in the targeted temporal cluster.</p>
<div class="md-typeset__scrollwrap">
<div class="md-typeset__table">
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metadata</code><br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br>
<em>
<a href="#temporal.io/v1beta1.TemporalClusterClientSpec">
TemporalClusterClientSpec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>clusterRef</code><br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#localobjectreference-v1-core">
Kubernetes core/v1.LocalObjectReference
</a>
</em>
</td>
<td>
<p>Reference to the temporal cluster the client will get access to.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br>
<em>
<a href="#temporal.io/v1beta1.TemporalClusterClientStatus">
TemporalClusterClientStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
</div>
</div>
<h3 id="temporal.io/v1beta1.TemporalClusterClientSpec">TemporalClusterClientSpec
</h3>
<p>
(<em>Appears on:</em>
<a href="#temporal.io/v1beta1.TemporalClusterClient">TemporalClusterClient</a>)
</p>
<p>TemporalClusterClientSpec defines the desired state of ClusterClient</p>
<div class="md-typeset__scrollwrap">
<div class="md-typeset__table">
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>clusterRef</code><br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#localobjectreference-v1-core">
Kubernetes core/v1.LocalObjectReference
</a>
</em>
</td>
<td>
<p>Reference to the temporal cluster the client will get access to.</p>
</td>
</tr>
</tbody>
</table>
</div>
</div>
<h3 id="temporal.io/v1beta1.TemporalClusterClientStatus">TemporalClusterClientStatus
</h3>
<p>
(<em>Appears on:</em>
<a href="#temporal.io/v1beta1.TemporalClusterClient">TemporalClusterClient</a>)
</p>
<p>TemporalClusterClientStatus defines the observed state of ClusterClient</p>
<div class="md-typeset__scrollwrap">
<div class="md-typeset__table">
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>serverName</code><br>
<em>
string
</em>
</td>
<td>
<p>ServerName is the hostname returned by the certificate.</p>
</td>
</tr>
<tr>
<td>
<code>secretRef</code><br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#localobjectreference-v1-core">
Kubernetes core/v1.LocalObjectReference
</a>
</em>
</td>
<td>
<p>Reference to the Kubernetes Secret containing the certificate for the client.</p>
</td>
</tr>
</tbody>
</table>
</div>
</div>
<h3 id="temporal.io/v1beta1.TemporalClusterSpec">TemporalClusterSpec
</h3>
<p>
(<em>Appears on:</em>
<a href="#temporal.io/v1beta1.TemporalCluster">TemporalCluster</a>)
</p>
<p>TemporalClusterSpec defines the desired state of Cluster.</p>
<div class="md-typeset__scrollwrap">
<div class="md-typeset__table">
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>image</code><br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Image defines the temporal server docker image the cluster should use for each services.</p>
</td>
</tr>
<tr>
<td>
<code>version</code><br>
<em>
github.com/alexandrevilain/temporal-operator/pkg/version.Version
</em>
</td>
<td>
<em>(Optional)</em>
<p>Version defines the temporal version the cluster to be deployed.
This version impacts the underlying persistence schemas versions.</p>
</td>
</tr>
<tr>
<td>
<code>jobTtlSecondsAfterFinished</code><br>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>JobTtlSecondsAfterFinished is amount of time to keep job pods after jobs are completed.
Defaults to 300 seconds.</p>
</td>
</tr>
<tr>
<td>
<code>numHistoryShards</code><br>
<em>
int32
</em>
</td>
<td>
<p>NumHistoryShards is the desired number of history shards.
This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>services</code><br>
<em>
<a href="#temporal.io/v1beta1.ServicesSpec">
ServicesSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Services allows customizations for for each temporal services deployment.</p>
</td>
</tr>
<tr>
<td>
<code>persistence</code><br>
<em>
<a href="#temporal.io/v1beta1.TemporalPersistenceSpec">
TemporalPersistenceSpec
</a>
</em>
</td>
<td>
<p>Persistence defines temporal persistence configuration.</p>
</td>
</tr>
<tr>
<td>
<code>imagePullSecrets</code><br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#localobjectreference-v1-core">
[]Kubernetes core/v1.LocalObjectReference
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>An optional list of references to secrets in the same namespace
to use for pulling temporal images from registries.</p>
</td>
</tr>
<tr>
<td>
<code>ui</code><br>
<em>
<a href="#temporal.io/v1beta1.TemporalUISpec">
TemporalUISpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>UI allows configuration of the optional temporal web ui deployed alongside the cluster.</p>
</td>
</tr>
<tr>
<td>
<code>admintools</code><br>
<em>
<a href="#temporal.io/v1beta1.TemporalAdminToolsSpec">
TemporalAdminToolsSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>AdminTools allows configuration of the optional admin tool pod deployed alongside the cluster.</p>
</td>
</tr>
<tr>
<td>
<code>mTLS</code><br>
<em>
<a href="#temporal.io/v1beta1.MTLSSpec">
MTLSSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>MTLS allows configuration of the network traffic encryption for the cluster.</p>
</td>
</tr>
<tr>
<td>
<code>metrics</code><br>
<em>
<a href="#temporal.io/v1beta1.MetricsSpec">
MetricsSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Metrics allows configuration of scraping endpoints for stats. prometheus or m3.</p>
</td>
</tr>
</tbody>
</table>
</div>
</div>
<h3 id="temporal.io/v1beta1.TemporalClusterStatus">TemporalClusterStatus
</h3>
<p>
(<em>Appears on:</em>
<a href="#temporal.io/v1beta1.TemporalCluster">TemporalCluster</a>)
</p>
<p>TemporalClusterStatus defines the observed state of Cluster.</p>
<div class="md-typeset__scrollwrap">
<div class="md-typeset__table">
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>version</code><br>
<em>
string
</em>
</td>
<td>
<p>Version holds the current temporal version.</p>
</td>
</tr>
<tr>
<td>
<code>services</code><br>
<em>
<a href="#temporal.io/v1beta1.ServiceStatus">
[]ServiceStatus
</a>
</em>
</td>
<td>
<p>Services holds all services statuses.</p>
</td>
</tr>
<tr>
<td>
<code>persistence</code><br>
<em>
<a href="#temporal.io/v1beta1.TemporalPersistenceStatus">
TemporalPersistenceStatus
</a>
</em>
</td>
<td>
<p>Persistence holds all datastores statuses.</p>
</td>
</tr>
<tr>
<td>
<code>conditions</code><br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#condition-v1-meta">
[]Kubernetes meta/v1.Condition
</a>
</em>
</td>
<td>
<p>Conditions represent the latest available observations of the Cluster state.</p>
</td>
</tr>
</tbody>
</table>
</div>
</div>
<h3 id="temporal.io/v1beta1.TemporalNamespace">TemporalNamespace
</h3>
<p>A TemporalNamespace creates a namespace in the targeted temporal cluster.</p>
<div class="md-typeset__scrollwrap">
<div class="md-typeset__table">
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metadata</code><br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br>
<em>
<a href="#temporal.io/v1beta1.TemporalNamespaceSpec">
TemporalNamespaceSpec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>clusterRef</code><br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#localobjectreference-v1-core">
Kubernetes core/v1.LocalObjectReference
</a>
</em>
</td>
<td>
<p>Reference to the temporal cluster the namespace will be created.</p>
</td>
</tr>
<tr>
<td>
<code>description</code><br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Namespace description.</p>
</td>
</tr>
<tr>
<td>
<code>ownerEmail</code><br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Namespace owner email.</p>
</td>
</tr>
<tr>
<td>
<code>retentionPeriod</code><br>
<em>
<a href="https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#Duration">
Kubernetes meta/v1.Duration
</a>
</em>
</td>
<td>
<p>RetentionPeriod to apply on closed workflow executions.</p>
</td>
</tr>
<tr>
<td>
<code>data</code><br>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Data is a key-value map for any customized purpose.</p>
</td>
</tr>
<tr>
<td>
<code>securityToken</code><br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
<tr>
<td>
<code>isGlobalNamespace</code><br>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>IsGlobalNamespace defines whether the namespace is a global namespace.</p>
</td>
</tr>
<tr>
<td>
<code>clusters</code><br>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>List of clusters names to which the namespace can fail over.
Only applicable if the namespace is a global namespace.</p>
</td>
</tr>
<tr>
<td>
<code>activeClusterName</code><br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>The name of active Temporal Cluster.
Only applicable if the namespace is a global namespace.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br>
<em>
<a href="#temporal.io/v1beta1.TemporalNamespaceStatus">
TemporalNamespaceStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
</div>
</div>
<h3 id="temporal.io/v1beta1.TemporalNamespaceSpec">TemporalNamespaceSpec
</h3>
<p>
(<em>Appears on:</em>
<a href="#temporal.io/v1beta1.TemporalNamespace">TemporalNamespace</a>)
</p>
<p>TemporalNamespaceSpec defines the desired state of Namespace</p>
<div class="md-typeset__scrollwrap">
<div class="md-typeset__table">
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>clusterRef</code><br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#localobjectreference-v1-core">
Kubernetes core/v1.LocalObjectReference
</a>
</em>
</td>
<td>
<p>Reference to the temporal cluster the namespace will be created.</p>
</td>
</tr>
<tr>
<td>
<code>description</code><br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Namespace description.</p>
</td>
</tr>
<tr>
<td>
<code>ownerEmail</code><br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Namespace owner email.</p>
</td>
</tr>
<tr>
<td>
<code>retentionPeriod</code><br>
<em>
<a href="https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#Duration">
Kubernetes meta/v1.Duration
</a>
</em>
</td>
<td>
<p>RetentionPeriod to apply on closed workflow executions.</p>
</td>
</tr>
<tr>
<td>
<code>data</code><br>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Data is a key-value map for any customized purpose.</p>
</td>
</tr>
<tr>
<td>
<code>securityToken</code><br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
<tr>
<td>
<code>isGlobalNamespace</code><br>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>IsGlobalNamespace defines whether the namespace is a global namespace.</p>
</td>
</tr>
<tr>
<td>
<code>clusters</code><br>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>List of clusters names to which the namespace can fail over.
Only applicable if the namespace is a global namespace.</p>
</td>
</tr>
<tr>
<td>
<code>activeClusterName</code><br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>The name of active Temporal Cluster.
Only applicable if the namespace is a global namespace.</p>
</td>
</tr>
</tbody>
</table>
</div>
</div>
<h3 id="temporal.io/v1beta1.TemporalNamespaceStatus">TemporalNamespaceStatus
</h3>
<p>
(<em>Appears on:</em>
<a href="#temporal.io/v1beta1.TemporalNamespace">TemporalNamespace</a>)
</p>
<p>TemporalNamespaceStatus defines the observed state of Namespace</p>
<div class="md-typeset__scrollwrap">
<div class="md-typeset__table">
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>conditions</code><br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#condition-v1-meta">
[]Kubernetes meta/v1.Condition
</a>
</em>
</td>
<td>
<p>Conditions represent the latest available observations of the Namespace state.</p>
</td>
</tr>
</tbody>
</table>
</div>
</div>
<h3 id="temporal.io/v1beta1.TemporalPersistenceSpec">TemporalPersistenceSpec
</h3>
<p>
(<em>Appears on:</em>
<a href="#temporal.io/v1beta1.TemporalClusterSpec">TemporalClusterSpec</a>)
</p>
<p>TemporalPersistenceSpec contains temporal persistence specifications.</p>
<div class="md-typeset__scrollwrap">
<div class="md-typeset__table">
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>defaultStore</code><br>
<em>
<a href="#temporal.io/v1beta1.DatastoreSpec">
DatastoreSpec
</a>
</em>
</td>
<td>
<p>DefaultStore holds the default datastore specs.</p>
</td>
</tr>
<tr>
<td>
<code>visibilityStore</code><br>
<em>
<a href="#temporal.io/v1beta1.DatastoreSpec">
DatastoreSpec
</a>
</em>
</td>
<td>
<p>VisibilityStore holds the visibility datastore specs.</p>
</td>
</tr>
<tr>
<td>
<code>advancedVisibilityStore</code><br>
<em>
<a href="#temporal.io/v1beta1.DatastoreSpec">
DatastoreSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>AdvancedVisibilityStore holds the avanced visibility datastore specs.</p>
</td>
</tr>
</tbody>
</table>
</div>
</div>
<h3 id="temporal.io/v1beta1.TemporalPersistenceStatus">TemporalPersistenceStatus
</h3>
<p>
(<em>Appears on:</em>
<a href="#temporal.io/v1beta1.TemporalClusterStatus">TemporalClusterStatus</a>)
</p>
<p>TemporalPersistenceStatus contains temporal persistence status.</p>
<div class="md-typeset__scrollwrap">
<div class="md-typeset__table">
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>defaultStore</code><br>
<em>
<a href="#temporal.io/v1beta1.DatastoreStatus">
DatastoreStatus
</a>
</em>
</td>
<td>
<p>DefaultStore holds the default datastore status.</p>
</td>
</tr>
<tr>
<td>
<code>visibilityStore</code><br>
<em>
<a href="#temporal.io/v1beta1.DatastoreStatus">
DatastoreStatus
</a>
</em>
</td>
<td>
<p>VisibilityStore holds the visibility datastore status.</p>
</td>
</tr>
<tr>
<td>
<code>advancedVisibilityStore</code><br>
<em>
<a href="#temporal.io/v1beta1.DatastoreStatus">
DatastoreStatus
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>AdvancedVisibilityStore holds the avanced visibility datastore status.</p>
</td>
</tr>
</tbody>
</table>
</div>
</div>
<h3 id="temporal.io/v1beta1.TemporalUIIngressSpec">TemporalUIIngressSpec
</h3>
<p>
(<em>Appears on:</em>
<a href="#temporal.io/v1beta1.TemporalUISpec">TemporalUISpec</a>)
</p>
<p>TemporalUIIngressSpec contains all configurations options for the UI ingress.</p>
<div class="md-typeset__scrollwrap">
<div class="md-typeset__table">
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>annotations</code><br>
<em>
map[string]string
</em>
</td>
<td>
<p>Annotations allows custom annotations on the ingress ressource.</p>
</td>
</tr>
<tr>
<td>
<code>ingressClassName</code><br>
<em>
string
</em>
</td>
<td>
<p>IngressClassName is the name of the IngressClass the deployed ingress resource should use.</p>
</td>
</tr>
<tr>
<td>
<code>hosts</code><br>
<em>
[]string
</em>
</td>
<td>
<p>Host is the list of host the ingress should use.</p>
</td>
</tr>
<tr>
<td>
<code>tls</code><br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#ingresstls-v1-networking">
[]Kubernetes networking/v1.IngressTLS
</a>
</em>
</td>
<td>
<p>TLS configuration.</p>
</td>
</tr>
</tbody>
</table>
</div>
</div>
<h3 id="temporal.io/v1beta1.TemporalUISpec">TemporalUISpec
</h3>
<p>
(<em>Appears on:</em>
<a href="#temporal.io/v1beta1.TemporalClusterSpec">TemporalClusterSpec</a>)
</p>
<p>TemporalUISpec defines parameters for the temporal UI within a Temporal cluster deployment.</p>
<div class="md-typeset__scrollwrap">
<div class="md-typeset__table">
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Enabled defines if the operator should deploy the web ui alongside the cluster.</p>
</td>
</tr>
<tr>
<td>
<code>version</code><br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Version defines the temporal ui version the instance should run.</p>
</td>
</tr>
<tr>
<td>
<code>image</code><br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Image defines the temporal ui docker image the instance should run.</p>
</td>
</tr>
<tr>
<td>
<code>ingress</code><br>
<em>
<a href="#temporal.io/v1beta1.TemporalUIIngressSpec">
TemporalUIIngressSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Ingress is an optional ingress configuration for the UI.
If lived empty, no ingress configuration will be created and the UI will only by available trough ClusterIP service.</p>
</td>
</tr>
</tbody>
</table>
</div>
</div>
<div class="admonition note">
<p class="last">This page was automatically generated with <code>gen-crd-api-reference-docs</code></p>
</div>
