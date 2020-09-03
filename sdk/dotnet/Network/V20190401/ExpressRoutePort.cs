// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20190401
{
    /// <summary>
    /// ExpressRoutePort resource definition.
    /// 
    /// ## Example Usage
    /// ### ExpressRoutePortCreate
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var expressRoutePort = new AzureRM.Network.V20190401.ExpressRoutePort("expressRoutePort", new AzureRM.Network.V20190401.ExpressRoutePortArgs
    ///         {
    ///             BandwidthInGbps = 100,
    ///             Encapsulation = "QinQ",
    ///             ExpressRoutePortName = "portName",
    ///             Location = "westus",
    ///             PeeringLocation = "peeringLocationName",
    ///             ResourceGroupName = "rg1",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// ### ExpressRoutePortUpdateLink
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var expressRoutePort = new AzureRM.Network.V20190401.ExpressRoutePort("expressRoutePort", new AzureRM.Network.V20190401.ExpressRoutePortArgs
    ///         {
    ///             BandwidthInGbps = 100,
    ///             Encapsulation = "QinQ",
    ///             ExpressRoutePortName = "portName",
    ///             Links = 
    ///             {
    ///                 new AzureRM.Network.V20190401.Inputs.ExpressRouteLinkArgs
    ///                 {
    ///                     Name = "link1",
    ///                 },
    ///             },
    ///             Location = "westus",
    ///             PeeringLocation = "peeringLocationName",
    ///             ResourceGroupName = "rg1",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class ExpressRoutePort : Pulumi.CustomResource
    {
        /// <summary>
        /// Date of the physical port allocation to be used in Letter of Authorization.
        /// </summary>
        [Output("allocationDate")]
        public Output<string> AllocationDate { get; private set; } = null!;

        /// <summary>
        /// Bandwidth of procured ports in Gbps.
        /// </summary>
        [Output("bandwidthInGbps")]
        public Output<int?> BandwidthInGbps { get; private set; } = null!;

        /// <summary>
        /// Reference the ExpressRoute circuit(s) that are provisioned on this ExpressRoutePort resource.
        /// </summary>
        [Output("circuits")]
        public Output<ImmutableArray<Outputs.SubResourceResponseResult>> Circuits { get; private set; } = null!;

        /// <summary>
        /// Encapsulation method on physical ports.
        /// </summary>
        [Output("encapsulation")]
        public Output<string?> Encapsulation { get; private set; } = null!;

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// Ether type of the physical port.
        /// </summary>
        [Output("etherType")]
        public Output<string> EtherType { get; private set; } = null!;

        /// <summary>
        /// The set of physical links of the ExpressRoutePort resource.
        /// </summary>
        [Output("links")]
        public Output<ImmutableArray<Outputs.ExpressRouteLinkResponseResult>> Links { get; private set; } = null!;

        /// <summary>
        /// Resource location.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// Maximum transmission unit of the physical port pair(s).
        /// </summary>
        [Output("mtu")]
        public Output<string> Mtu { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the peering location that the ExpressRoutePort is mapped to physically.
        /// </summary>
        [Output("peeringLocation")]
        public Output<string?> PeeringLocation { get; private set; } = null!;

        /// <summary>
        /// Aggregate Gbps of associated circuit bandwidths.
        /// </summary>
        [Output("provisionedBandwidthInGbps")]
        public Output<double> ProvisionedBandwidthInGbps { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the ExpressRoutePort resource. Possible values are: 'Succeeded', 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The resource GUID property of the ExpressRoutePort resource.
        /// </summary>
        [Output("resourceGuid")]
        public Output<string?> ResourceGuid { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ExpressRoutePort resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ExpressRoutePort(string name, ExpressRoutePortArgs args, CustomResourceOptions? options = null)
            : base("azurerm:network/v20190401:ExpressRoutePort", name, args ?? new ExpressRoutePortArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ExpressRoutePort(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:network/v20190401:ExpressRoutePort", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:network/latest:ExpressRoutePort"},
                    new Pulumi.Alias { Type = "azurerm:network/v20180801:ExpressRoutePort"},
                    new Pulumi.Alias { Type = "azurerm:network/v20181001:ExpressRoutePort"},
                    new Pulumi.Alias { Type = "azurerm:network/v20181101:ExpressRoutePort"},
                    new Pulumi.Alias { Type = "azurerm:network/v20181201:ExpressRoutePort"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190201:ExpressRoutePort"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190601:ExpressRoutePort"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190701:ExpressRoutePort"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190801:ExpressRoutePort"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190901:ExpressRoutePort"},
                    new Pulumi.Alias { Type = "azurerm:network/v20191101:ExpressRoutePort"},
                    new Pulumi.Alias { Type = "azurerm:network/v20191201:ExpressRoutePort"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200301:ExpressRoutePort"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200401:ExpressRoutePort"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200501:ExpressRoutePort"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200601:ExpressRoutePort"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ExpressRoutePort resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ExpressRoutePort Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ExpressRoutePort(name, id, options);
        }
    }

    public sealed class ExpressRoutePortArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Bandwidth of procured ports in Gbps.
        /// </summary>
        [Input("bandwidthInGbps")]
        public Input<int>? BandwidthInGbps { get; set; }

        /// <summary>
        /// Encapsulation method on physical ports.
        /// </summary>
        [Input("encapsulation")]
        public Input<string>? Encapsulation { get; set; }

        /// <summary>
        /// The name of the ExpressRoutePort resource.
        /// </summary>
        [Input("expressRoutePortName", required: true)]
        public Input<string> ExpressRoutePortName { get; set; } = null!;

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("links")]
        private InputList<Inputs.ExpressRouteLinkArgs>? _links;

        /// <summary>
        /// The set of physical links of the ExpressRoutePort resource.
        /// </summary>
        public InputList<Inputs.ExpressRouteLinkArgs> Links
        {
            get => _links ?? (_links = new InputList<Inputs.ExpressRouteLinkArgs>());
            set => _links = value;
        }

        /// <summary>
        /// Resource location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the peering location that the ExpressRoutePort is mapped to physically.
        /// </summary>
        [Input("peeringLocation")]
        public Input<string>? PeeringLocation { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The resource GUID property of the ExpressRoutePort resource.
        /// </summary>
        [Input("resourceGuid")]
        public Input<string>? ResourceGuid { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ExpressRoutePortArgs()
        {
        }
    }
}
