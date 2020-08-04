# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class Account(pulumi.CustomResource):
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
    NetApp Account properties
      * `active_directories` (`list`) - Active Directories
        * `active_directory_id` (`str`) - Id of the Active Directory
        * `dns` (`str`) - Comma separated list of DNS server IP addresses for the Active Directory domain
        * `domain` (`str`) - Name of the Active Directory domain
        * `organizational_unit` (`str`) - The Organizational Unit (OU) within the Windows Active Directory
        * `password` (`str`) - Plain text password of Active Directory domain administrator
        * `smb_server_name` (`str`) - NetBIOS name of the SMB server. This name will be registered as a computer account in the AD and used to mount volumes
        * `status` (`str`) - Status of the Active Directory
        * `username` (`str`) - Username of Active Directory domain administrator

      * `provisioning_state` (`str`) - Azure lifecycle management
    """
    tags: pulumi.Output[dict]
    """
    Resource tags
    """
    type: pulumi.Output[str]
    """
    Resource type
    """
    def __init__(__self__, resource_name, opts=None, active_directories=None, location=None, name=None, resource_group_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        NetApp account resource

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] active_directories: Active Directories
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[str] name: The name of the NetApp account
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[dict] tags: Resource tags

        The **active_directories** object supports the following:

          * `active_directory_id` (`pulumi.Input[str]`) - Id of the Active Directory
          * `dns` (`pulumi.Input[str]`) - Comma separated list of DNS server IP addresses for the Active Directory domain
          * `domain` (`pulumi.Input[str]`) - Name of the Active Directory domain
          * `organizational_unit` (`pulumi.Input[str]`) - The Organizational Unit (OU) within the Windows Active Directory
          * `password` (`pulumi.Input[str]`) - Plain text password of Active Directory domain administrator
          * `smb_server_name` (`pulumi.Input[str]`) - NetBIOS name of the SMB server. This name will be registered as a computer account in the AD and used to mount volumes
          * `status` (`pulumi.Input[str]`) - Status of the Active Directory
          * `username` (`pulumi.Input[str]`) - Username of Active Directory domain administrator
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

            __props__['active_directories'] = active_directories
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['properties'] = None
            __props__['type'] = None
        super(Account, __self__).__init__(
            'azurerm:netapp/v20190701:Account',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Account resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Account(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop