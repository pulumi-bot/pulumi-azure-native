// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190301preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The cluster resource
type Cluster struct {
	pulumi.CustomResourceState

	// The list of add-on features to enable in the cluster.
	AddOnFeatures pulumi.StringArrayOutput `pulumi:"addOnFeatures"`
	// The Service Fabric runtime versions available for this cluster.
	AvailableClusterVersions ClusterVersionDetailsResponseArrayOutput `pulumi:"availableClusterVersions"`
	// The AAD authentication settings of the cluster.
	AzureActiveDirectory AzureActiveDirectoryResponsePtrOutput `pulumi:"azureActiveDirectory"`
	// The certificate to use for securing the cluster. The certificate provided will be used for node to node security within the cluster, SSL certificate for cluster management endpoint and default admin client.
	Certificate CertificateDescriptionResponsePtrOutput `pulumi:"certificate"`
	// Describes a list of server certificates referenced by common name that are used to secure the cluster.
	CertificateCommonNames ServerCertificateCommonNamesResponsePtrOutput `pulumi:"certificateCommonNames"`
	// The list of client certificates referenced by common name that are allowed to manage the cluster.
	ClientCertificateCommonNames ClientCertificateCommonNameResponseArrayOutput `pulumi:"clientCertificateCommonNames"`
	// The list of client certificates referenced by thumbprint that are allowed to manage the cluster.
	ClientCertificateThumbprints ClientCertificateThumbprintResponseArrayOutput `pulumi:"clientCertificateThumbprints"`
	// The Service Fabric runtime version of the cluster. This property can only by set the user when **upgradeMode** is set to 'Manual'. To get list of available Service Fabric versions for new clusters use [ClusterVersion API](./ClusterVersion.md). To get the list of available version for existing clusters use **availableClusterVersions**.
	ClusterCodeVersion pulumi.StringPtrOutput `pulumi:"clusterCodeVersion"`
	// The Azure Resource Provider endpoint. A system service in the cluster connects to this  endpoint.
	ClusterEndpoint pulumi.StringOutput `pulumi:"clusterEndpoint"`
	// A service generated unique identifier for the cluster resource.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// The current state of the cluster.
	//
	//   - WaitingForNodes - Indicates that the cluster resource is created and the resource provider is waiting for Service Fabric VM extension to boot up and report to it.
	//   - Deploying - Indicates that the Service Fabric runtime is being installed on the VMs. Cluster resource will be in this state until the cluster boots up and system services are up.
	//   - BaselineUpgrade - Indicates that the cluster is upgrading to establishes the cluster version. This upgrade is automatically initiated when the cluster boots up for the first time.
	//   - UpdatingUserConfiguration - Indicates that the cluster is being upgraded with the user provided configuration.
	//   - UpdatingUserCertificate - Indicates that the cluster is being upgraded with the user provided certificate.
	//   - UpdatingInfrastructure - Indicates that the cluster is being upgraded with the latest Service Fabric runtime version. This happens only when the **upgradeMode** is set to 'Automatic'.
	//   - EnforcingClusterVersion - Indicates that cluster is on a different version than expected and the cluster is being upgraded to the expected version.
	//   - UpgradeServiceUnreachable - Indicates that the system service in the cluster is no longer polling the Resource Provider. Clusters in this state cannot be managed by the Resource Provider.
	//   - AutoScale - Indicates that the ReliabilityLevel of the cluster is being adjusted.
	//   - Ready - Indicates that the cluster is in a stable state.
	ClusterState pulumi.StringOutput `pulumi:"clusterState"`
	// The storage account information for storing Service Fabric diagnostic logs.
	DiagnosticsStorageAccountConfig DiagnosticsStorageAccountConfigResponsePtrOutput `pulumi:"diagnosticsStorageAccountConfig"`
	// Azure resource etag.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Indicates if the event store service is enabled.
	EventStoreServiceEnabled pulumi.BoolPtrOutput `pulumi:"eventStoreServiceEnabled"`
	// The list of custom fabric settings to configure the cluster.
	FabricSettings SettingsSectionDescriptionResponseArrayOutput `pulumi:"fabricSettings"`
	// Azure resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// The http management endpoint of the cluster.
	ManagementEndpoint pulumi.StringOutput `pulumi:"managementEndpoint"`
	// Azure resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The list of node types in the cluster.
	NodeTypes NodeTypeDescriptionResponseArrayOutput `pulumi:"nodeTypes"`
	// The provisioning state of the cluster resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The reliability level sets the replica set size of system services. Learn about [ReliabilityLevel](https://docs.microsoft.com/azure/service-fabric/service-fabric-cluster-capacity).
	//
	//   - None - Run the System services with a target replica set count of 1. This should only be used for test clusters.
	//   - Bronze - Run the System services with a target replica set count of 3. This should only be used for test clusters.
	//   - Silver - Run the System services with a target replica set count of 5.
	//   - Gold - Run the System services with a target replica set count of 7.
	//   - Platinum - Run the System services with a target replica set count of 9.
	ReliabilityLevel pulumi.StringPtrOutput `pulumi:"reliabilityLevel"`
	// The server certificate used by reverse proxy.
	ReverseProxyCertificate CertificateDescriptionResponsePtrOutput `pulumi:"reverseProxyCertificate"`
	// Describes a list of server certificates referenced by common name that are used to secure the cluster.
	ReverseProxyCertificateCommonNames ServerCertificateCommonNamesResponsePtrOutput `pulumi:"reverseProxyCertificateCommonNames"`
	// Azure resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Azure resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The policy to use when upgrading the cluster.
	UpgradeDescription ClusterUpgradePolicyResponsePtrOutput `pulumi:"upgradeDescription"`
	// The upgrade mode of the cluster when new Service Fabric runtime version is available.
	//
	//   - Automatic - The cluster will be automatically upgraded to the latest Service Fabric runtime version as soon as it is available.
	//   - Manual - The cluster will not be automatically upgraded to the latest Service Fabric runtime version. The cluster is upgraded by setting the **clusterCodeVersion** property in the cluster resource.
	UpgradeMode pulumi.StringPtrOutput `pulumi:"upgradeMode"`
	// The VM image VMSS has been configured with. Generic names such as Windows or Linux can be used.
	VmImage pulumi.StringPtrOutput `pulumi:"vmImage"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil || args.ClusterName == nil {
		return nil, errors.New("missing required argument 'ClusterName'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ManagementEndpoint == nil {
		return nil, errors.New("missing required argument 'ManagementEndpoint'")
	}
	if args == nil || args.NodeTypes == nil {
		return nil, errors.New("missing required argument 'NodeTypes'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ClusterArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:servicefabric/latest:Cluster"),
		},
		{
			Type: pulumi.String("azurerm:servicefabric/v20160901:Cluster"),
		},
		{
			Type: pulumi.String("azurerm:servicefabric/v20170701preview:Cluster"),
		},
		{
			Type: pulumi.String("azurerm:servicefabric/v20180201:Cluster"),
		},
		{
			Type: pulumi.String("azurerm:servicefabric/v20190301:Cluster"),
		},
		{
			Type: pulumi.String("azurerm:servicefabric/v20190601preview:Cluster"),
		},
		{
			Type: pulumi.String("azurerm:servicefabric/v20191101preview:Cluster"),
		},
		{
			Type: pulumi.String("azurerm:servicefabric/v20200301:Cluster"),
		},
	})
	opts = append(opts, aliases)
	var resource Cluster
	err := ctx.RegisterResource("azurerm:servicefabric/v20190301preview:Cluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCluster gets an existing Cluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterState, opts ...pulumi.ResourceOption) (*Cluster, error) {
	var resource Cluster
	err := ctx.ReadResource("azurerm:servicefabric/v20190301preview:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
	// The list of add-on features to enable in the cluster.
	AddOnFeatures []string `pulumi:"addOnFeatures"`
	// The Service Fabric runtime versions available for this cluster.
	AvailableClusterVersions []ClusterVersionDetailsResponse `pulumi:"availableClusterVersions"`
	// The AAD authentication settings of the cluster.
	AzureActiveDirectory *AzureActiveDirectoryResponse `pulumi:"azureActiveDirectory"`
	// The certificate to use for securing the cluster. The certificate provided will be used for node to node security within the cluster, SSL certificate for cluster management endpoint and default admin client.
	Certificate *CertificateDescriptionResponse `pulumi:"certificate"`
	// Describes a list of server certificates referenced by common name that are used to secure the cluster.
	CertificateCommonNames *ServerCertificateCommonNamesResponse `pulumi:"certificateCommonNames"`
	// The list of client certificates referenced by common name that are allowed to manage the cluster.
	ClientCertificateCommonNames []ClientCertificateCommonNameResponse `pulumi:"clientCertificateCommonNames"`
	// The list of client certificates referenced by thumbprint that are allowed to manage the cluster.
	ClientCertificateThumbprints []ClientCertificateThumbprintResponse `pulumi:"clientCertificateThumbprints"`
	// The Service Fabric runtime version of the cluster. This property can only by set the user when **upgradeMode** is set to 'Manual'. To get list of available Service Fabric versions for new clusters use [ClusterVersion API](./ClusterVersion.md). To get the list of available version for existing clusters use **availableClusterVersions**.
	ClusterCodeVersion *string `pulumi:"clusterCodeVersion"`
	// The Azure Resource Provider endpoint. A system service in the cluster connects to this  endpoint.
	ClusterEndpoint *string `pulumi:"clusterEndpoint"`
	// A service generated unique identifier for the cluster resource.
	ClusterId *string `pulumi:"clusterId"`
	// The current state of the cluster.
	//
	//   - WaitingForNodes - Indicates that the cluster resource is created and the resource provider is waiting for Service Fabric VM extension to boot up and report to it.
	//   - Deploying - Indicates that the Service Fabric runtime is being installed on the VMs. Cluster resource will be in this state until the cluster boots up and system services are up.
	//   - BaselineUpgrade - Indicates that the cluster is upgrading to establishes the cluster version. This upgrade is automatically initiated when the cluster boots up for the first time.
	//   - UpdatingUserConfiguration - Indicates that the cluster is being upgraded with the user provided configuration.
	//   - UpdatingUserCertificate - Indicates that the cluster is being upgraded with the user provided certificate.
	//   - UpdatingInfrastructure - Indicates that the cluster is being upgraded with the latest Service Fabric runtime version. This happens only when the **upgradeMode** is set to 'Automatic'.
	//   - EnforcingClusterVersion - Indicates that cluster is on a different version than expected and the cluster is being upgraded to the expected version.
	//   - UpgradeServiceUnreachable - Indicates that the system service in the cluster is no longer polling the Resource Provider. Clusters in this state cannot be managed by the Resource Provider.
	//   - AutoScale - Indicates that the ReliabilityLevel of the cluster is being adjusted.
	//   - Ready - Indicates that the cluster is in a stable state.
	ClusterState *string `pulumi:"clusterState"`
	// The storage account information for storing Service Fabric diagnostic logs.
	DiagnosticsStorageAccountConfig *DiagnosticsStorageAccountConfigResponse `pulumi:"diagnosticsStorageAccountConfig"`
	// Azure resource etag.
	Etag *string `pulumi:"etag"`
	// Indicates if the event store service is enabled.
	EventStoreServiceEnabled *bool `pulumi:"eventStoreServiceEnabled"`
	// The list of custom fabric settings to configure the cluster.
	FabricSettings []SettingsSectionDescriptionResponse `pulumi:"fabricSettings"`
	// Azure resource location.
	Location *string `pulumi:"location"`
	// The http management endpoint of the cluster.
	ManagementEndpoint *string `pulumi:"managementEndpoint"`
	// Azure resource name.
	Name *string `pulumi:"name"`
	// The list of node types in the cluster.
	NodeTypes []NodeTypeDescriptionResponse `pulumi:"nodeTypes"`
	// The provisioning state of the cluster resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The reliability level sets the replica set size of system services. Learn about [ReliabilityLevel](https://docs.microsoft.com/azure/service-fabric/service-fabric-cluster-capacity).
	//
	//   - None - Run the System services with a target replica set count of 1. This should only be used for test clusters.
	//   - Bronze - Run the System services with a target replica set count of 3. This should only be used for test clusters.
	//   - Silver - Run the System services with a target replica set count of 5.
	//   - Gold - Run the System services with a target replica set count of 7.
	//   - Platinum - Run the System services with a target replica set count of 9.
	ReliabilityLevel *string `pulumi:"reliabilityLevel"`
	// The server certificate used by reverse proxy.
	ReverseProxyCertificate *CertificateDescriptionResponse `pulumi:"reverseProxyCertificate"`
	// Describes a list of server certificates referenced by common name that are used to secure the cluster.
	ReverseProxyCertificateCommonNames *ServerCertificateCommonNamesResponse `pulumi:"reverseProxyCertificateCommonNames"`
	// Azure resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Azure resource type.
	Type *string `pulumi:"type"`
	// The policy to use when upgrading the cluster.
	UpgradeDescription *ClusterUpgradePolicyResponse `pulumi:"upgradeDescription"`
	// The upgrade mode of the cluster when new Service Fabric runtime version is available.
	//
	//   - Automatic - The cluster will be automatically upgraded to the latest Service Fabric runtime version as soon as it is available.
	//   - Manual - The cluster will not be automatically upgraded to the latest Service Fabric runtime version. The cluster is upgraded by setting the **clusterCodeVersion** property in the cluster resource.
	UpgradeMode *string `pulumi:"upgradeMode"`
	// The VM image VMSS has been configured with. Generic names such as Windows or Linux can be used.
	VmImage *string `pulumi:"vmImage"`
}

type ClusterState struct {
	// The list of add-on features to enable in the cluster.
	AddOnFeatures pulumi.StringArrayInput
	// The Service Fabric runtime versions available for this cluster.
	AvailableClusterVersions ClusterVersionDetailsResponseArrayInput
	// The AAD authentication settings of the cluster.
	AzureActiveDirectory AzureActiveDirectoryResponsePtrInput
	// The certificate to use for securing the cluster. The certificate provided will be used for node to node security within the cluster, SSL certificate for cluster management endpoint and default admin client.
	Certificate CertificateDescriptionResponsePtrInput
	// Describes a list of server certificates referenced by common name that are used to secure the cluster.
	CertificateCommonNames ServerCertificateCommonNamesResponsePtrInput
	// The list of client certificates referenced by common name that are allowed to manage the cluster.
	ClientCertificateCommonNames ClientCertificateCommonNameResponseArrayInput
	// The list of client certificates referenced by thumbprint that are allowed to manage the cluster.
	ClientCertificateThumbprints ClientCertificateThumbprintResponseArrayInput
	// The Service Fabric runtime version of the cluster. This property can only by set the user when **upgradeMode** is set to 'Manual'. To get list of available Service Fabric versions for new clusters use [ClusterVersion API](./ClusterVersion.md). To get the list of available version for existing clusters use **availableClusterVersions**.
	ClusterCodeVersion pulumi.StringPtrInput
	// The Azure Resource Provider endpoint. A system service in the cluster connects to this  endpoint.
	ClusterEndpoint pulumi.StringPtrInput
	// A service generated unique identifier for the cluster resource.
	ClusterId pulumi.StringPtrInput
	// The current state of the cluster.
	//
	//   - WaitingForNodes - Indicates that the cluster resource is created and the resource provider is waiting for Service Fabric VM extension to boot up and report to it.
	//   - Deploying - Indicates that the Service Fabric runtime is being installed on the VMs. Cluster resource will be in this state until the cluster boots up and system services are up.
	//   - BaselineUpgrade - Indicates that the cluster is upgrading to establishes the cluster version. This upgrade is automatically initiated when the cluster boots up for the first time.
	//   - UpdatingUserConfiguration - Indicates that the cluster is being upgraded with the user provided configuration.
	//   - UpdatingUserCertificate - Indicates that the cluster is being upgraded with the user provided certificate.
	//   - UpdatingInfrastructure - Indicates that the cluster is being upgraded with the latest Service Fabric runtime version. This happens only when the **upgradeMode** is set to 'Automatic'.
	//   - EnforcingClusterVersion - Indicates that cluster is on a different version than expected and the cluster is being upgraded to the expected version.
	//   - UpgradeServiceUnreachable - Indicates that the system service in the cluster is no longer polling the Resource Provider. Clusters in this state cannot be managed by the Resource Provider.
	//   - AutoScale - Indicates that the ReliabilityLevel of the cluster is being adjusted.
	//   - Ready - Indicates that the cluster is in a stable state.
	ClusterState pulumi.StringPtrInput
	// The storage account information for storing Service Fabric diagnostic logs.
	DiagnosticsStorageAccountConfig DiagnosticsStorageAccountConfigResponsePtrInput
	// Azure resource etag.
	Etag pulumi.StringPtrInput
	// Indicates if the event store service is enabled.
	EventStoreServiceEnabled pulumi.BoolPtrInput
	// The list of custom fabric settings to configure the cluster.
	FabricSettings SettingsSectionDescriptionResponseArrayInput
	// Azure resource location.
	Location pulumi.StringPtrInput
	// The http management endpoint of the cluster.
	ManagementEndpoint pulumi.StringPtrInput
	// Azure resource name.
	Name pulumi.StringPtrInput
	// The list of node types in the cluster.
	NodeTypes NodeTypeDescriptionResponseArrayInput
	// The provisioning state of the cluster resource.
	ProvisioningState pulumi.StringPtrInput
	// The reliability level sets the replica set size of system services. Learn about [ReliabilityLevel](https://docs.microsoft.com/azure/service-fabric/service-fabric-cluster-capacity).
	//
	//   - None - Run the System services with a target replica set count of 1. This should only be used for test clusters.
	//   - Bronze - Run the System services with a target replica set count of 3. This should only be used for test clusters.
	//   - Silver - Run the System services with a target replica set count of 5.
	//   - Gold - Run the System services with a target replica set count of 7.
	//   - Platinum - Run the System services with a target replica set count of 9.
	ReliabilityLevel pulumi.StringPtrInput
	// The server certificate used by reverse proxy.
	ReverseProxyCertificate CertificateDescriptionResponsePtrInput
	// Describes a list of server certificates referenced by common name that are used to secure the cluster.
	ReverseProxyCertificateCommonNames ServerCertificateCommonNamesResponsePtrInput
	// Azure resource tags.
	Tags pulumi.StringMapInput
	// Azure resource type.
	Type pulumi.StringPtrInput
	// The policy to use when upgrading the cluster.
	UpgradeDescription ClusterUpgradePolicyResponsePtrInput
	// The upgrade mode of the cluster when new Service Fabric runtime version is available.
	//
	//   - Automatic - The cluster will be automatically upgraded to the latest Service Fabric runtime version as soon as it is available.
	//   - Manual - The cluster will not be automatically upgraded to the latest Service Fabric runtime version. The cluster is upgraded by setting the **clusterCodeVersion** property in the cluster resource.
	UpgradeMode pulumi.StringPtrInput
	// The VM image VMSS has been configured with. Generic names such as Windows or Linux can be used.
	VmImage pulumi.StringPtrInput
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	// The list of add-on features to enable in the cluster.
	AddOnFeatures []string `pulumi:"addOnFeatures"`
	// The AAD authentication settings of the cluster.
	AzureActiveDirectory *AzureActiveDirectory `pulumi:"azureActiveDirectory"`
	// The certificate to use for securing the cluster. The certificate provided will be used for node to node security within the cluster, SSL certificate for cluster management endpoint and default admin client.
	Certificate *CertificateDescription `pulumi:"certificate"`
	// Describes a list of server certificates referenced by common name that are used to secure the cluster.
	CertificateCommonNames *ServerCertificateCommonNames `pulumi:"certificateCommonNames"`
	// The list of client certificates referenced by common name that are allowed to manage the cluster.
	ClientCertificateCommonNames []ClientCertificateCommonName `pulumi:"clientCertificateCommonNames"`
	// The list of client certificates referenced by thumbprint that are allowed to manage the cluster.
	ClientCertificateThumbprints []ClientCertificateThumbprint `pulumi:"clientCertificateThumbprints"`
	// The Service Fabric runtime version of the cluster. This property can only by set the user when **upgradeMode** is set to 'Manual'. To get list of available Service Fabric versions for new clusters use [ClusterVersion API](./ClusterVersion.md). To get the list of available version for existing clusters use **availableClusterVersions**.
	ClusterCodeVersion *string `pulumi:"clusterCodeVersion"`
	// The name of the cluster resource.
	ClusterName string `pulumi:"clusterName"`
	// The storage account information for storing Service Fabric diagnostic logs.
	DiagnosticsStorageAccountConfig *DiagnosticsStorageAccountConfig `pulumi:"diagnosticsStorageAccountConfig"`
	// Indicates if the event store service is enabled.
	EventStoreServiceEnabled *bool `pulumi:"eventStoreServiceEnabled"`
	// The list of custom fabric settings to configure the cluster.
	FabricSettings []SettingsSectionDescription `pulumi:"fabricSettings"`
	// Azure resource location.
	Location string `pulumi:"location"`
	// The http management endpoint of the cluster.
	ManagementEndpoint string `pulumi:"managementEndpoint"`
	// The list of node types in the cluster.
	NodeTypes []NodeTypeDescription `pulumi:"nodeTypes"`
	// The reliability level sets the replica set size of system services. Learn about [ReliabilityLevel](https://docs.microsoft.com/azure/service-fabric/service-fabric-cluster-capacity).
	//
	//   - None - Run the System services with a target replica set count of 1. This should only be used for test clusters.
	//   - Bronze - Run the System services with a target replica set count of 3. This should only be used for test clusters.
	//   - Silver - Run the System services with a target replica set count of 5.
	//   - Gold - Run the System services with a target replica set count of 7.
	//   - Platinum - Run the System services with a target replica set count of 9.
	ReliabilityLevel *string `pulumi:"reliabilityLevel"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The server certificate used by reverse proxy.
	ReverseProxyCertificate *CertificateDescription `pulumi:"reverseProxyCertificate"`
	// Describes a list of server certificates referenced by common name that are used to secure the cluster.
	ReverseProxyCertificateCommonNames *ServerCertificateCommonNames `pulumi:"reverseProxyCertificateCommonNames"`
	// Azure resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The policy to use when upgrading the cluster.
	UpgradeDescription *ClusterUpgradePolicy `pulumi:"upgradeDescription"`
	// The upgrade mode of the cluster when new Service Fabric runtime version is available.
	//
	//   - Automatic - The cluster will be automatically upgraded to the latest Service Fabric runtime version as soon as it is available.
	//   - Manual - The cluster will not be automatically upgraded to the latest Service Fabric runtime version. The cluster is upgraded by setting the **clusterCodeVersion** property in the cluster resource.
	UpgradeMode *string `pulumi:"upgradeMode"`
	// The VM image VMSS has been configured with. Generic names such as Windows or Linux can be used.
	VmImage *string `pulumi:"vmImage"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// The list of add-on features to enable in the cluster.
	AddOnFeatures pulumi.StringArrayInput
	// The AAD authentication settings of the cluster.
	AzureActiveDirectory AzureActiveDirectoryPtrInput
	// The certificate to use for securing the cluster. The certificate provided will be used for node to node security within the cluster, SSL certificate for cluster management endpoint and default admin client.
	Certificate CertificateDescriptionPtrInput
	// Describes a list of server certificates referenced by common name that are used to secure the cluster.
	CertificateCommonNames ServerCertificateCommonNamesPtrInput
	// The list of client certificates referenced by common name that are allowed to manage the cluster.
	ClientCertificateCommonNames ClientCertificateCommonNameArrayInput
	// The list of client certificates referenced by thumbprint that are allowed to manage the cluster.
	ClientCertificateThumbprints ClientCertificateThumbprintArrayInput
	// The Service Fabric runtime version of the cluster. This property can only by set the user when **upgradeMode** is set to 'Manual'. To get list of available Service Fabric versions for new clusters use [ClusterVersion API](./ClusterVersion.md). To get the list of available version for existing clusters use **availableClusterVersions**.
	ClusterCodeVersion pulumi.StringPtrInput
	// The name of the cluster resource.
	ClusterName pulumi.StringInput
	// The storage account information for storing Service Fabric diagnostic logs.
	DiagnosticsStorageAccountConfig DiagnosticsStorageAccountConfigPtrInput
	// Indicates if the event store service is enabled.
	EventStoreServiceEnabled pulumi.BoolPtrInput
	// The list of custom fabric settings to configure the cluster.
	FabricSettings SettingsSectionDescriptionArrayInput
	// Azure resource location.
	Location pulumi.StringInput
	// The http management endpoint of the cluster.
	ManagementEndpoint pulumi.StringInput
	// The list of node types in the cluster.
	NodeTypes NodeTypeDescriptionArrayInput
	// The reliability level sets the replica set size of system services. Learn about [ReliabilityLevel](https://docs.microsoft.com/azure/service-fabric/service-fabric-cluster-capacity).
	//
	//   - None - Run the System services with a target replica set count of 1. This should only be used for test clusters.
	//   - Bronze - Run the System services with a target replica set count of 3. This should only be used for test clusters.
	//   - Silver - Run the System services with a target replica set count of 5.
	//   - Gold - Run the System services with a target replica set count of 7.
	//   - Platinum - Run the System services with a target replica set count of 9.
	ReliabilityLevel pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The server certificate used by reverse proxy.
	ReverseProxyCertificate CertificateDescriptionPtrInput
	// Describes a list of server certificates referenced by common name that are used to secure the cluster.
	ReverseProxyCertificateCommonNames ServerCertificateCommonNamesPtrInput
	// Azure resource tags.
	Tags pulumi.StringMapInput
	// The policy to use when upgrading the cluster.
	UpgradeDescription ClusterUpgradePolicyPtrInput
	// The upgrade mode of the cluster when new Service Fabric runtime version is available.
	//
	//   - Automatic - The cluster will be automatically upgraded to the latest Service Fabric runtime version as soon as it is available.
	//   - Manual - The cluster will not be automatically upgraded to the latest Service Fabric runtime version. The cluster is upgraded by setting the **clusterCodeVersion** property in the cluster resource.
	UpgradeMode pulumi.StringPtrInput
	// The VM image VMSS has been configured with. Generic names such as Windows or Linux can be used.
	VmImage pulumi.StringPtrInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}
