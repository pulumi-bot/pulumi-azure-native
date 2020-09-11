# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'IdentityResponse',
    'MaintenanceWindowResponse',
    'ServerPropertiesResponseDelegatedSubnetArguments',
    'SkuResponse',
    'StorageProfileResponse',
]

@pulumi.output_type
class IdentityResponse(dict):
    """
    Identity for the resource.
    """
    def __init__(__self__, *,
                 principal_id: str,
                 tenant_id: str,
                 type: Optional[str] = None):
        """
        Identity for the resource.
        :param str principal_id: The principal ID of resource identity.
        :param str tenant_id: The tenant ID of resource.
        :param str type: The identity type.
        """
        pulumi.set(__self__, "principal_id", principal_id)
        pulumi.set(__self__, "tenant_id", tenant_id)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> str:
        """
        The principal ID of resource identity.
        """
        return pulumi.get(self, "principal_id")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> str:
        """
        The tenant ID of resource.
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        The identity type.
        """
        return pulumi.get(self, "type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class MaintenanceWindowResponse(dict):
    """
    Maintenance window of a server.
    """
    def __init__(__self__, *,
                 custom_window: Optional[str] = None,
                 day_of_week: Optional[float] = None,
                 start_hour: Optional[float] = None,
                 start_minute: Optional[float] = None):
        """
        Maintenance window of a server.
        :param str custom_window: indicates whether custom window is enabled or disabled
        :param float day_of_week: day of week for maintenance window
        :param float start_hour: start hour for maintenance window
        :param float start_minute: start minute for maintenance window
        """
        if custom_window is not None:
            pulumi.set(__self__, "custom_window", custom_window)
        if day_of_week is not None:
            pulumi.set(__self__, "day_of_week", day_of_week)
        if start_hour is not None:
            pulumi.set(__self__, "start_hour", start_hour)
        if start_minute is not None:
            pulumi.set(__self__, "start_minute", start_minute)

    @property
    @pulumi.getter(name="customWindow")
    def custom_window(self) -> Optional[str]:
        """
        indicates whether custom window is enabled or disabled
        """
        return pulumi.get(self, "custom_window")

    @property
    @pulumi.getter(name="dayOfWeek")
    def day_of_week(self) -> Optional[float]:
        """
        day of week for maintenance window
        """
        return pulumi.get(self, "day_of_week")

    @property
    @pulumi.getter(name="startHour")
    def start_hour(self) -> Optional[float]:
        """
        start hour for maintenance window
        """
        return pulumi.get(self, "start_hour")

    @property
    @pulumi.getter(name="startMinute")
    def start_minute(self) -> Optional[float]:
        """
        start minute for maintenance window
        """
        return pulumi.get(self, "start_minute")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServerPropertiesResponseDelegatedSubnetArguments(dict):
    def __init__(__self__, *,
                 subnet_arm_resource_id: Optional[str] = None):
        """
        :param str subnet_arm_resource_id: delegated subnet arm resource id.
        """
        if subnet_arm_resource_id is not None:
            pulumi.set(__self__, "subnet_arm_resource_id", subnet_arm_resource_id)

    @property
    @pulumi.getter(name="subnetArmResourceId")
    def subnet_arm_resource_id(self) -> Optional[str]:
        """
        delegated subnet arm resource id.
        """
        return pulumi.get(self, "subnet_arm_resource_id")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class SkuResponse(dict):
    """
    Sku information related properties of a server.
    """
    def __init__(__self__, *,
                 name: str,
                 tier: str):
        """
        Sku information related properties of a server.
        :param str name: The name of the sku, typically, tier + family + cores, e.g. Standard_D4s_v3.
        :param str tier: The tier of the particular SKU, e.g. Burstable.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "tier", tier)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the sku, typically, tier + family + cores, e.g. Standard_D4s_v3.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tier(self) -> str:
        """
        The tier of the particular SKU, e.g. Burstable.
        """
        return pulumi.get(self, "tier")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class StorageProfileResponse(dict):
    """
    Storage Profile properties of a server
    """
    def __init__(__self__, *,
                 backup_retention_days: Optional[float] = None,
                 storage_mb: Optional[float] = None):
        """
        Storage Profile properties of a server
        :param float backup_retention_days: Backup retention days for the server.
        :param float storage_mb: Max storage allowed for a server.
        """
        if backup_retention_days is not None:
            pulumi.set(__self__, "backup_retention_days", backup_retention_days)
        if storage_mb is not None:
            pulumi.set(__self__, "storage_mb", storage_mb)

    @property
    @pulumi.getter(name="backupRetentionDays")
    def backup_retention_days(self) -> Optional[float]:
        """
        Backup retention days for the server.
        """
        return pulumi.get(self, "backup_retention_days")

    @property
    @pulumi.getter(name="storageMB")
    def storage_mb(self) -> Optional[float]:
        """
        Max storage allowed for a server.
        """
        return pulumi.get(self, "storage_mb")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


