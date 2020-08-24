# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'ListTopicSharedAccessKeysResult',
    'AwaitableListTopicSharedAccessKeysResult',
    'list_topic_shared_access_keys',
]

@pulumi.output_type
class ListTopicSharedAccessKeysResult:
    """
    Shared access keys of the Topic
    """
    def __init__(__self__, key1=None, key2=None):
        if key1 and not isinstance(key1, str):
            raise TypeError("Expected argument 'key1' to be a str")
        pulumi.set(__self__, "key1", key1)
        if key2 and not isinstance(key2, str):
            raise TypeError("Expected argument 'key2' to be a str")
        pulumi.set(__self__, "key2", key2)

    @property
    @pulumi.getter
    def key1(self) -> Optional[str]:
        """
        Shared access key1 for the topic.
        """
        return pulumi.get(self, "key1")

    @property
    @pulumi.getter
    def key2(self) -> Optional[str]:
        """
        Shared access key2 for the topic.
        """
        return pulumi.get(self, "key2")


class AwaitableListTopicSharedAccessKeysResult(ListTopicSharedAccessKeysResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListTopicSharedAccessKeysResult(
            key1=self.key1,
            key2=self.key2)


def list_topic_shared_access_keys(resource_group_name: Optional[str] = None,
                                  topic_name: Optional[str] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListTopicSharedAccessKeysResult:
    """
    Use this data source to access information about an existing resource.

    :param str resource_group_name: The name of the resource group within the user's subscription.
    :param str topic_name: Name of the topic
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['topicName'] = topic_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:eventgrid/v20180101:listTopicSharedAccessKeys', __args__, opts=opts, typ=ListTopicSharedAccessKeysResult).value

    return AwaitableListTopicSharedAccessKeysResult(
        key1=__ret__.key1,
        key2=__ret__.key2)
