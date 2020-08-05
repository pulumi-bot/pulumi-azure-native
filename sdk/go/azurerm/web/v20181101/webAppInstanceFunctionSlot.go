// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20181101

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Web Job Information.
type WebAppInstanceFunctionSlot struct {
	pulumi.CustomResourceState

	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// FunctionEnvelope resource specific properties
	Properties FunctionEnvelopeResponsePropertiesOutput `pulumi:"properties"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWebAppInstanceFunctionSlot registers a new resource with the given unique name, arguments, and options.
func NewWebAppInstanceFunctionSlot(ctx *pulumi.Context,
	name string, args *WebAppInstanceFunctionSlotArgs, opts ...pulumi.ResourceOption) (*WebAppInstanceFunctionSlot, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Slot == nil {
		return nil, errors.New("missing required argument 'Slot'")
	}
	if args == nil {
		args = &WebAppInstanceFunctionSlotArgs{}
	}
	var resource WebAppInstanceFunctionSlot
	err := ctx.RegisterResource("azurerm:web/v20181101:WebAppInstanceFunctionSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAppInstanceFunctionSlot gets an existing WebAppInstanceFunctionSlot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAppInstanceFunctionSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppInstanceFunctionSlotState, opts ...pulumi.ResourceOption) (*WebAppInstanceFunctionSlot, error) {
	var resource WebAppInstanceFunctionSlot
	err := ctx.ReadResource("azurerm:web/v20181101:WebAppInstanceFunctionSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAppInstanceFunctionSlot resources.
type webAppInstanceFunctionSlotState struct {
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name *string `pulumi:"name"`
	// FunctionEnvelope resource specific properties
	Properties *FunctionEnvelopeResponseProperties `pulumi:"properties"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type WebAppInstanceFunctionSlotState struct {
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Resource Name.
	Name pulumi.StringPtrInput
	// FunctionEnvelope resource specific properties
	Properties FunctionEnvelopeResponsePropertiesPtrInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (WebAppInstanceFunctionSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppInstanceFunctionSlotState)(nil)).Elem()
}

type webAppInstanceFunctionSlotArgs struct {
	// Config information.
	Config map[string]interface{} `pulumi:"config"`
	// Config URI.
	Config_href *string `pulumi:"config_href"`
	// File list.
	Files map[string]string `pulumi:"files"`
	// Function App ID.
	Function_app_id *string `pulumi:"function_app_id"`
	// Function URI.
	Href *string `pulumi:"href"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Function name.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Script URI.
	Script_href *string `pulumi:"script_href"`
	// Script root path URI.
	Script_root_path_href *string `pulumi:"script_root_path_href"`
	// Secrets file URI.
	Secrets_file_href *string `pulumi:"secrets_file_href"`
	// Name of the deployment slot. If a slot is not specified, the API deletes a deployment for the production slot.
	Slot string `pulumi:"slot"`
	// Test data used when testing via the Azure Portal.
	Test_data *string `pulumi:"test_data"`
}

// The set of arguments for constructing a WebAppInstanceFunctionSlot resource.
type WebAppInstanceFunctionSlotArgs struct {
	// Config information.
	Config pulumi.MapInput
	// Config URI.
	Config_href pulumi.StringPtrInput
	// File list.
	Files pulumi.StringMapInput
	// Function App ID.
	Function_app_id pulumi.StringPtrInput
	// Function URI.
	Href pulumi.StringPtrInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Function name.
	Name pulumi.StringInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
	// Script URI.
	Script_href pulumi.StringPtrInput
	// Script root path URI.
	Script_root_path_href pulumi.StringPtrInput
	// Secrets file URI.
	Secrets_file_href pulumi.StringPtrInput
	// Name of the deployment slot. If a slot is not specified, the API deletes a deployment for the production slot.
	Slot pulumi.StringInput
	// Test data used when testing via the Azure Portal.
	Test_data pulumi.StringPtrInput
}

func (WebAppInstanceFunctionSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppInstanceFunctionSlotArgs)(nil)).Elem()
}
