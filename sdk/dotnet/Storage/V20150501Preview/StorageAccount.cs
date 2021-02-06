// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Storage.V20150501Preview
{
    /// <summary>
    /// The storage account.
    /// </summary>
    [AzureNextGenResourceType("azure-nextgen:storage/v20150501preview:StorageAccount")]
    public partial class StorageAccount : Pulumi.CustomResource
    {
        /// <summary>
        /// Gets the type of the storage account.
        /// </summary>
        [Output("accountType")]
        public Output<string?> AccountType { get; private set; } = null!;

        /// <summary>
        /// Gets the creation date and time of the storage account in UTC.
        /// </summary>
        [Output("creationTime")]
        public Output<string?> CreationTime { get; private set; } = null!;

        /// <summary>
        /// Gets the user assigned custom domain assigned to this storage account.
        /// </summary>
        [Output("customDomain")]
        public Output<Outputs.CustomDomainResponse?> CustomDomain { get; private set; } = null!;

        /// <summary>
        /// Gets the timestamp of the most recent instance of a failover to the secondary location. Only the most recent timestamp is retained. This element is not returned if there has never been a failover instance. Only available if the accountType is StandardGRS or StandardRAGRS.
        /// </summary>
        [Output("lastGeoFailoverTime")]
        public Output<string?> LastGeoFailoverTime { get; private set; } = null!;

        /// <summary>
        /// Resource location
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Gets the URLs that are used to perform a retrieval of a public blob, queue or table object.Note that StandardZRS and PremiumLRS accounts only return the blob endpoint.
        /// </summary>
        [Output("primaryEndpoints")]
        public Output<Outputs.EndpointsResponse?> PrimaryEndpoints { get; private set; } = null!;

        /// <summary>
        /// Gets the location of the primary for the storage account.
        /// </summary>
        [Output("primaryLocation")]
        public Output<string?> PrimaryLocation { get; private set; } = null!;

        /// <summary>
        /// Gets the status of the storage account at the time the operation was called.
        /// </summary>
        [Output("provisioningState")]
        public Output<string?> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Gets the URLs that are used to perform a retrieval of a public blob, queue or table object from the secondary location of the storage account. Only available if the accountType is StandardRAGRS.
        /// </summary>
        [Output("secondaryEndpoints")]
        public Output<Outputs.EndpointsResponse?> SecondaryEndpoints { get; private set; } = null!;

        /// <summary>
        /// Gets the location of the geo replicated secondary for the storage account. Only available if the accountType is StandardGRS or StandardRAGRS.
        /// </summary>
        [Output("secondaryLocation")]
        public Output<string?> SecondaryLocation { get; private set; } = null!;

        /// <summary>
        /// Gets the status indicating whether the primary location of the storage account is available or unavailable.
        /// </summary>
        [Output("statusOfPrimary")]
        public Output<string?> StatusOfPrimary { get; private set; } = null!;

        /// <summary>
        /// Gets the status indicating whether the secondary location of the storage account is available or unavailable. Only available if the accountType is StandardGRS or StandardRAGRS.
        /// </summary>
        [Output("statusOfSecondary")]
        public Output<string?> StatusOfSecondary { get; private set; } = null!;

        /// <summary>
        /// Resource tags
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
            : base("azure-nextgen:storage/v20150501preview:StorageAccount", name, args ?? new StorageAccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StorageAccount(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:storage/v20150501preview:StorageAccount", name, null, MakeResourceOptions(options, id))
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
                    new Pulumi.Alias { Type = "azure-nextgen:storage/v20150615:StorageAccount"},
                    new Pulumi.Alias { Type = "azure-nextgen:storage/v20160101:StorageAccount"},
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
        /// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.  
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// Gets or sets the account type.
        /// </summary>
        [Input("accountType")]
        public Input<Pulumi.AzureNextGen.Storage.V20150501Preview.AccountType>? AccountType { get; set; }

        /// <summary>
        /// Resource location
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the resource group within the user's subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags
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
