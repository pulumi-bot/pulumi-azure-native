// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Hybrid Connection for an App Service app.
type WebAppRelayServiceConnection struct {
	pulumi.CustomResourceState

	BiztalkUri             pulumi.StringPtrOutput `pulumi:"biztalkUri"`
	EntityConnectionString pulumi.StringPtrOutput `pulumi:"entityConnectionString"`
	EntityName             pulumi.StringPtrOutput `pulumi:"entityName"`
	Hostname               pulumi.StringPtrOutput `pulumi:"hostname"`
	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Name.
	Name                     pulumi.StringOutput    `pulumi:"name"`
	Port                     pulumi.IntPtrOutput    `pulumi:"port"`
	ResourceConnectionString pulumi.StringPtrOutput `pulumi:"resourceConnectionString"`
	ResourceType             pulumi.StringPtrOutput `pulumi:"resourceType"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWebAppRelayServiceConnection registers a new resource with the given unique name, arguments, and options.
func NewWebAppRelayServiceConnection(ctx *pulumi.Context,
	name string, args *WebAppRelayServiceConnectionArgs, opts ...pulumi.ResourceOption) (*WebAppRelayServiceConnection, error) {
	if args == nil || args.EntityName == nil {
		return nil, errors.New("missing required argument 'EntityName'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &WebAppRelayServiceConnectionArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:web/latest:WebAppRelayServiceConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20150801:WebAppRelayServiceConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20160801:WebAppRelayServiceConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20181101:WebAppRelayServiceConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20190801:WebAppRelayServiceConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200601:WebAppRelayServiceConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppRelayServiceConnection
	err := ctx.RegisterResource("azure-nextgen:web/v20180201:WebAppRelayServiceConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAppRelayServiceConnection gets an existing WebAppRelayServiceConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAppRelayServiceConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppRelayServiceConnectionState, opts ...pulumi.ResourceOption) (*WebAppRelayServiceConnection, error) {
	var resource WebAppRelayServiceConnection
	err := ctx.ReadResource("azure-nextgen:web/v20180201:WebAppRelayServiceConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAppRelayServiceConnection resources.
type webAppRelayServiceConnectionState struct {
	BiztalkUri             *string `pulumi:"biztalkUri"`
	EntityConnectionString *string `pulumi:"entityConnectionString"`
	EntityName             *string `pulumi:"entityName"`
	Hostname               *string `pulumi:"hostname"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name                     *string `pulumi:"name"`
	Port                     *int    `pulumi:"port"`
	ResourceConnectionString *string `pulumi:"resourceConnectionString"`
	ResourceType             *string `pulumi:"resourceType"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type WebAppRelayServiceConnectionState struct {
	BiztalkUri             pulumi.StringPtrInput
	EntityConnectionString pulumi.StringPtrInput
	EntityName             pulumi.StringPtrInput
	Hostname               pulumi.StringPtrInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Resource Name.
	Name                     pulumi.StringPtrInput
	Port                     pulumi.IntPtrInput
	ResourceConnectionString pulumi.StringPtrInput
	ResourceType             pulumi.StringPtrInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (WebAppRelayServiceConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppRelayServiceConnectionState)(nil)).Elem()
}

type webAppRelayServiceConnectionArgs struct {
	BiztalkUri             *string `pulumi:"biztalkUri"`
	EntityConnectionString *string `pulumi:"entityConnectionString"`
	EntityName             string  `pulumi:"entityName"`
	Hostname               *string `pulumi:"hostname"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Name of the app.
	Name                     string  `pulumi:"name"`
	Port                     *int    `pulumi:"port"`
	ResourceConnectionString *string `pulumi:"resourceConnectionString"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ResourceType      *string `pulumi:"resourceType"`
}

// The set of arguments for constructing a WebAppRelayServiceConnection resource.
type WebAppRelayServiceConnectionArgs struct {
	BiztalkUri             pulumi.StringPtrInput
	EntityConnectionString pulumi.StringPtrInput
	EntityName             pulumi.StringInput
	Hostname               pulumi.StringPtrInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Name of the app.
	Name                     pulumi.StringInput
	Port                     pulumi.IntPtrInput
	ResourceConnectionString pulumi.StringPtrInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
	ResourceType      pulumi.StringPtrInput
}

func (WebAppRelayServiceConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppRelayServiceConnectionArgs)(nil)).Elem()
}

type WebAppRelayServiceConnectionInput interface {
	pulumi.Input

	ToWebAppRelayServiceConnectionOutput() WebAppRelayServiceConnectionOutput
	ToWebAppRelayServiceConnectionOutputWithContext(ctx context.Context) WebAppRelayServiceConnectionOutput
}

func (WebAppRelayServiceConnection) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAppRelayServiceConnection)(nil)).Elem()
}

func (i WebAppRelayServiceConnection) ToWebAppRelayServiceConnectionOutput() WebAppRelayServiceConnectionOutput {
	return i.ToWebAppRelayServiceConnectionOutputWithContext(context.Background())
}

func (i WebAppRelayServiceConnection) ToWebAppRelayServiceConnectionOutputWithContext(ctx context.Context) WebAppRelayServiceConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppRelayServiceConnectionOutput)
}

type WebAppRelayServiceConnectionOutput struct {
	*pulumi.OutputState
}

func (WebAppRelayServiceConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAppRelayServiceConnectionOutput)(nil)).Elem()
}

func (o WebAppRelayServiceConnectionOutput) ToWebAppRelayServiceConnectionOutput() WebAppRelayServiceConnectionOutput {
	return o
}

func (o WebAppRelayServiceConnectionOutput) ToWebAppRelayServiceConnectionOutputWithContext(ctx context.Context) WebAppRelayServiceConnectionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WebAppRelayServiceConnectionOutput{})
}
