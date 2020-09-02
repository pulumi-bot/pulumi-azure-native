// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.EventGrid.V20200401Preview.Outputs
{

    [OutputType]
    public sealed class EventChannelFilterResponseResult
    {
        /// <summary>
        /// An array of advanced filters that are used for filtering event channels.
        /// </summary>
        public readonly ImmutableArray<Outputs.AdvancedFilterResponseResult> AdvancedFilters;

        [OutputConstructor]
        private EventChannelFilterResponseResult(ImmutableArray<Outputs.AdvancedFilterResponseResult> advancedFilters)
        {
            AdvancedFilters = advancedFilters;
        }
    }
}
