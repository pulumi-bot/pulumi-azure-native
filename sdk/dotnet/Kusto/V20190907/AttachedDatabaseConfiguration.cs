// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Kusto.V20190907
{
    /// <summary>
    /// Class representing an attached database configuration.
    /// </summary>
    public partial class AttachedDatabaseConfiguration : Pulumi.CustomResource
    {
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
        /// The properties of the attached database configuration.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.AttachedDatabaseConfigurationPropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a AttachedDatabaseConfiguration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AttachedDatabaseConfiguration(string name, AttachedDatabaseConfigurationArgs args, CustomResourceOptions? options = null)
            : base("azurerm:kusto/v20190907:AttachedDatabaseConfiguration", name, args ?? new AttachedDatabaseConfigurationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AttachedDatabaseConfiguration(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:kusto/v20190907:AttachedDatabaseConfiguration", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AttachedDatabaseConfiguration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AttachedDatabaseConfiguration Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AttachedDatabaseConfiguration(name, id, options);
        }
    }

    public sealed class AttachedDatabaseConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Kusto cluster.
        /// </summary>
        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        /// <summary>
        /// The resource id of the cluster where the databases you would like to attach reside.
        /// </summary>
        [Input("clusterResourceId", required: true)]
        public Input<string> ClusterResourceId { get; set; } = null!;

        /// <summary>
        /// The name of the database which you would like to attach, use * if you want to follow all current and future databases.
        /// </summary>
        [Input("databaseName", required: true)]
        public Input<string> DatabaseName { get; set; } = null!;

        /// <summary>
        /// The default principals modification kind
        /// </summary>
        [Input("defaultPrincipalsModificationKind", required: true)]
        public Input<string> DefaultPrincipalsModificationKind { get; set; } = null!;

        /// <summary>
        /// Resource location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the attached database configuration.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group containing the Kusto cluster.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public AttachedDatabaseConfigurationArgs()
        {
        }
    }
}
