// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20171201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Bot channel resource definition
type Channel struct {
	pulumi.CustomResourceState

	// Entity Tag
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Required. Gets or sets the Kind of the resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Specifies the location of the resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Specifies the name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The set of properties specific to bot channel resource
	Properties pulumi.AnyOutput `pulumi:"properties"`
	// Gets or sets the SKU of the resource.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Contains resource tags defined as key/value pairs.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewChannel registers a new resource with the given unique name, arguments, and options.
func NewChannel(ctx *pulumi.Context,
	name string, args *ChannelArgs, opts ...pulumi.ResourceOption) (*Channel, error) {
	if args == nil || args.ChannelName == nil {
		return nil, errors.New("missing required argument 'ChannelName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ResourceName == nil {
		return nil, errors.New("missing required argument 'ResourceName'")
	}
	if args == nil {
		args = &ChannelArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:botservice/latest:Channel"),
		},
		{
			Type: pulumi.String("azure-nextgen:botservice/v20180712:Channel"),
		},
		{
			Type: pulumi.String("azure-nextgen:botservice/v20200602:Channel"),
		},
	})
	opts = append(opts, aliases)
	var resource Channel
	err := ctx.RegisterResource("azure-nextgen:botservice/v20171201:Channel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetChannel gets an existing Channel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ChannelState, opts ...pulumi.ResourceOption) (*Channel, error) {
	var resource Channel
	err := ctx.ReadResource("azure-nextgen:botservice/v20171201:Channel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Channel resources.
type channelState struct {
	// Entity Tag
	Etag *string `pulumi:"etag"`
	// Required. Gets or sets the Kind of the resource.
	Kind *string `pulumi:"kind"`
	// Specifies the location of the resource.
	Location *string `pulumi:"location"`
	// Specifies the name of the resource.
	Name *string `pulumi:"name"`
	// The set of properties specific to bot channel resource
	Properties interface{} `pulumi:"properties"`
	// Gets or sets the SKU of the resource.
	Sku *SkuResponse `pulumi:"sku"`
	// Contains resource tags defined as key/value pairs.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the type of the resource.
	Type *string `pulumi:"type"`
}

type ChannelState struct {
	// Entity Tag
	Etag pulumi.StringPtrInput
	// Required. Gets or sets the Kind of the resource.
	Kind pulumi.StringPtrInput
	// Specifies the location of the resource.
	Location pulumi.StringPtrInput
	// Specifies the name of the resource.
	Name pulumi.StringPtrInput
	// The set of properties specific to bot channel resource
	Properties pulumi.Input
	// Gets or sets the SKU of the resource.
	Sku SkuResponsePtrInput
	// Contains resource tags defined as key/value pairs.
	Tags pulumi.StringMapInput
	// Specifies the type of the resource.
	Type pulumi.StringPtrInput
}

func (ChannelState) ElementType() reflect.Type {
	return reflect.TypeOf((*channelState)(nil)).Elem()
}

type channelArgs struct {
	// The name of the Channel resource.
	ChannelName string `pulumi:"channelName"`
	// Entity Tag
	Etag *string `pulumi:"etag"`
	// Required. Gets or sets the Kind of the resource.
	Kind *string `pulumi:"kind"`
	// Specifies the location of the resource.
	Location *string `pulumi:"location"`
	// The set of properties specific to bot channel resource
	Properties interface{} `pulumi:"properties"`
	// The name of the Bot resource group in the user subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Bot resource.
	ResourceName string `pulumi:"resourceName"`
	// Gets or sets the SKU of the resource.
	Sku *Sku `pulumi:"sku"`
	// Contains resource tags defined as key/value pairs.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Channel resource.
type ChannelArgs struct {
	// The name of the Channel resource.
	ChannelName pulumi.StringInput
	// Entity Tag
	Etag pulumi.StringPtrInput
	// Required. Gets or sets the Kind of the resource.
	Kind pulumi.StringPtrInput
	// Specifies the location of the resource.
	Location pulumi.StringPtrInput
	// The set of properties specific to bot channel resource
	Properties pulumi.Input
	// The name of the Bot resource group in the user subscription.
	ResourceGroupName pulumi.StringInput
	// The name of the Bot resource.
	ResourceName pulumi.StringInput
	// Gets or sets the SKU of the resource.
	Sku SkuPtrInput
	// Contains resource tags defined as key/value pairs.
	Tags pulumi.StringMapInput
}

func (ChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*channelArgs)(nil)).Elem()
}

type ChannelInput interface {
	pulumi.Input

	ToChannelOutput() ChannelOutput
	ToChannelOutputWithContext(ctx context.Context) ChannelOutput
}

func (Channel) ElementType() reflect.Type {
	return reflect.TypeOf((*Channel)(nil)).Elem()
}

func (i Channel) ToChannelOutput() ChannelOutput {
	return i.ToChannelOutputWithContext(context.Background())
}

func (i Channel) ToChannelOutputWithContext(ctx context.Context) ChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChannelOutput)
}

type ChannelOutput struct {
	*pulumi.OutputState
}

func (ChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ChannelOutput)(nil)).Elem()
}

func (o ChannelOutput) ToChannelOutput() ChannelOutput {
	return o
}

func (o ChannelOutput) ToChannelOutputWithContext(ctx context.Context) ChannelOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ChannelOutput{})
}
