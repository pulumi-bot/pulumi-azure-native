// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * A budget resource.
 *
 * ## Example Usage
 * ### CreateOrUpdateBudget
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const budget = new azurerm.consumption.latest.Budget("budget", {
 *     amount: 100.65,
 *     budgetName: "TestBudget",
 *     category: "Cost",
 *     eTag: "\"1d34d016a593709\"",
 *     filter: {
 *         and: [
 *             {
 *                 dimensions: {
 *                     name: "ResourceId",
 *                     operator: "In",
 *                     values: [
 *                         "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/Microsoft.Compute/virtualMachines/MSVM2",
 *                         "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/Microsoft.Compute/virtualMachines/platformcloudplatformGeneric1",
 *                     ],
 *                 },
 *             },
 *             {
 *                 tags: {
 *                     name: "category",
 *                     operator: "In",
 *                     values: [
 *                         "Dev",
 *                         "Prod",
 *                     ],
 *                 },
 *             },
 *             {
 *                 tags: {
 *                     name: "department",
 *                     operator: "In",
 *                     values: [
 *                         "engineering",
 *                         "sales",
 *                     ],
 *                 },
 *             },
 *         ],
 *     },
 *     notifications: {
 *         Actual_GreaterThan_80_Percent: {
 *             contactEmails: [
 *                 "johndoe@contoso.com",
 *                 "janesmith@contoso.com",
 *             ],
 *             contactGroups: ["/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/microsoft.insights/actionGroups/SampleActionGroup"],
 *             contactRoles: [
 *                 "Contributor",
 *                 "Reader",
 *             ],
 *             enabled: true,
 *             operator: "GreaterThan",
 *             threshold: 80,
 *             thresholdType: "Actual",
 *         },
 *     },
 *     scope: "subscriptions/00000000-0000-0000-0000-000000000000",
 *     timeGrain: "Monthly",
 *     timePeriod: {
 *         endDate: "2018-10-31T00:00:00Z",
 *         startDate: "2017-10-01T00:00:00Z",
 *     },
 * });
 *
 * ```
 */
export class Budget extends pulumi.CustomResource {
    /**
     * Get an existing Budget resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Budget {
        return new Budget(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:consumption/latest:Budget';

    /**
     * Returns true if the given object is an instance of Budget.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Budget {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Budget.__pulumiType;
    }

    /**
     * The total amount of cost to track with the budget
     */
    public readonly amount!: pulumi.Output<number>;
    /**
     * The category of the budget, whether the budget tracks cost or usage.
     */
    public readonly category!: pulumi.Output<string>;
    /**
     * The current amount of cost which is being tracked for a budget.
     */
    public /*out*/ readonly currentSpend!: pulumi.Output<outputs.consumption.latest.CurrentSpendResponse>;
    /**
     * eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
     */
    public readonly eTag!: pulumi.Output<string | undefined>;
    /**
     * May be used to filter budgets by resource group, resource, or meter.
     */
    public readonly filter!: pulumi.Output<outputs.consumption.latest.BudgetFilterResponse | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Dictionary of notifications associated with the budget. Budget can have up to five notifications.
     */
    public readonly notifications!: pulumi.Output<{[key: string]: outputs.consumption.latest.NotificationResponse} | undefined>;
    /**
     * The time covered by a budget. Tracking of the amount will be reset based on the time grain. BillingMonth, BillingQuarter, and BillingAnnual are only supported by WD customers
     */
    public readonly timeGrain!: pulumi.Output<string>;
    /**
     * Has start and end date of the budget. The start date must be first of the month and should be less than the end date. Budget start date must be on or after June 1, 2017. Future start date should not be more than twelve months. Past start date should  be selected within the timegrain period. There are no restrictions on the end date.
     */
    public readonly timePeriod!: pulumi.Output<outputs.consumption.latest.BudgetTimePeriodResponse>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Budget resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BudgetArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.amount === undefined) {
                throw new Error("Missing required property 'amount'");
            }
            if (!args || args.budgetName === undefined) {
                throw new Error("Missing required property 'budgetName'");
            }
            if (!args || args.category === undefined) {
                throw new Error("Missing required property 'category'");
            }
            if (!args || args.scope === undefined) {
                throw new Error("Missing required property 'scope'");
            }
            if (!args || args.timeGrain === undefined) {
                throw new Error("Missing required property 'timeGrain'");
            }
            if (!args || args.timePeriod === undefined) {
                throw new Error("Missing required property 'timePeriod'");
            }
            inputs["amount"] = args ? args.amount : undefined;
            inputs["budgetName"] = args ? args.budgetName : undefined;
            inputs["category"] = args ? args.category : undefined;
            inputs["eTag"] = args ? args.eTag : undefined;
            inputs["filter"] = args ? args.filter : undefined;
            inputs["notifications"] = args ? args.notifications : undefined;
            inputs["scope"] = args ? args.scope : undefined;
            inputs["timeGrain"] = args ? args.timeGrain : undefined;
            inputs["timePeriod"] = args ? args.timePeriod : undefined;
            inputs["currentSpend"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["amount"] = undefined /*out*/;
            inputs["category"] = undefined /*out*/;
            inputs["currentSpend"] = undefined /*out*/;
            inputs["eTag"] = undefined /*out*/;
            inputs["filter"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["notifications"] = undefined /*out*/;
            inputs["timeGrain"] = undefined /*out*/;
            inputs["timePeriod"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:consumption/v20190101:Budget" }, { type: "azurerm:consumption/v20190401preview:Budget" }, { type: "azurerm:consumption/v20190501:Budget" }, { type: "azurerm:consumption/v20190501preview:Budget" }, { type: "azurerm:consumption/v20190601:Budget" }, { type: "azurerm:consumption/v20191001:Budget" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Budget.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Budget resource.
 */
export interface BudgetArgs {
    /**
     * The total amount of cost to track with the budget
     */
    readonly amount: pulumi.Input<number>;
    /**
     * Budget Name.
     */
    readonly budgetName: pulumi.Input<string>;
    /**
     * The category of the budget, whether the budget tracks cost or usage.
     */
    readonly category: pulumi.Input<string>;
    /**
     * eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
     */
    readonly eTag?: pulumi.Input<string>;
    /**
     * May be used to filter budgets by resource group, resource, or meter.
     */
    readonly filter?: pulumi.Input<inputs.consumption.latest.BudgetFilter>;
    /**
     * Dictionary of notifications associated with the budget. Budget can have up to five notifications.
     */
    readonly notifications?: pulumi.Input<{[key: string]: pulumi.Input<inputs.consumption.latest.Notification>}>;
    /**
     * The scope associated with budget operations. This includes '/subscriptions/{subscriptionId}/' for subscription scope, '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, '/providers/Microsoft.Management/managementGroups/{managementGroupId}' for Management Group scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for billingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for invoiceSection scope.
     */
    readonly scope: pulumi.Input<string>;
    /**
     * The time covered by a budget. Tracking of the amount will be reset based on the time grain. BillingMonth, BillingQuarter, and BillingAnnual are only supported by WD customers
     */
    readonly timeGrain: pulumi.Input<string>;
    /**
     * Has start and end date of the budget. The start date must be first of the month and should be less than the end date. Budget start date must be on or after June 1, 2017. Future start date should not be more than twelve months. Past start date should  be selected within the timegrain period. There are no restrictions on the end date.
     */
    readonly timePeriod: pulumi.Input<inputs.consumption.latest.BudgetTimePeriod>;
}
