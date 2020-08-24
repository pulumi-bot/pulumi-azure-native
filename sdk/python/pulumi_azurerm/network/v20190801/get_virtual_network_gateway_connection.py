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
    'GetVirtualNetworkGatewayConnectionResult',
    'AwaitableGetVirtualNetworkGatewayConnectionResult',
    'get_virtual_network_gateway_connection',
]

@pulumi.output_type
class GetVirtualNetworkGatewayConnectionResult:
    """
    A common class for general resource information.
    """
    def __init__(__self__, authorization_key=None, connection_protocol=None, connection_status=None, connection_type=None, egress_bytes_transferred=None, enable_bgp=None, etag=None, express_route_gateway_bypass=None, ingress_bytes_transferred=None, ipsec_policies=None, local_network_gateway2=None, location=None, name=None, peer=None, provisioning_state=None, resource_guid=None, routing_weight=None, shared_key=None, tags=None, traffic_selector_policies=None, tunnel_connection_status=None, type=None, use_policy_based_traffic_selectors=None, virtual_network_gateway1=None, virtual_network_gateway2=None):
        if authorization_key and not isinstance(authorization_key, str):
            raise TypeError("Expected argument 'authorization_key' to be a str")
        pulumi.set(__self__, "authorization_key", authorization_key)
        if connection_protocol and not isinstance(connection_protocol, str):
            raise TypeError("Expected argument 'connection_protocol' to be a str")
        pulumi.set(__self__, "connection_protocol", connection_protocol)
        if connection_status and not isinstance(connection_status, str):
            raise TypeError("Expected argument 'connection_status' to be a str")
        pulumi.set(__self__, "connection_status", connection_status)
        if connection_type and not isinstance(connection_type, str):
            raise TypeError("Expected argument 'connection_type' to be a str")
        pulumi.set(__self__, "connection_type", connection_type)
        if egress_bytes_transferred and not isinstance(egress_bytes_transferred, float):
            raise TypeError("Expected argument 'egress_bytes_transferred' to be a float")
        pulumi.set(__self__, "egress_bytes_transferred", egress_bytes_transferred)
        if enable_bgp and not isinstance(enable_bgp, bool):
            raise TypeError("Expected argument 'enable_bgp' to be a bool")
        pulumi.set(__self__, "enable_bgp", enable_bgp)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if express_route_gateway_bypass and not isinstance(express_route_gateway_bypass, bool):
            raise TypeError("Expected argument 'express_route_gateway_bypass' to be a bool")
        pulumi.set(__self__, "express_route_gateway_bypass", express_route_gateway_bypass)
        if ingress_bytes_transferred and not isinstance(ingress_bytes_transferred, float):
            raise TypeError("Expected argument 'ingress_bytes_transferred' to be a float")
        pulumi.set(__self__, "ingress_bytes_transferred", ingress_bytes_transferred)
        if ipsec_policies and not isinstance(ipsec_policies, list):
            raise TypeError("Expected argument 'ipsec_policies' to be a list")
        pulumi.set(__self__, "ipsec_policies", ipsec_policies)
        if local_network_gateway2 and not isinstance(local_network_gateway2, dict):
            raise TypeError("Expected argument 'local_network_gateway2' to be a dict")
        pulumi.set(__self__, "local_network_gateway2", local_network_gateway2)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if peer and not isinstance(peer, dict):
            raise TypeError("Expected argument 'peer' to be a dict")
        pulumi.set(__self__, "peer", peer)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if resource_guid and not isinstance(resource_guid, str):
            raise TypeError("Expected argument 'resource_guid' to be a str")
        pulumi.set(__self__, "resource_guid", resource_guid)
        if routing_weight and not isinstance(routing_weight, float):
            raise TypeError("Expected argument 'routing_weight' to be a float")
        pulumi.set(__self__, "routing_weight", routing_weight)
        if shared_key and not isinstance(shared_key, str):
            raise TypeError("Expected argument 'shared_key' to be a str")
        pulumi.set(__self__, "shared_key", shared_key)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if traffic_selector_policies and not isinstance(traffic_selector_policies, list):
            raise TypeError("Expected argument 'traffic_selector_policies' to be a list")
        pulumi.set(__self__, "traffic_selector_policies", traffic_selector_policies)
        if tunnel_connection_status and not isinstance(tunnel_connection_status, list):
            raise TypeError("Expected argument 'tunnel_connection_status' to be a list")
        pulumi.set(__self__, "tunnel_connection_status", tunnel_connection_status)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if use_policy_based_traffic_selectors and not isinstance(use_policy_based_traffic_selectors, bool):
            raise TypeError("Expected argument 'use_policy_based_traffic_selectors' to be a bool")
        pulumi.set(__self__, "use_policy_based_traffic_selectors", use_policy_based_traffic_selectors)
        if virtual_network_gateway1 and not isinstance(virtual_network_gateway1, dict):
            raise TypeError("Expected argument 'virtual_network_gateway1' to be a dict")
        pulumi.set(__self__, "virtual_network_gateway1", virtual_network_gateway1)
        if virtual_network_gateway2 and not isinstance(virtual_network_gateway2, dict):
            raise TypeError("Expected argument 'virtual_network_gateway2' to be a dict")
        pulumi.set(__self__, "virtual_network_gateway2", virtual_network_gateway2)

    @property
    @pulumi.getter(name="authorizationKey")
    def authorization_key(self) -> Optional[str]:
        """
        The authorizationKey.
        """
        return pulumi.get(self, "authorization_key")

    @property
    @pulumi.getter(name="connectionProtocol")
    def connection_protocol(self) -> Optional[str]:
        """
        Connection protocol used for this connection.
        """
        return pulumi.get(self, "connection_protocol")

    @property
    @pulumi.getter(name="connectionStatus")
    def connection_status(self) -> str:
        """
        Virtual Network Gateway connection status.
        """
        return pulumi.get(self, "connection_status")

    @property
    @pulumi.getter(name="connectionType")
    def connection_type(self) -> str:
        """
        Gateway connection type.
        """
        return pulumi.get(self, "connection_type")

    @property
    @pulumi.getter(name="egressBytesTransferred")
    def egress_bytes_transferred(self) -> float:
        """
        The egress bytes transferred in this connection.
        """
        return pulumi.get(self, "egress_bytes_transferred")

    @property
    @pulumi.getter(name="enableBgp")
    def enable_bgp(self) -> Optional[bool]:
        """
        EnableBgp flag.
        """
        return pulumi.get(self, "enable_bgp")

    @property
    @pulumi.getter
    def etag(self) -> Optional[str]:
        """
        A unique read-only string that changes whenever the resource is updated.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="expressRouteGatewayBypass")
    def express_route_gateway_bypass(self) -> Optional[bool]:
        """
        Bypass ExpressRoute Gateway for data forwarding.
        """
        return pulumi.get(self, "express_route_gateway_bypass")

    @property
    @pulumi.getter(name="ingressBytesTransferred")
    def ingress_bytes_transferred(self) -> float:
        """
        The ingress bytes transferred in this connection.
        """
        return pulumi.get(self, "ingress_bytes_transferred")

    @property
    @pulumi.getter(name="ipsecPolicies")
    def ipsec_policies(self) -> Optional[List['outputs.IpsecPolicyResponse']]:
        """
        The IPSec Policies to be considered by this connection.
        """
        return pulumi.get(self, "ipsec_policies")

    @property
    @pulumi.getter(name="localNetworkGateway2")
    def local_network_gateway2(self) -> Optional['outputs.LocalNetworkGatewayResponse']:
        """
        The reference to local network gateway resource.
        """
        return pulumi.get(self, "local_network_gateway2")

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
    @pulumi.getter
    def peer(self) -> Optional['outputs.SubResourceResponse']:
        """
        The reference to peerings resource.
        """
        return pulumi.get(self, "peer")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        The provisioning state of the virtual network gateway connection resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="resourceGuid")
    def resource_guid(self) -> Optional[str]:
        """
        The resource GUID property of the virtual network gateway connection resource.
        """
        return pulumi.get(self, "resource_guid")

    @property
    @pulumi.getter(name="routingWeight")
    def routing_weight(self) -> Optional[float]:
        """
        The routing weight.
        """
        return pulumi.get(self, "routing_weight")

    @property
    @pulumi.getter(name="sharedKey")
    def shared_key(self) -> Optional[str]:
        """
        The IPSec shared key.
        """
        return pulumi.get(self, "shared_key")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="trafficSelectorPolicies")
    def traffic_selector_policies(self) -> Optional[List['outputs.TrafficSelectorPolicyResponse']]:
        """
        The Traffic Selector Policies to be considered by this connection.
        """
        return pulumi.get(self, "traffic_selector_policies")

    @property
    @pulumi.getter(name="tunnelConnectionStatus")
    def tunnel_connection_status(self) -> List['outputs.TunnelConnectionHealthResponse']:
        """
        Collection of all tunnels' connection health status.
        """
        return pulumi.get(self, "tunnel_connection_status")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="usePolicyBasedTrafficSelectors")
    def use_policy_based_traffic_selectors(self) -> Optional[bool]:
        """
        Enable policy-based traffic selectors.
        """
        return pulumi.get(self, "use_policy_based_traffic_selectors")

    @property
    @pulumi.getter(name="virtualNetworkGateway1")
    def virtual_network_gateway1(self) -> 'outputs.VirtualNetworkGatewayResponse':
        """
        The reference to virtual network gateway resource.
        """
        return pulumi.get(self, "virtual_network_gateway1")

    @property
    @pulumi.getter(name="virtualNetworkGateway2")
    def virtual_network_gateway2(self) -> Optional['outputs.VirtualNetworkGatewayResponse']:
        """
        The reference to virtual network gateway resource.
        """
        return pulumi.get(self, "virtual_network_gateway2")


class AwaitableGetVirtualNetworkGatewayConnectionResult(GetVirtualNetworkGatewayConnectionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVirtualNetworkGatewayConnectionResult(
            authorization_key=self.authorization_key,
            connection_protocol=self.connection_protocol,
            connection_status=self.connection_status,
            connection_type=self.connection_type,
            egress_bytes_transferred=self.egress_bytes_transferred,
            enable_bgp=self.enable_bgp,
            etag=self.etag,
            express_route_gateway_bypass=self.express_route_gateway_bypass,
            ingress_bytes_transferred=self.ingress_bytes_transferred,
            ipsec_policies=self.ipsec_policies,
            local_network_gateway2=self.local_network_gateway2,
            location=self.location,
            name=self.name,
            peer=self.peer,
            provisioning_state=self.provisioning_state,
            resource_guid=self.resource_guid,
            routing_weight=self.routing_weight,
            shared_key=self.shared_key,
            tags=self.tags,
            traffic_selector_policies=self.traffic_selector_policies,
            tunnel_connection_status=self.tunnel_connection_status,
            type=self.type,
            use_policy_based_traffic_selectors=self.use_policy_based_traffic_selectors,
            virtual_network_gateway1=self.virtual_network_gateway1,
            virtual_network_gateway2=self.virtual_network_gateway2)


def get_virtual_network_gateway_connection(name: Optional[str] = None,
                                           resource_group_name: Optional[str] = None,
                                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVirtualNetworkGatewayConnectionResult:
    """
    Use this data source to access information about an existing resource.

    :param str name: The name of the virtual network gateway connection.
    :param str resource_group_name: The name of the resource group.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:network/v20190801:getVirtualNetworkGatewayConnection', __args__, opts=opts, typ=GetVirtualNetworkGatewayConnectionResult).value

    return AwaitableGetVirtualNetworkGatewayConnectionResult(
        authorization_key=__ret__.authorization_key,
        connection_protocol=__ret__.connection_protocol,
        connection_status=__ret__.connection_status,
        connection_type=__ret__.connection_type,
        egress_bytes_transferred=__ret__.egress_bytes_transferred,
        enable_bgp=__ret__.enable_bgp,
        etag=__ret__.etag,
        express_route_gateway_bypass=__ret__.express_route_gateway_bypass,
        ingress_bytes_transferred=__ret__.ingress_bytes_transferred,
        ipsec_policies=__ret__.ipsec_policies,
        local_network_gateway2=__ret__.local_network_gateway2,
        location=__ret__.location,
        name=__ret__.name,
        peer=__ret__.peer,
        provisioning_state=__ret__.provisioning_state,
        resource_guid=__ret__.resource_guid,
        routing_weight=__ret__.routing_weight,
        shared_key=__ret__.shared_key,
        tags=__ret__.tags,
        traffic_selector_policies=__ret__.traffic_selector_policies,
        tunnel_connection_status=__ret__.tunnel_connection_status,
        type=__ret__.type,
        use_policy_based_traffic_selectors=__ret__.use_policy_based_traffic_selectors,
        virtual_network_gateway1=__ret__.virtual_network_gateway1,
        virtual_network_gateway2=__ret__.virtual_network_gateway2)
