// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getProtectionIntent(args: GetProtectionIntentArgs, opts?: pulumi.InvokeOptions): Promise<GetProtectionIntentResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:recoveryservices/latest:getProtectionIntent", {
        "fabricName": args.fabricName,
        "intentObjectName": args.intentObjectName,
        "resourceGroupName": args.resourceGroupName,
        "vaultName": args.vaultName,
    }, opts);
}

export interface GetProtectionIntentArgs {
    /**
     * Fabric name associated with the backed up item.
     */
    readonly fabricName: string;
    /**
     * Backed up item name whose details are to be fetched.
     */
    readonly intentObjectName: string;
    /**
     * The name of the resource group where the recovery services vault is present.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the recovery services vault.
     */
    readonly vaultName: string;
}

/**
 * Base class for backup ProtectionIntent.
 */
export interface GetProtectionIntentResult {
    /**
     * Optional ETag.
     */
    readonly eTag?: string;
    /**
     * Resource location.
     */
    readonly location?: string;
    /**
     * Resource name associated with the resource.
     */
    readonly name: string;
    /**
     * ProtectionIntentResource properties
     */
    readonly properties: outputs.recoveryservices.latest.AzureRecoveryServiceVaultProtectionIntentResponse | outputs.recoveryservices.latest.AzureResourceProtectionIntentResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
     */
    readonly type: string;
}
