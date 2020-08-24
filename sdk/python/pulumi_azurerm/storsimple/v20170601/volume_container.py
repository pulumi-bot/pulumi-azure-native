# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['VolumeContainer']


class VolumeContainer(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 band_width_rate_in_mbps: Optional[pulumi.Input[float]] = None,
                 bandwidth_setting_id: Optional[pulumi.Input[str]] = None,
                 device_name: Optional[pulumi.Input[str]] = None,
                 encryption_key: Optional[pulumi.Input[pulumi.InputType['AsymmetricEncryptedSecretArgs']]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 manager_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 storage_account_credential_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The volume container.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] band_width_rate_in_mbps: The bandwidth-rate set on the volume container.
        :param pulumi.Input[str] bandwidth_setting_id: The ID of the bandwidth setting associated with the volume container.
        :param pulumi.Input[str] device_name: The device name
        :param pulumi.Input[pulumi.InputType['AsymmetricEncryptedSecretArgs']] encryption_key: The key used to encrypt data in the volume container. It is required when property 'EncryptionStatus' is "Enabled".
        :param pulumi.Input[str] kind: The Kind of the object. Currently only Series8000 is supported
        :param pulumi.Input[str] manager_name: The manager name
        :param pulumi.Input[str] name: The name of the volume container.
        :param pulumi.Input[str] resource_group_name: The resource group name
        :param pulumi.Input[str] storage_account_credential_id: The path ID of storage account associated with the volume container.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['band_width_rate_in_mbps'] = band_width_rate_in_mbps
            __props__['bandwidth_setting_id'] = bandwidth_setting_id
            if device_name is None:
                raise TypeError("Missing required property 'device_name'")
            __props__['device_name'] = device_name
            __props__['encryption_key'] = encryption_key
            __props__['kind'] = kind
            if manager_name is None:
                raise TypeError("Missing required property 'manager_name'")
            __props__['manager_name'] = manager_name
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if storage_account_credential_id is None:
                raise TypeError("Missing required property 'storage_account_credential_id'")
            __props__['storage_account_credential_id'] = storage_account_credential_id
            __props__['encryption_status'] = None
            __props__['owner_ship_status'] = None
            __props__['total_cloud_storage_usage_in_bytes'] = None
            __props__['type'] = None
            __props__['volume_count'] = None
        super(VolumeContainer, __self__).__init__(
            'azurerm:storsimple/v20170601:VolumeContainer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'VolumeContainer':
        """
        Get an existing VolumeContainer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return VolumeContainer(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="bandWidthRateInMbps")
    def band_width_rate_in_mbps(self) -> Optional[float]:
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
    def volume_count(self) -> float:
        """
        The number of volumes in the volume Container.
        """
        return pulumi.get(self, "volume_count")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

