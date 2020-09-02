// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191001preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The role assignment
type BillingRoleAssignmentByBillingAccount struct {
	pulumi.CustomResourceState

	// The principal Id of the user who created the role assignment.
	CreatedByPrincipalId pulumi.StringOutput `pulumi:"createdByPrincipalId"`
	// The tenant Id of the user who created the role assignment.
	CreatedByPrincipalTenantId pulumi.StringOutput `pulumi:"createdByPrincipalTenantId"`
	// The email address of the user who created the role assignment. This is supported only for billing accounts with agreement type Enterprise Agreement.
	CreatedByUserEmailAddress pulumi.StringOutput `pulumi:"createdByUserEmailAddress"`
	// The date the role assignment was created.
	CreatedOn pulumi.StringOutput `pulumi:"createdOn"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The principal id of the user to whom the role was assigned.
	PrincipalId pulumi.StringPtrOutput `pulumi:"principalId"`
	// The principal tenant id of the user to whom the role was assigned.
	PrincipalTenantId pulumi.StringPtrOutput `pulumi:"principalTenantId"`
	// The ID of the role definition.
	RoleDefinitionId pulumi.StringPtrOutput `pulumi:"roleDefinitionId"`
	// The scope at which the role was assigned.
	Scope pulumi.StringOutput `pulumi:"scope"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The authentication type of the user, whether Organization or MSA, of the user to whom the role was assigned. This is supported only for billing accounts with agreement type Enterprise Agreement.
	UserAuthenticationType pulumi.StringPtrOutput `pulumi:"userAuthenticationType"`
	// The email address of the user to whom the role was assigned. This is supported only for billing accounts with agreement type Enterprise Agreement.
	UserEmailAddress pulumi.StringPtrOutput `pulumi:"userEmailAddress"`
}

