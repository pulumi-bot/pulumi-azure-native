// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20171101preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupManagementGroup(ctx *pulumi.Context, args *LookupManagementGroupArgs, opts ...pulumi.InvokeOption) (*LookupManagementGroupResult, error) {
	var rv LookupManagementGroupResult
	err := ctx.Invoke("azurerm:management/v20171101preview:getManagementGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagementGroupArgs struct {
	// The $expand=children query string parameter allows clients to request inclusion of children in the response payload.
	Expand *string `pulumi:"expand"`
	// Management Group ID.
	GroupId string `pulumi:"groupId"`
	// The $recurse=true query string parameter allows clients to request inclusion of entire hierarchy in the response payload.
	Recurse *bool `pulumi:"recurse"`
}

// The management group details.
type LookupManagementGroupResult struct {
	// The list of children.
	Children []ManagementGroupChildInfoResponse `pulumi:"children"`
	// The details of a management group.
	Details *ManagementGroupDetailsResponse `pulumi:"details"`
	// The friendly name of the management group.
	DisplayName *string `pulumi:"displayName"`
	// The name of the management group. For example, 00000000-0000-0000-0000-000000000000
	Name string `pulumi:"name"`
	// The AAD Tenant ID associated with the management group. For example, 00000000-0000-0000-0000-000000000000
	TenantId *string `pulumi:"tenantId"`
	// The type of the resource.  For example, /providers/Microsoft.Management/managementGroups
	Type string `pulumi:"type"`
}
