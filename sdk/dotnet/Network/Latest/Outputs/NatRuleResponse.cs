// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.Latest.Outputs
{

    [OutputType]
    public sealed class NatRuleResponse
    {
        /// <summary>
        /// Description of the rule.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// List of destination IP addresses or Service Tags.
        /// </summary>
        public readonly ImmutableArray<string> DestinationAddresses;
        /// <summary>
        /// List of destination ports.
        /// </summary>
        public readonly ImmutableArray<string> DestinationPorts;
        /// <summary>
        /// Array of FirewallPolicyRuleNetworkProtocols.
        /// </summary>
        public readonly ImmutableArray<string> IpProtocols;
        /// <summary>
        /// Name of the rule.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Rule Type.
        /// </summary>
        public readonly string RuleType;
        /// <summary>
        /// List of source IP addresses for this rule.
        /// </summary>
        public readonly ImmutableArray<string> SourceAddresses;
        /// <summary>
        /// List of source IpGroups for this rule.
        /// </summary>
        public readonly ImmutableArray<string> SourceIpGroups;
        /// <summary>
        /// The translated address for this NAT rule.
        /// </summary>
        public readonly string? TranslatedAddress;
        /// <summary>
        /// The translated FQDN for this NAT rule.
        /// </summary>
        public readonly string? TranslatedFqdn;
        /// <summary>
        /// The translated port for this NAT rule.
        /// </summary>
        public readonly string? TranslatedPort;

        [OutputConstructor]
        private NatRuleResponse(
            string? description,

            ImmutableArray<string> destinationAddresses,

            ImmutableArray<string> destinationPorts,

            ImmutableArray<string> ipProtocols,

            string? name,

            string ruleType,

            ImmutableArray<string> sourceAddresses,

            ImmutableArray<string> sourceIpGroups,

            string? translatedAddress,

            string? translatedFqdn,

            string? translatedPort)
        {
            Description = description;
            DestinationAddresses = destinationAddresses;
            DestinationPorts = destinationPorts;
            IpProtocols = ipProtocols;
            Name = name;
            RuleType = ruleType;
            SourceAddresses = sourceAddresses;
            SourceIpGroups = sourceIpGroups;
            TranslatedAddress = translatedAddress;
            TranslatedFqdn = translatedFqdn;
            TranslatedPort = translatedPort;
        }
    }
}
