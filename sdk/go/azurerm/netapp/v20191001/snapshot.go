// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191001

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Snapshot of a Volume
type Snapshot struct {
	pulumi.CustomResourceState

	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Snapshot Properties
	Properties SnapshotPropertiesResponseOutput `pulumi:"properties"`
	// Resource tags
	Tags ResourceTagsResponsePtrOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSnapshot registers a new resource with the given unique name, arguments, and options.
func NewSnapshot(ctx *pulumi.Context,
	name string, args *SnapshotArgs, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	if args == nil || args.AccountName == nil {
		return nil, errors.New("missing required argument 'AccountName'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.PoolName == nil {
		return nil, errors.New("missing required argument 'PoolName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.VolumeName == nil {
		return nil, errors.New("missing required argument 'VolumeName'")
	}
	if args == nil {
		args = &SnapshotArgs{}
	}
	var resource Snapshot
	err := ctx.RegisterResource("azurerm:netapp/v20191001:Snapshot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnapshot gets an existing Snapshot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnapshot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotState, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	var resource Snapshot
	err := ctx.ReadResource("azurerm:netapp/v20191001:Snapshot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Snapshot resources.
type snapshotState struct {
	// Resource location
	Location *string `pulumi:"location"`
	// Resource name
	Name *string `pulumi:"name"`
	// Snapshot Properties
	Properties *SnapshotPropertiesResponse `pulumi:"properties"`
	// Resource tags
	Tags *ResourceTagsResponse `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

type SnapshotState struct {
	// Resource location
	Location pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// Snapshot Properties
	Properties SnapshotPropertiesResponsePtrInput
	// Resource tags
	Tags ResourceTagsResponsePtrInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (SnapshotState) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotState)(nil)).Elem()
}

type snapshotArgs struct {
	// The name of the NetApp account
	AccountName string `pulumi:"accountName"`
	// Resource location
	Location string `pulumi:"location"`
	// The name of the mount target
	Name string `pulumi:"name"`
	// The name of the capacity pool
	PoolName string `pulumi:"poolName"`
	// Snapshot Properties
	Properties *SnapshotProperties `pulumi:"properties"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags
	Tags *ResourceTags `pulumi:"tags"`
	// The name of the volume
	VolumeName string `pulumi:"volumeName"`
}

// The set of arguments for constructing a Snapshot resource.
type SnapshotArgs struct {
	// The name of the NetApp account
	AccountName pulumi.StringInput
	// Resource location
	Location pulumi.StringInput
	// The name of the mount target
	Name pulumi.StringInput
	// The name of the capacity pool
	PoolName pulumi.StringInput
	// Snapshot Properties
	Properties SnapshotPropertiesPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags
	Tags ResourceTagsPtrInput
	// The name of the volume
	VolumeName pulumi.StringInput
}

func (SnapshotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotArgs)(nil)).Elem()
}