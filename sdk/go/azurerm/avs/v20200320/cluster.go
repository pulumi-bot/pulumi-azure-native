// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200320

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A cluster resource
type Cluster struct {
	pulumi.CustomResourceState

	// The identity
	ClusterId pulumi.IntOutput `pulumi:"clusterId"`
	// The cluster size
	ClusterSize pulumi.IntOutput `pulumi:"clusterSize"`
	// The hosts
	Hosts pulumi.StringArrayOutput `pulumi:"hosts"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The state of the cluster provisioning
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The cluster SKU
	Sku SkuResponseOutput `pulumi:"sku"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil || args.ClusterName == nil {
		return nil, errors.New("missing required argument 'ClusterName'")
	}
	if args == nil || args.ClusterSize == nil {
		return nil, errors.New("missing required argument 'ClusterSize'")
	}
	if args == nil || args.PrivateCloudName == nil {
		return nil, errors.New("missing required argument 'PrivateCloudName'")
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
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:avs/latest:Cluster"),
		},
		{
			Type: pulumi.String("azurerm:avs/v20190809preview:Cluster"),
		},
	})
	opts = append(opts, aliases)
	var resource Cluster
	err := ctx.RegisterResource("azurerm:avs/v20200320:Cluster", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azurerm:avs/v20200320:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
	// The identity
	ClusterId *int `pulumi:"clusterId"`
	// The cluster size
	ClusterSize *int `pulumi:"clusterSize"`
	// The hosts
	Hosts []string `pulumi:"hosts"`
	// Resource name.
	Name *string `pulumi:"name"`
	// The state of the cluster provisioning
	ProvisioningState *string `pulumi:"provisioningState"`
	// The cluster SKU
	Sku *SkuResponse `pulumi:"sku"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type ClusterState struct {
	// The identity
	ClusterId pulumi.IntPtrInput
	// The cluster size
	ClusterSize pulumi.IntPtrInput
	// The hosts
	Hosts pulumi.StringArrayInput
	// Resource name.
	Name pulumi.StringPtrInput
	// The state of the cluster provisioning
	ProvisioningState pulumi.StringPtrInput
	// The cluster SKU
	Sku SkuResponsePtrInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	// Name of the cluster in the private cloud
	ClusterName string `pulumi:"clusterName"`
	// The cluster size
	ClusterSize int `pulumi:"clusterSize"`
	// The name of the private cloud.
	PrivateCloudName string `pulumi:"privateCloudName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The cluster SKU
	Sku Sku `pulumi:"sku"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// Name of the cluster in the private cloud
	ClusterName pulumi.StringInput
	// The cluster size
	ClusterSize pulumi.IntInput
	// The name of the private cloud.
	PrivateCloudName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The cluster SKU
	Sku SkuInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}
