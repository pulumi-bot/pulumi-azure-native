// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getContentType(args: GetContentTypeArgs, opts?: pulumi.InvokeOptions): Promise<GetContentTypeResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:apimanagement/v20191201:getContentType", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
        "serviceName": args.serviceName,
    }, opts);
}

export interface GetContentTypeArgs {
    /**
     * Content type identifier.
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
 * Content type contract details.
 */
export interface GetContentTypeResult {
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * Properties of the content type.
     */
    readonly properties: outputs.apimanagement.v20191201.ContentTypeContractPropertiesResponse;
    /**
     * Resource type for API Management resource.
     */
    readonly type: string;
}