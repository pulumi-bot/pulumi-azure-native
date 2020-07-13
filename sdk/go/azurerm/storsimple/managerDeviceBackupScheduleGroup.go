// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storsimple

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The Backup Schedule Group
type ManagerDeviceBackupScheduleGroup struct {
	pulumi.CustomResourceState

	// The name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of BackupScheduleGroup
	Properties BackupScheduleGroupPropertiesResponseOutput `pulumi:"properties"`
	// The type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewManagerDeviceBackupScheduleGroup registers a new resource with the given unique name, arguments, and options.
func NewManagerDeviceBackupScheduleGroup(ctx *pulumi.Context,
	name string, args *ManagerDeviceBackupScheduleGroupArgs, opts ...pulumi.ResourceOption) (*ManagerDeviceBackupScheduleGroup, error) {
	if args == nil || args.DeviceName == nil {
		return nil, errors.New("missing required argument 'DeviceName'")
	}
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
		args = &ManagerDeviceBackupScheduleGroupArgs{}
	}
	var resource ManagerDeviceBackupScheduleGroup
	err := ctx.RegisterResource("azurerm:storsimple:ManagerDeviceBackupScheduleGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagerDeviceBackupScheduleGroup gets an existing ManagerDeviceBackupScheduleGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagerDeviceBackupScheduleGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagerDeviceBackupScheduleGroupState, opts ...pulumi.ResourceOption) (*ManagerDeviceBackupScheduleGroup, error) {
	var resource ManagerDeviceBackupScheduleGroup
	err := ctx.ReadResource("azurerm:storsimple:ManagerDeviceBackupScheduleGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagerDeviceBackupScheduleGroup resources.
type managerDeviceBackupScheduleGroupState struct {
	// The name.
	Name *string `pulumi:"name"`
	// Properties of BackupScheduleGroup
	Properties *BackupScheduleGroupPropertiesResponse `pulumi:"properties"`
	// The type.
	Type *string `pulumi:"type"`
}

type ManagerDeviceBackupScheduleGroupState struct {
	// The name.
	Name pulumi.StringPtrInput
	// Properties of BackupScheduleGroup
	Properties BackupScheduleGroupPropertiesResponsePtrInput
	// The type.
	Type pulumi.StringPtrInput
}

func (ManagerDeviceBackupScheduleGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*managerDeviceBackupScheduleGroupState)(nil)).Elem()
}

type managerDeviceBackupScheduleGroupArgs struct {
	// The name of the device.
	DeviceName string `pulumi:"deviceName"`
	// The manager name
	ManagerName string `pulumi:"managerName"`
	// The name of the schedule group.
	Name string `pulumi:"name"`
	// Properties of BackupScheduleGroup
	Properties BackupScheduleGroupProperties `pulumi:"properties"`
	// The resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ManagerDeviceBackupScheduleGroup resource.
type ManagerDeviceBackupScheduleGroupArgs struct {
	// The name of the device.
	DeviceName pulumi.StringInput
	// The manager name
	ManagerName pulumi.StringInput
	// The name of the schedule group.
	Name pulumi.StringInput
	// Properties of BackupScheduleGroup
	Properties BackupScheduleGroupPropertiesInput
	// The resource group name
	ResourceGroupName pulumi.StringInput
}

func (ManagerDeviceBackupScheduleGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managerDeviceBackupScheduleGroupArgs)(nil)).Elem()
}