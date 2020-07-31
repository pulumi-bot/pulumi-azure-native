// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180701

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A Streaming Locator resource
type StreamingLocator struct {
	pulumi.CustomResourceState

	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the Streaming Locator.
	Properties StreamingLocatorPropertiesResponseOutput `pulumi:"properties"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewStreamingLocator registers a new resource with the given unique name, arguments, and options.
func NewStreamingLocator(ctx *pulumi.Context,
	name string, args *StreamingLocatorArgs, opts ...pulumi.ResourceOption) (*StreamingLocator, error) {
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
		args = &StreamingLocatorArgs{}
	}
	var resource StreamingLocator
	err := ctx.RegisterResource("azurerm:media/v20180701:StreamingLocator", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStreamingLocator gets an existing StreamingLocator resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStreamingLocator(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StreamingLocatorState, opts ...pulumi.ResourceOption) (*StreamingLocator, error) {
	var resource StreamingLocator
	err := ctx.ReadResource("azurerm:media/v20180701:StreamingLocator", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StreamingLocator resources.
type streamingLocatorState struct {
	// The name of the resource.
	Name *string `pulumi:"name"`
	// Properties of the Streaming Locator.
	Properties *StreamingLocatorPropertiesResponse `pulumi:"properties"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

type StreamingLocatorState struct {
	// The name of the resource.
	Name pulumi.StringPtrInput
	// Properties of the Streaming Locator.
	Properties StreamingLocatorPropertiesResponsePtrInput
	// The type of the resource.
	Type pulumi.StringPtrInput
}

func (StreamingLocatorState) ElementType() reflect.Type {
	return reflect.TypeOf((*streamingLocatorState)(nil)).Elem()
}

type streamingLocatorArgs struct {
	// The Media Services account name.
	AccountName string `pulumi:"accountName"`
	// The Streaming Locator name.
	Name string `pulumi:"name"`
	// Properties of the Streaming Locator.
	Properties *StreamingLocatorProperties `pulumi:"properties"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a StreamingLocator resource.
type StreamingLocatorArgs struct {
	// The Media Services account name.
	AccountName pulumi.StringInput
	// The Streaming Locator name.
	Name pulumi.StringInput
	// Properties of the Streaming Locator.
	Properties StreamingLocatorPropertiesPtrInput
	// The name of the resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
}

func (StreamingLocatorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*streamingLocatorArgs)(nil)).Elem()
}