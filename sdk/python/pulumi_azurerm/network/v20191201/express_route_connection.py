# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['ExpressRouteConnection']


class ExpressRouteConnection(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorization_key: Optional[pulumi.Input[str]] = None,
                 connection_name: Optional[pulumi.Input[str]] = None,
                 enable_internet_security: Optional[pulumi.Input[bool]] = None,
                 express_route_circuit_peering: Optional[pulumi.Input[pulumi.InputType['ExpressRouteCircuitPeeringIdArgs']]] = None,
                 express_route_gateway_name: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 routing_weight: Optional[pulumi.Input[float]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        ExpressRouteConnection resource.

        ## Example Usage
        ### ExpressRouteConnectionCreate

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        express_route_connection = azurerm.network.v20191201.ExpressRouteConnection("expressRouteConnection",
            authorization_key="authorizationKey",
            connection_name="connectionName",
            express_route_circuit_peering={
                "id": "/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/expressRouteCircuits/circuitName/peerings/AzurePrivatePeering",
            },
            express_route_gateway_name="gateway-2",
            id="/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/expressRouteGateways/gateway-2/expressRouteConnections/connectionName",
            name="connectionName",
            resource_group_name="resourceGroupName",
            routing_weight=2)

        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] authorization_key: Authorization key to establish the connection.
        :param pulumi.Input[str] connection_name: The name of the connection subresource.
        :param pulumi.Input[bool] enable_internet_security: Enable internet security.
        :param pulumi.Input[pulumi.InputType['ExpressRouteCircuitPeeringIdArgs']] express_route_circuit_peering: The ExpressRoute circuit peering.
        :param pulumi.Input[str] express_route_gateway_name: The name of the ExpressRoute gateway.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[str] name: The name of the resource.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[float] routing_weight: The routing weight associated to the connection.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['authorization_key'] = authorization_key
            if connection_name is None:
                raise TypeError("Missing required property 'connection_name'")
            __props__['connection_name'] = connection_name
            __props__['enable_internet_security'] = enable_internet_security
            if express_route_circuit_peering is None:
                raise TypeError("Missing required property 'express_route_circuit_peering'")
            __props__['express_route_circuit_peering'] = express_route_circuit_peering
            if express_route_gateway_name is None:
                raise TypeError("Missing required property 'express_route_gateway_name'")
            __props__['express_route_gateway_name'] = express_route_gateway_name
            __props__['id'] = id
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['routing_weight'] = routing_weight
            __props__['provisioning_state'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:network/latest:ExpressRouteConnection"), pulumi.Alias(type_="azurerm:network/v20180801:ExpressRouteConnection"), pulumi.Alias(type_="azurerm:network/v20181001:ExpressRouteConnection"), pulumi.Alias(type_="azurerm:network/v20181101:ExpressRouteConnection"), pulumi.Alias(type_="azurerm:network/v20181201:ExpressRouteConnection"), pulumi.Alias(type_="azurerm:network/v20190201:ExpressRouteConnection"), pulumi.Alias(type_="azurerm:network/v20190401:ExpressRouteConnection"), pulumi.Alias(type_="azurerm:network/v20190601:ExpressRouteConnection"), pulumi.Alias(type_="azurerm:network/v20190701:ExpressRouteConnection"), pulumi.Alias(type_="azurerm:network/v20190801:ExpressRouteConnection"), pulumi.Alias(type_="azurerm:network/v20190901:ExpressRouteConnection"), pulumi.Alias(type_="azurerm:network/v20191101:ExpressRouteConnection"), pulumi.Alias(type_="azurerm:network/v20200301:ExpressRouteConnection"), pulumi.Alias(type_="azurerm:network/v20200401:ExpressRouteConnection"), pulumi.Alias(type_="azurerm:network/v20200501:ExpressRouteConnection"), pulumi.Alias(type_="azurerm:network/v20200601:ExpressRouteConnection")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(ExpressRouteConnection, __self__).__init__(
            'azurerm:network/v20191201:ExpressRouteConnection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ExpressRouteConnection':
        """
        Get an existing ExpressRouteConnection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ExpressRouteConnection(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="authorizationKey")
    def authorization_key(self) -> pulumi.Output[Optional[str]]:
        """
        Authorization key to establish the connection.
        """
        return pulumi.get(self, "authorization_key")

    @property
    @pulumi.getter(name="enableInternetSecurity")
    def enable_internet_security(self) -> pulumi.Output[Optional[bool]]:
        """
        Enable internet security.
        """
        return pulumi.get(self, "enable_internet_security")

    @property
    @pulumi.getter(name="expressRouteCircuitPeering")
    def express_route_circuit_peering(self) -> pulumi.Output['outputs.ExpressRouteCircuitPeeringIdResponse']:
        """
        The ExpressRoute circuit peering.
        """
        return pulumi.get(self, "express_route_circuit_peering")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        The provisioning state of the express route connection resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="routingWeight")
    def routing_weight(self) -> pulumi.Output[Optional[float]]:
        """
        The routing weight associated to the connection.
        """
        return pulumi.get(self, "routing_weight")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

