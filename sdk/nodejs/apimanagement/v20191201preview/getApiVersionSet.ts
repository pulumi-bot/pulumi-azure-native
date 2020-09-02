// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getApiVersionSet(args: GetApiVersionSetArgs, opts?: pulumi.InvokeOptions): Promise<GetApiVersionSetResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:apimanagement/v20191201preview:getApiVersionSet", {
        "resourceGroupName": args.resourceGroupName,
        "serviceName": args.serviceName,
        "versionSetId": args.versionSetId,
    }, opts);
}

export interface GetApiVersionSetArgs {
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the API Management service.
     */
    readonly serviceName: string;
    /**
     * Api Version Set identifier. Must be unique in the current API Management service instance.
     */
    readonly versionSetId: string;
}

/**
 * Api Version Set Contract details.
 */
export interface GetApiVersionSetResult {
    /**
     * Description of API Version Set.
     */
    readonly description?: string;
    /**
     * Name of API Version Set
     */
    readonly displayName: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * Resource type for API Management resource.
     */
    readonly type: string;
    /**
     * Name of HTTP header parameter that indicates the API Version if versioningScheme is set to `header`.
     */
    readonly versionHeaderName?: string;
    /**
     * Name of query parameter that indicates the API Version if versioningScheme is set to `query`.
     */
    readonly versionQueryName?: string;
    /**
     * An value that determines where the API Version identifer will be located in a HTTP request.
     */
    readonly versioningScheme: string;
}
