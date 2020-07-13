// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getAccountShareSubscriptionDataSetMapping(args: GetAccountShareSubscriptionDataSetMappingArgs, opts?: pulumi.InvokeOptions): Promise<GetAccountShareSubscriptionDataSetMappingResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:datashare:getAccountShareSubscriptionDataSetMapping", {
        "accountName": args.accountName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
        "shareSubscriptionName": args.shareSubscriptionName,
    }, opts);
}

export interface GetAccountShareSubscriptionDataSetMappingArgs {
    /**
     * The name of the share account.
     */
    readonly accountName: string;
    /**
     * The name of the dataSetMapping.
     */
    readonly name: string;
    /**
     * The resource group name.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the shareSubscription.
     */
    readonly shareSubscriptionName: string;
}

/**
 * A data set mapping data transfer object.
 */
export interface GetAccountShareSubscriptionDataSetMappingResult {
    /**
     * Kind of data set mapping.
     */
    readonly kind: string;
    /**
     * Name of the azure resource
     */
    readonly name: string;
    /**
     * Type of the azure resource
     */
    readonly type: string;
}