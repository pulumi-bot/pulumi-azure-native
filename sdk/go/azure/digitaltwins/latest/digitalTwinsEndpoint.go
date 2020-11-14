// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// DigitalTwinsInstance endpoint resource.
type DigitalTwinsEndpoint struct {
	pulumi.CustomResourceState

	// Extension resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// DigitalTwinsInstance endpoint resource properties.
	Properties pulumi.AnyOutput `pulumi:"properties"`
	// The resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDigitalTwinsEndpoint registers a new resource with the given unique name, arguments, and options.
func NewDigitalTwinsEndpoint(ctx *pulumi.Context,
	name string, args *DigitalTwinsEndpointArgs, opts ...pulumi.ResourceOption) (*DigitalTwinsEndpoint, error) {
	if args == nil || args.EndpointName == nil {
		return nil, errors.New("missing required argument 'EndpointName'")
	}
	if args == nil || args.Properties == nil {
		return nil, errors.New("missing required argument 'Properties'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ResourceName == nil {
		return nil, errors.New("missing required argument 'ResourceName'")
	}
	if args == nil {
		args = &DigitalTwinsEndpointArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:digitaltwins/v20200301preview:DigitalTwinsEndpoint"),
		},
		{
			Type: pulumi.String("azure-nextgen:digitaltwins/v20201031:DigitalTwinsEndpoint"),
		},
	})
	opts = append(opts, aliases)
	var resource DigitalTwinsEndpoint
	err := ctx.RegisterResource("azure-nextgen:digitaltwins/latest:DigitalTwinsEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDigitalTwinsEndpoint gets an existing DigitalTwinsEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDigitalTwinsEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DigitalTwinsEndpointState, opts ...pulumi.ResourceOption) (*DigitalTwinsEndpoint, error) {
	var resource DigitalTwinsEndpoint
	err := ctx.ReadResource("azure-nextgen:digitaltwins/latest:DigitalTwinsEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DigitalTwinsEndpoint resources.
type digitalTwinsEndpointState struct {
	// Extension resource name.
	Name *string `pulumi:"name"`
	// DigitalTwinsInstance endpoint resource properties.
	Properties interface{} `pulumi:"properties"`
	// The resource type.
	Type *string `pulumi:"type"`
}

type DigitalTwinsEndpointState struct {
	// Extension resource name.
	Name pulumi.StringPtrInput
	// DigitalTwinsInstance endpoint resource properties.
	Properties pulumi.Input
	// The resource type.
	Type pulumi.StringPtrInput
}

func (DigitalTwinsEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*digitalTwinsEndpointState)(nil)).Elem()
}

type digitalTwinsEndpointArgs struct {
	// Name of Endpoint Resource.
	EndpointName string `pulumi:"endpointName"`
	// DigitalTwinsInstance endpoint resource properties.
	Properties interface{} `pulumi:"properties"`
	// The name of the resource group that contains the DigitalTwinsInstance.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the DigitalTwinsInstance.
	ResourceName string `pulumi:"resourceName"`
}

// The set of arguments for constructing a DigitalTwinsEndpoint resource.
type DigitalTwinsEndpointArgs struct {
	// Name of Endpoint Resource.
	EndpointName pulumi.StringInput
	// DigitalTwinsInstance endpoint resource properties.
	Properties pulumi.Input
	// The name of the resource group that contains the DigitalTwinsInstance.
	ResourceGroupName pulumi.StringInput
	// The name of the DigitalTwinsInstance.
	ResourceName pulumi.StringInput
}

func (DigitalTwinsEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*digitalTwinsEndpointArgs)(nil)).Elem()
}

type DigitalTwinsEndpointInput interface {
	pulumi.Input

	ToDigitalTwinsEndpointOutput() DigitalTwinsEndpointOutput
	ToDigitalTwinsEndpointOutputWithContext(ctx context.Context) DigitalTwinsEndpointOutput
}

func (DigitalTwinsEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((*DigitalTwinsEndpoint)(nil)).Elem()
}

func (i DigitalTwinsEndpoint) ToDigitalTwinsEndpointOutput() DigitalTwinsEndpointOutput {
	return i.ToDigitalTwinsEndpointOutputWithContext(context.Background())
}

func (i DigitalTwinsEndpoint) ToDigitalTwinsEndpointOutputWithContext(ctx context.Context) DigitalTwinsEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DigitalTwinsEndpointOutput)
}

type DigitalTwinsEndpointOutput struct {
	*pulumi.OutputState
}

func (DigitalTwinsEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DigitalTwinsEndpointOutput)(nil)).Elem()
}

func (o DigitalTwinsEndpointOutput) ToDigitalTwinsEndpointOutput() DigitalTwinsEndpointOutput {
	return o
}

func (o DigitalTwinsEndpointOutput) ToDigitalTwinsEndpointOutputWithContext(ctx context.Context) DigitalTwinsEndpointOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DigitalTwinsEndpointOutput{})
}
