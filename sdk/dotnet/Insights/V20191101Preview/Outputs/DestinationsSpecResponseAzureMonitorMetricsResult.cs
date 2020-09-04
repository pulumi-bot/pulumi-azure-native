// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Insights.V20191101Preview.Outputs
{

    [OutputType]
    public sealed class DestinationsSpecResponseAzureMonitorMetricsResult
    {
        /// <summary>
        /// A friendly name for the destination. 
        /// This name should be unique across all destinations (regardless of type) within the data collection rule.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private DestinationsSpecResponseAzureMonitorMetricsResult(string name)
        {
            Name = name;
        }
    }
}
