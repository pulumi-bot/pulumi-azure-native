// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191109

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Class representing a Kusto cluster.
type Cluster struct {
	pulumi.CustomResourceState

	// The identity of the cluster, if configured.
	Identity IdentityResponsePtrOutput `pulumi:"identity"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The cluster properties.
	Properties ClusterPropertiesResponseOutput `pulumi:"properties"`
	// The SKU of the cluster.
	Sku AzureSkuResponseOutput `pulumi:"sku"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
	// The availability zones of the cluster.
	Zones ZonesResponsePtrOutput `pulumi:"zones"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Sku == nil {
		return nil, errors.New("missing required argument 'Sku'")
	}
	if args == nil {
		args = &ClusterArgs{}
	}
	var resource Cluster
	err := ctx.RegisterResource("azurerm:kusto/v20191109:Cluster", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azurerm:kusto/v20191109:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
	// The identity of the cluster, if configured.
	Identity *IdentityResponse `pulumi:"identity"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// The cluster properties.
	Properties *ClusterPropertiesResponse `pulumi:"properties"`
	// The SKU of the cluster.
	Sku *AzureSkuResponse `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `pulumi:"type"`
	// The availability zones of the cluster.
	Zones *ZonesResponse `pulumi:"zones"`
}

type ClusterState struct {
	// The identity of the cluster, if configured.
	Identity IdentityResponsePtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// The cluster properties.
	Properties ClusterPropertiesResponsePtrInput
	// The SKU of the cluster.
	Sku AzureSkuResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringPtrInput
	// The availability zones of the cluster.
	Zones ZonesResponsePtrInput
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	// A boolean value that indicates if the cluster's disks are encrypted.
	EnableDiskEncryption *bool `pulumi:"enableDiskEncryption"`
	// A boolean value that indicates if the streaming ingest is enabled.
	EnableStreamingIngest *bool `pulumi:"enableStreamingIngest"`
	// The identity of the cluster, if configured.
	Identity *Identity `pulumi:"identity"`
	// KeyVault properties for the cluster encryption.
	KeyVaultProperties *KeyVaultProperties `pulumi:"keyVaultProperties"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the Kusto cluster.
	Name string `pulumi:"name"`
	// Optimized auto scale definition.
	OptimizedAutoscale *OptimizedAutoscale `pulumi:"optimizedAutoscale"`
	// The name of the resource group containing the Kusto cluster.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The SKU of the cluster.
	Sku AzureSku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The cluster's external tenants.
	TrustedExternalTenants []TrustedExternalTenant `pulumi:"trustedExternalTenants"`
	// Virtual network definition.
	VirtualNetworkConfiguration *VirtualNetworkConfiguration `pulumi:"virtualNetworkConfiguration"`
	// The availability zones of the cluster.
	Zones *Zones `pulumi:"zones"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// A boolean value that indicates if the cluster's disks are encrypted.
	EnableDiskEncryption pulumi.BoolPtrInput
	// A boolean value that indicates if the streaming ingest is enabled.
	EnableStreamingIngest pulumi.BoolPtrInput
	// The identity of the cluster, if configured.
	Identity IdentityPtrInput
	// KeyVault properties for the cluster encryption.
	KeyVaultProperties KeyVaultPropertiesPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringInput
	// The name of the Kusto cluster.
	Name pulumi.StringInput
	// Optimized auto scale definition.
	OptimizedAutoscale OptimizedAutoscalePtrInput
	// The name of the resource group containing the Kusto cluster.
	ResourceGroupName pulumi.StringInput
	// The SKU of the cluster.
	Sku AzureSkuInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The cluster's external tenants.
	TrustedExternalTenants TrustedExternalTenantArrayInput
	// Virtual network definition.
	VirtualNetworkConfiguration VirtualNetworkConfigurationPtrInput
	// The availability zones of the cluster.
	Zones ZonesPtrInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}
