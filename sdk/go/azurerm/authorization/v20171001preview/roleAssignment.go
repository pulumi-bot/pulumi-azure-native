// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20171001preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Role Assignments
//
// ## Example Usage
// ### GetConfigurations
//
// ```go
// package main
//
// import (
// 	authorization "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/authorization/v20171001preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := authorization.NewRoleAssignment(ctx, "roleAssignment", &authorization.RoleAssignmentArgs{
// 			CanDelegate:        pulumi.Bool(false),
// 			PrincipalId:        pulumi.String("d93a38bc-d029-4160-bfb0-fbda779ac214"),
// 			RoleAssignmentName: pulumi.String("roleAssignmentName"),
// 			RoleDefinitionId:   pulumi.String("/subscriptions/4004a9fd-d58e-48dc-aeb2-4a4aec58606f/providers/Microsoft.Authorization/roleDefinitions/de139f84-1756-47ae-9be6-808fbbe84772"),
// 			Scope:              pulumi.String("scope"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type RoleAssignment struct {
	pulumi.CustomResourceState

	// The Delegation flag for the role assignment
	CanDelegate pulumi.BoolPtrOutput `pulumi:"canDelegate"`
	// The role assignment name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The principal ID.
	PrincipalId pulumi.StringPtrOutput `pulumi:"principalId"`
	// The role definition ID.
	RoleDefinitionId pulumi.StringPtrOutput `pulumi:"roleDefinitionId"`
	// The role assignment scope.
	Scope pulumi.StringPtrOutput `pulumi:"scope"`
	// The role assignment type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRoleAssignment registers a new resource with the given unique name, arguments, and options.
func NewRoleAssignment(ctx *pulumi.Context,
	name string, args *RoleAssignmentArgs, opts ...pulumi.ResourceOption) (*RoleAssignment, error) {
	if args == nil || args.RoleAssignmentName == nil {
		return nil, errors.New("missing required argument 'RoleAssignmentName'")
	}
	if args == nil || args.Scope == nil {
		return nil, errors.New("missing required argument 'Scope'")
	}
	if args == nil {
		args = &RoleAssignmentArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:authorization/latest:RoleAssignment"),
		},
		{
			Type: pulumi.String("azurerm:authorization/v20150701:RoleAssignment"),
		},
		{
			Type: pulumi.String("azurerm:authorization/v20180101preview:RoleAssignment"),
		},
		{
			Type: pulumi.String("azurerm:authorization/v20180901preview:RoleAssignment"),
		},
		{
			Type: pulumi.String("azurerm:authorization/v20200401preview:RoleAssignment"),
		},
	})
	opts = append(opts, aliases)
	var resource RoleAssignment
	err := ctx.RegisterResource("azurerm:authorization/v20171001preview:RoleAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoleAssignment gets an existing RoleAssignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoleAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoleAssignmentState, opts ...pulumi.ResourceOption) (*RoleAssignment, error) {
	var resource RoleAssignment
	err := ctx.ReadResource("azurerm:authorization/v20171001preview:RoleAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RoleAssignment resources.
type roleAssignmentState struct {
	// The Delegation flag for the role assignment
	CanDelegate *bool `pulumi:"canDelegate"`
	// The role assignment name.
	Name *string `pulumi:"name"`
	// The principal ID.
	PrincipalId *string `pulumi:"principalId"`
	// The role definition ID.
	RoleDefinitionId *string `pulumi:"roleDefinitionId"`
	// The role assignment scope.
	Scope *string `pulumi:"scope"`
	// The role assignment type.
	Type *string `pulumi:"type"`
}

type RoleAssignmentState struct {
	// The Delegation flag for the role assignment
	CanDelegate pulumi.BoolPtrInput
	// The role assignment name.
	Name pulumi.StringPtrInput
	// The principal ID.
	PrincipalId pulumi.StringPtrInput
	// The role definition ID.
	RoleDefinitionId pulumi.StringPtrInput
	// The role assignment scope.
	Scope pulumi.StringPtrInput
	// The role assignment type.
	Type pulumi.StringPtrInput
}

func (RoleAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*roleAssignmentState)(nil)).Elem()
}

type roleAssignmentArgs struct {
	// The delegation flag used for creating a role assignment
	CanDelegate *bool `pulumi:"canDelegate"`
	// The principal ID assigned to the role. This maps to the ID inside the Active Directory. It can point to a user, service principal, or security group.
	PrincipalId *string `pulumi:"principalId"`
	// The name of the role assignment to create. It can be any valid GUID.
	RoleAssignmentName string `pulumi:"roleAssignmentName"`
	// The role definition ID used in the role assignment.
	RoleDefinitionId *string `pulumi:"roleDefinitionId"`
	// The scope of the role assignment to create. The scope can be any REST resource instance. For example, use '/subscriptions/{subscription-id}/' for a subscription, '/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}' for a resource group, and '/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}/providers/{resource-provider}/{resource-type}/{resource-name}' for a resource.
	Scope string `pulumi:"scope"`
}

// The set of arguments for constructing a RoleAssignment resource.
type RoleAssignmentArgs struct {
	// The delegation flag used for creating a role assignment
	CanDelegate pulumi.BoolPtrInput
	// The principal ID assigned to the role. This maps to the ID inside the Active Directory. It can point to a user, service principal, or security group.
	PrincipalId pulumi.StringPtrInput
	// The name of the role assignment to create. It can be any valid GUID.
	RoleAssignmentName pulumi.StringInput
	// The role definition ID used in the role assignment.
	RoleDefinitionId pulumi.StringPtrInput
	// The scope of the role assignment to create. The scope can be any REST resource instance. For example, use '/subscriptions/{subscription-id}/' for a subscription, '/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}' for a resource group, and '/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}/providers/{resource-provider}/{resource-type}/{resource-name}' for a resource.
	Scope pulumi.StringInput
}

func (RoleAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*roleAssignmentArgs)(nil)).Elem()
}
