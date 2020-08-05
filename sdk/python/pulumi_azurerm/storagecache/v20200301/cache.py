# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class Cache(pulumi.CustomResource):
    identity: pulumi.Output[dict]
    """
    The identity of the cache, if configured.
      * `principal_id` (`str`) - The principal id of the cache.
      * `tenant_id` (`str`) - The tenant id associated with the cache.
      * `type` (`str`) - The type of identity used for the cache
    """
    location: pulumi.Output[str]
    """
    Region name string.
    """
    name: pulumi.Output[str]
    """
    Name of Cache.
    """
    properties: pulumi.Output[dict]
    """
    Properties of the Cache.
      * `cache_size_gb` (`float`) - The size of this Cache, in GB.
      * `encryption_settings` (`dict`) - Specifies encryption settings of the cache.
        * `key_encryption_key` (`dict`) - Specifies the location of the key encryption key in Key Vault.
          * `key_url` (`str`) - The URL referencing a key encryption key in Key Vault.
          * `source_vault` (`dict`) - Describes a resource Id to source Key Vault.
            * `id` (`str`) - Resource Id.

      * `health` (`dict`) - Health of the Cache.
        * `state` (`str`) - List of Cache health states.
        * `status_description` (`str`) - Describes explanation of state.

      * `mount_addresses` (`list`) - Array of IP addresses that can be used by clients mounting this Cache.
      * `network_settings` (`dict`) - Specifies network settings of the cache.
        * `mtu` (`float`) - The IPv4 maximum transmission unit configured for the subnet.
        * `utility_addresses` (`list`) - Array of additional IP addresses used by this Cache.

      * `provisioning_state` (`str`) - ARM provisioning state, see https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/Addendum.md#provisioningstate-property
      * `security_settings` (`dict`) - Specifies security settings of the cache.
        * `root_squash` (`bool`) - root squash of cache property.

      * `subnet` (`str`) - Subnet used for the Cache.
      * `upgrade_status` (`dict`) - Upgrade status of the Cache.
        * `current_firmware_version` (`str`) - Version string of the firmware currently installed on this Cache.
        * `firmware_update_deadline` (`str`) - Time at which the pending firmware update will automatically be installed on the Cache.
        * `firmware_update_status` (`str`) - True if there is a firmware update ready to install on this Cache. The firmware will automatically be installed after firmwareUpdateDeadline if not triggered earlier via the upgrade operation.
        * `last_firmware_update` (`str`) - Time of the last successful firmware update.
        * `pending_firmware_version` (`str`) - When firmwareUpdateAvailable is true, this field holds the version string for the update.
    """
    sku: pulumi.Output[dict]
    """
    SKU for the Cache.
      * `cache_size_gb` (`float`) - The size of this Cache, in GB.
      * `encryption_settings` (`dict`) - Specifies encryption settings of the cache.
        * `key_encryption_key` (`dict`) - Specifies the location of the key encryption key in Key Vault.
          * `key_url` (`str`) - The URL referencing a key encryption key in Key Vault.
          * `source_vault` (`dict`) - Describes a resource Id to source Key Vault.
            * `id` (`str`) - Resource Id.

      * `health` (`dict`) - Health of the Cache.
        * `state` (`str`) - List of Cache health states.
        * `status_description` (`str`) - Describes explanation of state.

      * `mount_addresses` (`list`) - Array of IP addresses that can be used by clients mounting this Cache.
      * `network_settings` (`dict`) - Specifies network settings of the cache.
        * `mtu` (`float`) - The IPv4 maximum transmission unit configured for the subnet.
        * `utility_addresses` (`list`) - Array of additional IP addresses used by this Cache.

      * `provisioning_state` (`str`) - ARM provisioning state, see https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/Addendum.md#provisioningstate-property
      * `security_settings` (`dict`) - Specifies security settings of the cache.
        * `root_squash` (`bool`) - root squash of cache property.

      * `subnet` (`str`) - Subnet used for the Cache.
      * `upgrade_status` (`dict`) - Upgrade status of the Cache.
        * `current_firmware_version` (`str`) - Version string of the firmware currently installed on this Cache.
        * `firmware_update_deadline` (`str`) - Time at which the pending firmware update will automatically be installed on the Cache.
        * `firmware_update_status` (`str`) - True if there is a firmware update ready to install on this Cache. The firmware will automatically be installed after firmwareUpdateDeadline if not triggered earlier via the upgrade operation.
        * `last_firmware_update` (`str`) - Time of the last successful firmware update.
        * `pending_firmware_version` (`str`) - When firmwareUpdateAvailable is true, this field holds the version string for the update.
    """
    tags: pulumi.Output[dict]
    """
    ARM tags as name/value pairs.
    """
    type: pulumi.Output[str]
    """
    Type of the Cache; Microsoft.StorageCache/Cache
    """
    def __init__(__self__, resource_name, opts=None, cache_size_gb=None, encryption_settings=None, identity=None, location=None, name=None, network_settings=None, provisioning_state=None, resource_group_name=None, security_settings=None, sku=None, subnet=None, tags=None, upgrade_status=None, __props__=None, __name__=None, __opts__=None):
        """
        A Cache instance. Follows Azure Resource Manager standards: https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/resource-api-reference.md

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] cache_size_gb: The size of this Cache, in GB.
        :param pulumi.Input[dict] encryption_settings: Specifies encryption settings of the cache.
        :param pulumi.Input[dict] identity: The identity of the cache, if configured.
        :param pulumi.Input[str] location: Region name string.
        :param pulumi.Input[str] name: Name of Cache. Length of name must be not greater than 80 and chars must be in list of [-0-9a-zA-Z_] char class.
        :param pulumi.Input[dict] network_settings: Specifies network settings of the cache.
        :param pulumi.Input[str] provisioning_state: ARM provisioning state, see https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/Addendum.md#provisioningstate-property
        :param pulumi.Input[str] resource_group_name: Target resource group.
        :param pulumi.Input[dict] security_settings: Specifies security settings of the cache.
        :param pulumi.Input[dict] sku: SKU for the Cache.
        :param pulumi.Input[str] subnet: Subnet used for the Cache.
        :param pulumi.Input[dict] tags: ARM tags as name/value pairs.
        :param pulumi.Input[dict] upgrade_status: Upgrade status of the Cache.

        The **encryption_settings** object supports the following:

          * `key_encryption_key` (`pulumi.Input[dict]`) - Specifies the location of the key encryption key in Key Vault.
            * `key_url` (`pulumi.Input[str]`) - The URL referencing a key encryption key in Key Vault.
            * `source_vault` (`pulumi.Input[dict]`) - Describes a resource Id to source Key Vault.
              * `id` (`pulumi.Input[str]`) - Resource Id.

        The **identity** object supports the following:

          * `type` (`pulumi.Input[str]`) - The type of identity used for the cache

        The **network_settings** object supports the following:

          * `mtu` (`pulumi.Input[float]`) - The IPv4 maximum transmission unit configured for the subnet.

        The **security_settings** object supports the following:

          * `root_squash` (`pulumi.Input[bool]`) - root squash of cache property.

        The **sku** object supports the following:

          * `name` (`pulumi.Input[str]`) - SKU name for this Cache.
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

            __props__['cache_size_gb'] = cache_size_gb
            __props__['encryption_settings'] = encryption_settings
            __props__['identity'] = identity
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['network_settings'] = network_settings
            __props__['provisioning_state'] = provisioning_state
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['security_settings'] = security_settings
            __props__['sku'] = sku
            __props__['subnet'] = subnet
            __props__['tags'] = tags
            __props__['upgrade_status'] = upgrade_status
            __props__['properties'] = None
            __props__['type'] = None
        super(Cache, __self__).__init__(
            'azurerm:storagecache/v20200301:Cache',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Cache resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Cache(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
