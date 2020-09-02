// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.StorageSync.V20190301
{
    /// <summary>
    /// Cloud Endpoint object.
    /// </summary>
    public partial class CloudEndpoint : Pulumi.CustomResource
    {
        /// <summary>
        /// Azure file share name
        /// </summary>
        [Output("azureFileShareName")]
        public Output<string?> AzureFileShareName { get; private set; } = null!;

        /// <summary>
        /// Backup Enabled
        /// </summary>
        [Output("backupEnabled")]
        public Output<string> BackupEnabled { get; private set; } = null!;

        /// <summary>
        /// Friendly Name
        /// </summary>
        [Output("friendlyName")]
        public Output<string?> FriendlyName { get; private set; } = null!;

        /// <summary>
        /// Resource Last Operation Name
        /// </summary>
        [Output("lastOperationName")]
        public Output<string?> LastOperationName { get; private set; } = null!;

        /// <summary>
        /// CloudEndpoint lastWorkflowId
        /// </summary>
        [Output("lastWorkflowId")]
        public Output<string?> LastWorkflowId { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Partnership Id
        /// </summary>
        [Output("partnershipId")]
        public Output<string?> PartnershipId { get; private set; } = null!;

        /// <summary>
        /// CloudEndpoint Provisioning State
        /// </summary>
        [Output("provisioningState")]
        public Output<string?> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Storage Account Resource Id
        /// </summary>
        [Output("storageAccountResourceId")]
        public Output<string?> StorageAccountResourceId { get; private set; } = null!;

        /// <summary>
        /// Storage Account Tenant Id
        /// </summary>
        [Output("storageAccountTenantId")]
        public Output<string?> StorageAccountTenantId { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a CloudEndpoint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CloudEndpoint(string name, CloudEndpointArgs args, CustomResourceOptions? options = null)
            : base("azurerm:storagesync/v20190301:CloudEndpoint", name, args ?? new CloudEndpointArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CloudEndpoint(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:storagesync/v20190301:CloudEndpoint", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:storagesync/latest:CloudEndpoint"},
                    new Pulumi.Alias { Type = "azurerm:storagesync/v20170605preview:CloudEndpoint"},
                    new Pulumi.Alias { Type = "azurerm:storagesync/v20180402:CloudEndpoint"},
                    new Pulumi.Alias { Type = "azurerm:storagesync/v20180701:CloudEndpoint"},
                    new Pulumi.Alias { Type = "azurerm:storagesync/v20181001:CloudEndpoint"},
                    new Pulumi.Alias { Type = "azurerm:storagesync/v20190201:CloudEndpoint"},
                    new Pulumi.Alias { Type = "azurerm:storagesync/v20190601:CloudEndpoint"},
                    new Pulumi.Alias { Type = "azurerm:storagesync/v20191001:CloudEndpoint"},
                    new Pulumi.Alias { Type = "azurerm:storagesync/v20200301:CloudEndpoint"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CloudEndpoint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CloudEndpoint Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new CloudEndpoint(name, id, options);
        }
    }

    public sealed class CloudEndpointArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Azure file share name
        /// </summary>
        [Input("azureFileShareName")]
        public Input<string>? AzureFileShareName { get; set; }

        /// <summary>
        /// Name of Cloud Endpoint object.
        /// </summary>
        [Input("cloudEndpointName", required: true)]
        public Input<string> CloudEndpointName { get; set; } = null!;

        /// <summary>
        /// Friendly Name
        /// </summary>
        [Input("friendlyName")]
        public Input<string>? FriendlyName { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Storage Account Resource Id
        /// </summary>
        [Input("storageAccountResourceId")]
        public Input<string>? StorageAccountResourceId { get; set; }

        /// <summary>
        /// Storage Account Tenant Id
        /// </summary>
        [Input("storageAccountTenantId")]
        public Input<string>? StorageAccountTenantId { get; set; }

        /// <summary>
        /// Name of Storage Sync Service resource.
        /// </summary>
        [Input("storageSyncServiceName", required: true)]
        public Input<string> StorageSyncServiceName { get; set; } = null!;

        /// <summary>
        /// Name of Sync Group resource.
        /// </summary>
        [Input("syncGroupName", required: true)]
        public Input<string> SyncGroupName { get; set; } = null!;

        public CloudEndpointArgs()
        {
        }
    }
}
