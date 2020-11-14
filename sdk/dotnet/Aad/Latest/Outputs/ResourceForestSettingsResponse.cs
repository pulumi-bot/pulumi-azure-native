// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Aad.Latest.Outputs
{

    [OutputType]
    public sealed class ResourceForestSettingsResponse
    {
        /// <summary>
        /// Resource Forest
        /// </summary>
        public readonly string? ResourceForest;
        /// <summary>
        /// List of settings for Resource Forest
        /// </summary>
        public readonly ImmutableArray<Outputs.ForestTrustResponse> Settings;

        [OutputConstructor]
        private ResourceForestSettingsResponse(
            string? resourceForest,

            ImmutableArray<Outputs.ForestTrustResponse> settings)
        {
            ResourceForest = resourceForest;
            Settings = settings;
        }
    }
}
