// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200301preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupStorageInsightConfig(ctx *pulumi.Context, args *LookupStorageInsightConfigArgs, opts ...pulumi.InvokeOption) (*LookupStorageInsightConfigResult, error) {
	var rv LookupStorageInsightConfigResult
	err := ctx.Invoke("azure-nextgen:operationalinsights/v20200301preview:getStorageInsightConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStorageInsightConfigArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the storageInsightsConfigs resource
	StorageInsightName string `pulumi:"storageInsightName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The top level storage insight resource container.
type LookupStorageInsightConfigResult struct {
	// The names of the blob containers that the workspace should read
	Containers []string `pulumi:"containers"`
	// The ETag of the storage insight.
	ETag *string `pulumi:"eTag"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The status of the storage insight
	Status StorageInsightStatusResponse `pulumi:"status"`
	// The storage account connection details
	StorageAccount StorageAccountResponse `pulumi:"storageAccount"`
	// The names of the Azure tables that the workspace should read
	Tables []string `pulumi:"tables"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}
