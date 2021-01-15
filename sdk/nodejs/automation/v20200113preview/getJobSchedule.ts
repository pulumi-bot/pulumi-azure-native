// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

export function getJobSchedule(args: GetJobScheduleArgs, opts?: pulumi.InvokeOptions): Promise<GetJobScheduleResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:automation/v20200113preview:getJobSchedule", {
        "automationAccountName": args.automationAccountName,
        "jobScheduleId": args.jobScheduleId,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetJobScheduleArgs {
    /**
     * The name of the automation account.
     */
    readonly automationAccountName: string;
    /**
     * The job schedule name.
     */
    readonly jobScheduleId: string;
    /**
     * Name of an Azure Resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * Definition of the job schedule.
 */
export interface GetJobScheduleResult {
    /**
     * Gets the id of the resource.
     */
    readonly id: string;
    /**
     * Gets or sets the id of job schedule.
     */
    readonly jobScheduleId?: string;
    /**
     * Gets the name of the variable.
     */
    readonly name: string;
    /**
     * Gets or sets the parameters of the job schedule.
     */
    readonly parameters?: {[key: string]: string};
    /**
     * Gets or sets the hybrid worker group that the scheduled job should run on.
     */
    readonly runOn?: string;
    /**
     * Gets or sets the runbook.
     */
    readonly runbook?: outputs.automation.v20200113preview.RunbookAssociationPropertyResponse;
    /**
     * Gets or sets the schedule.
     */
    readonly schedule?: outputs.automation.v20200113preview.ScheduleAssociationPropertyResponse;
    /**
     * Resource type
     */
    readonly type: string;
}
