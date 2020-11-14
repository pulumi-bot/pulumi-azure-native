// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getPublicIPPrefix(args: GetPublicIPPrefixArgs, opts?: pulumi.InvokeOptions): Promise<GetPublicIPPrefixResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:network/v20180701:getPublicIPPrefix", {
        "expand": args.expand,
        "publicIpPrefixName": args.publicIpPrefixName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetPublicIPPrefixArgs {
    /**
     * Expands referenced resources.
     */
    readonly expand?: string;
    /**
     * The name of the Public IP Prefix.
     */
    readonly publicIpPrefixName: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * Public IP prefix resource.
 */
export interface GetPublicIPPrefixResult {
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    readonly etag?: string;
    /**
     * The allocated Prefix
     */
    readonly ipPrefix?: string;
    /**
     * The list of tags associated with the public IP prefix.
     */
    readonly ipTags?: outputs.network.v20180701.IpTagResponse[];
    /**
     * The reference to load balancer frontend IP configuration associated with the public IP prefix.
     */
    readonly loadBalancerFrontendIpConfiguration: outputs.network.v20180701.SubResourceResponse;
    /**
     * Resource location.
     */
    readonly location?: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * The Length of the Public IP Prefix.
     */
    readonly prefixLength?: number;
    /**
     * The provisioning state of the Public IP prefix resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
     */
    readonly provisioningState?: string;
    /**
     * The public IP address version. Possible values are: 'IPv4' and 'IPv6'.
     */
    readonly publicIPAddressVersion?: string;
    /**
     * The list of all referenced PublicIPAddresses
     */
    readonly publicIPAddresses?: outputs.network.v20180701.ReferencedPublicIpAddressResponse[];
    /**
     * The resource GUID property of the public IP prefix resource.
     */
    readonly resourceGuid?: string;
    /**
     * The public IP prefix SKU.
     */
    readonly sku?: outputs.network.v20180701.PublicIPPrefixSkuResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type.
     */
    readonly type: string;
    /**
     * A list of availability zones denoting the IP allocated for the resource needs to come from.
     */
    readonly zones?: string[];
}
