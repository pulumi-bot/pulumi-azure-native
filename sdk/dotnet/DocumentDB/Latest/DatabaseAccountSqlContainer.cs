// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DocumentDB.Latest
{
    /// <summary>
    /// An Azure Cosmos DB container.
    /// 
    /// ## Example Usage
    /// ### CosmosDBSqlContainerCreateUpdate
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var databaseAccountSqlContainer = new AzureRM.DocumentDB.Latest.DatabaseAccountSqlContainer("databaseAccountSqlContainer", new AzureRM.DocumentDB.Latest.DatabaseAccountSqlContainerArgs
    ///         {
    ///             AccountName = "ddb1",
    ///             ContainerName = "containerName",
    ///             DatabaseName = "databaseName",
    ///             Options = ,
    ///             Resource = new AzureRM.DocumentDB.Latest.Inputs.SqlContainerResourceArgs
    ///             {
    ///                 ConflictResolutionPolicy = new AzureRM.DocumentDB.Latest.Inputs.ConflictResolutionPolicyArgs
    ///                 {
    ///                     ConflictResolutionPath = "/path",
    ///                     Mode = "LastWriterWins",
    ///                 },
    ///                 DefaultTtl = 100,
    ///                 Id = "containerName",
    ///                 IndexingPolicy = new AzureRM.DocumentDB.Latest.Inputs.IndexingPolicyArgs
    ///                 {
    ///                     Automatic = true,
    ///                     ExcludedPaths = {},
    ///                     IncludedPaths = 
    ///                     {
    ///                         new AzureRM.DocumentDB.Latest.Inputs.IncludedPathArgs
    ///                         {
    ///                             Indexes = 
    ///                             {
    ///                                 new AzureRM.DocumentDB.Latest.Inputs.IndexesArgs
    ///                                 {
    ///                                     DataType = "String",
    ///                                     Kind = "Range",
    ///                                     Precision = -1,
    ///                                 },
    ///                                 new AzureRM.DocumentDB.Latest.Inputs.IndexesArgs
    ///                                 {
    ///                                     DataType = "Number",
    ///                                     Kind = "Range",
    ///                                     Precision = -1,
    ///                                 },
    ///                             },
    ///                             Path = "/*",
    ///                         },
    ///                     },
    ///                     IndexingMode = "Consistent",
    ///                 },
    ///                 PartitionKey = new AzureRM.DocumentDB.Latest.Inputs.ContainerPartitionKeyArgs
    ///                 {
    ///                     Kind = "Hash",
    ///                     Paths = 
    ///                     {
    ///                         "/AccountNumber",
    ///                     },
    ///                 },
    ///                 UniqueKeyPolicy = new AzureRM.DocumentDB.Latest.Inputs.UniqueKeyPolicyArgs
    ///                 {
    ///                     UniqueKeys = 
    ///                     {
    ///                         new AzureRM.DocumentDB.Latest.Inputs.UniqueKeyArgs
    ///                         {
    ///                             Paths = 
    ///                             {
    ///                                 "/testPath",
    ///                             },
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///             ResourceGroupName = "rg1",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class DatabaseAccountSqlContainer : Pulumi.CustomResource
    {
        /// <summary>
        /// The conflict resolution policy for the container.
        /// </summary>
        [Output("conflictResolutionPolicy")]
        public Output<Outputs.ConflictResolutionPolicyResponseResult?> ConflictResolutionPolicy { get; private set; } = null!;

        /// <summary>
        /// Default time to live
        /// </summary>
        [Output("defaultTtl")]
        public Output<int?> DefaultTtl { get; private set; } = null!;

        /// <summary>
        /// A system generated property representing the resource etag required for optimistic concurrency control.
        /// </summary>
        [Output("etag")]
        public Output<string?> Etag { get; private set; } = null!;

        /// <summary>
        /// The configuration of the indexing policy. By default, the indexing is automatic for all document paths within the container
        /// </summary>
        [Output("indexingPolicy")]
        public Output<Outputs.IndexingPolicyResponseResult?> IndexingPolicy { get; private set; } = null!;

        /// <summary>
        /// The location of the resource group to which the resource belongs.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the database account.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The configuration of the partition key to be used for partitioning data into multiple partitions
        /// </summary>
        [Output("partitionKey")]
        public Output<Outputs.ContainerPartitionKeyResponseResult?> PartitionKey { get; private set; } = null!;

        /// <summary>
        /// A system generated property. A unique identifier.
        /// </summary>
        [Output("rid")]
        public Output<string?> Rid { get; private set; } = null!;

        /// <summary>
        /// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A system generated property that denotes the last updated timestamp of the resource.
        /// </summary>
        [Output("ts")]
        public Output<ImmutableDictionary<string, object>?> Ts { get; private set; } = null!;

        /// <summary>
        /// The type of Azure resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The unique key policy configuration for specifying uniqueness constraints on documents in the collection in the Azure Cosmos DB service.
        /// </summary>
        [Output("uniqueKeyPolicy")]
        public Output<Outputs.UniqueKeyPolicyResponseResult?> UniqueKeyPolicy { get; private set; } = null!;


        /// <summary>
        /// Create a DatabaseAccountSqlContainer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DatabaseAccountSqlContainer(string name, DatabaseAccountSqlContainerArgs args, CustomResourceOptions? options = null)
            : base("azurerm:documentdb/latest:DatabaseAccountSqlContainer", name, args ?? new DatabaseAccountSqlContainerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DatabaseAccountSqlContainer(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:documentdb/latest:DatabaseAccountSqlContainer", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:documentdb/v20150401:DatabaseAccountSqlContainer"},
                    new Pulumi.Alias { Type = "azurerm:documentdb/v20150408:DatabaseAccountSqlContainer"},
                    new Pulumi.Alias { Type = "azurerm:documentdb/v20151106:DatabaseAccountSqlContainer"},
                    new Pulumi.Alias { Type = "azurerm:documentdb/v20160319:DatabaseAccountSqlContainer"},
                    new Pulumi.Alias { Type = "azurerm:documentdb/v20160331:DatabaseAccountSqlContainer"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DatabaseAccountSqlContainer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DatabaseAccountSqlContainer Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DatabaseAccountSqlContainer(name, id, options);
        }
    }

    public sealed class DatabaseAccountSqlContainerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cosmos DB database account name.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// Cosmos DB container name.
        /// </summary>
        [Input("containerName", required: true)]
        public Input<string> ContainerName { get; set; } = null!;

        /// <summary>
        /// Cosmos DB database name.
        /// </summary>
        [Input("databaseName", required: true)]
        public Input<string> DatabaseName { get; set; } = null!;

        [Input("options", required: true)]
        private InputMap<string>? _options;

        /// <summary>
        /// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
        /// </summary>
        public InputMap<string> Options
        {
            get => _options ?? (_options = new InputMap<string>());
            set => _options = value;
        }

        /// <summary>
        /// The standard JSON format of a container
        /// </summary>
        [Input("resource", required: true)]
        public Input<Inputs.SqlContainerResourceArgs> Resource { get; set; } = null!;

        /// <summary>
        /// Name of an Azure resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public DatabaseAccountSqlContainerArgs()
        {
        }
    }
}
