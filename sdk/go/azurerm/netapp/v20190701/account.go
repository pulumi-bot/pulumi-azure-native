// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190701

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// NetApp account resource
type Account struct {
	pulumi.CustomResourceState

	// Active Directories
	ActiveDirectories ActiveDirectoryResponseArrayOutput `pulumi:"activeDirectories"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure lifecycle management
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Resource tags
	Tags pulumi.MapOutput `pulumi:"tags"`
	// Resource type
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
	if args == nil {
		args = &AccountArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:netapp/latest:Account"),
		},
		{
			Type: pulumi.String("azurerm:netapp/v20170815:Account"),
		},
		{
			Type: pulumi.String("azurerm:netapp/v20190501:Account"),
		},
		{
			Type: pulumi.String("azurerm:netapp/v20190601:Account"),
		},
		{
			Type: pulumi.String("azurerm:netapp/v20190801:Account"),
		},
		{
			Type: pulumi.String("azurerm:netapp/v20191001:Account"),
		},
		{
			Type: pulumi.String("azurerm:netapp/v20191101:Account"),
		},
		{
			Type: pulumi.String("azurerm:netapp/v20200201:Account"),
		},
		{
			Type: pulumi.String("azurerm:netapp/v20200601:Account"),
		},
	})
	opts = append(opts, aliases)
	var resource Account
	err := ctx.RegisterResource("azurerm:netapp/v20190701:Account", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azurerm:netapp/v20190701:Account", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Account resources.
type accountState struct {
	// Active Directories
	ActiveDirectories []ActiveDirectoryResponse `pulumi:"activeDirectories"`
	// Resource location
	Location *string `pulumi:"location"`
	// Resource name
	Name *string `pulumi:"name"`
	// Azure lifecycle management
	ProvisioningState *string `pulumi:"provisioningState"`
	// Resource tags
	Tags map[string]interface{} `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

type AccountState struct {
	// Active Directories
	ActiveDirectories ActiveDirectoryResponseArrayInput
	// Resource location
	Location pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// Azure lifecycle management
	ProvisioningState pulumi.StringPtrInput
	// Resource tags
	Tags pulumi.MapInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (AccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountState)(nil)).Elem()
}

type accountArgs struct {
	// The name of the NetApp account
	AccountName string `pulumi:"accountName"`
	// Active Directories
	ActiveDirectories []ActiveDirectory `pulumi:"activeDirectories"`
	// Resource location
	Location string `pulumi:"location"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a Account resource.
type AccountArgs struct {
	// The name of the NetApp account
	AccountName pulumi.StringInput
	// Active Directories
	ActiveDirectories ActiveDirectoryArrayInput
	// Resource location
	Location pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags
	Tags pulumi.MapInput
}

func (AccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountArgs)(nil)).Elem()
}
