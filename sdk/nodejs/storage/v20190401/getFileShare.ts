// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getFileShare(args: GetFileShareArgs, opts?: pulumi.InvokeOptions): Promise<GetFileShareResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:storage/v20190401:getFileShare", {
        "accountName": args.accountName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetFileShareArgs {
    /**
     * The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
     */
    readonly accountName: string;
    /**
     * The name of the file share within the specified storage account. File share names must be between 3 and 63 characters in length and use numbers, lower-case letters and dash (-) only. Every dash (-) character must be immediately preceded and followed by a letter or number.
     */
    readonly name: string;
    /**
     * The name of the resource group within the user's subscription. The name is case insensitive.
     */
    readonly resourceGroupName: string;
}

/**
 * Properties of the file share, including Id, resource name, resource type, Etag.
 */
export interface GetFileShareResult {
    /**
     * Resource Etag.
     */
    readonly etag: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * Properties of the file share.
     */
    readonly properties: outputs.storage.v20190401.FileSharePropertiesResponse;
    /**
     * The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
     */
    readonly type: string;
}