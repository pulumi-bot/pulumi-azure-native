// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Storage.V20180301Preview.Inputs
{

    /// <summary>
    /// Network rule set
    /// </summary>
    public sealed class NetworkRuleSetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether traffic is bypassed for Logging/Metrics/AzureServices. Possible values are any combination of Logging|Metrics|AzureServices (For example, "Logging, Metrics"), or None to bypass none of those traffics.
        /// </summary>
        [Input("bypass")]
        public InputUnion<string, Pulumi.AzureNextGen.Storage.V20180301Preview.Bypass>? Bypass { get; set; }

        /// <summary>
        /// Specifies the default action of allow or deny when no other rules match.
        /// </summary>
        [Input("defaultAction", required: true)]
        public Input<Pulumi.AzureNextGen.Storage.V20180301Preview.DefaultAction> DefaultAction { get; set; } = null!;

        [Input("ipRules")]
        private InputList<Inputs.IPRuleArgs>? _ipRules;

        /// <summary>
        /// Sets the IP ACL rules
        /// </summary>
        public InputList<Inputs.IPRuleArgs> IpRules
        {
            get => _ipRules ?? (_ipRules = new InputList<Inputs.IPRuleArgs>());
            set => _ipRules = value;
        }

        [Input("virtualNetworkRules")]
        private InputList<Inputs.VirtualNetworkRuleArgs>? _virtualNetworkRules;

        /// <summary>
        /// Sets the virtual network rules
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
