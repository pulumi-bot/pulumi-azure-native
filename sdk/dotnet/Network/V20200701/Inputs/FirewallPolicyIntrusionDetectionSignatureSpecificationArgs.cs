// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20200701.Inputs
{

    /// <summary>
    /// Intrusion detection signatures specification states.
    /// </summary>
    public sealed class FirewallPolicyIntrusionDetectionSignatureSpecificationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Signature id.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The signature state.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        public FirewallPolicyIntrusionDetectionSignatureSpecificationArgs()
        {
        }
    }
}
