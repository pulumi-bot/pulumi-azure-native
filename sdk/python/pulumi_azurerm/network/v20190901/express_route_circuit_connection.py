# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class ExpressRouteCircuitConnection(pulumi.CustomResource):
    etag: pulumi.Output[str]
    """
    A unique read-only string that changes whenever the resource is updated.
    """
    name: pulumi.Output[str]
    """
    The name of the resource that is unique within a resource group. This name can be used to access the resource.
    """
    properties: pulumi.Output[dict]
    """
    Properties of the express route circuit connection.
      * `address_prefix` (`str`) - /29 IP address space to carve out Customer addresses for tunnels.
      * `authorization_key` (`str`) - The authorization key.
      * `circuit_connection_status` (`str`) - Express Route Circuit connection state.
      * `express_route_circuit_peering` (`dict`) - Reference to Express Route Circuit Private Peering Resource of the circuit initiating connection.
        * `id` (`str`) - Resource ID.

      * `peer_express_route_circuit_peering` (`dict`) - Reference to Express Route Circuit Private Peering Resource of the peered circuit.
      * `provisioning_state` (`str`) - The provisioning state of the express route circuit connection resource.
    """
    type: pulumi.Output[str]
    """
    Type of the resource.
    """
    def __init__(__self__, resource_name, opts=None, address_prefix=None, authorization_key=None, circuit_connection_status=None, circuit_name=None, express_route_circuit_peering=None, id=None, name=None, peer_express_route_circuit_peering=None, peering_name=None, resource_group_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Express Route Circuit Connection in an ExpressRouteCircuitPeering resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address_prefix: /29 IP address space to carve out Customer addresses for tunnels.
        :param pulumi.Input[str] authorization_key: The authorization key.
        :param pulumi.Input[str] circuit_connection_status: Express Route Circuit connection state.
        :param pulumi.Input[str] circuit_name: The name of the express route circuit.
        :param pulumi.Input[dict] express_route_circuit_peering: Reference to Express Route Circuit Private Peering Resource of the circuit initiating connection.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[str] name: The name of the express route circuit connection.
        :param pulumi.Input[dict] peer_express_route_circuit_peering: Reference to Express Route Circuit Private Peering Resource of the peered circuit.
        :param pulumi.Input[str] peering_name: The name of the peering.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.

        The **express_route_circuit_peering** object supports the following:

          * `id` (`pulumi.Input[str]`) - Resource ID.
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
            __props__['circuit_connection_status'] = circuit_connection_status
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
            __props__['etag'] = None
            __props__['properties'] = None
            __props__['type'] = None
        super(ExpressRouteCircuitConnection, __self__).__init__(
            'azurerm:network/v20190901:ExpressRouteCircuitConnection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing ExpressRouteCircuitConnection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ExpressRouteCircuitConnection(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
