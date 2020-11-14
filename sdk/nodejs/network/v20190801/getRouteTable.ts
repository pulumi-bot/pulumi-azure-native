// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getRouteTable(args: GetRouteTableArgs, opts?: pulumi.InvokeOptions): Promise<GetRouteTableResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:network/v20190801:getRouteTable", {
        "expand": args.expand,
        "resourceGroupName": args.resourceGroupName,
        "routeTableName": args.routeTableName,
    }, opts);
}

export interface GetRouteTableArgs {
    /**
     * Expands referenced resources.
     */
    readonly expand?: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the route table.
     */
    readonly routeTableName: string;
}

/**
 * Route table resource.
 */
export interface GetRouteTableResult {
    /**
     * Whether to disable the routes learned by BGP on that route table. True means disable.
     */
    readonly disableBgpRoutePropagation?: boolean;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    readonly etag?: string;
    /**
     * Resource location.
     */
    readonly location?: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * The provisioning state of the route table resource.
     */
    readonly provisioningState: string;
    /**
     * Collection of routes contained within a route table.
     */
    readonly routes?: outputs.network.v20190801.RouteResponse[];
    /**
     * A collection of references to subnets.
     */
    readonly subnets: outputs.network.v20190801.SubnetResponse[];
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type.
     */
    readonly type: string;
}
