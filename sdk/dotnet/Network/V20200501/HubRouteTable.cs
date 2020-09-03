// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20200501
{
    /// <summary>
    /// RouteTable resource in a virtual hub.
    /// 
    /// ## Example Usage
    /// ### RouteTablePut
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var hubRouteTable = new AzureRM.Network.V20200501.HubRouteTable("hubRouteTable", new AzureRM.Network.V20200501.HubRouteTableArgs
    ///         {
    ///             Labels = 
    ///             {
    ///                 "label1",
    ///                 "label2",
    ///             },
    ///             ResourceGroupName = "rg1",
    ///             RouteTableName = "hubRouteTable1",
    ///             Routes = 
    ///             {
    ///                 new AzureRM.Network.V20200501.Inputs.HubRouteArgs
    ///                 {
    ///                     DestinationType = "CIDR",
    ///                     Destinations = 
    ///                     {
    ///                         "10.0.0.0/8",
    ///                         "20.0.0.0/8",
    ///                         "30.0.0.0/8",
    ///                     },
    ///                     Name = "route1",
    ///                     NextHop = "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azureFirewall1",
    ///                     NextHopType = "ResourceId",
    ///                 },
    ///             },
    ///             VirtualHubName = "virtualHub1",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class HubRouteTable : Pulumi.CustomResource
    {
        /// <summary>
        /// List of all connections associated with this route table.
        /// </summary>
        [Output("associatedConnections")]
        public Output<ImmutableArray<string>> AssociatedConnections { get; private set; } = null!;

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// List of labels associated with this route table.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableArray<string>> Labels { get; private set; } = null!;

        /// <summary>
        /// The name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// List of all connections that advertise to this route table.
        /// </summary>
        [Output("propagatingConnections")]
        public Output<ImmutableArray<string>> PropagatingConnections { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the RouteTable resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// List of all routes.
        /// </summary>
        [Output("routes")]
        public Output<ImmutableArray<Outputs.HubRouteResponseResult>> Routes { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a HubRouteTable resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HubRouteTable(string name, HubRouteTableArgs args, CustomResourceOptions? options = null)
            : base("azurerm:network/v20200501:HubRouteTable", name, args ?? new HubRouteTableArgs(), MakeResourceOptions(options, ""))
        {
        }

        private HubRouteTable(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:network/v20200501:HubRouteTable", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:network/latest:HubRouteTable"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200401:HubRouteTable"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200601:HubRouteTable"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing HubRouteTable resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HubRouteTable Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new HubRouteTable(name, id, options);
        }
    }

    public sealed class HubRouteTableArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("labels")]
        private InputList<string>? _labels;

        /// <summary>
        /// List of labels associated with this route table.
        /// </summary>
        public InputList<string> Labels
        {
            get => _labels ?? (_labels = new InputList<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The resource group name of the VirtualHub.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the RouteTable.
        /// </summary>
        [Input("routeTableName", required: true)]
        public Input<string> RouteTableName { get; set; } = null!;

        [Input("routes")]
        private InputList<Inputs.HubRouteArgs>? _routes;

        /// <summary>
        /// List of all routes.
        /// </summary>
        public InputList<Inputs.HubRouteArgs> Routes
        {
            get => _routes ?? (_routes = new InputList<Inputs.HubRouteArgs>());
            set => _routes = value;
        }

        /// <summary>
        /// The name of the VirtualHub.
        /// </summary>
        [Input("virtualHubName", required: true)]
        public Input<string> VirtualHubName { get; set; } = null!;

        public HubRouteTableArgs()
        {
        }
    }
}
