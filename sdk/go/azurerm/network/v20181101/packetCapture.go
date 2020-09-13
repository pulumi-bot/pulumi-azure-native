// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20181101

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Information about packet capture session.
//
// ## Example Usage
// ### Create packet capture
//
// ```go
// package main
//
// import (
// 	network "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/network/v20181101"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := network.NewPacketCapture(ctx, "packetCapture", &network.PacketCaptureArgs{
// 			BytesToCapturePerPacket: pulumi.Int(10000),
// 			Filters: network.PacketCaptureFilterArray{
// 				&network.PacketCaptureFilterArgs{
// 					LocalIPAddress: pulumi.String("10.0.0.4"),
// 					LocalPort:      pulumi.String("80"),
// 					Protocol:       pulumi.String("TCP"),
// 				},
// 			},
// 			NetworkWatcherName: pulumi.String("nw1"),
// 			PacketCaptureName:  pulumi.String("pc1"),
// 			ResourceGroupName:  pulumi.String("rg1"),
// 			StorageLocation: &network.PacketCaptureStorageLocationArgs{
// 				FilePath:    pulumi.String("D:\\capture\\pc1.cap"),
// 				StorageId:   pulumi.String("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Storage/storageAccounts/pcstore"),
// 				StoragePath: pulumi.String("https://mytestaccountname.blob.core.windows.net/capture/pc1.cap"),
// 			},
// 			Target:               pulumi.String("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Compute/virtualMachines/vm1"),
// 			TimeLimitInSeconds:   pulumi.Int(100),
// 			TotalBytesPerSession: pulumi.Int(100000),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
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
	if args == nil || args.NetworkWatcherName == nil {
		return nil, errors.New("missing required argument 'NetworkWatcherName'")
	}
	if args == nil || args.PacketCaptureName == nil {
		return nil, errors.New("missing required argument 'PacketCaptureName'")
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
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:network/latest:PacketCapture"),
		},
		{
			Type: pulumi.String("azurerm:network/v20160901:PacketCapture"),
		},
		{
			Type: pulumi.String("azurerm:network/v20161201:PacketCapture"),
		},
		{
			Type: pulumi.String("azurerm:network/v20170301:PacketCapture"),
		},
		{
			Type: pulumi.String("azurerm:network/v20170601:PacketCapture"),
		},
		{
			Type: pulumi.String("azurerm:network/v20170801:PacketCapture"),
		},
		{
			Type: pulumi.String("azurerm:network/v20170901:PacketCapture"),
		},
		{
			Type: pulumi.String("azurerm:network/v20171001:PacketCapture"),
		},
		{
			Type: pulumi.String("azurerm:network/v20171101:PacketCapture"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180101:PacketCapture"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180201:PacketCapture"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180401:PacketCapture"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180601:PacketCapture"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180701:PacketCapture"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180801:PacketCapture"),
		},
		{
			Type: pulumi.String("azurerm:network/v20181001:PacketCapture"),
		},
		{
			Type: pulumi.String("azurerm:network/v20181201:PacketCapture"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190201:PacketCapture"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190401:PacketCapture"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190601:PacketCapture"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190701:PacketCapture"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190801:PacketCapture"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190901:PacketCapture"),
		},
		{
			Type: pulumi.String("azurerm:network/v20191101:PacketCapture"),
		},
		{
			Type: pulumi.String("azurerm:network/v20191201:PacketCapture"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200301:PacketCapture"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200401:PacketCapture"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200501:PacketCapture"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200601:PacketCapture"),
		},
	})
	opts = append(opts, aliases)
	var resource PacketCapture
	err := ctx.RegisterResource("azurerm:network/v20181101:PacketCapture", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azurerm:network/v20181101:PacketCapture", name, id, state, &resource, opts...)
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
