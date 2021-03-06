// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.Latest.Inputs
{

    /// <summary>
    /// SKU of an Azure Firewall.
    /// </summary>
    public sealed class AzureFirewallSkuArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of an Azure Firewall SKU.
        /// </summary>
        [Input("name")]
        public InputUnion<string, Pulumi.AzureNative.Network.Latest.AzureFirewallSkuName>? Name { get; set; }

        /// <summary>
        /// Tier of an Azure Firewall.
        /// </summary>
        [Input("tier")]
        public InputUnion<string, Pulumi.AzureNative.Network.Latest.AzureFirewallSkuTier>? Tier { get; set; }

        public AzureFirewallSkuArgs()
        {
        }
    }
}
