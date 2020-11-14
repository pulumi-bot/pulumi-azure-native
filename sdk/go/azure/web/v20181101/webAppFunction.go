// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20181101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Web Job Information.
type WebAppFunction struct {
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
	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
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
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWebAppFunction registers a new resource with the given unique name, arguments, and options.
func NewWebAppFunction(ctx *pulumi.Context,
	name string, args *WebAppFunctionArgs, opts ...pulumi.ResourceOption) (*WebAppFunction, error) {
	if args == nil || args.FunctionName == nil {
		return nil, errors.New("missing required argument 'FunctionName'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &WebAppFunctionArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:web/latest:WebAppFunction"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20160801:WebAppFunction"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20180201:WebAppFunction"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20190801:WebAppFunction"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200601:WebAppFunction"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppFunction
	err := ctx.RegisterResource("azure-nextgen:web/v20181101:WebAppFunction", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAppFunction gets an existing WebAppFunction resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAppFunction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppFunctionState, opts ...pulumi.ResourceOption) (*WebAppFunction, error) {
	var resource WebAppFunction
	err := ctx.ReadResource("azure-nextgen:web/v20181101:WebAppFunction", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAppFunction resources.
type webAppFunctionState struct {
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
	// Kind of resource.
	Kind *string `pulumi:"kind"`
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
	// Resource type.
	Type *string `pulumi:"type"`
}

type WebAppFunctionState struct {
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
	// Kind of resource.
	Kind pulumi.StringPtrInput
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
	// Resource type.
	Type pulumi.StringPtrInput
}

func (WebAppFunctionState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppFunctionState)(nil)).Elem()
}

type webAppFunctionArgs struct {
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
	// Kind of resource.
	Kind *string `pulumi:"kind"`
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
	// Test data used when testing via the Azure Portal.
	TestData *string `pulumi:"testData"`
}

// The set of arguments for constructing a WebAppFunction resource.
type WebAppFunctionArgs struct {
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
	// Kind of resource.
	Kind pulumi.StringPtrInput
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
	// Test data used when testing via the Azure Portal.
	TestData pulumi.StringPtrInput
}

func (WebAppFunctionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppFunctionArgs)(nil)).Elem()
}

type WebAppFunctionInput interface {
	pulumi.Input

	ToWebAppFunctionOutput() WebAppFunctionOutput
	ToWebAppFunctionOutputWithContext(ctx context.Context) WebAppFunctionOutput
}

func (WebAppFunction) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAppFunction)(nil)).Elem()
}

func (i WebAppFunction) ToWebAppFunctionOutput() WebAppFunctionOutput {
	return i.ToWebAppFunctionOutputWithContext(context.Background())
}

func (i WebAppFunction) ToWebAppFunctionOutputWithContext(ctx context.Context) WebAppFunctionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppFunctionOutput)
}

type WebAppFunctionOutput struct {
	*pulumi.OutputState
}

func (WebAppFunctionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAppFunctionOutput)(nil)).Elem()
}

func (o WebAppFunctionOutput) ToWebAppFunctionOutput() WebAppFunctionOutput {
	return o
}

func (o WebAppFunctionOutput) ToWebAppFunctionOutputWithContext(ctx context.Context) WebAppFunctionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WebAppFunctionOutput{})
}
