// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Kusto.V20191109
{
    /// <summary>
    /// Class representing a Kusto database.
    /// 
    /// ## Example Usage
    /// ### KustoDatabasesCreateOrUpdate
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var database = new AzureRM.Kusto.V20191109.Database("database", new AzureRM.Kusto.V20191109.DatabaseArgs
    ///         {
    ///             ClusterName = "KustoClusterRPTest4",
    ///             DatabaseName = "KustoDatabase8",
    ///             Location = "westus",
    ///             ResourceGroupName = "kustorptest",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class Database : Pulumi.CustomResource
    {
        /// <summary>
        /// Kind of the database
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Resource location.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Database resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Database(string name, DatabaseArgs args, CustomResourceOptions? options = null)
            : base("azurerm:kusto/v20191109:Database", name, args ?? new DatabaseArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Database(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:kusto/v20191109:Database", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:kusto/latest:Database"},
                    new Pulumi.Alias { Type = "azurerm:kusto/v20170907privatepreview:Database"},
                    new Pulumi.Alias { Type = "azurerm:kusto/v20180907preview:Database"},
                    new Pulumi.Alias { Type = "azurerm:kusto/v20190121:Database"},
                    new Pulumi.Alias { Type = "azurerm:kusto/v20190515:Database"},
                    new Pulumi.Alias { Type = "azurerm:kusto/v20190907:Database"},
                    new Pulumi.Alias { Type = "azurerm:kusto/v20200215:Database"},
                    new Pulumi.Alias { Type = "azurerm:kusto/v20200614:Database"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Database resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Database Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Database(name, id, options);
        }
    }

    public sealed class DatabaseArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Kusto cluster.
        /// </summary>
        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        /// <summary>
        /// The name of the database in the Kusto cluster.
        /// </summary>
        [Input("databaseName", required: true)]
        public Input<string> DatabaseName { get; set; } = null!;

        /// <summary>
        /// Kind of the database
        /// </summary>
        [Input("kind", required: true)]
        public Input<string> Kind { get; set; } = null!;

        /// <summary>
        /// Resource location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the resource group containing the Kusto cluster.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public DatabaseArgs()
        {
        }
    }
}
