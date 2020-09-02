// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Represents a database.
 *
 * ## Example Usage
 * ### Create a database as a copy
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const database = new azurerm.sql.v20140401.Database("database", {
 *     createMode: "Copy",
 *     databaseName: "testdb",
 *     location: "Japan East",
 *     resourceGroupName: "sqlcrudtest-4799",
 *     serverName: "sqlcrudtest-6440",
 *     sourceDatabaseId: "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-4799/providers/Microsoft.Sql/servers/sqlcrudtest-3782/databases/testdb",
 * });
 *
 * ```
 * ### Create a database as a dropped database restore to a specific point in time
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const database = new azurerm.sql.v20140401.Database("database", {
 *     createMode: "Restore",
 *     databaseName: "testdb",
 *     location: "Japan East",
 *     resourceGroupName: "sqlcrudtest-8412",
 *     restorePointInTime: "2017-05-20T21:24:37.467Z",
 *     serverName: "sqlcrudtest-3584",
 *     sourceDatabaseId: "/subscriptions/00000000-1111-2222-3333-444444444444 /resourceGroups/sqlcrudtest-8412/providers/Microsoft.Sql/servers/sqlcrudtest-3782/restorableDroppedDatabases/sourcedb,131403269876900000",
 * });
 *
 * ```
 * ### Create a database as a dropped database restore to deletion time
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const database = new azurerm.sql.v20140401.Database("database", {
 *     createMode: "Restore",
 *     databaseName: "testdb",
 *     location: "Japan East",
 *     resourceGroupName: "sqlcrudtest-8412",
 *     serverName: "sqlcrudtest-3584",
 *     sourceDatabaseDeletionDate: "2017-05-27T02:49:47.69Z",
 *     sourceDatabaseId: "/subscriptions/00000000-1111-2222-3333-444444444444 /resourceGroups/sqlcrudtest-8412/providers/Microsoft.Sql/servers/sqlcrudtest-3782/databases/sourcedb",
 * });
 *
 * ```
 * ### Create a database as a geo restore
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const database = new azurerm.sql.v20140401.Database("database", {
 *     createMode: "Recovery",
 *     databaseName: "testdb",
 *     location: "Japan East",
 *     resourceGroupName: "sqlcrudtest-8412",
 *     serverName: "sqlcrudtest-3584",
 *     sourceDatabaseId: "/subscriptions/00000000-1111-2222-3333-444444444444 /resourceGroups/sqlcrudtest-8412/providers/Microsoft.Sql/servers/sqlcrudtest-3782/recoverableDatabases/sourcedb",
 * });
 *
 * ```
 * ### Create a database as a non-readable secondary
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const database = new azurerm.sql.v20140401.Database("database", {
 *     createMode: "NonReadableSecondary",
 *     databaseName: "testdb",
 *     location: "Japan East",
 *     resourceGroupName: "sqlcrudtest-4799",
 *     serverName: "sqlcrudtest-6440",
 *     sourceDatabaseId: "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-4799/providers/Microsoft.Sql/servers/sqlcrudtest-3782/databases/testdb",
 * });
 *
 * ```
 * ### Create a database as a point in time restore
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const database = new azurerm.sql.v20140401.Database("database", {
 *     createMode: "PointInTimeRestore",
 *     databaseName: "testdb",
 *     location: "Japan East",
 *     resourceGroupName: "sqlcrudtest-8412",
 *     restorePointInTime: "2017-02-16T21:24:37.467Z",
 *     serverName: "sqlcrudtest-3584",
 *     sourceDatabaseId: "/subscriptions/00000000-1111-2222-3333-444444444444 /resourceGroups/sqlcrudtest-8412/providers/Microsoft.Sql/servers/sqlcrudtest-3782/databases/testdb",
 * });
 *
 * ```
 * ### Create a database as an online secondary
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const database = new azurerm.sql.v20140401.Database("database", {
 *     createMode: "OnlineSecondary",
 *     databaseName: "testdb",
 *     location: "Japan East",
 *     resourceGroupName: "sqlcrudtest-4799",
 *     serverName: "sqlcrudtest-6440",
 *     sourceDatabaseId: "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-4799/providers/Microsoft.Sql/servers/sqlcrudtest-3782/databases/testdb",
 * });
 *
 * ```
 * ### Create a database from a long term retention backup
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const database = new azurerm.sql.v20140401.Database("database", {
 *     createMode: "RestoreLongTermRetentionBackup",
 *     databaseName: "testdb",
 *     location: "Japan East",
 *     recoveryServicesRecoveryPointResourceId: "/subscriptions/00000000-1111-2222-3333-444444444444 /resourceGroups/sqlcrudtest-8412/providers/Microsoft.RecoveryServices/vaults/testVault/backupFabrics/Azure/protectionContainers/AzureSqlContainer;Sql;sqlcrudtest-8412;testsvr/protectedItems/AzureSqlDb;dsName;testdb;9dafcc99-7c84-4727-88ee-1a4fdb89afd7/RecoveryPoints/16043455089734",
 *     resourceGroupName: "sqlcrudtest-8412",
 *     serverName: "sqlcrudtest-3584",
 * });
 *
 * ```
 * ### Create a database max
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const database = new azurerm.sql.v20140401.Database("database", {
 *     collation: "SQL_Latin1_General_CP1_CI_AS",
 *     createMode: "Default",
 *     databaseName: "testdb",
 *     edition: "Standard",
 *     location: "Japan East",
 *     maxSizeBytes: "268435456000",
 *     readScale: "Disabled",
 *     requestedServiceObjectiveId: "f1173c43-91bd-4aaa-973c-54e79e15235b",
 *     requestedServiceObjectiveName: "S0",
 *     resourceGroupName: "sqlcrudtest-4799",
 *     sampleName: "AdventureWorksLT",
 *     serverName: "sqlcrudtest-6440",
 * });
 *
 * ```
 * ### Create a database min
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const database = new azurerm.sql.v20140401.Database("database", {
 *     databaseName: "testdb",
 *     location: "Japan East",
 *     resourceGroupName: "sqlcrudtest-4799",
 *     serverName: "sqlcrudtest-5961",
 * });
 *
 * ```
 * ### Update a database max
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const database = new azurerm.sql.v20140401.Database("database", {
 *     collation: "SQL_Latin1_General_CP1_CI_AS",
 *     createMode: "Default",
 *     databaseName: "testdb",
 *     edition: "Standard",
 *     location: "Japan East",
 *     maxSizeBytes: "268435456000",
 *     readScale: "Disabled",
 *     requestedServiceObjectiveId: "f1173c43-91bd-4aaa-973c-54e79e15235b",
 *     requestedServiceObjectiveName: "S0",
 *     resourceGroupName: "sqlcrudtest-4799",
 *     serverName: "sqlcrudtest-5961",
 * });
 *
 * ```
 * ### Update a database's elastic pool'
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const database = new azurerm.sql.v20140401.Database("database", {
 *     databaseName: "testdb",
 *     elasticPoolName: "7537",
 *     location: "Japan East",
 *     resourceGroupName: "sqlcrudtest-4799",
 *     serverName: "sqlcrudtest-6440",
 * });
 *
 * ```
 */
