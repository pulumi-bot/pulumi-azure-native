// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160801

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Hybrid Connection for an App Service app.
type WebAppRelayServiceConnectionSlot struct {
	pulumi.CustomResourceState

	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// RelayServiceConnectionEntity resource specific properties
	Properties RelayServiceConnectionEntityResponsePropertiesOutput `pulumi:"properties"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWebAppRelayServiceConnectionSlot registers a new resource with the given unique name, arguments, and options.
func NewWebAppRelayServiceConnectionSlot(ctx *pulumi.Context,
	name string, args *WebAppRelayServiceConnectionSlotArgs, opts ...pulumi.ResourceOption) (*WebAppRelayServiceConnectionSlot, error) {
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
	var resource WebAppRelayServiceConnectionSlot
	err := ctx.RegisterResource("azurerm:web/v20160801:WebAppRelayServiceConnectionSlot", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azurerm:web/v20160801:WebAppRelayServiceConnectionSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAppRelayServiceConnectionSlot resources.
type webAppRelayServiceConnectionSlotState struct {
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name *string `pulumi:"name"`
	// RelayServiceConnectionEntity resource specific properties
	Properties *RelayServiceConnectionEntityResponseProperties `pulumi:"properties"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type WebAppRelayServiceConnectionSlotState struct {
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Resource Name.
	Name pulumi.StringPtrInput
	// RelayServiceConnectionEntity resource specific properties
	Properties RelayServiceConnectionEntityResponsePropertiesPtrInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (WebAppRelayServiceConnectionSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppRelayServiceConnectionSlotState)(nil)).Elem()
}

type webAppRelayServiceConnectionSlotArgs struct {
	BiztalkUri             *string `pulumi:"biztalkUri"`
	EntityConnectionString *string `pulumi:"entityConnectionString"`
	Hostname               *string `pulumi:"hostname"`
	// Kind of resource.
	Kind                     *string `pulumi:"kind"`
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
	Hostname               pulumi.StringPtrInput
	// Kind of resource.
	Kind                     pulumi.StringPtrInput
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
