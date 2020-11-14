// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getView(args: GetViewArgs, opts?: pulumi.InvokeOptions): Promise<GetViewResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:costmanagement/v20200601:getView", {
        "viewName": args.viewName,
    }, opts);
}

export interface GetViewArgs {
    /**
     * View name
     */
    readonly viewName: string;
}

/**
 * States and configurations of Cost Analysis.
 */
export interface GetViewResult {
    /**
     * Show costs accumulated over time.
     */
    readonly accumulated?: string;
    /**
     * Chart type of the main view in Cost Analysis. Required.
     */
    readonly chart?: string;
    /**
     * Date the user created this view.
     */
    readonly createdOn: string;
    /**
     * Has definition for data in this report config.
     */
    readonly dataset?: outputs.costmanagement.v20200601.ReportConfigDatasetResponse;
    /**
     * User input name of the view. Required.
     */
    readonly displayName?: string;
    /**
     * eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
     */
    readonly eTag?: string;
    /**
     * List of KPIs to show in Cost Analysis UI.
     */
    readonly kpis?: outputs.costmanagement.v20200601.KpiPropertiesResponse[];
    /**
     * Metric to use when displaying costs.
     */
    readonly metric?: string;
    /**
     * Date when the user last modified this view.
     */
    readonly modifiedOn: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * Configuration of 3 sub-views in the Cost Analysis UI.
     */
    readonly pivots?: outputs.costmanagement.v20200601.PivotPropertiesResponse[];
    /**
     * Cost Management scope to save the view on. This includes 'subscriptions/{subscriptionId}' for subscription scope, 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for BillingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for InvoiceSection scope, 'providers/Microsoft.Management/managementGroups/{managementGroupId}' for Management Group scope, '/providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountName}' for ExternalBillingAccount scope, and '/providers/Microsoft.CostManagement/externalSubscriptions/{externalSubscriptionName}' for ExternalSubscription scope.
     */
    readonly scope?: string;
    /**
     * Has time period for pulling data for the report.
     */
    readonly timePeriod?: outputs.costmanagement.v20200601.ReportConfigTimePeriodResponse;
    /**
     * The time frame for pulling data for the report. If custom, then a specific time period must be provided.
     */
    readonly timeframe: string;
    /**
     * Resource type.
     */
    readonly type: string;
}
