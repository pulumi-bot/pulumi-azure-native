// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getScheduleResource(args: GetScheduleResourceArgs, opts?: pulumi.InvokeOptions): Promise<GetScheduleResourceResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:devtestlab/v20150521preview:getScheduleResource", {
        "labName": args.labName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetScheduleResourceArgs {
    /**
     * The name of the lab.
     */
    readonly labName: string;
    /**
     * The name of the schedule.
     */
    readonly name: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * A schedule.
 */
export interface GetScheduleResourceResult {
    /**
     * The daily recurrence of the schedule.
     */
    readonly dailyRecurrence?: outputs.devtestlab.v20150521preview.DayDetailsResponse;
    /**
     * The hourly recurrence of the schedule.
     */
    readonly hourlyRecurrence?: outputs.devtestlab.v20150521preview.HourDetailsResponse;
    /**
     * The location of the resource.
     */
    readonly location?: string;
    /**
     * The name of the resource.
     */
    readonly name?: string;
    /**
     * The provisioning status of the resource.
     */
    readonly provisioningState?: string;
    /**
     * The status of the schedule.
     */
    readonly status?: string;
    /**
     * The tags of the resource.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The task type of the schedule.
     */
    readonly taskType?: string;
    /**
     * The time zone id.
     */
    readonly timeZoneId?: string;
    /**
     * The type of the resource.
     */
    readonly type?: string;
    /**
     * The weekly recurrence of the schedule.
     */
    readonly weeklyRecurrence?: outputs.devtestlab.v20150521preview.WeekDetailsResponse;
}
