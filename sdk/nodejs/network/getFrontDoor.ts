// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getFrontDoor(args: GetFrontDoorArgs, opts?: pulumi.InvokeOptions): Promise<GetFrontDoorResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:network:getFrontDoor", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetFrontDoorArgs {
    /**
     * Name of the Front Door which is globally unique.
     */
    readonly name: string;
    /**
     * Name of the Resource group within the Azure subscription.
     */
    readonly resourceGroupName: string;
}

/**
 * Front Door represents a collection of backend endpoints to route traffic to along with rules that specify how traffic is sent there.
 */
export interface GetFrontDoorResult {
    /**
     * Resource location.
     */
    readonly location?: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * Properties of the Front Door Load Balancer
     */
    readonly properties: outputs.network.FrontDoorPropertiesResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type.
     */
    readonly type: string;
}