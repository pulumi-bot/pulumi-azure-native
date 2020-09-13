// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160810

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// vCenter definition.
//
// ## Example Usage
// ### Add vCenter.
//
// ```go
// package main
//
// import (
// 	recoveryservices "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/recoveryservices/v20160810"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := recoveryservices.NewReplicationvCenter(ctx, "replicationvCenter", &recoveryservices.ReplicationvCenterArgs{
// 			FabricName:        pulumi.String("MadhaviFabric"),
// 			ResourceGroupName: pulumi.String("MadhaviVRG"),
// 			ResourceName:      pulumi.String("MadhaviVault"),
// 			VCenterName:       pulumi.String("esx-78"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type ReplicationvCenter struct {
	pulumi.CustomResourceState

	// Resource Location
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource Name
	Name pulumi.StringOutput `pulumi:"name"`
	// VCenter related data.
	Properties VCenterPropertiesResponseOutput `pulumi:"properties"`
	// Resource Type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewReplicationvCenter registers a new resource with the given unique name, arguments, and options.
func NewReplicationvCenter(ctx *pulumi.Context,
	name string, args *ReplicationvCenterArgs, opts ...pulumi.ResourceOption) (*ReplicationvCenter, error) {
	if args == nil || args.FabricName == nil {
		return nil, errors.New("missing required argument 'FabricName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ResourceName == nil {
		return nil, errors.New("missing required argument 'ResourceName'")
	}
	if args == nil || args.VCenterName == nil {
		return nil, errors.New("missing required argument 'VCenterName'")
	}
	if args == nil {
		args = &ReplicationvCenterArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:recoveryservices/latest:ReplicationvCenter"),
		},
		{
			Type: pulumi.String("azurerm:recoveryservices/v20180110:ReplicationvCenter"),
		},
		{
			Type: pulumi.String("azurerm:recoveryservices/v20180710:ReplicationvCenter"),
		},
	})
	opts = append(opts, aliases)
	var resource ReplicationvCenter
	err := ctx.RegisterResource("azurerm:recoveryservices/v20160810:ReplicationvCenter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplicationvCenter gets an existing ReplicationvCenter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplicationvCenter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicationvCenterState, opts ...pulumi.ResourceOption) (*ReplicationvCenter, error) {
	var resource ReplicationvCenter
	err := ctx.ReadResource("azurerm:recoveryservices/v20160810:ReplicationvCenter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReplicationvCenter resources.
type replicationvCenterState struct {
	// Resource Location
	Location *string `pulumi:"location"`
	// Resource Name
	Name *string `pulumi:"name"`
	// VCenter related data.
	Properties *VCenterPropertiesResponse `pulumi:"properties"`
	// Resource Type
	Type *string `pulumi:"type"`
}

type ReplicationvCenterState struct {
	// Resource Location
	Location pulumi.StringPtrInput
	// Resource Name
	Name pulumi.StringPtrInput
	// VCenter related data.
	Properties VCenterPropertiesResponsePtrInput
	// Resource Type
	Type pulumi.StringPtrInput
}

func (ReplicationvCenterState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationvCenterState)(nil)).Elem()
}

type replicationvCenterArgs struct {
	// Fabric name.
	FabricName string `pulumi:"fabricName"`
	// The properties of an add vCenter request.
	Properties *AddVCenterRequestProperties `pulumi:"properties"`
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the recovery services vault.
	ResourceName string `pulumi:"resourceName"`
	// vCenter name.
	VCenterName string `pulumi:"vCenterName"`
}

// The set of arguments for constructing a ReplicationvCenter resource.
type ReplicationvCenterArgs struct {
	// Fabric name.
	FabricName pulumi.StringInput
	// The properties of an add vCenter request.
	Properties AddVCenterRequestPropertiesPtrInput
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName pulumi.StringInput
	// The name of the recovery services vault.
	ResourceName pulumi.StringInput
	// vCenter name.
	VCenterName pulumi.StringInput
}

func (ReplicationvCenterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationvCenterArgs)(nil)).Elem()
}
