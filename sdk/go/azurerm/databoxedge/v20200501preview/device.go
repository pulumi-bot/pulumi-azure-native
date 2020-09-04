// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200501preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The Data Box Edge/Gateway device.
type Device struct {
	pulumi.CustomResourceState

	// Type of compute roles configured.
	ConfiguredRoleTypes pulumi.StringArrayOutput `pulumi:"configuredRoleTypes"`
	// The Data Box Edge/Gateway device culture.
	Culture pulumi.StringOutput `pulumi:"culture"`
	// The status of the Data Box Edge/Gateway device.
	DataBoxEdgeDeviceStatus pulumi.StringPtrOutput `pulumi:"dataBoxEdgeDeviceStatus"`
	// The Description of the Data Box Edge/Gateway device.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The device software version number of the device (eg: 1.2.18105.6).
	DeviceHcsVersion pulumi.StringOutput `pulumi:"deviceHcsVersion"`
	// The Data Box Edge/Gateway device local capacity in MB.
	DeviceLocalCapacity pulumi.IntOutput `pulumi:"deviceLocalCapacity"`
	// The Data Box Edge/Gateway device model.
	DeviceModel pulumi.StringOutput `pulumi:"deviceModel"`
	// The Data Box Edge/Gateway device software version.
	DeviceSoftwareVersion pulumi.StringOutput `pulumi:"deviceSoftwareVersion"`
	// The type of the Data Box Edge/Gateway device.
	DeviceType pulumi.StringOutput `pulumi:"deviceType"`
	// The etag for the devices.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The Data Box Edge/Gateway device name.
	FriendlyName pulumi.StringPtrOutput `pulumi:"friendlyName"`
	// The location of the device. This is a supported and registered Azure geographical region (for example, West US, East US, or Southeast Asia). The geographical region of a device cannot be changed once it is created, but if an identical geographical region is specified on update, the request will succeed.
	Location pulumi.StringOutput `pulumi:"location"`
	// The description of the Data Box Edge/Gateway device model.
	ModelDescription pulumi.StringPtrOutput `pulumi:"modelDescription"`
	// The object name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The number of nodes in the cluster.
	NodeCount pulumi.IntOutput `pulumi:"nodeCount"`
	// The Serial Number of Data Box Edge/Gateway device.
	SerialNumber pulumi.StringOutput `pulumi:"serialNumber"`
	// The SKU type.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// The list of tags that describe the device. These tags can be used to view and group this device (across resource groups).
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The Data Box Edge/Gateway device timezone.
	TimeZone pulumi.StringOutput `pulumi:"timeZone"`
	// The hierarchical type of the object.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDevice registers a new resource with the given unique name, arguments, and options.
func NewDevice(ctx *pulumi.Context,
	name string, args *DeviceArgs, opts ...pulumi.ResourceOption) (*Device, error) {
	if args == nil || args.DeviceName == nil {
		return nil, errors.New("missing required argument 'DeviceName'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &DeviceArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:databoxedge/latest:Device"),
		},
		{
			Type: pulumi.String("azurerm:databoxedge/v20190301:Device"),
		},
		{
			Type: pulumi.String("azurerm:databoxedge/v20190701:Device"),
		},
		{
			Type: pulumi.String("azurerm:databoxedge/v20190801:Device"),
		},
	})
	opts = append(opts, aliases)
	var resource Device
	err := ctx.RegisterResource("azurerm:databoxedge/v20200501preview:Device", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDevice gets an existing Device resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDevice(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeviceState, opts ...pulumi.ResourceOption) (*Device, error) {
	var resource Device
	err := ctx.ReadResource("azurerm:databoxedge/v20200501preview:Device", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Device resources.
type deviceState struct {
	// Type of compute roles configured.
	ConfiguredRoleTypes []string `pulumi:"configuredRoleTypes"`
	// The Data Box Edge/Gateway device culture.
	Culture *string `pulumi:"culture"`
	// The status of the Data Box Edge/Gateway device.
	DataBoxEdgeDeviceStatus *string `pulumi:"dataBoxEdgeDeviceStatus"`
	// The Description of the Data Box Edge/Gateway device.
	Description *string `pulumi:"description"`
	// The device software version number of the device (eg: 1.2.18105.6).
	DeviceHcsVersion *string `pulumi:"deviceHcsVersion"`
	// The Data Box Edge/Gateway device local capacity in MB.
	DeviceLocalCapacity *int `pulumi:"deviceLocalCapacity"`
	// The Data Box Edge/Gateway device model.
	DeviceModel *string `pulumi:"deviceModel"`
	// The Data Box Edge/Gateway device software version.
	DeviceSoftwareVersion *string `pulumi:"deviceSoftwareVersion"`
	// The type of the Data Box Edge/Gateway device.
	DeviceType *string `pulumi:"deviceType"`
	// The etag for the devices.
	Etag *string `pulumi:"etag"`
	// The Data Box Edge/Gateway device name.
	FriendlyName *string `pulumi:"friendlyName"`
	// The location of the device. This is a supported and registered Azure geographical region (for example, West US, East US, or Southeast Asia). The geographical region of a device cannot be changed once it is created, but if an identical geographical region is specified on update, the request will succeed.
	Location *string `pulumi:"location"`
	// The description of the Data Box Edge/Gateway device model.
	ModelDescription *string `pulumi:"modelDescription"`
	// The object name.
	Name *string `pulumi:"name"`
	// The number of nodes in the cluster.
	NodeCount *int `pulumi:"nodeCount"`
	// The Serial Number of Data Box Edge/Gateway device.
	SerialNumber *string `pulumi:"serialNumber"`
	// The SKU type.
	Sku *SkuResponse `pulumi:"sku"`
	// The list of tags that describe the device. These tags can be used to view and group this device (across resource groups).
	Tags map[string]string `pulumi:"tags"`
	// The Data Box Edge/Gateway device timezone.
	TimeZone *string `pulumi:"timeZone"`
	// The hierarchical type of the object.
	Type *string `pulumi:"type"`
}

type DeviceState struct {
	// Type of compute roles configured.
	ConfiguredRoleTypes pulumi.StringArrayInput
	// The Data Box Edge/Gateway device culture.
	Culture pulumi.StringPtrInput
	// The status of the Data Box Edge/Gateway device.
	DataBoxEdgeDeviceStatus pulumi.StringPtrInput
	// The Description of the Data Box Edge/Gateway device.
	Description pulumi.StringPtrInput
	// The device software version number of the device (eg: 1.2.18105.6).
	DeviceHcsVersion pulumi.StringPtrInput
	// The Data Box Edge/Gateway device local capacity in MB.
	DeviceLocalCapacity pulumi.IntPtrInput
	// The Data Box Edge/Gateway device model.
	DeviceModel pulumi.StringPtrInput
	// The Data Box Edge/Gateway device software version.
	DeviceSoftwareVersion pulumi.StringPtrInput
	// The type of the Data Box Edge/Gateway device.
	DeviceType pulumi.StringPtrInput
	// The etag for the devices.
	Etag pulumi.StringPtrInput
	// The Data Box Edge/Gateway device name.
	FriendlyName pulumi.StringPtrInput
	// The location of the device. This is a supported and registered Azure geographical region (for example, West US, East US, or Southeast Asia). The geographical region of a device cannot be changed once it is created, but if an identical geographical region is specified on update, the request will succeed.
	Location pulumi.StringPtrInput
	// The description of the Data Box Edge/Gateway device model.
	ModelDescription pulumi.StringPtrInput
	// The object name.
	Name pulumi.StringPtrInput
	// The number of nodes in the cluster.
	NodeCount pulumi.IntPtrInput
	// The Serial Number of Data Box Edge/Gateway device.
	SerialNumber pulumi.StringPtrInput
	// The SKU type.
	Sku SkuResponsePtrInput
	// The list of tags that describe the device. These tags can be used to view and group this device (across resource groups).
	Tags pulumi.StringMapInput
	// The Data Box Edge/Gateway device timezone.
	TimeZone pulumi.StringPtrInput
	// The hierarchical type of the object.
	Type pulumi.StringPtrInput
}

func (DeviceState) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceState)(nil)).Elem()
}

