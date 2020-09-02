// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getDiagnosticSetting(args: GetDiagnosticSettingArgs, opts?: pulumi.InvokeOptions): Promise<GetDiagnosticSettingResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:insights/v20170501preview:getDiagnosticSetting", {
        "name": args.name,
        "resourceUri": args.resourceUri,
    }, opts);
}

export interface GetDiagnosticSettingArgs {
    /**
     * The name of the diagnostic setting.
     */
    readonly name: string;
    /**
     * The identifier of the resource.
     */
    readonly resourceUri: string;
}

/**
 * The diagnostic setting resource.
 */
export interface GetDiagnosticSettingResult {
    /**
     * The resource Id for the event hub authorization rule.
     */
    readonly eventHubAuthorizationRuleId?: string;
    /**
     * The name of the event hub. If none is specified, the default event hub will be selected.
     */
    readonly eventHubName?: string;
    /**
     * A string indicating whether the export to Log Analytics should use the default destination type, i.e. AzureDiagnostics, or use a destination type constructed as follows: <normalized service identity>_<normalized category name>. Possible values are: Dedicated and null (null is default.)
     */
    readonly logAnalyticsDestinationType?: string;
    /**
     * The list of logs settings.
     */
    readonly logs?: outputs.insights.v20170501preview.LogSettingsResponse[];
    /**
     * The list of metric settings.
     */
    readonly metrics?: outputs.insights.v20170501preview.MetricSettingsResponse[];
    /**
     * Azure resource name
     */
    readonly name: string;
    /**
     * The service bus rule Id of the diagnostic setting. This is here to maintain backwards compatibility.
     */
    readonly serviceBusRuleId?: string;
    /**
     * The resource ID of the storage account to which you would like to send Diagnostic Logs.
     */
    readonly storageAccountId?: string;
    /**
     * Azure resource type
     */
    readonly type: string;
    /**
     * The full ARM resource ID of the Log Analytics workspace to which you would like to send Diagnostic Logs. Example: /subscriptions/4b9e8510-67ab-4e9a-95a9-e2f1e570ea9c/resourceGroups/insights-integration/providers/Microsoft.OperationalInsights/workspaces/viruela2
     */
    readonly workspaceId?: string;
}
