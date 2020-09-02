// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The storage account.
 *
 * ## Example Usage
 * ### StorageAccountCreate
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const storageAccount = new azurerm.storage.latest.StorageAccount("storageAccount", {
 *     accountName: "sto4445",
 *     allowBlobPublicAccess: false,
 *     encryption: {
 *         keySource: "Microsoft.Storage",
 *         requireInfrastructureEncryption: false,
 *         services: {
 *             blob: {
 *                 enabled: true,
 *                 keyType: "Account",
 *             },
 *             file: {
 *                 enabled: true,
 *                 keyType: "Account",
 *             },
 *         },
 *     },
 *     isHnsEnabled: true,
 *     kind: "Storage",
 *     location: "eastus",
 *     minimumTlsVersion: "TLS1_2",
 *     resourceGroupName: "res9101",
 *     routingPreference: {
 *         publishInternetEndpoints: true,
 *         publishMicrosoftEndpoints: true,
 *         routingChoice: "MicrosoftRouting",
 *     },
 *     sku: {
 *         name: "Standard_GRS",
 *     },
 *     tags: {
 *         key1: "value1",
 *         key2: "value2",
 *     },
 * });
 *
 * ```
 */
export class StorageAccount extends pulumi.CustomResource {
    /**
     * Get an existing StorageAccount resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): StorageAccount {
        return new StorageAccount(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:storage/latest:StorageAccount';

    /**
     * Returns true if the given object is an instance of StorageAccount.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StorageAccount {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StorageAccount.__pulumiType;
    }

    /**
     * Required for storage accounts where kind = BlobStorage. The access tier used for billing.
     */
    public readonly accessTier!: pulumi.Output<string>;
    /**
     * Allow or disallow public access to all blobs or containers in the storage account. The default interpretation is true for this property.
     */
    public readonly allowBlobPublicAccess!: pulumi.Output<boolean | undefined>;
    /**
     * Provides the identity based authentication settings for Azure Files.
     */
    public readonly azureFilesIdentityBasedAuthentication!: pulumi.Output<outputs.storage.latest.AzureFilesIdentityBasedAuthenticationResponse | undefined>;
    /**
     * Blob restore status
     */
    public /*out*/ readonly blobRestoreStatus!: pulumi.Output<outputs.storage.latest.BlobRestoreStatusResponse>;
    /**
     * Gets the creation date and time of the storage account in UTC.
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    /**
     * Gets the custom domain the user assigned to this storage account.
     */
    public readonly customDomain!: pulumi.Output<outputs.storage.latest.CustomDomainResponse>;
    /**
     * Allows https traffic only to storage service if sets to true.
     */
    public readonly enableHttpsTrafficOnly!: pulumi.Output<boolean | undefined>;
    /**
     * Gets the encryption settings on the account. If unspecified, the account is unencrypted.
     */
    public readonly encryption!: pulumi.Output<outputs.storage.latest.EncryptionResponse>;
    /**
     * If the failover is in progress, the value will be true, otherwise, it will be null.
     */
    public /*out*/ readonly failoverInProgress!: pulumi.Output<boolean>;
    /**
     * Geo Replication Stats
     */
    public /*out*/ readonly geoReplicationStats!: pulumi.Output<outputs.storage.latest.GeoReplicationStatsResponse>;
    /**
     * The identity of the resource.
     */
    public readonly identity!: pulumi.Output<outputs.storage.latest.IdentityResponse | undefined>;
    /**
     * Account HierarchicalNamespace enabled if sets to true.
     */
    public readonly isHnsEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Gets the Kind.
     */
    public readonly kind!: pulumi.Output<string>;
    /**
     * Allow large file shares if sets to Enabled. It cannot be disabled once it is enabled.
     */
    public readonly largeFileSharesState!: pulumi.Output<string | undefined>;
    /**
     * Gets the timestamp of the most recent instance of a failover to the secondary location. Only the most recent timestamp is retained. This element is not returned if there has never been a failover instance. Only available if the accountType is Standard_GRS or Standard_RAGRS.
     */
    public /*out*/ readonly lastGeoFailoverTime!: pulumi.Output<string>;
    /**
     * The geo-location where the resource lives
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Set the minimum TLS version to be permitted on requests to storage. The default interpretation is TLS 1.0 for this property.
     */
    public readonly minimumTlsVersion!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Network rule set
     */
    public readonly networkRuleSet!: pulumi.Output<outputs.storage.latest.NetworkRuleSetResponse>;
    /**
     * Gets the URLs that are used to perform a retrieval of a public blob, queue, or table object. Note that Standard_ZRS and Premium_LRS accounts only return the blob endpoint.
     */
    public /*out*/ readonly primaryEndpoints!: pulumi.Output<outputs.storage.latest.EndpointsResponse>;
    /**
     * Gets the location of the primary data center for the storage account.
     */
    public /*out*/ readonly primaryLocation!: pulumi.Output<string>;
    /**
     * List of private endpoint connection associated with the specified storage account
     */
    public /*out*/ readonly privateEndpointConnections!: pulumi.Output<outputs.storage.latest.PrivateEndpointConnectionResponse[]>;
    /**
     * Gets the status of the storage account at the time the operation was called.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Maintains information about the network routing choice opted by the user for data transfer
     */
    public readonly routingPreference!: pulumi.Output<outputs.storage.latest.RoutingPreferenceResponse | undefined>;
    /**
     * Gets the URLs that are used to perform a retrieval of a public blob, queue, or table object from the secondary location of the storage account. Only available if the SKU name is Standard_RAGRS.
     */
    public /*out*/ readonly secondaryEndpoints!: pulumi.Output<outputs.storage.latest.EndpointsResponse>;
    /**
     * Gets the location of the geo-replicated secondary for the storage account. Only available if the accountType is Standard_GRS or Standard_RAGRS.
     */
    public /*out*/ readonly secondaryLocation!: pulumi.Output<string>;
    /**
     * Gets the SKU.
     */
    public readonly sku!: pulumi.Output<outputs.storage.latest.SkuResponse>;
    /**
     * Gets the status indicating whether the primary location of the storage account is available or unavailable.
     */
    public /*out*/ readonly statusOfPrimary!: pulumi.Output<string>;
    /**
     * Gets the status indicating whether the secondary location of the storage account is available or unavailable. Only available if the SKU name is Standard_GRS or Standard_RAGRS.
     */
    public /*out*/ readonly statusOfSecondary!: pulumi.Output<string>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a StorageAccount resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StorageAccountArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StorageAccountArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as StorageAccountArgs | undefined;
            if (!args || args.accountName === undefined) {
                throw new Error("Missing required property 'accountName'");
            }
            if (!args || args.kind === undefined) {
                throw new Error("Missing required property 'kind'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.sku === undefined) {
                throw new Error("Missing required property 'sku'");
            }
            inputs["accessTier"] = args ? args.accessTier : undefined;
            inputs["accountName"] = args ? args.accountName : undefined;
            inputs["allowBlobPublicAccess"] = args ? args.allowBlobPublicAccess : undefined;
            inputs["azureFilesIdentityBasedAuthentication"] = args ? args.azureFilesIdentityBasedAuthentication : undefined;
            inputs["customDomain"] = args ? args.customDomain : undefined;
            inputs["enableHttpsTrafficOnly"] = args ? args.enableHttpsTrafficOnly : undefined;
            inputs["encryption"] = args ? args.encryption : undefined;
            inputs["identity"] = args ? args.identity : undefined;
            inputs["isHnsEnabled"] = args ? args.isHnsEnabled : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["largeFileSharesState"] = args ? args.largeFileSharesState : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["minimumTlsVersion"] = args ? args.minimumTlsVersion : undefined;
            inputs["networkRuleSet"] = args ? args.networkRuleSet : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["routingPreference"] = args ? args.routingPreference : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["blobRestoreStatus"] = undefined /*out*/;
            inputs["creationTime"] = undefined /*out*/;
            inputs["failoverInProgress"] = undefined /*out*/;
            inputs["geoReplicationStats"] = undefined /*out*/;
            inputs["lastGeoFailoverTime"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["primaryEndpoints"] = undefined /*out*/;
            inputs["primaryLocation"] = undefined /*out*/;
            inputs["privateEndpointConnections"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["secondaryEndpoints"] = undefined /*out*/;
            inputs["secondaryLocation"] = undefined /*out*/;
            inputs["statusOfPrimary"] = undefined /*out*/;
            inputs["statusOfSecondary"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:storage/v20150615:StorageAccount" }, { type: "azurerm:storage/v20160101:StorageAccount" }, { type: "azurerm:storage/v20160501:StorageAccount" }, { type: "azurerm:storage/v20161201:StorageAccount" }, { type: "azurerm:storage/v20170601:StorageAccount" }, { type: "azurerm:storage/v20171001:StorageAccount" }, { type: "azurerm:storage/v20180201:StorageAccount" }, { type: "azurerm:storage/v20180701:StorageAccount" }, { type: "azurerm:storage/v20181101:StorageAccount" }, { type: "azurerm:storage/v20190401:StorageAccount" }, { type: "azurerm:storage/v20190601:StorageAccount" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(StorageAccount.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a StorageAccount resource.
 */
export interface StorageAccountArgs {
    /**
     * Required for storage accounts where kind = BlobStorage. The access tier used for billing.
     */
    readonly accessTier?: pulumi.Input<string>;
    /**
     * The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
     */
    readonly accountName: pulumi.Input<string>;
    /**
     * Allow or disallow public access to all blobs or containers in the storage account. The default interpretation is true for this property.
     */
    readonly allowBlobPublicAccess?: pulumi.Input<boolean>;
    /**
     * Provides the identity based authentication settings for Azure Files.
     */
    readonly azureFilesIdentityBasedAuthentication?: pulumi.Input<inputs.storage.latest.AzureFilesIdentityBasedAuthentication>;
    /**
     * User domain assigned to the storage account. Name is the CNAME source. Only one custom domain is supported per storage account at this time. To clear the existing custom domain, use an empty string for the custom domain name property.
     */
    readonly customDomain?: pulumi.Input<inputs.storage.latest.CustomDomain>;
    /**
     * Allows https traffic only to storage service if sets to true. The default value is true since API version 2019-04-01.
     */
    readonly enableHttpsTrafficOnly?: pulumi.Input<boolean>;
    /**
     * Not applicable. Azure Storage encryption is enabled for all storage accounts and cannot be disabled.
     */
    readonly encryption?: pulumi.Input<inputs.storage.latest.Encryption>;
    /**
     * The identity of the resource.
     */
    readonly identity?: pulumi.Input<inputs.storage.latest.Identity>;
    /**
     * Account HierarchicalNamespace enabled if sets to true.
     */
    readonly isHnsEnabled?: pulumi.Input<boolean>;
    /**
     * Required. Indicates the type of storage account.
     */
    readonly kind: pulumi.Input<string>;
    /**
     * Allow large file shares if sets to Enabled. It cannot be disabled once it is enabled.
     */
    readonly largeFileSharesState?: pulumi.Input<string>;
    /**
     * Required. Gets or sets the location of the resource. This will be one of the supported and registered Azure Geo Regions (e.g. West US, East US, Southeast Asia, etc.). The geo region of a resource cannot be changed once it is created, but if an identical geo region is specified on update, the request will succeed.
     */
    readonly location: pulumi.Input<string>;
    /**
     * Set the minimum TLS version to be permitted on requests to storage. The default interpretation is TLS 1.0 for this property.
     */
    readonly minimumTlsVersion?: pulumi.Input<string>;
    /**
     * Network rule set
     */
    readonly networkRuleSet?: pulumi.Input<inputs.storage.latest.NetworkRuleSet>;
    /**
     * The name of the resource group within the user's subscription. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Maintains information about the network routing choice opted by the user for data transfer
     */
    readonly routingPreference?: pulumi.Input<inputs.storage.latest.RoutingPreference>;
    /**
     * Required. Gets or sets the SKU name.
     */
    readonly sku: pulumi.Input<inputs.storage.latest.Sku>;
    /**
     * Gets or sets a list of key value pairs that describe the resource. These tags can be used for viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key with a length no greater than 128 characters and a value with a length no greater than 256 characters.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
