// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Insights.V20180601Preview.Outputs
{

    [OutputType]
    public sealed class PerformanceCounterConfigurationResponseResult
    {
        public readonly string? Instance;
        public readonly string Name;
        public readonly string SamplingPeriod;

        [OutputConstructor]
        private PerformanceCounterConfigurationResponseResult(
            string? instance,

            string name,

            string samplingPeriod)
        {
            Instance = instance;
            Name = name;
            SamplingPeriod = samplingPeriod;
        }
    }
}
