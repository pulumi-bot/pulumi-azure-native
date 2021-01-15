// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Automation.Latest.Outputs
{

    [OutputType]
    public sealed class TargetPropertiesResponse
    {
        /// <summary>
        /// List of Azure queries in the software update configuration.
        /// </summary>
        public readonly ImmutableArray<Outputs.AzureQueryPropertiesResponse> AzureQueries;
        /// <summary>
        /// List of non Azure queries in the software update configuration.
        /// </summary>
        public readonly ImmutableArray<Outputs.NonAzureQueryPropertiesResponse> NonAzureQueries;

        [OutputConstructor]
        private TargetPropertiesResponse(
            ImmutableArray<Outputs.AzureQueryPropertiesResponse> azureQueries,

            ImmutableArray<Outputs.NonAzureQueryPropertiesResponse> nonAzureQueries)
        {
            AzureQueries = azureQueries;
            NonAzureQueries = nonAzureQueries;
        }
    }
}
