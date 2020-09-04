// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Billing.V20200301Preview.Inputs
{

    /// <summary>
    /// Resource details of the cost allocation rule
    /// </summary>
    public sealed class CostAllocationRuleDetailsArgs : Pulumi.ResourceArgs
    {
        [Input("sourceResources")]
        private InputList<Inputs.SourceCostAllocationResourceArgs>? _sourceResources;

        /// <summary>
        /// Source resources for cost allocation. At this time, this list can contain no more than one element.
        /// </summary>
        public InputList<Inputs.SourceCostAllocationResourceArgs> SourceResources
        {
            get => _sourceResources ?? (_sourceResources = new InputList<Inputs.SourceCostAllocationResourceArgs>());
            set => _sourceResources = value;
        }

        [Input("targetResources")]
        private InputList<Inputs.TargetCostAllocationResourceArgs>? _targetResources;

        /// <summary>
        /// Target resources for cost allocation. At this time, this list can contain no more than one element.
        /// </summary>
        public InputList<Inputs.TargetCostAllocationResourceArgs> TargetResources
        {
            get => _targetResources ?? (_targetResources = new InputList<Inputs.TargetCostAllocationResourceArgs>());
            set => _targetResources = value;
        }

        public CostAllocationRuleDetailsArgs()
        {
        }
    }
}
