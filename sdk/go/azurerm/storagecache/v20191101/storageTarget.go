// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191101

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A storage system being cached by a Cache.
type StorageTarget struct {
	pulumi.CustomResourceState

	// Name of the Storage Target.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the Storage Target.
	Properties StorageTargetResponsePropertiesOutput `pulumi:"properties"`
	// Type of the Storage Target; Microsoft.StorageCache/Cache/StorageTarget
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewStorageTarget registers a new resource with the given unique name, arguments, and options.
func NewStorageTarget(ctx *pulumi.Context,
	name string, args *StorageTargetArgs, opts ...pulumi.ResourceOption) (*StorageTarget, error) {
	if args == nil || args.CacheName == nil {
		return nil, errors.New("missing required argument 'CacheName'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &StorageTargetArgs{}
	}
	var resource StorageTarget
	err := ctx.RegisterResource("azurerm:storagecache/v20191101:StorageTarget", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStorageTarget gets an existing StorageTarget resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStorageTarget(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageTargetState, opts ...pulumi.ResourceOption) (*StorageTarget, error) {
	var resource StorageTarget
	err := ctx.ReadResource("azurerm:storagecache/v20191101:StorageTarget", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StorageTarget resources.
type storageTargetState struct {
	// Name of the Storage Target.
	Name *string `pulumi:"name"`
	// Properties of the Storage Target.
	Properties *StorageTargetResponseProperties `pulumi:"properties"`
	// Type of the Storage Target; Microsoft.StorageCache/Cache/StorageTarget
	Type *string `pulumi:"type"`
}

type StorageTargetState struct {
	// Name of the Storage Target.
	Name pulumi.StringPtrInput
	// Properties of the Storage Target.
	Properties StorageTargetResponsePropertiesPtrInput
	// Type of the Storage Target; Microsoft.StorageCache/Cache/StorageTarget
	Type pulumi.StringPtrInput
}

func (StorageTargetState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageTargetState)(nil)).Elem()
}

type storageTargetArgs struct {
	// Name of Cache.
	CacheName string `pulumi:"cacheName"`
	// Properties when targetType is clfs.
	Clfs *ClfsTarget `pulumi:"clfs"`
	// List of Cache namespace junctions to target for namespace associations.
	Junctions []NamespaceJunction `pulumi:"junctions"`
	// Name of the Storage Target.
	Name string `pulumi:"name"`
	// Properties when targetType is nfs3.
	Nfs3 *Nfs3Target `pulumi:"nfs3"`
	// ARM provisioning state, see https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/Addendum.md#provisioningstate-property
	ProvisioningState *string `pulumi:"provisioningState"`
	// Target resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Type of the Storage Target.
	TargetType *string `pulumi:"targetType"`
	// Properties when targetType is unknown.
	Unknown *UnknownTarget `pulumi:"unknown"`
}

// The set of arguments for constructing a StorageTarget resource.
type StorageTargetArgs struct {
	// Name of Cache.
	CacheName pulumi.StringInput
	// Properties when targetType is clfs.
	Clfs ClfsTargetPtrInput
	// List of Cache namespace junctions to target for namespace associations.
	Junctions NamespaceJunctionArrayInput
	// Name of the Storage Target.
	Name pulumi.StringInput
	// Properties when targetType is nfs3.
	Nfs3 Nfs3TargetPtrInput
	// ARM provisioning state, see https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/Addendum.md#provisioningstate-property
	ProvisioningState pulumi.StringPtrInput
	// Target resource group.
	ResourceGroupName pulumi.StringInput
	// Type of the Storage Target.
	TargetType pulumi.StringPtrInput
	// Properties when targetType is unknown.
	Unknown UnknownTargetPtrInput
}

func (StorageTargetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageTargetArgs)(nil)).Elem()
}
