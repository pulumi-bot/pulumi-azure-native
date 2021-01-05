// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The volume container.
type VolumeContainer struct {
	pulumi.CustomResourceState

	// The bandwidth-rate set on the volume container.
	BandWidthRateInMbps pulumi.IntPtrOutput `pulumi:"bandWidthRateInMbps"`
	// The ID of the bandwidth setting associated with the volume container.
	BandwidthSettingId pulumi.StringPtrOutput `pulumi:"bandwidthSettingId"`
	// The key used to encrypt data in the volume container. It is required when property 'EncryptionStatus' is "Enabled".
	EncryptionKey AsymmetricEncryptedSecretResponsePtrOutput `pulumi:"encryptionKey"`
	// The flag to denote whether encryption is enabled or not.
	EncryptionStatus pulumi.StringOutput `pulumi:"encryptionStatus"`
	// The Kind of the object. Currently only Series8000 is supported
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// The name of the object.
	Name pulumi.StringOutput `pulumi:"name"`
	// The owner ship status of the volume container. Only when the status is "NotOwned", the delete operation on the volume container is permitted.
	OwnerShipStatus pulumi.StringOutput `pulumi:"ownerShipStatus"`
	// The path ID of storage account associated with the volume container.
	StorageAccountCredentialId pulumi.StringOutput `pulumi:"storageAccountCredentialId"`
	// The total cloud storage for the volume container.
	TotalCloudStorageUsageInBytes pulumi.Float64Output `pulumi:"totalCloudStorageUsageInBytes"`
	// The hierarchical type of the object.
	Type pulumi.StringOutput `pulumi:"type"`
	// The number of volumes in the volume Container.
	VolumeCount pulumi.IntOutput `pulumi:"volumeCount"`
}

// NewVolumeContainer registers a new resource with the given unique name, arguments, and options.
func NewVolumeContainer(ctx *pulumi.Context,
	name string, args *VolumeContainerArgs, opts ...pulumi.ResourceOption) (*VolumeContainer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
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
	if args.StorageAccountCredentialId == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountCredentialId'")
	}
	if args.VolumeContainerName == nil {
		return nil, errors.New("invalid value for required argument 'VolumeContainerName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:storsimple/latest:VolumeContainer"),
		},
	})
	opts = append(opts, aliases)
	var resource VolumeContainer
	err := ctx.RegisterResource("azure-nextgen:storsimple/v20170601:VolumeContainer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVolumeContainer gets an existing VolumeContainer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVolumeContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VolumeContainerState, opts ...pulumi.ResourceOption) (*VolumeContainer, error) {
	var resource VolumeContainer
	err := ctx.ReadResource("azure-nextgen:storsimple/v20170601:VolumeContainer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VolumeContainer resources.
type volumeContainerState struct {
	// The bandwidth-rate set on the volume container.
	BandWidthRateInMbps *int `pulumi:"bandWidthRateInMbps"`
	// The ID of the bandwidth setting associated with the volume container.
	BandwidthSettingId *string `pulumi:"bandwidthSettingId"`
	// The key used to encrypt data in the volume container. It is required when property 'EncryptionStatus' is "Enabled".
	EncryptionKey *AsymmetricEncryptedSecretResponse `pulumi:"encryptionKey"`
	// The flag to denote whether encryption is enabled or not.
	EncryptionStatus *string `pulumi:"encryptionStatus"`
	// The Kind of the object. Currently only Series8000 is supported
	Kind *string `pulumi:"kind"`
	// The name of the object.
	Name *string `pulumi:"name"`
	// The owner ship status of the volume container. Only when the status is "NotOwned", the delete operation on the volume container is permitted.
	OwnerShipStatus *string `pulumi:"ownerShipStatus"`
	// The path ID of storage account associated with the volume container.
	StorageAccountCredentialId *string `pulumi:"storageAccountCredentialId"`
	// The total cloud storage for the volume container.
	TotalCloudStorageUsageInBytes *float64 `pulumi:"totalCloudStorageUsageInBytes"`
	// The hierarchical type of the object.
	Type *string `pulumi:"type"`
	// The number of volumes in the volume Container.
	VolumeCount *int `pulumi:"volumeCount"`
}

type VolumeContainerState struct {
	// The bandwidth-rate set on the volume container.
	BandWidthRateInMbps pulumi.IntPtrInput
	// The ID of the bandwidth setting associated with the volume container.
	BandwidthSettingId pulumi.StringPtrInput
	// The key used to encrypt data in the volume container. It is required when property 'EncryptionStatus' is "Enabled".
	EncryptionKey AsymmetricEncryptedSecretResponsePtrInput
	// The flag to denote whether encryption is enabled or not.
	EncryptionStatus pulumi.StringPtrInput
	// The Kind of the object. Currently only Series8000 is supported
	Kind pulumi.StringPtrInput
	// The name of the object.
	Name pulumi.StringPtrInput
	// The owner ship status of the volume container. Only when the status is "NotOwned", the delete operation on the volume container is permitted.
	OwnerShipStatus pulumi.StringPtrInput
	// The path ID of storage account associated with the volume container.
	StorageAccountCredentialId pulumi.StringPtrInput
	// The total cloud storage for the volume container.
	TotalCloudStorageUsageInBytes pulumi.Float64PtrInput
	// The hierarchical type of the object.
	Type pulumi.StringPtrInput
	// The number of volumes in the volume Container.
	VolumeCount pulumi.IntPtrInput
}

func (VolumeContainerState) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeContainerState)(nil)).Elem()
}

