// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getVirtualHubBgpConnection(args: GetVirtualHubBgpConnectionArgs, opts?: pulumi.InvokeOptions): Promise<GetVirtualHubBgpConnectionResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:network/v20200601:getVirtualHubBgpConnection", {
        "connectionName": args.connectionName,
        "resourceGroupName": args.resourceGroupName,
        "virtualHubName": args.virtualHubName,
    }, opts);
}

export interface GetVirtualHubBgpConnectionArgs {
    /**
     * The name of the connection.
     */
    readonly connectionName: string;
    /**
     * The resource group name of the VirtualHub.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the VirtualHub.
     */
    readonly virtualHubName: string;
}

/**
 * Virtual Appliance Site resource.
 */
export interface GetVirtualHubBgpConnectionResult {
    /**
     * The current state of the VirtualHub to Peer.
     */
    readonly connectionState: string;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    readonly etag: string;
    /**
     * Name of the connection.
     */
    readonly name?: string;
    /**
     * Peer ASN.
     */
    readonly peerAsn?: number;
    /**
     * Peer IP.
     */
    readonly peerIp?: string;
    /**
     * The provisioning state of the resource.
     */
    readonly provisioningState: string;
    /**
     * Connection type.
     */
    readonly type: string;
}