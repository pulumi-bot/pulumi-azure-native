// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Devices.V20200831Preview.Inputs
{

    /// <summary>
    /// Network Rule Set Properties of IotHub
    /// </summary>
    public sealed class NetworkRuleSetPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// If True, then Network Rule Set is also applied to BuiltIn EventHub EndPoint of IotHub
        /// </summary>
        [Input("applyToBuiltInEventHubEndpoint", required: true)]
        public Input<bool> ApplyToBuiltInEventHubEndpoint { get; set; } = null!;

        /// <summary>
        /// Default Action for Network Rule Set
        /// </summary>
        [Input("defaultAction")]
        public InputUnion<string, Pulumi.AzureNextGen.Devices.V20200831Preview.DefaultAction>? DefaultAction { get; set; }

        [Input("ipRules", required: true)]
        private InputList<Inputs.NetworkRuleSetIpRuleArgs>? _ipRules;

        /// <summary>
        /// List of IP Rules
        /// </summary>
        public InputList<Inputs.NetworkRuleSetIpRuleArgs> IpRules
        {
            get => _ipRules ?? (_ipRules = new InputList<Inputs.NetworkRuleSetIpRuleArgs>());
            set => _ipRules = value;
        }

        public NetworkRuleSetPropertiesArgs()
        {
        }
    }
}
