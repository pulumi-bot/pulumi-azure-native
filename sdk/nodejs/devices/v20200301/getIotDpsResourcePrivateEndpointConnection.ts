// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getIotDpsResourcePrivateEndpointConnection(args: GetIotDpsResourcePrivateEndpointConnectionArgs, opts?: pulumi.InvokeOptions): Promise<GetIotDpsResourcePrivateEndpointConnectionResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:devices/v20200301:getIotDpsResourcePrivateEndpointConnection", {
        "privateEndpointConnectionName": args.privateEndpointConnectionName,
        "resourceGroupName": args.resourceGroupName,
        "resourceName": args.resourceName,
    }, opts);
}

export interface GetIotDpsResourcePrivateEndpointConnectionArgs {
    /**
     * The name of the private endpoint connection
     */
    readonly privateEndpointConnectionName: string;
    /**
     * The name of the resource group that contains the provisioning service.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the provisioning service.
     */
    readonly resourceName: string;
}

/**
 * The private endpoint connection of a provisioning service
 */
export interface GetIotDpsResourcePrivateEndpointConnectionResult {
    /**
     * The resource name.
     */
    readonly name: string;
    /**
     * The properties of a private endpoint connection
     */
    readonly properties: outputs.devices.v20200301.PrivateEndpointConnectionPropertiesResponse;
    /**
     * The resource type.
     */
    readonly type: string;
}
