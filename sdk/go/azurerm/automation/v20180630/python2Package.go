// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180630

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Definition of the module type.
type Python2Package struct {
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

// NewPython2Package registers a new resource with the given unique name, arguments, and options.
func NewPython2Package(ctx *pulumi.Context,
	name string, args *Python2PackageArgs, opts ...pulumi.ResourceOption) (*Python2Package, error) {
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
		args = &Python2PackageArgs{}
	}
	var resource Python2Package
	err := ctx.RegisterResource("azurerm:automation/v20180630:Python2Package", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPython2Package gets an existing Python2Package resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPython2Package(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Python2PackageState, opts ...pulumi.ResourceOption) (*Python2Package, error) {
	var resource Python2Package
	err := ctx.ReadResource("azurerm:automation/v20180630:Python2Package", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Python2Package resources.
type python2PackageState struct {
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

type Python2PackageState struct {
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

func (Python2PackageState) ElementType() reflect.Type {
	return reflect.TypeOf((*python2PackageState)(nil)).Elem()
}

type python2PackageArgs struct {
	// The name of the automation account.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// Gets or sets the module content link.
	ContentLink ContentLink `pulumi:"contentLink"`
	// The name of python package.
	Name string `pulumi:"name"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Gets or sets the tags attached to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Python2Package resource.
type Python2PackageArgs struct {
	// The name of the automation account.
	AutomationAccountName pulumi.StringInput
	// Gets or sets the module content link.
	ContentLink ContentLinkInput
	// The name of python package.
	Name pulumi.StringInput
	// Name of an Azure Resource group.
	ResourceGroupName pulumi.StringInput
	// Gets or sets the tags attached to the resource.
	Tags pulumi.StringMapInput
}

func (Python2PackageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*python2PackageArgs)(nil)).Elem()
}
