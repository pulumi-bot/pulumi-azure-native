// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Synapse.V20190601Preview
{
    /// <summary>
    /// A SQL Analytics pool
    /// </summary>
    [AzureNextGenResourceType("azure-nextgen:synapse/v20190601preview:SqlPool")]
    public partial class SqlPool : Pulumi.CustomResource
    {
        /// <summary>
        /// Collation mode
        /// </summary>
        [Output("collation")]
        public Output<string?> Collation { get; private set; } = null!;

        /// <summary>
        /// What is this?
        /// </summary>
        [Output("createMode")]
        public Output<string?> CreateMode { get; private set; } = null!;

        /// <summary>
        /// Date the SQL pool was created
        /// </summary>
        [Output("creationDate")]
        public Output<string?> CreationDate { get; private set; } = null!;

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Maximum size in bytes
        /// </summary>
        [Output("maxSizeBytes")]
        public Output<double?> MaxSizeBytes { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Resource state
        /// </summary>
        [Output("provisioningState")]
        public Output<string?> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Backup database to restore from
        /// </summary>
        [Output("recoverableDatabaseId")]
        public Output<string?> RecoverableDatabaseId { get; private set; } = null!;

        /// <summary>
        /// Snapshot time to restore
        /// </summary>
        [Output("restorePointInTime")]
        public Output<string?> RestorePointInTime { get; private set; } = null!;

        /// <summary>
        /// SQL pool SKU
        /// </summary>
        [Output("sku")]
        public Output<Outputs.SkuResponse?> Sku { get; private set; } = null!;

        /// <summary>
        /// Source database to create from
        /// </summary>
        [Output("sourceDatabaseId")]
        public Output<string?> SourceDatabaseId { get; private set; } = null!;

        /// <summary>
        /// Resource status
        /// </summary>
        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a SqlPool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SqlPool(string name, SqlPoolArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:synapse/v20190601preview:SqlPool", name, args ?? new SqlPoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SqlPool(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:synapse/v20190601preview:SqlPool", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:synapse/v20200401preview:SqlPool"},
                    new Pulumi.Alias { Type = "azure-nextgen:synapse/v20201201:SqlPool"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SqlPool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SqlPool Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new SqlPool(name, id, options);
        }
    }

    public sealed class SqlPoolArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Collation mode
        /// </summary>
        [Input("collation")]
        public Input<string>? Collation { get; set; }

        /// <summary>
        /// What is this?
        /// </summary>
        [Input("createMode")]
        public Input<string>? CreateMode { get; set; }

        /// <summary>
        /// Date the SQL pool was created
        /// </summary>
        [Input("creationDate")]
        public Input<string>? CreationDate { get; set; }

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Maximum size in bytes
        /// </summary>
        [Input("maxSizeBytes")]
        public Input<double>? MaxSizeBytes { get; set; }

        /// <summary>
        /// Resource state
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        /// <summary>
        /// Backup database to restore from
        /// </summary>
        [Input("recoverableDatabaseId")]
        public Input<string>? RecoverableDatabaseId { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Snapshot time to restore
        /// </summary>
        [Input("restorePointInTime")]
        public Input<string>? RestorePointInTime { get; set; }

        /// <summary>
        /// SQL pool SKU
        /// </summary>
        [Input("sku")]
        public Input<Inputs.SkuArgs>? Sku { get; set; }

        /// <summary>
        /// Source database to create from
        /// </summary>
        [Input("sourceDatabaseId")]
        public Input<string>? SourceDatabaseId { get; set; }

        /// <summary>
        /// SQL pool name
        /// </summary>
        [Input("sqlPoolName", required: true)]
        public Input<string> SqlPoolName { get; set; } = null!;

        /// <summary>
        /// Resource status
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The name of the workspace
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public SqlPoolArgs()
        {
        }
    }
}
