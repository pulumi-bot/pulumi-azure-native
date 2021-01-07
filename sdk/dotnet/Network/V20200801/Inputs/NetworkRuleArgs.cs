// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20200801.Inputs
{

    /// <summary>
    /// Rule of type network.
    /// </summary>
    public sealed class NetworkRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the rule.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("destinationAddresses")]
        private InputList<string>? _destinationAddresses;

        /// <summary>
        /// List of destination IP addresses or Service Tags.
        /// </summary>
        public InputList<string> DestinationAddresses
        {
            get => _destinationAddresses ?? (_destinationAddresses = new InputList<string>());
            set => _destinationAddresses = value;
        }

        [Input("destinationFqdns")]
        private InputList<string>? _destinationFqdns;

        /// <summary>
        /// List of destination FQDNs.
        /// </summary>
        public InputList<string> DestinationFqdns
        {
            get => _destinationFqdns ?? (_destinationFqdns = new InputList<string>());
            set => _destinationFqdns = value;
        }

        [Input("destinationIpGroups")]
        private InputList<string>? _destinationIpGroups;

        /// <summary>
        /// List of destination IpGroups for this rule.
        /// </summary>
        public InputList<string> DestinationIpGroups
        {
            get => _destinationIpGroups ?? (_destinationIpGroups = new InputList<string>());
            set => _destinationIpGroups = value;
        }

        [Input("destinationPorts")]
        private InputList<string>? _destinationPorts;

        /// <summary>
        /// List of destination ports.
        /// </summary>
        public InputList<string> DestinationPorts
        {
            get => _destinationPorts ?? (_destinationPorts = new InputList<string>());
            set => _destinationPorts = value;
        }

        [Input("ipProtocols")]
        private InputList<Union<string, Pulumi.AzureNextGen.Network.V20200801.FirewallPolicyRuleNetworkProtocol>>? _ipProtocols;

        /// <summary>
        /// Array of FirewallPolicyRuleNetworkProtocols.
        /// </summary>
        public InputList<Union<string, Pulumi.AzureNextGen.Network.V20200801.FirewallPolicyRuleNetworkProtocol>> IpProtocols
        {
            get => _ipProtocols ?? (_ipProtocols = new InputList<Union<string, Pulumi.AzureNextGen.Network.V20200801.FirewallPolicyRuleNetworkProtocol>>());
            set => _ipProtocols = value;
        }

        /// <summary>
        /// Name of the rule.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Rule Type.
        /// Expected value is 'NetworkRule'.
        /// </summary>
        [Input("ruleType", required: true)]
        public Input<string> RuleType { get; set; } = null!;

        [Input("sourceAddresses")]
        private InputList<string>? _sourceAddresses;

        /// <summary>
        /// List of source IP addresses for this rule.
        /// </summary>
        public InputList<string> SourceAddresses
        {
            get => _sourceAddresses ?? (_sourceAddresses = new InputList<string>());
            set => _sourceAddresses = value;
        }

        [Input("sourceIpGroups")]
        private InputList<string>? _sourceIpGroups;

        /// <summary>
        /// List of source IpGroups for this rule.
        /// </summary>
        public InputList<string> SourceIpGroups
        {
            get => _sourceIpGroups ?? (_sourceIpGroups = new InputList<string>());
            set => _sourceIpGroups = value;
        }

        public NetworkRuleArgs()
        {
        }
    }
}
