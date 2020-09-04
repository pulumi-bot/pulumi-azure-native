// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getPublicIpAddress(args: GetPublicIpAddressArgs, opts?: pulumi.InvokeOptions): Promise<GetPublicIpAddressResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:network/v20150501preview:getPublicIpAddress", {
        "publicIpAddressName": args.publicIpAddressName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetPublicIpAddressArgs {
    /**
     * The name of the subnet.
     */
    readonly publicIpAddressName: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * PublicIPAddress resource
 */
export interface GetPublicIpAddressResult {
    /**
     * Gets or sets FQDN of the DNS record associated with the public IP address
     */
    readonly dnsSettings?: outputs.network.v20150501preview.PublicIpAddressDnsSettingsResponse;
    /**
     * Gets a unique read-only string that changes whenever the resource is updated
     */
    readonly etag?: string;
    /**
     * Gets or sets the idle timeout of the public IP address
     */
    readonly idleTimeoutInMinutes?: number;
    /**
     * Gets the assigned public IP address
     */
    readonly ipAddress?: string;
    /**
     * Gets a reference to the network interface IP configurations using this public IP address
     */
    readonly ipConfiguration?: outputs.network.v20150501preview.SubResourceResponse;
    /**
     * Resource location
     */
    readonly location: string;
    /**
     * Resource name
     */
    readonly name: string;
    /**
     * Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed
     */
    readonly provisioningState?: string;
    /**
     * Gets or sets PublicIP allocation method (Static/Dynamic)
     */
    readonly publicIPAllocationMethod: string;
    /**
     * Gets or sets resource guid property of the PublicIP resource
     */
    readonly resourceGuid?: string;
    /**
     * Resource tags
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type
     */
    readonly type: string;
}
