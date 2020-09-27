// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getSchedule(args: GetScheduleArgs, opts?: pulumi.InvokeOptions): Promise<GetScheduleResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:automation/v20151031:getSchedule", {
        "automationAccountName": args.automationAccountName,
        "resourceGroupName": args.resourceGroupName,
        "scheduleName": args.scheduleName,
    }, opts);
}

export interface GetScheduleArgs {
    /**
     * The name of the automation account.
     */
    readonly automationAccountName: string;
    /**
     * Name of an Azure Resource group.
     */
    readonly resourceGroupName: string;
    /**
     * The schedule name.
     */
    readonly scheduleName: string;
}

/**
 * Definition of the schedule.
 */
export interface GetScheduleResult {
    /**
     * Gets or sets the advanced schedule.
     */
    readonly advancedSchedule?: outputs.automation.v20151031.AdvancedScheduleResponse;
    /**
     * Gets or sets the creation time.
     */
    readonly creationTime?: string;
    /**
     * Gets or sets the description.
     */
    readonly description?: string;
    /**
     * Gets or sets the end time of the schedule.
     */
    readonly expiryTime?: string;
    /**
     * Gets or sets the expiry time's offset in minutes.
     */
    readonly expiryTimeOffsetMinutes?: number;
    /**
     * Gets or sets the frequency of the schedule.
     */
    readonly frequency?: string;
    /**
     * Gets or sets a value indicating whether this schedule is enabled.
     */
    readonly isEnabled?: boolean;
    /**
     * Gets or sets the last modified time.
     */
    readonly lastModifiedTime?: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * Gets or sets the next run time of the schedule.
     */
    readonly nextRun?: string;
    /**
     * Gets or sets the next run time's offset in minutes.
     */
    readonly nextRunOffsetMinutes?: number;
    /**
     * Gets or sets the start time of the schedule.
     */
    readonly startTime?: string;
    /**
     * Gets the start time's offset in minutes.
     */
    readonly startTimeOffsetMinutes: number;
    /**
     * Gets or sets the time zone of the schedule.
     */
    readonly timeZone?: string;
    /**
     * The type of the resource.
     */
    readonly type: string;
}
