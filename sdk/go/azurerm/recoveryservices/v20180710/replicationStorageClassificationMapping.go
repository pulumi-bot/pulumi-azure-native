// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180710

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Storage mapping object.
//
// ## Example Usage
// ### Create storage classification mapping.
//
// ```go
// package main
//
// import (
// 	recoveryservices "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/recoveryservices/v20180710"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := recoveryservices.NewReplicationStorageClassificationMapping(ctx, "replicationStorageClassificationMapping", &recoveryservices.ReplicationStorageClassificationMappingArgs{
// 			FabricName:                       pulumi.String("2a48e3770ac08aa2be8bfbd94fcfb1cbf2dcc487b78fb9d3bd778304441b06a0"),
// 			ResourceGroupName:                pulumi.String("resourceGroupPS1"),
// 			ResourceName:                     pulumi.String("vault1"),
// 			StorageClassificationMappingName: pulumi.String("testStorageMapping"),
// 			StorageClassificationName:        pulumi.String("8891569e-aaef-4a46-a4a0-78c14f2d7b09"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type ReplicationStorageClassificationMapping struct {
	pulumi.CustomResourceState

	// Resource Location
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource Name
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the storage mapping object.
	Properties StorageClassificationMappingPropertiesResponseOutput `pulumi:"properties"`
	// Resource Type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewReplicationStorageClassificationMapping registers a new resource with the given unique name, arguments, and options.
func NewReplicationStorageClassificationMapping(ctx *pulumi.Context,
	name string, args *ReplicationStorageClassificationMappingArgs, opts ...pulumi.ResourceOption) (*ReplicationStorageClassificationMapping, error) {
	if args == nil || args.FabricName == nil {
		return nil, errors.New("missing required argument 'FabricName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ResourceName == nil {
		return nil, errors.New("missing required argument 'ResourceName'")
	}
	if args == nil || args.StorageClassificationMappingName == nil {
		return nil, errors.New("missing required argument 'StorageClassificationMappingName'")
	}
	if args == nil || args.StorageClassificationName == nil {
		return nil, errors.New("missing required argument 'StorageClassificationName'")
	}
	if args == nil {
		args = &ReplicationStorageClassificationMappingArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:recoveryservices/latest:ReplicationStorageClassificationMapping"),
		},
		{
			Type: pulumi.String("azurerm:recoveryservices/v20160810:ReplicationStorageClassificationMapping"),
		},
		{
			Type: pulumi.String("azurerm:recoveryservices/v20180110:ReplicationStorageClassificationMapping"),
		},
	})
	opts = append(opts, aliases)
	var resource ReplicationStorageClassificationMapping
	err := ctx.RegisterResource("azurerm:recoveryservices/v20180710:ReplicationStorageClassificationMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplicationStorageClassificationMapping gets an existing ReplicationStorageClassificationMapping resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplicationStorageClassificationMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicationStorageClassificationMappingState, opts ...pulumi.ResourceOption) (*ReplicationStorageClassificationMapping, error) {
	var resource ReplicationStorageClassificationMapping
	err := ctx.ReadResource("azurerm:recoveryservices/v20180710:ReplicationStorageClassificationMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReplicationStorageClassificationMapping resources.
type replicationStorageClassificationMappingState struct {
	// Resource Location
	Location *string `pulumi:"location"`
	// Resource Name
	Name *string `pulumi:"name"`
	// Properties of the storage mapping object.
	Properties *StorageClassificationMappingPropertiesResponse `pulumi:"properties"`
	// Resource Type
	Type *string `pulumi:"type"`
}

type ReplicationStorageClassificationMappingState struct {
	// Resource Location
	Location pulumi.StringPtrInput
	// Resource Name
	Name pulumi.StringPtrInput
	// Properties of the storage mapping object.
	Properties StorageClassificationMappingPropertiesResponsePtrInput
	// Resource Type
	Type pulumi.StringPtrInput
}

func (ReplicationStorageClassificationMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationStorageClassificationMappingState)(nil)).Elem()
}

type replicationStorageClassificationMappingArgs struct {
	// Fabric name.
	FabricName string `pulumi:"fabricName"`
	// Storage mapping input properties.
	Properties *StorageMappingInputProperties `pulumi:"properties"`
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the recovery services vault.
	ResourceName string `pulumi:"resourceName"`
	// Storage classification mapping name.
	StorageClassificationMappingName string `pulumi:"storageClassificationMappingName"`
	// Storage classification name.
	StorageClassificationName string `pulumi:"storageClassificationName"`
}

// The set of arguments for constructing a ReplicationStorageClassificationMapping resource.
type ReplicationStorageClassificationMappingArgs struct {
	// Fabric name.
	FabricName pulumi.StringInput
	// Storage mapping input properties.
	Properties StorageMappingInputPropertiesPtrInput
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName pulumi.StringInput
	// The name of the recovery services vault.
	ResourceName pulumi.StringInput
	// Storage classification mapping name.
	StorageClassificationMappingName pulumi.StringInput
	// Storage classification name.
	StorageClassificationName pulumi.StringInput
}

func (ReplicationStorageClassificationMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationStorageClassificationMappingArgs)(nil)).Elem()
}
