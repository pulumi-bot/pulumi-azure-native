// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getJitRequest(args: GetJitRequestArgs, opts?: pulumi.InvokeOptions): Promise<GetJitRequestResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:solutions:getJitRequest", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetJitRequestArgs {
    /**
     * The name of the JIT request.
     */
    readonly name: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: string;
}

/**
 * Information about JIT request definition.
 */
export interface GetJitRequestResult {
    /**
     * Resource location
     */
    readonly location?: string;
    /**
     * Resource name
     */
    readonly name: string;
    /**
     * The JIT request properties.
     */
    readonly properties: outputs.solutions.JitRequestPropertiesResponse;
    /**
     * Resource tags
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type
     */
    readonly type: string;
}