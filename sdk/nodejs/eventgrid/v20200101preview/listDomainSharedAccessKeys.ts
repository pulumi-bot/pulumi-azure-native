// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function listDomainSharedAccessKeys(args: ListDomainSharedAccessKeysArgs, opts?: pulumi.InvokeOptions): Promise<ListDomainSharedAccessKeysResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:eventgrid/v20200101preview:listDomainSharedAccessKeys", {
        "domainName": args.domainName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListDomainSharedAccessKeysArgs {
    /**
     * Name of the domain
     */
    readonly domainName: string;
    /**
     * The name of the resource group within the user's subscription.
     */
    readonly resourceGroupName: string;
}

/**
 * Shared access keys of the Domain
 */
export interface ListDomainSharedAccessKeysResult {
    /**
     * Shared access key1 for the domain.
     */
    readonly key1?: string;
    /**
     * Shared access key2 for the domain.
     */
    readonly key2?: string;
}
