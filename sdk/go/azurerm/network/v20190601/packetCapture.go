// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190601

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Information about packet capture session.
type PacketCapture struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Name of the packet capture session.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the packet capture result.
	Properties PacketCaptureResultPropertiesResponseOutput `pulumi:"properties"`
}

// NewPacketCapture registers a new resource with the given unique name, arguments, and options.
func NewPacketCapture(ctx *pulumi.Context,
	name string, args *PacketCaptureArgs, opts ...pulumi.ResourceOption) (*PacketCapture, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.NetworkWatcherName == nil {
		return nil, errors.New("missing required argument 'NetworkWatcherName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.StorageLocation == nil {
		return nil, errors.New("missing required argument 'StorageLocation'")
	}
	if args == nil || args.Target == nil {
		return nil, errors.New("missing required argument 'Target'")
	}
	if args == nil {
		args = &PacketCaptureArgs{}
	}
	var resource PacketCapture
	err := ctx.RegisterResource("azurerm:network/v20190601:PacketCapture", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPacketCapture gets an existing PacketCapture resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPacketCapture(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PacketCaptureState, opts ...pulumi.ResourceOption) (*PacketCapture, error) {
	var resource PacketCapture
	err := ctx.ReadResource("azurerm:network/v20190601:PacketCapture", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PacketCapture resources.
type packetCaptureState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Name of the packet capture session.
	Name *string `pulumi:"name"`
	// Properties of the packet capture result.
	Properties *PacketCaptureResultPropertiesResponse `pulumi:"properties"`
}

type PacketCaptureState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Name of the packet capture session.
	Name pulumi.StringPtrInput
	// Properties of the packet capture result.
	Properties PacketCaptureResultPropertiesResponsePtrInput
}

func (PacketCaptureState) ElementType() reflect.Type {
	return reflect.TypeOf((*packetCaptureState)(nil)).Elem()
}

type packetCaptureArgs struct {
	// Number of bytes captured per packet, the remaining bytes are truncated.
	BytesToCapturePerPacket *int `pulumi:"bytesToCapturePerPacket"`
	// A list of packet capture filters.
	Filters []PacketCaptureFilter `pulumi:"filters"`
	// The name of the packet capture session.
	Name string `pulumi:"name"`
	// The name of the network watcher.
	NetworkWatcherName string `pulumi:"networkWatcherName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Describes the storage location for a packet capture session.
	StorageLocation PacketCaptureStorageLocation `pulumi:"storageLocation"`
	// The ID of the targeted resource, only VM is currently supported.
	Target string `pulumi:"target"`
	// Maximum duration of the capture session in seconds.
	TimeLimitInSeconds *int `pulumi:"timeLimitInSeconds"`
	// Maximum size of the capture output.
	TotalBytesPerSession *int `pulumi:"totalBytesPerSession"`
}

// The set of arguments for constructing a PacketCapture resource.
type PacketCaptureArgs struct {
	// Number of bytes captured per packet, the remaining bytes are truncated.
	BytesToCapturePerPacket pulumi.IntPtrInput
	// A list of packet capture filters.
	Filters PacketCaptureFilterArrayInput
	// The name of the packet capture session.
	Name pulumi.StringInput
	// The name of the network watcher.
	NetworkWatcherName pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Describes the storage location for a packet capture session.
	StorageLocation PacketCaptureStorageLocationInput
	// The ID of the targeted resource, only VM is currently supported.
	Target pulumi.StringInput
	// Maximum duration of the capture session in seconds.
	TimeLimitInSeconds pulumi.IntPtrInput
	// Maximum size of the capture output.
	TotalBytesPerSession pulumi.IntPtrInput
}

func (PacketCaptureArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*packetCaptureArgs)(nil)).Elem()
}
