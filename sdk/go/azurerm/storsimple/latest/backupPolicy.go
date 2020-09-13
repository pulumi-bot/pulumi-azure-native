// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The backup policy.
//
// ## Example Usage
// ### BackupPoliciesCreateOrUpdate
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
// 		_, err := storsimple.NewBackupPolicy(ctx, "backupPolicy", &storsimple.BackupPolicyArgs{
// 			BackupPolicyName:  pulumi.String("BkUpPolicy01ForSDKTest"),
// 			DeviceName:        pulumi.String("Device05ForSDKTest"),
// 			Kind:              pulumi.String("Series8000"),
// 			ManagerName:       pulumi.String("ManagerForSDKTest1"),
// 			ResourceGroupName: pulumi.String("ResourceGroupForSDKTest"),
// 			VolumeIds: pulumi.StringArray{
// 				pulumi.String("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/volumeContainers/volumeContainerForSDKTest/volumes/Clonedvolume1"),
// 				pulumi.String("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/volumeContainers/volumeContainerForSDKTest/volumes/volume1"),
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
type BackupPolicy struct {
	pulumi.CustomResourceState

	// The backup policy creation type. Indicates whether this was created through SaaS or through StorSimple Snapshot Manager.
	BackupPolicyCreationType pulumi.StringOutput `pulumi:"backupPolicyCreationType"`
	// The Kind of the object. Currently only Series8000 is supported
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// The time of the last backup for the backup policy.
	LastBackupTime pulumi.StringOutput `pulumi:"lastBackupTime"`
	// The name of the object.
	Name pulumi.StringOutput `pulumi:"name"`
	// The time of the next backup for the backup policy.
	NextBackupTime pulumi.StringOutput `pulumi:"nextBackupTime"`
	// Indicates whether at least one of the schedules in the backup policy is active or not.
	ScheduledBackupStatus pulumi.StringOutput `pulumi:"scheduledBackupStatus"`
	// The count of schedules the backup policy contains.
	SchedulesCount pulumi.IntOutput `pulumi:"schedulesCount"`
	// If the backup policy was created by StorSimple Snapshot Manager, then this field indicates the hostname of the StorSimple Snapshot Manager.
	SsmHostName pulumi.StringOutput `pulumi:"ssmHostName"`
	// The hierarchical type of the object.
	Type pulumi.StringOutput `pulumi:"type"`
	// The path IDs of the volumes which are part of the backup policy.
	VolumeIds pulumi.StringArrayOutput `pulumi:"volumeIds"`
}

// NewBackupPolicy registers a new resource with the given unique name, arguments, and options.
func NewBackupPolicy(ctx *pulumi.Context,
	name string, args *BackupPolicyArgs, opts ...pulumi.ResourceOption) (*BackupPolicy, error) {
	if args == nil || args.BackupPolicyName == nil {
		return nil, errors.New("missing required argument 'BackupPolicyName'")
	}
	if args == nil || args.DeviceName == nil {
		return nil, errors.New("missing required argument 'DeviceName'")
	}
	if args == nil || args.ManagerName == nil {
		return nil, errors.New("missing required argument 'ManagerName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.VolumeIds == nil {
		return nil, errors.New("missing required argument 'VolumeIds'")
	}
	if args == nil {
		args = &BackupPolicyArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:storsimple/v20170601:BackupPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource BackupPolicy
	err := ctx.RegisterResource("azurerm:storsimple/latest:BackupPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackupPolicy gets an existing BackupPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackupPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackupPolicyState, opts ...pulumi.ResourceOption) (*BackupPolicy, error) {
	var resource BackupPolicy
	err := ctx.ReadResource("azurerm:storsimple/latest:BackupPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackupPolicy resources.
type backupPolicyState struct {
	// The backup policy creation type. Indicates whether this was created through SaaS or through StorSimple Snapshot Manager.
	BackupPolicyCreationType *string `pulumi:"backupPolicyCreationType"`
	// The Kind of the object. Currently only Series8000 is supported
	Kind *string `pulumi:"kind"`
	// The time of the last backup for the backup policy.
	LastBackupTime *string `pulumi:"lastBackupTime"`
	// The name of the object.
	Name *string `pulumi:"name"`
	// The time of the next backup for the backup policy.
	NextBackupTime *string `pulumi:"nextBackupTime"`
	// Indicates whether at least one of the schedules in the backup policy is active or not.
	ScheduledBackupStatus *string `pulumi:"scheduledBackupStatus"`
	// The count of schedules the backup policy contains.
	SchedulesCount *int `pulumi:"schedulesCount"`
	// If the backup policy was created by StorSimple Snapshot Manager, then this field indicates the hostname of the StorSimple Snapshot Manager.
	SsmHostName *string `pulumi:"ssmHostName"`
	// The hierarchical type of the object.
	Type *string `pulumi:"type"`
	// The path IDs of the volumes which are part of the backup policy.
	VolumeIds []string `pulumi:"volumeIds"`
}

type BackupPolicyState struct {
	// The backup policy creation type. Indicates whether this was created through SaaS or through StorSimple Snapshot Manager.
	BackupPolicyCreationType pulumi.StringPtrInput
	// The Kind of the object. Currently only Series8000 is supported
	Kind pulumi.StringPtrInput
	// The time of the last backup for the backup policy.
	LastBackupTime pulumi.StringPtrInput
	// The name of the object.
	Name pulumi.StringPtrInput
	// The time of the next backup for the backup policy.
	NextBackupTime pulumi.StringPtrInput
	// Indicates whether at least one of the schedules in the backup policy is active or not.
	ScheduledBackupStatus pulumi.StringPtrInput
	// The count of schedules the backup policy contains.
	SchedulesCount pulumi.IntPtrInput
	// If the backup policy was created by StorSimple Snapshot Manager, then this field indicates the hostname of the StorSimple Snapshot Manager.
	SsmHostName pulumi.StringPtrInput
	// The hierarchical type of the object.
	Type pulumi.StringPtrInput
	// The path IDs of the volumes which are part of the backup policy.
	VolumeIds pulumi.StringArrayInput
}

func (BackupPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*backupPolicyState)(nil)).Elem()
}

type backupPolicyArgs struct {
	// The name of the backup policy to be created/updated.
	BackupPolicyName string `pulumi:"backupPolicyName"`
	// The device name
	DeviceName string `pulumi:"deviceName"`
	// The Kind of the object. Currently only Series8000 is supported
	Kind *string `pulumi:"kind"`
	// The manager name
	ManagerName string `pulumi:"managerName"`
	// The resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The path IDs of the volumes which are part of the backup policy.
	VolumeIds []string `pulumi:"volumeIds"`
}

// The set of arguments for constructing a BackupPolicy resource.
type BackupPolicyArgs struct {
	// The name of the backup policy to be created/updated.
	BackupPolicyName pulumi.StringInput
	// The device name
	DeviceName pulumi.StringInput
	// The Kind of the object. Currently only Series8000 is supported
	Kind pulumi.StringPtrInput
	// The manager name
	ManagerName pulumi.StringInput
	// The resource group name
	ResourceGroupName pulumi.StringInput
	// The path IDs of the volumes which are part of the backup policy.
	VolumeIds pulumi.StringArrayInput
}

func (BackupPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backupPolicyArgs)(nil)).Elem()
}
