// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Storage.V20160101
{
    /// <summary>
    /// The storage account.
    /// </summary>
    [AzureNextGenResourceType("azure-nextgen:storage/v20160101:StorageAccount")]
    public partial class StorageAccount : Pulumi.CustomResource
    {
        /// <summary>
        /// Required for storage accounts where kind = BlobStorage. The access tier used for billing.
        /// </summary>
        [Output("accessTier")]
        public Output<string> AccessTier { get; private set; } = null!;

        /// <summary>
        /// Gets the creation date and time of the storage account in UTC.
        /// </summary>
        [Output("creationTime")]
        public Output<string> CreationTime { get; private set; } = null!;

        /// <summary>
        /// Gets the custom domain the user assigned to this storage account.
        /// </summary>
        [Output("customDomain")]
        public Output<Outputs.CustomDomainResponse> CustomDomain { get; private set; } = null!;

        /// <summary>
        /// Gets the encryption settings on the account. If unspecified, the account is unencrypted.
        /// </summary>
        [Output("encryption")]
        public Output<Outputs.EncryptionResponse> Encryption { get; private set; } = null!;

        /// <summary>
        /// Gets the Kind.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Gets the timestamp of the most recent instance of a failover to the secondary location. Only the most recent timestamp is retained. This element is not returned if there has never been a failover instance. Only available if the accountType is Standard_GRS or Standard_RAGRS.
        /// </summary>
        [Output("lastGeoFailoverTime")]
        public Output<string> LastGeoFailoverTime { get; private set; } = null!;

        /// <summary>
        /// Resource location
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// Resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Gets the URLs that are used to perform a retrieval of a public blob, queue, or table object. Note that Standard_ZRS and Premium_LRS accounts only return the blob endpoint.
        /// </summary>
        [Output("primaryEndpoints")]
        public Output<Outputs.EndpointsResponse> PrimaryEndpoints { get; private set; } = null!;

        /// <summary>
        /// Gets the location of the primary data center for the storage account.
        /// </summary>
        [Output("primaryLocation")]
        public Output<string> PrimaryLocation { get; private set; } = null!;

        /// <summary>
        /// Gets the status of the storage account at the time the operation was called.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Gets the URLs that are used to perform a retrieval of a public blob, queue, or table object from the secondary location of the storage account. Only available if the SKU name is Standard_RAGRS.
        /// </summary>
        [Output("secondaryEndpoints")]
        public Output<Outputs.EndpointsResponse> SecondaryEndpoints { get; private set; } = null!;

        /// <summary>
        /// Gets the location of the geo-replicated secondary for the storage account. Only available if the accountType is Standard_GRS or Standard_RAGRS.
        /// </summary>
        [Output("secondaryLocation")]
        public Output<string> SecondaryLocation { get; private set; } = null!;

        /// <summary>
        /// Gets the SKU.
        /// </summary>
        [Output("sku")]
        public Output<Outputs.SkuResponse> Sku { get; private set; } = null!;

        /// <summary>
        /// Gets the status indicating whether the primary location of the storage account is available or unavailable.
        /// </summary>
        [Output("statusOfPrimary")]
        public Output<string> StatusOfPrimary { get; private set; } = null!;

        /// <summary>
        /// Gets the status indicating whether the secondary location of the storage account is available or unavailable. Only available if the SKU name is Standard_GRS or Standard_RAGRS.
        /// </summary>
        [Output("statusOfSecondary")]
        public Output<string> StatusOfSecondary { get; private set; } = null!;

        /// <summary>
        /// Tags assigned to a resource; can be used for viewing and grouping a resource (across resource groups).
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Resource type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a StorageAccount resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StorageAccount(string name, StorageAccountArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:storage/v20160101:StorageAccount", name, args ?? new StorageAccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StorageAccount(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:storage/v20160101:StorageAccount", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:storage/latest:StorageAccount"},
                    new Pulumi.Alias { Type = "azure-nextgen:storage/v20150501preview:StorageAccount"},
                    new Pulumi.Alias { Type = "azure-nextgen:storage/v20150615:StorageAccount"},
                    new Pulumi.Alias { Type = "azure-nextgen:storage/v20160501:StorageAccount"},
                    new Pulumi.Alias { Type = "azure-nextgen:storage/v20161201:StorageAccount"},
                    new Pulumi.Alias { Type = "azure-nextgen:storage/v20170601:StorageAccount"},
                    new Pulumi.Alias { Type = "azure-nextgen:storage/v20171001:StorageAccount"},
                    new Pulumi.Alias { Type = "azure-nextgen:storage/v20180201:StorageAccount"},
                    new Pulumi.Alias { Type = "azure-nextgen:storage/v20180301preview:StorageAccount"},
                    new Pulumi.Alias { Type = "azure-nextgen:storage/v20180701:StorageAccount"},
                    new Pulumi.Alias { Type = "azure-nextgen:storage/v20181101:StorageAccount"},
                    new Pulumi.Alias { Type = "azure-nextgen:storage/v20190401:StorageAccount"},
                    new Pulumi.Alias { Type = "azure-nextgen:storage/v20190601:StorageAccount"},
                    new Pulumi.Alias { Type = "azure-nextgen:storage/v20200801preview:StorageAccount"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing StorageAccount resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StorageAccount Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new StorageAccount(name, id, options);
        }
    }

    public sealed class StorageAccountArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required for storage accounts where kind = BlobStorage. The access tier used for billing.
        /// </summary>
        [Input("accessTier")]
        public Input<Pulumi.AzureNextGen.Storage.V20160101.AccessTier>? AccessTier { get; set; }

        /// <summary>
        /// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// User domain assigned to the storage account. Name is the CNAME source. Only one custom domain is supported per storage account at this time. To clear the existing custom domain, use an empty string for the custom domain name property.
        /// </summary>
        [Input("customDomain")]
        public Input<Inputs.CustomDomainArgs>? CustomDomain { get; set; }

        /// <summary>
        /// Provides the encryption settings on the account. If left unspecified the account encryption settings will remain the same. The default setting is unencrypted.
        /// </summary>
        [Input("encryption")]
        public Input<Inputs.EncryptionArgs>? Encryption { get; set; }

        /// <summary>
        /// Required. Indicates the type of storage account.
        /// </summary>
        [Input("kind", required: true)]
        public Input<Pulumi.AzureNextGen.Storage.V20160101.Kind> Kind { get; set; } = null!;

        /// <summary>
        /// Required. Gets or sets the location of the resource. This will be one of the supported and registered Azure Geo Regions (e.g. West US, East US, Southeast Asia, etc.). The geo region of a resource cannot be changed once it is created, but if an identical geo region is specified on update, the request will succeed.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the resource group within the user's subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Required. Gets or sets the sku name.
        /// </summary>
        [Input("sku", required: true)]
        public Input<Inputs.SkuArgs> Sku { get; set; } = null!;

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

        public StorageAccountArgs()
        {
        }
    }
}
