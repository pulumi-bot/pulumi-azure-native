# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'RedisAccessKeysResponse',
    'ScheduleEntryResponse',
    'SkuResponse',
]

@pulumi.output_type
class RedisAccessKeysResponse(dict):
    """
    Redis cache access keys.
    """
    def __init__(__self__, *,
                 primary_key: str,
                 secondary_key: str):
        """
        Redis cache access keys.
        :param str primary_key: The current primary key that clients can use to authenticate with Redis cache.
        :param str secondary_key: The current secondary key that clients can use to authenticate with Redis cache.
        """
        pulumi.set(__self__, "primary_key", primary_key)
        pulumi.set(__self__, "secondary_key", secondary_key)

    @property
    @pulumi.getter(name="primaryKey")
    def primary_key(self) -> str:
        """
        The current primary key that clients can use to authenticate with Redis cache.
        """
        return pulumi.get(self, "primary_key")

    @property
    @pulumi.getter(name="secondaryKey")
    def secondary_key(self) -> str:
        """
        The current secondary key that clients can use to authenticate with Redis cache.
        """
        return pulumi.get(self, "secondary_key")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ScheduleEntryResponse(dict):
    """
    Patch schedule entry for a Premium Redis Cache.
    """
    def __init__(__self__, *,
                 day_of_week: str,
                 start_hour_utc: float,
                 maintenance_window: Optional[str] = None):
        """
        Patch schedule entry for a Premium Redis Cache.
        :param str day_of_week: Day of the week when a cache can be patched.
        :param float start_hour_utc: Start hour after which cache patching can start.
        :param str maintenance_window: ISO8601 timespan specifying how much time cache patching can take. 
        """
        pulumi.set(__self__, "day_of_week", day_of_week)
        pulumi.set(__self__, "start_hour_utc", start_hour_utc)
        if maintenance_window is not None:
            pulumi.set(__self__, "maintenance_window", maintenance_window)

    @property
    @pulumi.getter(name="dayOfWeek")
    def day_of_week(self) -> str:
        """
        Day of the week when a cache can be patched.
        """
        return pulumi.get(self, "day_of_week")

    @property
    @pulumi.getter(name="startHourUtc")
    def start_hour_utc(self) -> float:
        """
        Start hour after which cache patching can start.
        """
        return pulumi.get(self, "start_hour_utc")

    @property
    @pulumi.getter(name="maintenanceWindow")
    def maintenance_window(self) -> Optional[str]:
        """
        ISO8601 timespan specifying how much time cache patching can take. 
        """
        return pulumi.get(self, "maintenance_window")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class SkuResponse(dict):
    """
    SKU parameters supplied to the create Redis operation.
    """
    def __init__(__self__, *,
                 capacity: float,
                 family: str,
                 name: str):
        """
        SKU parameters supplied to the create Redis operation.
        :param float capacity: The size of the Redis cache to deploy. Valid values: for C (Basic/Standard) family (0, 1, 2, 3, 4, 5, 6), for P (Premium) family (1, 2, 3, 4).
        :param str family: The SKU family to use. Valid values: (C, P). (C = Basic/Standard, P = Premium).
        :param str name: The type of Redis cache to deploy. Valid values: (Basic, Standard, Premium)
        """
        pulumi.set(__self__, "capacity", capacity)
        pulumi.set(__self__, "family", family)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def capacity(self) -> float:
        """
        The size of the Redis cache to deploy. Valid values: for C (Basic/Standard) family (0, 1, 2, 3, 4, 5, 6), for P (Premium) family (1, 2, 3, 4).
        """
        return pulumi.get(self, "capacity")

    @property
    @pulumi.getter
    def family(self) -> str:
        """
        The SKU family to use. Valid values: (C, P). (C = Basic/Standard, P = Premium).
        """
        return pulumi.get(self, "family")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The type of Redis cache to deploy. Valid values: (Basic, Standard, Premium)
        """
        return pulumi.get(self, "name")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


