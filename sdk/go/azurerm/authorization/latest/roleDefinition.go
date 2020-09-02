// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Role definition.
type RoleDefinition struct {
	pulumi.CustomResourceState

	// Role definition assignable scopes.
	AssignableScopes pulumi.StringArrayOutput `pulumi:"assignableScopes"`
	// The role definition description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The role definition name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Role definition permissions.
	Permissions PermissionResponseArrayOutput `pulumi:"permissions"`
	// The role name.
	RoleName pulumi.StringPtrOutput `pulumi:"roleName"`
	// The role type.
	RoleType pulumi.StringPtrOutput `pulumi:"roleType"`
	// The role definition type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRoleDefinition registers a new resource with the given unique name, arguments, and options.
func NewRoleDefinition(ctx *pulumi.Context,
	name string, args *RoleDefinitionArgs, opts ...pulumi.ResourceOption) (*RoleDefinition, error) {
	if args == nil || args.RoleDefinitionId == nil {
		return nil, errors.New("missing required argument 'RoleDefinitionId'")
	}
	if args == nil || args.Scope == nil {
		return nil, errors.New("missing required argument 'Scope'")
	}
	if args == nil {
		args = &RoleDefinitionArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:authorization/v20150701:RoleDefinition"),
		},
		{
			Type: pulumi.String("azurerm:authorization/v20180101preview:RoleDefinition"),
		},
	})
	opts = append(opts, aliases)
	var resource RoleDefinition
	err := ctx.RegisterResource("azurerm:authorization/latest:RoleDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoleDefinition gets an existing RoleDefinition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoleDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoleDefinitionState, opts ...pulumi.ResourceOption) (*RoleDefinition, error) {
	var resource RoleDefinition
	err := ctx.ReadResource("azurerm:authorization/latest:RoleDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RoleDefinition resources.
type roleDefinitionState struct {
	// Role definition assignable scopes.
	AssignableScopes []string `pulumi:"assignableScopes"`
	// The role definition description.
	Description *string `pulumi:"description"`
	// The role definition name.
	Name *string `pulumi:"name"`
	// Role definition permissions.
	Permissions []PermissionResponse `pulumi:"permissions"`
	// The role name.
	RoleName *string `pulumi:"roleName"`
	// The role type.
	RoleType *string `pulumi:"roleType"`
	// The role definition type.
	Type *string `pulumi:"type"`
}

type RoleDefinitionState struct {
	// Role definition assignable scopes.
	AssignableScopes pulumi.StringArrayInput
	// The role definition description.
	Description pulumi.StringPtrInput
	// The role definition name.
	Name pulumi.StringPtrInput
	// Role definition permissions.
	Permissions PermissionResponseArrayInput
	// The role name.
	RoleName pulumi.StringPtrInput
	// The role type.
	RoleType pulumi.StringPtrInput
	// The role definition type.
	Type pulumi.StringPtrInput
}

func (RoleDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*roleDefinitionState)(nil)).Elem()
}

type roleDefinitionArgs struct {
	// Role definition assignable scopes.
	AssignableScopes []string `pulumi:"assignableScopes"`
	// The role definition description.
	Description *string `pulumi:"description"`
	// Role definition permissions.
	Permissions []Permission `pulumi:"permissions"`
	// The ID of the role definition.
	RoleDefinitionId string `pulumi:"roleDefinitionId"`
	// The role name.
	RoleName *string `pulumi:"roleName"`
	// The role type.
	RoleType *string `pulumi:"roleType"`
	// The scope of the role definition.
	Scope string `pulumi:"scope"`
}

// The set of arguments for constructing a RoleDefinition resource.
type RoleDefinitionArgs struct {
	// Role definition assignable scopes.
	AssignableScopes pulumi.StringArrayInput
	// The role definition description.
	Description pulumi.StringPtrInput
	// Role definition permissions.
	Permissions PermissionArrayInput
	// The ID of the role definition.
	RoleDefinitionId pulumi.StringInput
	// The role name.
	RoleName pulumi.StringPtrInput
	// The role type.
	RoleType pulumi.StringPtrInput
	// The scope of the role definition.
	Scope pulumi.StringInput
}

func (RoleDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*roleDefinitionArgs)(nil)).Elem()
}
