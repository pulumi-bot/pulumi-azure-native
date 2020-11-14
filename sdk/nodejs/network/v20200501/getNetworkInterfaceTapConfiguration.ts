// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getNetworkInterfaceTapConfiguration(args: GetNetworkInterfaceTapConfigurationArgs, opts?: pulumi.InvokeOptions): Promise<GetNetworkInterfaceTapConfigurationResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:network/v20200501:getNetworkInterfaceTapConfiguration", {
        "networkInterfaceName": args.networkInterfaceName,
        "resourceGroupName": args.resourceGroupName,
        "tapConfigurationName": args.tapConfigurationName,
    }, opts);
}

export interface GetNetworkInterfaceTapConfigurationArgs {
    /**
     * The name of the network interface.
     */
    readonly networkInterfaceName: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the tap configuration.
     */
    readonly tapConfigurationName: string;
}

/**
 * Tap configuration in a Network Interface.
 */
export interface GetNetworkInterfaceTapConfigurationResult {
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    readonly etag: string;
    /**
     * The name of the resource that is unique within a resource group. This name can be used to access the resource.
     */
    readonly name?: string;
    /**
     * The provisioning state of the network interface tap configuration resource.
     */
    readonly provisioningState: string;
    /**
     * Sub Resource type.
     */
    readonly type: string;
    /**
     * The reference to the Virtual Network Tap resource.
     */
    readonly virtualNetworkTap?: outputs.network.v20200501.VirtualNetworkTapResponse;
}
