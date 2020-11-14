// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getVpnConnection(args: GetVpnConnectionArgs, opts?: pulumi.InvokeOptions): Promise<GetVpnConnectionResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:network/v20180601:getVpnConnection", {
        "connectionName": args.connectionName,
        "gatewayName": args.gatewayName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetVpnConnectionArgs {
    /**
     * The name of the vpn connection.
     */
    readonly connectionName: string;
    /**
     * The name of the gateway.
     */
    readonly gatewayName: string;
    /**
     * The resource group name of the VpnGateway.
     */
    readonly resourceGroupName: string;
}

/**
 * VpnConnection Resource.
 */
export interface GetVpnConnectionResult {
    /**
     * Expected bandwidth in MBPS.
     */
    readonly connectionBandwidthInMbps: number;
    /**
     * The connection status.
     */
    readonly connectionStatus: string;
    /**
     * Egress bytes transferred.
     */
    readonly egressBytesTransferred: number;
    /**
     * EnableBgp flag
     */
    readonly enableBgp?: boolean;
    /**
     * Gets a unique read-only string that changes whenever the resource is updated.
     */
    readonly etag: string;
    /**
     * Ingress bytes transferred.
     */
    readonly ingressBytesTransferred: number;
    /**
     * The IPSec Policies to be considered by this connection.
     */
    readonly ipsecPolicies?: outputs.network.v20180601.IpsecPolicyResponse[];
    /**
     * The name of the resource that is unique within a resource group. This name can be used to access the resource.
     */
    readonly name?: string;
    /**
     * The provisioning state of the resource.
     */
    readonly provisioningState: string;
    /**
     * Id of the connected vpn site.
     */
    readonly remoteVpnSite?: outputs.network.v20180601.SubResourceResponse;
    /**
     * routing weight for vpn connection.
     */
    readonly routingWeight?: number;
    /**
     * SharedKey for the vpn connection.
     */
    readonly sharedKey?: string;
}