export class Database extends pulumi.CustomResource {
    /**
     * Get an existing Database resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Database {
        return new Database(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:sql/v20140401:Database';

    /**
     * Returns true if the given object is an instance of Database.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Database {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Database.__pulumiType;
    }

    /**
     * The collation of the database. If createMode is not Default, this value is ignored.
     */
    public readonly collation!: pulumi.Output<string | undefined>;
    /**
     * The containment state of the database.
     */
    public /*out*/ readonly containmentState!: pulumi.Output<number>;
    /**
     * Specifies the mode of database creation.
     *
     * Default: regular database creation.
     *
     * Copy: creates a database as a copy of an existing database. sourceDatabaseId must be specified as the resource ID of the source database.
     *
     * OnlineSecondary/NonReadableSecondary: creates a database as a (readable or nonreadable) secondary replica of an existing database. sourceDatabaseId must be specified as the resource ID of the existing primary database.
     *
     * PointInTimeRestore: Creates a database by restoring a point in time backup of an existing database. sourceDatabaseId must be specified as the resource ID of the existing database, and restorePointInTime must be specified.
     *
     * Recovery: Creates a database by restoring a geo-replicated backup. sourceDatabaseId must be specified as the recoverable database resource ID to restore.
     *
     * Restore: Creates a database by restoring a backup of a deleted database. sourceDatabaseId must be specified. If sourceDatabaseId is the database's original resource ID, then sourceDatabaseDeletionDate must be specified. Otherwise sourceDatabaseId must be the restorable dropped database resource ID and sourceDatabaseDeletionDate is ignored. restorePointInTime may also be specified to restore from an earlier point in time.
     *
     * RestoreLongTermRetentionBackup: Creates a database by restoring from a long term retention vault. recoveryServicesRecoveryPointResourceId must be specified as the recovery point resource ID.
     *
     * Copy, NonReadableSecondary, OnlineSecondary and RestoreLongTermRetentionBackup are not supported for DataWarehouse edition.
     */
    public readonly createMode!: pulumi.Output<string | undefined>;
    /**
     * The creation date of the database (ISO8601 format).
     */
    public /*out*/ readonly creationDate!: pulumi.Output<string>;
    /**
     * The current service level objective ID of the database. This is the ID of the service level objective that is currently active.
     */
    public /*out*/ readonly currentServiceObjectiveId!: pulumi.Output<string>;
    /**
     * The ID of the database.
     */
    public /*out*/ readonly databaseId!: pulumi.Output<string>;
    /**
     * The default secondary region for this database.
     */
    public /*out*/ readonly defaultSecondaryLocation!: pulumi.Output<string>;
    /**
     * This records the earliest start date and time that restore is available for this database (ISO8601 format).
     */
    public /*out*/ readonly earliestRestoreDate!: pulumi.Output<string>;
    /**
     * The edition of the database. The DatabaseEditions enumeration contains all the valid editions. If createMode is NonReadableSecondary or OnlineSecondary, this value is ignored.
     * 
     * The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition, family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API or one of the following commands:
     * 
     * ```azurecli
     * az sql db list-editions -l <location> -o table
     * ````
     * 
     * ```powershell
     * Get-AzSqlServerServiceObjective -Location <location>
     * ````
     */
    public readonly edition!: pulumi.Output<string | undefined>;
    /**
     * The name of the elastic pool the database is in. If elasticPoolName and requestedServiceObjectiveName are both updated, the value of requestedServiceObjectiveName is ignored. Not supported for DataWarehouse edition.
     */
    public readonly elasticPoolName!: pulumi.Output<string | undefined>;
    /**
     * The resource identifier of the failover group containing this database.
     */
    public /*out*/ readonly failoverGroupId!: pulumi.Output<string>;
    /**
     * Kind of database.  This is metadata used for the Azure portal experience.
     */
    public /*out*/ readonly kind!: pulumi.Output<string>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The max size of the database expressed in bytes. If createMode is not Default, this value is ignored. To see possible values, query the capabilities API (/subscriptions/{subscriptionId}/providers/Microsoft.Sql/locations/{locationID}/capabilities) referred to by operationId: "Capabilities_ListByLocation."
     */
    public readonly maxSizeBytes!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Conditional. If the database is a geo-secondary, readScale indicates whether read-only connections are allowed to this database or not. Not supported for DataWarehouse edition.
     */
    public readonly readScale!: pulumi.Output<string | undefined>;
    /**
     * The recommended indices for this database.
     */
    public /*out*/ readonly recommendedIndex!: pulumi.Output<outputs.sql.v20140401.RecommendedIndexResponse[]>;
    /**
     * Conditional. If createMode is RestoreLongTermRetentionBackup, then this value is required. Specifies the resource ID of the recovery point to restore from.
     */
    public readonly recoveryServicesRecoveryPointResourceId!: pulumi.Output<string | undefined>;
    /**
     * The configured service level objective ID of the database. This is the service level objective that is in the process of being applied to the database. Once successfully updated, it will match the value of currentServiceObjectiveId property. If requestedServiceObjectiveId and requestedServiceObjectiveName are both updated, the value of requestedServiceObjectiveId overrides the value of requestedServiceObjectiveName.
     * 
     * The list of SKUs may vary by region and support offer. To determine the service objective ids that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API.
     */
    public readonly requestedServiceObjectiveId!: pulumi.Output<string | undefined>;
    /**
     * The name of the configured service level objective of the database. This is the service level objective that is in the process of being applied to the database. Once successfully updated, it will match the value of serviceLevelObjective property. 
     * 
     * The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition, family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API or one of the following commands:
     * 
     * ```azurecli
     * az sql db list-editions -l <location> -o table
     * ````
     * 
     * ```powershell
     * Get-AzSqlServerServiceObjective -Location <location>
     * ````
     */
    public readonly requestedServiceObjectiveName!: pulumi.Output<string | undefined>;
    /**
     * Conditional. If createMode is PointInTimeRestore, this value is required. If createMode is Restore, this value is optional. Specifies the point in time (ISO8601 format) of the source database that will be restored to create the new database. Must be greater than or equal to the source database's earliestRestoreDate value.
     */
    public readonly restorePointInTime!: pulumi.Output<string | undefined>;
    /**
     * Indicates the name of the sample schema to apply when creating this database. If createMode is not Default, this value is ignored. Not supported for DataWarehouse edition.
     */
    public readonly sampleName!: pulumi.Output<string | undefined>;
    /**
     * The current service level objective of the database.
     */
    public /*out*/ readonly serviceLevelObjective!: pulumi.Output<string>;
    /**
     * The list of service tier advisors for this database. Expanded property
     */
    public /*out*/ readonly serviceTierAdvisors!: pulumi.Output<outputs.sql.v20140401.ServiceTierAdvisorResponse[]>;
    /**
     * Conditional. If createMode is Restore and sourceDatabaseId is the deleted database's original resource id when it existed (as opposed to its current restorable dropped database id), then this value is required. Specifies the time that the database was deleted.
     */
    public readonly sourceDatabaseDeletionDate!: pulumi.Output<string | undefined>;
    /**
     * Conditional. If createMode is Copy, NonReadableSecondary, OnlineSecondary, PointInTimeRestore, Recovery, or Restore, then this value is required. Specifies the resource ID of the source database. If createMode is NonReadableSecondary or OnlineSecondary, the name of the source database must be the same as the new database being created.
     */
    public readonly sourceDatabaseId!: pulumi.Output<string | undefined>;
    /**
     * The status of the database.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The transparent data encryption info for this database.
     */
    public /*out*/ readonly transparentDataEncryption!: pulumi.Output<outputs.sql.v20140401.TransparentDataEncryptionResponse[]>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
     */
    public readonly zoneRedundant!: pulumi.Output<boolean | undefined>;

