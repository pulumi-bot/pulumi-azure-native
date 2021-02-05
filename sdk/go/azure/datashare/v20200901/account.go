// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An account data transfer object.
type Account struct {
	pulumi.CustomResourceState

	// Time at which the account was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Identity Info on the Account
	Identity IdentityResponseOutput `pulumi:"identity"`
	// Location of the azure resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Name of the azure resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of the Account
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// System Data of the Azure resource.
	SystemData DefaultDtoResponseSystemDataOutput `pulumi:"systemData"`
	// Tags on the azure resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Type of the azure resource
	Type pulumi.StringOutput `pulumi:"type"`
	// Email of the user who created the resource
	UserEmail pulumi.StringOutput `pulumi:"userEmail"`
	// Name of the user who created the resource
	UserName pulumi.StringOutput `pulumi:"userName"`
}

// NewAccount registers a new resource with the given unique name, arguments, and options.
func NewAccount(ctx *pulumi.Context,
	name string, args *AccountArgs, opts ...pulumi.ResourceOption) (*Account, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.Identity == nil {
		return nil, errors.New("invalid value for required argument 'Identity'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:datashare/latest:Account"),
		},
		{
			Type: pulumi.String("azure-nextgen:datashare/v20181101preview:Account"),
		},
		{
			Type: pulumi.String("azure-nextgen:datashare/v20191101:Account"),
		},
	})
	opts = append(opts, aliases)
	var resource Account
	err := ctx.RegisterResource("azure-nextgen:datashare/v20200901:Account", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-nextgen:datashare/v20200901:Account", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Account resources.
type accountState struct {
	// Time at which the account was created.
	CreatedAt *string `pulumi:"createdAt"`
	// Identity Info on the Account
	Identity *IdentityResponse `pulumi:"identity"`
	// Location of the azure resource.
	Location *string `pulumi:"location"`
	// Name of the azure resource
	Name *string `pulumi:"name"`
	// Provisioning state of the Account
	ProvisioningState *string `pulumi:"provisioningState"`
	// System Data of the Azure resource.
	SystemData *DefaultDtoResponseSystemData `pulumi:"systemData"`
	// Tags on the azure resource.
	Tags map[string]string `pulumi:"tags"`
	// Type of the azure resource
	Type *string `pulumi:"type"`
	// Email of the user who created the resource
	UserEmail *string `pulumi:"userEmail"`
	// Name of the user who created the resource
	UserName *string `pulumi:"userName"`
}

type AccountState struct {
	// Time at which the account was created.
	CreatedAt pulumi.StringPtrInput
	// Identity Info on the Account
	Identity IdentityResponsePtrInput
	// Location of the azure resource.
	Location pulumi.StringPtrInput
	// Name of the azure resource
	Name pulumi.StringPtrInput
	// Provisioning state of the Account
	ProvisioningState pulumi.StringPtrInput
	// System Data of the Azure resource.
	SystemData DefaultDtoResponseSystemDataPtrInput
	// Tags on the azure resource.
	Tags pulumi.StringMapInput
	// Type of the azure resource
	Type pulumi.StringPtrInput
	// Email of the user who created the resource
	UserEmail pulumi.StringPtrInput
	// Name of the user who created the resource
	UserName pulumi.StringPtrInput
}

func (AccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountState)(nil)).Elem()
}

type accountArgs struct {
	// The name of the share account.
	AccountName string `pulumi:"accountName"`
	// Identity Info on the Account
	Identity Identity `pulumi:"identity"`
	// Location of the azure resource.
	Location *string `pulumi:"location"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Tags on the azure resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Account resource.
type AccountArgs struct {
	// The name of the share account.
	AccountName pulumi.StringInput
	// Identity Info on the Account
	Identity IdentityInput
	// Location of the azure resource.
	Location pulumi.StringPtrInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// Tags on the azure resource.
	Tags pulumi.StringMapInput
}

func (AccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountArgs)(nil)).Elem()
}

type AccountInput interface {
	pulumi.Input

	ToAccountOutput() AccountOutput
	ToAccountOutputWithContext(ctx context.Context) AccountOutput
}

func (*Account) ElementType() reflect.Type {
	return reflect.TypeOf((*Account)(nil))
}

func (i *Account) ToAccountOutput() AccountOutput {
	return i.ToAccountOutputWithContext(context.Background())
}

func (i *Account) ToAccountOutputWithContext(ctx context.Context) AccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountOutput)
}

type AccountOutput struct {
	*pulumi.OutputState
}

func (AccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Account)(nil))
}

func (o AccountOutput) ToAccountOutput() AccountOutput {
	return o
}

func (o AccountOutput) ToAccountOutputWithContext(ctx context.Context) AccountOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AccountOutput{})
}
