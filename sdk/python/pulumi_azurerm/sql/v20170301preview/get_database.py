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
    A database resource.
    """
    def __init__(__self__, catalog_collation=None, collation=None, create_mode=None, creation_date=None, current_service_objective_name=None, database_id=None, default_secondary_location=None, elastic_pool_id=None, failover_group_id=None, kind=None, location=None, long_term_retention_backup_resource_id=None, max_size_bytes=None, name=None, recoverable_database_id=None, recovery_services_recovery_point_id=None, restorable_dropped_database_id=None, restore_point_in_time=None, sample_name=None, sku=None, source_database_deletion_date=None, source_database_id=None, status=None, tags=None, type=None, zone_redundant=None):
        if catalog_collation and not isinstance(catalog_collation, str):
            raise TypeError("Expected argument 'catalog_collation' to be a str")
        pulumi.set(__self__, "catalog_collation", catalog_collation)
        if collation and not isinstance(collation, str):
            raise TypeError("Expected argument 'collation' to be a str")
        pulumi.set(__self__, "collation", collation)
        if create_mode and not isinstance(create_mode, str):
            raise TypeError("Expected argument 'create_mode' to be a str")
        pulumi.set(__self__, "create_mode", create_mode)
        if creation_date and not isinstance(creation_date, str):
            raise TypeError("Expected argument 'creation_date' to be a str")
        pulumi.set(__self__, "creation_date", creation_date)
        if current_service_objective_name and not isinstance(current_service_objective_name, str):
            raise TypeError("Expected argument 'current_service_objective_name' to be a str")
        pulumi.set(__self__, "current_service_objective_name", current_service_objective_name)
        if database_id and not isinstance(database_id, str):
            raise TypeError("Expected argument 'database_id' to be a str")
        pulumi.set(__self__, "database_id", database_id)
        if default_secondary_location and not isinstance(default_secondary_location, str):
            raise TypeError("Expected argument 'default_secondary_location' to be a str")
        pulumi.set(__self__, "default_secondary_location", default_secondary_location)
        if elastic_pool_id and not isinstance(elastic_pool_id, str):
            raise TypeError("Expected argument 'elastic_pool_id' to be a str")
        pulumi.set(__self__, "elastic_pool_id", elastic_pool_id)
        if failover_group_id and not isinstance(failover_group_id, str):
            raise TypeError("Expected argument 'failover_group_id' to be a str")
        pulumi.set(__self__, "failover_group_id", failover_group_id)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if long_term_retention_backup_resource_id and not isinstance(long_term_retention_backup_resource_id, str):
            raise TypeError("Expected argument 'long_term_retention_backup_resource_id' to be a str")
        pulumi.set(__self__, "long_term_retention_backup_resource_id", long_term_retention_backup_resource_id)
        if max_size_bytes and not isinstance(max_size_bytes, float):
            raise TypeError("Expected argument 'max_size_bytes' to be a float")
        pulumi.set(__self__, "max_size_bytes", max_size_bytes)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if recoverable_database_id and not isinstance(recoverable_database_id, str):
            raise TypeError("Expected argument 'recoverable_database_id' to be a str")
        pulumi.set(__self__, "recoverable_database_id", recoverable_database_id)
        if recovery_services_recovery_point_id and not isinstance(recovery_services_recovery_point_id, str):
            raise TypeError("Expected argument 'recovery_services_recovery_point_id' to be a str")
        pulumi.set(__self__, "recovery_services_recovery_point_id", recovery_services_recovery_point_id)
        if restorable_dropped_database_id and not isinstance(restorable_dropped_database_id, str):
            raise TypeError("Expected argument 'restorable_dropped_database_id' to be a str")
        pulumi.set(__self__, "restorable_dropped_database_id", restorable_dropped_database_id)
        if restore_point_in_time and not isinstance(restore_point_in_time, str):
            raise TypeError("Expected argument 'restore_point_in_time' to be a str")
        pulumi.set(__self__, "restore_point_in_time", restore_point_in_time)
        if sample_name and not isinstance(sample_name, str):
            raise TypeError("Expected argument 'sample_name' to be a str")
        pulumi.set(__self__, "sample_name", sample_name)
        if sku and not isinstance(sku, dict):
            raise TypeError("Expected argument 'sku' to be a dict")
        pulumi.set(__self__, "sku", sku)
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
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if zone_redundant and not isinstance(zone_redundant, bool):
            raise TypeError("Expected argument 'zone_redundant' to be a bool")
        pulumi.set(__self__, "zone_redundant", zone_redundant)

    @property
    @pulumi.getter(name="catalogCollation")
    def catalog_collation(self) -> Optional[str]:
        """
        Collation of the metadata catalog.
        """
        return pulumi.get(self, "catalog_collation")

    @property
    @pulumi.getter
    def collation(self) -> Optional[str]:
        """
        The collation of the database.
        """
        return pulumi.get(self, "collation")

    @property
    @pulumi.getter(name="createMode")
    def create_mode(self) -> Optional[str]:
        """
        Specifies the mode of database creation.
        
        Default: regular database creation.
        
        Copy: creates a database as a copy of an existing database. sourceDatabaseId must be specified as the resource ID of the source database.
        
        Secondary: creates a database as a secondary replica of an existing database. sourceDatabaseId must be specified as the resource ID of the existing primary database.
        
        PointInTimeRestore: Creates a database by restoring a point in time backup of an existing database. sourceDatabaseId must be specified as the resource ID of the existing database, and restorePointInTime must be specified.
        
        Recovery: Creates a database by restoring a geo-replicated backup. sourceDatabaseId must be specified as the recoverable database resource ID to restore.
        
        Restore: Creates a database by restoring a backup of a deleted database. sourceDatabaseId must be specified. If sourceDatabaseId is the database's original resource ID, then sourceDatabaseDeletionDate must be specified. Otherwise sourceDatabaseId must be the restorable dropped database resource ID and sourceDatabaseDeletionDate is ignored. restorePointInTime may also be specified to restore from an earlier point in time.
        
        RestoreLongTermRetentionBackup: Creates a database by restoring from a long term retention vault. recoveryServicesRecoveryPointResourceId must be specified as the recovery point resource ID.
        
        Copy, Secondary, and RestoreLongTermRetentionBackup are not supported for DataWarehouse edition.
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
    @pulumi.getter(name="currentServiceObjectiveName")
    def current_service_objective_name(self) -> str:
        """
        The current service level objective name of the database.
        """
        return pulumi.get(self, "current_service_objective_name")

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
    @pulumi.getter(name="elasticPoolId")
    def elastic_pool_id(self) -> Optional[str]:
        """
        The resource identifier of the elastic pool containing this database.
        """
        return pulumi.get(self, "elastic_pool_id")

    @property
    @pulumi.getter(name="failoverGroupId")
    def failover_group_id(self) -> str:
        """
        Failover Group resource identifier that this database belongs to.
        """
        return pulumi.get(self, "failover_group_id")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        Kind of database. This is metadata used for the Azure portal experience.
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
    @pulumi.getter(name="longTermRetentionBackupResourceId")
    def long_term_retention_backup_resource_id(self) -> Optional[str]:
        """
        The resource identifier of the long term retention backup associated with create operation of this database.
        """
        return pulumi.get(self, "long_term_retention_backup_resource_id")

    @property
    @pulumi.getter(name="maxSizeBytes")
    def max_size_bytes(self) -> Optional[float]:
        """
        The max size of the database expressed in bytes.
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
    @pulumi.getter(name="recoverableDatabaseId")
    def recoverable_database_id(self) -> Optional[str]:
        """
        The resource identifier of the recoverable database associated with create operation of this database.
        """
        return pulumi.get(self, "recoverable_database_id")

    @property
    @pulumi.getter(name="recoveryServicesRecoveryPointId")
    def recovery_services_recovery_point_id(self) -> Optional[str]:
        """
        The resource identifier of the recovery point associated with create operation of this database.
        """
        return pulumi.get(self, "recovery_services_recovery_point_id")

    @property
    @pulumi.getter(name="restorableDroppedDatabaseId")
    def restorable_dropped_database_id(self) -> Optional[str]:
        """
        The resource identifier of the restorable dropped database associated with create operation of this database.
        """
        return pulumi.get(self, "restorable_dropped_database_id")

    @property
    @pulumi.getter(name="restorePointInTime")
    def restore_point_in_time(self) -> Optional[str]:
        """
        Specifies the point in time (ISO8601 format) of the source database that will be restored to create the new database.
        """
        return pulumi.get(self, "restore_point_in_time")

    @property
    @pulumi.getter(name="sampleName")
    def sample_name(self) -> Optional[str]:
        """
        The name of the sample schema to apply when creating this database.
        """
        return pulumi.get(self, "sample_name")

    @property
    @pulumi.getter
    def sku(self) -> Optional['outputs.SkuResponse']:
        """
        The name and tier of the SKU.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter(name="sourceDatabaseDeletionDate")
    def source_database_deletion_date(self) -> Optional[str]:
        """
        Specifies the time that the database was deleted.
        """
        return pulumi.get(self, "source_database_deletion_date")

    @property
    @pulumi.getter(name="sourceDatabaseId")
    def source_database_id(self) -> Optional[str]:
        """
        The resource identifier of the source database associated with create operation of this database.
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
            catalog_collation=self.catalog_collation,
            collation=self.collation,
            create_mode=self.create_mode,
            creation_date=self.creation_date,
            current_service_objective_name=self.current_service_objective_name,
            database_id=self.database_id,
            default_secondary_location=self.default_secondary_location,
            elastic_pool_id=self.elastic_pool_id,
            failover_group_id=self.failover_group_id,
            kind=self.kind,
            location=self.location,
            long_term_retention_backup_resource_id=self.long_term_retention_backup_resource_id,
            max_size_bytes=self.max_size_bytes,
            name=self.name,
            recoverable_database_id=self.recoverable_database_id,
            recovery_services_recovery_point_id=self.recovery_services_recovery_point_id,
            restorable_dropped_database_id=self.restorable_dropped_database_id,
            restore_point_in_time=self.restore_point_in_time,
            sample_name=self.sample_name,
            sku=self.sku,
            source_database_deletion_date=self.source_database_deletion_date,
            source_database_id=self.source_database_id,
            status=self.status,
            tags=self.tags,
            type=self.type,
            zone_redundant=self.zone_redundant)


def get_database(database_name: Optional[str] = None,
                 resource_group_name: Optional[str] = None,
                 server_name: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatabaseResult:
    """
    Use this data source to access information about an existing resource.

    :param str database_name: The name of the database.
    :param str resource_group_name: The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
    :param str server_name: The name of the server.
    """
    __args__ = dict()
    __args__['databaseName'] = database_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['serverName'] = server_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:sql/v20170301preview:getDatabase', __args__, opts=opts, typ=GetDatabaseResult).value

    return AwaitableGetDatabaseResult(
        catalog_collation=__ret__.catalog_collation,
        collation=__ret__.collation,
        create_mode=__ret__.create_mode,
        creation_date=__ret__.creation_date,
        current_service_objective_name=__ret__.current_service_objective_name,
        database_id=__ret__.database_id,
        default_secondary_location=__ret__.default_secondary_location,
        elastic_pool_id=__ret__.elastic_pool_id,
        failover_group_id=__ret__.failover_group_id,
        kind=__ret__.kind,
        location=__ret__.location,
        long_term_retention_backup_resource_id=__ret__.long_term_retention_backup_resource_id,
        max_size_bytes=__ret__.max_size_bytes,
        name=__ret__.name,
        recoverable_database_id=__ret__.recoverable_database_id,
        recovery_services_recovery_point_id=__ret__.recovery_services_recovery_point_id,
        restorable_dropped_database_id=__ret__.restorable_dropped_database_id,
        restore_point_in_time=__ret__.restore_point_in_time,
        sample_name=__ret__.sample_name,
        sku=__ret__.sku,
        source_database_deletion_date=__ret__.source_database_deletion_date,
        source_database_id=__ret__.source_database_id,
        status=__ret__.status,
        tags=__ret__.tags,
        type=__ret__.type,
        zone_redundant=__ret__.zone_redundant)
