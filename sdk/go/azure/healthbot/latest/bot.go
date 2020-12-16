// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// HealthBot resource definition
// Latest API Version: 2020-12-08.
type Bot struct {
	pulumi.CustomResourceState

	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The set of properties specific to healthcare bot resource.
	Properties HealthBotPropertiesResponseOutput `pulumi:"properties"`
	// SKU of the HealthBot.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Metadata pertaining to creation and last modification of the resource
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewBot registers a new resource with the given unique name, arguments, and options.
func NewBot(ctx *pulumi.Context,
	name string, args *BotArgs, opts ...pulumi.ResourceOption) (*Bot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BotName == nil {
		return nil, errors.New("invalid value for required argument 'BotName'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:healthbot/v20201020:Bot"),
		},
		{
			Type: pulumi.String("azure-nextgen:healthbot/v20201020preview:Bot"),
		},
		{
			Type: pulumi.String("azure-nextgen:healthbot/v20201208:Bot"),
		},
		{
			Type: pulumi.String("azure-nextgen:healthbot/v20201208preview:Bot"),
		},
	})
	opts = append(opts, aliases)
	var resource Bot
	err := ctx.RegisterResource("azure-nextgen:healthbot/latest:Bot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBot gets an existing Bot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BotState, opts ...pulumi.ResourceOption) (*Bot, error) {
	var resource Bot
	err := ctx.ReadResource("azure-nextgen:healthbot/latest:Bot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Bot resources.
type botState struct {
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// The set of properties specific to healthcare bot resource.
	Properties *HealthBotPropertiesResponse `pulumi:"properties"`
	// SKU of the HealthBot.
	Sku *SkuResponse `pulumi:"sku"`
	// Metadata pertaining to creation and last modification of the resource
	SystemData *SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

type BotState struct {
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// The set of properties specific to healthcare bot resource.
	Properties HealthBotPropertiesResponsePtrInput
	// SKU of the HealthBot.
	Sku SkuResponsePtrInput
	// Metadata pertaining to creation and last modification of the resource
	SystemData SystemDataResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The type of the resource.
	Type pulumi.StringPtrInput
}

func (BotState) ElementType() reflect.Type {
	return reflect.TypeOf((*botState)(nil)).Elem()
}

type botArgs struct {
	// The name of the Bot resource.
	BotName string `pulumi:"botName"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the Bot resource group in the user subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// SKU of the HealthBot.
	Sku *Sku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Bot resource.
type BotArgs struct {
	// The name of the Bot resource.
	BotName pulumi.StringInput
	// The geo-location where the resource lives
	Location pulumi.StringInput
	// The name of the Bot resource group in the user subscription.
	ResourceGroupName pulumi.StringInput
	// SKU of the HealthBot.
	Sku SkuPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (BotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*botArgs)(nil)).Elem()
}

type BotInput interface {
	pulumi.Input

	ToBotOutput() BotOutput
	ToBotOutputWithContext(ctx context.Context) BotOutput
}

func (Bot) ElementType() reflect.Type {
	return reflect.TypeOf((*Bot)(nil)).Elem()
}

func (i Bot) ToBotOutput() BotOutput {
	return i.ToBotOutputWithContext(context.Background())
}

func (i Bot) ToBotOutputWithContext(ctx context.Context) BotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BotOutput)
}

type BotOutput struct {
	*pulumi.OutputState
}

func (BotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BotOutput)(nil)).Elem()
}

func (o BotOutput) ToBotOutput() BotOutput {
	return o
}

func (o BotOutput) ToBotOutputWithContext(ctx context.Context) BotOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BotOutput{})
}
