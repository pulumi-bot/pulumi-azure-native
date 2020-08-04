# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class GetExpressRouteConnectionResult:
    """
    ExpressRouteConnection resource.
    """
    def __init__(__self__, name=None, properties=None):
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        The name of the resource.
        """
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        __self__.properties = properties
        """
        Properties of the ExpressRouteConnection subresource.
        """


class AwaitableGetExpressRouteConnectionResult(GetExpressRouteConnectionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetExpressRouteConnectionResult(
            name=self.name,
            properties=self.properties)


def get_express_route_connection(express_route_gateway_name=None, name=None, resource_group_name=None, opts=None):
    """
    Use this data source to access information about an existing resource.

    :param str express_route_gateway_name: The name of the ExpressRoute gateway.
    :param str name: The name of the ExpressRoute connection.
    :param str resource_group_name: The name of the resource group.
    """
    __args__ = dict()
    __args__['expressRouteGatewayName'] = express_route_gateway_name
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:network/v20180801:getExpressRouteConnection', __args__, opts=opts).value

    return AwaitableGetExpressRouteConnectionResult(
        name=__ret__.get('name'),
        properties=__ret__.get('properties'))
