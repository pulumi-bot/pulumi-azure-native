// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.ContainerRegistry.V20171001.Inputs
{

    /// <summary>
    /// Virtual network rule.
    /// </summary>
    public sealed class VirtualNetworkRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The action of virtual network rule.
        /// </summary>
        [Input("action")]
        public InputUnion<string, Pulumi.AzureNextGen.ContainerRegistry.V20171001.Action>? Action { get; set; }

        /// <summary>
        /// Resource ID of a subnet, for example: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{vnetName}/subnets/{subnetName}.
        /// </summary>
        [Input("virtualNetworkResourceId", required: true)]
        public Input<string> VirtualNetworkResourceId { get; set; } = null!;

        public VirtualNetworkRuleArgs()
        {
        }
    }
}
