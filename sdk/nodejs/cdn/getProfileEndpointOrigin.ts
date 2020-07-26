// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getProfileEndpointOrigin(args: GetProfileEndpointOriginArgs, opts?: pulumi.InvokeOptions): Promise<GetProfileEndpointOriginResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:cdn:getProfileEndpointOrigin", {
        "endpointName": args.endpointName,
        "name": args.name,
        "profileName": args.profileName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetProfileEndpointOriginArgs {
    /**
     * Name of the endpoint under the profile which is unique globally.
     */
    readonly endpointName: string;
    /**
     * Name of the origin which is unique within the endpoint.
     */
    readonly name: string;
    /**
     * Name of the CDN profile which is unique within the resource group.
     */
    readonly profileName: string;
    /**
     * Name of the Resource group within the Azure subscription.
     */
    readonly resourceGroupName: string;
}

/**
 * CDN origin is the source of the content being delivered via CDN. When the edge nodes represented by an endpoint do not have the requested content cached, they attempt to fetch it from one or more of the configured origins.
 */
export interface GetProfileEndpointOriginResult {
    /**
     * Resource location.
     */
    readonly location: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * The JSON object that contains the properties of the origin.
     */
    readonly properties: outputs.cdn.OriginPropertiesResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type.
     */
    readonly type: string;
}
