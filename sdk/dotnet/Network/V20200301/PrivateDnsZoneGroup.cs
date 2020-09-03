// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20200301
{
    /// <summary>
    /// Private dns zone group resource.
    /// 
    /// ## Example Usage
    /// ### Create private dns zone group
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var privateDnsZoneGroup = new AzureRM.Network.V20200301.PrivateDnsZoneGroup("privateDnsZoneGroup", new AzureRM.Network.V20200301.PrivateDnsZoneGroupArgs
    ///         {
    ///             PrivateDnsZoneConfigs = 
    ///             {
    ///                 ,
    ///             },
    ///             PrivateDnsZoneGroupName = "testPdnsgroup",
    ///             PrivateEndpointName = "testPe",
    ///             ResourceGroupName = "rg1",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class PrivateDnsZoneGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// Name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// A collection of private dns zone configurations of the private dns zone group.
        /// </summary>
        [Output("privateDnsZoneConfigs")]
        public Output<ImmutableArray<Outputs.PrivateDnsZoneConfigResponseResult>> PrivateDnsZoneConfigs { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the private dns zone group resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;


        /// <summary>
        /// Create a PrivateDnsZoneGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PrivateDnsZoneGroup(string name, PrivateDnsZoneGroupArgs args, CustomResourceOptions? options = null)
            : base("azurerm:network/v20200301:PrivateDnsZoneGroup", name, args ?? new PrivateDnsZoneGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PrivateDnsZoneGroup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:network/v20200301:PrivateDnsZoneGroup", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:network/latest:PrivateDnsZoneGroup"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200401:PrivateDnsZoneGroup"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200501:PrivateDnsZoneGroup"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200601:PrivateDnsZoneGroup"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PrivateDnsZoneGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PrivateDnsZoneGroup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PrivateDnsZoneGroup(name, id, options);
        }
    }

    public sealed class PrivateDnsZoneGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("privateDnsZoneConfigs")]
        private InputList<Inputs.PrivateDnsZoneConfigArgs>? _privateDnsZoneConfigs;

        /// <summary>
        /// A collection of private dns zone configurations of the private dns zone group.
        /// </summary>
        public InputList<Inputs.PrivateDnsZoneConfigArgs> PrivateDnsZoneConfigs
        {
            get => _privateDnsZoneConfigs ?? (_privateDnsZoneConfigs = new InputList<Inputs.PrivateDnsZoneConfigArgs>());
            set => _privateDnsZoneConfigs = value;
        }

        /// <summary>
        /// The name of the private dns zone group.
        /// </summary>
        [Input("privateDnsZoneGroupName", required: true)]
        public Input<string> PrivateDnsZoneGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the private endpoint.
        /// </summary>
        [Input("privateEndpointName", required: true)]
        public Input<string> PrivateEndpointName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public PrivateDnsZoneGroupArgs()
        {
        }
    }
}
