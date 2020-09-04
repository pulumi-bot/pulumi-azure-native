// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Consumption.V20171230Preview
{
    public static class GetBudget
    {
        public static Task<GetBudgetResult> InvokeAsync(GetBudgetArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBudgetResult>("azurerm:consumption/v20171230preview:getBudget", args ?? new GetBudgetArgs(), options.WithVersion());
    }


    public sealed class GetBudgetArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Budget name.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetBudgetArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetBudgetResult
    {
        /// <summary>
        /// The total amount of cost to track with the budget
        /// </summary>
        public readonly double Amount;
        /// <summary>
        /// The category of the budget, whether the budget tracks cost or something else.
        /// </summary>
        public readonly string Category;
        /// <summary>
        /// The current amount of cost which is being tracked for a budget.
        /// </summary>
        public readonly Outputs.CurrentSpendResponseResult CurrentSpend;
        /// <summary>
        /// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
        /// </summary>
        public readonly string? ETag;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Dictionary of notifications associated with the budget. Budget can have up to five notifications.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.NotificationResponseResult>? Notifications;
        /// <summary>
        /// The time covered by a budget. Tracking of the amount will be reset based on the time grain.
        /// </summary>
        public readonly string TimeGrain;
        /// <summary>
        /// Has start and end date of the budget. The start date must be first of the month and should be less than the end date. Budget start date must be on or after June 1, 2017. Future start date should not be more than three months. Past start date should  be selected within the timegrain period. There are no restrictions on the end date.
        /// </summary>
        public readonly Outputs.BudgetTimePeriodResponseResult TimePeriod;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetBudgetResult(
            double amount,

            string category,

            Outputs.CurrentSpendResponseResult currentSpend,

            string? eTag,

            string name,

            ImmutableDictionary<string, Outputs.NotificationResponseResult>? notifications,

            string timeGrain,

            Outputs.BudgetTimePeriodResponseResult timePeriod,

            string type)
        {
            Amount = amount;
            Category = category;
            CurrentSpend = currentSpend;
            ETag = eTag;
            Name = name;
            Notifications = notifications;
            TimeGrain = timeGrain;
            TimePeriod = timePeriod;
            Type = type;
        }
    }
}
