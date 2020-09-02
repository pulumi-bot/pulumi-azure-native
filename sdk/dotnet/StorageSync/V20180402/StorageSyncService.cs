// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.StorageSync.V20180402
{
    /// <summary>
    /// Storage Sync Service object.
    /// </summary>
    public partial class StorageSyncService : Pulumi.CustomResource
    {
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Storage Sync service status.
        /// </summary>
        [Output("storageSyncServiceStatus")]
        public Output<int> StorageSyncServiceStatus { get; private set; } = null!;

        /// <summary>
        /// Storage Sync service Uid
        /// </summary>
        [Output("storageSyncServiceUid")]
        public Output<string> StorageSyncServiceUid { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a StorageSyncService resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StorageSyncService(string name, StorageSyncServiceArgs args, CustomResourceOptions? options = null)
            : base("azurerm:storagesync/v20180402:StorageSyncService", name, args ?? new StorageSyncServiceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StorageSyncService(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:storagesync/v20180402:StorageSyncService", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:storagesync/latest:StorageSyncService"},
                    new Pulumi.Alias { Type = "azurerm:storagesync/v20170605preview:StorageSyncService"},
                    new Pulumi.Alias { Type = "azurerm:storagesync/v20180701:StorageSyncService"},
                    new Pulumi.Alias { Type = "azurerm:storagesync/v20181001:StorageSyncService"},
                    new Pulumi.Alias { Type = "azurerm:storagesync/v20190201:StorageSyncService"},
                    new Pulumi.Alias { Type = "azurerm:storagesync/v20190301:StorageSyncService"},
                    new Pulumi.Alias { Type = "azurerm:storagesync/v20190601:StorageSyncService"},
                    new Pulumi.Alias { Type = "azurerm:storagesync/v20191001:StorageSyncService"},
                    new Pulumi.Alias { Type = "azurerm:storagesync/v20200301:StorageSyncService"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing StorageSyncService resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StorageSyncService Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new StorageSyncService(name, id, options);
        }
    }

    public sealed class StorageSyncServiceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required. Gets or sets the location of the resource. This will be one of the supported and registered Azure Geo Regions (e.g. West US, East US, Southeast Asia, etc.). The geo region of a resource cannot be changed once it is created, but if an identical geo region is specified on update, the request will succeed.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of Storage Sync Service resource.
        /// </summary>
        [Input("storageSyncServiceName", required: true)]
        public Input<string> StorageSyncServiceName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Gets or sets a list of key value pairs that describe the resource. These tags can be used for viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key with a length no greater than 128 characters and a value with a length no greater than 256 characters.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public StorageSyncServiceArgs()
        {
        }
    }
}
