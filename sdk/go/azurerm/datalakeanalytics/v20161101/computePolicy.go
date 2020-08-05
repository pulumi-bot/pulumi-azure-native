// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20161101

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Data Lake Analytics compute policy information.
type ComputePolicy struct {
	pulumi.CustomResourceState

	// The resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The compute policy properties.
	Properties ComputePolicyPropertiesResponseOutput `pulumi:"properties"`
	// The resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewComputePolicy registers a new resource with the given unique name, arguments, and options.
func NewComputePolicy(ctx *pulumi.Context,
	name string, args *ComputePolicyArgs, opts ...pulumi.ResourceOption) (*ComputePolicy, error) {
	if args == nil || args.AccountName == nil {
		return nil, errors.New("missing required argument 'AccountName'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ObjectId == nil {
		return nil, errors.New("missing required argument 'ObjectId'")
	}
	if args == nil || args.ObjectType == nil {
		return nil, errors.New("missing required argument 'ObjectType'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ComputePolicyArgs{}
	}
	var resource ComputePolicy
	err := ctx.RegisterResource("azurerm:datalakeanalytics/v20161101:ComputePolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComputePolicy gets an existing ComputePolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputePolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputePolicyState, opts ...pulumi.ResourceOption) (*ComputePolicy, error) {
	var resource ComputePolicy
	err := ctx.ReadResource("azurerm:datalakeanalytics/v20161101:ComputePolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComputePolicy resources.
type computePolicyState struct {
	// The resource name.
	Name *string `pulumi:"name"`
	// The compute policy properties.
	Properties *ComputePolicyPropertiesResponse `pulumi:"properties"`
	// The resource type.
	Type *string `pulumi:"type"`
}

type ComputePolicyState struct {
	// The resource name.
	Name pulumi.StringPtrInput
	// The compute policy properties.
	Properties ComputePolicyPropertiesResponsePtrInput
	// The resource type.
	Type pulumi.StringPtrInput
}

func (ComputePolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*computePolicyState)(nil)).Elem()
}

type computePolicyArgs struct {
	// The name of the Data Lake Analytics account.
	AccountName string `pulumi:"accountName"`
	// The maximum degree of parallelism per job this user can use to submit jobs. This property, the min priority per job property, or both must be passed.
	MaxDegreeOfParallelismPerJob *int `pulumi:"maxDegreeOfParallelismPerJob"`
	// The minimum priority per job this user can use to submit jobs. This property, the max degree of parallelism per job property, or both must be passed.
	MinPriorityPerJob *int `pulumi:"minPriorityPerJob"`
	// The name of the compute policy to create or update.
	Name string `pulumi:"name"`
	// The AAD object identifier for the entity to create a policy for.
	ObjectId string `pulumi:"objectId"`
	// The type of AAD object the object identifier refers to.
	ObjectType string `pulumi:"objectType"`
	// The name of the Azure resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ComputePolicy resource.
type ComputePolicyArgs struct {
	// The name of the Data Lake Analytics account.
	AccountName pulumi.StringInput
	// The maximum degree of parallelism per job this user can use to submit jobs. This property, the min priority per job property, or both must be passed.
	MaxDegreeOfParallelismPerJob pulumi.IntPtrInput
	// The minimum priority per job this user can use to submit jobs. This property, the max degree of parallelism per job property, or both must be passed.
	MinPriorityPerJob pulumi.IntPtrInput
	// The name of the compute policy to create or update.
	Name pulumi.StringInput
	// The AAD object identifier for the entity to create a policy for.
	ObjectId pulumi.StringInput
	// The type of AAD object the object identifier refers to.
	ObjectType pulumi.StringInput
	// The name of the Azure resource group.
	ResourceGroupName pulumi.StringInput
}

func (ComputePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computePolicyArgs)(nil)).Elem()
}
