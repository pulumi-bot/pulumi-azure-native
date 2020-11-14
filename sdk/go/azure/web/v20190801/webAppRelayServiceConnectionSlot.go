// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Hybrid Connection for an App Service app.
type WebAppRelayServiceConnectionSlot struct {
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

// NewWebAppRelayServiceConnectionSlot registers a new resource with the given unique name, arguments, and options.
func NewWebAppRelayServiceConnectionSlot(ctx *pulumi.Context,
	name string, args *WebAppRelayServiceConnectionSlotArgs, opts ...pulumi.ResourceOption) (*WebAppRelayServiceConnectionSlot, error) {
	if args == nil || args.EntityName == nil {
		return nil, errors.New("missing required argument 'EntityName'")
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
		args = &WebAppRelayServiceConnectionSlotArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:web/latest:WebAppRelayServiceConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20150801:WebAppRelayServiceConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20160801:WebAppRelayServiceConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20180201:WebAppRelayServiceConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20181101:WebAppRelayServiceConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200601:WebAppRelayServiceConnectionSlot"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppRelayServiceConnectionSlot
	err := ctx.RegisterResource("azure-nextgen:web/v20190801:WebAppRelayServiceConnectionSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAppRelayServiceConnectionSlot gets an existing WebAppRelayServiceConnectionSlot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAppRelayServiceConnectionSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppRelayServiceConnectionSlotState, opts ...pulumi.ResourceOption) (*WebAppRelayServiceConnectionSlot, error) {
	var resource WebAppRelayServiceConnectionSlot
	err := ctx.ReadResource("azure-nextgen:web/v20190801:WebAppRelayServiceConnectionSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAppRelayServiceConnectionSlot resources.
type webAppRelayServiceConnectionSlotState struct {
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

type WebAppRelayServiceConnectionSlotState struct {
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

func (WebAppRelayServiceConnectionSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppRelayServiceConnectionSlotState)(nil)).Elem()
}

type webAppRelayServiceConnectionSlotArgs struct {
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
	// Name of the deployment slot. If a slot is not specified, the API will create or update a hybrid connection for the production slot.
	Slot string `pulumi:"slot"`
}

// The set of arguments for constructing a WebAppRelayServiceConnectionSlot resource.
type WebAppRelayServiceConnectionSlotArgs struct {
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
	// Name of the deployment slot. If a slot is not specified, the API will create or update a hybrid connection for the production slot.
	Slot pulumi.StringInput
}

func (WebAppRelayServiceConnectionSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppRelayServiceConnectionSlotArgs)(nil)).Elem()
}

type WebAppRelayServiceConnectionSlotInput interface {
	pulumi.Input

	ToWebAppRelayServiceConnectionSlotOutput() WebAppRelayServiceConnectionSlotOutput
	ToWebAppRelayServiceConnectionSlotOutputWithContext(ctx context.Context) WebAppRelayServiceConnectionSlotOutput
}

func (WebAppRelayServiceConnectionSlot) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAppRelayServiceConnectionSlot)(nil)).Elem()
}

func (i WebAppRelayServiceConnectionSlot) ToWebAppRelayServiceConnectionSlotOutput() WebAppRelayServiceConnectionSlotOutput {
	return i.ToWebAppRelayServiceConnectionSlotOutputWithContext(context.Background())
}

func (i WebAppRelayServiceConnectionSlot) ToWebAppRelayServiceConnectionSlotOutputWithContext(ctx context.Context) WebAppRelayServiceConnectionSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppRelayServiceConnectionSlotOutput)
}

type WebAppRelayServiceConnectionSlotOutput struct {
	*pulumi.OutputState
}

func (WebAppRelayServiceConnectionSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAppRelayServiceConnectionSlotOutput)(nil)).Elem()
}

func (o WebAppRelayServiceConnectionSlotOutput) ToWebAppRelayServiceConnectionSlotOutput() WebAppRelayServiceConnectionSlotOutput {
	return o
}

func (o WebAppRelayServiceConnectionSlotOutput) ToWebAppRelayServiceConnectionSlotOutputWithContext(ctx context.Context) WebAppRelayServiceConnectionSlotOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WebAppRelayServiceConnectionSlotOutput{})
}
