// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupWorkspaceCollection(ctx *pulumi.Context, args *LookupWorkspaceCollectionArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceCollectionResult, error) {
	var rv LookupWorkspaceCollectionResult
	err := ctx.Invoke("azurerm:powerbi/latest:getWorkspaceCollection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkspaceCollectionArgs struct {
	// Azure resource group
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Power BI Embedded Workspace Collection name
	WorkspaceCollectionName string `pulumi:"workspaceCollectionName"`
}

type LookupWorkspaceCollectionResult struct {
	// Azure location
	Location *string `pulumi:"location"`
	// Workspace collection name
	Name *string `pulumi:"name"`
	// Properties
	Properties map[string]interface{} `pulumi:"properties"`
	Sku        *AzureSkuResponse      `pulumi:"sku"`
	Tags       map[string]string      `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}