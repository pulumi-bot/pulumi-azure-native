// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getVirtualNetworkTap(args: GetVirtualNetworkTapArgs, opts?: pulumi.InvokeOptions): Promise<GetVirtualNetworkTapResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:network/latest:getVirtualNetworkTap", {
        "resourceGroupName": args.resourceGroupName,
        "tapName": args.tapName,
    }, opts);
}

export interface GetVirtualNetworkTapArgs {
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
    /**
     * The name of virtual network tap.
     */
    readonly tapName: string;
}

/**
 * Virtual Network Tap resource.
 */
export interface GetVirtualNetworkTapResult {
    /**
     * The reference to the private IP address on the internal Load Balancer that will receive the tap.
     */
    readonly destinationLoadBalancerFrontEndIPConfiguration?: outputs.network.latest.FrontendIPConfigurationResponse;
    /**
     * The reference to the private IP Address of the collector nic that will receive the tap.
     */
    readonly destinationNetworkInterfaceIPConfiguration?: outputs.network.latest.NetworkInterfaceIPConfigurationResponse;
    /**
     * The VXLAN destination port that will receive the tapped traffic.
     */
    readonly destinationPort?: number;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    readonly etag: string;
    /**
     * Resource location.
     */
    readonly location?: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * Specifies the list of resource IDs for the network interface IP configuration that needs to be tapped.
     */
    readonly networkInterfaceTapConfigurations: outputs.network.latest.NetworkInterfaceTapConfigurationResponse[];
    /**
     * The provisioning state of the virtual network tap resource.
     */
    readonly provisioningState: string;
    /**
     * The resource GUID property of the virtual network tap resource.
     */
    readonly resourceGuid: string;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type.
     */
    readonly type: string;
}