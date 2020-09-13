// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190901

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Resource information with extended details.
//
// ## Example Usage
// ### Create a new vault or update an existing vault
//
// ```go
// package main
//
// import (
// 	keyvault "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/keyvault/v20190901"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := keyvault.NewVault(ctx, "vault", &keyvault.VaultArgs{
// 			Location:          pulumi.String("westus"),
// 			ResourceGroupName: pulumi.String("sample-resource-group"),
// 			VaultName:         pulumi.String("sample-vault"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
// ### Create or update a vault with network acls
//
// ```go
// package main
//
// import (
// 	keyvault "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/keyvault/v20190901"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := keyvault.NewVault(ctx, "vault", &keyvault.VaultArgs{
// 			Location:          pulumi.String("westus"),
// 			ResourceGroupName: pulumi.String("sample-resource-group"),
// 			VaultName:         pulumi.String("sample-vault"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type Vault struct {
	pulumi.CustomResourceState

	// Azure location of the key vault resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Name of the key vault resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the vault
	Properties VaultPropertiesResponseOutput `pulumi:"properties"`
	// Tags assigned to the key vault resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type of the key vault resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewVault registers a new resource with the given unique name, arguments, and options.
func NewVault(ctx *pulumi.Context,
	name string, args *VaultArgs, opts ...pulumi.ResourceOption) (*Vault, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Properties == nil {
		return nil, errors.New("missing required argument 'Properties'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.VaultName == nil {
		return nil, errors.New("missing required argument 'VaultName'")
	}
	if args == nil {
		args = &VaultArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:keyvault/latest:Vault"),
		},
		{
			Type: pulumi.String("azurerm:keyvault/v20150601:Vault"),
		},
		{
			Type: pulumi.String("azurerm:keyvault/v20161001:Vault"),
		},
		{
			Type: pulumi.String("azurerm:keyvault/v20180214:Vault"),
		},
		{
			Type: pulumi.String("azurerm:keyvault/v20180214preview:Vault"),
		},
		{
			Type: pulumi.String("azurerm:keyvault/v20200401preview:Vault"),
		},
	})
	opts = append(opts, aliases)
	var resource Vault
	err := ctx.RegisterResource("azurerm:keyvault/v20190901:Vault", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVault gets an existing Vault resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVault(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VaultState, opts ...pulumi.ResourceOption) (*Vault, error) {
	var resource Vault
	err := ctx.ReadResource("azurerm:keyvault/v20190901:Vault", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Vault resources.
type vaultState struct {
	// Azure location of the key vault resource.
	Location *string `pulumi:"location"`
	// Name of the key vault resource.
	Name *string `pulumi:"name"`
	// Properties of the vault
	Properties *VaultPropertiesResponse `pulumi:"properties"`
	// Tags assigned to the key vault resource.
	Tags map[string]string `pulumi:"tags"`
	// Resource type of the key vault resource.
	Type *string `pulumi:"type"`
}

type VaultState struct {
	// Azure location of the key vault resource.
	Location pulumi.StringPtrInput
	// Name of the key vault resource.
	Name pulumi.StringPtrInput
	// Properties of the vault
	Properties VaultPropertiesResponsePtrInput
	// Tags assigned to the key vault resource.
	Tags pulumi.StringMapInput
	// Resource type of the key vault resource.
	Type pulumi.StringPtrInput
}

func (VaultState) ElementType() reflect.Type {
	return reflect.TypeOf((*vaultState)(nil)).Elem()
}

type vaultArgs struct {
	// The supported Azure location where the key vault should be created.
	Location string `pulumi:"location"`
	// Properties of the vault
	Properties VaultProperties `pulumi:"properties"`
	// The name of the Resource Group to which the server belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The tags that will be assigned to the key vault.
	Tags map[string]string `pulumi:"tags"`
	// Name of the vault
	VaultName string `pulumi:"vaultName"`
}

// The set of arguments for constructing a Vault resource.
type VaultArgs struct {
	// The supported Azure location where the key vault should be created.
	Location pulumi.StringInput
	// Properties of the vault
	Properties VaultPropertiesInput
	// The name of the Resource Group to which the server belongs.
	ResourceGroupName pulumi.StringInput
	// The tags that will be assigned to the key vault.
	Tags pulumi.StringMapInput
	// Name of the vault
	VaultName pulumi.StringInput
}

func (VaultArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vaultArgs)(nil)).Elem()
}
