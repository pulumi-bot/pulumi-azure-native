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
    'GetPatchScheduleResult',
    'AwaitableGetPatchScheduleResult',
    'get_patch_schedule',
]

@pulumi.output_type
class GetPatchScheduleResult:
    """
    Response to put/get patch schedules for Redis cache.
    """
    def __init__(__self__, name=None, schedule_entries=None, type=None):
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if schedule_entries and not isinstance(schedule_entries, list):
            raise TypeError("Expected argument 'schedule_entries' to be a list")
        pulumi.set(__self__, "schedule_entries", schedule_entries)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="scheduleEntries")
    def schedule_entries(self) -> List['outputs.ScheduleEntryResponse']:
        """
        List of patch schedules for a Redis cache.
        """
        return pulumi.get(self, "schedule_entries")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type.
        """
        return pulumi.get(self, "type")


class AwaitableGetPatchScheduleResult(GetPatchScheduleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPatchScheduleResult(
            name=self.name,
            schedule_entries=self.schedule_entries,
            type=self.type)


def get_patch_schedule(name: Optional[str] = None,
                       resource_group_name: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPatchScheduleResult:
    """
    Use this data source to access information about an existing resource.

    :param str name: Default string modeled as parameter for auto generation to work correctly.
    :param str resource_group_name: The name of the resource group.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:cache/v20171001:getPatchSchedule', __args__, opts=opts, typ=GetPatchScheduleResult).value

    return AwaitableGetPatchScheduleResult(
        name=__ret__.name,
        schedule_entries=__ret__.schedule_entries,
        type=__ret__.type)