    /**
     * Create a Database resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatabaseArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatabaseArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as DatabaseArgs | undefined;
            if (!args || args.databaseName === undefined) {
                throw new Error("Missing required property 'databaseName'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.serverName === undefined) {
                throw new Error("Missing required property 'serverName'");
            }
            inputs["collation"] = args ? args.collation : undefined;
            inputs["createMode"] = args ? args.createMode : undefined;
            inputs["databaseName"] = args ? args.databaseName : undefined;
            inputs["edition"] = args ? args.edition : undefined;
            inputs["elasticPoolName"] = args ? args.elasticPoolName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["maxSizeBytes"] = args ? args.maxSizeBytes : undefined;
            inputs["readScale"] = args ? args.readScale : undefined;
            inputs["recoveryServicesRecoveryPointResourceId"] = args ? args.recoveryServicesRecoveryPointResourceId : undefined;
            inputs["requestedServiceObjectiveId"] = args ? args.requestedServiceObjectiveId : undefined;
            inputs["requestedServiceObjectiveName"] = args ? args.requestedServiceObjectiveName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["restorePointInTime"] = args ? args.restorePointInTime : undefined;
            inputs["sampleName"] = args ? args.sampleName : undefined;
            inputs["serverName"] = args ? args.serverName : undefined;
            inputs["sourceDatabaseDeletionDate"] = args ? args.sourceDatabaseDeletionDate : undefined;
            inputs["sourceDatabaseId"] = args ? args.sourceDatabaseId : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["zoneRedundant"] = args ? args.zoneRedundant : undefined;
            inputs["containmentState"] = undefined /*out*/;
            inputs["creationDate"] = undefined /*out*/;
            inputs["currentServiceObjectiveId"] = undefined /*out*/;
            inputs["databaseId"] = undefined /*out*/;
            inputs["defaultSecondaryLocation"] = undefined /*out*/;
            inputs["earliestRestoreDate"] = undefined /*out*/;
            inputs["failoverGroupId"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["recommendedIndex"] = undefined /*out*/;
            inputs["serviceLevelObjective"] = undefined /*out*/;
            inputs["serviceTierAdvisors"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["transparentDataEncryption"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:sql/latest:Database" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Database.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Database resource.
 */
export interface DatabaseArgs {
    /**
     * The collation of the database. If createMode is not Default, this value is ignored.
     */
    readonly collation?: pulumi.Input<string>;
    /**
     * Specifies the mode of database creation.
     *
     * Default: regular database creation.
     *
     * Copy: creates a database as a copy of an existing database. sourceDatabaseId must be specified as the resource ID of the source database.
     *
     * OnlineSecondary/NonReadableSecondary: creates a database as a (readable or nonreadable) secondary replica of an existing database. sourceDatabaseId must be specified as the resource ID of the existing primary database.
     *
     * PointInTimeRestore: Creates a database by restoring a point in time backup of an existing database. sourceDatabaseId must be specified as the resource ID of the existing database, and restorePointInTime must be specified.
     *
     * Recovery: Creates a database by restoring a geo-replicated backup. sourceDatabaseId must be specified as the recoverable database resource ID to restore.
     *
     * Restore: Creates a database by restoring a backup of a deleted database. sourceDatabaseId must be specified. If sourceDatabaseId is the database's original resource ID, then sourceDatabaseDeletionDate must be specified. Otherwise sourceDatabaseId must be the restorable dropped database resource ID and sourceDatabaseDeletionDate is ignored. restorePointInTime may also be specified to restore from an earlier point in time.
     *
     * RestoreLongTermRetentionBackup: Creates a database by restoring from a long term retention vault. recoveryServicesRecoveryPointResourceId must be specified as the recovery point resource ID.
     *
     * Copy, NonReadableSecondary, OnlineSecondary and RestoreLongTermRetentionBackup are not supported for DataWarehouse edition.
     */
    readonly createMode?: pulumi.Input<string>;
    /**
     * The name of the database to be operated on (updated or created).
     */
    readonly databaseName: pulumi.Input<string>;
    /**
     * The edition of the database. The DatabaseEditions enumeration contains all the valid editions. If createMode is NonReadableSecondary or OnlineSecondary, this value is ignored.
     * 
     * The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition, family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API or one of the following commands:
     * 
     * ```azurecli
     * az sql db list-editions -l <location> -o table
     * ````
     * 
     * ```powershell
     * Get-AzSqlServerServiceObjective -Location <location>
     * ````
     */
    readonly edition?: pulumi.Input<string>;
    /**
     * The name of the elastic pool the database is in. If elasticPoolName and requestedServiceObjectiveName are both updated, the value of requestedServiceObjectiveName is ignored. Not supported for DataWarehouse edition.
     */
    readonly elasticPoolName?: pulumi.Input<string>;
    /**
     * Resource location.
     */
    readonly location: pulumi.Input<string>;
    /**
     * The max size of the database expressed in bytes. If createMode is not Default, this value is ignored. To see possible values, query the capabilities API (/subscriptions/{subscriptionId}/providers/Microsoft.Sql/locations/{locationID}/capabilities) referred to by operationId: "Capabilities_ListByLocation."
     */
    readonly maxSizeBytes?: pulumi.Input<string>;
    /**
     * Conditional. If the database is a geo-secondary, readScale indicates whether read-only connections are allowed to this database or not. Not supported for DataWarehouse edition.
     */
    readonly readScale?: pulumi.Input<string>;
    /**
     * Conditional. If createMode is RestoreLongTermRetentionBackup, then this value is required. Specifies the resource ID of the recovery point to restore from.
     */
    readonly recoveryServicesRecoveryPointResourceId?: pulumi.Input<string>;
    /**
     * The configured service level objective ID of the database. This is the service level objective that is in the process of being applied to the database. Once successfully updated, it will match the value of currentServiceObjectiveId property. If requestedServiceObjectiveId and requestedServiceObjectiveName are both updated, the value of requestedServiceObjectiveId overrides the value of requestedServiceObjectiveName.
     * 
     * The list of SKUs may vary by region and support offer. To determine the service objective ids that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API.
     */
    readonly requestedServiceObjectiveId?: pulumi.Input<string>;
    /**
     * The name of the configured service level objective of the database. This is the service level objective that is in the process of being applied to the database. Once successfully updated, it will match the value of serviceLevelObjective property. 
     * 
     * The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition, family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API or one of the following commands:
     * 
     * ```azurecli
     * az sql db list-editions -l <location> -o table
     * ````
     * 
     * ```powershell
     * Get-AzSqlServerServiceObjective -Location <location>
     * ````
     */
    readonly requestedServiceObjectiveName?: pulumi.Input<string>;
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Conditional. If createMode is PointInTimeRestore, this value is required. If createMode is Restore, this value is optional. Specifies the point in time (ISO8601 format) of the source database that will be restored to create the new database. Must be greater than or equal to the source database's earliestRestoreDate value.
     */
    readonly restorePointInTime?: pulumi.Input<string>;
    /**
     * Indicates the name of the sample schema to apply when creating this database. If createMode is not Default, this value is ignored. Not supported for DataWarehouse edition.
     */
    readonly sampleName?: pulumi.Input<string>;
    /**
     * The name of the server.
     */
    readonly serverName: pulumi.Input<string>;
    /**
     * Conditional. If createMode is Restore and sourceDatabaseId is the deleted database's original resource id when it existed (as opposed to its current restorable dropped database id), then this value is required. Specifies the time that the database was deleted.
     */
    readonly sourceDatabaseDeletionDate?: pulumi.Input<string>;
    /**
     * Conditional. If createMode is Copy, NonReadableSecondary, OnlineSecondary, PointInTimeRestore, Recovery, or Restore, then this value is required. Specifies the resource ID of the source database. If createMode is NonReadableSecondary or OnlineSecondary, the name of the source database must be the same as the new database being created.
     */
    readonly sourceDatabaseId?: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
     */
    readonly zoneRedundant?: pulumi.Input<boolean>;
}
