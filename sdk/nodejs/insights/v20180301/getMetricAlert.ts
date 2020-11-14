// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getMetricAlert(args: GetMetricAlertArgs, opts?: pulumi.InvokeOptions): Promise<GetMetricAlertResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:insights/v20180301:getMetricAlert", {
        "resourceGroupName": args.resourceGroupName,
        "ruleName": args.ruleName,
    }, opts);
}

export interface GetMetricAlertArgs {
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the rule.
     */
    readonly ruleName: string;
}

/**
 * The metric alert resource.
 */
export interface GetMetricAlertResult {
    /**
     * the array of actions that are performed when the alert rule becomes active, and when an alert condition is resolved.
     */
    readonly actions?: outputs.insights.v20180301.MetricAlertActionResponse[];
    /**
     * the flag that indicates whether the alert should be auto resolved or not. The default is true.
     */
    readonly autoMitigate?: boolean;
    /**
     * defines the specific alert criteria information.
     */
    readonly criteria: outputs.insights.v20180301.MetricAlertMultipleResourceMultipleMetricCriteriaResponse | outputs.insights.v20180301.MetricAlertSingleResourceMultipleMetricCriteriaResponse | outputs.insights.v20180301.WebtestLocationAvailabilityCriteriaResponse;
    /**
     * the description of the metric alert that will be included in the alert email.
     */
    readonly description: string;
    /**
     * the flag that indicates whether the metric alert is enabled.
     */
    readonly enabled: boolean;
    /**
     * how often the metric alert is evaluated represented in ISO 8601 duration format.
     */
    readonly evaluationFrequency: string;
    /**
     * Last time the rule was updated in ISO8601 format.
     */
    readonly lastUpdatedTime: string;
    /**
     * Resource location
     */
    readonly location: string;
    /**
     * Azure resource name
     */
    readonly name: string;
    /**
     * the list of resource id's that this metric alert is scoped to.
     */
    readonly scopes?: string[];
    /**
     * Alert severity {0, 1, 2, 3, 4}
     */
    readonly severity: number;
    /**
     * Resource tags
     */
    readonly tags?: {[key: string]: string};
    /**
     * the region of the target resource(s) on which the alert is created/updated. Mandatory for MultipleResourceMultipleMetricCriteria.
     */
    readonly targetResourceRegion?: string;
    /**
     * the resource type of the target resource(s) on which the alert is created/updated. Mandatory for MultipleResourceMultipleMetricCriteria.
     */
    readonly targetResourceType?: string;
    /**
     * Azure resource type
     */
    readonly type: string;
    /**
     * the period of time (in ISO 8601 duration format) that is used to monitor alert activity based on the threshold.
     */
    readonly windowSize: string;
}
