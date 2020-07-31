// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170601

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The bandwidth setting.
type BandwidthSetting struct {
	pulumi.CustomResourceState

	// The Kind of the object. Currently only Series8000 is supported
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// The name of the object.
	Name pulumi.StringOutput `pulumi:"name"`
	// The properties of the bandwidth setting.
	Properties BandwidthRateSettingPropertiesResponseOutput `pulumi:"properties"`
	// The hierarchical type of the object.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewBandwidthSetting registers a new resource with the given unique name, arguments, and options.
func NewBandwidthSetting(ctx *pulumi.Context,
	name string, args *BandwidthSettingArgs, opts ...pulumi.ResourceOption) (*BandwidthSetting, error) {
	if args == nil || args.ManagerName == nil {
		return nil, errors.New("missing required argument 'ManagerName'")
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
		args = &BandwidthSettingArgs{}
	}
	var resource BandwidthSetting
	err := ctx.RegisterResource("azurerm:storsimple/v20170601:BandwidthSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBandwidthSetting gets an existing BandwidthSetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBandwidthSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BandwidthSettingState, opts ...pulumi.ResourceOption) (*BandwidthSetting, error) {
	var resource BandwidthSetting
	err := ctx.ReadResource("azurerm:storsimple/v20170601:BandwidthSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BandwidthSetting resources.
type bandwidthSettingState struct {
	// The Kind of the object. Currently only Series8000 is supported
	Kind *string `pulumi:"kind"`
	// The name of the object.
	Name *string `pulumi:"name"`
	// The properties of the bandwidth setting.
	Properties *BandwidthRateSettingPropertiesResponse `pulumi:"properties"`
	// The hierarchical type of the object.
	Type *string `pulumi:"type"`
}

type BandwidthSettingState struct {
	// The Kind of the object. Currently only Series8000 is supported
	Kind pulumi.StringPtrInput
	// The name of the object.
	Name pulumi.StringPtrInput
	// The properties of the bandwidth setting.
	Properties BandwidthRateSettingPropertiesResponsePtrInput
	// The hierarchical type of the object.
	Type pulumi.StringPtrInput
}

func (BandwidthSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*bandwidthSettingState)(nil)).Elem()
}

type bandwidthSettingArgs struct {
	// The Kind of the object. Currently only Series8000 is supported
	Kind *string `pulumi:"kind"`
	// The manager name
	ManagerName string `pulumi:"managerName"`
	// The bandwidth setting name.
	Name string `pulumi:"name"`
	// The properties of the bandwidth setting.
	Properties BandwidthRateSettingProperties `pulumi:"properties"`
	// The resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a BandwidthSetting resource.
type BandwidthSettingArgs struct {
	// The Kind of the object. Currently only Series8000 is supported
	Kind pulumi.StringPtrInput
	// The manager name
	ManagerName pulumi.StringInput
	// The bandwidth setting name.
	Name pulumi.StringInput
	// The properties of the bandwidth setting.
	Properties BandwidthRateSettingPropertiesInput
	// The resource group name
	ResourceGroupName pulumi.StringInput
}

func (BandwidthSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bandwidthSettingArgs)(nil)).Elem()
}