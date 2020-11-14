// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getStorageDomain(args: GetStorageDomainArgs, opts?: pulumi.InvokeOptions): Promise<GetStorageDomainResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:storsimple/latest:getStorageDomain", {
        "managerName": args.managerName,
        "resourceGroupName": args.resourceGroupName,
        "storageDomainName": args.storageDomainName,
    }, opts);
}

export interface GetStorageDomainArgs {
    /**
     * The manager name
     */
    readonly managerName: string;
    /**
     * The resource group name
     */
    readonly resourceGroupName: string;
    /**
     * The storage domain name.
     */
    readonly storageDomainName: string;
}

/**
 * The storage domain.
 */
export interface GetStorageDomainResult {
    /**
     * The encryption key used to encrypt the data. This is a user secret.
     */
    readonly encryptionKey?: outputs.storsimple.latest.AsymmetricEncryptedSecretResponse;
    /**
     * The encryption status "Enabled | Disabled".
     */
    readonly encryptionStatus: string;
    /**
     * The name.
     */
    readonly name: string;
    /**
     * The storage account credentials.
     */
    readonly storageAccountCredentialIds: string[];
    /**
     * The type.
     */
    readonly type: string;
}
