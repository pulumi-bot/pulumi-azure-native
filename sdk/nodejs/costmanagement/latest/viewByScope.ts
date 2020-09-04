// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * States and configurations of Cost Analysis.
 */
export class ViewByScope extends pulumi.CustomResource {
    /**
     * Get an existing ViewByScope resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ViewByScope {
        return new ViewByScope(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:costmanagement/latest:ViewByScope';

    /**
     * Returns true if the given object is an instance of ViewByScope.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ViewByScope {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ViewByScope.__pulumiType;
    }

    /**
     * Show costs accumulated over time.
     */
    public readonly accumulated!: pulumi.Output<string | undefined>;
    /**
     * Chart type of the main view in Cost Analysis. Required.
     */
    public readonly chart!: pulumi.Output<string | undefined>;
    /**
     * Date the user created this view.
     */
    public /*out*/ readonly createdOn!: pulumi.Output<string>;
    /**
     * Has definition for data in this report config.
     */
    public readonly dataset!: pulumi.Output<outputs.costmanagement.latest.ReportConfigDatasetResponse | undefined>;
    /**
     * User input name of the view. Required.
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
     */
    public readonly eTag!: pulumi.Output<string | undefined>;
    /**
     * List of KPIs to show in Cost Analysis UI.
     */
    public readonly kpis!: pulumi.Output<outputs.costmanagement.latest.KpiPropertiesResponse[] | undefined>;
    /**
     * Metric to use when displaying costs.
     */
    public readonly metric!: pulumi.Output<string | undefined>;
    /**
     * Date when the user last modified this view.
     */
    public /*out*/ readonly modifiedOn!: pulumi.Output<string>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Configuration of 3 sub-views in the Cost Analysis UI.
     */
    public readonly pivots!: pulumi.Output<outputs.costmanagement.latest.PivotPropertiesResponse[] | undefined>;
    /**
     * Cost Management scope to save the view on. This includes 'subscriptions/{subscriptionId}' for subscription scope, 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for BillingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for InvoiceSection scope, 'providers/Microsoft.Management/managementGroups/{managementGroupId}' for Management Group scope, '/providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountName}' for ExternalBillingAccount scope, and '/providers/Microsoft.CostManagement/externalSubscriptions/{externalSubscriptionName}' for ExternalSubscription scope.
     */
    public readonly scope!: pulumi.Output<string | undefined>;
    /**
     * Has time period for pulling data for the report.
     */
    public readonly timePeriod!: pulumi.Output<outputs.costmanagement.latest.ReportConfigTimePeriodResponse | undefined>;
    /**
     * The time frame for pulling data for the report. If custom, then a specific time period must be provided.
     */
    public readonly timeframe!: pulumi.Output<string>;
    /**
     * Resource type.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a ViewByScope resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ViewByScopeArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.scope === undefined) {
                throw new Error("Missing required property 'scope'");
            }
            if (!args || args.timeframe === undefined) {
                throw new Error("Missing required property 'timeframe'");
            }
            if (!args || args.type === undefined) {
                throw new Error("Missing required property 'type'");
            }
            if (!args || args.viewName === undefined) {
                throw new Error("Missing required property 'viewName'");
            }
            inputs["accumulated"] = args ? args.accumulated : undefined;
            inputs["chart"] = args ? args.chart : undefined;
            inputs["dataset"] = args ? args.dataset : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["eTag"] = args ? args.eTag : undefined;
            inputs["kpis"] = args ? args.kpis : undefined;
            inputs["metric"] = args ? args.metric : undefined;
            inputs["pivots"] = args ? args.pivots : undefined;
            inputs["scope"] = args ? args.scope : undefined;
            inputs["timePeriod"] = args ? args.timePeriod : undefined;
            inputs["timeframe"] = args ? args.timeframe : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["viewName"] = args ? args.viewName : undefined;
            inputs["createdOn"] = undefined /*out*/;
            inputs["modifiedOn"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
        } else {
            inputs["accumulated"] = undefined /*out*/;
            inputs["chart"] = undefined /*out*/;
            inputs["createdOn"] = undefined /*out*/;
            inputs["dataset"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["eTag"] = undefined /*out*/;
            inputs["kpis"] = undefined /*out*/;
            inputs["metric"] = undefined /*out*/;
            inputs["modifiedOn"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["pivots"] = undefined /*out*/;
            inputs["scope"] = undefined /*out*/;
            inputs["timePeriod"] = undefined /*out*/;
            inputs["timeframe"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:costmanagement/v20190401preview:ViewByScope" }, { type: "azurerm:costmanagement/v20191101:ViewByScope" }, { type: "azurerm:costmanagement/v20200601:ViewByScope" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ViewByScope.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ViewByScope resource.
 */
export interface ViewByScopeArgs {
    /**
     * Show costs accumulated over time.
     */
    readonly accumulated?: pulumi.Input<string>;
    /**
     * Chart type of the main view in Cost Analysis. Required.
     */
    readonly chart?: pulumi.Input<string>;
    /**
     * Has definition for data in this report config.
     */
    readonly dataset?: pulumi.Input<inputs.costmanagement.latest.ReportConfigDataset>;
    /**
     * User input name of the view. Required.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
     */
    readonly eTag?: pulumi.Input<string>;
    /**
     * List of KPIs to show in Cost Analysis UI.
     */
    readonly kpis?: pulumi.Input<pulumi.Input<inputs.costmanagement.latest.KpiProperties>[]>;
    /**
     * Metric to use when displaying costs.
     */
    readonly metric?: pulumi.Input<string>;
    /**
     * Configuration of 3 sub-views in the Cost Analysis UI.
     */
    readonly pivots?: pulumi.Input<pulumi.Input<inputs.costmanagement.latest.PivotProperties>[]>;
    /**
     * Cost Management scope to save the view on. This includes 'subscriptions/{subscriptionId}' for subscription scope, 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for BillingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for InvoiceSection scope, 'providers/Microsoft.Management/managementGroups/{managementGroupId}' for Management Group scope, '/providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountName}' for ExternalBillingAccount scope, and '/providers/Microsoft.CostManagement/externalSubscriptions/{externalSubscriptionName}' for ExternalSubscription scope.
     */
    readonly scope: pulumi.Input<string>;
    /**
     * Has time period for pulling data for the report.
     */
    readonly timePeriod?: pulumi.Input<inputs.costmanagement.latest.ReportConfigTimePeriod>;
    /**
     * The time frame for pulling data for the report. If custom, then a specific time period must be provided.
     */
    readonly timeframe: pulumi.Input<string>;
    /**
     * The type of the report. Usage represents actual usage, forecast represents forecasted data and UsageAndForecast represents both usage and forecasted data. Actual usage and forecasted data can be differentiated based on dates.
     */
    readonly type: pulumi.Input<string>;
    /**
     * View name
     */
    readonly viewName: pulumi.Input<string>;
}
