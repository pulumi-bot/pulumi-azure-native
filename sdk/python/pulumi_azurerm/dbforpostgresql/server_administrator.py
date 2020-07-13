# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class ServerAdministrator(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    The name of the resource
    """
    properties: pulumi.Output[dict]
    """
    Properties of the server AAD administrator.
      * `administrator_type` (`str`) - The type of administrator.
      * `login` (`str`) - The server administrator login account name.
      * `sid` (`str`) - The server administrator Sid (Secure ID).
      * `tenant_id` (`str`) - The server Active Directory Administrator tenant id.
    """
    type: pulumi.Output[str]
    """
    The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
    """
    def __init__(__self__, resource_name, opts=None, name=None, properties=None, resource_group_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Represents a and external administrator to be created.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the server.
        :param pulumi.Input[dict] properties: Properties of the server AAD administrator.
        :param pulumi.Input[str] resource_group_name: The name of the resource group. The name is case insensitive.

        The **properties** object supports the following:

          * `administrator_type` (`pulumi.Input[str]`) - The type of administrator.
          * `login` (`pulumi.Input[str]`) - The server administrator login account name.
          * `sid` (`pulumi.Input[str]`) - The server administrator Sid (Secure ID).
          * `tenant_id` (`pulumi.Input[str]`) - The server Active Directory Administrator tenant id.
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

            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['type'] = None
        super(ServerAdministrator, __self__).__init__(
            'azurerm:dbforpostgresql:ServerAdministrator',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing ServerAdministrator resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ServerAdministrator(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
