# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class Trigger(pulumi.CustomResource):
    kind: pulumi.Output[str]
    """
    Trigger Kind.
    """
    name: pulumi.Output[str]
    """
    The object name.
    """
    type: pulumi.Output[str]
    """
    The hierarchical type of the object.
    """
    def __init__(__self__, resource_name, opts=None, device_name=None, kind=None, name=None, resource_group_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Trigger details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] device_name: Creates or updates a trigger
        :param pulumi.Input[str] kind: Trigger Kind.
        :param pulumi.Input[str] name: The trigger name.
        :param pulumi.Input[str] resource_group_name: The resource group name.
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

            if device_name is None:
                raise TypeError("Missing required property 'device_name'")
            __props__['device_name'] = device_name
            if kind is None:
                raise TypeError("Missing required property 'kind'")
            __props__['kind'] = kind
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['type'] = None
        super(Trigger, __self__).__init__(
            'azurerm:databoxedge/v20190301:Trigger',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Trigger resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Trigger(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
