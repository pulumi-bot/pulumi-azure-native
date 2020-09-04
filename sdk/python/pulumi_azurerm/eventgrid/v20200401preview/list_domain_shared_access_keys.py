# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'ListDomainSharedAccessKeysResult',
    'AwaitableListDomainSharedAccessKeysResult',
    'list_domain_shared_access_keys',
]

@pulumi.output_type
class ListDomainSharedAccessKeysResult:
    """
    Shared access keys of the Domain.
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
        Shared access key1 for the domain.
        """
        return pulumi.get(self, "key1")

    @property
    @pulumi.getter
    def key2(self) -> Optional[str]:
        """
        Shared access key2 for the domain.
        """
        return pulumi.get(self, "key2")


class AwaitableListDomainSharedAccessKeysResult(ListDomainSharedAccessKeysResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListDomainSharedAccessKeysResult(
            key1=self.key1,
            key2=self.key2)


def list_domain_shared_access_keys(domain_name: Optional[str] = None,
                                   resource_group_name: Optional[str] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListDomainSharedAccessKeysResult:
    """
    Use this data source to access information about an existing resource.

    :param str domain_name: Name of the domain.
    :param str resource_group_name: The name of the resource group within the user's subscription.
    """
    __args__ = dict()
    __args__['domainName'] = domain_name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:eventgrid/v20200401preview:listDomainSharedAccessKeys', __args__, opts=opts, typ=ListDomainSharedAccessKeysResult).value

    return AwaitableListDomainSharedAccessKeysResult(
        key1=__ret__.key1,
        key2=__ret__.key2)
