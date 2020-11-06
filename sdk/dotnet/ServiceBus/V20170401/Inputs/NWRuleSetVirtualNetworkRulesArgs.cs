// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.ServiceBus.V20170401.Inputs
{

    /// <summary>
    /// Description of VirtualNetworkRules - NetworkRules resource.
    /// </summary>
    public sealed class NWRuleSetVirtualNetworkRulesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Value that indicates whether to ignore missing VNet Service Endpoint
        /// </summary>
        [Input("ignoreMissingVnetServiceEndpoint")]
        public Input<bool>? IgnoreMissingVnetServiceEndpoint { get; set; }

        /// <summary>
        /// Subnet properties
        /// </summary>
        [Input("subnet")]
        public Input<Inputs.SubnetArgs>? Subnet { get; set; }

        public NWRuleSetVirtualNetworkRulesArgs()
        {
        }
    }
}
