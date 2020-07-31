// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getWebAppVnetConnectionSlot(args: GetWebAppVnetConnectionSlotArgs, opts?: pulumi.InvokeOptions): Promise<GetWebAppVnetConnectionSlotResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:web/v20190801:getWebAppVnetConnectionSlot", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
        "slot": args.slot,
    }, opts);
}

export interface GetWebAppVnetConnectionSlotArgs {
    /**
     * Name of the virtual network.
     */
    readonly name: string;
    /**
     * Name of the resource group to which the resource belongs.
     */
    readonly resourceGroupName: string;
    /**
     * Name of the deployment slot. If a slot is not specified, the API will get the named virtual network for the production slot.
     */
    readonly slot: string;
}

/**
 * Virtual Network information contract.
 */
export interface GetWebAppVnetConnectionSlotResult {
    /**
     * Kind of resource.
     */
    readonly kind?: string;
    /**
     * Resource Name.
     */
    readonly name: string;
    /**
     * VnetInfo resource specific properties
     */
    readonly properties: outputs.web.v20190801.VnetInfoResponseProperties;
    /**
     * Resource type.
     */
    readonly type: string;
}