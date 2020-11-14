// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getVirtualNetwork(args: GetVirtualNetworkArgs, opts?: pulumi.InvokeOptions): Promise<GetVirtualNetworkResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:network/v20171101:getVirtualNetwork", {
        "expand": args.expand,
        "resourceGroupName": args.resourceGroupName,
        "virtualNetworkName": args.virtualNetworkName,
    }, opts);
}

export interface GetVirtualNetworkArgs {
    /**
     * Expands referenced resources.
     */
    readonly expand?: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the virtual network.
     */
    readonly virtualNetworkName: string;
}

/**
 * Virtual Network resource.
 */
export interface GetVirtualNetworkResult {
    /**
     * The AddressSpace that contains an array of IP address ranges that can be used by subnets.
     */
    readonly addressSpace?: outputs.network.v20171101.AddressSpaceResponse;
    /**
     * The dhcpOptions that contains an array of DNS servers available to VMs deployed in the virtual network.
     */
    readonly dhcpOptions?: outputs.network.v20171101.DhcpOptionsResponse;
    /**
     * Indicates if DDoS protection is enabled for all the protected resources in a Virtual Network.
     */
    readonly enableDdosProtection?: boolean;
    /**
     * Indicates if Vm protection is enabled for all the subnets in a Virtual Network.
     */
    readonly enableVmProtection?: boolean;
    /**
     * Gets a unique read-only string that changes whenever the resource is updated.
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
     * The provisioning state of the PublicIP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
     */
    readonly provisioningState?: string;
    /**
     * The resourceGuid property of the Virtual Network resource.
     */
    readonly resourceGuid?: string;
    /**
     * A list of subnets in a Virtual Network.
     */
    readonly subnets?: outputs.network.v20171101.SubnetResponse[];
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type.
     */
    readonly type: string;
    /**
     * A list of peerings in a Virtual Network.
     */
    readonly virtualNetworkPeerings?: outputs.network.v20171101.VirtualNetworkPeeringResponse[];
}
