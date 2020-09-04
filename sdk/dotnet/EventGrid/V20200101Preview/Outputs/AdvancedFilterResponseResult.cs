// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.EventGrid.V20200101Preview.Outputs
{

    [OutputType]
    public sealed class AdvancedFilterResponseResult
    {
        /// <summary>
        /// The field/property in the event based on which you want to filter.
        /// </summary>
        public readonly string? Key;
        /// <summary>
        /// The operator type used for filtering, e.g., NumberIn, StringContains, BoolEquals and others.
        /// </summary>
        public readonly string OperatorType;

        [OutputConstructor]
        private AdvancedFilterResponseResult(
            string? key,

            string operatorType)
        {
            Key = key;
            OperatorType = operatorType;
        }
    }
}
