// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Sql.V20171001Preview
{
    /// <summary>
    /// A database resource.
    /// </summary>
    public partial class Database : Pulumi.CustomResource
    {
        /// <summary>
        /// Time in minutes after which database is automatically paused. A value of -1 means that automatic pause is disabled
        /// </summary>
        [Output("autoPauseDelay")]
        public Output<int?> AutoPauseDelay { get; private set; } = null!;

        /// <summary>
        /// Collation of the metadata catalog.
        /// </summary>
        [Output("catalogCollation")]
        public Output<string?> CatalogCollation { get; private set; } = null!;

        /// <summary>
        /// The collation of the database.
        /// </summary>
        [Output("collation")]
        public Output<string?> Collation { get; private set; } = null!;

        /// <summary>
        /// Specifies the mode of database creation.
        /// 
        /// Default: regular database creation.
        /// 
        /// Copy: creates a database as a copy of an existing database. sourceDatabaseId must be specified as the resource ID of the source database.
        /// 
        /// Secondary: creates a database as a secondary replica of an existing database. sourceDatabaseId must be specified as the resource ID of the existing primary database.
        /// 
        /// PointInTimeRestore: Creates a database by restoring a point in time backup of an existing database. sourceDatabaseId must be specified as the resource ID of the existing database, and restorePointInTime must be specified.
        /// 
        /// Recovery: Creates a database by restoring a geo-replicated backup. sourceDatabaseId must be specified as the recoverable database resource ID to restore.
        /// 
        /// Restore: Creates a database by restoring a backup of a deleted database. sourceDatabaseId must be specified. If sourceDatabaseId is the database's original resource ID, then sourceDatabaseDeletionDate must be specified. Otherwise sourceDatabaseId must be the restorable dropped database resource ID and sourceDatabaseDeletionDate is ignored. restorePointInTime may also be specified to restore from an earlier point in time.
        /// 
        /// RestoreLongTermRetentionBackup: Creates a database by restoring from a long term retention vault. recoveryServicesRecoveryPointResourceId must be specified as the recovery point resource ID.
        /// 
        /// Copy, Secondary, and RestoreLongTermRetentionBackup are not supported for DataWarehouse edition.
        /// </summary>
        [Output("createMode")]
        public Output<string?> CreateMode { get; private set; } = null!;

        /// <summary>
        /// The creation date of the database (ISO8601 format).
        /// </summary>
        [Output("creationDate")]
        public Output<string> CreationDate { get; private set; } = null!;

        /// <summary>
        /// The current service level objective name of the database.
        /// </summary>
        [Output("currentServiceObjectiveName")]
        public Output<string> CurrentServiceObjectiveName { get; private set; } = null!;

        /// <summary>
        /// The name and tier of the SKU.
        /// </summary>
        [Output("currentSku")]
        public Output<Outputs.SkuResponse> CurrentSku { get; private set; } = null!;

        /// <summary>
        /// The ID of the database.
        /// </summary>
        [Output("databaseId")]
        public Output<string> DatabaseId { get; private set; } = null!;

        /// <summary>
        /// The default secondary region for this database.
        /// </summary>
        [Output("defaultSecondaryLocation")]
        public Output<string> DefaultSecondaryLocation { get; private set; } = null!;

        /// <summary>
        /// This records the earliest start date and time that restore is available for this database (ISO8601 format).
        /// </summary>
        [Output("earliestRestoreDate")]
        public Output<string> EarliestRestoreDate { get; private set; } = null!;

        /// <summary>
        /// The resource identifier of the elastic pool containing this database.
        /// </summary>
        [Output("elasticPoolId")]
        public Output<string?> ElasticPoolId { get; private set; } = null!;

        /// <summary>
        /// Failover Group resource identifier that this database belongs to.
        /// </summary>
        [Output("failoverGroupId")]
        public Output<string> FailoverGroupId { get; private set; } = null!;

        /// <summary>
        /// Kind of database. This is metadata used for the Azure portal experience.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// The license type to apply for this database. `LicenseIncluded` if you need a license, or `BasePrice` if you have a license and are eligible for the Azure Hybrid Benefit.
        /// </summary>
        [Output("licenseType")]
        public Output<string?> LicenseType { get; private set; } = null!;

        /// <summary>
        /// Resource location.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The resource identifier of the long term retention backup associated with create operation of this database.
        /// </summary>
        [Output("longTermRetentionBackupResourceId")]
        public Output<string?> LongTermRetentionBackupResourceId { get; private set; } = null!;

        /// <summary>
        /// Resource that manages the database.
        /// </summary>
        [Output("managedBy")]
        public Output<string> ManagedBy { get; private set; } = null!;

        /// <summary>
        /// The max log size for this database.
        /// </summary>
        [Output("maxLogSizeBytes")]
        public Output<int> MaxLogSizeBytes { get; private set; } = null!;

        /// <summary>
        /// The max size of the database expressed in bytes.
        /// </summary>
        [Output("maxSizeBytes")]
        public Output<int?> MaxSizeBytes { get; private set; } = null!;

        /// <summary>
        /// Minimal capacity that database will always have allocated, if not paused
        /// </summary>
        [Output("minCapacity")]
        public Output<double?> MinCapacity { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The date when database was paused by user configuration or action (ISO8601 format). Null if the database is ready.
        /// </summary>
        [Output("pausedDate")]
        public Output<string> PausedDate { get; private set; } = null!;

        /// <summary>
        /// The number of readonly secondary replicas associated with the database to which readonly application intent connections may be routed. This property is only settable for Hyperscale edition databases.
        /// </summary>
        [Output("readReplicaCount")]
        public Output<int?> ReadReplicaCount { get; private set; } = null!;

        /// <summary>
        /// If enabled, connections that have application intent set to readonly in their connection string may be routed to a readonly secondary replica. This property is only settable for Premium and Business Critical databases.
        /// </summary>
        [Output("readScale")]
        public Output<string?> ReadScale { get; private set; } = null!;

        /// <summary>
        /// The resource identifier of the recoverable database associated with create operation of this database.
        /// </summary>
        [Output("recoverableDatabaseId")]
        public Output<string?> RecoverableDatabaseId { get; private set; } = null!;

        /// <summary>
        /// The resource identifier of the recovery point associated with create operation of this database.
        /// </summary>
        [Output("recoveryServicesRecoveryPointId")]
        public Output<string?> RecoveryServicesRecoveryPointId { get; private set; } = null!;

        /// <summary>
        /// The requested service level objective name of the database.
        /// </summary>
        [Output("requestedServiceObjectiveName")]
        public Output<string> RequestedServiceObjectiveName { get; private set; } = null!;

        /// <summary>
        /// The resource identifier of the restorable dropped database associated with create operation of this database.
        /// </summary>
        [Output("restorableDroppedDatabaseId")]
        public Output<string?> RestorableDroppedDatabaseId { get; private set; } = null!;

        /// <summary>
        /// Specifies the point in time (ISO8601 format) of the source database that will be restored to create the new database.
        /// </summary>
        [Output("restorePointInTime")]
        public Output<string?> RestorePointInTime { get; private set; } = null!;

        /// <summary>
        /// The date when database was resumed by user action or database login (ISO8601 format). Null if the database is paused.
        /// </summary>
        [Output("resumedDate")]
        public Output<string> ResumedDate { get; private set; } = null!;

        /// <summary>
        /// The name of the sample schema to apply when creating this database.
        /// </summary>
        [Output("sampleName")]
        public Output<string?> SampleName { get; private set; } = null!;

        /// <summary>
        /// The database SKU.
        /// 
        /// The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition, family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API or one of the following commands:
        /// 
        /// ```azurecli
        /// az sql db list-editions -l &lt;location&gt; -o table
        /// ````
        /// 
        /// ```powershell
        /// Get-AzSqlServerServiceObjective -Location &lt;location&gt;
        /// ````
        /// </summary>
        [Output("sku")]
        public Output<Outputs.SkuResponse?> Sku { get; private set; } = null!;

        /// <summary>
        /// Specifies the time that the database was deleted.
        /// </summary>
        [Output("sourceDatabaseDeletionDate")]
        public Output<string?> SourceDatabaseDeletionDate { get; private set; } = null!;

        /// <summary>
        /// The resource identifier of the source database associated with create operation of this database.
        /// </summary>
        [Output("sourceDatabaseId")]
        public Output<string?> SourceDatabaseId { get; private set; } = null!;

        /// <summary>
        /// The status of the database.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
        /// </summary>
        [Output("zoneRedundant")]
        public Output<bool?> ZoneRedundant { get; private set; } = null!;


        /// <summary>
        /// Create a Database resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Database(string name, DatabaseArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:sql/v20171001preview:Database", name, args ?? new DatabaseArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Database(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:sql/v20171001preview:Database", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:sql/latest:Database"},
                    new Pulumi.Alias { Type = "azure-nextgen:sql/v20140401:Database"},
                    new Pulumi.Alias { Type = "azure-nextgen:sql/v20170301preview:Database"},
                    new Pulumi.Alias { Type = "azure-nextgen:sql/v20190601preview:Database"},
                    new Pulumi.Alias { Type = "azure-nextgen:sql/v20200801preview:Database"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Database resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Database Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Database(name, id, options);
        }
    }

    public sealed class DatabaseArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Time in minutes after which database is automatically paused. A value of -1 means that automatic pause is disabled
        /// </summary>
        [Input("autoPauseDelay")]
        public Input<int>? AutoPauseDelay { get; set; }

        /// <summary>
        /// Collation of the metadata catalog.
        /// </summary>
        [Input("catalogCollation")]
        public Input<string>? CatalogCollation { get; set; }

        /// <summary>
        /// The collation of the database.
        /// </summary>
        [Input("collation")]
        public Input<string>? Collation { get; set; }

        /// <summary>
        /// Specifies the mode of database creation.
        /// 
        /// Default: regular database creation.
        /// 
        /// Copy: creates a database as a copy of an existing database. sourceDatabaseId must be specified as the resource ID of the source database.
        /// 
        /// Secondary: creates a database as a secondary replica of an existing database. sourceDatabaseId must be specified as the resource ID of the existing primary database.
        /// 
        /// PointInTimeRestore: Creates a database by restoring a point in time backup of an existing database. sourceDatabaseId must be specified as the resource ID of the existing database, and restorePointInTime must be specified.
        /// 
        /// Recovery: Creates a database by restoring a geo-replicated backup. sourceDatabaseId must be specified as the recoverable database resource ID to restore.
        /// 
        /// Restore: Creates a database by restoring a backup of a deleted database. sourceDatabaseId must be specified. If sourceDatabaseId is the database's original resource ID, then sourceDatabaseDeletionDate must be specified. Otherwise sourceDatabaseId must be the restorable dropped database resource ID and sourceDatabaseDeletionDate is ignored. restorePointInTime may also be specified to restore from an earlier point in time.
        /// 
        /// RestoreLongTermRetentionBackup: Creates a database by restoring from a long term retention vault. recoveryServicesRecoveryPointResourceId must be specified as the recovery point resource ID.
        /// 
        /// Copy, Secondary, and RestoreLongTermRetentionBackup are not supported for DataWarehouse edition.
        /// </summary>
        [Input("createMode")]
        public Input<string>? CreateMode { get; set; }

        /// <summary>
        /// The name of the database.
        /// </summary>
        [Input("databaseName", required: true)]
        public Input<string> DatabaseName { get; set; } = null!;

        /// <summary>
        /// The resource identifier of the elastic pool containing this database.
        /// </summary>
        [Input("elasticPoolId")]
        public Input<string>? ElasticPoolId { get; set; }

        /// <summary>
        /// The license type to apply for this database. `LicenseIncluded` if you need a license, or `BasePrice` if you have a license and are eligible for the Azure Hybrid Benefit.
        /// </summary>
        [Input("licenseType")]
        public Input<string>? LicenseType { get; set; }

        /// <summary>
        /// Resource location.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The resource identifier of the long term retention backup associated with create operation of this database.
        /// </summary>
        [Input("longTermRetentionBackupResourceId")]
        public Input<string>? LongTermRetentionBackupResourceId { get; set; }

        /// <summary>
        /// The max size of the database expressed in bytes.
        /// </summary>
        [Input("maxSizeBytes")]
        public Input<int>? MaxSizeBytes { get; set; }

        /// <summary>
        /// Minimal capacity that database will always have allocated, if not paused
        /// </summary>
        [Input("minCapacity")]
        public Input<double>? MinCapacity { get; set; }

        /// <summary>
        /// The number of readonly secondary replicas associated with the database to which readonly application intent connections may be routed. This property is only settable for Hyperscale edition databases.
        /// </summary>
        [Input("readReplicaCount")]
        public Input<int>? ReadReplicaCount { get; set; }

        /// <summary>
        /// If enabled, connections that have application intent set to readonly in their connection string may be routed to a readonly secondary replica. This property is only settable for Premium and Business Critical databases.
        /// </summary>
        [Input("readScale")]
        public Input<string>? ReadScale { get; set; }

        /// <summary>
        /// The resource identifier of the recoverable database associated with create operation of this database.
        /// </summary>
        [Input("recoverableDatabaseId")]
        public Input<string>? RecoverableDatabaseId { get; set; }

        /// <summary>
        /// The resource identifier of the recovery point associated with create operation of this database.
        /// </summary>
        [Input("recoveryServicesRecoveryPointId")]
        public Input<string>? RecoveryServicesRecoveryPointId { get; set; }

        /// <summary>
        /// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The resource identifier of the restorable dropped database associated with create operation of this database.
        /// </summary>
        [Input("restorableDroppedDatabaseId")]
        public Input<string>? RestorableDroppedDatabaseId { get; set; }

        /// <summary>
        /// Specifies the point in time (ISO8601 format) of the source database that will be restored to create the new database.
        /// </summary>
        [Input("restorePointInTime")]
        public Input<string>? RestorePointInTime { get; set; }

        /// <summary>
        /// The name of the sample schema to apply when creating this database.
        /// </summary>
        [Input("sampleName")]
        public Input<string>? SampleName { get; set; }

        /// <summary>
        /// The name of the server.
        /// </summary>
        [Input("serverName", required: true)]
        public Input<string> ServerName { get; set; } = null!;

        /// <summary>
        /// The database SKU.
        /// 
        /// The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition, family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API or one of the following commands:
        /// 
        /// ```azurecli
        /// az sql db list-editions -l &lt;location&gt; -o table
        /// ````
        /// 
        /// ```powershell
        /// Get-AzSqlServerServiceObjective -Location &lt;location&gt;
        /// ````
        /// </summary>
        [Input("sku")]
        public Input<Inputs.SkuArgs>? Sku { get; set; }

        /// <summary>
        /// Specifies the time that the database was deleted.
        /// </summary>
        [Input("sourceDatabaseDeletionDate")]
        public Input<string>? SourceDatabaseDeletionDate { get; set; }

        /// <summary>
        /// The resource identifier of the source database associated with create operation of this database.
        /// </summary>
        [Input("sourceDatabaseId")]
        public Input<string>? SourceDatabaseId { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
        /// </summary>
        [Input("zoneRedundant")]
        public Input<bool>? ZoneRedundant { get; set; }

        public DatabaseArgs()
        {
        }
    }
}
