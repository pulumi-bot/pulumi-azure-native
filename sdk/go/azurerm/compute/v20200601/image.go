// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200601

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The source user image virtual hard disk. The virtual hard disk will be copied before being attached to the virtual machine. If SourceImage is provided, the destination virtual hard drive must not exist.
//
// ## Example Usage
// ### Create a virtual machine image from a blob with DiskEncryptionSet resource.
//
// ```go
// package main
//
// import (
// 	compute "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/compute/v20200601"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := compute.NewImage(ctx, "image", &compute.ImageArgs{
// 			ImageName:         pulumi.String("myImage"),
// 			Location:          pulumi.String("West US"),
// 			ResourceGroupName: pulumi.String("myResourceGroup"),
// 			StorageProfile: &compute.ImageStorageProfileArgs{
// 				OsDisk: &compute.ImageOSDiskArgs{
// 					BlobUri: pulumi.String("https://mystorageaccount.blob.core.windows.net/osimages/osimage.vhd"),
// 					DiskEncryptionSet: &compute.DiskEncryptionSetParametersArgs{
// 						Id: pulumi.String("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{existing-diskEncryptionSet-name}"),
// 					},
// 					OsState: pulumi.String("Generalized"),
// 					OsType:  pulumi.String("Linux"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
// ### Create a virtual machine image from a blob.
//
// ```go
// package main
//
// import (
// 	compute "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/compute/v20200601"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := compute.NewImage(ctx, "image", &compute.ImageArgs{
// 			ImageName:         pulumi.String("myImage"),
// 			Location:          pulumi.String("West US"),
// 			ResourceGroupName: pulumi.String("myResourceGroup"),
// 			StorageProfile: &compute.ImageStorageProfileArgs{
// 				OsDisk: &compute.ImageOSDiskArgs{
// 					BlobUri: pulumi.String("https://mystorageaccount.blob.core.windows.net/osimages/osimage.vhd"),
// 					OsState: pulumi.String("Generalized"),
// 					OsType:  pulumi.String("Linux"),
// 				},
// 				ZoneResilient: pulumi.Bool(true),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
// ### Create a virtual machine image from a managed disk with DiskEncryptionSet resource.
//
// ```go
// package main
//
// import (
// 	compute "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/compute/v20200601"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := compute.NewImage(ctx, "image", &compute.ImageArgs{
// 			ImageName:         pulumi.String("myImage"),
// 			Location:          pulumi.String("West US"),
// 			ResourceGroupName: pulumi.String("myResourceGroup"),
// 			StorageProfile: &compute.ImageStorageProfileArgs{
// 				OsDisk: &compute.ImageOSDiskArgs{
// 					DiskEncryptionSet: &compute.DiskEncryptionSetParametersArgs{
// 						Id: pulumi.String("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{existing-diskEncryptionSet-name}"),
// 					},
// 					ManagedDisk: &compute.SubResourceArgs{
// 						Id: pulumi.String("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myManagedDisk"),
// 					},
// 					OsState: pulumi.String("Generalized"),
// 					OsType:  pulumi.String("Linux"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
// ### Create a virtual machine image from a managed disk.
//
// ```go
// package main
//
// import (
// 	compute "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/compute/v20200601"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := compute.NewImage(ctx, "image", &compute.ImageArgs{
// 			ImageName:         pulumi.String("myImage"),
// 			Location:          pulumi.String("West US"),
// 			ResourceGroupName: pulumi.String("myResourceGroup"),
// 			StorageProfile: &compute.ImageStorageProfileArgs{
// 				OsDisk: &compute.ImageOSDiskArgs{
// 					ManagedDisk: &compute.SubResourceArgs{
// 						Id: pulumi.String("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myManagedDisk"),
// 					},
// 					OsState: pulumi.String("Generalized"),
// 					OsType:  pulumi.String("Linux"),
// 				},
// 				ZoneResilient: pulumi.Bool(true),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
// ### Create a virtual machine image from a snapshot with DiskEncryptionSet resource.
//
// ```go
// package main
//
// import (
// 	compute "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/compute/v20200601"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := compute.NewImage(ctx, "image", &compute.ImageArgs{
// 			ImageName:         pulumi.String("myImage"),
// 			Location:          pulumi.String("West US"),
// 			ResourceGroupName: pulumi.String("myResourceGroup"),
// 			StorageProfile: &compute.ImageStorageProfileArgs{
// 				OsDisk: &compute.ImageOSDiskArgs{
// 					DiskEncryptionSet: &compute.DiskEncryptionSetParametersArgs{
// 						Id: pulumi.String("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{existing-diskEncryptionSet-name}"),
// 					},
// 					OsState: pulumi.String("Generalized"),
// 					OsType:  pulumi.String("Linux"),
// 					Snapshot: &compute.SubResourceArgs{
// 						Id: pulumi.String("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot"),
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
// ### Create a virtual machine image from a snapshot.
//
// ```go
// package main
//
// import (
// 	compute "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/compute/v20200601"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := compute.NewImage(ctx, "image", &compute.ImageArgs{
// 			ImageName:         pulumi.String("myImage"),
// 			Location:          pulumi.String("West US"),
// 			ResourceGroupName: pulumi.String("myResourceGroup"),
// 			StorageProfile: &compute.ImageStorageProfileArgs{
// 				OsDisk: &compute.ImageOSDiskArgs{
// 					OsState: pulumi.String("Generalized"),
// 					OsType:  pulumi.String("Linux"),
// 					Snapshot: &compute.SubResourceArgs{
// 						Id: pulumi.String("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot"),
// 					},
// 				},
// 				ZoneResilient: pulumi.Bool(false),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
// ### Create a virtual machine image from an existing virtual machine.
//
// ```go
// package main
//
// import (
// 	compute "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/compute/v20200601"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := compute.NewImage(ctx, "image", &compute.ImageArgs{
// 			ImageName:         pulumi.String("myImage"),
// 			Location:          pulumi.String("West US"),
// 			ResourceGroupName: pulumi.String("myResourceGroup"),
// 			SourceVirtualMachine: &compute.SubResourceArgs{
// 				Id: pulumi.String("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
// ### Create a virtual machine image that includes a data disk from a blob.
//
// ```go
// package main
//
// import (
// 	compute "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/compute/v20200601"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := compute.NewImage(ctx, "image", &compute.ImageArgs{
// 			ImageName:         pulumi.String("myImage"),
// 			Location:          pulumi.String("West US"),
// 			ResourceGroupName: pulumi.String("myResourceGroup"),
// 			StorageProfile: &compute.ImageStorageProfileArgs{
// 				DataDisks: compute.ImageDataDiskArray{
// 					&compute.ImageDataDiskArgs{
// 						BlobUri: pulumi.String("https://mystorageaccount.blob.core.windows.net/dataimages/dataimage.vhd"),
// 						Lun:     pulumi.Int(1),
// 					},
// 				},
// 				OsDisk: &compute.ImageOSDiskArgs{
// 					BlobUri: pulumi.String("https://mystorageaccount.blob.core.windows.net/osimages/osimage.vhd"),
// 					OsState: pulumi.String("Generalized"),
// 					OsType:  pulumi.String("Linux"),
// 				},
// 				ZoneResilient: pulumi.Bool(false),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
// ### Create a virtual machine image that includes a data disk from a managed disk.
//
// ```go
// package main
//
// import (
// 	compute "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/compute/v20200601"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := compute.NewImage(ctx, "image", &compute.ImageArgs{
// 			ImageName:         pulumi.String("myImage"),
// 			Location:          pulumi.String("West US"),
// 			ResourceGroupName: pulumi.String("myResourceGroup"),
// 			StorageProfile: &compute.ImageStorageProfileArgs{
// 				DataDisks: compute.ImageDataDiskArray{
// 					&compute.ImageDataDiskArgs{
// 						Lun: pulumi.Int(1),
// 						ManagedDisk: &compute.SubResourceArgs{
// 							Id: pulumi.String("subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myManagedDisk2"),
// 						},
// 					},
// 				},
// 				OsDisk: &compute.ImageOSDiskArgs{
// 					ManagedDisk: &compute.SubResourceArgs{
// 						Id: pulumi.String("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myManagedDisk"),
// 					},
// 					OsState: pulumi.String("Generalized"),
// 					OsType:  pulumi.String("Linux"),
// 				},
// 				ZoneResilient: pulumi.Bool(false),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
// ### Create a virtual machine image that includes a data disk from a snapshot.
//
// ```go
// package main
//
// import (
// 	compute "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/compute/v20200601"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := compute.NewImage(ctx, "image", &compute.ImageArgs{
// 			ImageName:         pulumi.String("myImage"),
// 			Location:          pulumi.String("West US"),
// 			ResourceGroupName: pulumi.String("myResourceGroup"),
// 			StorageProfile: &compute.ImageStorageProfileArgs{
// 				DataDisks: compute.ImageDataDiskArray{
// 					&compute.ImageDataDiskArgs{
// 						Lun: pulumi.Int(1),
// 						Snapshot: &compute.SubResourceArgs{
// 							Id: pulumi.String("subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot2"),
// 						},
// 					},
// 				},
// 				OsDisk: &compute.ImageOSDiskArgs{
// 					OsState: pulumi.String("Generalized"),
// 					OsType:  pulumi.String("Linux"),
// 					Snapshot: &compute.SubResourceArgs{
// 						Id: pulumi.String("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot"),
// 					},
// 				},
// 				ZoneResilient: pulumi.Bool(true),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type Image struct {
	pulumi.CustomResourceState

	// Gets the HyperVGenerationType of the VirtualMachine created from the image
	HyperVGeneration pulumi.StringPtrOutput `pulumi:"hyperVGeneration"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The source virtual machine from which Image is created.
	SourceVirtualMachine SubResourceResponsePtrOutput `pulumi:"sourceVirtualMachine"`
	// Specifies the storage settings for the virtual machine disks.
	StorageProfile ImageStorageProfileResponsePtrOutput `pulumi:"storageProfile"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewImage registers a new resource with the given unique name, arguments, and options.
