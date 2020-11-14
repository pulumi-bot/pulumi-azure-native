// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getMediaService(args: GetMediaServiceArgs, opts?: pulumi.InvokeOptions): Promise<GetMediaServiceResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:media/latest:getMediaService", {
        "accountName": args.accountName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetMediaServiceArgs {
    /**
     * The Media Services account name.
     */
    readonly accountName: string;
    /**
     * The name of the resource group within the Azure subscription.
     */
    readonly resourceGroupName: string;
}

/**
 * A Media Services account.
 */
export interface GetMediaServiceResult {
    /**
     * The account encryption properties.
     */
    readonly encryption?: outputs.media.latest.AccountEncryptionResponse;
    /**
     * The Managed Identity for the Media Services account.
     */
    readonly identity?: outputs.media.latest.MediaServiceIdentityResponse;
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * The Media Services account ID.
     */
    readonly mediaServiceId: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * The storage accounts for this resource.
     */
    readonly storageAccounts?: outputs.media.latest.StorageAccountResponse[];
    readonly storageAuthentication?: string;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
}
