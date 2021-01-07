# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs

__all__ = [
    'GetVpnGatewayResult',
    'AwaitableGetVpnGatewayResult',
    'get_vpn_gateway',
]

@pulumi.output_type
class GetVpnGatewayResult:
    """
    VpnGateway Resource.
    """
    def __init__(__self__, bgp_settings=None, connections=None, etag=None, id=None, ip_configurations=None, is_routing_preference_internet=None, location=None, name=None, nat_rules=None, provisioning_state=None, tags=None, type=None, virtual_hub=None, vpn_gateway_scale_unit=None):
        if bgp_settings and not isinstance(bgp_settings, dict):
            raise TypeError("Expected argument 'bgp_settings' to be a dict")
        pulumi.set(__self__, "bgp_settings", bgp_settings)
        if connections and not isinstance(connections, list):
            raise TypeError("Expected argument 'connections' to be a list")
        pulumi.set(__self__, "connections", connections)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip_configurations and not isinstance(ip_configurations, list):
            raise TypeError("Expected argument 'ip_configurations' to be a list")
        pulumi.set(__self__, "ip_configurations", ip_configurations)
        if is_routing_preference_internet and not isinstance(is_routing_preference_internet, bool):
            raise TypeError("Expected argument 'is_routing_preference_internet' to be a bool")
        pulumi.set(__self__, "is_routing_preference_internet", is_routing_preference_internet)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if nat_rules and not isinstance(nat_rules, list):
            raise TypeError("Expected argument 'nat_rules' to be a list")
        pulumi.set(__self__, "nat_rules", nat_rules)
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
        if vpn_gateway_scale_unit and not isinstance(vpn_gateway_scale_unit, int):
            raise TypeError("Expected argument 'vpn_gateway_scale_unit' to be a int")
        pulumi.set(__self__, "vpn_gateway_scale_unit", vpn_gateway_scale_unit)

    @property
    @pulumi.getter(name="bgpSettings")
    def bgp_settings(self) -> Optional['outputs.BgpSettingsResponse']:
        """
        Local network gateway's BGP speaker settings.
        """
        return pulumi.get(self, "bgp_settings")

    @property
    @pulumi.getter
    def connections(self) -> Optional[Sequence['outputs.VpnConnectionResponse']]:
        """
        List of all vpn connections to the gateway.
        """
        return pulumi.get(self, "connections")

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        A unique read-only string that changes whenever the resource is updated.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        Resource ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ipConfigurations")
    def ip_configurations(self) -> Sequence['outputs.VpnGatewayIpConfigurationResponse']:
        """
        List of all IPs configured on the gateway.
        """
        return pulumi.get(self, "ip_configurations")

    @property
    @pulumi.getter(name="isRoutingPreferenceInternet")
    def is_routing_preference_internet(self) -> Optional[bool]:
        """
        Enable Routing Preference property for the Public IP Interface of the VpnGateway.
        """
        return pulumi.get(self, "is_routing_preference_internet")

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
    @pulumi.getter(name="natRules")
    def nat_rules(self) -> Optional[Sequence['outputs.VpnGatewayNatRuleResponse']]:
        """
        List of all the nat Rules associated with the gateway.
        """
        return pulumi.get(self, "nat_rules")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        The provisioning state of the VPN gateway resource.
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
    def virtual_hub(self) -> Optional['outputs.SubResourceResponse']:
        """
        The VirtualHub to which the gateway belongs.
        """
        return pulumi.get(self, "virtual_hub")

    @property
    @pulumi.getter(name="vpnGatewayScaleUnit")
    def vpn_gateway_scale_unit(self) -> Optional[int]:
        """
        The scale unit for this vpn gateway.
        """
        return pulumi.get(self, "vpn_gateway_scale_unit")


class AwaitableGetVpnGatewayResult(GetVpnGatewayResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVpnGatewayResult(
            bgp_settings=self.bgp_settings,
            connections=self.connections,
            etag=self.etag,
            id=self.id,
            ip_configurations=self.ip_configurations,
            is_routing_preference_internet=self.is_routing_preference_internet,
            location=self.location,
            name=self.name,
            nat_rules=self.nat_rules,
            provisioning_state=self.provisioning_state,
            tags=self.tags,
            type=self.type,
            virtual_hub=self.virtual_hub,
            vpn_gateway_scale_unit=self.vpn_gateway_scale_unit)


def get_vpn_gateway(gateway_name: Optional[str] = None,
                    resource_group_name: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVpnGatewayResult:
    """
    Use this data source to access information about an existing resource.

    :param str gateway_name: The name of the gateway.
    :param str resource_group_name: The resource group name of the VpnGateway.
    """
    __args__ = dict()
    __args__['gatewayName'] = gateway_name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure-nextgen:network/latest:getVpnGateway', __args__, opts=opts, typ=GetVpnGatewayResult).value

    return AwaitableGetVpnGatewayResult(
        bgp_settings=__ret__.bgp_settings,
        connections=__ret__.connections,
        etag=__ret__.etag,
        id=__ret__.id,
        ip_configurations=__ret__.ip_configurations,
        is_routing_preference_internet=__ret__.is_routing_preference_internet,
        location=__ret__.location,
        name=__ret__.name,
        nat_rules=__ret__.nat_rules,
        provisioning_state=__ret__.provisioning_state,
        tags=__ret__.tags,
        type=__ret__.type,
        virtual_hub=__ret__.virtual_hub,
        vpn_gateway_scale_unit=__ret__.vpn_gateway_scale_unit)
