// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * A database resource.
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
        return new Database(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:sql/v20171001preview:Database';

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
     * Time in minutes after which database is automatically paused. A value of -1 means that automatic pause is disabled
     */
    public readonly autoPauseDelay!: pulumi.Output<number | undefined>;
    /**
     * Collation of the metadata catalog.
     */
    public readonly catalogCollation!: pulumi.Output<string | undefined>;
    /**
     * The collation of the database.
     */
    public readonly collation!: pulumi.Output<string | undefined>;
    /**
     * Specifies the mode of database creation.
     * 
     * Default: regular database creation.
     * 
     * Copy: creates a database as a copy of an existing database. sourceDatabaseId must be specified as the resource ID of the source database.
     * 
     * Secondary: creates a database as a secondary replica of an existing database. sourceDatabaseId must be specified as the resource ID of the existing primary database.
     * 
     * PointInTimeRestore: Creates a database by restoring a point in time backup of an existing database. sourceDatabaseId must be specified as the resource ID of the existing database, and restorePointInTime must be specified.
     * 
     * Recovery: Creates a database by restoring a geo-replicated backup. sourceDatabaseId must be specified as the recoverable database resource ID to restore.
     * 
     * Restore: Creates a database by restoring a backup of a deleted database. sourceDatabaseId must be specified. If sourceDatabaseId is the database's original resource ID, then sourceDatabaseDeletionDate must be specified. Otherwise sourceDatabaseId must be the restorable dropped database resource ID and sourceDatabaseDeletionDate is ignored. restorePointInTime may also be specified to restore from an earlier point in time.
     * 
     * RestoreLongTermRetentionBackup: Creates a database by restoring from a long term retention vault. recoveryServicesRecoveryPointResourceId must be specified as the recovery point resource ID.
     * 
     * Copy, Secondary, and RestoreLongTermRetentionBackup are not supported for DataWarehouse edition.
     */
    public readonly createMode!: pulumi.Output<string | undefined>;
    /**
     * The creation date of the database (ISO8601 format).
     */
    public /*out*/ readonly creationDate!: pulumi.Output<string>;
    /**
     * The current service level objective name of the database.
     */
    public /*out*/ readonly currentServiceObjectiveName!: pulumi.Output<string>;
    /**
     * The name and tier of the SKU.
     */
    public /*out*/ readonly currentSku!: pulumi.Output<outputs.sql.v20171001preview.SkuResponse>;
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
     * The resource identifier of the elastic pool containing this database.
     */
    public readonly elasticPoolId!: pulumi.Output<string | undefined>;
    /**
     * Failover Group resource identifier that this database belongs to.
     */
    public /*out*/ readonly failoverGroupId!: pulumi.Output<string>;
    /**
     * Kind of database. This is metadata used for the Azure portal experience.
     */
    public /*out*/ readonly kind!: pulumi.Output<string>;
    /**
     * The license type to apply for this database. `LicenseIncluded` if you need a license, or `BasePrice` if you have a license and are eligible for the Azure Hybrid Benefit.
     */
    public readonly licenseType!: pulumi.Output<string | undefined>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The resource identifier of the long term retention backup associated with create operation of this database.
     */
    public readonly longTermRetentionBackupResourceId!: pulumi.Output<string | undefined>;
    /**
     * Resource that manages the database.
     */
    public /*out*/ readonly managedBy!: pulumi.Output<string>;
    /**
     * The max log size for this database.
     */
    public /*out*/ readonly maxLogSizeBytes!: pulumi.Output<number>;
    /**
     * The max size of the database expressed in bytes.
     */
    public readonly maxSizeBytes!: pulumi.Output<number | undefined>;
    /**
     * Minimal capacity that database will always have allocated, if not paused
     */
    public readonly minCapacity!: pulumi.Output<number | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The date when database was paused by user configuration or action (ISO8601 format). Null if the database is ready.
     */
    public /*out*/ readonly pausedDate!: pulumi.Output<string>;
    /**
     * The number of readonly secondary replicas associated with the database to which readonly application intent connections may be routed. This property is only settable for Hyperscale edition databases.
     */
    public readonly readReplicaCount!: pulumi.Output<number | undefined>;
    /**
     * If enabled, connections that have application intent set to readonly in their connection string may be routed to a readonly secondary replica. This property is only settable for Premium and Business Critical databases.
     */
    public readonly readScale!: pulumi.Output<string | undefined>;
    /**
     * The resource identifier of the recoverable database associated with create operation of this database.
     */
    public readonly recoverableDatabaseId!: pulumi.Output<string | undefined>;
    /**
     * The resource identifier of the recovery point associated with create operation of this database.
     */
    public readonly recoveryServicesRecoveryPointId!: pulumi.Output<string | undefined>;
    /**
     * The requested service level objective name of the database.
     */
    public /*out*/ readonly requestedServiceObjectiveName!: pulumi.Output<string>;
    /**
     * The resource identifier of the restorable dropped database associated with create operation of this database.
     */
    public readonly restorableDroppedDatabaseId!: pulumi.Output<string | undefined>;
    /**
     * Specifies the point in time (ISO8601 format) of the source database that will be restored to create the new database.
     */
    public readonly restorePointInTime!: pulumi.Output<string | undefined>;
    /**
     * The date when database was resumed by user action or database login (ISO8601 format). Null if the database is paused.
     */
    public /*out*/ readonly resumedDate!: pulumi.Output<string>;
    /**
     * The name of the sample schema to apply when creating this database.
     */
    public readonly sampleName!: pulumi.Output<string | undefined>;
    /**
     * The database SKU.
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
    public readonly sku!: pulumi.Output<outputs.sql.v20171001preview.SkuResponse | undefined>;
    /**
     * Specifies the time that the database was deleted.
     */
    public readonly sourceDatabaseDeletionDate!: pulumi.Output<string | undefined>;
    /**
     * The resource identifier of the source database associated with create operation of this database.
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
    constructor(name: string, args: DatabaseArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.databaseName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'databaseName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.serverName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'serverName'");
            }
            inputs["autoPauseDelay"] = args ? args.autoPauseDelay : undefined;
            inputs["catalogCollation"] = args ? args.catalogCollation : undefined;
            inputs["collation"] = args ? args.collation : undefined;
            inputs["createMode"] = args ? args.createMode : undefined;
            inputs["databaseName"] = args ? args.databaseName : undefined;
            inputs["elasticPoolId"] = args ? args.elasticPoolId : undefined;
            inputs["licenseType"] = args ? args.licenseType : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["longTermRetentionBackupResourceId"] = args ? args.longTermRetentionBackupResourceId : undefined;
            inputs["maxSizeBytes"] = args ? args.maxSizeBytes : undefined;
            inputs["minCapacity"] = args ? args.minCapacity : undefined;
            inputs["readReplicaCount"] = args ? args.readReplicaCount : undefined;
            inputs["readScale"] = args ? args.readScale : undefined;
            inputs["recoverableDatabaseId"] = args ? args.recoverableDatabaseId : undefined;
            inputs["recoveryServicesRecoveryPointId"] = args ? args.recoveryServicesRecoveryPointId : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["restorableDroppedDatabaseId"] = args ? args.restorableDroppedDatabaseId : undefined;
            inputs["restorePointInTime"] = args ? args.restorePointInTime : undefined;
            inputs["sampleName"] = args ? args.sampleName : undefined;
            inputs["serverName"] = args ? args.serverName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["sourceDatabaseDeletionDate"] = args ? args.sourceDatabaseDeletionDate : undefined;
            inputs["sourceDatabaseId"] = args ? args.sourceDatabaseId : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["zoneRedundant"] = args ? args.zoneRedundant : undefined;
            inputs["creationDate"] = undefined /*out*/;
            inputs["currentServiceObjectiveName"] = undefined /*out*/;
            inputs["currentSku"] = undefined /*out*/;
            inputs["databaseId"] = undefined /*out*/;
            inputs["defaultSecondaryLocation"] = undefined /*out*/;
            inputs["earliestRestoreDate"] = undefined /*out*/;
            inputs["failoverGroupId"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["managedBy"] = undefined /*out*/;
            inputs["maxLogSizeBytes"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["pausedDate"] = undefined /*out*/;
            inputs["requestedServiceObjectiveName"] = undefined /*out*/;
            inputs["resumedDate"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["autoPauseDelay"] = undefined /*out*/;
            inputs["catalogCollation"] = undefined /*out*/;
            inputs["collation"] = undefined /*out*/;
            inputs["createMode"] = undefined /*out*/;
            inputs["creationDate"] = undefined /*out*/;
            inputs["currentServiceObjectiveName"] = undefined /*out*/;
            inputs["currentSku"] = undefined /*out*/;
            inputs["databaseId"] = undefined /*out*/;
            inputs["defaultSecondaryLocation"] = undefined /*out*/;
            inputs["earliestRestoreDate"] = undefined /*out*/;
            inputs["elasticPoolId"] = undefined /*out*/;
            inputs["failoverGroupId"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["licenseType"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["longTermRetentionBackupResourceId"] = undefined /*out*/;
            inputs["managedBy"] = undefined /*out*/;
            inputs["maxLogSizeBytes"] = undefined /*out*/;
            inputs["maxSizeBytes"] = undefined /*out*/;
            inputs["minCapacity"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["pausedDate"] = undefined /*out*/;
            inputs["readReplicaCount"] = undefined /*out*/;
            inputs["readScale"] = undefined /*out*/;
            inputs["recoverableDatabaseId"] = undefined /*out*/;
            inputs["recoveryServicesRecoveryPointId"] = undefined /*out*/;
            inputs["requestedServiceObjectiveName"] = undefined /*out*/;
            inputs["restorableDroppedDatabaseId"] = undefined /*out*/;
            inputs["restorePointInTime"] = undefined /*out*/;
            inputs["resumedDate"] = undefined /*out*/;
            inputs["sampleName"] = undefined /*out*/;
            inputs["sku"] = undefined /*out*/;
            inputs["sourceDatabaseDeletionDate"] = undefined /*out*/;
            inputs["sourceDatabaseId"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["zoneRedundant"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:sql/latest:Database" }, { type: "azure-nextgen:sql/v20140401:Database" }, { type: "azure-nextgen:sql/v20170301preview:Database" }, { type: "azure-nextgen:sql/v20190601preview:Database" }, { type: "azure-nextgen:sql/v20200202preview:Database" }, { type: "azure-nextgen:sql/v20200801preview:Database" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Database.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Database resource.
 */
export interface DatabaseArgs {
    /**
     * Time in minutes after which database is automatically paused. A value of -1 means that automatic pause is disabled
     */
    readonly autoPauseDelay?: pulumi.Input<number>;
    /**
     * Collation of the metadata catalog.
     */
    readonly catalogCollation?: pulumi.Input<string | enums.sql.v20171001preview.CatalogCollationType>;
    /**
     * The collation of the database.
     */
    readonly collation?: pulumi.Input<string>;
    /**
     * Specifies the mode of database creation.
     * 
     * Default: regular database creation.
     * 
     * Copy: creates a database as a copy of an existing database. sourceDatabaseId must be specified as the resource ID of the source database.
     * 
     * Secondary: creates a database as a secondary replica of an existing database. sourceDatabaseId must be specified as the resource ID of the existing primary database.
     * 
     * PointInTimeRestore: Creates a database by restoring a point in time backup of an existing database. sourceDatabaseId must be specified as the resource ID of the existing database, and restorePointInTime must be specified.
     * 
     * Recovery: Creates a database by restoring a geo-replicated backup. sourceDatabaseId must be specified as the recoverable database resource ID to restore.
     * 
     * Restore: Creates a database by restoring a backup of a deleted database. sourceDatabaseId must be specified. If sourceDatabaseId is the database's original resource ID, then sourceDatabaseDeletionDate must be specified. Otherwise sourceDatabaseId must be the restorable dropped database resource ID and sourceDatabaseDeletionDate is ignored. restorePointInTime may also be specified to restore from an earlier point in time.
     * 
     * RestoreLongTermRetentionBackup: Creates a database by restoring from a long term retention vault. recoveryServicesRecoveryPointResourceId must be specified as the recovery point resource ID.
     * 
     * Copy, Secondary, and RestoreLongTermRetentionBackup are not supported for DataWarehouse edition.
     */
    readonly createMode?: pulumi.Input<string | enums.sql.v20171001preview.CreateMode>;
    /**
     * The name of the database.
     */
    readonly databaseName: pulumi.Input<string>;
    /**
     * The resource identifier of the elastic pool containing this database.
     */
    readonly elasticPoolId?: pulumi.Input<string>;
    /**
     * The license type to apply for this database. `LicenseIncluded` if you need a license, or `BasePrice` if you have a license and are eligible for the Azure Hybrid Benefit.
     */
    readonly licenseType?: pulumi.Input<string | enums.sql.v20171001preview.DatabaseLicenseType>;
    /**
     * Resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The resource identifier of the long term retention backup associated with create operation of this database.
     */
    readonly longTermRetentionBackupResourceId?: pulumi.Input<string>;
    /**
     * The max size of the database expressed in bytes.
     */
    readonly maxSizeBytes?: pulumi.Input<number>;
    /**
     * Minimal capacity that database will always have allocated, if not paused
     */
    readonly minCapacity?: pulumi.Input<number>;
    /**
     * The number of readonly secondary replicas associated with the database to which readonly application intent connections may be routed. This property is only settable for Hyperscale edition databases.
     */
    readonly readReplicaCount?: pulumi.Input<number>;
    /**
     * If enabled, connections that have application intent set to readonly in their connection string may be routed to a readonly secondary replica. This property is only settable for Premium and Business Critical databases.
     */
    readonly readScale?: pulumi.Input<string | enums.sql.v20171001preview.DatabaseReadScale>;
    /**
     * The resource identifier of the recoverable database associated with create operation of this database.
     */
    readonly recoverableDatabaseId?: pulumi.Input<string>;
    /**
     * The resource identifier of the recovery point associated with create operation of this database.
     */
    readonly recoveryServicesRecoveryPointId?: pulumi.Input<string>;
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The resource identifier of the restorable dropped database associated with create operation of this database.
     */
    readonly restorableDroppedDatabaseId?: pulumi.Input<string>;
    /**
     * Specifies the point in time (ISO8601 format) of the source database that will be restored to create the new database.
     */
    readonly restorePointInTime?: pulumi.Input<string>;
    /**
     * The name of the sample schema to apply when creating this database.
     */
    readonly sampleName?: pulumi.Input<string | enums.sql.v20171001preview.SampleName>;
    /**
     * The name of the server.
     */
    readonly serverName: pulumi.Input<string>;
    /**
     * The database SKU.
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
    readonly sku?: pulumi.Input<inputs.sql.v20171001preview.Sku>;
    /**
     * Specifies the time that the database was deleted.
     */
    readonly sourceDatabaseDeletionDate?: pulumi.Input<string>;
    /**
     * The resource identifier of the source database associated with create operation of this database.
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
