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
    'ListIotHubResourceKeysResult',
    'AwaitableListIotHubResourceKeysResult',
    'list_iot_hub_resource_keys',
]

@pulumi.output_type
class ListIotHubResourceKeysResult:
    """
    The list of shared access policies with a next link.
    """
    def __init__(__self__, next_link=None, value=None):
        if next_link and not isinstance(next_link, str):
            raise TypeError("Expected argument 'next_link' to be a str")
        pulumi.set(__self__, "next_link", next_link)
        if value and not isinstance(value, list):
            raise TypeError("Expected argument 'value' to be a list")
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter(name="nextLink")
    def next_link(self) -> str:
        """
        The next link.
        """
        return pulumi.get(self, "next_link")

    @property
    @pulumi.getter
    def value(self) -> Optional[List['outputs.SharedAccessSignatureAuthorizationRuleResponse']]:
        """
        The list of shared access policies.
        """
        return pulumi.get(self, "value")


class AwaitableListIotHubResourceKeysResult(ListIotHubResourceKeysResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListIotHubResourceKeysResult(
            next_link=self.next_link,
            value=self.value)


def list_iot_hub_resource_keys(resource_group_name: Optional[str] = None,
                               resource_name: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListIotHubResourceKeysResult:
    """
    Use this data source to access information about an existing resource.

    :param str resource_group_name: The name of the resource group that contains the IoT hub.
    :param str resource_name: The name of the IoT hub.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['resourceName'] = resource_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:devices/v20200801:listIotHubResourceKeys', __args__, opts=opts, typ=ListIotHubResourceKeysResult).value

    return AwaitableListIotHubResourceKeysResult(
        next_link=__ret__.next_link,
        value=__ret__.value)
