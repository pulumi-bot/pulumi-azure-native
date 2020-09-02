// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.AnalysisServices.V20170801beta.Inputs
{

    /// <summary>
    /// The detail of firewall rule.
    /// </summary>
    public sealed class IPv4FirewallRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The rule name.
        /// </summary>
        [Input("firewallRuleName")]
        public Input<string>? FirewallRuleName { get; set; }

        /// <summary>
        /// The end range of IPv4.
        /// </summary>
        [Input("rangeEnd")]
        public Input<string>? RangeEnd { get; set; }

        /// <summary>
        /// The start range of IPv4.
        /// </summary>
        [Input("rangeStart")]
        public Input<string>? RangeStart { get; set; }

        public IPv4FirewallRuleArgs()
        {
        }
    }
}
