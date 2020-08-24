# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs

__all__ = [
    'GetDatabaseResult',
    'AwaitableGetDatabaseResult',
    'get_database',
]

@pulumi.output_type
class GetDatabaseResult:
    """
    Represents a database.
    """
    def __init__(__self__, collation=None, containment_state=None, create_mode=None, creation_date=None, current_service_objective_id=None, database_id=None, default_secondary_location=None, earliest_restore_date=None, edition=None, elastic_pool_name=None, failover_group_id=None, kind=None, location=None, max_size_bytes=None, name=None, read_scale=None, recommended_index=None, recovery_services_recovery_point_resource_id=None, requested_service_objective_id=None, requested_service_objective_name=None, restore_point_in_time=None, sample_name=None, service_level_objective=None, service_tier_advisors=None, source_database_deletion_date=None, source_database_id=None, status=None, tags=None, transparent_data_encryption=None, type=None, zone_redundant=None):
        if collation and not isinstance(collation, str):
            raise TypeError("Expected argument 'collation' to be a str")
        pulumi.set(__self__, "collation", collation)
        if containment_state and not isinstance(containment_state, float):
            raise TypeError("Expected argument 'containment_state' to be a float")
        pulumi.set(__self__, "containment_state", containment_state)
        if create_mode and not isinstance(create_mode, str):
            raise TypeError("Expected argument 'create_mode' to be a str")
        pulumi.set(__self__, "create_mode", create_mode)
        if creation_date and not isinstance(creation_date, str):
            raise TypeError("Expected argument 'creation_date' to be a str")
        pulumi.set(__self__, "creation_date", creation_date)
        if current_service_objective_id and not isinstance(current_service_objective_id, str):
            raise TypeError("Expected argument 'current_service_objective_id' to be a str")
        pulumi.set(__self__, "current_service_objective_id", current_service_objective_id)
        if database_id and not isinstance(database_id, str):
            raise TypeError("Expected argument 'database_id' to be a str")
        pulumi.set(__self__, "database_id", database_id)
        if default_secondary_location and not isinstance(default_secondary_location, str):
            raise TypeError("Expected argument 'default_secondary_location' to be a str")
        pulumi.set(__self__, "default_secondary_location", default_secondary_location)
        if earliest_restore_date and not isinstance(earliest_restore_date, str):
            raise TypeError("Expected argument 'earliest_restore_date' to be a str")
        pulumi.set(__self__, "earliest_restore_date", earliest_restore_date)
        if edition and not isinstance(edition, str):
            raise TypeError("Expected argument 'edition' to be a str")
        pulumi.set(__self__, "edition", edition)
        if elastic_pool_name and not isinstance(elastic_pool_name, str):
            raise TypeError("Expected argument 'elastic_pool_name' to be a str")
        pulumi.set(__self__, "elastic_pool_name", elastic_pool_name)
        if failover_group_id and not isinstance(failover_group_id, str):
            raise TypeError("Expected argument 'failover_group_id' to be a str")
        pulumi.set(__self__, "failover_group_id", failover_group_id)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if max_size_bytes and not isinstance(max_size_bytes, str):
            raise TypeError("Expected argument 'max_size_bytes' to be a str")
        pulumi.set(__self__, "max_size_bytes", max_size_bytes)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if read_scale and not isinstance(read_scale, str):
            raise TypeError("Expected argument 'read_scale' to be a str")
        pulumi.set(__self__, "read_scale", read_scale)
        if recommended_index and not isinstance(recommended_index, list):
            raise TypeError("Expected argument 'recommended_index' to be a list")
        pulumi.set(__self__, "recommended_index", recommended_index)
        if recovery_services_recovery_point_resource_id and not isinstance(recovery_services_recovery_point_resource_id, str):
            raise TypeError("Expected argument 'recovery_services_recovery_point_resource_id' to be a str")
        pulumi.set(__self__, "recovery_services_recovery_point_resource_id", recovery_services_recovery_point_resource_id)
        if requested_service_objective_id and not isinstance(requested_service_objective_id, str):
            raise TypeError("Expected argument 'requested_service_objective_id' to be a str")
        pulumi.set(__self__, "requested_service_objective_id", requested_service_objective_id)
        if requested_service_objective_name and not isinstance(requested_service_objective_name, str):
            raise TypeError("Expected argument 'requested_service_objective_name' to be a str")
        pulumi.set(__self__, "requested_service_objective_name", requested_service_objective_name)
        if restore_point_in_time and not isinstance(restore_point_in_time, str):
            raise TypeError("Expected argument 'restore_point_in_time' to be a str")
        pulumi.set(__self__, "restore_point_in_time", restore_point_in_time)
        if sample_name and not isinstance(sample_name, str):
            raise TypeError("Expected argument 'sample_name' to be a str")
        pulumi.set(__self__, "sample_name", sample_name)
        if service_level_objective and not isinstance(service_level_objective, str):
            raise TypeError("Expected argument 'service_level_objective' to be a str")
        pulumi.set(__self__, "service_level_objective", service_level_objective)
        if service_tier_advisors and not isinstance(service_tier_advisors, list):
            raise TypeError("Expected argument 'service_tier_advisors' to be a list")
        pulumi.set(__self__, "service_tier_advisors", service_tier_advisors)
        if source_database_deletion_date and not isinstance(source_database_deletion_date, str):
            raise TypeError("Expected argument 'source_database_deletion_date' to be a str")
        pulumi.set(__self__, "source_database_deletion_date", source_database_deletion_date)
        if source_database_id and not isinstance(source_database_id, str):
            raise TypeError("Expected argument 'source_database_id' to be a str")
        pulumi.set(__self__, "source_database_id", source_database_id)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if transparent_data_encryption and not isinstance(transparent_data_encryption, list):
            raise TypeError("Expected argument 'transparent_data_encryption' to be a list")
        pulumi.set(__self__, "transparent_data_encryption", transparent_data_encryption)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if zone_redundant and not isinstance(zone_redundant, bool):
            raise TypeError("Expected argument 'zone_redundant' to be a bool")
        pulumi.set(__self__, "zone_redundant", zone_redundant)

    @property
    @pulumi.getter
    def collation(self) -> Optional[str]:
        """
        The collation of the database. If createMode is not Default, this value is ignored.
        """
        return pulumi.get(self, "collation")

    @property
    @pulumi.getter(name="containmentState")
    def containment_state(self) -> float:
        """
        The containment state of the database.
        """
        return pulumi.get(self, "containment_state")

    @property
    @pulumi.getter(name="createMode")
    def create_mode(self) -> Optional[str]:
        """
        Specifies the mode of database creation.

        Default: regular database creation.

        Copy: creates a database as a copy of an existing database. sourceDatabaseId must be specified as the resource ID of the source database.

        OnlineSecondary/NonReadableSecondary: creates a database as a (readable or nonreadable) secondary replica of an existing database. sourceDatabaseId must be specified as the resource ID of the existing primary database.

        PointInTimeRestore: Creates a database by restoring a point in time backup of an existing database. sourceDatabaseId must be specified as the resource ID of the existing database, and restorePointInTime must be specified.

        Recovery: Creates a database by restoring a geo-replicated backup. sourceDatabaseId must be specified as the recoverable database resource ID to restore.

        Restore: Creates a database by restoring a backup of a deleted database. sourceDatabaseId must be specified. If sourceDatabaseId is the database's original resource ID, then sourceDatabaseDeletionDate must be specified. Otherwise sourceDatabaseId must be the restorable dropped database resource ID and sourceDatabaseDeletionDate is ignored. restorePointInTime may also be specified to restore from an earlier point in time.

        RestoreLongTermRetentionBackup: Creates a database by restoring from a long term retention vault. recoveryServicesRecoveryPointResourceId must be specified as the recovery point resource ID.

        Copy, NonReadableSecondary, OnlineSecondary and RestoreLongTermRetentionBackup are not supported for DataWarehouse edition.
        """
        return pulumi.get(self, "create_mode")

    @property
    @pulumi.getter(name="creationDate")
    def creation_date(self) -> str:
        """
        The creation date of the database (ISO8601 format).
        """
        return pulumi.get(self, "creation_date")

    @property
    @pulumi.getter(name="currentServiceObjectiveId")
    def current_service_objective_id(self) -> str:
        """
        The current service level objective ID of the database. This is the ID of the service level objective that is currently active.
        """
        return pulumi.get(self, "current_service_objective_id")

    @property
    @pulumi.getter(name="databaseId")
    def database_id(self) -> str:
        """
        The ID of the database.
        """
        return pulumi.get(self, "database_id")

    @property
    @pulumi.getter(name="defaultSecondaryLocation")
    def default_secondary_location(self) -> str:
        """
        The default secondary region for this database.
        """
        return pulumi.get(self, "default_secondary_location")

    @property
    @pulumi.getter(name="earliestRestoreDate")
    def earliest_restore_date(self) -> str:
        """
        This records the earliest start date and time that restore is available for this database (ISO8601 format).
        """
        return pulumi.get(self, "earliest_restore_date")

    @property
    @pulumi.getter
    def edition(self) -> Optional[str]:
        """
        The edition of the database. The DatabaseEditions enumeration contains all the valid editions. If createMode is NonReadableSecondary or OnlineSecondary, this value is ignored.
        
        The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition, family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API or one of the following commands:
        
        ```azurecli
        az sql db list-editions -l <location> -o table
        ````
        
        ```powershell
        Get-AzSqlServerServiceObjective -Location <location>
        ````
        """
        return pulumi.get(self, "edition")

    @property
    @pulumi.getter(name="elasticPoolName")
    def elastic_pool_name(self) -> Optional[str]:
        """
        The name of the elastic pool the database is in. If elasticPoolName and requestedServiceObjectiveName are both updated, the value of requestedServiceObjectiveName is ignored. Not supported for DataWarehouse edition.
        """
        return pulumi.get(self, "elastic_pool_name")

    @property
    @pulumi.getter(name="failoverGroupId")
    def failover_group_id(self) -> str:
        """
        The resource identifier of the failover group containing this database.
        """
        return pulumi.get(self, "failover_group_id")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        Kind of database.  This is metadata used for the Azure portal experience.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        Resource location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="maxSizeBytes")
    def max_size_bytes(self) -> Optional[str]:
        """
        The max size of the database expressed in bytes. If createMode is not Default, this value is ignored. To see possible values, query the capabilities API (/subscriptions/{subscriptionId}/providers/Microsoft.Sql/locations/{locationID}/capabilities) referred to by operationId: "Capabilities_ListByLocation."
        """
        return pulumi.get(self, "max_size_bytes")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="readScale")
    def read_scale(self) -> Optional[str]:
        """
        Conditional. If the database is a geo-secondary, readScale indicates whether read-only connections are allowed to this database or not. Not supported for DataWarehouse edition.
        """
        return pulumi.get(self, "read_scale")

    @property
    @pulumi.getter(name="recommendedIndex")
    def recommended_index(self) -> List['outputs.RecommendedIndexResponse']:
        """
        The recommended indices for this database.
        """
        return pulumi.get(self, "recommended_index")

    @property
    @pulumi.getter(name="recoveryServicesRecoveryPointResourceId")
    def recovery_services_recovery_point_resource_id(self) -> Optional[str]:
        """
        Conditional. If createMode is RestoreLongTermRetentionBackup, then this value is required. Specifies the resource ID of the recovery point to restore from.
        """
        return pulumi.get(self, "recovery_services_recovery_point_resource_id")

    @property
    @pulumi.getter(name="requestedServiceObjectiveId")
    def requested_service_objective_id(self) -> Optional[str]:
        """
        The configured service level objective ID of the database. This is the service level objective that is in the process of being applied to the database. Once successfully updated, it will match the value of currentServiceObjectiveId property. If requestedServiceObjectiveId and requestedServiceObjectiveName are both updated, the value of requestedServiceObjectiveId overrides the value of requestedServiceObjectiveName.
        
        The list of SKUs may vary by region and support offer. To determine the service objective ids that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API.
        """
        return pulumi.get(self, "requested_service_objective_id")

    @property
    @pulumi.getter(name="requestedServiceObjectiveName")
    def requested_service_objective_name(self) -> Optional[str]:
        """
        The name of the configured service level objective of the database. This is the service level objective that is in the process of being applied to the database. Once successfully updated, it will match the value of serviceLevelObjective property. 
        
        The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition, family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API or one of the following commands:
        
        ```azurecli
        az sql db list-editions -l <location> -o table
        ````
        
        ```powershell
        Get-AzSqlServerServiceObjective -Location <location>
        ````
        """
        return pulumi.get(self, "requested_service_objective_name")

    @property
    @pulumi.getter(name="restorePointInTime")
    def restore_point_in_time(self) -> Optional[str]:
        """
        Conditional. If createMode is PointInTimeRestore, this value is required. If createMode is Restore, this value is optional. Specifies the point in time (ISO8601 format) of the source database that will be restored to create the new database. Must be greater than or equal to the source database's earliestRestoreDate value.
        """
        return pulumi.get(self, "restore_point_in_time")

    @property
    @pulumi.getter(name="sampleName")
    def sample_name(self) -> Optional[str]:
        """
        Indicates the name of the sample schema to apply when creating this database. If createMode is not Default, this value is ignored. Not supported for DataWarehouse edition.
        """
        return pulumi.get(self, "sample_name")

    @property
    @pulumi.getter(name="serviceLevelObjective")
    def service_level_objective(self) -> str:
        """
        The current service level objective of the database.
        """
        return pulumi.get(self, "service_level_objective")

    @property
    @pulumi.getter(name="serviceTierAdvisors")
    def service_tier_advisors(self) -> List['outputs.ServiceTierAdvisorResponse']:
        """
        The list of service tier advisors for this database. Expanded property
        """
        return pulumi.get(self, "service_tier_advisors")

    @property
    @pulumi.getter(name="sourceDatabaseDeletionDate")
    def source_database_deletion_date(self) -> Optional[str]:
        """
        Conditional. If createMode is Restore and sourceDatabaseId is the deleted database's original resource id when it existed (as opposed to its current restorable dropped database id), then this value is required. Specifies the time that the database was deleted.
        """
        return pulumi.get(self, "source_database_deletion_date")

    @property
    @pulumi.getter(name="sourceDatabaseId")
    def source_database_id(self) -> Optional[str]:
        """
        Conditional. If createMode is Copy, NonReadableSecondary, OnlineSecondary, PointInTimeRestore, Recovery, or Restore, then this value is required. Specifies the resource ID of the source database. If createMode is NonReadableSecondary or OnlineSecondary, the name of the source database must be the same as the new database being created.
        """
        return pulumi.get(self, "source_database_id")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The status of the database.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="transparentDataEncryption")
    def transparent_data_encryption(self) -> List['outputs.TransparentDataEncryptionResponse']:
        """
        The transparent data encryption info for this database.
        """
        return pulumi.get(self, "transparent_data_encryption")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="zoneRedundant")
    def zone_redundant(self) -> Optional[bool]:
        """
        Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
        """
        return pulumi.get(self, "zone_redundant")


