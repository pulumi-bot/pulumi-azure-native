// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The top level Log Analytics cluster resource container.
type Cluster struct {
	pulumi.CustomResourceState

	// The ID associated with the cluster.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// The identity of the resource.
	Identity IdentityResponsePtrOutput `pulumi:"identity"`
	// The associated key properties.
	KeyVaultProperties KeyVaultPropertiesResponsePtrOutput `pulumi:"keyVaultProperties"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The link used to get the next page of recommendations.
	NextLink pulumi.StringPtrOutput `pulumi:"nextLink"`
	// The provisioning state of the cluster.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The sku properties.
	Sku ClusterSkuResponsePtrOutput `pulumi:"sku"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
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
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ClusterArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:operationalinsights/v20190801preview:Cluster"),
		},
		{
			Type: pulumi.String("azurerm:operationalinsights/v20200301preview:Cluster"),
		},
		{
			Type: pulumi.String("azurerm:operationalinsights/v20200801:Cluster"),
		},
	})
	opts = append(opts, aliases)
	var resource Cluster
	err := ctx.RegisterResource("azurerm:operationalinsights/latest:Cluster", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azurerm:operationalinsights/latest:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
	// The ID associated with the cluster.
	ClusterId *string `pulumi:"clusterId"`
	// The identity of the resource.
	Identity *IdentityResponse `pulumi:"identity"`
	// The associated key properties.
	KeyVaultProperties *KeyVaultPropertiesResponse `pulumi:"keyVaultProperties"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// The link used to get the next page of recommendations.
	NextLink *string `pulumi:"nextLink"`
	// The provisioning state of the cluster.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The sku properties.
	Sku *ClusterSkuResponse `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `pulumi:"type"`
}

type ClusterState struct {
	// The ID associated with the cluster.
	ClusterId pulumi.StringPtrInput
	// The identity of the resource.
	Identity IdentityResponsePtrInput
	// The associated key properties.
	KeyVaultProperties KeyVaultPropertiesResponsePtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// The link used to get the next page of recommendations.
	NextLink pulumi.StringPtrInput
	// The provisioning state of the cluster.
	ProvisioningState pulumi.StringPtrInput
	// The sku properties.
	Sku ClusterSkuResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringPtrInput
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	// The name of the Log Analytics cluster.
	ClusterName string `pulumi:"clusterName"`
	// The identity of the resource.
	Identity *Identity `pulumi:"identity"`
	// The associated key properties.
	KeyVaultProperties *KeyVaultProperties `pulumi:"keyVaultProperties"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The link used to get the next page of recommendations.
	NextLink *string `pulumi:"nextLink"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The sku properties.
	Sku *ClusterSku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// The name of the Log Analytics cluster.
	ClusterName pulumi.StringInput
	// The identity of the resource.
	Identity IdentityPtrInput
	// The associated key properties.
	KeyVaultProperties KeyVaultPropertiesPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringInput
	// The link used to get the next page of recommendations.
	NextLink pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The sku properties.
	Sku ClusterSkuPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}
