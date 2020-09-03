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
    /// Network Mapping model. Ideally it should have been possible to inherit this class from prev version in InheritedModels as long as there is no difference in structure or method signature. Since there were no base Models for certain fields and methods viz NetworkMappingProperties and Load with required return type, the class has been introduced in its entirety with references to base models to facilitate extensions in subsequent versions.
    /// 
    /// ## Example Usage
    /// ### Creates network mapping.
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var replicationNetworkMapping = new AzureRM.RecoveryServices.V20180110.ReplicationNetworkMapping("replicationNetworkMapping", new AzureRM.RecoveryServices.V20180110.ReplicationNetworkMappingArgs
    ///         {
    ///             FabricName = "b0cef6e9a4437b81803d0b55ada4f700ab66caae59c35d62723a1589c0cd13ac",
    ///             NetworkMappingName = "corpe2amap",
    ///             NetworkName = "e2267b5c-2650-49bd-ab3f-d66aae694c06",
    ///             ResourceGroupName = "srcBvte2a14C27",
    ///             ResourceName = "srce2avaultbvtaC27",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class ReplicationNetworkMapping : Pulumi.CustomResource
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
        /// The Network Mapping Properties.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.NetworkMappingPropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// Resource Type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ReplicationNetworkMapping resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ReplicationNetworkMapping(string name, ReplicationNetworkMappingArgs args, CustomResourceOptions? options = null)
            : base("azurerm:recoveryservices/v20180110:ReplicationNetworkMapping", name, args ?? new ReplicationNetworkMappingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ReplicationNetworkMapping(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:recoveryservices/v20180110:ReplicationNetworkMapping", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:recoveryservices/latest:ReplicationNetworkMapping"},
                    new Pulumi.Alias { Type = "azurerm:recoveryservices/v20160810:ReplicationNetworkMapping"},
                    new Pulumi.Alias { Type = "azurerm:recoveryservices/v20180710:ReplicationNetworkMapping"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ReplicationNetworkMapping resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ReplicationNetworkMapping Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ReplicationNetworkMapping(name, id, options);
        }
    }

    public sealed class ReplicationNetworkMappingArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Primary fabric name.
        /// </summary>
        [Input("fabricName", required: true)]
        public Input<string> FabricName { get; set; } = null!;

        /// <summary>
        /// Network mapping name.
        /// </summary>
        [Input("networkMappingName", required: true)]
        public Input<string> NetworkMappingName { get; set; } = null!;

        /// <summary>
        /// Primary network name.
        /// </summary>
        [Input("networkName", required: true)]
        public Input<string> NetworkName { get; set; } = null!;

        /// <summary>
        /// Input properties for creating network mapping.
        /// </summary>
        [Input("properties")]
        public Input<Inputs.CreateNetworkMappingInputPropertiesArgs>? Properties { get; set; }

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

        public ReplicationNetworkMappingArgs()
        {
        }
    }
}