class AwaitableGetDatabaseResult(GetDatabaseResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatabaseResult(
            collation=self.collation,
            containment_state=self.containment_state,
            create_mode=self.create_mode,
            creation_date=self.creation_date,
            current_service_objective_id=self.current_service_objective_id,
            database_id=self.database_id,
            default_secondary_location=self.default_secondary_location,
            earliest_restore_date=self.earliest_restore_date,
            edition=self.edition,
            elastic_pool_name=self.elastic_pool_name,
            failover_group_id=self.failover_group_id,
            kind=self.kind,
            location=self.location,
            max_size_bytes=self.max_size_bytes,
            name=self.name,
            read_scale=self.read_scale,
            recommended_index=self.recommended_index,
            recovery_services_recovery_point_resource_id=self.recovery_services_recovery_point_resource_id,
            requested_service_objective_id=self.requested_service_objective_id,
            requested_service_objective_name=self.requested_service_objective_name,
            restore_point_in_time=self.restore_point_in_time,
            sample_name=self.sample_name,
            service_level_objective=self.service_level_objective,
            service_tier_advisors=self.service_tier_advisors,
            source_database_deletion_date=self.source_database_deletion_date,
            source_database_id=self.source_database_id,
            status=self.status,
            tags=self.tags,
            transparent_data_encryption=self.transparent_data_encryption,
            type=self.type,
            zone_redundant=self.zone_redundant)


def get_database(expand: Optional[str] = None,
                 name: Optional[str] = None,
                 resource_group_name: Optional[str] = None,
                 server_name: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatabaseResult:
    """
    Use this data source to access information about an existing resource.

    :param str expand: A comma separated list of child objects to expand in the response. Possible properties: serviceTierAdvisors, transparentDataEncryption.
    :param str name: The name of the database to be retrieved.
    :param str resource_group_name: The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
    :param str server_name: The name of the server.
    """
    __args__ = dict()
    __args__['expand'] = expand
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    __args__['serverName'] = server_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:sql/v20140401:getDatabase', __args__, opts=opts, typ=GetDatabaseResult).value

    return AwaitableGetDatabaseResult(
        collation=__ret__.collation,
        containment_state=__ret__.containment_state,
        create_mode=__ret__.create_mode,
        creation_date=__ret__.creation_date,
        current_service_objective_id=__ret__.current_service_objective_id,
        database_id=__ret__.database_id,
        default_secondary_location=__ret__.default_secondary_location,
        earliest_restore_date=__ret__.earliest_restore_date,
        edition=__ret__.edition,
        elastic_pool_name=__ret__.elastic_pool_name,
        failover_group_id=__ret__.failover_group_id,
        kind=__ret__.kind,
        location=__ret__.location,
        max_size_bytes=__ret__.max_size_bytes,
        name=__ret__.name,
        read_scale=__ret__.read_scale,
        recommended_index=__ret__.recommended_index,
        recovery_services_recovery_point_resource_id=__ret__.recovery_services_recovery_point_resource_id,
        requested_service_objective_id=__ret__.requested_service_objective_id,
        requested_service_objective_name=__ret__.requested_service_objective_name,
        restore_point_in_time=__ret__.restore_point_in_time,
        sample_name=__ret__.sample_name,
        service_level_objective=__ret__.service_level_objective,
        service_tier_advisors=__ret__.service_tier_advisors,
        source_database_deletion_date=__ret__.source_database_deletion_date,
        source_database_id=__ret__.source_database_id,
        status=__ret__.status,
        tags=__ret__.tags,
        transparent_data_encryption=__ret__.transparent_data_encryption,
        type=__ret__.type,
        zone_redundant=__ret__.zone_redundant)
