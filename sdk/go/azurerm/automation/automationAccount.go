// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package automation

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Definition of the automation account type.
type AutomationAccount struct {
	pulumi.CustomResourceState

	// Gets or sets the etag of the resource.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The Azure Region where the resource lives
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets or sets the automation account properties.
	Properties AutomationAccountPropertiesResponseOutput `pulumi:"properties"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAutomationAccount registers a new resource with the given unique name, arguments, and options.
func NewAutomationAccount(ctx *pulumi.Context,
	name string, args *AutomationAccountArgs, opts ...pulumi.ResourceOption) (*AutomationAccount, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &AutomationAccountArgs{}
	}
	var resource AutomationAccount
	err := ctx.RegisterResource("azurerm:automation:AutomationAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAutomationAccount gets an existing AutomationAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAutomationAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AutomationAccountState, opts ...pulumi.ResourceOption) (*AutomationAccount, error) {
	var resource AutomationAccount
	err := ctx.ReadResource("azurerm:automation:AutomationAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AutomationAccount resources.
type automationAccountState struct {
	// Gets or sets the etag of the resource.
	Etag *string `pulumi:"etag"`
	// The Azure Region where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// Gets or sets the automation account properties.
	Properties *AutomationAccountPropertiesResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

type AutomationAccountState struct {
	// Gets or sets the etag of the resource.
	Etag pulumi.StringPtrInput
	// The Azure Region where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// Gets or sets the automation account properties.
	Properties AutomationAccountPropertiesResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The type of the resource.
	Type pulumi.StringPtrInput
}

func (AutomationAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*automationAccountState)(nil)).Elem()
}

type automationAccountArgs struct {
	// Gets or sets the location of the resource.
	Location *string `pulumi:"location"`
	// The name of the automation account.
	Name string `pulumi:"name"`
	// Gets or sets account create or update properties.
	Properties *AutomationAccountCreateOrUpdateProperties `pulumi:"properties"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Gets or sets the tags attached to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a AutomationAccount resource.
type AutomationAccountArgs struct {
	// Gets or sets the location of the resource.
	Location pulumi.StringPtrInput
	// The name of the automation account.
	Name pulumi.StringInput
	// Gets or sets account create or update properties.
	Properties AutomationAccountCreateOrUpdatePropertiesPtrInput
	// Name of an Azure Resource group.
	ResourceGroupName pulumi.StringInput
	// Gets or sets the tags attached to the resource.
	Tags pulumi.StringMapInput
}

func (AutomationAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*automationAccountArgs)(nil)).Elem()
}
