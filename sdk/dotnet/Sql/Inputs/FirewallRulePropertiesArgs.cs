// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Sql.Inputs
{

    /// <summary>
    /// Represents the properties of a server firewall rule.
    /// </summary>
    public sealed class FirewallRulePropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The end IP address of the firewall rule. Must be IPv4 format. Must be greater than or equal to startIpAddress. Use value '0.0.0.0' to represent all Azure-internal IP addresses.
        /// </summary>
        [Input("endIpAddress", required: true)]
        public Input<string> EndIpAddress { get; set; } = null!;

        /// <summary>
        /// The start IP address of the firewall rule. Must be IPv4 format. Use value '0.0.0.0' to represent all Azure-internal IP addresses.
        /// </summary>
        [Input("startIpAddress", required: true)]
        public Input<string> StartIpAddress { get; set; } = null!;

        public FirewallRulePropertiesArgs()
        {
        }
    }
}