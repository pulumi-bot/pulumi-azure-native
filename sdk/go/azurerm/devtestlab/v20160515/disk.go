// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160515

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A Disk.
type Disk struct {
	pulumi.CustomResourceState

	// The location of the resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The properties of the resource.
	Properties DiskPropertiesResponseOutput `pulumi:"properties"`
	// The tags of the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDisk registers a new resource with the given unique name, arguments, and options.
func NewDisk(ctx *pulumi.Context,
	name string, args *DiskArgs, opts ...pulumi.ResourceOption) (*Disk, error) {
	if args == nil || args.LabName == nil {
		return nil, errors.New("missing required argument 'LabName'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.UserName == nil {
		return nil, errors.New("missing required argument 'UserName'")
	}
	if args == nil {
		args = &DiskArgs{}
	}
	var resource Disk
	err := ctx.RegisterResource("azurerm:devtestlab/v20160515:Disk", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azurerm:devtestlab/v20160515:Disk", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Disk resources.
type diskState struct {
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name *string `pulumi:"name"`
	// The properties of the resource.
	Properties *DiskPropertiesResponse `pulumi:"properties"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

type DiskState struct {
	// The location of the resource.
	Location pulumi.StringPtrInput
	// The name of the resource.
	Name pulumi.StringPtrInput
	// The properties of the resource.
	Properties DiskPropertiesResponsePtrInput
	// The tags of the resource.
	Tags pulumi.StringMapInput
	// The type of the resource.
	Type pulumi.StringPtrInput
}

func (DiskState) ElementType() reflect.Type {
	return reflect.TypeOf((*diskState)(nil)).Elem()
}

type diskArgs struct {
	// When backed by a blob, the name of the VHD blob without extension.
	DiskBlobName *string `pulumi:"diskBlobName"`
	// The size of the disk in Gibibytes.
	DiskSizeGiB *int `pulumi:"diskSizeGiB"`
	// The storage type for the disk (i.e. Standard, Premium).
	DiskType *string `pulumi:"diskType"`
	// When backed by a blob, the URI of underlying blob.
	DiskUri *string `pulumi:"diskUri"`
	// The host caching policy of the disk (i.e. None, ReadOnly, ReadWrite).
	HostCaching *string `pulumi:"hostCaching"`
	// The name of the lab.
	LabName string `pulumi:"labName"`
	// The resource ID of the VM to which this disk is leased.
	LeasedByLabVmId *string `pulumi:"leasedByLabVmId"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// When backed by managed disk, this is the ID of the compute disk resource.
	ManagedDiskId *string `pulumi:"managedDiskId"`
	// The name of the disk.
	Name string `pulumi:"name"`
	// The provisioning status of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The unique immutable identifier of a resource (Guid).
	UniqueIdentifier *string `pulumi:"uniqueIdentifier"`
	// The name of the user profile.
	UserName string `pulumi:"userName"`
}

// The set of arguments for constructing a Disk resource.
type DiskArgs struct {
	// When backed by a blob, the name of the VHD blob without extension.
	DiskBlobName pulumi.StringPtrInput
	// The size of the disk in Gibibytes.
	DiskSizeGiB pulumi.IntPtrInput
	// The storage type for the disk (i.e. Standard, Premium).
	DiskType pulumi.StringPtrInput
	// When backed by a blob, the URI of underlying blob.
	DiskUri pulumi.StringPtrInput
	// The host caching policy of the disk (i.e. None, ReadOnly, ReadWrite).
	HostCaching pulumi.StringPtrInput
	// The name of the lab.
	LabName pulumi.StringInput
	// The resource ID of the VM to which this disk is leased.
	LeasedByLabVmId pulumi.StringPtrInput
	// The location of the resource.
	Location pulumi.StringPtrInput
	// When backed by managed disk, this is the ID of the compute disk resource.
	ManagedDiskId pulumi.StringPtrInput
	// The name of the disk.
	Name pulumi.StringInput
	// The provisioning status of the resource.
	ProvisioningState pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The tags of the resource.
	Tags pulumi.StringMapInput
	// The unique immutable identifier of a resource (Guid).
	UniqueIdentifier pulumi.StringPtrInput
	// The name of the user profile.
	UserName pulumi.StringInput
}

func (DiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*diskArgs)(nil)).Elem()
}
