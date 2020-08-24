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
    'GetSubnetResult',
    'AwaitableGetSubnetResult',
    'get_subnet',
]

@pulumi.output_type
class GetSubnetResult:
    """
    Subnet in a VirtualNetwork resource
    """
    def __init__(__self__, address_prefix=None, etag=None, ip_configurations=None, name=None, network_security_group=None, provisioning_state=None, resource_navigation_links=None, route_table=None):
        if address_prefix and not isinstance(address_prefix, str):
            raise TypeError("Expected argument 'address_prefix' to be a str")
        pulumi.set(__self__, "address_prefix", address_prefix)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if ip_configurations and not isinstance(ip_configurations, list):
            raise TypeError("Expected argument 'ip_configurations' to be a list")
        pulumi.set(__self__, "ip_configurations", ip_configurations)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if network_security_group and not isinstance(network_security_group, dict):
            raise TypeError("Expected argument 'network_security_group' to be a dict")
        pulumi.set(__self__, "network_security_group", network_security_group)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if resource_navigation_links and not isinstance(resource_navigation_links, list):
            raise TypeError("Expected argument 'resource_navigation_links' to be a list")
        pulumi.set(__self__, "resource_navigation_links", resource_navigation_links)
        if route_table and not isinstance(route_table, dict):
            raise TypeError("Expected argument 'route_table' to be a dict")
        pulumi.set(__self__, "route_table", route_table)

    @property
    @pulumi.getter(name="addressPrefix")
    def address_prefix(self) -> Optional[str]:
        """
        Gets or sets Address prefix for the subnet.
        """
        return pulumi.get(self, "address_prefix")

    @property
    @pulumi.getter
    def etag(self) -> Optional[str]:
        """
        A unique read-only string that changes whenever the resource is updated
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="ipConfigurations")
    def ip_configurations(self) -> List['outputs.IPConfigurationResponse']:
        """
        Gets array of references to the network interface IP configurations using subnet
        """
        return pulumi.get(self, "ip_configurations")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Gets or sets the name of the resource that is unique within a resource group. This name can be used to access the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkSecurityGroup")
    def network_security_group(self) -> Optional['outputs.NetworkSecurityGroupResponse']:
        """
        Gets or sets the reference of the NetworkSecurityGroup resource
        """
        return pulumi.get(self, "network_security_group")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> Optional[str]:
        """
        Gets provisioning state of the resource
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="resourceNavigationLinks")
    def resource_navigation_links(self) -> Optional[List['outputs.ResourceNavigationLinkResponse']]:
        """
        Gets array of references to the external resources using subnet
        """
        return pulumi.get(self, "resource_navigation_links")

    @property
    @pulumi.getter(name="routeTable")
    def route_table(self) -> Optional['outputs.RouteTableResponse']:
        """
        Gets or sets the reference of the RouteTable resource
        """
        return pulumi.get(self, "route_table")


class AwaitableGetSubnetResult(GetSubnetResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSubnetResult(
            address_prefix=self.address_prefix,
            etag=self.etag,
            ip_configurations=self.ip_configurations,
            name=self.name,
            network_security_group=self.network_security_group,
            provisioning_state=self.provisioning_state,
            resource_navigation_links=self.resource_navigation_links,
            route_table=self.route_table)


def get_subnet(expand: Optional[str] = None,
               name: Optional[str] = None,
               resource_group_name: Optional[str] = None,
               virtual_network_name: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSubnetResult:
    """
    Use this data source to access information about an existing resource.

    :param str expand: expand references resources.
    :param str name: The name of the subnet.
    :param str resource_group_name: The name of the resource group.
    :param str virtual_network_name: The name of the virtual network.
    """
    __args__ = dict()
    __args__['expand'] = expand
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    __args__['virtualNetworkName'] = virtual_network_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:network/v20160601:getSubnet', __args__, opts=opts, typ=GetSubnetResult).value

    return AwaitableGetSubnetResult(
        address_prefix=__ret__.address_prefix,
        etag=__ret__.etag,
        ip_configurations=__ret__.ip_configurations,
        name=__ret__.name,
        network_security_group=__ret__.network_security_group,
        provisioning_state=__ret__.provisioning_state,
        resource_navigation_links=__ret__.resource_navigation_links,
        route_table=__ret__.route_table)
