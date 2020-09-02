// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.CostManagement.V20180801Preview.Outputs
{

    [OutputType]
    public sealed class ReportTimePeriodResponseResult
    {
        /// <summary>
        /// The start date to pull data from.
        /// </summary>
        public readonly string From;
        /// <summary>
        /// The end date to pull data to.
        /// </summary>
        public readonly string To;

        [OutputConstructor]
        private ReportTimePeriodResponseResult(
            string from,

            string to)
        {
            From = from;
            To = to;
        }
    }
}
