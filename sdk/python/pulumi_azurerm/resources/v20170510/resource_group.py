# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class ResourceGroup(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    The location of the resource group. It cannot be changed after the resource group has been created. It must be one of the supported Azure locations.
    """
    managed_by: pulumi.Output[str]
    """
    The ID of the resource that manages this resource group.
    """
    name: pulumi.Output[str]
    """
    The name of the resource group.
    """
    properties: pulumi.Output[dict]
    """
    The resource group properties.
      * `provisioning_state` (`str`) - The provisioning state. 
    """
    tags: pulumi.Output[dict]
    """
    The tags attached to the resource group.
    """
    def __init__(__self__, resource_name, opts=None, location=None, managed_by=None, name=None, properties=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Resource group information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: The location of the resource group. It cannot be changed after the resource group has been created. It must be one of the supported Azure locations.
        :param pulumi.Input[str] managed_by: The ID of the resource that manages this resource group.
        :param pulumi.Input[str] name: The name of the resource group to create or update.
        :param pulumi.Input[dict] properties: The resource group properties.
        :param pulumi.Input[dict] tags: The tags attached to the resource group.
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

            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            __props__['managed_by'] = managed_by
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['properties'] = properties
            __props__['tags'] = tags
        super(ResourceGroup, __self__).__init__(
            'azurerm:resources/v20170510:ResourceGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing ResourceGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ResourceGroup(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
