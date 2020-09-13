// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Protection profile details.
//
// ## Example Usage
// ### Creates the policy.
//
// ```go
// package main
//
// import (
// 	recoveryservices "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/recoveryservices/latest"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := recoveryservices.NewReplicationPolicy(ctx, "replicationPolicy", &recoveryservices.ReplicationPolicyArgs{
// 			PolicyName:        pulumi.String("protectionprofile1"),
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
type ReplicationPolicy struct {
	pulumi.CustomResourceState

	// Resource Location
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource Name
	Name pulumi.StringOutput `pulumi:"name"`
	// The custom data.
	Properties PolicyPropertiesResponseOutput `pulumi:"properties"`
	// Resource Type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewReplicationPolicy registers a new resource with the given unique name, arguments, and options.
func NewReplicationPolicy(ctx *pulumi.Context,
	name string, args *ReplicationPolicyArgs, opts ...pulumi.ResourceOption) (*ReplicationPolicy, error) {
	if args == nil || args.PolicyName == nil {
		return nil, errors.New("missing required argument 'PolicyName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ResourceName == nil {
		return nil, errors.New("missing required argument 'ResourceName'")
	}
	if args == nil {
		args = &ReplicationPolicyArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:recoveryservices/v20160810:ReplicationPolicy"),
		},
		{
			Type: pulumi.String("azurerm:recoveryservices/v20180110:ReplicationPolicy"),
		},
		{
			Type: pulumi.String("azurerm:recoveryservices/v20180710:ReplicationPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource ReplicationPolicy
	err := ctx.RegisterResource("azurerm:recoveryservices/latest:ReplicationPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplicationPolicy gets an existing ReplicationPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplicationPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicationPolicyState, opts ...pulumi.ResourceOption) (*ReplicationPolicy, error) {
	var resource ReplicationPolicy
	err := ctx.ReadResource("azurerm:recoveryservices/latest:ReplicationPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReplicationPolicy resources.
type replicationPolicyState struct {
	// Resource Location
	Location *string `pulumi:"location"`
	// Resource Name
	Name *string `pulumi:"name"`
	// The custom data.
	Properties *PolicyPropertiesResponse `pulumi:"properties"`
	// Resource Type
	Type *string `pulumi:"type"`
}

type ReplicationPolicyState struct {
	// Resource Location
	Location pulumi.StringPtrInput
	// Resource Name
	Name pulumi.StringPtrInput
	// The custom data.
	Properties PolicyPropertiesResponsePtrInput
	// Resource Type
	Type pulumi.StringPtrInput
}

func (ReplicationPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationPolicyState)(nil)).Elem()
}

type replicationPolicyArgs struct {
	// Replication policy name
	PolicyName string `pulumi:"policyName"`
	// Policy creation properties.
	Properties *CreatePolicyInputProperties `pulumi:"properties"`
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the recovery services vault.
	ResourceName string `pulumi:"resourceName"`
}

// The set of arguments for constructing a ReplicationPolicy resource.
type ReplicationPolicyArgs struct {
	// Replication policy name
	PolicyName pulumi.StringInput
	// Policy creation properties.
	Properties CreatePolicyInputPropertiesPtrInput
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName pulumi.StringInput
	// The name of the recovery services vault.
	ResourceName pulumi.StringInput
}

func (ReplicationPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationPolicyArgs)(nil)).Elem()
}
