// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getP2sVpnGatewayP2sVpnConnectionHealth(args: GetP2sVpnGatewayP2sVpnConnectionHealthArgs, opts?: pulumi.InvokeOptions): Promise<GetP2sVpnGatewayP2sVpnConnectionHealthResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:network/v20190701:getP2sVpnGatewayP2sVpnConnectionHealth", {
        "gatewayName": args.gatewayName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetP2sVpnGatewayP2sVpnConnectionHealthArgs {
    /**
     * The name of the P2SVpnGateway.
     */
    readonly gatewayName: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * P2SVpnGateway Resource.
 */
export interface GetP2sVpnGatewayP2sVpnConnectionHealthResult {
    /**
     * The reference of the address space resource which represents the custom routes specified by the customer for P2SVpnGateway and P2S VpnClient.
     */
    readonly customRoutes?: outputs.network.v20190701.AddressSpaceResponse;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    readonly etag: string;
    /**
     * Resource location.
     */
    readonly location: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * The P2SVpnServerConfiguration to which the p2sVpnGateway is attached to.
     */
    readonly p2SVpnServerConfiguration?: outputs.network.v20190701.SubResourceResponse;
    /**
     * The provisioning state of the P2S VPN gateway resource.
     */
    readonly provisioningState: string;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type.
     */
    readonly type: string;
    /**
     * The VirtualHub to which the gateway belongs.
     */
    readonly virtualHub?: outputs.network.v20190701.SubResourceResponse;
    /**
     * The reference of the address space resource which represents Address space for P2S VpnClient.
     */
    readonly vpnClientAddressPool?: outputs.network.v20190701.AddressSpaceResponse;
    /**
     * All P2S VPN clients' connection health status.
     */
    readonly vpnClientConnectionHealth: outputs.network.v20190701.VpnClientConnectionHealthResponse;
    /**
     * The scale unit for this p2s vpn gateway.
     */
    readonly vpnGatewayScaleUnit?: number;
}
