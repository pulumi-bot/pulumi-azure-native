// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The volume.
// Latest API Version: 2017-06-01.
type Volume struct {
	pulumi.CustomResourceState

	// The IDs of the access control records, associated with the volume.
	AccessControlRecordIds pulumi.StringArrayOutput `pulumi:"accessControlRecordIds"`
	// The IDs of the backup policies, in which this volume is part of.
	BackupPolicyIds pulumi.StringArrayOutput `pulumi:"backupPolicyIds"`
	// The backup status of the volume.
	BackupStatus pulumi.StringOutput `pulumi:"backupStatus"`
	// The Kind of the object. Currently only Series8000 is supported
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// The monitoring status of the volume.
	MonitoringStatus pulumi.StringOutput `pulumi:"monitoringStatus"`
	// The name of the object.
	Name pulumi.StringOutput `pulumi:"name"`
	// The operation status on the volume.
	OperationStatus pulumi.StringOutput `pulumi:"operationStatus"`
	// The size of the volume in bytes.
	SizeInBytes pulumi.Float64Output `pulumi:"sizeInBytes"`
	// The hierarchical type of the object.
	Type pulumi.StringOutput `pulumi:"type"`
	// The ID of the volume container, in which this volume is created.
	VolumeContainerId pulumi.StringOutput `pulumi:"volumeContainerId"`
	// The volume status.
	VolumeStatus pulumi.StringOutput `pulumi:"volumeStatus"`
	// The type of the volume.
	VolumeType pulumi.StringOutput `pulumi:"volumeType"`
}

// NewVolume registers a new resource with the given unique name, arguments, and options.
func NewVolume(ctx *pulumi.Context,
	name string, args *VolumeArgs, opts ...pulumi.ResourceOption) (*Volume, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccessControlRecordIds == nil {
		return nil, errors.New("invalid value for required argument 'AccessControlRecordIds'")
	}
	if args.DeviceName == nil {
		return nil, errors.New("invalid value for required argument 'DeviceName'")
	}
	if args.ManagerName == nil {
		return nil, errors.New("invalid value for required argument 'ManagerName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SizeInBytes == nil {
		return nil, errors.New("invalid value for required argument 'SizeInBytes'")
	}
	if args.VolumeContainerName == nil {
		return nil, errors.New("invalid value for required argument 'VolumeContainerName'")
	}
	if args.VolumeName == nil {
		return nil, errors.New("invalid value for required argument 'VolumeName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:storsimple/v20170601:Volume"),
		},
	})
	opts = append(opts, aliases)
	var resource Volume
	err := ctx.RegisterResource("azure-nextgen:storsimple/latest:Volume", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVolume gets an existing Volume resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVolume(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VolumeState, opts ...pulumi.ResourceOption) (*Volume, error) {
	var resource Volume
	err := ctx.ReadResource("azure-nextgen:storsimple/latest:Volume", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Volume resources.
type volumeState struct {
	// The IDs of the access control records, associated with the volume.
	AccessControlRecordIds []string `pulumi:"accessControlRecordIds"`
	// The IDs of the backup policies, in which this volume is part of.
	BackupPolicyIds []string `pulumi:"backupPolicyIds"`
	// The backup status of the volume.
	BackupStatus *string `pulumi:"backupStatus"`
	// The Kind of the object. Currently only Series8000 is supported
	Kind *string `pulumi:"kind"`
	// The monitoring status of the volume.
	MonitoringStatus *string `pulumi:"monitoringStatus"`
	// The name of the object.
	Name *string `pulumi:"name"`
	// The operation status on the volume.
	OperationStatus *string `pulumi:"operationStatus"`
	// The size of the volume in bytes.
	SizeInBytes *float64 `pulumi:"sizeInBytes"`
	// The hierarchical type of the object.
	Type *string `pulumi:"type"`
	// The ID of the volume container, in which this volume is created.
	VolumeContainerId *string `pulumi:"volumeContainerId"`
	// The volume status.
	VolumeStatus *string `pulumi:"volumeStatus"`
	// The type of the volume.
	VolumeType *string `pulumi:"volumeType"`
}

type VolumeState struct {
	// The IDs of the access control records, associated with the volume.
	AccessControlRecordIds pulumi.StringArrayInput
	// The IDs of the backup policies, in which this volume is part of.
	BackupPolicyIds pulumi.StringArrayInput
	// The backup status of the volume.
	BackupStatus pulumi.StringPtrInput
	// The Kind of the object. Currently only Series8000 is supported
	Kind pulumi.StringPtrInput
	// The monitoring status of the volume.
	MonitoringStatus pulumi.StringPtrInput
	// The name of the object.
	Name pulumi.StringPtrInput
	// The operation status on the volume.
	OperationStatus pulumi.StringPtrInput
	// The size of the volume in bytes.
	SizeInBytes pulumi.Float64PtrInput
	// The hierarchical type of the object.
	Type pulumi.StringPtrInput
	// The ID of the volume container, in which this volume is created.
	VolumeContainerId pulumi.StringPtrInput
	// The volume status.
	VolumeStatus pulumi.StringPtrInput
	// The type of the volume.
	VolumeType pulumi.StringPtrInput
}

func (VolumeState) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeState)(nil)).Elem()
}

