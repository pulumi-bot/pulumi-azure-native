// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.CognitiveServices.V20170418.Inputs
{

    /// <summary>
    /// A set of rules governing the network accessibility.
    /// </summary>
    public sealed class NetworkRuleSetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The default action when no rule from ipRules and from virtualNetworkRules match. This is only used after the bypass property has been evaluated.
        /// </summary>
        [Input("defaultAction")]
        public InputUnion<string, Pulumi.AzureNextGen.CognitiveServices.V20170418.NetworkRuleAction>? DefaultAction { get; set; }

        [Input("ipRules")]
        private InputList<Inputs.IpRuleArgs>? _ipRules;

        /// <summary>
        /// The list of IP address rules.
        /// </summary>
        public InputList<Inputs.IpRuleArgs> IpRules
        {
            get => _ipRules ?? (_ipRules = new InputList<Inputs.IpRuleArgs>());
            set => _ipRules = value;
        }

        [Input("virtualNetworkRules")]
        private InputList<Inputs.VirtualNetworkRuleArgs>? _virtualNetworkRules;

        /// <summary>
        /// The list of virtual network rules.
        /// </summary>
        public InputList<Inputs.VirtualNetworkRuleArgs> VirtualNetworkRules
        {
            get => _virtualNetworkRules ?? (_virtualNetworkRules = new InputList<Inputs.VirtualNetworkRuleArgs>());
            set => _virtualNetworkRules = value;
        }

        public NetworkRuleSetArgs()
        {
        }
    }
}
