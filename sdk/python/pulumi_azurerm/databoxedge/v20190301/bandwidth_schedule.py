# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class BandwidthSchedule(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    The object name.
    """
    properties: pulumi.Output[dict]
    """
    The properties of the bandwidth schedule.
      * `days` (`list`) - The days of the week when this schedule is applicable.
      * `rate_in_mbps` (`float`) - The bandwidth rate in Mbps.
      * `start` (`str`) - The start time of the schedule in UTC.
      * `stop` (`str`) - The stop time of the schedule in UTC.
    """
    type: pulumi.Output[str]
    """
    The hierarchical type of the object.
    """
    def __init__(__self__, resource_name, opts=None, device_name=None, name=None, properties=None, resource_group_name=None, __props__=None, __name__=None, __opts__=None):
        """
        The bandwidth schedule details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] device_name: The device name.
        :param pulumi.Input[str] name: The bandwidth schedule name which needs to be added/updated.
        :param pulumi.Input[dict] properties: The properties of the bandwidth schedule.
        :param pulumi.Input[str] resource_group_name: The resource group name.

        The **properties** object supports the following:

          * `days` (`pulumi.Input[list]`) - The days of the week when this schedule is applicable.
          * `rate_in_mbps` (`pulumi.Input[float]`) - The bandwidth rate in Mbps.
          * `start` (`pulumi.Input[str]`) - The start time of the schedule in UTC.
          * `stop` (`pulumi.Input[str]`) - The stop time of the schedule in UTC.
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
        super(BandwidthSchedule, __self__).__init__(
            'azurerm:databoxedge/v20190301:BandwidthSchedule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing BandwidthSchedule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return BandwidthSchedule(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
