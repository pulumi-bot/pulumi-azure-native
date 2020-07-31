// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getSignalRPrivateEndpointConnection(args: GetSignalRPrivateEndpointConnectionArgs, opts?: pulumi.InvokeOptions): Promise<GetSignalRPrivateEndpointConnectionResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:signalrservice/v20200501:getSignalRPrivateEndpointConnection", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
        "resourceName": args.resourceName,
    }, opts);
}

export interface GetSignalRPrivateEndpointConnectionArgs {
    /**
     * The name of the private endpoint connection associated with the SignalR resource.
     */
    readonly name: string;
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the SignalR resource.
     */
    readonly resourceName: string;
}

/**
 * A private endpoint connection to SignalR resource
 */
export interface GetSignalRPrivateEndpointConnectionResult {
    /**
     * The name of the resource.
     */
    readonly name: string;
    /**
     * Properties of the private endpoint connection
     */
    readonly properties: outputs.signalrservice.v20200501.PrivateEndpointConnectionPropertiesResponse;
    /**
     * The type of the resource - e.g. "Microsoft.SignalRService/SignalR"
     */
    readonly type: string;
}