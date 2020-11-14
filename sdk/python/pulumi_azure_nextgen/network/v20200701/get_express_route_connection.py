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
    'GetExpressRouteConnectionResult',
    'AwaitableGetExpressRouteConnectionResult',
    'get_express_route_connection',
]

@pulumi.output_type
class GetExpressRouteConnectionResult:
    """
    ExpressRouteConnection resource.
    """
    def __init__(__self__, authorization_key=None, enable_internet_security=None, express_route_circuit_peering=None, name=None, provisioning_state=None, routing_configuration=None, routing_weight=None):
        if authorization_key and not isinstance(authorization_key, str):
            raise TypeError("Expected argument 'authorization_key' to be a str")
        pulumi.set(__self__, "authorization_key", authorization_key)
        if enable_internet_security and not isinstance(enable_internet_security, bool):
            raise TypeError("Expected argument 'enable_internet_security' to be a bool")
        pulumi.set(__self__, "enable_internet_security", enable_internet_security)
        if express_route_circuit_peering and not isinstance(express_route_circuit_peering, dict):
            raise TypeError("Expected argument 'express_route_circuit_peering' to be a dict")
        pulumi.set(__self__, "express_route_circuit_peering", express_route_circuit_peering)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if routing_configuration and not isinstance(routing_configuration, dict):
            raise TypeError("Expected argument 'routing_configuration' to be a dict")
        pulumi.set(__self__, "routing_configuration", routing_configuration)
        if routing_weight and not isinstance(routing_weight, int):
            raise TypeError("Expected argument 'routing_weight' to be a int")
        pulumi.set(__self__, "routing_weight", routing_weight)

    @property
    @pulumi.getter(name="authorizationKey")
    def authorization_key(self) -> Optional[str]:
        """
        Authorization key to establish the connection.
        """
        return pulumi.get(self, "authorization_key")

    @property
    @pulumi.getter(name="enableInternetSecurity")
    def enable_internet_security(self) -> Optional[bool]:
        """
        Enable internet security.
        """
        return pulumi.get(self, "enable_internet_security")

    @property
    @pulumi.getter(name="expressRouteCircuitPeering")
    def express_route_circuit_peering(self) -> 'outputs.ExpressRouteCircuitPeeringIdResponse':
        """
        The ExpressRoute circuit peering.
        """
        return pulumi.get(self, "express_route_circuit_peering")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        The provisioning state of the express route connection resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="routingConfiguration")
    def routing_configuration(self) -> Optional['outputs.RoutingConfigurationResponse']:
        """
        The Routing Configuration indicating the associated and propagated route tables on this connection.
        """
        return pulumi.get(self, "routing_configuration")

    @property
    @pulumi.getter(name="routingWeight")
    def routing_weight(self) -> Optional[int]:
        """
        The routing weight associated to the connection.
        """
        return pulumi.get(self, "routing_weight")


class AwaitableGetExpressRouteConnectionResult(GetExpressRouteConnectionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetExpressRouteConnectionResult(
            authorization_key=self.authorization_key,
            enable_internet_security=self.enable_internet_security,
            express_route_circuit_peering=self.express_route_circuit_peering,
            name=self.name,
            provisioning_state=self.provisioning_state,
            routing_configuration=self.routing_configuration,
            routing_weight=self.routing_weight)


def get_express_route_connection(connection_name: Optional[str] = None,
                                 express_route_gateway_name: Optional[str] = None,
                                 resource_group_name: Optional[str] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetExpressRouteConnectionResult:
    """
    Use this data source to access information about an existing resource.

    :param str connection_name: The name of the ExpressRoute connection.
    :param str express_route_gateway_name: The name of the ExpressRoute gateway.
    :param str resource_group_name: The name of the resource group.
    """
    __args__ = dict()
    __args__['connectionName'] = connection_name
    __args__['expressRouteGatewayName'] = express_route_gateway_name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure-nextgen:network/v20200701:getExpressRouteConnection', __args__, opts=opts, typ=GetExpressRouteConnectionResult).value

    return AwaitableGetExpressRouteConnectionResult(
        authorization_key=__ret__.authorization_key,
        enable_internet_security=__ret__.enable_internet_security,
        express_route_circuit_peering=__ret__.express_route_circuit_peering,
        name=__ret__.name,
        provisioning_state=__ret__.provisioning_state,
        routing_configuration=__ret__.routing_configuration,
        routing_weight=__ret__.routing_weight)
