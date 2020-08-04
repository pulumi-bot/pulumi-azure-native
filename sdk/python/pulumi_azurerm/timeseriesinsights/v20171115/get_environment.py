# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class GetEnvironmentResult:
    """
    An environment is a set of time-series data available for query, and is the top level Azure Time Series Insights resource.
    """
    def __init__(__self__, location=None, name=None, properties=None, sku=None, tags=None, type=None):
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        __self__.location = location
        """
        Resource location
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        Resource name
        """
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        __self__.properties = properties
        """
        Properties of the environment.
        """
        if sku and not isinstance(sku, dict):
            raise TypeError("Expected argument 'sku' to be a dict")
        __self__.sku = sku
        """
        The sku determines the capacity of the environment, the SLA (in queries-per-minute and total capacity), and the billing rate.
        """
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        """
        Resource tags
        """
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        __self__.type = type
        """
        Resource type
        """


class AwaitableGetEnvironmentResult(GetEnvironmentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEnvironmentResult(
            location=self.location,
            name=self.name,
            properties=self.properties,
            sku=self.sku,
            tags=self.tags,
            type=self.type)


def get_environment(name=None, resource_group_name=None, opts=None):
    """
    Use this data source to access information about an existing resource.

    :param str name: The name of the Time Series Insights environment associated with the specified resource group.
    :param str resource_group_name: Name of an Azure Resource group.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:timeseriesinsights/v20171115:getEnvironment', __args__, opts=opts).value

    return AwaitableGetEnvironmentResult(
        location=__ret__.get('location'),
        name=__ret__.get('name'),
        properties=__ret__.get('properties'),
        sku=__ret__.get('sku'),
        tags=__ret__.get('tags'),
        type=__ret__.get('type'))
