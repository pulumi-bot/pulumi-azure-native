// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getApiRelease(args: GetApiReleaseArgs, opts?: pulumi.InvokeOptions): Promise<GetApiReleaseResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:apimanagement/v20190101:getApiRelease", {
        "apiId": args.apiId,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
        "serviceName": args.serviceName,
    }, opts);
}

export interface GetApiReleaseArgs {
    /**
     * API identifier. Must be unique in the current API Management service instance.
     */
    readonly apiId: string;
    /**
     * Release identifier within an API. Must be unique in the current API Management service instance.
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
 * ApiRelease details.
 */
export interface GetApiReleaseResult {
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * ApiRelease entity contract properties.
     */
    readonly properties: outputs.apimanagement.v20190101.ApiReleaseContractPropertiesResponse;
    /**
     * Resource type for API Management resource.
     */
    readonly type: string;
}