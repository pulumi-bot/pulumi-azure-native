# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class GetOriginResult:
    """
    CDN origin is the source of the content being delivered via CDN. When the edge nodes represented by an endpoint do not have the requested content cached, they attempt to fetch it from one or more of the configured origins.
    """
    def __init__(__self__, name=None, properties=None, type=None):
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        Resource name
        """
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        __self__.properties = properties
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        __self__.type = type
        """
        Resource type
        """


class AwaitableGetOriginResult(GetOriginResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOriginResult(
            name=self.name,
            properties=self.properties,
            type=self.type)


def get_origin(endpoint_name=None, name=None, profile_name=None, resource_group_name=None, opts=None):
    """
    Use this data source to access information about an existing resource.

    :param str endpoint_name: Name of the endpoint within the CDN profile.
    :param str name: Name of the origin, an arbitrary value but it needs to be unique under endpoint
    :param str profile_name: Name of the CDN profile within the resource group.
    :param str resource_group_name: Name of the resource group within the Azure subscription.
    """
    __args__ = dict()
    __args__['endpointName'] = endpoint_name
    __args__['name'] = name
    __args__['profileName'] = profile_name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:cdn/v20160402:getOrigin', __args__, opts=opts).value

    return AwaitableGetOriginResult(
        name=__ret__.get('name'),
        properties=__ret__.get('properties'),
        type=__ret__.get('type'))
