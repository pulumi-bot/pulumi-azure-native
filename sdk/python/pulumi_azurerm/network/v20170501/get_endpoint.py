# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class GetEndpointResult:
    """
    Class representing a Traffic Manager endpoint.
    """
    def __init__(__self__, name=None, properties=None, type=None):
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        The name of the resource
        """
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        __self__.properties = properties
        """
        The properties of the Traffic Manager endpoint.
        """
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        __self__.type = type
        """
        The type of the resource. Ex- Microsoft.Network/trafficManagerProfiles.
        """


class AwaitableGetEndpointResult(GetEndpointResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEndpointResult(
            name=self.name,
            properties=self.properties,
            type=self.type)


def get_endpoint(endpoint_type=None, name=None, profile_name=None, resource_group_name=None, opts=None):
    """
    Use this data source to access information about an existing resource.

    :param str endpoint_type: The type of the Traffic Manager endpoint.
    :param str name: The name of the Traffic Manager endpoint.
    :param str profile_name: The name of the Traffic Manager profile.
    :param str resource_group_name: The name of the resource group containing the Traffic Manager endpoint.
    """
    __args__ = dict()
    __args__['endpointType'] = endpoint_type
    __args__['name'] = name
    __args__['profileName'] = profile_name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:network/v20170501:getEndpoint', __args__, opts=opts).value

    return AwaitableGetEndpointResult(
        name=__ret__.get('name'),
        properties=__ret__.get('properties'),
        type=__ret__.get('type'))