// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190801

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The bandwidth schedule details.
type BandwidthSchedule struct {
	pulumi.CustomResourceState

	// The object name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The properties of the bandwidth schedule.
	Properties BandwidthSchedulePropertiesResponseOutput `pulumi:"properties"`
	// The hierarchical type of the object.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewBandwidthSchedule registers a new resource with the given unique name, arguments, and options.
func NewBandwidthSchedule(ctx *pulumi.Context,
	name string, args *BandwidthScheduleArgs, opts ...pulumi.ResourceOption) (*BandwidthSchedule, error) {
	if args == nil || args.DeviceName == nil {
		return nil, errors.New("missing required argument 'DeviceName'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.Properties == nil {
		return nil, errors.New("missing required argument 'Properties'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &BandwidthScheduleArgs{}
	}
	var resource BandwidthSchedule
	err := ctx.RegisterResource("azurerm:databoxedge/v20190801:BandwidthSchedule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBandwidthSchedule gets an existing BandwidthSchedule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBandwidthSchedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BandwidthScheduleState, opts ...pulumi.ResourceOption) (*BandwidthSchedule, error) {
	var resource BandwidthSchedule
	err := ctx.ReadResource("azurerm:databoxedge/v20190801:BandwidthSchedule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BandwidthSchedule resources.
type bandwidthScheduleState struct {
	// The object name.
	Name *string `pulumi:"name"`
	// The properties of the bandwidth schedule.
	Properties *BandwidthSchedulePropertiesResponse `pulumi:"properties"`
	// The hierarchical type of the object.
	Type *string `pulumi:"type"`
}

type BandwidthScheduleState struct {
	// The object name.
	Name pulumi.StringPtrInput
	// The properties of the bandwidth schedule.
	Properties BandwidthSchedulePropertiesResponsePtrInput
	// The hierarchical type of the object.
	Type pulumi.StringPtrInput
}

func (BandwidthScheduleState) ElementType() reflect.Type {
	return reflect.TypeOf((*bandwidthScheduleState)(nil)).Elem()
}

type bandwidthScheduleArgs struct {
	// The device name.
	DeviceName string `pulumi:"deviceName"`
	// The bandwidth schedule name which needs to be added/updated.
	Name string `pulumi:"name"`
	// The properties of the bandwidth schedule.
	Properties BandwidthScheduleProperties `pulumi:"properties"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a BandwidthSchedule resource.
type BandwidthScheduleArgs struct {
	// The device name.
	DeviceName pulumi.StringInput
	// The bandwidth schedule name which needs to be added/updated.
	Name pulumi.StringInput
	// The properties of the bandwidth schedule.
	Properties BandwidthSchedulePropertiesInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
}

func (BandwidthScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bandwidthScheduleArgs)(nil)).Elem()
}