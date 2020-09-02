// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180101preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupRoleAssignment(ctx *pulumi.Context, args *LookupRoleAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupRoleAssignmentResult, error) {
	var rv LookupRoleAssignmentResult
	err := ctx.Invoke("azurerm:authorization/v20180101preview:getRoleAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRoleAssignmentArgs struct {
	// The name of the role assignment to get.
	RoleAssignmentName string `pulumi:"roleAssignmentName"`
	// The scope of the role assignment.
	Scope string `pulumi:"scope"`
}

// Role Assignments
type LookupRoleAssignmentResult struct {
	// The Delegation flag for the role assignment
	CanDelegate *bool `pulumi:"canDelegate"`
	// The role assignment name.
	Name string `pulumi:"name"`
	// The principal ID.
	PrincipalId *string `pulumi:"principalId"`
	// The role definition ID.
	RoleDefinitionId *string `pulumi:"roleDefinitionId"`
	// The role assignment scope.
	Scope *string `pulumi:"scope"`
	// The role assignment type.
	Type string `pulumi:"type"`
}
