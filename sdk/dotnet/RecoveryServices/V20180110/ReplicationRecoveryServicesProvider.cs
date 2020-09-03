// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.RecoveryServices.V20180110
{
    /// <summary>
    /// Provider details.
    /// 
    /// ## Example Usage
    /// ### Adds a recovery services provider.
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var replicationRecoveryServicesProvider = new AzureRM.RecoveryServices.V20180110.ReplicationRecoveryServicesProvider("replicationRecoveryServicesProvider", new AzureRM.RecoveryServices.V20180110.ReplicationRecoveryServicesProviderArgs
    ///         {
    ///             FabricName = "vmwarefabric1",
    ///             ProviderName = "vmwareprovider1",
    ///             ResourceGroupName = "resourcegroup1",
    ///             ResourceName = "migrationvault",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class ReplicationRecoveryServicesProvider : Pulumi.CustomResource
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
        /// Provider properties.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.RecoveryServicesProviderPropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// Resource Type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ReplicationRecoveryServicesProvider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ReplicationRecoveryServicesProvider(string name, ReplicationRecoveryServicesProviderArgs args, CustomResourceOptions? options = null)
            : base("azurerm:recoveryservices/v20180110:ReplicationRecoveryServicesProvider", name, args ?? new ReplicationRecoveryServicesProviderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ReplicationRecoveryServicesProvider(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:recoveryservices/v20180110:ReplicationRecoveryServicesProvider", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:recoveryservices/latest:ReplicationRecoveryServicesProvider"},
                    new Pulumi.Alias { Type = "azurerm:recoveryservices/v20180710:ReplicationRecoveryServicesProvider"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ReplicationRecoveryServicesProvider resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ReplicationRecoveryServicesProvider Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ReplicationRecoveryServicesProvider(name, id, options);
        }
    }

    public sealed class ReplicationRecoveryServicesProviderArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Fabric name.
        /// </summary>
        [Input("fabricName", required: true)]
        public Input<string> FabricName { get; set; } = null!;

        /// <summary>
        /// The properties of an add provider request.
        /// </summary>
        [Input("properties", required: true)]
        public Input<Inputs.AddRecoveryServicesProviderInputPropertiesArgs> Properties { get; set; } = null!;

        /// <summary>
        /// Recovery services provider name.
        /// </summary>
        [Input("providerName", required: true)]
        public Input<string> ProviderName { get; set; } = null!;

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

        public ReplicationRecoveryServicesProviderArgs()
        {
        }
    }
}
