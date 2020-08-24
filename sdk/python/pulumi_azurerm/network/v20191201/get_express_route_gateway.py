# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs

__all__ = [
    'GetExpressRouteGatewayResult',
    'AwaitableGetExpressRouteGatewayResult',
    'get_express_route_gateway',
]

@pulumi.output_type
class GetExpressRouteGatewayResult:
    """
    ExpressRoute gateway resource.
    """
    def __init__(__self__, auto_scale_configuration=None, etag=None, express_route_connections=None, location=None, name=None, provisioning_state=None, tags=None, type=None, virtual_hub=None):
        if auto_scale_configuration and not isinstance(auto_scale_configuration, dict):
            raise TypeError("Expected argument 'auto_scale_configuration' to be a dict")
        pulumi.set(__self__, "auto_scale_configuration", auto_scale_configuration)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if express_route_connections and not isinstance(express_route_connections, list):
            raise TypeError("Expected argument 'express_route_connections' to be a list")
        pulumi.set(__self__, "express_route_connections", express_route_connections)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if virtual_hub and not isinstance(virtual_hub, dict):
            raise TypeError("Expected argument 'virtual_hub' to be a dict")
        pulumi.set(__self__, "virtual_hub", virtual_hub)

    @property
    @pulumi.getter(name="autoScaleConfiguration")
    def auto_scale_configuration(self) -> Optional['outputs.ExpressRouteGatewayPropertiesResponseAutoScaleConfiguration']:
        """
        Configuration for auto scaling.
        """
        return pulumi.get(self, "auto_scale_configuration")

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        A unique read-only string that changes whenever the resource is updated.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="expressRouteConnections")
    def express_route_connections(self) -> List['outputs.ExpressRouteConnectionResponse']:
        """
        List of ExpressRoute connections to the ExpressRoute gateway.
        """
        return pulumi.get(self, "express_route_connections")

    @property
    @pulumi.getter
    def location(self) -> Optional[str]:
        """
        Resource location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        The provisioning state of the express route gateway resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="virtualHub")
    def virtual_hub(self) -> 'outputs.VirtualHubIdResponse':
        """
        The Virtual Hub where the ExpressRoute gateway is or will be deployed.
        """
        return pulumi.get(self, "virtual_hub")


class AwaitableGetExpressRouteGatewayResult(GetExpressRouteGatewayResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetExpressRouteGatewayResult(
            auto_scale_configuration=self.auto_scale_configuration,
            etag=self.etag,
            express_route_connections=self.express_route_connections,
            location=self.location,
            name=self.name,
            provisioning_state=self.provisioning_state,
            tags=self.tags,
            type=self.type,
            virtual_hub=self.virtual_hub)


def get_express_route_gateway(name: Optional[str] = None,
                              resource_group_name: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetExpressRouteGatewayResult:
    """
    Use this data source to access information about an existing resource.

    :param str name: The name of the ExpressRoute gateway.
    :param str resource_group_name: The name of the resource group.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:network/v20191201:getExpressRouteGateway', __args__, opts=opts, typ=GetExpressRouteGatewayResult).value

    return AwaitableGetExpressRouteGatewayResult(
        auto_scale_configuration=__ret__.auto_scale_configuration,
        etag=__ret__.etag,
        express_route_connections=__ret__.express_route_connections,
        location=__ret__.location,
        name=__ret__.name,
        provisioning_state=__ret__.provisioning_state,
        tags=__ret__.tags,
        type=__ret__.type,
        virtual_hub=__ret__.virtual_hub)
