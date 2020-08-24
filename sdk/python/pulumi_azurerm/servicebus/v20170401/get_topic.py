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
    'GetTopicResult',
    'AwaitableGetTopicResult',
    'get_topic',
]

@pulumi.output_type
class GetTopicResult:
    """
    Description of topic resource.
    """
    def __init__(__self__, accessed_at=None, auto_delete_on_idle=None, count_details=None, created_at=None, default_message_time_to_live=None, duplicate_detection_history_time_window=None, enable_batched_operations=None, enable_express=None, enable_partitioning=None, max_size_in_megabytes=None, name=None, requires_duplicate_detection=None, size_in_bytes=None, status=None, subscription_count=None, support_ordering=None, type=None, updated_at=None):
        if accessed_at and not isinstance(accessed_at, str):
            raise TypeError("Expected argument 'accessed_at' to be a str")
        pulumi.set(__self__, "accessed_at", accessed_at)
        if auto_delete_on_idle and not isinstance(auto_delete_on_idle, str):
            raise TypeError("Expected argument 'auto_delete_on_idle' to be a str")
        pulumi.set(__self__, "auto_delete_on_idle", auto_delete_on_idle)
        if count_details and not isinstance(count_details, dict):
            raise TypeError("Expected argument 'count_details' to be a dict")
        pulumi.set(__self__, "count_details", count_details)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if default_message_time_to_live and not isinstance(default_message_time_to_live, str):
            raise TypeError("Expected argument 'default_message_time_to_live' to be a str")
        pulumi.set(__self__, "default_message_time_to_live", default_message_time_to_live)
        if duplicate_detection_history_time_window and not isinstance(duplicate_detection_history_time_window, str):
            raise TypeError("Expected argument 'duplicate_detection_history_time_window' to be a str")
        pulumi.set(__self__, "duplicate_detection_history_time_window", duplicate_detection_history_time_window)
        if enable_batched_operations and not isinstance(enable_batched_operations, bool):
            raise TypeError("Expected argument 'enable_batched_operations' to be a bool")
        pulumi.set(__self__, "enable_batched_operations", enable_batched_operations)
        if enable_express and not isinstance(enable_express, bool):
            raise TypeError("Expected argument 'enable_express' to be a bool")
        pulumi.set(__self__, "enable_express", enable_express)
        if enable_partitioning and not isinstance(enable_partitioning, bool):
            raise TypeError("Expected argument 'enable_partitioning' to be a bool")
        pulumi.set(__self__, "enable_partitioning", enable_partitioning)
        if max_size_in_megabytes and not isinstance(max_size_in_megabytes, float):
            raise TypeError("Expected argument 'max_size_in_megabytes' to be a float")
        pulumi.set(__self__, "max_size_in_megabytes", max_size_in_megabytes)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if requires_duplicate_detection and not isinstance(requires_duplicate_detection, bool):
            raise TypeError("Expected argument 'requires_duplicate_detection' to be a bool")
        pulumi.set(__self__, "requires_duplicate_detection", requires_duplicate_detection)
        if size_in_bytes and not isinstance(size_in_bytes, float):
            raise TypeError("Expected argument 'size_in_bytes' to be a float")
        pulumi.set(__self__, "size_in_bytes", size_in_bytes)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if subscription_count and not isinstance(subscription_count, float):
            raise TypeError("Expected argument 'subscription_count' to be a float")
        pulumi.set(__self__, "subscription_count", subscription_count)
        if support_ordering and not isinstance(support_ordering, bool):
            raise TypeError("Expected argument 'support_ordering' to be a bool")
        pulumi.set(__self__, "support_ordering", support_ordering)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="accessedAt")
    def accessed_at(self) -> str:
        """
        Last time the message was sent, or a request was received, for this topic.
        """
        return pulumi.get(self, "accessed_at")

    @property
    @pulumi.getter(name="autoDeleteOnIdle")
    def auto_delete_on_idle(self) -> Optional[str]:
        """
        ISO 8601 timespan idle interval after which the topic is automatically deleted. The minimum duration is 5 minutes.
        """
        return pulumi.get(self, "auto_delete_on_idle")

    @property
    @pulumi.getter(name="countDetails")
    def count_details(self) -> 'outputs.MessageCountDetailsResponse':
        """
        Message count details
        """
        return pulumi.get(self, "count_details")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        Exact time the message was created.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="defaultMessageTimeToLive")
    def default_message_time_to_live(self) -> Optional[str]:
        """
        ISO 8601 Default message timespan to live value. This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
        """
        return pulumi.get(self, "default_message_time_to_live")

    @property
    @pulumi.getter(name="duplicateDetectionHistoryTimeWindow")
    def duplicate_detection_history_time_window(self) -> Optional[str]:
        """
        ISO8601 timespan structure that defines the duration of the duplicate detection history. The default value is 10 minutes.
        """
        return pulumi.get(self, "duplicate_detection_history_time_window")

    @property
    @pulumi.getter(name="enableBatchedOperations")
    def enable_batched_operations(self) -> Optional[bool]:
        """
        Value that indicates whether server-side batched operations are enabled.
        """
        return pulumi.get(self, "enable_batched_operations")

    @property
    @pulumi.getter(name="enableExpress")
    def enable_express(self) -> Optional[bool]:
        """
        Value that indicates whether Express Entities are enabled. An express topic holds a message in memory temporarily before writing it to persistent storage.
        """
        return pulumi.get(self, "enable_express")

    @property
    @pulumi.getter(name="enablePartitioning")
    def enable_partitioning(self) -> Optional[bool]:
        """
        Value that indicates whether the topic to be partitioned across multiple message brokers is enabled.
        """
        return pulumi.get(self, "enable_partitioning")

    @property
    @pulumi.getter(name="maxSizeInMegabytes")
    def max_size_in_megabytes(self) -> Optional[float]:
        """
        Maximum size of the topic in megabytes, which is the size of the memory allocated for the topic. Default is 1024.
        """
        return pulumi.get(self, "max_size_in_megabytes")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="requiresDuplicateDetection")
    def requires_duplicate_detection(self) -> Optional[bool]:
        """
        Value indicating if this topic requires duplicate detection.
        """
        return pulumi.get(self, "requires_duplicate_detection")

    @property
    @pulumi.getter(name="sizeInBytes")
    def size_in_bytes(self) -> float:
        """
        Size of the topic, in bytes.
        """
        return pulumi.get(self, "size_in_bytes")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        Enumerates the possible values for the status of a messaging entity.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="subscriptionCount")
    def subscription_count(self) -> float:
        """
        Number of subscriptions.
        """
        return pulumi.get(self, "subscription_count")

    @property
    @pulumi.getter(name="supportOrdering")
    def support_ordering(self) -> Optional[bool]:
        """
        Value that indicates whether the topic supports ordering.
        """
        return pulumi.get(self, "support_ordering")

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


