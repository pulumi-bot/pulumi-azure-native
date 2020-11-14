// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getVirtualMachineSchedule(args: GetVirtualMachineScheduleArgs, opts?: pulumi.InvokeOptions): Promise<GetVirtualMachineScheduleResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:devtestlab/v20180915:getVirtualMachineSchedule", {
        "expand": args.expand,
        "labName": args.labName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
        "virtualMachineName": args.virtualMachineName,
    }, opts);
}

export interface GetVirtualMachineScheduleArgs {
    /**
     * Specify the $expand query. Example: 'properties($select=status)'
     */
    readonly expand?: string;
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
    /**
     * The name of the virtual machine.
     */
    readonly virtualMachineName: string;
}

/**
 * A schedule.
 */
export interface GetVirtualMachineScheduleResult {
    /**
     * The creation date of the schedule.
     */
    readonly createdDate: string;
    /**
     * If the schedule will occur once each day of the week, specify the daily recurrence.
     */
    readonly dailyRecurrence?: outputs.devtestlab.v20180915.DayDetailsResponse;
    /**
     * If the schedule will occur multiple times a day, specify the hourly recurrence.
     */
    readonly hourlyRecurrence?: outputs.devtestlab.v20180915.HourDetailsResponse;
    /**
     * The location of the resource.
     */
    readonly location?: string;
    /**
     * The name of the resource.
     */
    readonly name: string;
    /**
     * Notification settings.
     */
    readonly notificationSettings?: outputs.devtestlab.v20180915.NotificationSettingsResponse;
    /**
     * The provisioning status of the resource.
     */
    readonly provisioningState: string;
    /**
     * The status of the schedule (i.e. Enabled, Disabled)
     */
    readonly status?: string;
    /**
     * The tags of the resource.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The resource ID to which the schedule belongs
     */
    readonly targetResourceId?: string;
    /**
     * The task type of the schedule (e.g. LabVmsShutdownTask, LabVmAutoStart).
     */
    readonly taskType?: string;
    /**
     * The time zone ID (e.g. Pacific Standard time).
     */
    readonly timeZoneId?: string;
    /**
     * The type of the resource.
     */
    readonly type: string;
    /**
     * The unique immutable identifier of a resource (Guid).
     */
    readonly uniqueIdentifier: string;
    /**
     * If the schedule will occur only some days of the week, specify the weekly recurrence.
     */
    readonly weeklyRecurrence?: outputs.devtestlab.v20180915.WeekDetailsResponse;
}
