// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getAutomationAccountWatcher(args: GetAutomationAccountWatcherArgs, opts?: pulumi.InvokeOptions): Promise<GetAutomationAccountWatcherResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:automation:getAutomationAccountWatcher", {
        "automationAccountName": args.automationAccountName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetAutomationAccountWatcherArgs {
    /**
     * The name of the automation account.
     */
    readonly automationAccountName: string;
    /**
     * The watcher name.
     */
    readonly name: string;
    /**
     * Name of an Azure Resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * Definition of the watcher type.
 */
export interface GetAutomationAccountWatcherResult {
    /**
     * Gets or sets the etag of the resource.
     */
    readonly etag?: string;
    /**
     * The Azure Region where the resource lives
     */
    readonly location?: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * Gets or sets the watcher properties.
     */
    readonly properties: outputs.automation.WatcherPropertiesResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource.
     */
    readonly type: string;
}