// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180701

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A Streaming Policy resource
type StreamingPolicy struct {
	pulumi.CustomResourceState

	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Class to specify properties of Streaming Policy
	Properties StreamingPolicyPropertiesResponseOutput `pulumi:"properties"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewStreamingPolicy registers a new resource with the given unique name, arguments, and options.
func NewStreamingPolicy(ctx *pulumi.Context,
	name string, args *StreamingPolicyArgs, opts ...pulumi.ResourceOption) (*StreamingPolicy, error) {
	if args == nil || args.AccountName == nil {
		return nil, errors.New("missing required argument 'AccountName'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &StreamingPolicyArgs{}
	}
	var resource StreamingPolicy
	err := ctx.RegisterResource("azurerm:media/v20180701:StreamingPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStreamingPolicy gets an existing StreamingPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStreamingPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StreamingPolicyState, opts ...pulumi.ResourceOption) (*StreamingPolicy, error) {
	var resource StreamingPolicy
	err := ctx.ReadResource("azurerm:media/v20180701:StreamingPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StreamingPolicy resources.
type streamingPolicyState struct {
	// The name of the resource.
	Name *string `pulumi:"name"`
	// Class to specify properties of Streaming Policy
	Properties *StreamingPolicyPropertiesResponse `pulumi:"properties"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

type StreamingPolicyState struct {
	// The name of the resource.
	Name pulumi.StringPtrInput
	// Class to specify properties of Streaming Policy
	Properties StreamingPolicyPropertiesResponsePtrInput
	// The type of the resource.
	Type pulumi.StringPtrInput
}

func (StreamingPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*streamingPolicyState)(nil)).Elem()
}

type streamingPolicyArgs struct {
	// The Media Services account name.
	AccountName string `pulumi:"accountName"`
	// The Streaming Policy name.
	Name string `pulumi:"name"`
	// Class to specify properties of Streaming Policy
	Properties *StreamingPolicyProperties `pulumi:"properties"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a StreamingPolicy resource.
type StreamingPolicyArgs struct {
	// The Media Services account name.
	AccountName pulumi.StringInput
	// The Streaming Policy name.
	Name pulumi.StringInput
	// Class to specify properties of Streaming Policy
	Properties StreamingPolicyPropertiesPtrInput
	// The name of the resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
}

func (StreamingPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*streamingPolicyArgs)(nil)).Elem()
}