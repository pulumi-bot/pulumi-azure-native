# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables

__all__ = [
    'GetLinkedStorageAccountResult',
    'AwaitableGetLinkedStorageAccountResult',
    'get_linked_storage_account',
]

@pulumi.output_type
class GetLinkedStorageAccountResult:
    """
    Linked storage accounts top level resource container.
    """
    def __init__(__self__, data_source_type=None, name=None, storage_account_ids=None, type=None):
        if data_source_type and not isinstance(data_source_type, str):
            raise TypeError("Expected argument 'data_source_type' to be a str")
        pulumi.set(__self__, "data_source_type", data_source_type)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if storage_account_ids and not isinstance(storage_account_ids, list):
            raise TypeError("Expected argument 'storage_account_ids' to be a list")
        pulumi.set(__self__, "storage_account_ids", storage_account_ids)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="dataSourceType")
    def data_source_type(self) -> str:
        """
        Linked storage accounts type.
        """
        return pulumi.get(self, "data_source_type")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="storageAccountIds")
    def storage_account_ids(self) -> Optional[Sequence[str]]:
        """
        Linked storage accounts resources ids.
        """
        return pulumi.get(self, "storage_account_ids")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")


class AwaitableGetLinkedStorageAccountResult(GetLinkedStorageAccountResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLinkedStorageAccountResult(
            data_source_type=self.data_source_type,
            name=self.name,
            storage_account_ids=self.storage_account_ids,
            type=self.type)


def get_linked_storage_account(data_source_type: Optional[str] = None,
                               resource_group_name: Optional[str] = None,
                               workspace_name: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLinkedStorageAccountResult:
    """
    Use this data source to access information about an existing resource.

    :param str data_source_type: Linked storage accounts type.
    :param str resource_group_name: The name of the resource group. The name is case insensitive.
    :param str workspace_name: The name of the workspace.
    """
    __args__ = dict()
    __args__['dataSourceType'] = data_source_type
    __args__['resourceGroupName'] = resource_group_name
    __args__['workspaceName'] = workspace_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure-nextgen:operationalinsights/v20200801:getLinkedStorageAccount', __args__, opts=opts, typ=GetLinkedStorageAccountResult).value

    return AwaitableGetLinkedStorageAccountResult(
        data_source_type=__ret__.data_source_type,
        name=__ret__.name,
        storage_account_ids=__ret__.storage_account_ids,
        type=__ret__.type)