type volumeContainerArgs struct {
	// The bandwidth-rate set on the volume container.
	BandWidthRateInMbps *int `pulumi:"bandWidthRateInMbps"`
	// The ID of the bandwidth setting associated with the volume container.
	BandwidthSettingId *string `pulumi:"bandwidthSettingId"`
	// The device name
	DeviceName string `pulumi:"deviceName"`
	// The key used to encrypt data in the volume container. It is required when property 'EncryptionStatus' is "Enabled".
	EncryptionKey *AsymmetricEncryptedSecret `pulumi:"encryptionKey"`
	// The Kind of the object. Currently only Series8000 is supported
	Kind *string `pulumi:"kind"`
	// The manager name
	ManagerName string `pulumi:"managerName"`
	// The resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The path ID of storage account associated with the volume container.
	StorageAccountCredentialId string `pulumi:"storageAccountCredentialId"`
	// The name of the volume container.
	VolumeContainerName string `pulumi:"volumeContainerName"`
}

// The set of arguments for constructing a VolumeContainer resource.
type VolumeContainerArgs struct {
	// The bandwidth-rate set on the volume container.
	BandWidthRateInMbps pulumi.IntPtrInput
	// The ID of the bandwidth setting associated with the volume container.
	BandwidthSettingId pulumi.StringPtrInput
	// The device name
	DeviceName pulumi.StringInput
	// The key used to encrypt data in the volume container. It is required when property 'EncryptionStatus' is "Enabled".
	EncryptionKey AsymmetricEncryptedSecretPtrInput
	// The Kind of the object. Currently only Series8000 is supported
	Kind Kind
	// The manager name
	ManagerName pulumi.StringInput
	// The resource group name
	ResourceGroupName pulumi.StringInput
	// The path ID of storage account associated with the volume container.
	StorageAccountCredentialId pulumi.StringInput
	// The name of the volume container.
	VolumeContainerName pulumi.StringInput
}

func (VolumeContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeContainerArgs)(nil)).Elem()
}

type VolumeContainerInput interface {
	pulumi.Input

	ToVolumeContainerOutput() VolumeContainerOutput
	ToVolumeContainerOutputWithContext(ctx context.Context) VolumeContainerOutput
}

func (*VolumeContainer) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeContainer)(nil))
}

func (i *VolumeContainer) ToVolumeContainerOutput() VolumeContainerOutput {
	return i.ToVolumeContainerOutputWithContext(context.Background())
}

func (i *VolumeContainer) ToVolumeContainerOutputWithContext(ctx context.Context) VolumeContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeContainerOutput)
}

type VolumeContainerOutput struct {
	*pulumi.OutputState
}

func (VolumeContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeContainer)(nil))
}

func (o VolumeContainerOutput) ToVolumeContainerOutput() VolumeContainerOutput {
	return o
}

func (o VolumeContainerOutput) ToVolumeContainerOutputWithContext(ctx context.Context) VolumeContainerOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(VolumeContainerOutput{})
}
