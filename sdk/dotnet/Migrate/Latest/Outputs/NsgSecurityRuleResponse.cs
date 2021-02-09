// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Migrate.Latest.Outputs
{

    [OutputType]
    public sealed class NsgSecurityRuleResponse
    {
        /// <summary>
        /// Gets or sets whether network traffic is allowed or denied.
        /// Possible values are “Allow” and “Deny”.
        /// </summary>
        public readonly string? Access;
        /// <summary>
        /// Gets or sets a description for this rule. Restricted to 140 chars.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Gets or sets destination address prefix. CIDR or source IP range.
        ///  A “*” can also be used to match all source IPs. Default tags such
        /// as ‘VirtualNetwork’, ‘AzureLoadBalancer’ and ‘Internet’ can also be used.
        /// </summary>
        public readonly string? DestinationAddressPrefix;
        /// <summary>
        /// Gets or sets Destination Port or Range. Integer or range between
        /// 0 and 65535. A “*” can also be used to match all ports.
        /// </summary>
        public readonly string? DestinationPortRange;
        /// <summary>
        /// Gets or sets the direction of the rule.InBound or Outbound. The
        /// direction specifies if rule will be evaluated on incoming or outgoing traffic.
        /// </summary>
        public readonly string? Direction;
        /// <summary>
        /// Gets or sets the Security rule name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Gets or sets the priority of the rule. The value can be between
        /// 100 and 4096. The priority number must be unique for each rule in the collection.
        /// The lower the priority number, the higher the priority of the rule.
        /// </summary>
        public readonly int? Priority;
        /// <summary>
        /// Gets or sets Network protocol this rule applies to. Can be Tcp, Udp or All(*).
        /// </summary>
        public readonly string? Protocol;
        /// <summary>
        /// Gets or sets source address prefix. CIDR or source IP range. A
        /// “*” can also be used to match all source IPs.  Default tags such as ‘VirtualNetwork’,
        /// ‘AzureLoadBalancer’ and ‘Internet’ can also be used. If this is an ingress
        /// rule, specifies where network traffic originates from.
        /// </summary>
        public readonly string? SourceAddressPrefix;
        /// <summary>
        /// Gets or sets Source Port or Range. Integer or range between 0 and
        /// 65535. A “*” can also be used to match all ports.
        /// </summary>
        public readonly string? SourcePortRange;

        [OutputConstructor]
        private NsgSecurityRuleResponse(
            string? access,

            string? description,

            string? destinationAddressPrefix,

            string? destinationPortRange,

            string? direction,

            string? name,

            int? priority,

            string? protocol,

            string? sourceAddressPrefix,

            string? sourcePortRange)
        {
            Access = access;
            Description = description;
            DestinationAddressPrefix = destinationAddressPrefix;
            DestinationPortRange = destinationPortRange;
            Direction = direction;
            Name = name;
            Priority = priority;
            Protocol = protocol;
            SourceAddressPrefix = sourceAddressPrefix;
            SourcePortRange = sourcePortRange;
        }
    }
}
