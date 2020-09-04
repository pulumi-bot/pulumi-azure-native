// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200301preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupWorkspace(ctx *pulumi.Context, args *LookupWorkspaceArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceResult, error) {
	var rv LookupWorkspaceResult
	err := ctx.Invoke("azurerm:operationalinsights/v20200301preview:getWorkspace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkspaceArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The top level Workspace resource container.
type LookupWorkspaceResult struct {
	// This is a read-only property. Represents the ID associated with the workspace.
	CustomerId string `pulumi:"customerId"`
	// The ETag of the workspace.
	ETag *string `pulumi:"eTag"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// List of linked private link scope resources.
	PrivateLinkScopedResources []PrivateLinkScopedResourceResponse `pulumi:"privateLinkScopedResources"`
	// The provisioning state of the workspace.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The network access type for accessing Log Analytics ingestion.
	PublicNetworkAccessForIngestion *string `pulumi:"publicNetworkAccessForIngestion"`
	// The network access type for accessing Log Analytics query.
	PublicNetworkAccessForQuery *string `pulumi:"publicNetworkAccessForQuery"`
	// The workspace data retention in days. -1 means Unlimited retention for the Unlimited Sku. 730 days is the maximum allowed for all other Skus.
	RetentionInDays *int `pulumi:"retentionInDays"`
	// The SKU of the workspace.
	Sku *WorkspaceSkuResponse `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type string `pulumi:"type"`
	// The daily volume cap for ingestion.
	WorkspaceCapping *WorkspaceCappingResponse `pulumi:"workspaceCapping"`
}
