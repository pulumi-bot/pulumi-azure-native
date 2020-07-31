// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getManagementPolicy(args: GetManagementPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetManagementPolicyResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:storage/v20181101:getManagementPolicy", {
        "accountName": args.accountName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetManagementPolicyArgs {
    /**
     * The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
     */
    readonly accountName: string;
    /**
     * The name of the Storage Account Management Policy. It should always be 'default'
     */
    readonly name: string;
    /**
     * The name of the resource group within the user's subscription. The name is case insensitive.
     */
    readonly resourceGroupName: string;
}

/**
 * The Get Storage Account ManagementPolicies operation response.
 */
export interface GetManagementPolicyResult {
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * Returns the Storage Account Data Policies Rules.
     */
    readonly properties: outputs.storage.v20181101.ManagementPolicyPropertiesResponse;
    /**
     * The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
     */
    readonly type: string;
}