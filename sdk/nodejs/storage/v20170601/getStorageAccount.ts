// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getStorageAccount(args: GetStorageAccountArgs, opts?: pulumi.InvokeOptions): Promise<GetStorageAccountResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:storage/v20170601:getStorageAccount", {
        "accountName": args.accountName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetStorageAccountArgs {
    /**
     * The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.  
     */
    readonly accountName: string;
    /**
     * The name of the resource group within the user's subscription. The name is case insensitive.
     */
    readonly resourceGroupName: string;
}

/**
 * The storage account.
 */
export interface GetStorageAccountResult {
    /**
     * Required for storage accounts where kind = BlobStorage. The access tier used for billing.
     */
    readonly accessTier: AccessTier;
    /**
     * Gets the creation date and time of the storage account in UTC.
     */
    readonly creationTime: string;
    /**
     * Gets the custom domain the user assigned to this storage account.
     */
    readonly customDomain: outputs.storage.v20170601.CustomDomainResponse;
    /**
     * Allows https traffic only to storage service if sets to true.
     */
    readonly enableHttpsTrafficOnly?: boolean;
    /**
     * Gets the encryption settings on the account. If unspecified, the account is unencrypted.
     */
    readonly encryption: outputs.storage.v20170601.EncryptionResponse;
    /**
     * The identity of the resource.
     */
    readonly identity?: outputs.storage.v20170601.IdentityResponse;
    /**
     * Gets the Kind.
     */
    readonly kind: Kind;
    /**
     * Gets the timestamp of the most recent instance of a failover to the secondary location. Only the most recent timestamp is retained. This element is not returned if there has never been a failover instance. Only available if the accountType is Standard_GRS or Standard_RAGRS.
     */
    readonly lastGeoFailoverTime: string;
    /**
     * Resource location
     */
    readonly location?: string;
    /**
     * Resource name
     */
    readonly name: string;
    /**
     * Network rule set
     */
    readonly networkRuleSet: outputs.storage.v20170601.NetworkRuleSetResponse;
    /**
     * Gets the URLs that are used to perform a retrieval of a public blob, queue, or table object. Note that Standard_ZRS and Premium_LRS accounts only return the blob endpoint.
     */
    readonly primaryEndpoints: outputs.storage.v20170601.EndpointsResponse;
    /**
     * Gets the location of the primary data center for the storage account.
     */
    readonly primaryLocation: string;
    /**
     * Gets the status of the storage account at the time the operation was called.
     */
    readonly provisioningState: ProvisioningState;
    /**
     * Gets the URLs that are used to perform a retrieval of a public blob, queue, or table object from the secondary location of the storage account. Only available if the SKU name is Standard_RAGRS.
     */
    readonly secondaryEndpoints: outputs.storage.v20170601.EndpointsResponse;
    /**
     * Gets the location of the geo-replicated secondary for the storage account. Only available if the accountType is Standard_GRS or Standard_RAGRS.
     */
    readonly secondaryLocation: string;
    /**
     * Gets the SKU.
     */
    readonly sku: outputs.storage.v20170601.SkuResponse;
    /**
     * Gets the status indicating whether the primary location of the storage account is available or unavailable.
     */
    readonly statusOfPrimary: AccountStatus;
    /**
     * Gets the status indicating whether the secondary location of the storage account is available or unavailable. Only available if the SKU name is Standard_GRS or Standard_RAGRS.
     */
    readonly statusOfSecondary: AccountStatus;
    /**
     * Tags assigned to a resource; can be used for viewing and grouping a resource (across resource groups).
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type
     */
    readonly type: string;
}
