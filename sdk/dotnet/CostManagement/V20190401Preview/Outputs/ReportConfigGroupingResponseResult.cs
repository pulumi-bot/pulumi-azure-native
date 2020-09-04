// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.CostManagement.V20190401Preview.Outputs
{

    [OutputType]
    public sealed class ReportConfigGroupingResponseResult
    {
        /// <summary>
        /// The name of the column to group. This version supports subscription lowest possible grain.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Has type of the column to group.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ReportConfigGroupingResponseResult(
            string name,

            string type)
        {
            Name = name;
            Type = type;
        }
    }
}
