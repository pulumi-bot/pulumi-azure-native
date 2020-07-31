// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20191201.Inputs
{

    /// <summary>
    /// VirtualNetworkGatewayConnection properties.
    /// </summary>
    public sealed class VirtualNetworkGatewayConnectionPropertiesFormatArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The authorizationKey.
        /// </summary>
        [Input("authorizationKey")]
        public Input<string>? AuthorizationKey { get; set; }

        /// <summary>
        /// Connection protocol used for this connection.
        /// </summary>
        [Input("connectionProtocol")]
        public Input<string>? ConnectionProtocol { get; set; }

        /// <summary>
        /// Gateway connection type.
        /// </summary>
        [Input("connectionType", required: true)]
        public Input<string> ConnectionType { get; set; } = null!;

        /// <summary>
        /// EnableBgp flag.
        /// </summary>
        [Input("enableBgp")]
        public Input<bool>? EnableBgp { get; set; }

        /// <summary>
        /// Bypass ExpressRoute Gateway for data forwarding.
        /// </summary>
        [Input("expressRouteGatewayBypass")]
        public Input<bool>? ExpressRouteGatewayBypass { get; set; }

        [Input("ipsecPolicies")]
        private InputList<Inputs.IpsecPolicyArgs>? _ipsecPolicies;

        /// <summary>
        /// The IPSec Policies to be considered by this connection.
        /// </summary>
        public InputList<Inputs.IpsecPolicyArgs> IpsecPolicies
        {
            get => _ipsecPolicies ?? (_ipsecPolicies = new InputList<Inputs.IpsecPolicyArgs>());
            set => _ipsecPolicies = value;
        }

        /// <summary>
        /// The reference to local network gateway resource.
        /// </summary>
        [Input("localNetworkGateway2")]
        public Input<Inputs.LocalNetworkGatewayArgs>? LocalNetworkGateway2 { get; set; }

        /// <summary>
        /// The reference to peerings resource.
        /// </summary>
        [Input("peer")]
        public Input<Inputs.SubResourceArgs>? Peer { get; set; }

        /// <summary>
        /// The routing weight.
        /// </summary>
        [Input("routingWeight")]
        public Input<int>? RoutingWeight { get; set; }

        /// <summary>
        /// The IPSec shared key.
        /// </summary>
        [Input("sharedKey")]
        public Input<string>? SharedKey { get; set; }

        [Input("trafficSelectorPolicies")]
        private InputList<Inputs.TrafficSelectorPolicyArgs>? _trafficSelectorPolicies;

        /// <summary>
        /// The Traffic Selector Policies to be considered by this connection.
        /// </summary>
        public InputList<Inputs.TrafficSelectorPolicyArgs> TrafficSelectorPolicies
        {
            get => _trafficSelectorPolicies ?? (_trafficSelectorPolicies = new InputList<Inputs.TrafficSelectorPolicyArgs>());
            set => _trafficSelectorPolicies = value;
        }

        /// <summary>
        /// Use private local Azure IP for the connection.
        /// </summary>
        [Input("useLocalAzureIpAddress")]
        public Input<bool>? UseLocalAzureIpAddress { get; set; }

        /// <summary>
        /// Enable policy-based traffic selectors.
        /// </summary>
        [Input("usePolicyBasedTrafficSelectors")]
        public Input<bool>? UsePolicyBasedTrafficSelectors { get; set; }

        /// <summary>
        /// The reference to virtual network gateway resource.
        /// </summary>
        [Input("virtualNetworkGateway1", required: true)]
        public Input<Inputs.VirtualNetworkGatewayArgs> VirtualNetworkGateway1 { get; set; } = null!;

        /// <summary>
        /// The reference to virtual network gateway resource.
        /// </summary>
        [Input("virtualNetworkGateway2")]
        public Input<Inputs.VirtualNetworkGatewayArgs>? VirtualNetworkGateway2 { get; set; }

        public VirtualNetworkGatewayConnectionPropertiesFormatArgs()
        {
        }
    }
}