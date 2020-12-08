// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.KeyVault.V20180214Preview.Inputs
{

    /// <summary>
    /// A set of rules governing the network accessibility of a vault.
    /// </summary>
    public sealed class NetworkRuleSetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Tells what traffic can bypass network rules. This can be 'AzureServices' or 'None'.  If not specified the default is 'AzureServices'.
        /// </summary>
        [Input("bypass")]
        public InputUnion<string, Pulumi.AzureNextGen.KeyVault.V20180214Preview.NetworkRuleBypassOptions>? Bypass { get; set; }

        /// <summary>
        /// The default action when no rule from ipRules and from virtualNetworkRules match. This is only used after the bypass property has been evaluated.
        /// </summary>
        [Input("defaultAction")]
        public InputUnion<string, Pulumi.AzureNextGen.KeyVault.V20180214Preview.NetworkRuleAction>? DefaultAction { get; set; }

        [Input("ipRules")]
        private InputList<Inputs.IPRuleArgs>? _ipRules;

        /// <summary>
        /// The list of IP address rules.
        /// </summary>
        public InputList<Inputs.IPRuleArgs> IpRules
        {
            get => _ipRules ?? (_ipRules = new InputList<Inputs.IPRuleArgs>());
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
