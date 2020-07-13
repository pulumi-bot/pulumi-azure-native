# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class ManagerDeviceVolumeContainer(pulumi.CustomResource):
    kind: pulumi.Output[str]
    """
    The Kind of the object. Currently only Series8000 is supported
    """
    name: pulumi.Output[str]
    """
    The name of the object.
    """
    properties: pulumi.Output[dict]
    """
    The volume container properties.
      * `band_width_rate_in_mbps` (`float`) - The bandwidth-rate set on the volume container.
      * `bandwidth_setting_id` (`str`) - The ID of the bandwidth setting associated with the volume container.
      * `encryption_key` (`dict`) - The key used to encrypt data in the volume container. It is required when property 'EncryptionStatus' is "Enabled".
        * `encryption_algorithm` (`str`) - The algorithm used to encrypt "Value".
        * `encryption_cert_thumbprint` (`str`) - Thumbprint certificate that was used to encrypt "Value". If the value in unencrypted, it will be null.
        * `value` (`str`) - The value of the secret.

      * `encryption_status` (`str`) - The flag to denote whether encryption is enabled or not.
      * `owner_ship_status` (`str`) - The owner ship status of the volume container. Only when the status is "NotOwned", the delete operation on the volume container is permitted.
      * `storage_account_credential_id` (`str`) - The path ID of storage account associated with the volume container.
      * `total_cloud_storage_usage_in_bytes` (`float`) - The total cloud storage for the volume container.
      * `volume_count` (`float`) - The number of volumes in the volume Container.
    """
    type: pulumi.Output[str]
    """
    The hierarchical type of the object.
    """
    def __init__(__self__, resource_name, opts=None, device_name=None, kind=None, manager_name=None, name=None, properties=None, resource_group_name=None, __props__=None, __name__=None, __opts__=None):
        """
        The volume container.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] device_name: The device name
        :param pulumi.Input[str] kind: The Kind of the object. Currently only Series8000 is supported
        :param pulumi.Input[str] manager_name: The manager name
        :param pulumi.Input[str] name: The name of the volume container.
        :param pulumi.Input[dict] properties: The volume container properties.
        :param pulumi.Input[str] resource_group_name: The resource group name

        The **properties** object supports the following:

          * `band_width_rate_in_mbps` (`pulumi.Input[float]`) - The bandwidth-rate set on the volume container.
          * `bandwidth_setting_id` (`pulumi.Input[str]`) - The ID of the bandwidth setting associated with the volume container.
          * `encryption_key` (`pulumi.Input[dict]`) - The key used to encrypt data in the volume container. It is required when property 'EncryptionStatus' is "Enabled".
            * `encryption_algorithm` (`pulumi.Input[str]`) - The algorithm used to encrypt "Value".
            * `encryption_cert_thumbprint` (`pulumi.Input[str]`) - Thumbprint certificate that was used to encrypt "Value". If the value in unencrypted, it will be null.
            * `value` (`pulumi.Input[str]`) - The value of the secret.

          * `storage_account_credential_id` (`pulumi.Input[str]`) - The path ID of storage account associated with the volume container.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if device_name is None:
                raise TypeError("Missing required property 'device_name'")
            __props__['device_name'] = device_name
            __props__['kind'] = kind
            if manager_name is None:
                raise TypeError("Missing required property 'manager_name'")
            __props__['manager_name'] = manager_name
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if properties is None:
                raise TypeError("Missing required property 'properties'")
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['type'] = None
        super(ManagerDeviceVolumeContainer, __self__).__init__(
            'azurerm:storsimple:ManagerDeviceVolumeContainer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing ManagerDeviceVolumeContainer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ManagerDeviceVolumeContainer(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
