// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * An Azure Cosmos DB database account.
 */
export class DatabaseAccount extends pulumi.CustomResource {
    /**
     * Get an existing DatabaseAccount resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DatabaseAccount {
        return new DatabaseAccount(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:documentdb/v20200401:DatabaseAccount';

    /**
     * Returns true if the given object is an instance of DatabaseAccount.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatabaseAccount {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatabaseAccount.__pulumiType;
    }

    /**
     * API specific properties.
     */
    public readonly apiProperties!: pulumi.Output<outputs.documentdb.v20200401.ApiPropertiesResponse | undefined>;
    /**
     * List of Cosmos DB capabilities for the account
     */
    public readonly capabilities!: pulumi.Output<outputs.documentdb.v20200401.CapabilityResponse[] | undefined>;
    /**
     * The cassandra connector offer type for the Cosmos DB database C* account.
     */
    public readonly connectorOffer!: pulumi.Output<string | undefined>;
    /**
     * The consistency policy for the Cosmos DB database account.
     */
    public readonly consistencyPolicy!: pulumi.Output<outputs.documentdb.v20200401.ConsistencyPolicyResponse | undefined>;
    /**
     * The CORS policy for the Cosmos DB database account.
     */
    public readonly cors!: pulumi.Output<outputs.documentdb.v20200401.CorsPolicyResponse[] | undefined>;
    /**
     * The offer type for the Cosmos DB database account. Default value: Standard.
     */
    public readonly databaseAccountOfferType!: pulumi.Output<string>;
    /**
     * Disable write operations on metadata resources (databases, containers, throughput) via account keys
     */
    public readonly disableKeyBasedMetadataWriteAccess!: pulumi.Output<boolean | undefined>;
    /**
     * The connection endpoint for the Cosmos DB database account.
     */
    public /*out*/ readonly documentEndpoint!: pulumi.Output<string>;
    /**
     * Flag to indicate whether to enable storage analytics.
     */
    public readonly enableAnalyticalStorage!: pulumi.Output<boolean | undefined>;
    /**
     * Enables automatic failover of the write region in the rare event that the region is unavailable due to an outage. Automatic failover will result in a new write region for the account and is chosen based on the failover priorities configured for the account.
     */
    public readonly enableAutomaticFailover!: pulumi.Output<boolean | undefined>;
    /**
     * Enables the cassandra connector on the Cosmos DB C* account
     */
    public readonly enableCassandraConnector!: pulumi.Output<boolean | undefined>;
    /**
     * Flag to indicate whether Free Tier is enabled.
     */
    public readonly enableFreeTier!: pulumi.Output<boolean | undefined>;
    /**
     * Enables the account to write in multiple locations
     */
    public readonly enableMultipleWriteLocations!: pulumi.Output<boolean | undefined>;
    /**
     * An array that contains the regions ordered by their failover priorities.
     */
    public /*out*/ readonly failoverPolicies!: pulumi.Output<outputs.documentdb.v20200401.FailoverPolicyResponse[]>;
    /**
     * List of IpRules.
     */
    public readonly ipRules!: pulumi.Output<outputs.documentdb.v20200401.IpAddressOrRangeResponse[] | undefined>;
    /**
     * Flag to indicate whether to enable/disable Virtual Network ACL rules.
     */
    public readonly isVirtualNetworkFilterEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The URI of the key vault
     */
    public readonly keyVaultKeyUri!: pulumi.Output<string | undefined>;
    /**
     * Indicates the type of database account. This can only be set at database account creation.
     */
    public readonly kind!: pulumi.Output<string | undefined>;
    /**
     * The location of the resource group to which the resource belongs.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * An array that contains all of the locations enabled for the Cosmos DB account.
     */
    public readonly locations!: pulumi.Output<outputs.documentdb.v20200401.LocationResponse[]>;
    /**
     * The name of the ARM resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * List of Private Endpoint Connections configured for the Cosmos DB account.
     */
    public /*out*/ readonly privateEndpointConnections!: pulumi.Output<outputs.documentdb.v20200401.PrivateEndpointConnectionResponse[]>;
    /**
     * The status of the Cosmos DB account at the time the operation was called. The status can be one of following. 'Creating' – the Cosmos DB account is being created. When an account is in Creating state, only properties that are specified as input for the Create Cosmos DB account operation are returned. 'Succeeded' – the Cosmos DB account is active for use. 'Updating' – the Cosmos DB account is being updated. 'Deleting' – the Cosmos DB account is being deleted. 'Failed' – the Cosmos DB account failed creation. 'DeletionFailed' – the Cosmos DB account deletion failed.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string | undefined>;
    /**
     * Whether requests from Public Network are allowed
     */
    public readonly publicNetworkAccess!: pulumi.Output<string | undefined>;
    /**
     * An array that contains of the read locations enabled for the Cosmos DB account.
     */
    public /*out*/ readonly readLocations!: pulumi.Output<outputs.documentdb.v20200401.LocationResponse[]>;
    /**
     * Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of Azure resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * List of Virtual Network ACL rules configured for the Cosmos DB account.
     */
    public readonly virtualNetworkRules!: pulumi.Output<outputs.documentdb.v20200401.VirtualNetworkRuleResponse[] | undefined>;
    /**
     * An array that contains the write location for the Cosmos DB account.
     */
    public /*out*/ readonly writeLocations!: pulumi.Output<outputs.documentdb.v20200401.LocationResponse[]>;

