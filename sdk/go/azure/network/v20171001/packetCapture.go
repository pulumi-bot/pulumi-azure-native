// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20171001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Information about packet capture session.
type PacketCapture struct {
	pulumi.CustomResourceState

	// Number of bytes captured per packet, the remaining bytes are truncated.
	BytesToCapturePerPacket pulumi.IntPtrOutput                    `pulumi:"bytesToCapturePerPacket"`
	Etag                    pulumi.StringPtrOutput                 `pulumi:"etag"`
	Filters                 PacketCaptureFilterResponseArrayOutput `pulumi:"filters"`
	// Name of the packet capture session.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the packet capture session.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// Describes the storage location for a packet capture session.
	StorageLocation PacketCaptureStorageLocationResponseOutput `pulumi:"storageLocation"`
	// The ID of the targeted resource, only VM is currently supported.
	Target pulumi.StringOutput `pulumi:"target"`
	// Maximum duration of the capture session in seconds.
	TimeLimitInSeconds pulumi.IntPtrOutput `pulumi:"timeLimitInSeconds"`
	// Maximum size of the capture output.
	TotalBytesPerSession pulumi.IntPtrOutput `pulumi:"totalBytesPerSession"`
}

// NewPacketCapture registers a new resource with the given unique name, arguments, and options.
func NewPacketCapture(ctx *pulumi.Context,
	name string, args *PacketCaptureArgs, opts ...pulumi.ResourceOption) (*PacketCapture, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkWatcherName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkWatcherName'")
	}
	if args.PacketCaptureName == nil {
		return nil, errors.New("invalid value for required argument 'PacketCaptureName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StorageLocation == nil {
		return nil, errors.New("invalid value for required argument 'StorageLocation'")
	}
	if args.Target == nil {
		return nil, errors.New("invalid value for required argument 'Target'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:network/latest:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20160901:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20161201:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20170301:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20170601:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20170801:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20170901:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20171101:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180101:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180201:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180401:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180601:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180701:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180801:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181001:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181101:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181201:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190201:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190401:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190601:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190701:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190801:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190901:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191101:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191201:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200301:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200401:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200501:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200601:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200701:PacketCapture"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200801:PacketCapture"),
		},
	})
	opts = append(opts, aliases)
	var resource PacketCapture
	err := ctx.RegisterResource("azure-nextgen:network/v20171001:PacketCapture", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-nextgen:network/v20171001:PacketCapture", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PacketCapture resources.
type packetCaptureState struct {
	// Number of bytes captured per packet, the remaining bytes are truncated.
	BytesToCapturePerPacket *int                          `pulumi:"bytesToCapturePerPacket"`
	Etag                    *string                       `pulumi:"etag"`
	Filters                 []PacketCaptureFilterResponse `pulumi:"filters"`
	// Name of the packet capture session.
	Name *string `pulumi:"name"`
	// The provisioning state of the packet capture session.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Describes the storage location for a packet capture session.
	StorageLocation *PacketCaptureStorageLocationResponse `pulumi:"storageLocation"`
	// The ID of the targeted resource, only VM is currently supported.
	Target *string `pulumi:"target"`
	// Maximum duration of the capture session in seconds.
	TimeLimitInSeconds *int `pulumi:"timeLimitInSeconds"`
	// Maximum size of the capture output.
	TotalBytesPerSession *int `pulumi:"totalBytesPerSession"`
}

type PacketCaptureState struct {
	// Number of bytes captured per packet, the remaining bytes are truncated.
	BytesToCapturePerPacket pulumi.IntPtrInput
	Etag                    pulumi.StringPtrInput
	Filters                 PacketCaptureFilterResponseArrayInput
	// Name of the packet capture session.
	Name pulumi.StringPtrInput
	// The provisioning state of the packet capture session.
	ProvisioningState pulumi.StringPtrInput
	// Describes the storage location for a packet capture session.
	StorageLocation PacketCaptureStorageLocationResponsePtrInput
	// The ID of the targeted resource, only VM is currently supported.
	Target pulumi.StringPtrInput
	// Maximum duration of the capture session in seconds.
	TimeLimitInSeconds pulumi.IntPtrInput
	// Maximum size of the capture output.
	TotalBytesPerSession pulumi.IntPtrInput
}

func (PacketCaptureState) ElementType() reflect.Type {
	return reflect.TypeOf((*packetCaptureState)(nil)).Elem()
}

type packetCaptureArgs struct {
	// Number of bytes captured per packet, the remaining bytes are truncated.
	BytesToCapturePerPacket *int                  `pulumi:"bytesToCapturePerPacket"`
	Filters                 []PacketCaptureFilter `pulumi:"filters"`
	// The name of the network watcher.
	NetworkWatcherName string `pulumi:"networkWatcherName"`
	// The name of the packet capture session.
	PacketCaptureName string `pulumi:"packetCaptureName"`
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
	// The name of the network watcher.
	NetworkWatcherName pulumi.StringInput
	// The name of the packet capture session.
	PacketCaptureName pulumi.StringInput
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

type PacketCaptureInput interface {
	pulumi.Input

	ToPacketCaptureOutput() PacketCaptureOutput
	ToPacketCaptureOutputWithContext(ctx context.Context) PacketCaptureOutput
}

func (*PacketCapture) ElementType() reflect.Type {
	return reflect.TypeOf((*PacketCapture)(nil))
}

func (i *PacketCapture) ToPacketCaptureOutput() PacketCaptureOutput {
	return i.ToPacketCaptureOutputWithContext(context.Background())
}

func (i *PacketCapture) ToPacketCaptureOutputWithContext(ctx context.Context) PacketCaptureOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PacketCaptureOutput)
}

type PacketCaptureOutput struct {
	*pulumi.OutputState
}

func (PacketCaptureOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PacketCapture)(nil))
}

func (o PacketCaptureOutput) ToPacketCaptureOutput() PacketCaptureOutput {
	return o
}

func (o PacketCaptureOutput) ToPacketCaptureOutputWithContext(ctx context.Context) PacketCaptureOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PacketCaptureOutput{})
}