type deviceArgs struct {
	// The status of the Data Box Edge/Gateway device.
	DataBoxEdgeDeviceStatus *string `pulumi:"dataBoxEdgeDeviceStatus"`
	// The Description of the Data Box Edge/Gateway device.
	Description *string `pulumi:"description"`
	// The device name.
	DeviceName string `pulumi:"deviceName"`
	// The etag for the devices.
	Etag *string `pulumi:"etag"`
	// The Data Box Edge/Gateway device name.
	FriendlyName *string `pulumi:"friendlyName"`
	// The location of the device. This is a supported and registered Azure geographical region (for example, West US, East US, or Southeast Asia). The geographical region of a device cannot be changed once it is created, but if an identical geographical region is specified on update, the request will succeed.
	Location string `pulumi:"location"`
	// The description of the Data Box Edge/Gateway device model.
	ModelDescription *string `pulumi:"modelDescription"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The SKU type.
	Sku *Sku `pulumi:"sku"`
	// The list of tags that describe the device. These tags can be used to view and group this device (across resource groups).
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Device resource.
type DeviceArgs struct {
	// The status of the Data Box Edge/Gateway device.
	DataBoxEdgeDeviceStatus pulumi.StringPtrInput
	// The Description of the Data Box Edge/Gateway device.
	Description pulumi.StringPtrInput
	// The device name.
	DeviceName pulumi.StringInput
	// The etag for the devices.
	Etag pulumi.StringPtrInput
	// The Data Box Edge/Gateway device name.
	FriendlyName pulumi.StringPtrInput
	// The location of the device. This is a supported and registered Azure geographical region (for example, West US, East US, or Southeast Asia). The geographical region of a device cannot be changed once it is created, but if an identical geographical region is specified on update, the request will succeed.
	Location pulumi.StringInput
	// The description of the Data Box Edge/Gateway device model.
	ModelDescription pulumi.StringPtrInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The SKU type.
	Sku SkuPtrInput
	// The list of tags that describe the device. These tags can be used to view and group this device (across resource groups).
	Tags pulumi.StringMapInput
}

func (DeviceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceArgs)(nil)).Elem()
}
