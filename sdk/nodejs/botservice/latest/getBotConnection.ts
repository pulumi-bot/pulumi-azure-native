// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getBotConnection(args: GetBotConnectionArgs, opts?: pulumi.InvokeOptions): Promise<GetBotConnectionResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:botservice/latest:getBotConnection", {
        "connectionName": args.connectionName,
        "resourceGroupName": args.resourceGroupName,
        "resourceName": args.resourceName,
    }, opts);
}

export interface GetBotConnectionArgs {
    /**
     * The name of the Bot Service Connection Setting resource.
     */
    readonly connectionName: string;
    /**
     * The name of the Bot resource group in the user subscription.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the Bot resource.
     */
    readonly resourceName: string;
}

/**
 * Bot channel resource definition
 */
export interface GetBotConnectionResult {
    /**
     * Entity Tag
     */
    readonly etag?: string;
    /**
     * Required. Gets or sets the Kind of the resource.
     */
    readonly kind?: string;
    /**
     * Specifies the location of the resource.
     */
    readonly location?: string;
    /**
     * Specifies the name of the resource.
     */
    readonly name: string;
    /**
     * The set of properties specific to bot channel resource
     */
    readonly properties: outputs.botservice.latest.ConnectionSettingPropertiesResponse;
    /**
     * Gets or sets the SKU of the resource.
     */
    readonly sku?: outputs.botservice.latest.SkuResponse;
    /**
     * Contains resource tags defined as key/value pairs.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Specifies the type of the resource.
     */
    readonly type: string;
}
