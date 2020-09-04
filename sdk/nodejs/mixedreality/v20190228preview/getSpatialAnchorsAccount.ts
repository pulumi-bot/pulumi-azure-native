// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getSpatialAnchorsAccount(args: GetSpatialAnchorsAccountArgs, opts?: pulumi.InvokeOptions): Promise<GetSpatialAnchorsAccountResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:mixedreality/v20190228preview:getSpatialAnchorsAccount", {
        "resourceGroupName": args.resourceGroupName,
        "spatialAnchorsAccountName": args.spatialAnchorsAccountName,
    }, opts);
}

export interface GetSpatialAnchorsAccountArgs {
    /**
     * Name of an Azure resource group.
     */
    readonly resourceGroupName: string;
    /**
     * Name of an Mixed Reality Spatial Anchors Account.
     */
    readonly spatialAnchorsAccountName: string;
}

/**
 * SpatialAnchorsAccount Response.
 */
export interface GetSpatialAnchorsAccountResult {
    /**
     * Correspond domain name of certain Spatial Anchors Account
     */
    readonly accountDomain: string;
    /**
     * unique id of certain Spatial Anchors Account data contract.
     */
    readonly accountId: string;
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
     */
    readonly type: string;
}
