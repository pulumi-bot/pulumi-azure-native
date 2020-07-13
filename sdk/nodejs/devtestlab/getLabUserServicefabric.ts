// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getLabUserServicefabric(args: GetLabUserServicefabricArgs, opts?: pulumi.InvokeOptions): Promise<GetLabUserServicefabricResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:devtestlab:getLabUserServicefabric", {
        "labName": args.labName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
        "userName": args.userName,
    }, opts);
}

export interface GetLabUserServicefabricArgs {
    /**
     * The name of the lab.
     */
    readonly labName: string;
    /**
     * The name of the service fabric.
     */
    readonly name: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the user profile.
     */
    readonly userName: string;
}

/**
 * A Service Fabric.
 */
export interface GetLabUserServicefabricResult {
    /**
     * The location of the resource.
     */
    readonly location?: string;
    /**
     * The name of the resource.
     */
    readonly name: string;
    /**
     * The properties of the resource.
     */
    readonly properties: outputs.devtestlab.ServiceFabricPropertiesResponse;
    /**
     * The tags of the resource.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource.
     */
    readonly type: string;
}