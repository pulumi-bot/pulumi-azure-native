// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getReportConfig(args: GetReportConfigArgs, opts?: pulumi.InvokeOptions): Promise<GetReportConfigResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:costmanagement/latest:getReportConfig", {
        "reportConfigName": args.reportConfigName,
    }, opts);
}

export interface GetReportConfigArgs {
    /**
     * Report Config Name.
     */
    readonly reportConfigName: string;
}

/**
 * A report config resource.
 */
export interface GetReportConfigResult {
    /**
     * Has definition for the report config.
     */
    readonly definition: outputs.costmanagement.latest.ReportConfigDefinitionResponse;
    /**
     * Has delivery information for the report config.
     */
    readonly deliveryInfo: outputs.costmanagement.latest.ReportConfigDeliveryInfoResponse;
    /**
     * The format of the report being delivered.
     */
    readonly format?: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * Has schedule information for the report config.
     */
    readonly schedule?: outputs.costmanagement.latest.ReportConfigScheduleResponse;
    /**
     * Resource tags.
     */
    readonly tags: {[key: string]: string};
    /**
     * Resource type.
     */
    readonly type: string;
}
