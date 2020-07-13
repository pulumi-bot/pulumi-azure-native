// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getAppServicePremieraddon(args: GetAppServicePremieraddonArgs, opts?: pulumi.InvokeOptions): Promise<GetAppServicePremieraddonResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:web:getAppServicePremieraddon", {
        "name": args.name,
        "premierAddOnName": args.premierAddOnName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetAppServicePremieraddonArgs {
    /**
     * Name of the app.
     */
    readonly name: string;
    /**
     * Add-on name.
     */
    readonly premierAddOnName: string;
    /**
     * Name of the resource group to which the resource belongs.
     */
    readonly resourceGroupName: string;
}

/**
 * Premier add-on.
 */
export interface GetAppServicePremieraddonResult {
    /**
     * Kind of resource.
     */
    readonly kind?: string;
    /**
     * Resource Location.
     */
    readonly location: string;
    /**
     * Resource Name.
     */
    readonly name: string;
    /**
     * PremierAddOn resource specific properties
     */
    readonly properties: outputs.web.PremierAddOnResponseProperties;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type.
     */
    readonly type: string;
}