// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180901preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The EngagementFabric account
type Account struct {
	pulumi.CustomResourceState

	// The location of the resource
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The SKU of the resource
	Sku SKUResponseOutput `pulumi:"sku"`
	// The tags of the resource
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The fully qualified type of the resource
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
	var resource Account
	err := ctx.RegisterResource("azurerm:engagementfabric/v20180901preview:Account", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azurerm:engagementfabric/v20180901preview:Account", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Account resources.
type accountState struct {
	// The location of the resource
	Location *string `pulumi:"location"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// The SKU of the resource
	Sku *SKUResponse `pulumi:"sku"`
	// The tags of the resource
	Tags map[string]string `pulumi:"tags"`
	// The fully qualified type of the resource
	Type *string `pulumi:"type"`
}

type AccountState struct {
	// The location of the resource
	Location pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// The SKU of the resource
	Sku SKUResponsePtrInput
	// The tags of the resource
	Tags pulumi.StringMapInput
	// The fully qualified type of the resource
	Type pulumi.StringPtrInput
}

func (AccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountState)(nil)).Elem()
}

type accountArgs struct {
	// Account Name
	AccountName string `pulumi:"accountName"`
	// The location of the resource
	Location string `pulumi:"location"`
	// Resource Group Name
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The SKU of the resource
	Sku SKU `pulumi:"sku"`
	// The tags of the resource
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Account resource.
type AccountArgs struct {
	// Account Name
	AccountName pulumi.StringInput
	// The location of the resource
	Location pulumi.StringInput
	// Resource Group Name
	ResourceGroupName pulumi.StringInput
	// The SKU of the resource
	Sku SKUInput
	// The tags of the resource
	Tags pulumi.StringMapInput
}

func (AccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountArgs)(nil)).Elem()
}
