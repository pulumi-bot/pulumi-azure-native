# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class GetBandwidthScheduleResult:
    """
    The bandwidth schedule details.
    """
    def __init__(__self__, name=None, properties=None, type=None):
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        The object name.
        """
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        __self__.properties = properties
        """
        The properties of the bandwidth schedule.
        """
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        __self__.type = type
        """
        The hierarchical type of the object.
        """


class AwaitableGetBandwidthScheduleResult(GetBandwidthScheduleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBandwidthScheduleResult(
            name=self.name,
            properties=self.properties,
            type=self.type)


def get_bandwidth_schedule(device_name=None, name=None, resource_group_name=None, opts=None):
    """
    Use this data source to access information about an existing resource.

    :param str device_name: The device name.
    :param str name: The bandwidth schedule name.
    :param str resource_group_name: The resource group name.
    """
    __args__ = dict()
    __args__['deviceName'] = device_name
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:databoxedge/v20190701:getBandwidthSchedule', __args__, opts=opts).value

    return AwaitableGetBandwidthScheduleResult(
        name=__ret__.get('name'),
        properties=__ret__.get('properties'),
        type=__ret__.get('type'))
