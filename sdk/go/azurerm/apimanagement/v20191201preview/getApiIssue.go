// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191201preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupApiIssue(ctx *pulumi.Context, args *LookupApiIssueArgs, opts ...pulumi.InvokeOption) (*LookupApiIssueResult, error) {
	var rv LookupApiIssueResult
	err := ctx.Invoke("azurerm:apimanagement/v20191201preview:getApiIssue", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApiIssueArgs struct {
	// API identifier. Must be unique in the current API Management service instance.
	ApiId string `pulumi:"apiId"`
	// Expand the comment attachments.
	ExpandCommentsAttachments *bool `pulumi:"expandCommentsAttachments"`
	// Issue identifier. Must be unique in the current API Management service instance.
	IssueId string `pulumi:"issueId"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// Issue Contract details.
type LookupApiIssueResult struct {
	// A resource identifier for the API the issue was created for.
	ApiId *string `pulumi:"apiId"`
	// Date and time when the issue was created.
	CreatedDate *string `pulumi:"createdDate"`
	// Text describing the issue.
	Description string `pulumi:"description"`
	// Resource name.
	Name string `pulumi:"name"`
	// Status of the issue.
	State *string `pulumi:"state"`
	// The issue title.
	Title string `pulumi:"title"`
	// Resource type for API Management resource.
	Type string `pulumi:"type"`
	// A resource identifier for the user created the issue.
	UserId string `pulumi:"userId"`
}