type volumeArgs struct {
	// The IDs of the access control records, associated with the volume.
	AccessControlRecordIds []string `pulumi:"accessControlRecordIds"`
	// The device name
	DeviceName string `pulumi:"deviceName"`
	// The Kind of the object. Currently only Series8000 is supported
	Kind *string `pulumi:"kind"`
	// The manager name
	ManagerName string `pulumi:"managerName"`
	// The monitoring status of the volume.
	MonitoringStatus string `pulumi:"monitoringStatus"`
	// The resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The size of the volume in bytes.
	SizeInBytes float64 `pulumi:"sizeInBytes"`
	// The volume container name.
	VolumeContainerName string `pulumi:"volumeContainerName"`
	// The volume name.
	VolumeName string `pulumi:"volumeName"`
	// The volume status.
	VolumeStatus string `pulumi:"volumeStatus"`
	// The type of the volume.
	VolumeType string `pulumi:"volumeType"`
}

// The set of arguments for constructing a Volume resource.
type VolumeArgs struct {
	// The IDs of the access control records, associated with the volume.
	AccessControlRecordIds pulumi.StringArrayInput
	// The device name
	DeviceName pulumi.StringInput
	// The Kind of the object. Currently only Series8000 is supported
	Kind Kind
	// The manager name
	ManagerName pulumi.StringInput
	// The monitoring status of the volume.
	MonitoringStatus MonitoringStatus
	// The resource group name
	ResourceGroupName pulumi.StringInput
	// The size of the volume in bytes.
	SizeInBytes pulumi.Float64Input
	// The volume container name.
	VolumeContainerName pulumi.StringInput
	// The volume name.
	VolumeName pulumi.StringInput
	// The volume status.
	VolumeStatus VolumeStatus
	// The type of the volume.
	VolumeType VolumeType
}

func (VolumeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeArgs)(nil)).Elem()
}

type VolumeInput interface {
	pulumi.Input

	ToVolumeOutput() VolumeOutput
	ToVolumeOutputWithContext(ctx context.Context) VolumeOutput
}

func (*Volume) ElementType() reflect.Type {
	return reflect.TypeOf((*Volume)(nil))
}

func (i *Volume) ToVolumeOutput() VolumeOutput {
	return i.ToVolumeOutputWithContext(context.Background())
}

func (i *Volume) ToVolumeOutputWithContext(ctx context.Context) VolumeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeOutput)
}

type VolumeOutput struct {
	*pulumi.OutputState
}

func (VolumeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Volume)(nil))
}

func (o VolumeOutput) ToVolumeOutput() VolumeOutput {
	return o
}

func (o VolumeOutput) ToVolumeOutputWithContext(ctx context.Context) VolumeOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(VolumeOutput{})
}
