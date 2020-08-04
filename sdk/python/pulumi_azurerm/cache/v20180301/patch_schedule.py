# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class PatchSchedule(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    Resource name.
    """
    properties: pulumi.Output[dict]
    """
    List of patch schedules for a Redis cache.
      * `schedule_entries` (`list`) - List of patch schedules for a Redis cache.
        * `day_of_week` (`str`) - Day of the week when a cache can be patched.
        * `maintenance_window` (`str`) - ISO8601 timespan specifying how much time cache patching can take. 
        * `start_hour_utc` (`float`) - Start hour after which cache patching can start.
    """
    type: pulumi.Output[str]
    """
    Resource type.
    """
    def __init__(__self__, resource_name, opts=None, name=None, properties=None, resource_group_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Response to put/get patch schedules for Redis cache.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Default string modeled as parameter for auto generation to work correctly.
        :param pulumi.Input[dict] properties: List of patch schedules for a Redis cache.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.

        The **properties** object supports the following:

          * `schedule_entries` (`pulumi.Input[list]`) - List of patch schedules for a Redis cache.
            * `day_of_week` (`pulumi.Input[str]`) - Day of the week when a cache can be patched.
            * `maintenance_window` (`pulumi.Input[str]`) - ISO8601 timespan specifying how much time cache patching can take. 
            * `start_hour_utc` (`pulumi.Input[float]`) - Start hour after which cache patching can start.
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
        super(PatchSchedule, __self__).__init__(
            'azurerm:cache/v20180301:PatchSchedule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing PatchSchedule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return PatchSchedule(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
