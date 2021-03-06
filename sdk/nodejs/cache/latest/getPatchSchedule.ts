// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Response to put/get patch schedules for Redis cache.
 * Latest API Version: 2020-06-01.
 */
/** @deprecated The 'latest' version is deprecated. Please migrate to the function in the top-level module: 'azure-native:cache:getPatchSchedule'. */
export function getPatchSchedule(args: GetPatchScheduleArgs, opts?: pulumi.InvokeOptions): Promise<GetPatchScheduleResult> {
    pulumi.log.warn("getPatchSchedule is deprecated: The 'latest' version is deprecated. Please migrate to the function in the top-level module: 'azure-native:cache:getPatchSchedule'.")
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:cache/latest:getPatchSchedule", {
        "default": args.default,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetPatchScheduleArgs {
    /**
     * Default string modeled as parameter for auto generation to work correctly.
     */
    readonly default: string;
    /**
     * The name of the redis cache.
     */
    readonly name: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * Response to put/get patch schedules for Redis cache.
 */
export interface GetPatchScheduleResult {
    /**
     * Resource ID.
     */
    readonly id: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * List of patch schedules for a Redis cache.
     */
    readonly scheduleEntries: outputs.cache.latest.ScheduleEntryResponse[];
    /**
     * Resource type.
     */
    readonly type: string;
}
