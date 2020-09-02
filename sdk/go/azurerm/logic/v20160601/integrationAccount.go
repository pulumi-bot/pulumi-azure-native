// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160601

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The integration account.
type IntegrationAccount struct {
	pulumi.CustomResourceState

	// The resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Gets the resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The sku.
	Sku IntegrationAccountSkuResponsePtrOutput `pulumi:"sku"`
	// The resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Gets the resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewIntegrationAccount registers a new resource with the given unique name, arguments, and options.
func NewIntegrationAccount(ctx *pulumi.Context,
	name string, args *IntegrationAccountArgs, opts ...pulumi.ResourceOption) (*IntegrationAccount, error) {
	if args == nil || args.IntegrationAccountName == nil {
		return nil, errors.New("missing required argument 'IntegrationAccountName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &IntegrationAccountArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:logic/latest:IntegrationAccount"),
		},
		{
			Type: pulumi.String("azurerm:logic/v20150801preview:IntegrationAccount"),
		},
		{
			Type: pulumi.String("azurerm:logic/v20180701preview:IntegrationAccount"),
		},
		{
			Type: pulumi.String("azurerm:logic/v20190501:IntegrationAccount"),
		},
	})
	opts = append(opts, aliases)
	var resource IntegrationAccount
	err := ctx.RegisterResource("azurerm:logic/v20160601:IntegrationAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntegrationAccount gets an existing IntegrationAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegrationAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationAccountState, opts ...pulumi.ResourceOption) (*IntegrationAccount, error) {
	var resource IntegrationAccount
	err := ctx.ReadResource("azurerm:logic/v20160601:IntegrationAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IntegrationAccount resources.
type integrationAccountState struct {
	// The resource location.
	Location *string `pulumi:"location"`
	// Gets the resource name.
	Name *string `pulumi:"name"`
	// The sku.
	Sku *IntegrationAccountSkuResponse `pulumi:"sku"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Gets the resource type.
	Type *string `pulumi:"type"`
}

type IntegrationAccountState struct {
	// The resource location.
	Location pulumi.StringPtrInput
	// Gets the resource name.
	Name pulumi.StringPtrInput
	// The sku.
	Sku IntegrationAccountSkuResponsePtrInput
	// The resource tags.
	Tags pulumi.StringMapInput
	// Gets the resource type.
	Type pulumi.StringPtrInput
}

func (IntegrationAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountState)(nil)).Elem()
}

type integrationAccountArgs struct {
	// The integration account name.
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The sku.
	Sku *IntegrationAccountSku `pulumi:"sku"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a IntegrationAccount resource.
type IntegrationAccountArgs struct {
	// The integration account name.
	IntegrationAccountName pulumi.StringInput
	// The resource location.
	Location pulumi.StringPtrInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The sku.
	Sku IntegrationAccountSkuPtrInput
	// The resource tags.
	Tags pulumi.StringMapInput
}

func (IntegrationAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountArgs)(nil)).Elem()
}
