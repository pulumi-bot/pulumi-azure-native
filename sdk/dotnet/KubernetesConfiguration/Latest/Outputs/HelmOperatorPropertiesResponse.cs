// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.KubernetesConfiguration.Latest.Outputs
{

    [OutputType]
    public sealed class HelmOperatorPropertiesResponse
    {
        /// <summary>
        /// Values override for the operator Helm chart.
        /// </summary>
        public readonly string? ChartValues;
        /// <summary>
        /// Version of the operator Helm chart.
        /// </summary>
        public readonly string? ChartVersion;

        [OutputConstructor]
        private HelmOperatorPropertiesResponse(
            string? chartValues,

            string? chartVersion)
        {
            ChartValues = chartValues;
            ChartVersion = chartVersion;
        }
    }
}
