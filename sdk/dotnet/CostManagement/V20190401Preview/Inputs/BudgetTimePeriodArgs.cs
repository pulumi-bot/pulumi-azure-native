// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.CostManagement.V20190401Preview.Inputs
{

    /// <summary>
    /// The start and end date for a budget.
    /// </summary>
    public sealed class BudgetTimePeriodArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The end date for the budget. If not provided, we default this to 10 years from the start date.
        /// </summary>
        [Input("endDate")]
        public Input<string>? EndDate { get; set; }

        /// <summary>
        /// The start date for the budget.
        /// </summary>
        [Input("startDate", required: true)]
        public Input<string> StartDate { get; set; } = null!;

        public BudgetTimePeriodArgs()
        {
        }
    }
}
