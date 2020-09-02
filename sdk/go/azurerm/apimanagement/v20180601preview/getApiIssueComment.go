// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180601preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupApiIssueComment(ctx *pulumi.Context, args *LookupApiIssueCommentArgs, opts ...pulumi.InvokeOption) (*LookupApiIssueCommentResult, error) {
	var rv LookupApiIssueCommentResult
	err := ctx.Invoke("azurerm:apimanagement/v20180601preview:getApiIssueComment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApiIssueCommentArgs struct {
	// API identifier. Must be unique in the current API Management service instance.
	ApiId string `pulumi:"apiId"`
	// Comment identifier within an Issue. Must be unique in the current Issue.
	CommentId string `pulumi:"commentId"`
	// Issue identifier. Must be unique in the current API Management service instance.
	IssueId string `pulumi:"issueId"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// Issue Comment Contract details.
type LookupApiIssueCommentResult struct {
	// Date and time when the comment was created.
	CreatedDate *string `pulumi:"createdDate"`
	// Resource name.
	Name string `pulumi:"name"`
	// Comment text.
	Text string `pulumi:"text"`
	// Resource type for API Management resource.
	Type string `pulumi:"type"`
	// A resource identifier for the user who left the comment.
	UserId string `pulumi:"userId"`
}
