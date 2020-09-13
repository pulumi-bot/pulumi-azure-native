// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Bot channel resource definition
//
// ## Example Usage
// ### Create Connection Setting
//
// ```go
// package main
//
// import (
// 	botservice "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/botservice/latest"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := botservice.NewBotConnection(ctx, "botConnection", &botservice.BotConnectionArgs{
// 			ConnectionName:    pulumi.String("sampleConnection"),
// 			Etag:              pulumi.String("etag1"),
// 			Location:          pulumi.String("West US"),
// 			ResourceGroupName: pulumi.String("OneResourceGroupName"),
// 			ResourceName:      pulumi.String("samplebotname"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type BotConnection struct {
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
	Properties ConnectionSettingPropertiesResponseOutput `pulumi:"properties"`
	// Gets or sets the SKU of the resource.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Contains resource tags defined as key/value pairs.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewBotConnection registers a new resource with the given unique name, arguments, and options.
func NewBotConnection(ctx *pulumi.Context,
	name string, args *BotConnectionArgs, opts ...pulumi.ResourceOption) (*BotConnection, error) {
	if args == nil || args.ConnectionName == nil {
		return nil, errors.New("missing required argument 'ConnectionName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ResourceName == nil {
		return nil, errors.New("missing required argument 'ResourceName'")
	}
	if args == nil {
		args = &BotConnectionArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:botservice/v20171201:BotConnection"),
		},
		{
			Type: pulumi.String("azurerm:botservice/v20180712:BotConnection"),
		},
		{
			Type: pulumi.String("azurerm:botservice/v20200602:BotConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource BotConnection
	err := ctx.RegisterResource("azurerm:botservice/latest:BotConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBotConnection gets an existing BotConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBotConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BotConnectionState, opts ...pulumi.ResourceOption) (*BotConnection, error) {
	var resource BotConnection
	err := ctx.ReadResource("azurerm:botservice/latest:BotConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BotConnection resources.
type botConnectionState struct {
	// Entity Tag
	Etag *string `pulumi:"etag"`
	// Required. Gets or sets the Kind of the resource.
	Kind *string `pulumi:"kind"`
	// Specifies the location of the resource.
	Location *string `pulumi:"location"`
	// Specifies the name of the resource.
	Name *string `pulumi:"name"`
	// The set of properties specific to bot channel resource
	Properties *ConnectionSettingPropertiesResponse `pulumi:"properties"`
	// Gets or sets the SKU of the resource.
	Sku *SkuResponse `pulumi:"sku"`
	// Contains resource tags defined as key/value pairs.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the type of the resource.
	Type *string `pulumi:"type"`
}

type BotConnectionState struct {
	// Entity Tag
	Etag pulumi.StringPtrInput
	// Required. Gets or sets the Kind of the resource.
	Kind pulumi.StringPtrInput
	// Specifies the location of the resource.
	Location pulumi.StringPtrInput
	// Specifies the name of the resource.
	Name pulumi.StringPtrInput
	// The set of properties specific to bot channel resource
	Properties ConnectionSettingPropertiesResponsePtrInput
	// Gets or sets the SKU of the resource.
	Sku SkuResponsePtrInput
	// Contains resource tags defined as key/value pairs.
	Tags pulumi.StringMapInput
	// Specifies the type of the resource.
	Type pulumi.StringPtrInput
}

func (BotConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*botConnectionState)(nil)).Elem()
}

type botConnectionArgs struct {
	// The name of the Bot Service Connection Setting resource.
	ConnectionName string `pulumi:"connectionName"`
	// Entity Tag
	Etag *string `pulumi:"etag"`
	// Required. Gets or sets the Kind of the resource.
	Kind *string `pulumi:"kind"`
	// Specifies the location of the resource.
	Location *string `pulumi:"location"`
	// The set of properties specific to bot channel resource
	Properties *ConnectionSettingProperties `pulumi:"properties"`
	// The name of the Bot resource group in the user subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Bot resource.
	ResourceName string `pulumi:"resourceName"`
	// Gets or sets the SKU of the resource.
	Sku *Sku `pulumi:"sku"`
	// Contains resource tags defined as key/value pairs.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a BotConnection resource.
type BotConnectionArgs struct {
	// The name of the Bot Service Connection Setting resource.
	ConnectionName pulumi.StringInput
	// Entity Tag
	Etag pulumi.StringPtrInput
	// Required. Gets or sets the Kind of the resource.
	Kind pulumi.StringPtrInput
	// Specifies the location of the resource.
	Location pulumi.StringPtrInput
	// The set of properties specific to bot channel resource
	Properties ConnectionSettingPropertiesPtrInput
	// The name of the Bot resource group in the user subscription.
	ResourceGroupName pulumi.StringInput
	// The name of the Bot resource.
	ResourceName pulumi.StringInput
	// Gets or sets the SKU of the resource.
	Sku SkuPtrInput
	// Contains resource tags defined as key/value pairs.
	Tags pulumi.StringMapInput
}

func (BotConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*botConnectionArgs)(nil)).Elem()
}
