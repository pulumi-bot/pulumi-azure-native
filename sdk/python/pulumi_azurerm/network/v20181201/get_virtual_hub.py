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
    'GetVirtualHubResult',
    'AwaitableGetVirtualHubResult',
    'get_virtual_hub',
]

@pulumi.output_type
class GetVirtualHubResult:
    """
    VirtualHub Resource.
    """
    def __init__(__self__, address_prefix=None, etag=None, express_route_gateway=None, location=None, name=None, p2_s_vpn_gateway=None, provisioning_state=None, route_table=None, tags=None, type=None, virtual_network_connections=None, virtual_wan=None, vpn_gateway=None):
        if address_prefix and not isinstance(address_prefix, str):
            raise TypeError("Expected argument 'address_prefix' to be a str")
        pulumi.set(__self__, "address_prefix", address_prefix)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if express_route_gateway and not isinstance(express_route_gateway, dict):
            raise TypeError("Expected argument 'express_route_gateway' to be a dict")
        pulumi.set(__self__, "express_route_gateway", express_route_gateway)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if p2_s_vpn_gateway and not isinstance(p2_s_vpn_gateway, dict):
            raise TypeError("Expected argument 'p2_s_vpn_gateway' to be a dict")
        pulumi.set(__self__, "p2_s_vpn_gateway", p2_s_vpn_gateway)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if route_table and not isinstance(route_table, dict):
            raise TypeError("Expected argument 'route_table' to be a dict")
        pulumi.set(__self__, "route_table", route_table)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if virtual_network_connections and not isinstance(virtual_network_connections, list):
            raise TypeError("Expected argument 'virtual_network_connections' to be a list")
        pulumi.set(__self__, "virtual_network_connections", virtual_network_connections)
        if virtual_wan and not isinstance(virtual_wan, dict):
            raise TypeError("Expected argument 'virtual_wan' to be a dict")
        pulumi.set(__self__, "virtual_wan", virtual_wan)
        if vpn_gateway and not isinstance(vpn_gateway, dict):
            raise TypeError("Expected argument 'vpn_gateway' to be a dict")
        pulumi.set(__self__, "vpn_gateway", vpn_gateway)

    @property
    @pulumi.getter(name="addressPrefix")
    def address_prefix(self) -> Optional[str]:
        """
        Address-prefix for this VirtualHub.
        """
        return pulumi.get(self, "address_prefix")

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        Gets a unique read-only string that changes whenever the resource is updated.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="expressRouteGateway")
    def express_route_gateway(self) -> Optional['outputs.SubResourceResponse']:
        """
        The expressRouteGateway associated with this VirtualHub
        """
        return pulumi.get(self, "express_route_gateway")

    @property
    @pulumi.getter
    def location(self) -> str:
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
    @pulumi.getter(name="p2SVpnGateway")
    def p2_s_vpn_gateway(self) -> Optional['outputs.SubResourceResponse']:
        """
        The P2SVpnGateway associated with this VirtualHub
        """
        return pulumi.get(self, "p2_s_vpn_gateway")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> Optional[str]:
        """
        The provisioning state of the resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="routeTable")
    def route_table(self) -> Optional['outputs.VirtualHubRouteTableResponse']:
        """
        The routeTable associated with this virtual hub.
        """
        return pulumi.get(self, "route_table")

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
    @pulumi.getter(name="virtualNetworkConnections")
    def virtual_network_connections(self) -> Optional[List['outputs.HubVirtualNetworkConnectionResponse']]:
        """
        List of all vnet connections with this VirtualHub.
        """
        return pulumi.get(self, "virtual_network_connections")

    @property
    @pulumi.getter(name="virtualWan")
    def virtual_wan(self) -> Optional['outputs.SubResourceResponse']:
        """
        The VirtualWAN to which the VirtualHub belongs
        """
        return pulumi.get(self, "virtual_wan")

    @property
    @pulumi.getter(name="vpnGateway")
    def vpn_gateway(self) -> Optional['outputs.SubResourceResponse']:
        """
        The VpnGateway associated with this VirtualHub
        """
        return pulumi.get(self, "vpn_gateway")


class AwaitableGetVirtualHubResult(GetVirtualHubResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVirtualHubResult(
            address_prefix=self.address_prefix,
            etag=self.etag,
            express_route_gateway=self.express_route_gateway,
            location=self.location,
            name=self.name,
            p2_s_vpn_gateway=self.p2_s_vpn_gateway,
            provisioning_state=self.provisioning_state,
            route_table=self.route_table,
            tags=self.tags,
            type=self.type,
            virtual_network_connections=self.virtual_network_connections,
            virtual_wan=self.virtual_wan,
            vpn_gateway=self.vpn_gateway)


def get_virtual_hub(name: Optional[str] = None,
                    resource_group_name: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVirtualHubResult:
    """
    Use this data source to access information about an existing resource.

    :param str name: The name of the VirtualHub.
    :param str resource_group_name: The resource group name of the VirtualHub.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:network/v20181201:getVirtualHub', __args__, opts=opts, typ=GetVirtualHubResult).value

    return AwaitableGetVirtualHubResult(
        address_prefix=__ret__.address_prefix,
        etag=__ret__.etag,
        express_route_gateway=__ret__.express_route_gateway,
        location=__ret__.location,
        name=__ret__.name,
        p2_s_vpn_gateway=__ret__.p2_s_vpn_gateway,
        provisioning_state=__ret__.provisioning_state,
        route_table=__ret__.route_table,
        tags=__ret__.tags,
        type=__ret__.type,
        virtual_network_connections=__ret__.virtual_network_connections,
        virtual_wan=__ret__.virtual_wan,
        vpn_gateway=__ret__.vpn_gateway)
