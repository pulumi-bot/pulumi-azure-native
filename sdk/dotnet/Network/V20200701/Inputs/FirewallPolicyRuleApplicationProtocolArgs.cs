// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20200701.Inputs
{

    /// <summary>
    /// Properties of the application rule protocol.
    /// </summary>
    public sealed class FirewallPolicyRuleApplicationProtocolArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Port number for the protocol, cannot be greater than 64000.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Protocol type.
        /// </summary>
        [Input("protocolType")]
        public InputUnion<string, Pulumi.AzureNextGen.Network.V20200701.FirewallPolicyRuleApplicationProtocolType>? ProtocolType { get; set; }

        public FirewallPolicyRuleApplicationProtocolArgs()
        {
        }
    }
}
