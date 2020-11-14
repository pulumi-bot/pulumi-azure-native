// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getTagByProduct(args: GetTagByProductArgs, opts?: pulumi.InvokeOptions): Promise<GetTagByProductResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:apimanagement/latest:getTagByProduct", {
        "productId": args.productId,
        "resourceGroupName": args.resourceGroupName,
        "serviceName": args.serviceName,
        "tagId": args.tagId,
    }, opts);
}

export interface GetTagByProductArgs {
    /**
     * Product identifier. Must be unique in the current API Management service instance.
     */
    readonly productId: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the API Management service.
     */
    readonly serviceName: string;
    /**
     * Tag identifier. Must be unique in the current API Management service instance.
     */
    readonly tagId: string;
}

/**
 * Tag Contract details.
 */
export interface GetTagByProductResult {
    /**
     * Tag name.
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
}
