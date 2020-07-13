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
    public sealed class BackendPoolPropertiesResponse
    {
        /// <summary>
        /// The set of backends for this pool
        /// </summary>
        public readonly ImmutableArray<Outputs.BackendResponse> Backends;
        /// <summary>
        /// L7 health probe settings for a backend pool
        /// </summary>
        public readonly Outputs.SubResourceResponse? HealthProbeSettings;
        /// <summary>
        /// Load balancing settings for a backend pool
        /// </summary>
        public readonly Outputs.SubResourceResponse? LoadBalancingSettings;
        /// <summary>
        /// Resource status.
        /// </summary>
        public readonly string? ResourceState;

        [OutputConstructor]
        private BackendPoolPropertiesResponse(
            ImmutableArray<Outputs.BackendResponse> backends,

            Outputs.SubResourceResponse? healthProbeSettings,

            Outputs.SubResourceResponse? loadBalancingSettings,

            string? resourceState)
        {
            Backends = backends;
            HealthProbeSettings = healthProbeSettings;
            LoadBalancingSettings = loadBalancingSettings;
            ResourceState = resourceState;
        }
    }
}