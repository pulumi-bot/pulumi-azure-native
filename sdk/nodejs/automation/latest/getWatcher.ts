// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getWatcher(args: GetWatcherArgs, opts?: pulumi.InvokeOptions): Promise<GetWatcherResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:automation/latest:getWatcher", {
        "automationAccountName": args.automationAccountName,
        "resourceGroupName": args.resourceGroupName,
        "watcherName": args.watcherName,
    }, opts);
}

export interface GetWatcherArgs {
    /**
     * The name of the automation account.
     */
    readonly automationAccountName: string;
    /**
     * Name of an Azure Resource group.
     */
    readonly resourceGroupName: string;
    /**
     * The watcher name.
     */
    readonly watcherName: string;
}

/**
 * Definition of the watcher type.
 */
export interface GetWatcherResult {
    /**
     * Gets or sets the creation time.
     */
    readonly creationTime: string;
    /**
     * Gets or sets the description.
     */
    readonly description?: string;
    /**
     * Gets or sets the etag of the resource.
     */
    readonly etag?: string;
    /**
     * Gets or sets the frequency at which the watcher is invoked.
     */
    readonly executionFrequencyInSeconds?: number;
    /**
     * Details of the user who last modified the watcher.
     */
    readonly lastModifiedBy: string;
    /**
     * Gets or sets the last modified time.
     */
    readonly lastModifiedTime: string;
    /**
     * The Azure Region where the resource lives
     */
    readonly location?: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * Gets or sets the name of the script the watcher is attached to, i.e. the name of an existing runbook.
     */
    readonly scriptName?: string;
    /**
     * Gets or sets the parameters of the script.
     */
    readonly scriptParameters?: {[key: string]: string};
    /**
     * Gets or sets the name of the hybrid worker group the watcher will run on.
     */
    readonly scriptRunOn?: string;
    /**
     * Gets the current status of the watcher.
     */
    readonly status: string;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource.
     */
    readonly type: string;
}
