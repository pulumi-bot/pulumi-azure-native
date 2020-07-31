// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getHostingEnvironment(args: GetHostingEnvironmentArgs, opts?: pulumi.InvokeOptions): Promise<GetHostingEnvironmentResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:web/v20150801:getHostingEnvironment", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetHostingEnvironmentArgs {
    /**
     * Name of hostingEnvironment (App Service Environment)
     */
    readonly name: string;
    /**
     * Name of resource group
     */
    readonly resourceGroupName: string;
}

/**
 * Description of an hostingEnvironment (App Service Environment)
 */
export interface GetHostingEnvironmentResult {
    /**
     * Kind of resource
     */
    readonly kind?: string;
    /**
     * Resource Location
     */
    readonly location: string;
    /**
     * Resource Name
     */
    readonly name?: string;
    readonly properties: outputs.web.v20150801.HostingEnvironmentResponseProperties;
    /**
     * Resource tags
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type
     */
    readonly type?: string;
}