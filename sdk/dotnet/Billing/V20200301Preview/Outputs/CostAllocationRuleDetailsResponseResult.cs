// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Billing.V20200301Preview.Outputs
{

    [OutputType]
    public sealed class CostAllocationRuleDetailsResponseResult
    {
        /// <summary>
        /// Source resources for cost allocation. At this time, this list can contain no more than one element.
        /// </summary>
        public readonly ImmutableArray<Outputs.SourceCostAllocationResourceResponseResult> SourceResources;
        /// <summary>
        /// Target resources for cost allocation. At this time, this list can contain no more than one element.
        /// </summary>
        public readonly ImmutableArray<Outputs.TargetCostAllocationResourceResponseResult> TargetResources;

        [OutputConstructor]
        private CostAllocationRuleDetailsResponseResult(
            ImmutableArray<Outputs.SourceCostAllocationResourceResponseResult> sourceResources,

            ImmutableArray<Outputs.TargetCostAllocationResourceResponseResult> targetResources)
        {
            SourceResources = sourceResources;
            TargetResources = targetResources;
        }
    }
}
