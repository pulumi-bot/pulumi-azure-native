// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

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
// 	authorization "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/authorization/latest"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := authorization.NewRoleAssignment(ctx, "roleAssignment", &authorization.RoleAssignmentArgs{
// 			RoleAssignmentName: pulumi.String("roleAssignmentName"),
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

	// The role assignment name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Role assignment properties.
	Properties RoleAssignmentPropertiesWithScopeResponseOutput `pulumi:"properties"`
	// The role assignment type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRoleAssignment registers a new resource with the given unique name, arguments, and options.
func NewRoleAssignment(ctx *pulumi.Context,
	name string, args *RoleAssignmentArgs, opts ...pulumi.ResourceOption) (*RoleAssignment, error) {
	if args == nil || args.Properties == nil {
		return nil, errors.New("missing required argument 'Properties'")
	}
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
			Type: pulumi.String("azurerm:authorization/v20150701:RoleAssignment"),
		},
		{
			Type: pulumi.String("azurerm:authorization/v20171001preview:RoleAssignment"),
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
	err := ctx.RegisterResource("azurerm:authorization/latest:RoleAssignment", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azurerm:authorization/latest:RoleAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RoleAssignment resources.
type roleAssignmentState struct {
	// The role assignment name.
	Name *string `pulumi:"name"`
	// Role assignment properties.
	Properties *RoleAssignmentPropertiesWithScopeResponse `pulumi:"properties"`
	// The role assignment type.
	Type *string `pulumi:"type"`
}

type RoleAssignmentState struct {
	// The role assignment name.
	Name pulumi.StringPtrInput
	// Role assignment properties.
	Properties RoleAssignmentPropertiesWithScopeResponsePtrInput
	// The role assignment type.
	Type pulumi.StringPtrInput
}

func (RoleAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*roleAssignmentState)(nil)).Elem()
}

type roleAssignmentArgs struct {
	// Role assignment properties.
	Properties RoleAssignmentProperties `pulumi:"properties"`
	// The name of the role assignment to create. It can be any valid GUID.
	RoleAssignmentName string `pulumi:"roleAssignmentName"`
	// The scope of the role assignment to create. The scope can be any REST resource instance. For example, use '/subscriptions/{subscription-id}/' for a subscription, '/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}' for a resource group, and '/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}/providers/{resource-provider}/{resource-type}/{resource-name}' for a resource.
	Scope string `pulumi:"scope"`
}

// The set of arguments for constructing a RoleAssignment resource.
type RoleAssignmentArgs struct {
	// Role assignment properties.
	Properties RoleAssignmentPropertiesInput
	// The name of the role assignment to create. It can be any valid GUID.
	RoleAssignmentName pulumi.StringInput
	// The scope of the role assignment to create. The scope can be any REST resource instance. For example, use '/subscriptions/{subscription-id}/' for a subscription, '/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}' for a resource group, and '/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}/providers/{resource-provider}/{resource-type}/{resource-name}' for a resource.
	Scope pulumi.StringInput
}

func (RoleAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*roleAssignmentArgs)(nil)).Elem()
}
