# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'ListDatabaseAccountKeysResult',
    'AwaitableListDatabaseAccountKeysResult',
    'list_database_account_keys',
]

@pulumi.output_type
class ListDatabaseAccountKeysResult:
    """
    The access keys for the given database account.
    """
    def __init__(__self__, primary_master_key=None, primary_readonly_master_key=None, secondary_master_key=None, secondary_readonly_master_key=None):
        if primary_master_key and not isinstance(primary_master_key, str):
            raise TypeError("Expected argument 'primary_master_key' to be a str")
        pulumi.set(__self__, "primary_master_key", primary_master_key)
        if primary_readonly_master_key and not isinstance(primary_readonly_master_key, str):
            raise TypeError("Expected argument 'primary_readonly_master_key' to be a str")
        pulumi.set(__self__, "primary_readonly_master_key", primary_readonly_master_key)
        if secondary_master_key and not isinstance(secondary_master_key, str):
            raise TypeError("Expected argument 'secondary_master_key' to be a str")
        pulumi.set(__self__, "secondary_master_key", secondary_master_key)
        if secondary_readonly_master_key and not isinstance(secondary_readonly_master_key, str):
            raise TypeError("Expected argument 'secondary_readonly_master_key' to be a str")
        pulumi.set(__self__, "secondary_readonly_master_key", secondary_readonly_master_key)

    @property
    @pulumi.getter(name="primaryMasterKey")
    def primary_master_key(self) -> str:
        """
        Base 64 encoded value of the primary read-write key.
        """
        return pulumi.get(self, "primary_master_key")

    @property
    @pulumi.getter(name="primaryReadonlyMasterKey")
    def primary_readonly_master_key(self) -> str:
        """
        Base 64 encoded value of the primary read-only key.
        """
        return pulumi.get(self, "primary_readonly_master_key")

    @property
    @pulumi.getter(name="secondaryMasterKey")
    def secondary_master_key(self) -> str:
        """
        Base 64 encoded value of the secondary read-write key.
        """
        return pulumi.get(self, "secondary_master_key")

    @property
    @pulumi.getter(name="secondaryReadonlyMasterKey")
    def secondary_readonly_master_key(self) -> str:
        """
        Base 64 encoded value of the secondary read-only key.
        """
        return pulumi.get(self, "secondary_readonly_master_key")


class AwaitableListDatabaseAccountKeysResult(ListDatabaseAccountKeysResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListDatabaseAccountKeysResult(
            primary_master_key=self.primary_master_key,
            primary_readonly_master_key=self.primary_readonly_master_key,
            secondary_master_key=self.secondary_master_key,
            secondary_readonly_master_key=self.secondary_readonly_master_key)


def list_database_account_keys(account_name: Optional[str] = None,
                               resource_group_name: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListDatabaseAccountKeysResult:
    """
    Use this data source to access information about an existing resource.

    :param str account_name: Cosmos DB database account name.
    :param str resource_group_name: Name of an Azure resource group.
    """
    __args__ = dict()
    __args__['accountName'] = account_name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:documentdb/v20160331:listDatabaseAccountKeys', __args__, opts=opts, typ=ListDatabaseAccountKeysResult).value

    return AwaitableListDatabaseAccountKeysResult(
        primary_master_key=__ret__.primary_master_key,
        primary_readonly_master_key=__ret__.primary_readonly_master_key,
        secondary_master_key=__ret__.secondary_master_key,
        secondary_readonly_master_key=__ret__.secondary_readonly_master_key)
