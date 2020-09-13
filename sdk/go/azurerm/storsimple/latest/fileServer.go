// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The file server.
//
// ## Example Usage
// ### FileServersCreateOrUpdate
//
// ```go
// package main
//
// import (
// 	storsimple "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/storsimple/latest"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := storsimple.NewFileServer(ctx, "fileServer", &storsimple.FileServerArgs{
// 			BackupScheduleGroupId: pulumi.String("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/hsdk-4xy4fi2ivg/backupScheduleGroups/BackupSchGroupForSDKTest"),
// 			Description:           pulumi.String("Demo FileServer for SDK Test"),
// 			DeviceName:            pulumi.String("HSDK-4XY4FI2IVG"),
// 			DomainName:            pulumi.String("fareast.corp.microsoft.com"),
// 			FileServerName:        pulumi.String("HSDK-4XY4FI2IVG"),
// 			ManagerName:           pulumi.String("hAzureSDKOperations"),
// 			ResourceGroupName:     pulumi.String("ResourceGroupForSDKTest"),
// 			StorageDomainId:       pulumi.String("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/storageDomains/sd-fs-HSDK-4XY4FI2IVG"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type FileServer struct {
	pulumi.CustomResourceState

	// The backup policy id.
	BackupScheduleGroupId pulumi.StringOutput `pulumi:"backupScheduleGroupId"`
	// The description of the file server
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Domain of the file server
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// The name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The storage domain id.
	StorageDomainId pulumi.StringOutput `pulumi:"storageDomainId"`
	// The type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewFileServer registers a new resource with the given unique name, arguments, and options.
func NewFileServer(ctx *pulumi.Context,
	name string, args *FileServerArgs, opts ...pulumi.ResourceOption) (*FileServer, error) {
	if args == nil || args.BackupScheduleGroupId == nil {
		return nil, errors.New("missing required argument 'BackupScheduleGroupId'")
	}
	if args == nil || args.DeviceName == nil {
		return nil, errors.New("missing required argument 'DeviceName'")
	}
	if args == nil || args.DomainName == nil {
		return nil, errors.New("missing required argument 'DomainName'")
	}
	if args == nil || args.FileServerName == nil {
		return nil, errors.New("missing required argument 'FileServerName'")
	}
	if args == nil || args.ManagerName == nil {
		return nil, errors.New("missing required argument 'ManagerName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.StorageDomainId == nil {
		return nil, errors.New("missing required argument 'StorageDomainId'")
	}
	if args == nil {
		args = &FileServerArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:storsimple/v20161001:FileServer"),
		},
	})
	opts = append(opts, aliases)
	var resource FileServer
	err := ctx.RegisterResource("azurerm:storsimple/latest:FileServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFileServer gets an existing FileServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFileServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FileServerState, opts ...pulumi.ResourceOption) (*FileServer, error) {
	var resource FileServer
	err := ctx.ReadResource("azurerm:storsimple/latest:FileServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FileServer resources.
type fileServerState struct {
	// The backup policy id.
	BackupScheduleGroupId *string `pulumi:"backupScheduleGroupId"`
	// The description of the file server
	Description *string `pulumi:"description"`
	// Domain of the file server
	DomainName *string `pulumi:"domainName"`
	// The name.
	Name *string `pulumi:"name"`
	// The storage domain id.
	StorageDomainId *string `pulumi:"storageDomainId"`
	// The type.
	Type *string `pulumi:"type"`
}

type FileServerState struct {
	// The backup policy id.
	BackupScheduleGroupId pulumi.StringPtrInput
	// The description of the file server
	Description pulumi.StringPtrInput
	// Domain of the file server
	DomainName pulumi.StringPtrInput
	// The name.
	Name pulumi.StringPtrInput
	// The storage domain id.
	StorageDomainId pulumi.StringPtrInput
	// The type.
	Type pulumi.StringPtrInput
}

func (FileServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*fileServerState)(nil)).Elem()
}

type fileServerArgs struct {
	// The backup policy id.
	BackupScheduleGroupId string `pulumi:"backupScheduleGroupId"`
	// The description of the file server
	Description *string `pulumi:"description"`
	// The device name.
	DeviceName string `pulumi:"deviceName"`
	// Domain of the file server
	DomainName string `pulumi:"domainName"`
	// The file server name.
	FileServerName string `pulumi:"fileServerName"`
	// The manager name
	ManagerName string `pulumi:"managerName"`
	// The resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The storage domain id.
	StorageDomainId string `pulumi:"storageDomainId"`
}

// The set of arguments for constructing a FileServer resource.
type FileServerArgs struct {
	// The backup policy id.
	BackupScheduleGroupId pulumi.StringInput
	// The description of the file server
	Description pulumi.StringPtrInput
	// The device name.
	DeviceName pulumi.StringInput
	// Domain of the file server
	DomainName pulumi.StringInput
	// The file server name.
	FileServerName pulumi.StringInput
	// The manager name
	ManagerName pulumi.StringInput
	// The resource group name
	ResourceGroupName pulumi.StringInput
	// The storage domain id.
	StorageDomainId pulumi.StringInput
}

func (FileServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fileServerArgs)(nil)).Elem()
}
