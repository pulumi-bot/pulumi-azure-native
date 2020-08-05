// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190501

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
	// The integration account properties.
	Properties IntegrationAccountPropertiesResponseOutput `pulumi:"properties"`
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
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &IntegrationAccountArgs{}
	}
	var resource IntegrationAccount
	err := ctx.RegisterResource("azurerm:logic/v20190501:IntegrationAccount", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azurerm:logic/v20190501:IntegrationAccount", name, id, state, &resource, opts...)
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
	// The integration account properties.
	Properties *IntegrationAccountPropertiesResponse `pulumi:"properties"`
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
	// The integration account properties.
	Properties IntegrationAccountPropertiesResponsePtrInput
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
	// The integration service environment.
	IntegrationServiceEnvironment *IntegrationServiceEnvironmentType `pulumi:"integrationServiceEnvironment"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The integration account name.
	Name string `pulumi:"name"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The sku.
	Sku *IntegrationAccountSku `pulumi:"sku"`
	// The workflow state.
	State *string `pulumi:"state"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a IntegrationAccount resource.
type IntegrationAccountArgs struct {
	// The integration service environment.
	IntegrationServiceEnvironment IntegrationServiceEnvironmentTypePtrInput
	// The resource location.
	Location pulumi.StringPtrInput
	// The integration account name.
	Name pulumi.StringInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The sku.
	Sku IntegrationAccountSkuPtrInput
	// The workflow state.
	State pulumi.StringPtrInput
	// The resource tags.
	Tags pulumi.StringMapInput
}

func (IntegrationAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountArgs)(nil)).Elem()
}
