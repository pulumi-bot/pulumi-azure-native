// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getResourceGroup(args: GetResourceGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetResourceGroupResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:resources/v20200601:getResourceGroup", {
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetResourceGroupArgs {
    /**
     * The name of the resource group to get. The name is case insensitive.
     */
    readonly resourceGroupName: string;
}

/**
 * Resource group information.
 */
export interface GetResourceGroupResult {
    /**
     * The location of the resource group. It cannot be changed after the resource group has been created. It must be one of the supported Azure locations.
     */
    readonly location: string;
    /**
     * The ID of the resource that manages this resource group.
     */
    readonly managedBy?: string;
    /**
     * The name of the resource group.
     */
    readonly name: string;
    /**
     * The resource group properties.
     */
    readonly properties: outputs.resources.v20200601.ResourceGroupPropertiesResponse;
    /**
     * The tags attached to the resource group.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource group.
     */
    readonly type: string;
}
