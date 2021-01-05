# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs

__all__ = [
    'GetVolumeContainerResult',
    'AwaitableGetVolumeContainerResult',
    'get_volume_container',
]

@pulumi.output_type
class GetVolumeContainerResult:
    """
    The volume container.
    """
    def __init__(__self__, band_width_rate_in_mbps=None, bandwidth_setting_id=None, encryption_key=None, encryption_status=None, id=None, kind=None, name=None, owner_ship_status=None, storage_account_credential_id=None, total_cloud_storage_usage_in_bytes=None, type=None, volume_count=None):
        if band_width_rate_in_mbps and not isinstance(band_width_rate_in_mbps, int):
            raise TypeError("Expected argument 'band_width_rate_in_mbps' to be a int")
        pulumi.set(__self__, "band_width_rate_in_mbps", band_width_rate_in_mbps)
        if bandwidth_setting_id and not isinstance(bandwidth_setting_id, str):
            raise TypeError("Expected argument 'bandwidth_setting_id' to be a str")
        pulumi.set(__self__, "bandwidth_setting_id", bandwidth_setting_id)
        if encryption_key and not isinstance(encryption_key, dict):
            raise TypeError("Expected argument 'encryption_key' to be a dict")
        pulumi.set(__self__, "encryption_key", encryption_key)
        if encryption_status and not isinstance(encryption_status, str):
            raise TypeError("Expected argument 'encryption_status' to be a str")
        pulumi.set(__self__, "encryption_status", encryption_status)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if owner_ship_status and not isinstance(owner_ship_status, str):
            raise TypeError("Expected argument 'owner_ship_status' to be a str")
        pulumi.set(__self__, "owner_ship_status", owner_ship_status)
        if storage_account_credential_id and not isinstance(storage_account_credential_id, str):
            raise TypeError("Expected argument 'storage_account_credential_id' to be a str")
        pulumi.set(__self__, "storage_account_credential_id", storage_account_credential_id)
        if total_cloud_storage_usage_in_bytes and not isinstance(total_cloud_storage_usage_in_bytes, float):
            raise TypeError("Expected argument 'total_cloud_storage_usage_in_bytes' to be a float")
        pulumi.set(__self__, "total_cloud_storage_usage_in_bytes", total_cloud_storage_usage_in_bytes)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if volume_count and not isinstance(volume_count, int):
            raise TypeError("Expected argument 'volume_count' to be a int")
        pulumi.set(__self__, "volume_count", volume_count)

    @property
    @pulumi.getter(name="bandWidthRateInMbps")
    def band_width_rate_in_mbps(self) -> Optional[int]:
        """
        The bandwidth-rate set on the volume container.
        """
        return pulumi.get(self, "band_width_rate_in_mbps")

    @property
    @pulumi.getter(name="bandwidthSettingId")
    def bandwidth_setting_id(self) -> Optional[str]:
        """
        The ID of the bandwidth setting associated with the volume container.
        """
        return pulumi.get(self, "bandwidth_setting_id")

    @property
    @pulumi.getter(name="encryptionKey")
    def encryption_key(self) -> Optional['outputs.AsymmetricEncryptedSecretResponse']:
        """
        The key used to encrypt data in the volume container. It is required when property 'EncryptionStatus' is "Enabled".
        """
        return pulumi.get(self, "encryption_key")

    @property
    @pulumi.getter(name="encryptionStatus")
    def encryption_status(self) -> str:
        """
        The flag to denote whether encryption is enabled or not.
        """
        return pulumi.get(self, "encryption_status")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The path ID that uniquely identifies the object.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def kind(self) -> Optional[str]:
        """
        The Kind of the object. Currently only Series8000 is supported
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the object.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="ownerShipStatus")
    def owner_ship_status(self) -> str:
        """
        The owner ship status of the volume container. Only when the status is "NotOwned", the delete operation on the volume container is permitted.
        """
        return pulumi.get(self, "owner_ship_status")

    @property
    @pulumi.getter(name="storageAccountCredentialId")
    def storage_account_credential_id(self) -> str:
        """
        The path ID of storage account associated with the volume container.
        """
        return pulumi.get(self, "storage_account_credential_id")

    @property
    @pulumi.getter(name="totalCloudStorageUsageInBytes")
    def total_cloud_storage_usage_in_bytes(self) -> float:
        """
        The total cloud storage for the volume container.
        """
        return pulumi.get(self, "total_cloud_storage_usage_in_bytes")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The hierarchical type of the object.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="volumeCount")
    def volume_count(self) -> int:
        """
        The number of volumes in the volume Container.
        """
        return pulumi.get(self, "volume_count")


class AwaitableGetVolumeContainerResult(GetVolumeContainerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVolumeContainerResult(
            band_width_rate_in_mbps=self.band_width_rate_in_mbps,
            bandwidth_setting_id=self.bandwidth_setting_id,
            encryption_key=self.encryption_key,
            encryption_status=self.encryption_status,
            id=self.id,
            kind=self.kind,
            name=self.name,
            owner_ship_status=self.owner_ship_status,
            storage_account_credential_id=self.storage_account_credential_id,
            total_cloud_storage_usage_in_bytes=self.total_cloud_storage_usage_in_bytes,
            type=self.type,
            volume_count=self.volume_count)


def get_volume_container(device_name: Optional[str] = None,
                         manager_name: Optional[str] = None,
                         resource_group_name: Optional[str] = None,
                         volume_container_name: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVolumeContainerResult:
    """
    Use this data source to access information about an existing resource.

    :param str device_name: The device name
    :param str manager_name: The manager name
    :param str resource_group_name: The resource group name
    :param str volume_container_name: The name of the volume container.
    """
    __args__ = dict()
    __args__['deviceName'] = device_name
    __args__['managerName'] = manager_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['volumeContainerName'] = volume_container_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure-nextgen:storsimple/v20170601:getVolumeContainer', __args__, opts=opts, typ=GetVolumeContainerResult).value

    return AwaitableGetVolumeContainerResult(
        band_width_rate_in_mbps=__ret__.band_width_rate_in_mbps,
        bandwidth_setting_id=__ret__.bandwidth_setting_id,
        encryption_key=__ret__.encryption_key,
        encryption_status=__ret__.encryption_status,
        id=__ret__.id,
        kind=__ret__.kind,
        name=__ret__.name,
        owner_ship_status=__ret__.owner_ship_status,
        storage_account_credential_id=__ret__.storage_account_credential_id,
        total_cloud_storage_usage_in_bytes=__ret__.total_cloud_storage_usage_in_bytes,
        type=__ret__.type,
        volume_count=__ret__.volume_count)
