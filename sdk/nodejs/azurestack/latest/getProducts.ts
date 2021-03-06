// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Pageable list of products.
 * Latest API Version: 2017-06-01.
 */
/** @deprecated The 'latest' version is deprecated. Please migrate to the function in the top-level module: 'azure-native:azurestack:getProducts'. */
export function getProducts(args: GetProductsArgs, opts?: pulumi.InvokeOptions): Promise<GetProductsResult> {
    pulumi.log.warn("getProducts is deprecated: The 'latest' version is deprecated. Please migrate to the function in the top-level module: 'azure-native:azurestack:getProducts'.")
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:azurestack/latest:getProducts", {
        "productName": args.productName,
        "registrationName": args.registrationName,
        "resourceGroup": args.resourceGroup,
    }, opts);
}

export interface GetProductsArgs {
    /**
     * Name of the product.
     */
    readonly productName: string;
    /**
     * Name of the Azure Stack registration.
     */
    readonly registrationName: string;
    /**
     * Name of the resource group.
     */
    readonly resourceGroup: string;
}

/**
 * Pageable list of products.
 */
export interface GetProductsResult {
    /**
     * URI to the next page.
     */
    readonly nextLink?: string;
    /**
     * List of products.
     */
    readonly value?: outputs.azurestack.latest.ProductResponse[];
}
