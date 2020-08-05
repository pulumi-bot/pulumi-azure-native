// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180601

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Information about packet capture session.
type PacketCapture struct {
	pulumi.CustomResourceState

	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Name of the packet capture session.
	Name pulumi.StringOutput `pulumi:"name"`
	// Describes the properties of a packet capture session.
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
	err := ctx.RegisterResource("azurerm:network/v20180601:PacketCapture", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azurerm:network/v20180601:PacketCapture", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PacketCapture resources.
type packetCaptureState struct {
	Etag *string `pulumi:"etag"`
	// Name of the packet capture session.
	Name *string `pulumi:"name"`
	// Describes the properties of a packet capture session.
	Properties *PacketCaptureResultPropertiesResponse `pulumi:"properties"`
}

type PacketCaptureState struct {
	Etag pulumi.StringPtrInput
	// Name of the packet capture session.
	Name pulumi.StringPtrInput
	// Describes the properties of a packet capture session.
	Properties PacketCaptureResultPropertiesResponsePtrInput
}

func (PacketCaptureState) ElementType() reflect.Type {
	return reflect.TypeOf((*packetCaptureState)(nil)).Elem()
}

type packetCaptureArgs struct {
	// Number of bytes captured per packet, the remaining bytes are truncated.
	BytesToCapturePerPacket *int                  `pulumi:"bytesToCapturePerPacket"`
	Filters                 []PacketCaptureFilter `pulumi:"filters"`
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
	Filters                 PacketCaptureFilterArrayInput
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
