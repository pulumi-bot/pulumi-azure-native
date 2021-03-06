// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SecurityInsights.Outputs
{

    [OutputType]
    public sealed class InsightsTableResultResponseResult
    {
        /// <summary>
        /// Columns Metadata of the table
        /// </summary>
        public readonly ImmutableArray<Outputs.InsightsTableResultResponseColumnsResult> Columns;
        /// <summary>
        /// Rows data of the table
        /// </summary>
        public readonly ImmutableArray<ImmutableArray<string>> Rows;

        [OutputConstructor]
        private InsightsTableResultResponseResult(
            ImmutableArray<Outputs.InsightsTableResultResponseColumnsResult> columns,

            ImmutableArray<ImmutableArray<string>> rows)
        {
            Columns = columns;
            Rows = rows;
        }
    }
}
