// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20181101.Inputs
{

    /// <summary>
    /// Security rule resource.
    /// </summary>
    public sealed class SecurityRulePropertiesFormatResponseArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The network traffic is allowed or denied. Possible values are: 'Allow' and 'Deny'.
        /// </summary>
        [Input("access", required: true)]
        public string Access { get; set; } = null!;

        /// <summary>
        /// A description for this rule. Restricted to 140 chars.
        /// </summary>
        [Input("description")]
        public string? Description { get; set; }

        /// <summary>
        /// The destination address prefix. CIDR or destination IP range. Asterisks '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used.
        /// </summary>
        [Input("destinationAddressPrefix")]
        public string? DestinationAddressPrefix { get; set; }

        [Input("destinationAddressPrefixes")]
        private List<string>? _destinationAddressPrefixes;

        /// <summary>
        /// The destination address prefixes. CIDR or destination IP ranges.
        /// </summary>
        public List<string> DestinationAddressPrefixes
        {
            get => _destinationAddressPrefixes ?? (_destinationAddressPrefixes = new List<string>());
            set => _destinationAddressPrefixes = value;
        }

        [Input("destinationApplicationSecurityGroups")]
        private List<Inputs.ApplicationSecurityGroupResponseArgs>? _destinationApplicationSecurityGroups;

        /// <summary>
        /// The application security group specified as destination.
        /// </summary>
        public List<Inputs.ApplicationSecurityGroupResponseArgs> DestinationApplicationSecurityGroups
        {
            get => _destinationApplicationSecurityGroups ?? (_destinationApplicationSecurityGroups = new List<Inputs.ApplicationSecurityGroupResponseArgs>());
            set => _destinationApplicationSecurityGroups = value;
        }

        /// <summary>
        /// The destination port or range. Integer or range between 0 and 65535. Asterisks '*' can also be used to match all ports.
        /// </summary>
        [Input("destinationPortRange")]
        public string? DestinationPortRange { get; set; }

        [Input("destinationPortRanges")]
        private List<string>? _destinationPortRanges;

        /// <summary>
        /// The destination port ranges.
        /// </summary>
        public List<string> DestinationPortRanges
        {
            get => _destinationPortRanges ?? (_destinationPortRanges = new List<string>());
            set => _destinationPortRanges = value;
        }

        /// <summary>
        /// The direction of the rule. The direction specifies if rule will be evaluated on incoming or outgoing traffic. Possible values are: 'Inbound' and 'Outbound'.
        /// </summary>
        [Input("direction", required: true)]
        public string Direction { get; set; } = null!;

        /// <summary>
        /// The priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
        /// </summary>
        [Input("priority")]
        public int? Priority { get; set; }

        /// <summary>
        /// Network protocol this rule applies to. Possible values are 'Tcp', 'Udp', and '*'.
        /// </summary>
        [Input("protocol", required: true)]
        public string Protocol { get; set; } = null!;

        /// <summary>
        /// The provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        [Input("provisioningState")]
        public string? ProvisioningState { get; set; }

        /// <summary>
        /// The CIDR or source IP range. Asterisks '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used. If this is an ingress rule, specifies where network traffic originates from. 
        /// </summary>
        [Input("sourceAddressPrefix")]
        public string? SourceAddressPrefix { get; set; }

        [Input("sourceAddressPrefixes")]
        private List<string>? _sourceAddressPrefixes;

        /// <summary>
        /// The CIDR or source IP ranges.
        /// </summary>
        public List<string> SourceAddressPrefixes
        {
            get => _sourceAddressPrefixes ?? (_sourceAddressPrefixes = new List<string>());
            set => _sourceAddressPrefixes = value;
        }

        [Input("sourceApplicationSecurityGroups")]
        private List<Inputs.ApplicationSecurityGroupResponseArgs>? _sourceApplicationSecurityGroups;

        /// <summary>
        /// The application security group specified as source.
        /// </summary>
        public List<Inputs.ApplicationSecurityGroupResponseArgs> SourceApplicationSecurityGroups
        {
            get => _sourceApplicationSecurityGroups ?? (_sourceApplicationSecurityGroups = new List<Inputs.ApplicationSecurityGroupResponseArgs>());
            set => _sourceApplicationSecurityGroups = value;
        }

        /// <summary>
        /// The source port or range. Integer or range between 0 and 65535. Asterisks '*' can also be used to match all ports.
        /// </summary>
        [Input("sourcePortRange")]
        public string? SourcePortRange { get; set; }

        [Input("sourcePortRanges")]
        private List<string>? _sourcePortRanges;

        /// <summary>
        /// The source port ranges.
        /// </summary>
        public List<string> SourcePortRanges
        {
            get => _sourcePortRanges ?? (_sourcePortRanges = new List<string>());
            set => _sourcePortRanges = value;
        }

        public SecurityRulePropertiesFormatResponseArgs()
        {
        }
    }
}
