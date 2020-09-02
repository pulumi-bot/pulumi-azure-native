// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190801

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Trigger details.
type Trigger struct {
	pulumi.CustomResourceState

	// Trigger Kind.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The object name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The hierarchical type of the object.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewTrigger registers a new resource with the given unique name, arguments, and options.
func NewTrigger(ctx *pulumi.Context,
	name string, args *TriggerArgs, opts ...pulumi.ResourceOption) (*Trigger, error) {
	if args == nil || args.DeviceName == nil {
		return nil, errors.New("missing required argument 'DeviceName'")
	}
	if args == nil || args.Kind == nil {
		return nil, errors.New("missing required argument 'Kind'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &TriggerArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:databoxedge/latest:Trigger"),
		},
		{
			Type: pulumi.String("azurerm:databoxedge/v20190301:Trigger"),
		},
		{
			Type: pulumi.String("azurerm:databoxedge/v20190701:Trigger"),
		},
		{
			Type: pulumi.String("azurerm:databoxedge/v20200501preview:Trigger"),
		},
	})
	opts = append(opts, aliases)
	var resource Trigger
	err := ctx.RegisterResource("azurerm:databoxedge/v20190801:Trigger", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrigger gets an existing Trigger resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrigger(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TriggerState, opts ...pulumi.ResourceOption) (*Trigger, error) {
	var resource Trigger
	err := ctx.ReadResource("azurerm:databoxedge/v20190801:Trigger", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Trigger resources.
type triggerState struct {
	// Trigger Kind.
	Kind *string `pulumi:"kind"`
	// The object name.
	Name *string `pulumi:"name"`
	// The hierarchical type of the object.
	Type *string `pulumi:"type"`
}

type TriggerState struct {
	// Trigger Kind.
	Kind pulumi.StringPtrInput
	// The object name.
	Name pulumi.StringPtrInput
	// The hierarchical type of the object.
	Type pulumi.StringPtrInput
}

func (TriggerState) ElementType() reflect.Type {
	return reflect.TypeOf((*triggerState)(nil)).Elem()
}

type triggerArgs struct {
	// Creates or updates a trigger
	DeviceName string `pulumi:"deviceName"`
	// Trigger Kind.
	Kind string `pulumi:"kind"`
	// The trigger name.
	Name string `pulumi:"name"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a Trigger resource.
type TriggerArgs struct {
	// Creates or updates a trigger
	DeviceName pulumi.StringInput
	// Trigger Kind.
	Kind pulumi.StringInput
	// The trigger name.
	Name pulumi.StringInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
}

func (TriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*triggerArgs)(nil)).Elem()
}
