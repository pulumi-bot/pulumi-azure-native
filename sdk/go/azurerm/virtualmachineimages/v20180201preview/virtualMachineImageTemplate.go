// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180201preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Example Usage
// ### Create an Image Template.
//
// ```go
//
// ```
type VirtualMachineImageTemplate struct {
	pulumi.CustomResourceState

	// Specifies the properties used to describe the customization steps of the image, like Image source etc
	Customize ImageTemplateShellCustomizerResponseArrayOutput `pulumi:"customize"`
	// The distribution targets where the image output needs to go to.
	Distribute pulumi.ArrayOutput `pulumi:"distribute"`
	// State of 'run' that is currently executing or was last executed.
	LastRunStatus ImageTemplateLastRunStatusResponseOutput `pulumi:"lastRunStatus"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning error, if any
	ProvisioningError ProvisioningErrorResponseOutput `pulumi:"provisioningError"`
	// Provisioning state of the resource
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Specifies the properties used to describe the source image.
	Source pulumi.AnyOutput `pulumi:"source"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewVirtualMachineImageTemplate registers a new resource with the given unique name, arguments, and options.
func NewVirtualMachineImageTemplate(ctx *pulumi.Context,
	name string, args *VirtualMachineImageTemplateArgs, opts ...pulumi.ResourceOption) (*VirtualMachineImageTemplate, error) {
	if args == nil || args.Distribute == nil {
		return nil, errors.New("missing required argument 'Distribute'")
	}
	if args == nil || args.ImageTemplateName == nil {
		return nil, errors.New("missing required argument 'ImageTemplateName'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Source == nil {
		return nil, errors.New("missing required argument 'Source'")
	}
	if args == nil {
		args = &VirtualMachineImageTemplateArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:virtualmachineimages/latest:VirtualMachineImageTemplate"),
		},
		{
			Type: pulumi.String("azurerm:virtualmachineimages/v20190201preview:VirtualMachineImageTemplate"),
		},
		{
			Type: pulumi.String("azurerm:virtualmachineimages/v20190501preview:VirtualMachineImageTemplate"),
		},
		{
			Type: pulumi.String("azurerm:virtualmachineimages/v20200214:VirtualMachineImageTemplate"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualMachineImageTemplate
	err := ctx.RegisterResource("azurerm:virtualmachineimages/v20180201preview:VirtualMachineImageTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualMachineImageTemplate gets an existing VirtualMachineImageTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualMachineImageTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMachineImageTemplateState, opts ...pulumi.ResourceOption) (*VirtualMachineImageTemplate, error) {
	var resource VirtualMachineImageTemplate
	err := ctx.ReadResource("azurerm:virtualmachineimages/v20180201preview:VirtualMachineImageTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualMachineImageTemplate resources.
type virtualMachineImageTemplateState struct {
	// Specifies the properties used to describe the customization steps of the image, like Image source etc
	Customize []ImageTemplateShellCustomizerResponse `pulumi:"customize"`
	// The distribution targets where the image output needs to go to.
	Distribute []interface{} `pulumi:"distribute"`
	// State of 'run' that is currently executing or was last executed.
	LastRunStatus *ImageTemplateLastRunStatusResponse `pulumi:"lastRunStatus"`
	// Resource location
	Location *string `pulumi:"location"`
	// Resource name
	Name *string `pulumi:"name"`
	// Provisioning error, if any
	ProvisioningError *ProvisioningErrorResponse `pulumi:"provisioningError"`
	// Provisioning state of the resource
	ProvisioningState *string `pulumi:"provisioningState"`
	// Specifies the properties used to describe the source image.
	Source interface{} `pulumi:"source"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

type VirtualMachineImageTemplateState struct {
	// Specifies the properties used to describe the customization steps of the image, like Image source etc
	Customize ImageTemplateShellCustomizerResponseArrayInput
	// The distribution targets where the image output needs to go to.
	Distribute pulumi.ArrayInput
	// State of 'run' that is currently executing or was last executed.
	LastRunStatus ImageTemplateLastRunStatusResponsePtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// Provisioning error, if any
	ProvisioningError ProvisioningErrorResponsePtrInput
	// Provisioning state of the resource
	ProvisioningState pulumi.StringPtrInput
	// Specifies the properties used to describe the source image.
	Source pulumi.Input
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (VirtualMachineImageTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineImageTemplateState)(nil)).Elem()
}

type virtualMachineImageTemplateArgs struct {
	// Specifies the properties used to describe the customization steps of the image, like Image source etc
	Customize []ImageTemplateShellCustomizer `pulumi:"customize"`
	// The distribution targets where the image output needs to go to.
	Distribute []interface{} `pulumi:"distribute"`
	// The name of the image Template
	ImageTemplateName string `pulumi:"imageTemplateName"`
	// Resource location
	Location string `pulumi:"location"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the properties used to describe the source image.
	Source interface{} `pulumi:"source"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a VirtualMachineImageTemplate resource.
type VirtualMachineImageTemplateArgs struct {
	// Specifies the properties used to describe the customization steps of the image, like Image source etc
	Customize ImageTemplateShellCustomizerArrayInput
	// The distribution targets where the image output needs to go to.
	Distribute pulumi.ArrayInput
	// The name of the image Template
	ImageTemplateName pulumi.StringInput
	// Resource location
	Location pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Specifies the properties used to describe the source image.
	Source pulumi.Input
	// Resource tags
	Tags pulumi.StringMapInput
}

func (VirtualMachineImageTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineImageTemplateArgs)(nil)).Elem()
}
