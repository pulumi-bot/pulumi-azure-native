// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DocumentDB.V20190801
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
    ///         var sqlResourceSqlContainer = new AzureRM.DocumentDB.V20190801.SqlResourceSqlContainer("sqlResourceSqlContainer", new AzureRM.DocumentDB.V20190801.SqlResourceSqlContainerArgs
    ///         {
    ///             AccountName = "ddb1",
    ///             ContainerName = "containerName",
    ///             DatabaseName = "databaseName",
    ///             Location = "West US",
    ///             Options = ,
    ///             Resource = new AzureRM.DocumentDB.V20190801.Inputs.SqlContainerResourceArgs
    ///             {
    ///                 ConflictResolutionPolicy = new AzureRM.DocumentDB.V20190801.Inputs.ConflictResolutionPolicyArgs
    ///                 {
    ///                     ConflictResolutionPath = "/path",
    ///                     Mode = "LastWriterWins",
    ///                 },
    ///                 DefaultTtl = 100,
    ///                 Id = "containerName",
    ///                 IndexingPolicy = new AzureRM.DocumentDB.V20190801.Inputs.IndexingPolicyArgs
    ///                 {
    ///                     Automatic = true,
    ///                     ExcludedPaths = {},
    ///                     IncludedPaths = 
    ///                     {
    ///                         new AzureRM.DocumentDB.V20190801.Inputs.IncludedPathArgs
    ///                         {
    ///                             Indexes = 
    ///                             {
    ///                                 new AzureRM.DocumentDB.V20190801.Inputs.IndexesArgs
    ///                                 {
    ///                                     DataType = "String",
    ///                                     Kind = "Range",
    ///                                     Precision = -1,
    ///                                 },
    ///                                 new AzureRM.DocumentDB.V20190801.Inputs.IndexesArgs
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
    ///                 PartitionKey = new AzureRM.DocumentDB.V20190801.Inputs.ContainerPartitionKeyArgs
    ///                 {
    ///                     Kind = "Hash",
    ///                     Paths = 
    ///                     {
    ///                         "/AccountNumber",
    ///                     },
    ///                 },
    ///                 UniqueKeyPolicy = new AzureRM.DocumentDB.V20190801.Inputs.UniqueKeyPolicyArgs
    ///                 {
    ///                     UniqueKeys = 
    ///                     {
    ///                         new AzureRM.DocumentDB.V20190801.Inputs.UniqueKeyArgs
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
    ///             Tags = ,
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class SqlResourceSqlContainer : Pulumi.CustomResource
    {
        /// <summary>
        /// The location of the resource group to which the resource belongs.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the ARM resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("resource")]
        public Output<Outputs.SqlContainerGetPropertiesResponseResourceResult?> Resource { get; private set; } = null!;

        /// <summary>
        /// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of Azure resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a SqlResourceSqlContainer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SqlResourceSqlContainer(string name, SqlResourceSqlContainerArgs args, CustomResourceOptions? options = null)
            : base("azurerm:documentdb/v20190801:SqlResourceSqlContainer", name, args ?? new SqlResourceSqlContainerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SqlResourceSqlContainer(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:documentdb/v20190801:SqlResourceSqlContainer", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:documentdb/latest:SqlResourceSqlContainer"},
                    new Pulumi.Alias { Type = "azurerm:documentdb/v20191212:SqlResourceSqlContainer"},
                    new Pulumi.Alias { Type = "azurerm:documentdb/v20200301:SqlResourceSqlContainer"},
                    new Pulumi.Alias { Type = "azurerm:documentdb/v20200401:SqlResourceSqlContainer"},
                    new Pulumi.Alias { Type = "azurerm:documentdb/v20200601preview:SqlResourceSqlContainer"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SqlResourceSqlContainer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SqlResourceSqlContainer Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new SqlResourceSqlContainer(name, id, options);
        }
    }

    public sealed class SqlResourceSqlContainerArgs : Pulumi.ResourceArgs
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

        /// <summary>
        /// The location of the resource group to which the resource belongs.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

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

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public SqlResourceSqlContainerArgs()
        {
        }
    }
}
