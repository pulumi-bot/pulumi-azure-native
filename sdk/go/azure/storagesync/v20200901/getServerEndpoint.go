// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200901

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupServerEndpoint(ctx *pulumi.Context, args *LookupServerEndpointArgs, opts ...pulumi.InvokeOption) (*LookupServerEndpointResult, error) {
	var rv LookupServerEndpointResult
	err := ctx.Invoke("azure-nextgen:storagesync/v20200901:getServerEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerEndpointArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Server Endpoint object.
	ServerEndpointName string `pulumi:"serverEndpointName"`
	// Name of Storage Sync Service resource.
	StorageSyncServiceName string `pulumi:"storageSyncServiceName"`
	// Name of Sync Group resource.
	SyncGroupName string `pulumi:"syncGroupName"`
}

// Server Endpoint object.
type LookupServerEndpointResult struct {
	// Cloud Tiering.
	CloudTiering *string `pulumi:"cloudTiering"`
	// Cloud tiering status. Only populated if cloud tiering is enabled.
	CloudTieringStatus ServerEndpointCloudTieringStatusResponse `pulumi:"cloudTieringStatus"`
	// Friendly Name
	FriendlyName *string `pulumi:"friendlyName"`
	// Policy for how namespace and files are recalled during FastDr.
	InitialDownloadPolicy *string `pulumi:"initialDownloadPolicy"`
	// Resource Last Operation Name
	LastOperationName string `pulumi:"lastOperationName"`
	// ServerEndpoint lastWorkflowId
	LastWorkflowId string `pulumi:"lastWorkflowId"`
	// Policy for enabling follow-the-sun business models: link local cache to cloud behavior to pre-populate before local access.
	LocalCacheMode *string `pulumi:"localCacheMode"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Offline data transfer
	OfflineDataTransfer *string `pulumi:"offlineDataTransfer"`
	// Offline data transfer share name
	OfflineDataTransferShareName *string `pulumi:"offlineDataTransferShareName"`
	// Offline data transfer storage account resource ID
	OfflineDataTransferStorageAccountResourceId string `pulumi:"offlineDataTransferStorageAccountResourceId"`
	// Offline data transfer storage account tenant ID
	OfflineDataTransferStorageAccountTenantId string `pulumi:"offlineDataTransferStorageAccountTenantId"`
	// ServerEndpoint Provisioning State
	ProvisioningState string `pulumi:"provisioningState"`
	// Recall status. Only populated if cloud tiering is enabled.
	RecallStatus ServerEndpointRecallStatusResponse `pulumi:"recallStatus"`
	// Server Local path.
	ServerLocalPath *string `pulumi:"serverLocalPath"`
	// Server Resource Id.
	ServerResourceId *string `pulumi:"serverResourceId"`
	// Server Endpoint sync status
	SyncStatus ServerEndpointSyncStatusResponse `pulumi:"syncStatus"`
	// Tier files older than days.
	TierFilesOlderThanDays *int `pulumi:"tierFilesOlderThanDays"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type string `pulumi:"type"`
	// Level of free space to be maintained by Cloud Tiering if it is enabled.
	VolumeFreeSpacePercent *int `pulumi:"volumeFreeSpacePercent"`
}
