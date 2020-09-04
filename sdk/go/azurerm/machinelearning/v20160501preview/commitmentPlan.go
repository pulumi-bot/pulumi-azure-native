// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160501preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An Azure ML commitment plan resource.
type CommitmentPlan struct {
	pulumi.CustomResourceState

	// An entity tag used to enforce optimistic concurrency.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The commitment plan properties.
	Properties CommitmentPlanPropertiesResponseOutput `pulumi:"properties"`
	// The commitment plan SKU.
	Sku ResourceSkuResponsePtrOutput `pulumi:"sku"`
	// User-defined tags for the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCommitmentPlan registers a new resource with the given unique name, arguments, and options.
func NewCommitmentPlan(ctx *pulumi.Context,
	name string, args *CommitmentPlanArgs, opts ...pulumi.ResourceOption) (*CommitmentPlan, error) {
	if args == nil || args.CommitmentPlanName == nil {
		return nil, errors.New("missing required argument 'CommitmentPlanName'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &CommitmentPlanArgs{}
	}
	var resource CommitmentPlan
	err := ctx.RegisterResource("azurerm:machinelearning/v20160501preview:CommitmentPlan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCommitmentPlan gets an existing CommitmentPlan resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCommitmentPlan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CommitmentPlanState, opts ...pulumi.ResourceOption) (*CommitmentPlan, error) {
	var resource CommitmentPlan
	err := ctx.ReadResource("azurerm:machinelearning/v20160501preview:CommitmentPlan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CommitmentPlan resources.
type commitmentPlanState struct {
	// An entity tag used to enforce optimistic concurrency.
	Etag *string `pulumi:"etag"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name *string `pulumi:"name"`
	// The commitment plan properties.
	Properties *CommitmentPlanPropertiesResponse `pulumi:"properties"`
	// The commitment plan SKU.
	Sku *ResourceSkuResponse `pulumi:"sku"`
	// User-defined tags for the resource.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type CommitmentPlanState struct {
	// An entity tag used to enforce optimistic concurrency.
	Etag pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// The commitment plan properties.
	Properties CommitmentPlanPropertiesResponsePtrInput
	// The commitment plan SKU.
	Sku ResourceSkuResponsePtrInput
	// User-defined tags for the resource.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (CommitmentPlanState) ElementType() reflect.Type {
	return reflect.TypeOf((*commitmentPlanState)(nil)).Elem()
}

type commitmentPlanArgs struct {
	// The Azure ML commitment plan name.
	CommitmentPlanName string `pulumi:"commitmentPlanName"`
	// An entity tag used to enforce optimistic concurrency.
	Etag *string `pulumi:"etag"`
	// Resource location.
	Location string `pulumi:"location"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The commitment plan SKU.
	Sku *ResourceSku `pulumi:"sku"`
	// User-defined tags for the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a CommitmentPlan resource.
type CommitmentPlanArgs struct {
	// The Azure ML commitment plan name.
	CommitmentPlanName pulumi.StringInput
	// An entity tag used to enforce optimistic concurrency.
	Etag pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The commitment plan SKU.
	Sku ResourceSkuPtrInput
	// User-defined tags for the resource.
	Tags pulumi.StringMapInput
}

func (CommitmentPlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*commitmentPlanArgs)(nil)).Elem()
}
