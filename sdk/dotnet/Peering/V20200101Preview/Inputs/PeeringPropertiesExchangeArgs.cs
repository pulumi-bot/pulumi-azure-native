// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Peering.V20200101Preview.Inputs
{

    /// <summary>
    /// The properties that define an exchange peering.
    /// </summary>
    public sealed class PeeringPropertiesExchangeArgs : Pulumi.ResourceArgs
    {
        [Input("connections")]
        private InputList<Inputs.ExchangeConnectionArgs>? _connections;

        /// <summary>
        /// The set of connections that constitute an exchange peering.
        /// </summary>
        public InputList<Inputs.ExchangeConnectionArgs> Connections
        {
            get => _connections ?? (_connections = new InputList<Inputs.ExchangeConnectionArgs>());
            set => _connections = value;
        }

        /// <summary>
        /// The reference of the peer ASN.
        /// </summary>
        [Input("peerAsn")]
        public Input<Inputs.SubResourceArgs>? PeerAsn { get; set; }

        public PeeringPropertiesExchangeArgs()
        {
        }
    }
}
