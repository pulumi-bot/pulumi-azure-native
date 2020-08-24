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

__all__ = ['ExpressRouteCircuitConnection']


class ExpressRouteCircuitConnection(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address_prefix: Optional[pulumi.Input[str]] = None,
                 authorization_key: Optional[pulumi.Input[str]] = None,
                 circuit_name: Optional[pulumi.Input[str]] = None,
                 express_route_circuit_peering: Optional[pulumi.Input[pulumi.InputType['SubResourceArgs']]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 peer_express_route_circuit_peering: Optional[pulumi.Input[pulumi.InputType['SubResourceArgs']]] = None,
                 peering_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Express Route Circuit Connection in an ExpressRouteCircuitPeering resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address_prefix: /29 IP address space to carve out Customer addresses for tunnels.
        :param pulumi.Input[str] authorization_key: The authorization key.
        :param pulumi.Input[str] circuit_name: The name of the express route circuit.
        :param pulumi.Input[pulumi.InputType['SubResourceArgs']] express_route_circuit_peering: Reference to Express Route Circuit Private Peering Resource of the circuit initiating connection.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[str] name: The name of the express route circuit connection.
        :param pulumi.Input[pulumi.InputType['SubResourceArgs']] peer_express_route_circuit_peering: Reference to Express Route Circuit Private Peering Resource of the peered circuit.
        :param pulumi.Input[str] peering_name: The name of the peering.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
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

            __props__['address_prefix'] = address_prefix
            __props__['authorization_key'] = authorization_key
            if circuit_name is None:
                raise TypeError("Missing required property 'circuit_name'")
            __props__['circuit_name'] = circuit_name
            __props__['express_route_circuit_peering'] = express_route_circuit_peering
            __props__['id'] = id
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['peer_express_route_circuit_peering'] = peer_express_route_circuit_peering
            if peering_name is None:
                raise TypeError("Missing required property 'peering_name'")
            __props__['peering_name'] = peering_name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['circuit_connection_status'] = None
            __props__['etag'] = None
            __props__['provisioning_state'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:network/v20180201:ExpressRouteCircuitConnection"), pulumi.Alias(type_="azurerm:network/v20180401:ExpressRouteCircuitConnection"), pulumi.Alias(type_="azurerm:network/v20180601:ExpressRouteCircuitConnection"), pulumi.Alias(type_="azurerm:network/v20180701:ExpressRouteCircuitConnection"), pulumi.Alias(type_="azurerm:network/v20180801:ExpressRouteCircuitConnection"), pulumi.Alias(type_="azurerm:network/v20181101:ExpressRouteCircuitConnection"), pulumi.Alias(type_="azurerm:network/v20181201:ExpressRouteCircuitConnection"), pulumi.Alias(type_="azurerm:network/v20190201:ExpressRouteCircuitConnection"), pulumi.Alias(type_="azurerm:network/v20190401:ExpressRouteCircuitConnection"), pulumi.Alias(type_="azurerm:network/v20190601:ExpressRouteCircuitConnection"), pulumi.Alias(type_="azurerm:network/v20190701:ExpressRouteCircuitConnection"), pulumi.Alias(type_="azurerm:network/v20190801:ExpressRouteCircuitConnection"), pulumi.Alias(type_="azurerm:network/v20190901:ExpressRouteCircuitConnection"), pulumi.Alias(type_="azurerm:network/v20191101:ExpressRouteCircuitConnection"), pulumi.Alias(type_="azurerm:network/v20191201:ExpressRouteCircuitConnection"), pulumi.Alias(type_="azurerm:network/v20200301:ExpressRouteCircuitConnection"), pulumi.Alias(type_="azurerm:network/v20200401:ExpressRouteCircuitConnection"), pulumi.Alias(type_="azurerm:network/v20200501:ExpressRouteCircuitConnection"), pulumi.Alias(type_="azurerm:network/v20200601:ExpressRouteCircuitConnection")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(ExpressRouteCircuitConnection, __self__).__init__(
            'azurerm:network/v20181001:ExpressRouteCircuitConnection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ExpressRouteCircuitConnection':
        """
        Get an existing ExpressRouteCircuitConnection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ExpressRouteCircuitConnection(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="addressPrefix")
    def address_prefix(self) -> Optional[str]:
        """
        /29 IP address space to carve out Customer addresses for tunnels.
        """
        return pulumi.get(self, "address_prefix")

    @property
    @pulumi.getter(name="authorizationKey")
    def authorization_key(self) -> Optional[str]:
        """
        The authorization key.
        """
        return pulumi.get(self, "authorization_key")

    @property
    @pulumi.getter(name="circuitConnectionStatus")
    def circuit_connection_status(self) -> str:
        """
        Express Route Circuit Connection State. Possible values are: 'Connected' and 'Disconnected'.
        """
        return pulumi.get(self, "circuit_connection_status")

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        A unique read-only string that changes whenever the resource is updated.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="expressRouteCircuitPeering")
    def express_route_circuit_peering(self) -> Optional['outputs.SubResourceResponse']:
        """
        Reference to Express Route Circuit Private Peering Resource of the circuit initiating connection.
        """
        return pulumi.get(self, "express_route_circuit_peering")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Gets name of the resource that is unique within a resource group. This name can be used to access the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="peerExpressRouteCircuitPeering")
    def peer_express_route_circuit_peering(self) -> Optional['outputs.SubResourceResponse']:
        """
        Reference to Express Route Circuit Private Peering Resource of the peered circuit.
        """
        return pulumi.get(self, "peer_express_route_circuit_peering")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        Provisioning state of the circuit connection resource. Possible values are: 'Succeeded', 'Updating', 'Deleting', and 'Failed'.
        """
        return pulumi.get(self, "provisioning_state")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

