# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class Volume(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    Resource location
    """
    name: pulumi.Output[str]
    """
    Resource name
    """
    properties: pulumi.Output[dict]
    """
    Volume properties
      * `baremetal_tenant_id` (`str`) - Unique Baremetal Tenant Identifier.
      * `creation_token` (`str`) - A unique file path for the volume. Used when creating mount targets
      * `export_policy` (`dict`) - Set of export policy rules
        * `rules` (`list`) - Export policy rule
          * `allowed_clients` (`str`) - Client ingress specification as comma separated string with IPv4 CIDRs, IPv4 host addresses and host names
          * `cifs` (`bool`) - Allows CIFS protocol
          * `nfsv3` (`bool`) - Allows NFSv3 protocol
          * `nfsv4` (`bool`) - Deprecated: Will use the NFSv4.1 protocol, please use swagger version 2019-07-01 or later
          * `rule_index` (`float`) - Order index
          * `unix_read_only` (`bool`) - Read only access
          * `unix_read_write` (`bool`) - Read and write access

      * `file_system_id` (`str`) - Unique FileSystem Identifier.
      * `mount_targets` (`list`) - List of mount targets
        * `end_ip` (`str`) - The end of IPv4 address range to use when creating a new mount target
        * `file_system_id` (`str`) - UUID v4 used to identify the MountTarget
        * `gateway` (`str`) - The gateway of the IPv4 address range to use when creating a new mount target
        * `ip_address` (`str`) - The mount target's IPv4 address
        * `mount_target_id` (`str`) - UUID v4 used to identify the MountTarget
        * `netmask` (`str`) - The netmask of the IPv4 address range to use when creating a new mount target
        * `provisioning_state` (`str`) - Azure lifecycle management
        * `smb_server_fqdn` (`str`) - The SMB server's Fully Qualified Domain Name, FQDN
        * `start_ip` (`str`) - The start of IPv4 address range to use when creating a new mount target
        * `subnet` (`str`) - The subnet

      * `protocol_types` (`list`) - Set of protocol types
      * `provisioning_state` (`str`) - Azure lifecycle management
      * `service_level` (`str`) - The service level of the file system
      * `snapshot_id` (`str`) - UUID v4 or resource identifier used to identify the Snapshot.
      * `subnet_id` (`str`) - The Azure Resource URI for a delegated subnet. Must have the delegation Microsoft.NetApp/volumes
      * `usage_threshold` (`float`) - Maximum storage quota allowed for a file system in bytes. This is a soft quota used for alerting only. Minimum size is 100 GiB. Upper limit is 100TiB. Specified in bytes.
    """
    tags: pulumi.Output[dict]
    """
    Resource tags
    """
    type: pulumi.Output[str]
    """
    Resource type
    """
    def __init__(__self__, resource_name, opts=None, account_name=None, location=None, name=None, pool_name=None, properties=None, resource_group_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Volume resource

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: The name of the NetApp account
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[str] name: The name of the volume
        :param pulumi.Input[str] pool_name: The name of the capacity pool
        :param pulumi.Input[dict] properties: Volume properties
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[dict] tags: Resource tags

        The **properties** object supports the following:

          * `creation_token` (`pulumi.Input[str]`) - A unique file path for the volume. Used when creating mount targets
          * `export_policy` (`pulumi.Input[dict]`) - Set of export policy rules
            * `rules` (`pulumi.Input[list]`) - Export policy rule
              * `allowed_clients` (`pulumi.Input[str]`) - Client ingress specification as comma separated string with IPv4 CIDRs, IPv4 host addresses and host names
              * `cifs` (`pulumi.Input[bool]`) - Allows CIFS protocol
              * `nfsv3` (`pulumi.Input[bool]`) - Allows NFSv3 protocol
              * `nfsv4` (`pulumi.Input[bool]`) - Deprecated: Will use the NFSv4.1 protocol, please use swagger version 2019-07-01 or later
              * `rule_index` (`pulumi.Input[float]`) - Order index
              * `unix_read_only` (`pulumi.Input[bool]`) - Read only access
              * `unix_read_write` (`pulumi.Input[bool]`) - Read and write access

          * `mount_targets` (`pulumi.Input[list]`) - List of mount targets
            * `end_ip` (`pulumi.Input[str]`) - The end of IPv4 address range to use when creating a new mount target
            * `file_system_id` (`pulumi.Input[str]`) - UUID v4 used to identify the MountTarget
            * `gateway` (`pulumi.Input[str]`) - The gateway of the IPv4 address range to use when creating a new mount target
            * `netmask` (`pulumi.Input[str]`) - The netmask of the IPv4 address range to use when creating a new mount target
            * `smb_server_fqdn` (`pulumi.Input[str]`) - The SMB server's Fully Qualified Domain Name, FQDN
            * `start_ip` (`pulumi.Input[str]`) - The start of IPv4 address range to use when creating a new mount target
            * `subnet` (`pulumi.Input[str]`) - The subnet

          * `protocol_types` (`pulumi.Input[list]`) - Set of protocol types
          * `service_level` (`pulumi.Input[str]`) - The service level of the file system
          * `snapshot_id` (`pulumi.Input[str]`) - UUID v4 or resource identifier used to identify the Snapshot.
          * `subnet_id` (`pulumi.Input[str]`) - The Azure Resource URI for a delegated subnet. Must have the delegation Microsoft.NetApp/volumes
          * `usage_threshold` (`pulumi.Input[float]`) - Maximum storage quota allowed for a file system in bytes. This is a soft quota used for alerting only. Minimum size is 100 GiB. Upper limit is 100TiB. Specified in bytes.
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

            if account_name is None:
                raise TypeError("Missing required property 'account_name'")
            __props__['account_name'] = account_name
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if pool_name is None:
                raise TypeError("Missing required property 'pool_name'")
            __props__['pool_name'] = pool_name
            if properties is None:
                raise TypeError("Missing required property 'properties'")
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['type'] = None
        super(Volume, __self__).__init__(
            'azurerm:netapp/v20190601:Volume',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Volume resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Volume(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
