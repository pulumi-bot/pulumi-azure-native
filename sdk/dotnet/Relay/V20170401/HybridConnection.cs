// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Relay.V20170401
{
    /// <summary>
    /// Description of hybrid connection resource.
    /// 
    /// ## Example Usage
    /// ### RelayHybridConnectionCreate
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var hybridConnection = new AzureRM.Relay.V20170401.HybridConnection("hybridConnection", new AzureRM.Relay.V20170401.HybridConnectionArgs
    ///         {
    ///             HybridConnectionName = "sdk-Relay-Hybrid-01",
    ///             NamespaceName = "sdk-RelayNamespace-01",
    ///             RequiresClientAuthorization = true,
    ///             ResourceGroupName = "RG-eg",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class HybridConnection : Pulumi.CustomResource
    {
        /// <summary>
        /// The time the hybrid connection was created.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The number of listeners for this hybrid connection. Note that min : 1 and max:25 are supported.
        /// </summary>
        [Output("listenerCount")]
        public Output<int> ListenerCount { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Returns true if client authorization is needed for this hybrid connection; otherwise, false.
        /// </summary>
        [Output("requiresClientAuthorization")]
        public Output<bool?> RequiresClientAuthorization { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The time the namespace was updated.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// The usermetadata is a placeholder to store user-defined string data for the hybrid connection endpoint. For example, it can be used to store descriptive data, such as a list of teams and their contact information. Also, user-defined configuration settings can be stored.
        /// </summary>
        [Output("userMetadata")]
        public Output<string?> UserMetadata { get; private set; } = null!;


        /// <summary>
        /// Create a HybridConnection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HybridConnection(string name, HybridConnectionArgs args, CustomResourceOptions? options = null)
            : base("azurerm:relay/v20170401:HybridConnection", name, args ?? new HybridConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private HybridConnection(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:relay/v20170401:HybridConnection", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:relay/latest:HybridConnection"},
                    new Pulumi.Alias { Type = "azurerm:relay/v20160701:HybridConnection"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing HybridConnection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HybridConnection Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new HybridConnection(name, id, options);
        }
    }

    public sealed class HybridConnectionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The hybrid connection name.
        /// </summary>
        [Input("hybridConnectionName", required: true)]
        public Input<string> HybridConnectionName { get; set; } = null!;

        /// <summary>
        /// The namespace name
        /// </summary>
        [Input("namespaceName", required: true)]
        public Input<string> NamespaceName { get; set; } = null!;

        /// <summary>
        /// Returns true if client authorization is needed for this hybrid connection; otherwise, false.
        /// </summary>
        [Input("requiresClientAuthorization")]
        public Input<bool>? RequiresClientAuthorization { get; set; }

        /// <summary>
        /// Name of the Resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The usermetadata is a placeholder to store user-defined string data for the hybrid connection endpoint. For example, it can be used to store descriptive data, such as a list of teams and their contact information. Also, user-defined configuration settings can be stored.
        /// </summary>
        [Input("userMetadata")]
        public Input<string>? UserMetadata { get; set; }

        public HybridConnectionArgs()
        {
        }
    }
}
