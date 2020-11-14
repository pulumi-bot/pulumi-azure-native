// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function listSiteMetadataSlot(args: ListSiteMetadataSlotArgs, opts?: pulumi.InvokeOptions): Promise<ListSiteMetadataSlotResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:web/v20150801:listSiteMetadataSlot", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
        "slot": args.slot,
    }, opts);
}

export interface ListSiteMetadataSlotArgs {
    /**
     * Name of web app
     */
    readonly name: string;
    /**
     * Name of resource group
     */
    readonly resourceGroupName: string;
    /**
     * Name of web app slot. If not specified then will default to production slot.
     */
    readonly slot: string;
}

/**
 * String dictionary resource
 */
export interface ListSiteMetadataSlotResult {
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
    /**
     * Settings
     */
    readonly properties: {[key: string]: string};
    /**
     * Resource tags
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type
     */
    readonly type?: string;
}
