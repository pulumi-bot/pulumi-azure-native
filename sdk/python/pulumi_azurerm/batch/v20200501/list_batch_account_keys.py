# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class ListBatchAccountKeysResult:
    """
    A set of Azure Batch account keys.
    """
    def __init__(__self__, account_name=None, primary=None, secondary=None):
        if account_name and not isinstance(account_name, str):
            raise TypeError("Expected argument 'account_name' to be a str")
        __self__.account_name = account_name
        """
        The Batch account name.
        """
        if primary and not isinstance(primary, str):
            raise TypeError("Expected argument 'primary' to be a str")
        __self__.primary = primary
        """
        The primary key associated with the account.
        """
        if secondary and not isinstance(secondary, str):
            raise TypeError("Expected argument 'secondary' to be a str")
        __self__.secondary = secondary
        """
        The secondary key associated with the account.
        """


class AwaitableListBatchAccountKeysResult(ListBatchAccountKeysResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListBatchAccountKeysResult(
            account_name=self.account_name,
            primary=self.primary,
            secondary=self.secondary)


def list_batch_account_keys(account_name=None, resource_group_name=None, opts=None):
    """
    Use this data source to access information about an existing resource.

    :param str account_name: The name of the Batch account.
    :param str resource_group_name: The name of the resource group that contains the Batch account.
    """
    __args__ = dict()
    __args__['accountName'] = account_name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:batch/v20200501:listBatchAccountKeys', __args__, opts=opts).value

    return AwaitableListBatchAccountKeysResult(
        account_name=__ret__.get('accountName'),
        primary=__ret__.get('primary'),
        secondary=__ret__.get('secondary'))
