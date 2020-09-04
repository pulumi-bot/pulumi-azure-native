// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.AlertsManagement.V20190505Preview.Outputs
{

    [OutputType]
    public sealed class ConditionResponseResult
    {
        /// <summary>
        /// operator for a given condition
        /// </summary>
        public readonly string? Operator;
        /// <summary>
        /// list of values to match for a given condition.
        /// </summary>
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private ConditionResponseResult(
            string? @operator,

            ImmutableArray<string> values)
        {
            Operator = @operator;
            Values = values;
        }
    }
}
