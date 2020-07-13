// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.Outputs
{

    [OutputType]
    public sealed class HealthProbeSettingsPropertiesResponse
    {
        /// <summary>
        /// Whether to enable health probes to be made against backends defined under backendPools. Health probes can only be disabled if there is a single enabled backend in single enabled backend pool.
        /// </summary>
        public readonly string? EnabledState;
        /// <summary>
        /// Configures which HTTP method to use to probe the backends defined under backendPools.
        /// </summary>
        public readonly string? HealthProbeMethod;
        /// <summary>
        /// The number of seconds between health probes.
        /// </summary>
        public readonly int? IntervalInSeconds;
        /// <summary>
        /// The path to use for the health probe. Default is /
        /// </summary>
        public readonly string? Path;
        /// <summary>
        /// Protocol scheme to use for this probe
        /// </summary>
        public readonly string? Protocol;
        /// <summary>
        /// Resource status.
        /// </summary>
        public readonly string? ResourceState;

        [OutputConstructor]
        private HealthProbeSettingsPropertiesResponse(
            string? enabledState,

            string? healthProbeMethod,

            int? intervalInSeconds,

            string? path,

            string? protocol,

            string? resourceState)
        {
            EnabledState = enabledState;
            HealthProbeMethod = healthProbeMethod;
            IntervalInSeconds = intervalInSeconds;
            Path = path;
            Protocol = protocol;
            ResourceState = resourceState;
        }
    }
}