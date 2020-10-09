// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.StorageSync.V20200901
{
    public static class GetServerEndpoint
    {
        public static Task<GetServerEndpointResult> InvokeAsync(GetServerEndpointArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetServerEndpointResult>("azure-nextgen:storagesync/v20200901:getServerEndpoint", args ?? new GetServerEndpointArgs(), options.WithVersion());
    }


    public sealed class GetServerEndpointArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of Server Endpoint object.
        /// </summary>
        [Input("serverEndpointName", required: true)]
        public string ServerEndpointName { get; set; } = null!;

        /// <summary>
        /// Name of Storage Sync Service resource.
        /// </summary>
        [Input("storageSyncServiceName", required: true)]
        public string StorageSyncServiceName { get; set; } = null!;

        /// <summary>
        /// Name of Sync Group resource.
        /// </summary>
        [Input("syncGroupName", required: true)]
        public string SyncGroupName { get; set; } = null!;

        public GetServerEndpointArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetServerEndpointResult
    {
        /// <summary>
        /// Cloud Tiering.
        /// </summary>
        public readonly string? CloudTiering;
        /// <summary>
        /// Cloud tiering status. Only populated if cloud tiering is enabled.
        /// </summary>
        public readonly Outputs.ServerEndpointCloudTieringStatusResponse CloudTieringStatus;
        /// <summary>
        /// Friendly Name
        /// </summary>
        public readonly string? FriendlyName;
        /// <summary>
        /// Policy for how namespace and files are recalled during FastDr.
        /// </summary>
        public readonly string? InitialDownloadPolicy;
        /// <summary>
        /// Resource Last Operation Name
        /// </summary>
        public readonly string LastOperationName;
        /// <summary>
        /// ServerEndpoint lastWorkflowId
        /// </summary>
        public readonly string LastWorkflowId;
        /// <summary>
        /// Policy for enabling follow-the-sun business models: link local cache to cloud behavior to pre-populate before local access.
        /// </summary>
        public readonly string? LocalCacheMode;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Offline data transfer
        /// </summary>
        public readonly string? OfflineDataTransfer;
        /// <summary>
        /// Offline data transfer share name
        /// </summary>
        public readonly string? OfflineDataTransferShareName;
        /// <summary>
        /// Offline data transfer storage account resource ID
        /// </summary>
        public readonly string OfflineDataTransferStorageAccountResourceId;
        /// <summary>
        /// Offline data transfer storage account tenant ID
        /// </summary>
        public readonly string OfflineDataTransferStorageAccountTenantId;
        /// <summary>
        /// ServerEndpoint Provisioning State
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Recall status. Only populated if cloud tiering is enabled.
        /// </summary>
        public readonly Outputs.ServerEndpointRecallStatusResponse RecallStatus;
        /// <summary>
        /// Server Local path.
        /// </summary>
        public readonly string? ServerLocalPath;
        /// <summary>
        /// Server name
        /// </summary>
        public readonly string ServerName;
        /// <summary>
        /// Server Resource Id.
        /// </summary>
        public readonly string? ServerResourceId;
        /// <summary>
        /// Server Endpoint sync status
        /// </summary>
        public readonly Outputs.ServerEndpointSyncStatusResponse SyncStatus;
        /// <summary>
        /// Tier files older than days.
        /// </summary>
        public readonly int? TierFilesOlderThanDays;
        /// <summary>
        /// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Level of free space to be maintained by Cloud Tiering if it is enabled.
        /// </summary>
        public readonly int? VolumeFreeSpacePercent;

        [OutputConstructor]
        private GetServerEndpointResult(
            string? cloudTiering,

            Outputs.ServerEndpointCloudTieringStatusResponse cloudTieringStatus,

            string? friendlyName,

            string? initialDownloadPolicy,

            string lastOperationName,

            string lastWorkflowId,

            string? localCacheMode,

            string name,

            string? offlineDataTransfer,

            string? offlineDataTransferShareName,

            string offlineDataTransferStorageAccountResourceId,

            string offlineDataTransferStorageAccountTenantId,

            string provisioningState,

            Outputs.ServerEndpointRecallStatusResponse recallStatus,

            string? serverLocalPath,

            string serverName,

            string? serverResourceId,

            Outputs.ServerEndpointSyncStatusResponse syncStatus,

            int? tierFilesOlderThanDays,

            string type,

            int? volumeFreeSpacePercent)
        {
            CloudTiering = cloudTiering;
            CloudTieringStatus = cloudTieringStatus;
            FriendlyName = friendlyName;
            InitialDownloadPolicy = initialDownloadPolicy;
            LastOperationName = lastOperationName;
            LastWorkflowId = lastWorkflowId;
            LocalCacheMode = localCacheMode;
            Name = name;
            OfflineDataTransfer = offlineDataTransfer;
            OfflineDataTransferShareName = offlineDataTransferShareName;
            OfflineDataTransferStorageAccountResourceId = offlineDataTransferStorageAccountResourceId;
            OfflineDataTransferStorageAccountTenantId = offlineDataTransferStorageAccountTenantId;
            ProvisioningState = provisioningState;
            RecallStatus = recallStatus;
            ServerLocalPath = serverLocalPath;
            ServerName = serverName;
            ServerResourceId = serverResourceId;
            SyncStatus = syncStatus;
            TierFilesOlderThanDays = tierFilesOlderThanDays;
            Type = type;
            VolumeFreeSpacePercent = volumeFreeSpacePercent;
        }
    }
}
