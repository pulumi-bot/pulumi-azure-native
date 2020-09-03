// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.RecoveryServices.V20180710
{
    /// <summary>
    /// Migration item.
    /// 
    /// ## Example Usage
    /// ### Enables migration.
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var replicationMigrationItem = new AzureRM.RecoveryServices.V20180710.ReplicationMigrationItem("replicationMigrationItem", new AzureRM.RecoveryServices.V20180710.ReplicationMigrationItemArgs
    ///         {
    ///             FabricName = "vmwarefabric1",
    ///             MigrationItemName = "virtualmachine1",
    ///             ProtectionContainerName = "vmwareContainer1",
    ///             ResourceGroupName = "resourcegroup1",
    ///             ResourceName = "migrationvault",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class ReplicationMigrationItem : Pulumi.CustomResource
    {
        /// <summary>
        /// Resource Location
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// Resource Name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The migration item properties.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.MigrationItemPropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// Resource Type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ReplicationMigrationItem resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ReplicationMigrationItem(string name, ReplicationMigrationItemArgs args, CustomResourceOptions? options = null)
            : base("azurerm:recoveryservices/v20180710:ReplicationMigrationItem", name, args ?? new ReplicationMigrationItemArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ReplicationMigrationItem(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:recoveryservices/v20180710:ReplicationMigrationItem", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:recoveryservices/latest:ReplicationMigrationItem"},
                    new Pulumi.Alias { Type = "azurerm:recoveryservices/v20180110:ReplicationMigrationItem"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ReplicationMigrationItem resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ReplicationMigrationItem Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ReplicationMigrationItem(name, id, options);
        }
    }

    public sealed class ReplicationMigrationItemArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Fabric name.
        /// </summary>
        [Input("fabricName", required: true)]
        public Input<string> FabricName { get; set; } = null!;

        /// <summary>
        /// Migration item name.
        /// </summary>
        [Input("migrationItemName", required: true)]
        public Input<string> MigrationItemName { get; set; } = null!;

        /// <summary>
        /// Enable migration input properties.
        /// </summary>
        [Input("properties", required: true)]
        public Input<Inputs.EnableMigrationInputPropertiesArgs> Properties { get; set; } = null!;

        /// <summary>
        /// Protection container name.
        /// </summary>
        [Input("protectionContainerName", required: true)]
        public Input<string> ProtectionContainerName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group where the recovery services vault is present.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the recovery services vault.
        /// </summary>
        [Input("resourceName", required: true)]
        public Input<string> ResourceName { get; set; } = null!;

        public ReplicationMigrationItemArgs()
        {
        }
    }
}
