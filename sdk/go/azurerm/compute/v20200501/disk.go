// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200501

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Disk resource.
type Disk struct {
	pulumi.CustomResourceState

	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// A relative URI containing the ID of the VM that has the disk attached.
	ManagedBy pulumi.StringOutput `pulumi:"managedBy"`
	// List of relative URIs containing the IDs of the VMs that have the disk attached. maxShares should be set to a value greater than one for disks to allow attaching them to multiple VMs.
	ManagedByExtended pulumi.StringArrayOutput `pulumi:"managedByExtended"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Disk resource properties.
	Properties DiskPropertiesResponseOutput `pulumi:"properties"`
	// The disks sku name. Can be Standard_LRS, Premium_LRS, StandardSSD_LRS, or UltraSSD_LRS.
	Sku DiskSkuResponsePtrOutput `pulumi:"sku"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// The Logical zone list for Disk.
	Zones pulumi.StringArrayOutput `pulumi:"zones"`
}

// NewDisk registers a new resource with the given unique name, arguments, and options.
func NewDisk(ctx *pulumi.Context,
	name string, args *DiskArgs, opts ...pulumi.ResourceOption) (*Disk, error) {
	if args == nil || args.CreationData == nil {
		return nil, errors.New("missing required argument 'CreationData'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &DiskArgs{}
	}
	var resource Disk
	err := ctx.RegisterResource("azurerm:compute/v20200501:Disk", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDisk gets an existing Disk resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDisk(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DiskState, opts ...pulumi.ResourceOption) (*Disk, error) {
	var resource Disk
	err := ctx.ReadResource("azurerm:compute/v20200501:Disk", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Disk resources.
type diskState struct {
	// Resource location
	Location *string `pulumi:"location"`
	// A relative URI containing the ID of the VM that has the disk attached.
	ManagedBy *string `pulumi:"managedBy"`
	// List of relative URIs containing the IDs of the VMs that have the disk attached. maxShares should be set to a value greater than one for disks to allow attaching them to multiple VMs.
	ManagedByExtended []string `pulumi:"managedByExtended"`
	// Resource name
	Name *string `pulumi:"name"`
	// Disk resource properties.
	Properties *DiskPropertiesResponse `pulumi:"properties"`
	// The disks sku name. Can be Standard_LRS, Premium_LRS, StandardSSD_LRS, or UltraSSD_LRS.
	Sku *DiskSkuResponse `pulumi:"sku"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
	// The Logical zone list for Disk.
	Zones []string `pulumi:"zones"`
}

type DiskState struct {
	// Resource location
	Location pulumi.StringPtrInput
	// A relative URI containing the ID of the VM that has the disk attached.
	ManagedBy pulumi.StringPtrInput
	// List of relative URIs containing the IDs of the VMs that have the disk attached. maxShares should be set to a value greater than one for disks to allow attaching them to multiple VMs.
	ManagedByExtended pulumi.StringArrayInput
	// Resource name
	Name pulumi.StringPtrInput
	// Disk resource properties.
	Properties DiskPropertiesResponsePtrInput
	// The disks sku name. Can be Standard_LRS, Premium_LRS, StandardSSD_LRS, or UltraSSD_LRS.
	Sku DiskSkuResponsePtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
	// The Logical zone list for Disk.
	Zones pulumi.StringArrayInput
}

func (DiskState) ElementType() reflect.Type {
	return reflect.TypeOf((*diskState)(nil)).Elem()
}

type diskArgs struct {
	// Disk source information. CreationData information cannot be changed after the disk has been created.
	CreationData CreationData `pulumi:"creationData"`
	// ARM id of the DiskAccess resource for using private endpoints on disks.
	DiskAccessId *string `pulumi:"diskAccessId"`
	// The total number of IOPS that will be allowed across all VMs mounting the shared disk as ReadOnly. One operation can transfer between 4k and 256k bytes.
	DiskIOPSReadOnly *int `pulumi:"diskIOPSReadOnly"`
	// The number of IOPS allowed for this disk; only settable for UltraSSD disks. One operation can transfer between 4k and 256k bytes.
	DiskIOPSReadWrite *int `pulumi:"diskIOPSReadWrite"`
	// The total throughput (MBps) that will be allowed across all VMs mounting the shared disk as ReadOnly. MBps means millions of bytes per second - MB here uses the ISO notation, of powers of 10.
	DiskMBpsReadOnly *int `pulumi:"diskMBpsReadOnly"`
	// The bandwidth allowed for this disk; only settable for UltraSSD disks. MBps means millions of bytes per second - MB here uses the ISO notation, of powers of 10.
	DiskMBpsReadWrite *int `pulumi:"diskMBpsReadWrite"`
	// If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
	DiskSizeGB *int `pulumi:"diskSizeGB"`
	// Encryption property can be used to encrypt data at rest with customer managed keys or platform managed keys.
	Encryption *Encryption `pulumi:"encryption"`
	// Encryption settings collection used for Azure Disk Encryption, can contain multiple encryption settings per disk or snapshot.
	EncryptionSettingsCollection *EncryptionSettingsCollection `pulumi:"encryptionSettingsCollection"`
	// The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
	HyperVGeneration *string `pulumi:"hyperVGeneration"`
	// Resource location
	Location string `pulumi:"location"`
	// The maximum number of VMs that can attach to the disk at the same time. Value greater than one indicates a disk that can be mounted on multiple VMs at the same time.
	MaxShares *int `pulumi:"maxShares"`
	// The name of the managed disk that is being created. The name can't be changed after the disk is created. Supported characters for the name are a-z, A-Z, 0-9 and _. The maximum name length is 80 characters.
	Name string `pulumi:"name"`
	// Policy for accessing the disk via network.
	NetworkAccessPolicy *string `pulumi:"networkAccessPolicy"`
	// The Operating System type.
	OsType *string `pulumi:"osType"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The disks sku name. Can be Standard_LRS, Premium_LRS, StandardSSD_LRS, or UltraSSD_LRS.
	Sku *DiskSku `pulumi:"sku"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// The Logical zone list for Disk.
	Zones []string `pulumi:"zones"`
}

// The set of arguments for constructing a Disk resource.
type DiskArgs struct {
	// Disk source information. CreationData information cannot be changed after the disk has been created.
	CreationData CreationDataInput
	// ARM id of the DiskAccess resource for using private endpoints on disks.
	DiskAccessId pulumi.StringPtrInput
	// The total number of IOPS that will be allowed across all VMs mounting the shared disk as ReadOnly. One operation can transfer between 4k and 256k bytes.
	DiskIOPSReadOnly pulumi.IntPtrInput
	// The number of IOPS allowed for this disk; only settable for UltraSSD disks. One operation can transfer between 4k and 256k bytes.
	DiskIOPSReadWrite pulumi.IntPtrInput
	// The total throughput (MBps) that will be allowed across all VMs mounting the shared disk as ReadOnly. MBps means millions of bytes per second - MB here uses the ISO notation, of powers of 10.
	DiskMBpsReadOnly pulumi.IntPtrInput
	// The bandwidth allowed for this disk; only settable for UltraSSD disks. MBps means millions of bytes per second - MB here uses the ISO notation, of powers of 10.
	DiskMBpsReadWrite pulumi.IntPtrInput
	// If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
	DiskSizeGB pulumi.IntPtrInput
	// Encryption property can be used to encrypt data at rest with customer managed keys or platform managed keys.
	Encryption EncryptionPtrInput
	// Encryption settings collection used for Azure Disk Encryption, can contain multiple encryption settings per disk or snapshot.
	EncryptionSettingsCollection EncryptionSettingsCollectionPtrInput
	// The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
	HyperVGeneration pulumi.StringPtrInput
	// Resource location
	Location pulumi.StringInput
	// The maximum number of VMs that can attach to the disk at the same time. Value greater than one indicates a disk that can be mounted on multiple VMs at the same time.
	MaxShares pulumi.IntPtrInput
	// The name of the managed disk that is being created. The name can't be changed after the disk is created. Supported characters for the name are a-z, A-Z, 0-9 and _. The maximum name length is 80 characters.
	Name pulumi.StringInput
	// Policy for accessing the disk via network.
	NetworkAccessPolicy pulumi.StringPtrInput
	// The Operating System type.
	OsType pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The disks sku name. Can be Standard_LRS, Premium_LRS, StandardSSD_LRS, or UltraSSD_LRS.
	Sku DiskSkuPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// The Logical zone list for Disk.
	Zones pulumi.StringArrayInput
}

func (DiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*diskArgs)(nil)).Elem()
}
