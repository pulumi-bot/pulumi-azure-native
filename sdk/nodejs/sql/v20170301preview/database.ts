// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
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
    public static readonly __pulumiType = 'azure-nextgen:sql/v20170301preview:Database';

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
     * The ID of the database.
     */
    public /*out*/ readonly databaseId!: pulumi.Output<string>;
    /**
     * The default secondary region for this database.
     */
    public /*out*/ readonly defaultSecondaryLocation!: pulumi.Output<string>;
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
     * Resource location.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The resource identifier of the long term retention backup associated with create operation of this database.
     */
    public readonly longTermRetentionBackupResourceId!: pulumi.Output<string | undefined>;
    /**
     * The max size of the database expressed in bytes.
     */
    public readonly maxSizeBytes!: pulumi.Output<number | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The resource identifier of the recoverable database associated with create operation of this database.
     */
    public readonly recoverableDatabaseId!: pulumi.Output<string | undefined>;
    /**
     * The resource identifier of the recovery point associated with create operation of this database.
     */
    public readonly recoveryServicesRecoveryPointId!: pulumi.Output<string | undefined>;
    /**
     * The resource identifier of the restorable dropped database associated with create operation of this database.
     */
    public readonly restorableDroppedDatabaseId!: pulumi.Output<string | undefined>;
    /**
     * Specifies the point in time (ISO8601 format) of the source database that will be restored to create the new database.
     */
    public readonly restorePointInTime!: pulumi.Output<string | undefined>;
    /**
     * The name of the sample schema to apply when creating this database.
     */
    public readonly sampleName!: pulumi.Output<string | undefined>;
    /**
     * The name and tier of the SKU.
     */
    public readonly sku!: pulumi.Output<outputs.sql.v20170301preview.SkuResponse | undefined>;
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
            inputs["catalogCollation"] = args ? args.catalogCollation : undefined;
            inputs["collation"] = args ? args.collation : undefined;
            inputs["createMode"] = args ? args.createMode : undefined;
            inputs["databaseName"] = args ? args.databaseName : undefined;
            inputs["elasticPoolId"] = args ? args.elasticPoolId : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["longTermRetentionBackupResourceId"] = args ? args.longTermRetentionBackupResourceId : undefined;
            inputs["maxSizeBytes"] = args ? args.maxSizeBytes : undefined;
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
            inputs["databaseId"] = undefined /*out*/;
            inputs["defaultSecondaryLocation"] = undefined /*out*/;
            inputs["failoverGroupId"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["catalogCollation"] = undefined /*out*/;
            inputs["collation"] = undefined /*out*/;
            inputs["createMode"] = undefined /*out*/;
            inputs["creationDate"] = undefined /*out*/;
            inputs["currentServiceObjectiveName"] = undefined /*out*/;
            inputs["databaseId"] = undefined /*out*/;
            inputs["defaultSecondaryLocation"] = undefined /*out*/;
            inputs["elasticPoolId"] = undefined /*out*/;
            inputs["failoverGroupId"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["longTermRetentionBackupResourceId"] = undefined /*out*/;
            inputs["maxSizeBytes"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["recoverableDatabaseId"] = undefined /*out*/;
            inputs["recoveryServicesRecoveryPointId"] = undefined /*out*/;
            inputs["restorableDroppedDatabaseId"] = undefined /*out*/;
            inputs["restorePointInTime"] = undefined /*out*/;
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
        const aliasOpts = { aliases: [{ type: "azure-nextgen:sql/latest:Database" }, { type: "azure-nextgen:sql/v20140401:Database" }, { type: "azure-nextgen:sql/v20171001preview:Database" }, { type: "azure-nextgen:sql/v20190601preview:Database" }, { type: "azure-nextgen:sql/v20200801preview:Database" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Database.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Database resource.
 */
export interface DatabaseArgs {
    /**
     * Collation of the metadata catalog.
     */
    readonly catalogCollation?: pulumi.Input<string>;
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
    readonly createMode?: pulumi.Input<string>;
    /**
     * The name of the database.
     */
    readonly databaseName: pulumi.Input<string>;
    /**
     * The resource identifier of the elastic pool containing this database.
     */
    readonly elasticPoolId?: pulumi.Input<string>;
    /**
     * Resource location.
     */
    readonly location: pulumi.Input<string>;
    /**
     * The resource identifier of the long term retention backup associated with create operation of this database.
     */
    readonly longTermRetentionBackupResourceId?: pulumi.Input<string>;
    /**
     * The max size of the database expressed in bytes.
     */
    readonly maxSizeBytes?: pulumi.Input<number>;
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
    readonly sampleName?: pulumi.Input<string>;
    /**
     * The name of the server.
     */
    readonly serverName: pulumi.Input<string>;
    /**
     * The name and tier of the SKU.
     */
    readonly sku?: pulumi.Input<inputs.sql.v20170301preview.Sku>;
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
