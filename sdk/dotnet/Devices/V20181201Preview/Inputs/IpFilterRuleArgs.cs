// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Devices.V20181201Preview.Inputs
{

    /// <summary>
    /// The IP filter rules for the IoT hub.
    /// </summary>
    public sealed class IpFilterRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The desired action for requests captured by this rule.
        /// </summary>
        [Input("action", required: true)]
        public Input<string> Action { get; set; } = null!;

        /// <summary>
        /// The name of the IP filter rule.
        /// </summary>
        [Input("filterName", required: true)]
        public Input<string> FilterName { get; set; } = null!;

        /// <summary>
        /// A string that contains the IP address range in CIDR notation for the rule.
        /// </summary>
        [Input("ipMask", required: true)]
        public Input<string> IpMask { get; set; } = null!;

        public IpFilterRuleArgs()
        {
        }
    }
}
