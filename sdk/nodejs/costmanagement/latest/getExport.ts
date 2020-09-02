// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getExport(args: GetExportArgs, opts?: pulumi.InvokeOptions): Promise<GetExportResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:costmanagement/latest:getExport", {
        "expand": args.expand,
        "exportName": args.exportName,
        "scope": args.scope,
    }, opts);
}

export interface GetExportArgs {
    /**
     * May be used to expand the properties within an export. Currently only 'runHistory' is supported and will return information for the last 10 executions of the export.
     */
    readonly expand?: string;
    /**
     * Export Name.
     */
    readonly exportName: string;
    /**
     * The scope associated with query and export operations. This includes '/subscriptions/{subscriptionId}/' for subscription scope, '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope and '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, '/providers/Microsoft.Management/managementGroups/{managementGroupId} for Management Group scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for billingProfile scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}/invoiceSections/{invoiceSectionId}' for invoiceSection scope, and '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/customers/{customerId}' specific for partners.
     */
    readonly scope: string;
}

/**
 * An export resource.
 */
export interface GetExportResult {
    /**
     * Has the definition for the export.
     */
    readonly definition: outputs.costmanagement.latest.ExportDefinitionResponse;
    /**
     * Has delivery information for the export.
     */
    readonly deliveryInfo: outputs.costmanagement.latest.ExportDeliveryInfoResponse;
    /**
     * eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
     */
    readonly eTag?: string;
    /**
     * The format of the export being delivered. Currently only 'Csv' is supported.
     */
    readonly format?: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * If the export has an active schedule, provides an estimate of the next execution time.
     */
    readonly nextRunTimeEstimate: string;
    /**
     * If requested, has the most recent execution history for the export.
     */
    readonly runHistory?: outputs.costmanagement.latest.ExportExecutionListResultResponse;
    /**
     * Has schedule information for the export.
     */
    readonly schedule?: outputs.costmanagement.latest.ExportScheduleResponse;
    /**
     * Resource type.
     */
    readonly type: string;
}