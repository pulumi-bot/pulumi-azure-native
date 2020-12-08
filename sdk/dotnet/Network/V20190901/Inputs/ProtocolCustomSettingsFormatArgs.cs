// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20190901.Inputs
{

    /// <summary>
    /// DDoS custom policy properties.
    /// </summary>
    public sealed class ProtocolCustomSettingsFormatArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The protocol for which the DDoS protection policy is being customized.
        /// </summary>
        [Input("protocol")]
        public InputUnion<string, Pulumi.AzureNextGen.Network.V20190901.DdosCustomPolicyProtocol>? Protocol { get; set; }

        /// <summary>
        /// The customized DDoS protection source rate.
        /// </summary>
        [Input("sourceRateOverride")]
        public Input<string>? SourceRateOverride { get; set; }

        /// <summary>
        /// The customized DDoS protection trigger rate.
        /// </summary>
        [Input("triggerRateOverride")]
        public Input<string>? TriggerRateOverride { get; set; }

        /// <summary>
        /// The customized DDoS protection trigger rate sensitivity degrees. High: Trigger rate set with most sensitivity w.r.t. normal traffic. Default: Trigger rate set with moderate sensitivity w.r.t. normal traffic. Low: Trigger rate set with less sensitivity w.r.t. normal traffic. Relaxed: Trigger rate set with least sensitivity w.r.t. normal traffic.
        /// </summary>
        [Input("triggerSensitivityOverride")]
        public InputUnion<string, Pulumi.AzureNextGen.Network.V20190901.DdosCustomPolicyTriggerSensitivityOverride>? TriggerSensitivityOverride { get; set; }

        public ProtocolCustomSettingsFormatArgs()
        {
        }
    }
}
