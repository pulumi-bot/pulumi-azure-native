// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getAuthorization(args: GetAuthorizationArgs, opts?: pulumi.InvokeOptions): Promise<GetAuthorizationResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:avs/v20200717preview:getAuthorization", {
        "authorizationName": args.authorizationName,
        "privateCloudName": args.privateCloudName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetAuthorizationArgs {
    /**
     * Name of the ExpressRoute Circuit Authorization in the private cloud
     */
    readonly authorizationName: string;
    /**
     * Name of the private cloud
     */
    readonly privateCloudName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: string;
}

/**
 * ExpressRoute Circuit Authorization
 */
export interface GetAuthorizationResult {
    /**
     * The ID of the ExpressRoute Circuit Authorization
     */
    readonly expressRouteAuthorizationId: string;
    /**
     * The key of the ExpressRoute Circuit Authorization
     */
    readonly expressRouteAuthorizationKey: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * The state of the  ExpressRoute Circuit Authorization provisioning
     */
    readonly provisioningState: string;
    /**
     * Resource type.
     */
    readonly type: string;
}
