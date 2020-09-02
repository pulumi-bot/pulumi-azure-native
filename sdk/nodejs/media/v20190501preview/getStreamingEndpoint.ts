// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getStreamingEndpoint(args: GetStreamingEndpointArgs, opts?: pulumi.InvokeOptions): Promise<GetStreamingEndpointResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:media/v20190501preview:getStreamingEndpoint", {
        "accountName": args.accountName,
        "resourceGroupName": args.resourceGroupName,
        "streamingEndpointName": args.streamingEndpointName,
    }, opts);
}

export interface GetStreamingEndpointArgs {
    /**
     * The Media Services account name.
     */
    readonly accountName: string;
    /**
     * The name of the resource group within the Azure subscription.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the StreamingEndpoint.
     */
    readonly streamingEndpointName: string;
}

/**
 * The StreamingEndpoint.
 */
export interface GetStreamingEndpointResult {
    /**
     * The access control definition of the StreamingEndpoint.
     */
    readonly accessControl?: outputs.media.v20190501preview.StreamingEndpointAccessControlResponse;
    /**
     * The name of the AvailabilitySet used with this StreamingEndpoint for high availability streaming.  This value can only be set at creation time.
     */
    readonly availabilitySetName?: string;
    /**
     * The CDN enabled flag.
     */
    readonly cdnEnabled?: boolean;
    /**
     * The CDN profile name.
     */
    readonly cdnProfile?: string;
    /**
     * The CDN provider name.
     */
    readonly cdnProvider?: string;
    /**
     * The exact time the StreamingEndpoint was created.
     */
    readonly created: string;
    /**
     * The StreamingEndpoint access policies.
     */
    readonly crossSiteAccessPolicies?: outputs.media.v20190501preview.CrossSiteAccessPoliciesResponse;
    /**
     * The custom host names of the StreamingEndpoint
     */
    readonly customHostNames?: string[];
    /**
     * The StreamingEndpoint description.
     */
    readonly description?: string;
    /**
     * The free trial expiration time.
     */
    readonly freeTrialEndTime: string;
    /**
     * The StreamingEndpoint host name.
     */
    readonly hostName: string;
    /**
     * The exact time the StreamingEndpoint was last modified.
     */
    readonly lastModified: string;
    /**
     * The Azure Region of the resource.
     */
    readonly location?: string;
    /**
     * Max cache age
     */
    readonly maxCacheAge?: number;
    /**
     * The name of the resource.
     */
    readonly name: string;
    /**
     * The provisioning state of the StreamingEndpoint.
     */
    readonly provisioningState: string;
    /**
     * The resource state of the StreamingEndpoint.
     */
    readonly resourceState: string;
    /**
     * The number of scale units.  Use the Scale operation to adjust this value.
     */
    readonly scaleUnits: number;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource.
     */
    readonly type: string;
}
