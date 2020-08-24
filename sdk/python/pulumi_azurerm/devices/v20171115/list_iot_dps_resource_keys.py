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
    'ListIotDpsResourceKeysResult',
    'AwaitableListIotDpsResourceKeysResult',
    'list_iot_dps_resource_keys',
]

@pulumi.output_type
class ListIotDpsResourceKeysResult:
    """
    List of shared access keys.
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
    def value(self) -> Optional[List['outputs.SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponse']]:
        """
        The list of shared access policies.
        """
        return pulumi.get(self, "value")


class AwaitableListIotDpsResourceKeysResult(ListIotDpsResourceKeysResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListIotDpsResourceKeysResult(
            next_link=self.next_link,
            value=self.value)


def list_iot_dps_resource_keys(provisioning_service_name: Optional[str] = None,
                               resource_group_name: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListIotDpsResourceKeysResult:
    """
    Use this data source to access information about an existing resource.

    :param str provisioning_service_name: The provisioning service name to get the shared access keys for.
    :param str resource_group_name: resource group name
    """
    __args__ = dict()
    __args__['provisioningServiceName'] = provisioning_service_name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:devices/v20171115:listIotDpsResourceKeys', __args__, opts=opts, typ=ListIotDpsResourceKeysResult).value

    return AwaitableListIotDpsResourceKeysResult(
        next_link=__ret__.next_link,
        value=__ret__.value)
