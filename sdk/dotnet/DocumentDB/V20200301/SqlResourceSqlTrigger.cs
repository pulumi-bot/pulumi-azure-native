// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DocumentDB.V20200301
{
    /// <summary>
    /// An Azure Cosmos DB trigger.
    /// 
    /// ## Example Usage
    /// ### CosmosDBSqlTriggerCreateUpdate
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var sqlResourceSqlTrigger = new AzureRM.DocumentDB.V20200301.SqlResourceSqlTrigger("sqlResourceSqlTrigger", new AzureRM.DocumentDB.V20200301.SqlResourceSqlTriggerArgs
    ///         {
    ///             AccountName = "ddb1",
    ///             ContainerName = "containerName",
    ///             DatabaseName = "databaseName",
    ///             Options = ,
    ///             Resource = new AzureRM.DocumentDB.V20200301.Inputs.SqlTriggerResourceArgs
    ///             {
    ///                 Body = "body",
    ///                 Id = "triggerName",
    ///                 TriggerOperation = "triggerOperation",
    ///                 TriggerType = "triggerType",
    ///             },
    ///             ResourceGroupName = "rg1",
    ///             TriggerName = "triggerName",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class SqlResourceSqlTrigger : Pulumi.CustomResource
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
        public Output<Outputs.SqlTriggerGetPropertiesResponseResourceResult?> Resource { get; private set; } = null!;

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
        /// Create a SqlResourceSqlTrigger resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SqlResourceSqlTrigger(string name, SqlResourceSqlTriggerArgs args, CustomResourceOptions? options = null)
            : base("azurerm:documentdb/v20200301:SqlResourceSqlTrigger", name, args ?? new SqlResourceSqlTriggerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SqlResourceSqlTrigger(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:documentdb/v20200301:SqlResourceSqlTrigger", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:documentdb/latest:SqlResourceSqlTrigger"},
                    new Pulumi.Alias { Type = "azurerm:documentdb/v20190801:SqlResourceSqlTrigger"},
                    new Pulumi.Alias { Type = "azurerm:documentdb/v20191212:SqlResourceSqlTrigger"},
                    new Pulumi.Alias { Type = "azurerm:documentdb/v20200401:SqlResourceSqlTrigger"},
                    new Pulumi.Alias { Type = "azurerm:documentdb/v20200601preview:SqlResourceSqlTrigger"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SqlResourceSqlTrigger resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SqlResourceSqlTrigger Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new SqlResourceSqlTrigger(name, id, options);
        }
    }

    public sealed class SqlResourceSqlTriggerArgs : Pulumi.ResourceArgs
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

        /// <summary>
        /// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
        /// </summary>
        [Input("options", required: true)]
        public Input<Inputs.CreateUpdateOptionsArgs> Options { get; set; } = null!;

        /// <summary>
        /// The standard JSON format of a trigger
        /// </summary>
        [Input("resource", required: true)]
        public Input<Inputs.SqlTriggerResourceArgs> Resource { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
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

        /// <summary>
        /// Cosmos DB trigger name.
        /// </summary>
        [Input("triggerName", required: true)]
        public Input<string> TriggerName { get; set; } = null!;

        public SqlResourceSqlTriggerArgs()
        {
        }
    }
}
