// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DocumentDB.V20150401
{
    /// <summary>
    /// An Azure Cosmos DB MongoDB collection.
    /// 
    /// ## Example Usage
    /// ### CosmosDBMongoDBCollectionCreateUpdate
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var databaseAccountMongoDBCollection = new AzureRM.DocumentDB.V20150401.DatabaseAccountMongoDBCollection("databaseAccountMongoDBCollection", new AzureRM.DocumentDB.V20150401.DatabaseAccountMongoDBCollectionArgs
    ///         {
    ///             AccountName = "ddb1",
    ///             CollectionName = "collectionName",
    ///             DatabaseName = "databaseName",
    ///             Options = ,
    ///             Resource = new AzureRM.DocumentDB.V20150401.Inputs.MongoDBCollectionResourceArgs
    ///             {
    ///                 Id = "testcoll",
    ///                 Indexes = 
    ///                 {
    ///                     new AzureRM.DocumentDB.V20150401.Inputs.MongoIndexArgs
    ///                     {
    ///                         Key = new AzureRM.DocumentDB.V20150401.Inputs.MongoIndexKeysArgs
    ///                         {
    ///                             Keys = 
    ///                             {
    ///                                 "testKey",
    ///                             },
    ///                         },
    ///                         Options = new AzureRM.DocumentDB.V20150401.Inputs.MongoIndexOptionsArgs
    ///                         {
    ///                             ExpireAfterSeconds = 100,
    ///                             Unique = true,
    ///                         },
    ///                     },
    ///                 },
    ///                 ShardKey = 
    ///                 {
    ///                     { "testKey", "Hash" },
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
    public partial class DatabaseAccountMongoDBCollection : Pulumi.CustomResource
    {
        /// <summary>
        /// List of index keys
        /// </summary>
        [Output("indexes")]
        public Output<ImmutableArray<Outputs.MongoIndexResponseResult>> Indexes { get; private set; } = null!;

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
        /// A key-value pair of shard keys to be applied for the request.
        /// </summary>
        [Output("shardKey")]
        public Output<ImmutableDictionary<string, string>?> ShardKey { get; private set; } = null!;

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
        /// Create a DatabaseAccountMongoDBCollection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DatabaseAccountMongoDBCollection(string name, DatabaseAccountMongoDBCollectionArgs args, CustomResourceOptions? options = null)
            : base("azurerm:documentdb/v20150401:DatabaseAccountMongoDBCollection", name, args ?? new DatabaseAccountMongoDBCollectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DatabaseAccountMongoDBCollection(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:documentdb/v20150401:DatabaseAccountMongoDBCollection", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:documentdb/latest:DatabaseAccountMongoDBCollection"},
                    new Pulumi.Alias { Type = "azurerm:documentdb/v20150408:DatabaseAccountMongoDBCollection"},
                    new Pulumi.Alias { Type = "azurerm:documentdb/v20151106:DatabaseAccountMongoDBCollection"},
                    new Pulumi.Alias { Type = "azurerm:documentdb/v20160319:DatabaseAccountMongoDBCollection"},
                    new Pulumi.Alias { Type = "azurerm:documentdb/v20160331:DatabaseAccountMongoDBCollection"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DatabaseAccountMongoDBCollection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DatabaseAccountMongoDBCollection Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DatabaseAccountMongoDBCollection(name, id, options);
        }
    }

    public sealed class DatabaseAccountMongoDBCollectionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cosmos DB database account name.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// Cosmos DB collection name.
        /// </summary>
        [Input("collectionName", required: true)]
        public Input<string> CollectionName { get; set; } = null!;

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
        /// The standard JSON format of a MongoDB collection
        /// </summary>
        [Input("resource", required: true)]
        public Input<Inputs.MongoDBCollectionResourceArgs> Resource { get; set; } = null!;

        /// <summary>
        /// Name of an Azure resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public DatabaseAccountMongoDBCollectionArgs()
        {
        }
    }
}
