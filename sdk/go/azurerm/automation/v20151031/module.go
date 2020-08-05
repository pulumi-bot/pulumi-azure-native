// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20151031

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Definition of the module type.
type Module struct {
	pulumi.CustomResourceState

	// Gets or sets the etag of the resource.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The Azure Region where the resource lives
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets or sets the module properties.
	Properties ModulePropertiesResponseOutput `pulumi:"properties"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewModule registers a new resource with the given unique name, arguments, and options.
func NewModule(ctx *pulumi.Context,
	name string, args *ModuleArgs, opts ...pulumi.ResourceOption) (*Module, error) {
	if args == nil || args.AutomationAccountName == nil {
		return nil, errors.New("missing required argument 'AutomationAccountName'")
	}
	if args == nil || args.ContentLink == nil {
		return nil, errors.New("missing required argument 'ContentLink'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ModuleArgs{}
	}
	var resource Module
	err := ctx.RegisterResource("azurerm:automation/v20151031:Module", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetModule gets an existing Module resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetModule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ModuleState, opts ...pulumi.ResourceOption) (*Module, error) {
	var resource Module
	err := ctx.ReadResource("azurerm:automation/v20151031:Module", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Module resources.
type moduleState struct {
	// Gets or sets the etag of the resource.
	Etag *string `pulumi:"etag"`
	// The Azure Region where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// Gets or sets the module properties.
	Properties *ModulePropertiesResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

type ModuleState struct {
	// Gets or sets the etag of the resource.
	Etag pulumi.StringPtrInput
	// The Azure Region where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// Gets or sets the module properties.
	Properties ModulePropertiesResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The type of the resource.
	Type pulumi.StringPtrInput
}

func (ModuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*moduleState)(nil)).Elem()
}

type moduleArgs struct {
	// The name of the automation account.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// Gets or sets the module content link.
	ContentLink ContentLink `pulumi:"contentLink"`
	// Gets or sets the location of the resource.
	Location *string `pulumi:"location"`
	// The name of module.
	Name string `pulumi:"name"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Gets or sets the tags attached to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Module resource.
type ModuleArgs struct {
	// The name of the automation account.
	AutomationAccountName pulumi.StringInput
	// Gets or sets the module content link.
	ContentLink ContentLinkInput
	// Gets or sets the location of the resource.
	Location pulumi.StringPtrInput
	// The name of module.
	Name pulumi.StringInput
	// Name of an Azure Resource group.
	ResourceGroupName pulumi.StringInput
	// Gets or sets the tags attached to the resource.
	Tags pulumi.StringMapInput
}

func (ModuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*moduleArgs)(nil)).Elem()
}