func NewImage(ctx *pulumi.Context,
	name string, args *ImageArgs, opts ...pulumi.ResourceOption) (*Image, error) {
	if args == nil || args.ImageName == nil {
		return nil, errors.New("missing required argument 'ImageName'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ImageArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:compute/latest:Image"),
		},
		{
			Type: pulumi.String("azurerm:compute/v20160430preview:Image"),
		},
		{
			Type: pulumi.String("azurerm:compute/v20170330:Image"),
		},
		{
			Type: pulumi.String("azurerm:compute/v20171201:Image"),
		},
		{
			Type: pulumi.String("azurerm:compute/v20180401:Image"),
		},
		{
			Type: pulumi.String("azurerm:compute/v20180601:Image"),
		},
		{
			Type: pulumi.String("azurerm:compute/v20181001:Image"),
		},
		{
			Type: pulumi.String("azurerm:compute/v20190301:Image"),
		},
		{
			Type: pulumi.String("azurerm:compute/v20190701:Image"),
		},
		{
			Type: pulumi.String("azurerm:compute/v20191201:Image"),
		},
	})
	opts = append(opts, aliases)
	var resource Image
	err := ctx.RegisterResource("azurerm:compute/v20200601:Image", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetImage gets an existing Image resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImageState, opts ...pulumi.ResourceOption) (*Image, error) {
	var resource Image
	err := ctx.ReadResource("azurerm:compute/v20200601:Image", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Image resources.
type imageState struct {
	// Gets the HyperVGenerationType of the VirtualMachine created from the image
	HyperVGeneration *string `pulumi:"hyperVGeneration"`
	// Resource location
	Location *string `pulumi:"location"`
	// Resource name
	Name *string `pulumi:"name"`
	// The provisioning state.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The source virtual machine from which Image is created.
	SourceVirtualMachine *SubResourceResponse `pulumi:"sourceVirtualMachine"`
	// Specifies the storage settings for the virtual machine disks.
	StorageProfile *ImageStorageProfileResponse `pulumi:"storageProfile"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

type ImageState struct {
	// Gets the HyperVGenerationType of the VirtualMachine created from the image
	HyperVGeneration pulumi.StringPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// The provisioning state.
	ProvisioningState pulumi.StringPtrInput
	// The source virtual machine from which Image is created.
	SourceVirtualMachine SubResourceResponsePtrInput
	// Specifies the storage settings for the virtual machine disks.
	StorageProfile ImageStorageProfileResponsePtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (ImageState) ElementType() reflect.Type {
	return reflect.TypeOf((*imageState)(nil)).Elem()
}

type imageArgs struct {
	// Gets the HyperVGenerationType of the VirtualMachine created from the image
	HyperVGeneration *string `pulumi:"hyperVGeneration"`
	// The name of the image.
	ImageName string `pulumi:"imageName"`
	// Resource location
	Location string `pulumi:"location"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The source virtual machine from which Image is created.
	SourceVirtualMachine *SubResource `pulumi:"sourceVirtualMachine"`
	// Specifies the storage settings for the virtual machine disks.
	StorageProfile *ImageStorageProfile `pulumi:"storageProfile"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Image resource.
type ImageArgs struct {
	// Gets the HyperVGenerationType of the VirtualMachine created from the image
	HyperVGeneration pulumi.StringPtrInput
	// The name of the image.
	ImageName pulumi.StringInput
	// Resource location
	Location pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The source virtual machine from which Image is created.
	SourceVirtualMachine SubResourcePtrInput
	// Specifies the storage settings for the virtual machine disks.
	StorageProfile ImageStorageProfilePtrInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (ImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*imageArgs)(nil)).Elem()
}
