// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Sql.Latest
{
    /// <summary>
    /// Represents a database.
    /// </summary>
    public partial class Database : Pulumi.CustomResource
    {
        /// <summary>
        /// The collation of the database. If createMode is not Default, this value is ignored.
        /// </summary>
        [Output("collation")]
        public Output<string?> Collation { get; private set; } = null!;

        /// <summary>
        /// The containment state of the database.
        /// </summary>
        [Output("containmentState")]
        public Output<int> ContainmentState { get; private set; } = null!;

        /// <summary>
        /// Specifies the mode of database creation.
        /// 
        /// Default: regular database creation.
        /// 
        /// Copy: creates a database as a copy of an existing database. sourceDatabaseId must be specified as the resource ID of the source database.
        /// 
        /// OnlineSecondary/NonReadableSecondary: creates a database as a (readable or nonreadable) secondary replica of an existing database. sourceDatabaseId must be specified as the resource ID of the existing primary database.
        /// 
        /// PointInTimeRestore: Creates a database by restoring a point in time backup of an existing database. sourceDatabaseId must be specified as the resource ID of the existing database, and restorePointInTime must be specified.
        /// 
        /// Recovery: Creates a database by restoring a geo-replicated backup. sourceDatabaseId must be specified as the recoverable database resource ID to restore.
        /// 
        /// Restore: Creates a database by restoring a backup of a deleted database. sourceDatabaseId must be specified. If sourceDatabaseId is the database's original resource ID, then sourceDatabaseDeletionDate must be specified. Otherwise sourceDatabaseId must be the restorable dropped database resource ID and sourceDatabaseDeletionDate is ignored. restorePointInTime may also be specified to restore from an earlier point in time.
        /// 
        /// RestoreLongTermRetentionBackup: Creates a database by restoring from a long term retention vault. recoveryServicesRecoveryPointResourceId must be specified as the recovery point resource ID.
        /// 
        /// Copy, NonReadableSecondary, OnlineSecondary and RestoreLongTermRetentionBackup are not supported for DataWarehouse edition.
        /// </summary>
        [Output("createMode")]
        public Output<string?> CreateMode { get; private set; } = null!;

        /// <summary>
        /// The creation date of the database (ISO8601 format).
        /// </summary>
        [Output("creationDate")]
        public Output<string> CreationDate { get; private set; } = null!;

        /// <summary>
        /// The current service level objective ID of the database. This is the ID of the service level objective that is currently active.
        /// </summary>
        [Output("currentServiceObjectiveId")]
        public Output<string> CurrentServiceObjectiveId { get; private set; } = null!;

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
        /// The edition of the database. The DatabaseEditions enumeration contains all the valid editions. If createMode is NonReadableSecondary or OnlineSecondary, this value is ignored.
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
        [Output("edition")]
        public Output<string?> Edition { get; private set; } = null!;

        /// <summary>
        /// The name of the elastic pool the database is in. If elasticPoolName and requestedServiceObjectiveName are both updated, the value of requestedServiceObjectiveName is ignored. Not supported for DataWarehouse edition.
        /// </summary>
        [Output("elasticPoolName")]
        public Output<string?> ElasticPoolName { get; private set; } = null!;

        /// <summary>
        /// The resource identifier of the failover group containing this database.
        /// </summary>
        [Output("failoverGroupId")]
        public Output<string> FailoverGroupId { get; private set; } = null!;

        /// <summary>
        /// Kind of database.  This is metadata used for the Azure portal experience.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Resource location.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The max size of the database expressed in bytes. If createMode is not Default, this value is ignored. To see possible values, query the capabilities API (/subscriptions/{subscriptionId}/providers/Microsoft.Sql/locations/{locationID}/capabilities) referred to by operationId: "Capabilities_ListByLocation."
        /// </summary>
        [Output("maxSizeBytes")]
        public Output<string?> MaxSizeBytes { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Conditional. If the database is a geo-secondary, readScale indicates whether read-only connections are allowed to this database or not. Not supported for DataWarehouse edition.
        /// </summary>
        [Output("readScale")]
        public Output<string?> ReadScale { get; private set; } = null!;

        /// <summary>
        /// The recommended indices for this database.
        /// </summary>
        [Output("recommendedIndex")]
        public Output<ImmutableArray<Outputs.RecommendedIndexResponse>> RecommendedIndex { get; private set; } = null!;

        /// <summary>
        /// Conditional. If createMode is RestoreLongTermRetentionBackup, then this value is required. Specifies the resource ID of the recovery point to restore from.
        /// </summary>
        [Output("recoveryServicesRecoveryPointResourceId")]
        public Output<string?> RecoveryServicesRecoveryPointResourceId { get; private set; } = null!;

        /// <summary>
        /// The configured service level objective ID of the database. This is the service level objective that is in the process of being applied to the database. Once successfully updated, it will match the value of currentServiceObjectiveId property. If requestedServiceObjectiveId and requestedServiceObjectiveName are both updated, the value of requestedServiceObjectiveId overrides the value of requestedServiceObjectiveName.
        /// 
        /// The list of SKUs may vary by region and support offer. To determine the service objective ids that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API.
        /// </summary>
        [Output("requestedServiceObjectiveId")]
        public Output<string?> RequestedServiceObjectiveId { get; private set; } = null!;

        /// <summary>
        /// The name of the configured service level objective of the database. This is the service level objective that is in the process of being applied to the database. Once successfully updated, it will match the value of serviceLevelObjective property. 
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
        [Output("requestedServiceObjectiveName")]
        public Output<string?> RequestedServiceObjectiveName { get; private set; } = null!;

        /// <summary>
        /// Conditional. If createMode is PointInTimeRestore, this value is required. If createMode is Restore, this value is optional. Specifies the point in time (ISO8601 format) of the source database that will be restored to create the new database. Must be greater than or equal to the source database's earliestRestoreDate value.
        /// </summary>
        [Output("restorePointInTime")]
        public Output<string?> RestorePointInTime { get; private set; } = null!;

        /// <summary>
        /// Indicates the name of the sample schema to apply when creating this database. If createMode is not Default, this value is ignored. Not supported for DataWarehouse edition.
        /// </summary>
        [Output("sampleName")]
        public Output<string?> SampleName { get; private set; } = null!;

        /// <summary>
        /// The current service level objective of the database.
        /// </summary>
        [Output("serviceLevelObjective")]
        public Output<string> ServiceLevelObjective { get; private set; } = null!;

        /// <summary>
        /// The list of service tier advisors for this database. Expanded property
        /// </summary>
        [Output("serviceTierAdvisors")]
        public Output<ImmutableArray<Outputs.ServiceTierAdvisorResponse>> ServiceTierAdvisors { get; private set; } = null!;

        /// <summary>
        /// Conditional. If createMode is Restore and sourceDatabaseId is the deleted database's original resource id when it existed (as opposed to its current restorable dropped database id), then this value is required. Specifies the time that the database was deleted.
        /// </summary>
        [Output("sourceDatabaseDeletionDate")]
        public Output<string?> SourceDatabaseDeletionDate { get; private set; } = null!;

        /// <summary>
        /// Conditional. If createMode is Copy, NonReadableSecondary, OnlineSecondary, PointInTimeRestore, Recovery, or Restore, then this value is required. Specifies the resource ID of the source database. If createMode is NonReadableSecondary or OnlineSecondary, the name of the source database must be the same as the new database being created.
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
        /// The transparent data encryption info for this database.
        /// </summary>
        [Output("transparentDataEncryption")]
        public Output<ImmutableArray<Outputs.TransparentDataEncryptionResponse>> TransparentDataEncryption { get; private set; } = null!;

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
            : base("azure-nextgen:sql/latest:Database", name, args ?? new DatabaseArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Database(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:sql/latest:Database", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:sql/v20140401:Database"},
                    new Pulumi.Alias { Type = "azure-nextgen:sql/v20170301preview:Database"},
                    new Pulumi.Alias { Type = "azure-nextgen:sql/v20171001preview:Database"},
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
        /// The collation of the database. If createMode is not Default, this value is ignored.
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
        /// OnlineSecondary/NonReadableSecondary: creates a database as a (readable or nonreadable) secondary replica of an existing database. sourceDatabaseId must be specified as the resource ID of the existing primary database.
        /// 
        /// PointInTimeRestore: Creates a database by restoring a point in time backup of an existing database. sourceDatabaseId must be specified as the resource ID of the existing database, and restorePointInTime must be specified.
        /// 
        /// Recovery: Creates a database by restoring a geo-replicated backup. sourceDatabaseId must be specified as the recoverable database resource ID to restore.
        /// 
        /// Restore: Creates a database by restoring a backup of a deleted database. sourceDatabaseId must be specified. If sourceDatabaseId is the database's original resource ID, then sourceDatabaseDeletionDate must be specified. Otherwise sourceDatabaseId must be the restorable dropped database resource ID and sourceDatabaseDeletionDate is ignored. restorePointInTime may also be specified to restore from an earlier point in time.
        /// 
        /// RestoreLongTermRetentionBackup: Creates a database by restoring from a long term retention vault. recoveryServicesRecoveryPointResourceId must be specified as the recovery point resource ID.
        /// 
        /// Copy, NonReadableSecondary, OnlineSecondary and RestoreLongTermRetentionBackup are not supported for DataWarehouse edition.
        /// </summary>
        [Input("createMode")]
        public Input<string>? CreateMode { get; set; }

        /// <summary>
        /// The name of the database to be operated on (updated or created).
        /// </summary>
        [Input("databaseName", required: true)]
        public Input<string> DatabaseName { get; set; } = null!;

        /// <summary>
        /// The edition of the database. The DatabaseEditions enumeration contains all the valid editions. If createMode is NonReadableSecondary or OnlineSecondary, this value is ignored.
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
        [Input("edition")]
        public Input<string>? Edition { get; set; }

        /// <summary>
        /// The name of the elastic pool the database is in. If elasticPoolName and requestedServiceObjectiveName are both updated, the value of requestedServiceObjectiveName is ignored. Not supported for DataWarehouse edition.
        /// </summary>
        [Input("elasticPoolName")]
        public Input<string>? ElasticPoolName { get; set; }

        /// <summary>
        /// Resource location.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The max size of the database expressed in bytes. If createMode is not Default, this value is ignored. To see possible values, query the capabilities API (/subscriptions/{subscriptionId}/providers/Microsoft.Sql/locations/{locationID}/capabilities) referred to by operationId: "Capabilities_ListByLocation."
        /// </summary>
        [Input("maxSizeBytes")]
        public Input<string>? MaxSizeBytes { get; set; }

        /// <summary>
        /// Conditional. If the database is a geo-secondary, readScale indicates whether read-only connections are allowed to this database or not. Not supported for DataWarehouse edition.
        /// </summary>
        [Input("readScale")]
        public Input<string>? ReadScale { get; set; }

        /// <summary>
        /// Conditional. If createMode is RestoreLongTermRetentionBackup, then this value is required. Specifies the resource ID of the recovery point to restore from.
        /// </summary>
        [Input("recoveryServicesRecoveryPointResourceId")]
        public Input<string>? RecoveryServicesRecoveryPointResourceId { get; set; }

        /// <summary>
        /// The configured service level objective ID of the database. This is the service level objective that is in the process of being applied to the database. Once successfully updated, it will match the value of currentServiceObjectiveId property. If requestedServiceObjectiveId and requestedServiceObjectiveName are both updated, the value of requestedServiceObjectiveId overrides the value of requestedServiceObjectiveName.
        /// 
        /// The list of SKUs may vary by region and support offer. To determine the service objective ids that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API.
        /// </summary>
        [Input("requestedServiceObjectiveId")]
        public Input<string>? RequestedServiceObjectiveId { get; set; }

        /// <summary>
        /// The name of the configured service level objective of the database. This is the service level objective that is in the process of being applied to the database. Once successfully updated, it will match the value of serviceLevelObjective property. 
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
        [Input("requestedServiceObjectiveName")]
        public Input<string>? RequestedServiceObjectiveName { get; set; }

        /// <summary>
        /// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Conditional. If createMode is PointInTimeRestore, this value is required. If createMode is Restore, this value is optional. Specifies the point in time (ISO8601 format) of the source database that will be restored to create the new database. Must be greater than or equal to the source database's earliestRestoreDate value.
        /// </summary>
        [Input("restorePointInTime")]
        public Input<string>? RestorePointInTime { get; set; }

        /// <summary>
        /// Indicates the name of the sample schema to apply when creating this database. If createMode is not Default, this value is ignored. Not supported for DataWarehouse edition.
        /// </summary>
        [Input("sampleName")]
        public Input<string>? SampleName { get; set; }

        /// <summary>
        /// The name of the server.
        /// </summary>
        [Input("serverName", required: true)]
        public Input<string> ServerName { get; set; } = null!;

        /// <summary>
        /// Conditional. If createMode is Restore and sourceDatabaseId is the deleted database's original resource id when it existed (as opposed to its current restorable dropped database id), then this value is required. Specifies the time that the database was deleted.
        /// </summary>
        [Input("sourceDatabaseDeletionDate")]
        public Input<string>? SourceDatabaseDeletionDate { get; set; }

        /// <summary>
        /// Conditional. If createMode is Copy, NonReadableSecondary, OnlineSecondary, PointInTimeRestore, Recovery, or Restore, then this value is required. Specifies the resource ID of the source database. If createMode is NonReadableSecondary or OnlineSecondary, the name of the source database must be the same as the new database being created.
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
