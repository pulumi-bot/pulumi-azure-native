// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function listSiteBackupConfigurationSlot(args: ListSiteBackupConfigurationSlotArgs, opts?: pulumi.InvokeOptions): Promise<ListSiteBackupConfigurationSlotResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:web/v20150801:listSiteBackupConfigurationSlot", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListSiteBackupConfigurationSlotArgs {
    /**
     * Name of web app slot. If not specified then will default to production slot.
     */
    readonly name: string;
    /**
     * Name of resource group
     */
    readonly resourceGroupName: string;
}

/**
 * Description of a backup which will be performed
 */
export interface ListSiteBackupConfigurationSlotResult {
    /**
     * Kind of resource
     */
    readonly kind?: string;
    /**
     * Resource Location
     */
    readonly location: string;
    /**
     * Resource Name
     */
    readonly name?: string;
    readonly properties: outputs.web.v20150801.BackupRequestResponseProperties;
    /**
     * Resource tags
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type
     */
    readonly type?: string;
}