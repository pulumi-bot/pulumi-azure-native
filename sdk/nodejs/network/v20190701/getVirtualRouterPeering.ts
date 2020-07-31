// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getVirtualRouterPeering(args: GetVirtualRouterPeeringArgs, opts?: pulumi.InvokeOptions): Promise<GetVirtualRouterPeeringResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:network/v20190701:getVirtualRouterPeering", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
        "virtualRouterName": args.virtualRouterName,
    }, opts);
}

export interface GetVirtualRouterPeeringArgs {
    /**
     * The name of the Virtual Router Peering.
     */
    readonly name: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the Virtual Router.
     */
    readonly virtualRouterName: string;
}

/**
 * Virtual Router Peering resource
 */
export interface GetVirtualRouterPeeringResult {
    /**
     * Gets a unique read-only string that changes whenever the resource is updated.
     */
    readonly etag: string;
    /**
     * Gets name of the peering unique to VirtualRouter. This name can be used to access the resource.
     */
    readonly name?: string;
    /**
     * The properties of the Virtual Router Peering.
     */
    readonly properties: outputs.network.v20190701.VirtualRouterPeeringPropertiesResponse;
    /**
     * Peering type.
     */
    readonly type: string;
}