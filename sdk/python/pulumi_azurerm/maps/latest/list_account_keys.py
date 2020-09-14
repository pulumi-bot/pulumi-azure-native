# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'ListAccountKeysResult',
    'AwaitableListAccountKeysResult',
    'list_account_keys',
]

@pulumi.output_type
class ListAccountKeysResult:
    """
    The set of keys which can be used to access the Maps REST APIs. Two keys are provided for key rotation without interruption.
    """
    def __init__(__self__, primary_key=None, secondary_key=None):
        if primary_key and not isinstance(primary_key, str):
            raise TypeError("Expected argument 'primary_key' to be a str")
        pulumi.set(__self__, "primary_key", primary_key)
        if secondary_key and not isinstance(secondary_key, str):
            raise TypeError("Expected argument 'secondary_key' to be a str")
        pulumi.set(__self__, "secondary_key", secondary_key)

    @property
    @pulumi.getter(name="primaryKey")
    def primary_key(self) -> str:
        """
        The primary key for accessing the Maps REST APIs.
        """
        return pulumi.get(self, "primary_key")

    @property
    @pulumi.getter(name="secondaryKey")
    def secondary_key(self) -> str:
        """
        The secondary key for accessing the Maps REST APIs.
        """
        return pulumi.get(self, "secondary_key")


class AwaitableListAccountKeysResult(ListAccountKeysResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListAccountKeysResult(
            primary_key=self.primary_key,
            secondary_key=self.secondary_key)


def list_account_keys(account_name: Optional[str] = None,
                      resource_group_name: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListAccountKeysResult:
    """
    Use this data source to access information about an existing resource.

    :param str account_name: The name of the Maps Account.
    :param str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['accountName'] = account_name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:maps/latest:listAccountKeys', __args__, opts=opts, typ=ListAccountKeysResult).value

    return AwaitableListAccountKeysResult(
        primary_key=__ret__.primary_key,
        secondary_key=__ret__.secondary_key)
