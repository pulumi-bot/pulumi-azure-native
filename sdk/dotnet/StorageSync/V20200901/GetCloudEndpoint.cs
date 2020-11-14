// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.StorageSync.V20200901
{
    public static class GetCloudEndpoint
    {
        public static Task<GetCloudEndpointResult> InvokeAsync(GetCloudEndpointArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCloudEndpointResult>("azure-nextgen:storagesync/v20200901:getCloudEndpoint", args ?? new GetCloudEndpointArgs(), options.WithVersion());
    }


    public sealed class GetCloudEndpointArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of Cloud Endpoint object.
        /// </summary>
        [Input("cloudEndpointName", required: true)]
        public string CloudEndpointName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

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

        public GetCloudEndpointArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetCloudEndpointResult
    {
        /// <summary>
        /// Azure file share name
        /// </summary>
        public readonly string? AzureFileShareName;
        /// <summary>
        /// Backup Enabled
        /// </summary>
        public readonly string BackupEnabled;
        /// <summary>
        /// Friendly Name
        /// </summary>
        public readonly string? FriendlyName;
        /// <summary>
        /// Resource Last Operation Name
        /// </summary>
        public readonly string? LastOperationName;
        /// <summary>
        /// CloudEndpoint lastWorkflowId
        /// </summary>
        public readonly string? LastWorkflowId;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Partnership Id
        /// </summary>
        public readonly string? PartnershipId;
        /// <summary>
        /// CloudEndpoint Provisioning State
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// Storage Account Resource Id
        /// </summary>
        public readonly string? StorageAccountResourceId;
        /// <summary>
        /// Storage Account Tenant Id
        /// </summary>
        public readonly string? StorageAccountTenantId;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetCloudEndpointResult(
            string? azureFileShareName,

            string backupEnabled,

            string? friendlyName,

            string? lastOperationName,

            string? lastWorkflowId,

            string name,

            string? partnershipId,

            string? provisioningState,

            string? storageAccountResourceId,

            string? storageAccountTenantId,

            string type)
        {
            AzureFileShareName = azureFileShareName;
            BackupEnabled = backupEnabled;
            FriendlyName = friendlyName;
            LastOperationName = lastOperationName;
            LastWorkflowId = lastWorkflowId;
            Name = name;
            PartnershipId = partnershipId;
            ProvisioningState = provisioningState;
            StorageAccountResourceId = storageAccountResourceId;
            StorageAccountTenantId = storageAccountTenantId;
            Type = type;
        }
    }
}
