# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'GetSyncGroupResult',
    'AwaitableGetSyncGroupResult',
    'get_sync_group',
]

@pulumi.output_type
class GetSyncGroupResult:
    """
    Sync Group object.
    """
    def __init__(__self__, name=None, sync_group_status=None, type=None, unique_id=None):
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if sync_group_status and not isinstance(sync_group_status, str):
            raise TypeError("Expected argument 'sync_group_status' to be a str")
        pulumi.set(__self__, "sync_group_status", sync_group_status)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if unique_id and not isinstance(unique_id, str):
            raise TypeError("Expected argument 'unique_id' to be a str")
        pulumi.set(__self__, "unique_id", unique_id)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="syncGroupStatus")
    def sync_group_status(self) -> str:
        """
        Sync group status
        """
        return pulumi.get(self, "sync_group_status")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="uniqueId")
    def unique_id(self) -> Optional[str]:
        """
        Unique Id
        """
        return pulumi.get(self, "unique_id")


class AwaitableGetSyncGroupResult(GetSyncGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSyncGroupResult(
            name=self.name,
            sync_group_status=self.sync_group_status,
            type=self.type,
            unique_id=self.unique_id)


def get_sync_group(name: Optional[str] = None,
                   resource_group_name: Optional[str] = None,
                   storage_sync_service_name: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSyncGroupResult:
    """
    Use this data source to access information about an existing resource.

    :param str name: Name of Sync Group resource.
    :param str resource_group_name: The name of the resource group. The name is case insensitive.
    :param str storage_sync_service_name: Name of Storage Sync Service resource.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    __args__['storageSyncServiceName'] = storage_sync_service_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:storagesync/v20180402:getSyncGroup', __args__, opts=opts, typ=GetSyncGroupResult).value

    return AwaitableGetSyncGroupResult(
        name=__ret__.name,
        sync_group_status=__ret__.sync_group_status,
        type=__ret__.type,
        unique_id=__ret__.unique_id)