class AwaitableGetTopicResult(GetTopicResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTopicResult(
            accessed_at=self.accessed_at,
            auto_delete_on_idle=self.auto_delete_on_idle,
            count_details=self.count_details,
            created_at=self.created_at,
            default_message_time_to_live=self.default_message_time_to_live,
            duplicate_detection_history_time_window=self.duplicate_detection_history_time_window,
            enable_batched_operations=self.enable_batched_operations,
            enable_express=self.enable_express,
            enable_partitioning=self.enable_partitioning,
            max_size_in_megabytes=self.max_size_in_megabytes,
            name=self.name,
            requires_duplicate_detection=self.requires_duplicate_detection,
            size_in_bytes=self.size_in_bytes,
            status=self.status,
            subscription_count=self.subscription_count,
            support_ordering=self.support_ordering,
            type=self.type,
            updated_at=self.updated_at)


def get_topic(name: Optional[str] = None,
              namespace_name: Optional[str] = None,
              resource_group_name: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTopicResult:
    """
    Use this data source to access information about an existing resource.

    :param str name: The topic name.
    :param str namespace_name: The namespace name
    :param str resource_group_name: Name of the Resource group within the Azure subscription.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['namespaceName'] = namespace_name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:servicebus/v20170401:getTopic', __args__, opts=opts, typ=GetTopicResult).value

    return AwaitableGetTopicResult(
        accessed_at=__ret__.accessed_at,
        auto_delete_on_idle=__ret__.auto_delete_on_idle,
        count_details=__ret__.count_details,
        created_at=__ret__.created_at,
        default_message_time_to_live=__ret__.default_message_time_to_live,
        duplicate_detection_history_time_window=__ret__.duplicate_detection_history_time_window,
        enable_batched_operations=__ret__.enable_batched_operations,
        enable_express=__ret__.enable_express,
        enable_partitioning=__ret__.enable_partitioning,
        max_size_in_megabytes=__ret__.max_size_in_megabytes,
        name=__ret__.name,
        requires_duplicate_detection=__ret__.requires_duplicate_detection,
        size_in_bytes=__ret__.size_in_bytes,
        status=__ret__.status,
        subscription_count=__ret__.subscription_count,
        support_ordering=__ret__.support_ordering,
        type=__ret__.type,
        updated_at=__ret__.updated_at)
