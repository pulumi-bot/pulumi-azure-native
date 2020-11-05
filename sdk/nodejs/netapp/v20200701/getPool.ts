// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getPool(args: GetPoolArgs, opts?: pulumi.InvokeOptions): Promise<GetPoolResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:netapp/v20200701:getPool", {
        "accountName": args.accountName,
        "poolName": args.poolName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetPoolArgs {
    /**
     * The name of the NetApp account
     */
    readonly accountName: string;
    /**
     * The name of the capacity pool
     */
    readonly poolName: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * Capacity pool resource
 */
export interface GetPoolResult {
    /**
     * Resource location
     */
    readonly location: string;
    /**
     * Resource name
     */
    readonly name: string;
    /**
     * UUID v4 used to identify the Pool
     */
    readonly poolId: string;
    /**
     * Azure lifecycle management
     */
    readonly provisioningState: string;
    /**
     * The qos type of the pool
     */
    readonly qosType?: string;
    /**
     * The service level of the file system
     */
    readonly serviceLevel: string;
    /**
     * Provisioned size of the pool (in bytes). Allowed values are in 4TiB chunks (value must be multiply of 4398046511104).
     */
    readonly size: number;
    /**
     * Resource tags
     */
    readonly tags?: {[key: string]: string};
    /**
     * Total throughput of pool in Mibps
     */
    readonly totalThroughputMibps: number;
    /**
     * Resource type
     */
    readonly type: string;
    /**
     * Utilized throughput of pool in Mibps
     */
    readonly utilizedThroughputMibps: number;
}
