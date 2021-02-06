// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180214preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Resource information with extended details.
type Vault struct {
	pulumi.CustomResourceState

	// The supported Azure location where the key vault should be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the key vault.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the vault
	Properties VaultPropertiesResponseOutput `pulumi:"properties"`
	// The tags that will be assigned to the key vault.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The resource type of the key vault.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewVault registers a new resource with the given unique name, arguments, and options.
func NewVault(ctx *pulumi.Context,
	name string, args *VaultArgs, opts ...pulumi.ResourceOption) (*Vault, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VaultName == nil {
		return nil, errors.New("invalid value for required argument 'VaultName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:keyvault/latest:Vault"),
		},
		{
			Type: pulumi.String("azure-nextgen:keyvault/v20150601:Vault"),
		},
		{
			Type: pulumi.String("azure-nextgen:keyvault/v20161001:Vault"),
		},
		{
			Type: pulumi.String("azure-nextgen:keyvault/v20180214:Vault"),
		},
		{
			Type: pulumi.String("azure-nextgen:keyvault/v20190901:Vault"),
		},
		{
			Type: pulumi.String("azure-nextgen:keyvault/v20200401preview:Vault"),
		},
	})
	opts = append(opts, aliases)
	var resource Vault
	err := ctx.RegisterResource("azure-nextgen:keyvault/v20180214preview:Vault", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-nextgen:keyvault/v20180214preview:Vault", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Vault resources.
type vaultState struct {
	// The supported Azure location where the key vault should be created.
	Location *string `pulumi:"location"`
	// The name of the key vault.
	Name *string `pulumi:"name"`
	// Properties of the vault
	Properties *VaultPropertiesResponse `pulumi:"properties"`
	// The tags that will be assigned to the key vault.
	Tags map[string]string `pulumi:"tags"`
	// The resource type of the key vault.
	Type *string `pulumi:"type"`
}

type VaultState struct {
	// The supported Azure location where the key vault should be created.
	Location pulumi.StringPtrInput
	// The name of the key vault.
	Name pulumi.StringPtrInput
	// Properties of the vault
	Properties VaultPropertiesResponsePtrInput
	// The tags that will be assigned to the key vault.
	Tags pulumi.StringMapInput
	// The resource type of the key vault.
	Type pulumi.StringPtrInput
}

func (VaultState) ElementType() reflect.Type {
	return reflect.TypeOf((*vaultState)(nil)).Elem()
}

type vaultArgs struct {
	// The supported Azure location where the key vault should be created.
	Location *string `pulumi:"location"`
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
	Location pulumi.StringPtrInput
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

type VaultInput interface {
	pulumi.Input

	ToVaultOutput() VaultOutput
	ToVaultOutputWithContext(ctx context.Context) VaultOutput
}

func (*Vault) ElementType() reflect.Type {
	return reflect.TypeOf((*Vault)(nil))
}

func (i *Vault) ToVaultOutput() VaultOutput {
	return i.ToVaultOutputWithContext(context.Background())
}

func (i *Vault) ToVaultOutputWithContext(ctx context.Context) VaultOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultOutput)
}

type VaultOutput struct {
	*pulumi.OutputState
}

func (VaultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Vault)(nil))
}

func (o VaultOutput) ToVaultOutput() VaultOutput {
	return o
}

func (o VaultOutput) ToVaultOutputWithContext(ctx context.Context) VaultOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(VaultOutput{})
}
