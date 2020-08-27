// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getManagerExtendedInfo(args: GetManagerExtendedInfoArgs, opts?: pulumi.InvokeOptions): Promise<GetManagerExtendedInfoResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:storsimple/v20170601:getManagerExtendedInfo", {
        "managerName": args.managerName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetManagerExtendedInfoArgs {
    /**
     * The manager name
     */
    readonly managerName: string;
    /**
     * The resource group name
     */
    readonly resourceGroupName: string;
}

/**
 * The extended info of the manager.
 */
export interface GetManagerExtendedInfoResult {
    /**
     * Represents the encryption algorithm used to encrypt the keys. None - if Key is saved in plain text format. Algorithm name - if key is encrypted
     */
    readonly algorithm: string;
    /**
     * Represents the CEK of the resource.
     */
    readonly encryptionKey?: string;
    /**
     * Represents the Cert thumbprint that was used to encrypt the CEK.
     */
    readonly encryptionKeyThumbprint?: string;
    /**
     * The etag of the resource.
     */
    readonly etag?: string;
    /**
     * Represents the CIK of the resource.
     */
    readonly integrityKey: string;
    /**
     * The Kind of the object. Currently only Series8000 is supported
     */
    readonly kind?: Kind;
    /**
     * The name of the object.
     */
    readonly name: string;
    /**
     * Represents the portal thumbprint which can be used optionally to encrypt the entire data before storing it.
     */
    readonly portalCertificateThumbprint?: string;
    /**
     * The hierarchical type of the object.
     */
    readonly type: string;
    /**
     * The version of the extended info being persisted.
     */
    readonly version?: string;
}
