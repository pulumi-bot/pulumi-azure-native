// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

export function getStorageAccount(args: GetStorageAccountArgs, opts?: pulumi.InvokeOptions): Promise<GetStorageAccountResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:storage/latest:getStorageAccount", {
        "accountName": args.accountName,
        "expand": args.expand,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetStorageAccountArgs {
    /**
     * The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
     */
    readonly accountName: string;
    /**
     * May be used to expand the properties within account's properties. By default, data is not included when fetching properties. Currently we only support geoReplicationStats and blobRestoreStatus.
     */
    readonly expand?: string;
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
    readonly accessTier: string;
    /**
     * Allow or disallow public access to all blobs or containers in the storage account. The default interpretation is true for this property.
     */
    readonly allowBlobPublicAccess?: boolean;
    /**
     * Indicates whether the storage account permits requests to be authorized with the account access key via Shared Key. If false, then all requests, including shared access signatures, must be authorized with Azure Active Directory (Azure AD). The default value is null, which is equivalent to true.
     */
    readonly allowSharedKeyAccess?: boolean;
    /**
     * Provides the identity based authentication settings for Azure Files.
     */
    readonly azureFilesIdentityBasedAuthentication?: outputs.storage.latest.AzureFilesIdentityBasedAuthenticationResponse;
    /**
     * Blob restore status
     */
    readonly blobRestoreStatus: outputs.storage.latest.BlobRestoreStatusResponse;
    /**
     * Gets the creation date and time of the storage account in UTC.
     */
    readonly creationTime: string;
    /**
     * Gets the custom domain the user assigned to this storage account.
     */
    readonly customDomain: outputs.storage.latest.CustomDomainResponse;
    /**
     * Allows https traffic only to storage service if sets to true.
     */
    readonly enableHttpsTrafficOnly?: boolean;
    /**
     * Gets the encryption settings on the account. If unspecified, the account is unencrypted.
     */
    readonly encryption: outputs.storage.latest.EncryptionResponse;
    /**
     * If the failover is in progress, the value will be true, otherwise, it will be null.
     */
    readonly failoverInProgress: boolean;
    /**
     * Geo Replication Stats
     */
    readonly geoReplicationStats: outputs.storage.latest.GeoReplicationStatsResponse;
    /**
     * Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
     */
    readonly id: string;
    /**
     * The identity of the resource.
     */
    readonly identity?: outputs.storage.latest.IdentityResponse;
    /**
     * Account HierarchicalNamespace enabled if sets to true.
     */
    readonly isHnsEnabled?: boolean;
    /**
     * Gets the Kind.
     */
    readonly kind: string;
    /**
     * Allow large file shares if sets to Enabled. It cannot be disabled once it is enabled.
     */
    readonly largeFileSharesState?: string;
    /**
     * Gets the timestamp of the most recent instance of a failover to the secondary location. Only the most recent timestamp is retained. This element is not returned if there has never been a failover instance. Only available if the accountType is Standard_GRS or Standard_RAGRS.
     */
    readonly lastGeoFailoverTime: string;
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * Set the minimum TLS version to be permitted on requests to storage. The default interpretation is TLS 1.0 for this property.
     */
    readonly minimumTlsVersion?: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * Network rule set
     */
    readonly networkRuleSet: outputs.storage.latest.NetworkRuleSetResponse;
    /**
     * Gets the URLs that are used to perform a retrieval of a public blob, queue, or table object. Note that Standard_ZRS and Premium_LRS accounts only return the blob endpoint.
     */
    readonly primaryEndpoints: outputs.storage.latest.EndpointsResponse;
    /**
     * Gets the location of the primary data center for the storage account.
     */
    readonly primaryLocation: string;
    /**
     * List of private endpoint connection associated with the specified storage account
     */
    readonly privateEndpointConnections: outputs.storage.latest.PrivateEndpointConnectionResponse[];
    /**
     * Gets the status of the storage account at the time the operation was called.
     */
    readonly provisioningState: string;
    /**
     * Maintains information about the network routing choice opted by the user for data transfer
     */
    readonly routingPreference?: outputs.storage.latest.RoutingPreferenceResponse;
    /**
     * Gets the URLs that are used to perform a retrieval of a public blob, queue, or table object from the secondary location of the storage account. Only available if the SKU name is Standard_RAGRS.
     */
    readonly secondaryEndpoints: outputs.storage.latest.EndpointsResponse;
    /**
     * Gets the location of the geo-replicated secondary for the storage account. Only available if the accountType is Standard_GRS or Standard_RAGRS.
     */
    readonly secondaryLocation: string;
    /**
     * Gets the SKU.
     */
    readonly sku: outputs.storage.latest.SkuResponse;
    /**
     * Gets the status indicating whether the primary location of the storage account is available or unavailable.
     */
    readonly statusOfPrimary: string;
    /**
     * Gets the status indicating whether the secondary location of the storage account is available or unavailable. Only available if the SKU name is Standard_GRS or Standard_RAGRS.
     */
    readonly statusOfSecondary: string;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
}
