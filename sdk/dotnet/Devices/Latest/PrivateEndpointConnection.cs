// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Devices.Latest
{
    /// <summary>
    /// The private endpoint connection of an IotHub
    /// 
    /// ## Example Usage
    /// ### PrivateEndpointConnection_Update
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var privateEndpointConnection = new AzureRM.Devices.Latest.PrivateEndpointConnection("privateEndpointConnection", new AzureRM.Devices.Latest.PrivateEndpointConnectionArgs
    ///         {
    ///             PrivateEndpointConnectionName = "myPrivateEndpointConnection",
    ///             ResourceGroupName = "myResourceGroup",
    ///             ResourceName = "testHub",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class PrivateEndpointConnection : Pulumi.CustomResource
    {
        /// <summary>
        /// The resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The properties of a private endpoint connection
        /// </summary>
        [Output("properties")]
        public Output<Outputs.PrivateEndpointConnectionPropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// The resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a PrivateEndpointConnection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PrivateEndpointConnection(string name, PrivateEndpointConnectionArgs args, CustomResourceOptions? options = null)
            : base("azurerm:devices/latest:PrivateEndpointConnection", name, args ?? new PrivateEndpointConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PrivateEndpointConnection(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:devices/latest:PrivateEndpointConnection", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:devices/v20200301:PrivateEndpointConnection"},
                    new Pulumi.Alias { Type = "azurerm:devices/v20200401:PrivateEndpointConnection"},
                    new Pulumi.Alias { Type = "azurerm:devices/v20200615:PrivateEndpointConnection"},
                    new Pulumi.Alias { Type = "azurerm:devices/v20200710preview:PrivateEndpointConnection"},
                    new Pulumi.Alias { Type = "azurerm:devices/v20200801:PrivateEndpointConnection"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PrivateEndpointConnection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PrivateEndpointConnection Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PrivateEndpointConnection(name, id, options);
        }
    }

    public sealed class PrivateEndpointConnectionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the private endpoint connection
        /// </summary>
        [Input("privateEndpointConnectionName", required: true)]
        public Input<string> PrivateEndpointConnectionName { get; set; } = null!;

        /// <summary>
        /// The properties of a private endpoint connection
        /// </summary>
        [Input("properties", required: true)]
        public Input<Inputs.PrivateEndpointConnectionPropertiesArgs> Properties { get; set; } = null!;

        /// <summary>
        /// The name of the resource group that contains the IoT hub.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the IoT hub.
        /// </summary>
        [Input("resourceName", required: true)]
        public Input<string> ResourceName { get; set; } = null!;

        public PrivateEndpointConnectionArgs()
        {
        }
    }
}
