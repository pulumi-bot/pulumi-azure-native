# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class GetResourceGroupResult:
    """
    Resource group information.
    """
    def __init__(__self__, location=None, name=None, properties=None, tags=None):
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        __self__.location = location
        """
        The location of the resource group. It cannot be changed after the resource group has been created. Has to be one of the supported Azure Locations, such as West US, East US, West Europe, East Asia, etc.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        The Name of the resource group.
        """
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        __self__.properties = properties
        """
        The resource group properties.
        """
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        """
        The tags attached to the resource group.
        """


class AwaitableGetResourceGroupResult(GetResourceGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetResourceGroupResult(
            location=self.location,
            name=self.name,
            properties=self.properties,
            tags=self.tags)


def get_resource_group(name=None, opts=None):
    """
    Use this data source to access information about an existing resource.

    :param str name: The name of the resource group to get. The name is case insensitive.
    """
    __args__ = dict()
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:resources/v20160701:getResourceGroup', __args__, opts=opts).value

    return AwaitableGetResourceGroupResult(
        location=__ret__.get('location'),
        name=__ret__.get('name'),
        properties=__ret__.get('properties'),
        tags=__ret__.get('tags'))
