// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getOriginGroup(args: GetOriginGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetOriginGroupResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:cdn/latest:getOriginGroup", {
        "endpointName": args.endpointName,
        "originGroupName": args.originGroupName,
        "profileName": args.profileName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetOriginGroupArgs {
    /**
     * Name of the endpoint under the profile which is unique globally.
     */
    readonly endpointName: string;
    /**
     * Name of the origin group which is unique within the endpoint.
     */
    readonly originGroupName: string;
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
 * Origin group comprising of origins is used for load balancing to origins when the content cannot be served from CDN.
 */
export interface GetOriginGroupResult {
    /**
     * Health probe settings to the origin that is used to determine the health of the origin.
     */
    readonly healthProbeSettings?: outputs.cdn.latest.HealthProbeParametersResponse;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * The source of the content being delivered via CDN within given origin group.
     */
    readonly origins: outputs.cdn.latest.ResourceReferenceResponse[];
    /**
     * Provisioning status of the origin group.
     */
    readonly provisioningState: string;
    /**
     * Resource status of the origin group.
     */
    readonly resourceState: string;
    /**
     * The JSON object that contains the properties to determine origin health using real requests/responses. This property is currently not supported.
     */
    readonly responseBasedOriginErrorDetectionSettings?: outputs.cdn.latest.ResponseBasedOriginErrorDetectionParametersResponse;
    /**
     * Time in minutes to shift the traffic to the endpoint gradually when an unhealthy endpoint comes healthy or a new endpoint is added. Default is 10 mins. This property is currently not supported.
     */
    readonly trafficRestorationTimeToHealedOrNewEndpointsInMinutes?: number;
    /**
     * Resource type.
     */
    readonly type: string;
}
