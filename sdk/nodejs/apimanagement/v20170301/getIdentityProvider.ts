// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getIdentityProvider(args: GetIdentityProviderArgs, opts?: pulumi.InvokeOptions): Promise<GetIdentityProviderResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:apimanagement/v20170301:getIdentityProvider", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
        "serviceName": args.serviceName,
    }, opts);
}

export interface GetIdentityProviderArgs {
    /**
     * Identity Provider Type identifier.
     */
    readonly name: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the API Management service.
     */
    readonly serviceName: string;
}

/**
 * Identity Provider details.
 */
export interface GetIdentityProviderResult {
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * Identity Provider contract properties.
     */
    readonly properties: outputs.apimanagement.v20170301.IdentityProviderContractPropertiesResponse;
    /**
     * Resource type for API Management resource.
     */
    readonly type: string;
}