// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20200701.Outputs
{

    [OutputType]
    public sealed class FirewallPolicyIntrusionDetectionResponse
    {
        /// <summary>
        /// Intrusion detection configuration properties.
        /// </summary>
        public readonly Outputs.FirewallPolicyIntrusionDetectionConfigurationResponse? Configuration;
        /// <summary>
        /// Intrusion detection general state.
        /// </summary>
        public readonly string? Mode;

        [OutputConstructor]
        private FirewallPolicyIntrusionDetectionResponse(
            Outputs.FirewallPolicyIntrusionDetectionConfigurationResponse? configuration,

            string? mode)
        {
            Configuration = configuration;
            Mode = mode;
        }
    }
}
