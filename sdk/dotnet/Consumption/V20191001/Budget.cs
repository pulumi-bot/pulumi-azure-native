// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Consumption.V20191001
{
    /// <summary>
    /// A budget resource.
    /// </summary>
    public partial class Budget : Pulumi.CustomResource
    {
        /// <summary>
        /// The total amount of cost to track with the budget
        /// </summary>
        [Output("amount")]
        public Output<double> Amount { get; private set; } = null!;

        /// <summary>
        /// The category of the budget, whether the budget tracks cost or usage.
        /// </summary>
        [Output("category")]
        public Output<string> Category { get; private set; } = null!;

        /// <summary>
        /// The current amount of cost which is being tracked for a budget.
        /// </summary>
        [Output("currentSpend")]
        public Output<Outputs.CurrentSpendResponseResult> CurrentSpend { get; private set; } = null!;

        /// <summary>
        /// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
        /// </summary>
        [Output("eTag")]
        public Output<string?> ETag { get; private set; } = null!;

        /// <summary>
        /// May be used to filter budgets by resource group, resource, or meter.
        /// </summary>
        [Output("filter")]
        public Output<Outputs.BudgetFilterResponseResult?> Filter { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Dictionary of notifications associated with the budget. Budget can have up to five notifications.
        /// </summary>
        [Output("notifications")]
        public Output<ImmutableDictionary<string, Outputs.NotificationResponseResult>?> Notifications { get; private set; } = null!;

        /// <summary>
        /// The time covered by a budget. Tracking of the amount will be reset based on the time grain. BillingMonth, BillingQuarter, and BillingAnnual are only supported by WD customers
        /// </summary>
        [Output("timeGrain")]
        public Output<string> TimeGrain { get; private set; } = null!;

        /// <summary>
        /// Has start and end date of the budget. The start date must be first of the month and should be less than the end date. Budget start date must be on or after June 1, 2017. Future start date should not be more than twelve months. Past start date should  be selected within the timegrain period. There are no restrictions on the end date.
        /// </summary>
        [Output("timePeriod")]
        public Output<Outputs.BudgetTimePeriodResponseResult> TimePeriod { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Budget resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Budget(string name, BudgetArgs args, CustomResourceOptions? options = null)
            : base("azurerm:consumption/v20191001:Budget", name, args ?? new BudgetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Budget(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:consumption/v20191001:Budget", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:consumption/latest:Budget"},
                    new Pulumi.Alias { Type = "azurerm:consumption/v20190101:Budget"},
                    new Pulumi.Alias { Type = "azurerm:consumption/v20190401preview:Budget"},
                    new Pulumi.Alias { Type = "azurerm:consumption/v20190501:Budget"},
                    new Pulumi.Alias { Type = "azurerm:consumption/v20190501preview:Budget"},
                    new Pulumi.Alias { Type = "azurerm:consumption/v20190601:Budget"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Budget resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Budget Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Budget(name, id, options);
        }
    }

    public sealed class BudgetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The total amount of cost to track with the budget
        /// </summary>
        [Input("amount", required: true)]
        public Input<double> Amount { get; set; } = null!;

        /// <summary>
        /// Budget Name.
        /// </summary>
        [Input("budgetName", required: true)]
        public Input<string> BudgetName { get; set; } = null!;

        /// <summary>
        /// The category of the budget, whether the budget tracks cost or usage.
        /// </summary>
        [Input("category", required: true)]
        public Input<string> Category { get; set; } = null!;

        /// <summary>
        /// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
        /// </summary>
        [Input("eTag")]
        public Input<string>? ETag { get; set; }

        /// <summary>
        /// May be used to filter budgets by resource group, resource, or meter.
        /// </summary>
        [Input("filter")]
        public Input<Inputs.BudgetFilterArgs>? Filter { get; set; }

        [Input("notifications")]
        private InputMap<Inputs.NotificationArgs>? _notifications;

        /// <summary>
        /// Dictionary of notifications associated with the budget. Budget can have up to five notifications.
        /// </summary>
        public InputMap<Inputs.NotificationArgs> Notifications
        {
            get => _notifications ?? (_notifications = new InputMap<Inputs.NotificationArgs>());
            set => _notifications = value;
        }

        /// <summary>
        /// The scope associated with budget operations. This includes '/subscriptions/{subscriptionId}/' for subscription scope, '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, '/providers/Microsoft.Management/managementGroups/{managementGroupId}' for Management Group scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for billingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for invoiceSection scope.
        /// </summary>
        [Input("scope", required: true)]
        public Input<string> Scope { get; set; } = null!;

        /// <summary>
        /// The time covered by a budget. Tracking of the amount will be reset based on the time grain. BillingMonth, BillingQuarter, and BillingAnnual are only supported by WD customers
        /// </summary>
        [Input("timeGrain", required: true)]
        public Input<string> TimeGrain { get; set; } = null!;

        /// <summary>
        /// Has start and end date of the budget. The start date must be first of the month and should be less than the end date. Budget start date must be on or after June 1, 2017. Future start date should not be more than twelve months. Past start date should  be selected within the timegrain period. There are no restrictions on the end date.
        /// </summary>
        [Input("timePeriod", required: true)]
        public Input<Inputs.BudgetTimePeriodArgs> TimePeriod { get; set; } = null!;

        public BudgetArgs()
        {
        }
    }
}
