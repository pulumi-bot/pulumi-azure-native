# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class P2sVpnGateway(pulumi.CustomResource):
    etag: pulumi.Output[str]
    """
    A unique read-only string that changes whenever the resource is updated.
    """
    location: pulumi.Output[str]
    """
    Resource location.
    """
    name: pulumi.Output[str]
    """
    Resource name.
    """
    properties: pulumi.Output[dict]
    """
    Properties of the P2SVpnGateway.
      * `p2_s_connection_configurations` (`list`) - List of all p2s connection configurations of the gateway.
        * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
        * `id` (`str`) - Resource ID.
        * `name` (`str`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
        * `properties` (`dict`) - Properties of the P2S connection configuration.
          * `provisioning_state` (`str`) - The provisioning state of the P2SConnectionConfiguration resource.
          * `routing_configuration` (`dict`) - The Routing Configuration indicating the associated and propagated route tables on this connection.
            * `associated_route_table` (`dict`) - The resource id RouteTable associated with this RoutingConfiguration.
              * `id` (`str`) - Resource ID.

            * `propagated_route_tables` (`dict`) - The list of RouteTables to advertise the routes to.
              * `ids` (`list`) - The list of resource ids of all the RouteTables.
              * `labels` (`list`) - The list of labels.

            * `vnet_routes` (`dict`) - List of routes that control routing from VirtualHub into a virtual network connection.
              * `static_routes` (`list`) - List of all Static Routes.
                * `address_prefixes` (`list`) - List of all address prefixes.
                * `name` (`str`) - The name of the StaticRoute that is unique within a VnetRoute.
                * `next_hop_ip_address` (`str`) - The ip address of the next hop.

          * `vpn_client_address_pool` (`dict`) - The reference to the address space resource which represents Address space for P2S VpnClient.
            * `address_prefixes` (`list`) - A list of address blocks reserved for this virtual network in CIDR notation.

      * `provisioning_state` (`str`) - The provisioning state of the P2S VPN gateway resource.
      * `virtual_hub` (`dict`) - The VirtualHub to which the gateway belongs.
      * `vpn_client_connection_health` (`dict`) - All P2S VPN clients' connection health status.
        * `allocated_ip_addresses` (`list`) - List of allocated ip addresses to the connected p2s vpn clients.
        * `total_egress_bytes_transferred` (`float`) - Total of the Egress Bytes Transferred in this connection.
        * `total_ingress_bytes_transferred` (`float`) - Total of the Ingress Bytes Transferred in this P2S Vpn connection.
        * `vpn_client_connections_count` (`float`) - The total of p2s vpn clients connected at this time to this P2SVpnGateway.

      * `vpn_gateway_scale_unit` (`float`) - The scale unit for this p2s vpn gateway.
      * `vpn_server_configuration` (`dict`) - The VpnServerConfiguration to which the p2sVpnGateway is attached to.
    """
    tags: pulumi.Output[dict]
    """
    Resource tags.
    """
    type: pulumi.Output[str]
    """
    Resource type.
    """
    def __init__(__self__, resource_name, opts=None, id=None, location=None, name=None, p2_s_connection_configurations=None, resource_group_name=None, tags=None, virtual_hub=None, vpn_gateway_scale_unit=None, vpn_server_configuration=None, __props__=None, __name__=None, __opts__=None):
        """
        P2SVpnGateway Resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[str] name: The name of the gateway.
        :param pulumi.Input[list] p2_s_connection_configurations: List of all p2s connection configurations of the gateway.
        :param pulumi.Input[str] resource_group_name: The resource group name of the P2SVpnGateway.
        :param pulumi.Input[dict] tags: Resource tags.
        :param pulumi.Input[dict] virtual_hub: The VirtualHub to which the gateway belongs.
        :param pulumi.Input[float] vpn_gateway_scale_unit: The scale unit for this p2s vpn gateway.
        :param pulumi.Input[dict] vpn_server_configuration: The VpnServerConfiguration to which the p2sVpnGateway is attached to.

        The **p2_s_connection_configurations** object supports the following:

          * `id` (`pulumi.Input[str]`) - Resource ID.
          * `name` (`pulumi.Input[str]`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
          * `routing_configuration` (`pulumi.Input[dict]`) - The Routing Configuration indicating the associated and propagated route tables on this connection.
            * `associated_route_table` (`pulumi.Input[dict]`) - The resource id RouteTable associated with this RoutingConfiguration.
              * `id` (`pulumi.Input[str]`) - Resource ID.

            * `propagated_route_tables` (`pulumi.Input[dict]`) - The list of RouteTables to advertise the routes to.
              * `ids` (`pulumi.Input[list]`) - The list of resource ids of all the RouteTables.
              * `labels` (`pulumi.Input[list]`) - The list of labels.

            * `vnet_routes` (`pulumi.Input[dict]`) - List of routes that control routing from VirtualHub into a virtual network connection.
              * `static_routes` (`pulumi.Input[list]`) - List of all Static Routes.
                * `address_prefixes` (`pulumi.Input[list]`) - List of all address prefixes.
                * `name` (`pulumi.Input[str]`) - The name of the StaticRoute that is unique within a VnetRoute.
                * `next_hop_ip_address` (`pulumi.Input[str]`) - The ip address of the next hop.

          * `vpn_client_address_pool` (`pulumi.Input[dict]`) - The reference to the address space resource which represents Address space for P2S VpnClient.
            * `address_prefixes` (`pulumi.Input[list]`) - A list of address blocks reserved for this virtual network in CIDR notation.
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

            __props__['id'] = id
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['p2_s_connection_configurations'] = p2_s_connection_configurations
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['virtual_hub'] = virtual_hub
            __props__['vpn_gateway_scale_unit'] = vpn_gateway_scale_unit
            __props__['vpn_server_configuration'] = vpn_server_configuration
            __props__['etag'] = None
            __props__['properties'] = None
            __props__['type'] = None
        super(P2sVpnGateway, __self__).__init__(
            'azurerm:network/v20200401:P2sVpnGateway',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing P2sVpnGateway resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return P2sVpnGateway(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
