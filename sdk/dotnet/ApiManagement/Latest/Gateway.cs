// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ApiManagement.Latest
{
    /// <summary>
    /// Gateway details.
    /// 
    /// ## Example Usage
    /// ### ApiManagementCreateGateway
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var gateway = new AzureRM.ApiManagement.Latest.Gateway("gateway", new AzureRM.ApiManagement.Latest.GatewayArgs
    ///         {
    ///             Description = "my gateway 1",
    ///             GatewayId = "gw1",
    ///             LocationData = new AzureRM.ApiManagement.Latest.Inputs.ResourceLocationDataContractArgs
    ///             {
    ///                 Name = "my location",
    ///             },
    ///             ResourceGroupName = "rg1",
    ///             ServiceName = "apimService1",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class Gateway : Pulumi.CustomResource
    {
        /// <summary>
        /// Gateway description
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Gateway location.
        /// </summary>
        [Output("locationData")]
        public Output<Outputs.ResourceLocationDataContractResponseResult?> LocationData { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Resource type for API Management resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Gateway resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Gateway(string name, GatewayArgs args, CustomResourceOptions? options = null)
            : base("azurerm:apimanagement/latest:Gateway", name, args ?? new GatewayArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Gateway(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:apimanagement/latest:Gateway", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:apimanagement/v20191201:Gateway"},
                    new Pulumi.Alias { Type = "azurerm:apimanagement/v20191201preview:Gateway"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Gateway resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Gateway Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Gateway(name, id, options);
        }
    }

    public sealed class GatewayArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gateway description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Gateway entity identifier. Must be unique in the current API Management service instance. Must not have value 'managed'
        /// </summary>
        [Input("gatewayId", required: true)]
        public Input<string> GatewayId { get; set; } = null!;

        /// <summary>
        /// Gateway location.
        /// </summary>
        [Input("locationData")]
        public Input<Inputs.ResourceLocationDataContractArgs>? LocationData { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the API Management service.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GatewayArgs()
        {
        }
    }
}
