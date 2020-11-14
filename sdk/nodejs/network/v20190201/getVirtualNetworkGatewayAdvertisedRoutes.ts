// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getVirtualNetworkGatewayAdvertisedRoutes(args: GetVirtualNetworkGatewayAdvertisedRoutesArgs, opts?: pulumi.InvokeOptions): Promise<GetVirtualNetworkGatewayAdvertisedRoutesResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:network/v20190201:getVirtualNetworkGatewayAdvertisedRoutes", {
        "peer": args.peer,
        "resourceGroupName": args.resourceGroupName,
        "virtualNetworkGatewayName": args.virtualNetworkGatewayName,
    }, opts);
}

export interface GetVirtualNetworkGatewayAdvertisedRoutesArgs {
    /**
     * The IP address of the peer
     */
    readonly peer: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the virtual network gateway.
     */
    readonly virtualNetworkGatewayName: string;
}

/**
 * List of virtual network gateway routes
 */
export interface GetVirtualNetworkGatewayAdvertisedRoutesResult {
    /**
     * List of gateway routes
     */
    readonly value?: outputs.network.v20190201.GatewayRouteResponse[];
}
