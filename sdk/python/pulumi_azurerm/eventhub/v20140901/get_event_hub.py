# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'GetEventHubResult',
    'AwaitableGetEventHubResult',
    'get_event_hub',
]

@pulumi.output_type
class GetEventHubResult:
    """
    Single item in List or Get Event Hub operation
    """
    def __init__(__self__, created_at=None, location=None, message_retention_in_days=None, name=None, partition_count=None, partition_ids=None, status=None, type=None, updated_at=None):
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if message_retention_in_days and not isinstance(message_retention_in_days, float):
            raise TypeError("Expected argument 'message_retention_in_days' to be a float")
        pulumi.set(__self__, "message_retention_in_days", message_retention_in_days)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if partition_count and not isinstance(partition_count, float):
            raise TypeError("Expected argument 'partition_count' to be a float")
        pulumi.set(__self__, "partition_count", partition_count)
        if partition_ids and not isinstance(partition_ids, list):
            raise TypeError("Expected argument 'partition_ids' to be a list")
        pulumi.set(__self__, "partition_ids", partition_ids)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        Exact time the Event Hub was created.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def location(self) -> Optional[str]:
        """
        Resource location
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="messageRetentionInDays")
    def message_retention_in_days(self) -> Optional[float]:
        """
        Number of days to retain the events for this Event Hub.
        """
        return pulumi.get(self, "message_retention_in_days")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="partitionCount")
    def partition_count(self) -> Optional[float]:
        """
        Number of partitions created for the Event Hub.
        """
        return pulumi.get(self, "partition_count")

    @property
    @pulumi.getter(name="partitionIds")
    def partition_ids(self) -> List[str]:
        """
        Current number of shards on the Event Hub.
        """
        return pulumi.get(self, "partition_ids")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        Enumerates the possible values for the status of the Event Hub.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        """
        The exact time the message was updated.
        """
        return pulumi.get(self, "updated_at")


class AwaitableGetEventHubResult(GetEventHubResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEventHubResult(
            created_at=self.created_at,
            location=self.location,
            message_retention_in_days=self.message_retention_in_days,
            name=self.name,
            partition_count=self.partition_count,
            partition_ids=self.partition_ids,
            status=self.status,
            type=self.type,
            updated_at=self.updated_at)


def get_event_hub(name: Optional[str] = None,
                  namespace_name: Optional[str] = None,
                  resource_group_name: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEventHubResult:
    """
    Use this data source to access information about an existing resource.

    :param str name: The Event Hub name
    :param str namespace_name: The Namespace name
    :param str resource_group_name: Name of the resource group within the azure subscription.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['namespaceName'] = namespace_name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:eventhub/v20140901:getEventHub', __args__, opts=opts, typ=GetEventHubResult).value

    return AwaitableGetEventHubResult(
        created_at=__ret__.created_at,
        location=__ret__.location,
        message_retention_in_days=__ret__.message_retention_in_days,
        name=__ret__.name,
        partition_count=__ret__.partition_count,
        partition_ids=__ret__.partition_ids,
        status=__ret__.status,
        type=__ret__.type,
        updated_at=__ret__.updated_at)