// NewBillingRoleAssignmentByBillingAccount registers a new resource with the given unique name, arguments, and options.
func NewBillingRoleAssignmentByBillingAccount(ctx *pulumi.Context,
	name string, args *BillingRoleAssignmentByBillingAccountArgs, opts ...pulumi.ResourceOption) (*BillingRoleAssignmentByBillingAccount, error) {
	if args == nil || args.BillingAccountName == nil {
		return nil, errors.New("missing required argument 'BillingAccountName'")
	}
	if args == nil || args.BillingRoleAssignmentName == nil {
		return nil, errors.New("missing required argument 'BillingRoleAssignmentName'")
	}
	if args == nil {
		args = &BillingRoleAssignmentByBillingAccountArgs{}
	}
	var resource BillingRoleAssignmentByBillingAccount
	err := ctx.RegisterResource("azurerm:billing/v20191001preview:BillingRoleAssignmentByBillingAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBillingRoleAssignmentByBillingAccount gets an existing BillingRoleAssignmentByBillingAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBillingRoleAssignmentByBillingAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BillingRoleAssignmentByBillingAccountState, opts ...pulumi.ResourceOption) (*BillingRoleAssignmentByBillingAccount, error) {
	var resource BillingRoleAssignmentByBillingAccount
	err := ctx.ReadResource("azurerm:billing/v20191001preview:BillingRoleAssignmentByBillingAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BillingRoleAssignmentByBillingAccount resources.
type billingRoleAssignmentByBillingAccountState struct {
	// The principal Id of the user who created the role assignment.
	CreatedByPrincipalId *string `pulumi:"createdByPrincipalId"`
	// The tenant Id of the user who created the role assignment.
	CreatedByPrincipalTenantId *string `pulumi:"createdByPrincipalTenantId"`
	// The email address of the user who created the role assignment. This is supported only for billing accounts with agreement type Enterprise Agreement.
	CreatedByUserEmailAddress *string `pulumi:"createdByUserEmailAddress"`
	// The date the role assignment was created.
	CreatedOn *string `pulumi:"createdOn"`
	// Resource name.
	Name *string `pulumi:"name"`
	// The principal id of the user to whom the role was assigned.
	PrincipalId *string `pulumi:"principalId"`
	// The principal tenant id of the user to whom the role was assigned.
	PrincipalTenantId *string `pulumi:"principalTenantId"`
	// The ID of the role definition.
	RoleDefinitionId *string `pulumi:"roleDefinitionId"`
	// The scope at which the role was assigned.
	Scope *string `pulumi:"scope"`
	// Resource type.
	Type *string `pulumi:"type"`
	// The authentication type of the user, whether Organization or MSA, of the user to whom the role was assigned. This is supported only for billing accounts with agreement type Enterprise Agreement.
	UserAuthenticationType *string `pulumi:"userAuthenticationType"`
	// The email address of the user to whom the role was assigned. This is supported only for billing accounts with agreement type Enterprise Agreement.
	UserEmailAddress *string `pulumi:"userEmailAddress"`
}

type BillingRoleAssignmentByBillingAccountState struct {
	// The principal Id of the user who created the role assignment.
	CreatedByPrincipalId pulumi.StringPtrInput
	// The tenant Id of the user who created the role assignment.
	CreatedByPrincipalTenantId pulumi.StringPtrInput
	// The email address of the user who created the role assignment. This is supported only for billing accounts with agreement type Enterprise Agreement.
	CreatedByUserEmailAddress pulumi.StringPtrInput
	// The date the role assignment was created.
	CreatedOn pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// The principal id of the user to whom the role was assigned.
	PrincipalId pulumi.StringPtrInput
	// The principal tenant id of the user to whom the role was assigned.
	PrincipalTenantId pulumi.StringPtrInput
	// The ID of the role definition.
	RoleDefinitionId pulumi.StringPtrInput
	// The scope at which the role was assigned.
	Scope pulumi.StringPtrInput
	// Resource type.
	Type pulumi.StringPtrInput
	// The authentication type of the user, whether Organization or MSA, of the user to whom the role was assigned. This is supported only for billing accounts with agreement type Enterprise Agreement.
	UserAuthenticationType pulumi.StringPtrInput
	// The email address of the user to whom the role was assigned. This is supported only for billing accounts with agreement type Enterprise Agreement.
	UserEmailAddress pulumi.StringPtrInput
}

func (BillingRoleAssignmentByBillingAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*billingRoleAssignmentByBillingAccountState)(nil)).Elem()
}

type billingRoleAssignmentByBillingAccountArgs struct {
	// The ID that uniquely identifies a billing account.
	BillingAccountName string `pulumi:"billingAccountName"`
	// The ID that uniquely identifies a role assignment.
	BillingRoleAssignmentName string `pulumi:"billingRoleAssignmentName"`
	// The principal id of the user to whom the role was assigned.
	PrincipalId *string `pulumi:"principalId"`
	// The principal tenant id of the user to whom the role was assigned.
	PrincipalTenantId *string `pulumi:"principalTenantId"`
	// The ID of the role definition.
	RoleDefinitionId *string `pulumi:"roleDefinitionId"`
	// The authentication type of the user, whether Organization or MSA, of the user to whom the role was assigned. This is supported only for billing accounts with agreement type Enterprise Agreement.
	UserAuthenticationType *string `pulumi:"userAuthenticationType"`
	// The email address of the user to whom the role was assigned. This is supported only for billing accounts with agreement type Enterprise Agreement.
	UserEmailAddress *string `pulumi:"userEmailAddress"`
}

// The set of arguments for constructing a BillingRoleAssignmentByBillingAccount resource.
type BillingRoleAssignmentByBillingAccountArgs struct {
	// The ID that uniquely identifies a billing account.
	BillingAccountName pulumi.StringInput
	// The ID that uniquely identifies a role assignment.
	BillingRoleAssignmentName pulumi.StringInput
	// The principal id of the user to whom the role was assigned.
	PrincipalId pulumi.StringPtrInput
	// The principal tenant id of the user to whom the role was assigned.
	PrincipalTenantId pulumi.StringPtrInput
	// The ID of the role definition.
	RoleDefinitionId pulumi.StringPtrInput
	// The authentication type of the user, whether Organization or MSA, of the user to whom the role was assigned. This is supported only for billing accounts with agreement type Enterprise Agreement.
	UserAuthenticationType pulumi.StringPtrInput
	// The email address of the user to whom the role was assigned. This is supported only for billing accounts with agreement type Enterprise Agreement.
	UserEmailAddress pulumi.StringPtrInput
}

func (BillingRoleAssignmentByBillingAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*billingRoleAssignmentByBillingAccountArgs)(nil)).Elem()
}