    /**
     * Create a DatabaseAccount resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatabaseAccountArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatabaseAccountArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as DatabaseAccountArgs | undefined;
            if (!args || args.accountName === undefined) {
                throw new Error("Missing required property 'accountName'");
            }
            if (!args || args.databaseAccountOfferType === undefined) {
                throw new Error("Missing required property 'databaseAccountOfferType'");
            }
            if (!args || args.locations === undefined) {
                throw new Error("Missing required property 'locations'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["accountName"] = args ? args.accountName : undefined;
            inputs["apiProperties"] = args ? args.apiProperties : undefined;
            inputs["capabilities"] = args ? args.capabilities : undefined;
            inputs["connectorOffer"] = args ? args.connectorOffer : undefined;
            inputs["consistencyPolicy"] = args ? args.consistencyPolicy : undefined;
            inputs["cors"] = args ? args.cors : undefined;
            inputs["databaseAccountOfferType"] = args ? args.databaseAccountOfferType : undefined;
            inputs["disableKeyBasedMetadataWriteAccess"] = args ? args.disableKeyBasedMetadataWriteAccess : undefined;
            inputs["enableAnalyticalStorage"] = args ? args.enableAnalyticalStorage : undefined;
            inputs["enableAutomaticFailover"] = args ? args.enableAutomaticFailover : undefined;
            inputs["enableCassandraConnector"] = args ? args.enableCassandraConnector : undefined;
            inputs["enableFreeTier"] = args ? args.enableFreeTier : undefined;
            inputs["enableMultipleWriteLocations"] = args ? args.enableMultipleWriteLocations : undefined;
            inputs["ipRules"] = args ? args.ipRules : undefined;
            inputs["isVirtualNetworkFilterEnabled"] = args ? args.isVirtualNetworkFilterEnabled : undefined;
            inputs["keyVaultKeyUri"] = args ? args.keyVaultKeyUri : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["locations"] = args ? args.locations : undefined;
            inputs["publicNetworkAccess"] = args ? args.publicNetworkAccess : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["virtualNetworkRules"] = args ? args.virtualNetworkRules : undefined;
            inputs["documentEndpoint"] = undefined /*out*/;
            inputs["failoverPolicies"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["privateEndpointConnections"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["readLocations"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["writeLocations"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:documentdb/latest:DatabaseAccount" }, { type: "azurerm:documentdb/v20150401:DatabaseAccount" }, { type: "azurerm:documentdb/v20150408:DatabaseAccount" }, { type: "azurerm:documentdb/v20151106:DatabaseAccount" }, { type: "azurerm:documentdb/v20160319:DatabaseAccount" }, { type: "azurerm:documentdb/v20160331:DatabaseAccount" }, { type: "azurerm:documentdb/v20190801:DatabaseAccount" }, { type: "azurerm:documentdb/v20191212:DatabaseAccount" }, { type: "azurerm:documentdb/v20200301:DatabaseAccount" }, { type: "azurerm:documentdb/v20200601preview:DatabaseAccount" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(DatabaseAccount.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DatabaseAccount resource.
 */
export interface DatabaseAccountArgs {
    /**
     * Cosmos DB database account name.
     */
    readonly accountName: pulumi.Input<string>;
    /**
     * API specific properties. Currently, supported only for MongoDB API.
     */
    readonly apiProperties?: pulumi.Input<inputs.documentdb.v20200401.ApiProperties>;
    /**
     * List of Cosmos DB capabilities for the account
     */
    readonly capabilities?: pulumi.Input<pulumi.Input<inputs.documentdb.v20200401.Capability>[]>;
    /**
     * The cassandra connector offer type for the Cosmos DB database C* account.
     */
    readonly connectorOffer?: pulumi.Input<string>;
    /**
     * The consistency policy for the Cosmos DB account.
     */
    readonly consistencyPolicy?: pulumi.Input<inputs.documentdb.v20200401.ConsistencyPolicy>;
    /**
     * The CORS policy for the Cosmos DB database account.
     */
    readonly cors?: pulumi.Input<pulumi.Input<inputs.documentdb.v20200401.CorsPolicy>[]>;
    /**
     * The offer type for the database
     */
    readonly databaseAccountOfferType: pulumi.Input<string>;
    /**
     * Disable write operations on metadata resources (databases, containers, throughput) via account keys
     */
    readonly disableKeyBasedMetadataWriteAccess?: pulumi.Input<boolean>;
    /**
     * Flag to indicate whether to enable storage analytics.
     */
    readonly enableAnalyticalStorage?: pulumi.Input<boolean>;
    /**
     * Enables automatic failover of the write region in the rare event that the region is unavailable due to an outage. Automatic failover will result in a new write region for the account and is chosen based on the failover priorities configured for the account.
     */
    readonly enableAutomaticFailover?: pulumi.Input<boolean>;
    /**
     * Enables the cassandra connector on the Cosmos DB C* account
     */
    readonly enableCassandraConnector?: pulumi.Input<boolean>;
    /**
     * Flag to indicate whether Free Tier is enabled.
     */
    readonly enableFreeTier?: pulumi.Input<boolean>;
    /**
     * Enables the account to write in multiple locations
     */
    readonly enableMultipleWriteLocations?: pulumi.Input<boolean>;
    /**
     * List of IpRules.
     */
    readonly ipRules?: pulumi.Input<pulumi.Input<inputs.documentdb.v20200401.IpAddressOrRange>[]>;
    /**
     * Flag to indicate whether to enable/disable Virtual Network ACL rules.
     */
    readonly isVirtualNetworkFilterEnabled?: pulumi.Input<boolean>;
    /**
     * The URI of the key vault
     */
    readonly keyVaultKeyUri?: pulumi.Input<string>;
    /**
     * Indicates the type of database account. This can only be set at database account creation.
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * The location of the resource group to which the resource belongs.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * An array that contains the georeplication locations enabled for the Cosmos DB account.
     */
    readonly locations: pulumi.Input<pulumi.Input<inputs.documentdb.v20200401.Location>[]>;
    /**
     * Whether requests from Public Network are allowed
     */
    readonly publicNetworkAccess?: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * List of Virtual Network ACL rules configured for the Cosmos DB account.
     */
    readonly virtualNetworkRules?: pulumi.Input<pulumi.Input<inputs.documentdb.v20200401.VirtualNetworkRule>[]>;
}
