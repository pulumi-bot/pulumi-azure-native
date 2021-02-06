// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Single Event Hubs Cluster resource in List or Get operations.
type Cluster struct {
	pulumi.CustomResourceState

	// The UTC time when the Event Hubs Cluster was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The metric ID of the cluster resource. Provided by the service and not modifiable by the user.
	MetricId pulumi.StringOutput `pulumi:"metricId"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the cluster SKU.
	Sku ClusterSkuResponsePtrOutput `pulumi:"sku"`
	// Status of the Cluster resource
	Status pulumi.StringOutput `pulumi:"status"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The UTC time when the Event Hubs Cluster was last updated.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource Cluster
	err := ctx.RegisterResource("azure-nextgen:eventhub/v20180101preview:Cluster", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-nextgen:eventhub/v20180101preview:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
	// The UTC time when the Event Hubs Cluster was created.
	CreatedAt *string `pulumi:"createdAt"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The metric ID of the cluster resource. Provided by the service and not modifiable by the user.
	MetricId *string `pulumi:"metricId"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Properties of the cluster SKU.
	Sku *ClusterSkuResponse `pulumi:"sku"`
	// Status of the Cluster resource
	Status *string `pulumi:"status"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
	// The UTC time when the Event Hubs Cluster was last updated.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type ClusterState struct {
	// The UTC time when the Event Hubs Cluster was created.
	CreatedAt pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The metric ID of the cluster resource. Provided by the service and not modifiable by the user.
	MetricId pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Properties of the cluster SKU.
	Sku ClusterSkuResponsePtrInput
	// Status of the Cluster resource
	Status pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
	// The UTC time when the Event Hubs Cluster was last updated.
	UpdatedAt pulumi.StringPtrInput
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	// The name of the Event Hubs Cluster.
	ClusterName string `pulumi:"clusterName"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Name of the resource group within the azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Properties of the cluster SKU.
	Sku *ClusterSku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// The name of the Event Hubs Cluster.
	ClusterName pulumi.StringInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Name of the resource group within the azure subscription.
	ResourceGroupName pulumi.StringInput
	// Properties of the cluster SKU.
	Sku ClusterSkuPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}

type ClusterInput interface {
	pulumi.Input

	ToClusterOutput() ClusterOutput
	ToClusterOutputWithContext(ctx context.Context) ClusterOutput
}

func (*Cluster) ElementType() reflect.Type {
	return reflect.TypeOf((*Cluster)(nil))
}

func (i *Cluster) ToClusterOutput() ClusterOutput {
	return i.ToClusterOutputWithContext(context.Background())
}

func (i *Cluster) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterOutput)
}

type ClusterOutput struct {
	*pulumi.OutputState
}

func (ClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Cluster)(nil))
}

func (o ClusterOutput) ToClusterOutput() ClusterOutput {
	return o
}

func (o ClusterOutput) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ClusterOutput{})
}
