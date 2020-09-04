// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.AlertsManagement.V20181102privatePreview.Outputs
{

    [OutputType]
    public sealed class ScopeResponseResult
    {
        /// <summary>
        /// type of target scope
        /// </summary>
        public readonly string? Type;
        /// <summary>
        /// list of ARM IDs of the given scope type which will be the target of the given action rule.
        /// </summary>
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private ScopeResponseResult(
            string? type,

            ImmutableArray<string> values)
        {
            Type = type;
            Values = values;
        }
    }
}
