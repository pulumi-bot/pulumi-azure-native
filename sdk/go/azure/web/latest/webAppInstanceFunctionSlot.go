// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Function information.
type WebAppInstanceFunctionSlot struct {
	pulumi.CustomResourceState

	// Config information.
	Config pulumi.AnyOutput `pulumi:"config"`
	// Config URI.
	ConfigHref pulumi.StringPtrOutput `pulumi:"configHref"`
	// File list.
	Files pulumi.StringMapOutput `pulumi:"files"`
	// Function App ID.
	FunctionAppId pulumi.StringPtrOutput `pulumi:"functionAppId"`
	// Function URI.
	Href pulumi.StringPtrOutput `pulumi:"href"`
	// The invocation URL
	InvokeUrlTemplate pulumi.StringPtrOutput `pulumi:"invokeUrlTemplate"`
	// Gets or sets a value indicating whether the function is disabled
	IsDisabled pulumi.BoolPtrOutput `pulumi:"isDisabled"`
	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// The function language
	Language pulumi.StringPtrOutput `pulumi:"language"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Script URI.
	ScriptHref pulumi.StringPtrOutput `pulumi:"scriptHref"`
	// Script root path URI.
	ScriptRootPathHref pulumi.StringPtrOutput `pulumi:"scriptRootPathHref"`
	// Secrets file URI.
	SecretsFileHref pulumi.StringPtrOutput `pulumi:"secretsFileHref"`
	// Test data used when testing via the Azure Portal.
	TestData pulumi.StringPtrOutput `pulumi:"testData"`
	// Test data URI.
	TestDataHref pulumi.StringPtrOutput `pulumi:"testDataHref"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWebAppInstanceFunctionSlot registers a new resource with the given unique name, arguments, and options.
func NewWebAppInstanceFunctionSlot(ctx *pulumi.Context,
	name string, args *WebAppInstanceFunctionSlotArgs, opts ...pulumi.ResourceOption) (*WebAppInstanceFunctionSlot, error) {
	if args == nil || args.FunctionName == nil {
		return nil, errors.New("missing required argument 'FunctionName'")
	}
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
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:web/v20160801:WebAppInstanceFunctionSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20180201:WebAppInstanceFunctionSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20181101:WebAppInstanceFunctionSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20190801:WebAppInstanceFunctionSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200601:WebAppInstanceFunctionSlot"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppInstanceFunctionSlot
	err := ctx.RegisterResource("azure-nextgen:web/latest:WebAppInstanceFunctionSlot", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-nextgen:web/latest:WebAppInstanceFunctionSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAppInstanceFunctionSlot resources.
type webAppInstanceFunctionSlotState struct {
	// Config information.
	Config interface{} `pulumi:"config"`
	// Config URI.
	ConfigHref *string `pulumi:"configHref"`
	// File list.
	Files map[string]string `pulumi:"files"`
	// Function App ID.
	FunctionAppId *string `pulumi:"functionAppId"`
	// Function URI.
	Href *string `pulumi:"href"`
	// The invocation URL
	InvokeUrlTemplate *string `pulumi:"invokeUrlTemplate"`
	// Gets or sets a value indicating whether the function is disabled
	IsDisabled *bool `pulumi:"isDisabled"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// The function language
	Language *string `pulumi:"language"`
	// Resource Name.
	Name *string `pulumi:"name"`
	// Script URI.
	ScriptHref *string `pulumi:"scriptHref"`
	// Script root path URI.
	ScriptRootPathHref *string `pulumi:"scriptRootPathHref"`
	// Secrets file URI.
	SecretsFileHref *string `pulumi:"secretsFileHref"`
	// Test data used when testing via the Azure Portal.
	TestData *string `pulumi:"testData"`
	// Test data URI.
	TestDataHref *string `pulumi:"testDataHref"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type WebAppInstanceFunctionSlotState struct {
	// Config information.
	Config pulumi.Input
	// Config URI.
	ConfigHref pulumi.StringPtrInput
	// File list.
	Files pulumi.StringMapInput
	// Function App ID.
	FunctionAppId pulumi.StringPtrInput
	// Function URI.
	Href pulumi.StringPtrInput
	// The invocation URL
	InvokeUrlTemplate pulumi.StringPtrInput
	// Gets or sets a value indicating whether the function is disabled
	IsDisabled pulumi.BoolPtrInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// The function language
	Language pulumi.StringPtrInput
	// Resource Name.
	Name pulumi.StringPtrInput
	// Script URI.
	ScriptHref pulumi.StringPtrInput
	// Script root path URI.
	ScriptRootPathHref pulumi.StringPtrInput
	// Secrets file URI.
	SecretsFileHref pulumi.StringPtrInput
	// Test data used when testing via the Azure Portal.
	TestData pulumi.StringPtrInput
	// Test data URI.
	TestDataHref pulumi.StringPtrInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (WebAppInstanceFunctionSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppInstanceFunctionSlotState)(nil)).Elem()
}

type webAppInstanceFunctionSlotArgs struct {
	// Config information.
	Config interface{} `pulumi:"config"`
	// Config URI.
	ConfigHref *string `pulumi:"configHref"`
	// File list.
	Files map[string]string `pulumi:"files"`
	// Function App ID.
	FunctionAppId *string `pulumi:"functionAppId"`
	// Function name.
	FunctionName string `pulumi:"functionName"`
	// Function URI.
	Href *string `pulumi:"href"`
	// The invocation URL
	InvokeUrlTemplate *string `pulumi:"invokeUrlTemplate"`
	// Gets or sets a value indicating whether the function is disabled
	IsDisabled *bool `pulumi:"isDisabled"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// The function language
	Language *string `pulumi:"language"`
	// Site name.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Script URI.
	ScriptHref *string `pulumi:"scriptHref"`
	// Script root path URI.
	ScriptRootPathHref *string `pulumi:"scriptRootPathHref"`
	// Secrets file URI.
	SecretsFileHref *string `pulumi:"secretsFileHref"`
	// Name of the deployment slot.
	Slot string `pulumi:"slot"`
	// Test data used when testing via the Azure Portal.
	TestData *string `pulumi:"testData"`
	// Test data URI.
	TestDataHref *string `pulumi:"testDataHref"`
}

// The set of arguments for constructing a WebAppInstanceFunctionSlot resource.
type WebAppInstanceFunctionSlotArgs struct {
	// Config information.
	Config pulumi.Input
	// Config URI.
	ConfigHref pulumi.StringPtrInput
	// File list.
	Files pulumi.StringMapInput
	// Function App ID.
	FunctionAppId pulumi.StringPtrInput
	// Function name.
	FunctionName pulumi.StringInput
	// Function URI.
	Href pulumi.StringPtrInput
	// The invocation URL
	InvokeUrlTemplate pulumi.StringPtrInput
	// Gets or sets a value indicating whether the function is disabled
	IsDisabled pulumi.BoolPtrInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// The function language
	Language pulumi.StringPtrInput
	// Site name.
	Name pulumi.StringInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
	// Script URI.
	ScriptHref pulumi.StringPtrInput
	// Script root path URI.
	ScriptRootPathHref pulumi.StringPtrInput
	// Secrets file URI.
	SecretsFileHref pulumi.StringPtrInput
	// Name of the deployment slot.
	Slot pulumi.StringInput
	// Test data used when testing via the Azure Portal.
	TestData pulumi.StringPtrInput
	// Test data URI.
	TestDataHref pulumi.StringPtrInput
}

func (WebAppInstanceFunctionSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppInstanceFunctionSlotArgs)(nil)).Elem()
}

type WebAppInstanceFunctionSlotInput interface {
	pulumi.Input

	ToWebAppInstanceFunctionSlotOutput() WebAppInstanceFunctionSlotOutput
	ToWebAppInstanceFunctionSlotOutputWithContext(ctx context.Context) WebAppInstanceFunctionSlotOutput
}

func (WebAppInstanceFunctionSlot) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAppInstanceFunctionSlot)(nil)).Elem()
}

func (i WebAppInstanceFunctionSlot) ToWebAppInstanceFunctionSlotOutput() WebAppInstanceFunctionSlotOutput {
	return i.ToWebAppInstanceFunctionSlotOutputWithContext(context.Background())
}

func (i WebAppInstanceFunctionSlot) ToWebAppInstanceFunctionSlotOutputWithContext(ctx context.Context) WebAppInstanceFunctionSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppInstanceFunctionSlotOutput)
}

type WebAppInstanceFunctionSlotOutput struct {
	*pulumi.OutputState
}

func (WebAppInstanceFunctionSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAppInstanceFunctionSlotOutput)(nil)).Elem()
}

func (o WebAppInstanceFunctionSlotOutput) ToWebAppInstanceFunctionSlotOutput() WebAppInstanceFunctionSlotOutput {
	return o
}

func (o WebAppInstanceFunctionSlotOutput) ToWebAppInstanceFunctionSlotOutputWithContext(ctx context.Context) WebAppInstanceFunctionSlotOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WebAppInstanceFunctionSlotOutput{})
}
