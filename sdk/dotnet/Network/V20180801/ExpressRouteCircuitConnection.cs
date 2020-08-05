// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20180801
{
    /// <summary>
    /// Express Route Circuit Connection in an ExpressRouteCircuitPeering resource.
    /// </summary>
    public partial class ExpressRouteCircuitConnection : Pulumi.CustomResource
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// Gets name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        [Output("properties")]
        public Output<Outputs.ExpressRouteCircuitConnectionPropertiesFormatResponseResult> Properties { get; private set; } = null!;


        /// <summary>
        /// Create a ExpressRouteCircuitConnection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ExpressRouteCircuitConnection(string name, ExpressRouteCircuitConnectionArgs args, CustomResourceOptions? options = null)
            : base("azurerm:network/v20180801:ExpressRouteCircuitConnection", name, args ?? new ExpressRouteCircuitConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ExpressRouteCircuitConnection(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:network/v20180801:ExpressRouteCircuitConnection", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ExpressRouteCircuitConnection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ExpressRouteCircuitConnection Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ExpressRouteCircuitConnection(name, id, options);
        }
    }

    public sealed class ExpressRouteCircuitConnectionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// /29 IP address space to carve out Customer addresses for tunnels.
        /// </summary>
        [Input("addressPrefix")]
        public Input<string>? AddressPrefix { get; set; }

        /// <summary>
        /// The authorization key.
        /// </summary>
        [Input("authorizationKey")]
        public Input<string>? AuthorizationKey { get; set; }

        /// <summary>
        /// The name of the express route circuit.
        /// </summary>
        [Input("circuitName", required: true)]
        public Input<string> CircuitName { get; set; } = null!;

        /// <summary>
        /// Reference to Express Route Circuit Private Peering Resource of the circuit initiating connection.
        /// </summary>
        [Input("expressRouteCircuitPeering")]
        public Input<Inputs.SubResourceArgs>? ExpressRouteCircuitPeering { get; set; }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The name of the express route circuit connection.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Reference to Express Route Circuit Private Peering Resource of the peered circuit.
        /// </summary>
        [Input("peerExpressRouteCircuitPeering")]
        public Input<Inputs.SubResourceArgs>? PeerExpressRouteCircuitPeering { get; set; }

        /// <summary>
        /// The name of the peering.
        /// </summary>
        [Input("peeringName", required: true)]
        public Input<string> PeeringName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public ExpressRouteCircuitConnectionArgs()
        {
        }
    }
}
