// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180501

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An Azure resource which represents access to a suite of Maps REST APIs.
//
// ## Example Usage
// ### CreateAccount
//
// ```go
// package main
//
// import (
// 	maps "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/maps/v20180501"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := maps.NewAccount(ctx, "account", &maps.AccountArgs{
// 			AccountName:       pulumi.String("myMapsAccount"),
// 			Location:          pulumi.String("global"),
// 			ResourceGroupName: pulumi.String("myResourceGroup"),
// 			Sku: &maps.SkuArgs{
// 				Name: pulumi.String("S0"),
// 			},
// 			Tags: pulumi.StringMap{
// 				"test": pulumi.String("true"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type Account struct {
	pulumi.CustomResourceState

	// The location of the resource.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the Maps Account, which is unique within a Resource Group.
	Name pulumi.StringOutput `pulumi:"name"`
	// The map account properties.
	Properties MapsAccountPropertiesResponseOutput `pulumi:"properties"`
	// The SKU of this account.
	Sku SkuResponseOutput `pulumi:"sku"`
	// Gets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Azure resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAccount registers a new resource with the given unique name, arguments, and options.
func NewAccount(ctx *pulumi.Context,
	name string, args *AccountArgs, opts ...pulumi.ResourceOption) (*Account, error) {
	if args == nil || args.AccountName == nil {
		return nil, errors.New("missing required argument 'AccountName'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Sku == nil {
		return nil, errors.New("missing required argument 'Sku'")
	}
	if args == nil {
		args = &AccountArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:maps/latest:Account"),
		},
		{
			Type: pulumi.String("azurerm:maps/v20170101preview:Account"),
		},
		{
			Type: pulumi.String("azurerm:maps/v20200201preview:Account"),
		},
	})
	opts = append(opts, aliases)
	var resource Account
	err := ctx.RegisterResource("azurerm:maps/v20180501:Account", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccount gets an existing Account resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountState, opts ...pulumi.ResourceOption) (*Account, error) {
	var resource Account
	err := ctx.ReadResource("azurerm:maps/v20180501:Account", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Account resources.
type accountState struct {
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The name of the Maps Account, which is unique within a Resource Group.
	Name *string `pulumi:"name"`
	// The map account properties.
	Properties *MapsAccountPropertiesResponse `pulumi:"properties"`
	// The SKU of this account.
	Sku *SkuResponse `pulumi:"sku"`
	// Gets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters.
	Tags map[string]string `pulumi:"tags"`
	// Azure resource type.
	Type *string `pulumi:"type"`
}

type AccountState struct {
	// The location of the resource.
	Location pulumi.StringPtrInput
	// The name of the Maps Account, which is unique within a Resource Group.
	Name pulumi.StringPtrInput
	// The map account properties.
	Properties MapsAccountPropertiesResponsePtrInput
	// The SKU of this account.
	Sku SkuResponsePtrInput
	// Gets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters.
	Tags pulumi.StringMapInput
	// Azure resource type.
	Type pulumi.StringPtrInput
}

func (AccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountState)(nil)).Elem()
}

type accountArgs struct {
	// The name of the Maps Account.
	AccountName string `pulumi:"accountName"`
	// The location of the resource.
	Location string `pulumi:"location"`
	// The name of the Azure Resource Group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The SKU of this account.
	Sku Sku `pulumi:"sku"`
	// Gets or sets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Account resource.
type AccountArgs struct {
	// The name of the Maps Account.
	AccountName pulumi.StringInput
	// The location of the resource.
	Location pulumi.StringInput
	// The name of the Azure Resource Group.
	ResourceGroupName pulumi.StringInput
	// The SKU of this account.
	Sku SkuInput
	// Gets or sets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters.
	Tags pulumi.StringMapInput
}

func (AccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountArgs)(nil)).Elem()
}
