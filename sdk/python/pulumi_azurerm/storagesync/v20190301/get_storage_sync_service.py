# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'GetStorageSyncServiceResult',
    'AwaitableGetStorageSyncServiceResult',
    'get_storage_sync_service',
]

@pulumi.output_type
class GetStorageSyncServiceResult:
    """
    Storage Sync Service object.
    """
    def __init__(__self__, location=None, name=None, storage_sync_service_status=None, storage_sync_service_uid=None, tags=None, type=None):
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if storage_sync_service_status and not isinstance(storage_sync_service_status, float):
            raise TypeError("Expected argument 'storage_sync_service_status' to be a float")
        pulumi.set(__self__, "storage_sync_service_status", storage_sync_service_status)
        if storage_sync_service_uid and not isinstance(storage_sync_service_uid, str):
            raise TypeError("Expected argument 'storage_sync_service_uid' to be a str")
        pulumi.set(__self__, "storage_sync_service_uid", storage_sync_service_uid)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="storageSyncServiceStatus")
    def storage_sync_service_status(self) -> float:
        """
        Storage Sync service status.
        """
        return pulumi.get(self, "storage_sync_service_status")

    @property
    @pulumi.getter(name="storageSyncServiceUid")
    def storage_sync_service_uid(self) -> str:
        """
        Storage Sync service Uid
        """
        return pulumi.get(self, "storage_sync_service_uid")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        """
        return pulumi.get(self, "type")


class AwaitableGetStorageSyncServiceResult(GetStorageSyncServiceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetStorageSyncServiceResult(
            location=self.location,
            name=self.name,
            storage_sync_service_status=self.storage_sync_service_status,
            storage_sync_service_uid=self.storage_sync_service_uid,
            tags=self.tags,
            type=self.type)


def get_storage_sync_service(name: Optional[str] = None,
                             resource_group_name: Optional[str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetStorageSyncServiceResult:
    """
    Use this data source to access information about an existing resource.

    :param str name: Name of Storage Sync Service resource.
    :param str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:storagesync/v20190301:getStorageSyncService', __args__, opts=opts, typ=GetStorageSyncServiceResult).value

    return AwaitableGetStorageSyncServiceResult(
        location=__ret__.location,
        name=__ret__.name,
        storage_sync_service_status=__ret__.storage_sync_service_status,
        storage_sync_service_uid=__ret__.storage_sync_service_uid,
        tags=__ret__.tags,
        type=__ret__.type)
