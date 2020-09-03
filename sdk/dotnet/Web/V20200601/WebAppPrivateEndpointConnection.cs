// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Web.V20200601
{
    /// <summary>
    /// Private Endpoint Connection ARM resource.
    /// 
    /// ## Example Usage
    /// ### Approves or rejects a private endpoint connection for a site.
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var webAppPrivateEndpointConnection = new AzureRM.Web.V20200601.WebAppPrivateEndpointConnection("webAppPrivateEndpointConnection", new AzureRM.Web.V20200601.WebAppPrivateEndpointConnectionArgs
    ///         {
    ///             Name = "testSite",
    ///             PrivateEndpointConnectionName = "connection",
    ///             PrivateLinkServiceConnectionState = new AzureRM.Web.V20200601.Inputs.PrivateLinkConnectionStateArgs
    ///             {
    ///                 ActionsRequired = "",
    ///                 Description = "Approved by admin.",
    ///                 Status = "Approved",
    ///             },
    ///             ResourceGroupName = "rg",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class WebAppPrivateEndpointConnection : Pulumi.CustomResource
    {
        /// <summary>
        /// Kind of resource.
        /// </summary>
        [Output("kind")]
        public Output<string?> Kind { get; private set; } = null!;

        /// <summary>
        /// Resource Name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// PrivateEndpoint of a remote private endpoint connection
        /// </summary>
        [Output("privateEndpoint")]
        public Output<Outputs.ArmIdWrapperResponseResult?> PrivateEndpoint { get; private set; } = null!;

        /// <summary>
        /// The state of a private link connection
        /// </summary>
        [Output("privateLinkServiceConnectionState")]
        public Output<Outputs.PrivateLinkConnectionStateResponseResult?> PrivateLinkServiceConnectionState { get; private set; } = null!;

        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a WebAppPrivateEndpointConnection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WebAppPrivateEndpointConnection(string name, WebAppPrivateEndpointConnectionArgs args, CustomResourceOptions? options = null)
            : base("azurerm:web/v20200601:WebAppPrivateEndpointConnection", name, args ?? new WebAppPrivateEndpointConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WebAppPrivateEndpointConnection(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:web/v20200601:WebAppPrivateEndpointConnection", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:web/latest:WebAppPrivateEndpointConnection"},
                    new Pulumi.Alias { Type = "azurerm:web/v20190801:WebAppPrivateEndpointConnection"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing WebAppPrivateEndpointConnection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WebAppPrivateEndpointConnection Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new WebAppPrivateEndpointConnection(name, id, options);
        }
    }

    public sealed class WebAppPrivateEndpointConnectionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Kind of resource.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Name of the site.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("privateEndpointConnectionName", required: true)]
        public Input<string> PrivateEndpointConnectionName { get; set; } = null!;

        /// <summary>
        /// The state of a private link connection
        /// </summary>
        [Input("privateLinkServiceConnectionState")]
        public Input<Inputs.PrivateLinkConnectionStateArgs>? PrivateLinkServiceConnectionState { get; set; }

        /// <summary>
        /// Name of the resource group to which the resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public WebAppPrivateEndpointConnectionArgs()
        {
        }
    }
}
