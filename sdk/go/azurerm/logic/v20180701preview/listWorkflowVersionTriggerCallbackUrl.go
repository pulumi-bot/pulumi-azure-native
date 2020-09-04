// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180701preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func ListWorkflowVersionTriggerCallbackUrl(ctx *pulumi.Context, args *ListWorkflowVersionTriggerCallbackUrlArgs, opts ...pulumi.InvokeOption) (*ListWorkflowVersionTriggerCallbackUrlResult, error) {
	var rv ListWorkflowVersionTriggerCallbackUrlResult
	err := ctx.Invoke("azurerm:logic/v20180701preview:listWorkflowVersionTriggerCallbackUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWorkflowVersionTriggerCallbackUrlArgs struct {
	// The key type.
	KeyType *string `pulumi:"keyType"`
	// The expiry time.
	NotAfter *string `pulumi:"notAfter"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The workflow trigger name.
	TriggerName string `pulumi:"triggerName"`
	// The workflow versionId.
	VersionId string `pulumi:"versionId"`
	// The workflow name.
	WorkflowName string `pulumi:"workflowName"`
}

// The workflow trigger callback URL.
type ListWorkflowVersionTriggerCallbackUrlResult struct {
	// Gets the workflow trigger callback URL base path.
	BasePath string `pulumi:"basePath"`
	// Gets the workflow trigger callback URL HTTP method.
	Method string `pulumi:"method"`
	// Gets the workflow trigger callback URL query parameters.
	Queries *WorkflowTriggerListCallbackUrlQueriesResponse `pulumi:"queries"`
	// Gets the workflow trigger callback URL relative path.
	RelativePath string `pulumi:"relativePath"`
	// Gets the workflow trigger callback URL relative path parameters.
	RelativePathParameters []string `pulumi:"relativePathParameters"`
	// Gets the workflow trigger callback URL.
	Value string `pulumi:"value"`
}
