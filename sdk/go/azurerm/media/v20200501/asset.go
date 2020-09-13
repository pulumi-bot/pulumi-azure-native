// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200501

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An Asset.
//
// ## Example Usage
// ### Create an Asset
//
// ```go
// package main
//
// import (
// 	media "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/media/v20200501"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := media.NewAsset(ctx, "asset", &media.AssetArgs{
// 			AccountName:        pulumi.String("contosomedia"),
// 			AssetName:          pulumi.String("ClimbingMountLogan"),
// 			Description:        pulumi.String("A documentary showing the ascent of Mount Logan"),
// 			ResourceGroupName:  pulumi.String("contoso"),
// 			StorageAccountName: pulumi.String("storage0"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type Asset struct {
	pulumi.CustomResourceState

	// The alternate ID of the Asset.
	AlternateId pulumi.StringPtrOutput `pulumi:"alternateId"`
	// The Asset ID.
	AssetId pulumi.StringOutput `pulumi:"assetId"`
	// The name of the asset blob container.
	Container pulumi.StringPtrOutput `pulumi:"container"`
	// The creation date of the Asset.
	Created pulumi.StringOutput `pulumi:"created"`
	// The Asset description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The last modified date of the Asset.
	LastModified pulumi.StringOutput `pulumi:"lastModified"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the storage account.
	StorageAccountName pulumi.StringPtrOutput `pulumi:"storageAccountName"`
	// The Asset encryption format. One of None or MediaStorageEncryption.
	StorageEncryptionFormat pulumi.StringOutput `pulumi:"storageEncryptionFormat"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAsset registers a new resource with the given unique name, arguments, and options.
func NewAsset(ctx *pulumi.Context,
	name string, args *AssetArgs, opts ...pulumi.ResourceOption) (*Asset, error) {
	if args == nil || args.AccountName == nil {
		return nil, errors.New("missing required argument 'AccountName'")
	}
	if args == nil || args.AssetName == nil {
		return nil, errors.New("missing required argument 'AssetName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &AssetArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:media/latest:Asset"),
		},
		{
			Type: pulumi.String("azurerm:media/v20180330preview:Asset"),
		},
		{
			Type: pulumi.String("azurerm:media/v20180601preview:Asset"),
		},
		{
			Type: pulumi.String("azurerm:media/v20180701:Asset"),
		},
	})
	opts = append(opts, aliases)
	var resource Asset
	err := ctx.RegisterResource("azurerm:media/v20200501:Asset", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAsset gets an existing Asset resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAsset(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AssetState, opts ...pulumi.ResourceOption) (*Asset, error) {
	var resource Asset
	err := ctx.ReadResource("azurerm:media/v20200501:Asset", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Asset resources.
type assetState struct {
	// The alternate ID of the Asset.
	AlternateId *string `pulumi:"alternateId"`
	// The Asset ID.
	AssetId *string `pulumi:"assetId"`
	// The name of the asset blob container.
	Container *string `pulumi:"container"`
	// The creation date of the Asset.
	Created *string `pulumi:"created"`
	// The Asset description.
	Description *string `pulumi:"description"`
	// The last modified date of the Asset.
	LastModified *string `pulumi:"lastModified"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// The name of the storage account.
	StorageAccountName *string `pulumi:"storageAccountName"`
	// The Asset encryption format. One of None or MediaStorageEncryption.
	StorageEncryptionFormat *string `pulumi:"storageEncryptionFormat"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `pulumi:"type"`
}

type AssetState struct {
	// The alternate ID of the Asset.
	AlternateId pulumi.StringPtrInput
	// The Asset ID.
	AssetId pulumi.StringPtrInput
	// The name of the asset blob container.
	Container pulumi.StringPtrInput
	// The creation date of the Asset.
	Created pulumi.StringPtrInput
	// The Asset description.
	Description pulumi.StringPtrInput
	// The last modified date of the Asset.
	LastModified pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// The name of the storage account.
	StorageAccountName pulumi.StringPtrInput
	// The Asset encryption format. One of None or MediaStorageEncryption.
	StorageEncryptionFormat pulumi.StringPtrInput
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringPtrInput
}

func (AssetState) ElementType() reflect.Type {
	return reflect.TypeOf((*assetState)(nil)).Elem()
}

type assetArgs struct {
	// The Media Services account name.
	AccountName string `pulumi:"accountName"`
	// The alternate ID of the Asset.
	AlternateId *string `pulumi:"alternateId"`
	// The Asset name.
	AssetName string `pulumi:"assetName"`
	// The name of the asset blob container.
	Container *string `pulumi:"container"`
	// The Asset description.
	Description *string `pulumi:"description"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the storage account.
	StorageAccountName *string `pulumi:"storageAccountName"`
}

// The set of arguments for constructing a Asset resource.
type AssetArgs struct {
	// The Media Services account name.
	AccountName pulumi.StringInput
	// The alternate ID of the Asset.
	AlternateId pulumi.StringPtrInput
	// The Asset name.
	AssetName pulumi.StringInput
	// The name of the asset blob container.
	Container pulumi.StringPtrInput
	// The Asset description.
	Description pulumi.StringPtrInput
	// The name of the resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// The name of the storage account.
	StorageAccountName pulumi.StringPtrInput
}

func (AssetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*assetArgs)(nil)).Elem()
}
