// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180110

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Recovery plan details.
//
// ## Example Usage
// ### Creates a recovery plan with the given details.
//
// ```go
// package main
//
// import (
// 	recoveryservices "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/recoveryservices/v20180110"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := recoveryservices.NewReplicationRecoveryPlan(ctx, "replicationRecoveryPlan", &recoveryservices.ReplicationRecoveryPlanArgs{
// 			RecoveryPlanName:  pulumi.String("RPtest1"),
// 			ResourceGroupName: pulumi.String("resourceGroupPS1"),
// 			ResourceName:      pulumi.String("vault1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type ReplicationRecoveryPlan struct {
	pulumi.CustomResourceState

	// Resource Location
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource Name
	Name pulumi.StringOutput `pulumi:"name"`
	// The custom details.
	Properties RecoveryPlanPropertiesResponseOutput `pulumi:"properties"`
	// Resource Type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewReplicationRecoveryPlan registers a new resource with the given unique name, arguments, and options.
func NewReplicationRecoveryPlan(ctx *pulumi.Context,
	name string, args *ReplicationRecoveryPlanArgs, opts ...pulumi.ResourceOption) (*ReplicationRecoveryPlan, error) {
	if args == nil || args.Properties == nil {
		return nil, errors.New("missing required argument 'Properties'")
	}
	if args == nil || args.RecoveryPlanName == nil {
		return nil, errors.New("missing required argument 'RecoveryPlanName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ResourceName == nil {
		return nil, errors.New("missing required argument 'ResourceName'")
	}
	if args == nil {
		args = &ReplicationRecoveryPlanArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:recoveryservices/latest:ReplicationRecoveryPlan"),
		},
		{
			Type: pulumi.String("azurerm:recoveryservices/v20160810:ReplicationRecoveryPlan"),
		},
		{
			Type: pulumi.String("azurerm:recoveryservices/v20180710:ReplicationRecoveryPlan"),
		},
	})
	opts = append(opts, aliases)
	var resource ReplicationRecoveryPlan
	err := ctx.RegisterResource("azurerm:recoveryservices/v20180110:ReplicationRecoveryPlan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplicationRecoveryPlan gets an existing ReplicationRecoveryPlan resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplicationRecoveryPlan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicationRecoveryPlanState, opts ...pulumi.ResourceOption) (*ReplicationRecoveryPlan, error) {
	var resource ReplicationRecoveryPlan
	err := ctx.ReadResource("azurerm:recoveryservices/v20180110:ReplicationRecoveryPlan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReplicationRecoveryPlan resources.
type replicationRecoveryPlanState struct {
	// Resource Location
	Location *string `pulumi:"location"`
	// Resource Name
	Name *string `pulumi:"name"`
	// The custom details.
	Properties *RecoveryPlanPropertiesResponse `pulumi:"properties"`
	// Resource Type
	Type *string `pulumi:"type"`
}

type ReplicationRecoveryPlanState struct {
	// Resource Location
	Location pulumi.StringPtrInput
	// Resource Name
	Name pulumi.StringPtrInput
	// The custom details.
	Properties RecoveryPlanPropertiesResponsePtrInput
	// Resource Type
	Type pulumi.StringPtrInput
}

func (ReplicationRecoveryPlanState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationRecoveryPlanState)(nil)).Elem()
}

type replicationRecoveryPlanArgs struct {
	// Recovery plan creation properties.
	Properties CreateRecoveryPlanInputProperties `pulumi:"properties"`
	// Recovery plan name.
	RecoveryPlanName string `pulumi:"recoveryPlanName"`
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the recovery services vault.
	ResourceName string `pulumi:"resourceName"`
}

// The set of arguments for constructing a ReplicationRecoveryPlan resource.
type ReplicationRecoveryPlanArgs struct {
	// Recovery plan creation properties.
	Properties CreateRecoveryPlanInputPropertiesInput
	// Recovery plan name.
	RecoveryPlanName pulumi.StringInput
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName pulumi.StringInput
	// The name of the recovery services vault.
	ResourceName pulumi.StringInput
}

func (ReplicationRecoveryPlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationRecoveryPlanArgs)(nil)).Elem()
}
