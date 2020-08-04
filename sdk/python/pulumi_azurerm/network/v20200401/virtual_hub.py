# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class VirtualHub(pulumi.CustomResource):
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
    Properties of the virtual hub.
      * `address_prefix` (`str`) - Address-prefix for this VirtualHub.
      * `azure_firewall` (`dict`) - The azureFirewall associated with this VirtualHub.
        * `id` (`str`) - Resource ID.

      * `express_route_gateway` (`dict`) - The expressRouteGateway associated with this VirtualHub.
      * `p2_s_vpn_gateway` (`dict`) - The P2SVpnGateway associated with this VirtualHub.
      * `provisioning_state` (`str`) - The provisioning state of the virtual hub resource.
      * `route_table` (`dict`) - The routeTable associated with this virtual hub.
        * `routes` (`list`) - List of all routes.
          * `address_prefixes` (`list`) - List of all addressPrefixes.
          * `next_hop_ip_address` (`str`) - NextHop ip address.

      * `security_partner_provider` (`dict`) - The securityPartnerProvider associated with this VirtualHub.
      * `security_provider_name` (`str`) - The Security Provider name.
      * `sku` (`str`) - The sku of this VirtualHub.
      * `virtual_hub_route_table_v2s` (`list`) - List of all virtual hub route table v2s associated with this VirtualHub.
        * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
        * `id` (`str`) - Resource ID.
        * `name` (`str`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
        * `properties` (`dict`) - Properties of the virtual hub route table v2.
          * `attached_connections` (`list`) - List of all connections attached to this route table v2.
          * `provisioning_state` (`str`) - The provisioning state of the virtual hub route table v2 resource.
          * `routes` (`list`) - List of all routes.
            * `destination_type` (`str`) - The type of destinations.
            * `destinations` (`list`) - List of all destinations.
            * `next_hop_type` (`str`) - The type of next hops.
            * `next_hops` (`list`) - NextHops ip address.

      * `virtual_network_connections` (`list`) - List of all vnet connections with this VirtualHub.
        * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
        * `id` (`str`) - Resource ID.
        * `name` (`str`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
        * `properties` (`dict`) - Properties of the hub virtual network connection.
          * `allow_hub_to_remote_vnet_transit` (`bool`) - VirtualHub to RemoteVnet transit to enabled or not.
          * `allow_remote_vnet_to_use_hub_vnet_gateways` (`bool`) - Allow RemoteVnet to use Virtual Hub's gateways.
          * `enable_internet_security` (`bool`) - Enable internet security.
          * `provisioning_state` (`str`) - The provisioning state of the hub virtual network connection resource.
          * `remote_virtual_network` (`dict`) - Reference to the remote virtual network.
          * `routing_configuration` (`dict`) - The Routing Configuration indicating the associated and propagated route tables on this connection.
            * `associated_route_table` (`dict`) - The resource id RouteTable associated with this RoutingConfiguration.
            * `propagated_route_tables` (`dict`) - The list of RouteTables to advertise the routes to.
              * `ids` (`list`) - The list of resource ids of all the RouteTables.
              * `labels` (`list`) - The list of labels.

            * `vnet_routes` (`dict`) - List of routes that control routing from VirtualHub into a virtual network connection.
              * `static_routes` (`list`) - List of all Static Routes.
                * `address_prefixes` (`list`) - List of all address prefixes.
                * `name` (`str`) - The name of the StaticRoute that is unique within a VnetRoute.
                * `next_hop_ip_address` (`str`) - The ip address of the next hop.

      * `virtual_wan` (`dict`) - The VirtualWAN to which the VirtualHub belongs.
      * `vpn_gateway` (`dict`) - The VpnGateway associated with this VirtualHub.
    """
    tags: pulumi.Output[dict]
    """
    Resource tags.
    """
    type: pulumi.Output[str]
    """
    Resource type.
    """
    def __init__(__self__, resource_name, opts=None, id=None, location=None, name=None, properties=None, resource_group_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        VirtualHub Resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[str] name: The name of the VirtualHub.
        :param pulumi.Input[dict] properties: Properties of the virtual hub.
        :param pulumi.Input[str] resource_group_name: The resource group name of the VirtualHub.
        :param pulumi.Input[dict] tags: Resource tags.

        The **properties** object supports the following:

          * `address_prefix` (`pulumi.Input[str]`) - Address-prefix for this VirtualHub.
          * `azure_firewall` (`pulumi.Input[dict]`) - The azureFirewall associated with this VirtualHub.
            * `id` (`pulumi.Input[str]`) - Resource ID.

          * `express_route_gateway` (`pulumi.Input[dict]`) - The expressRouteGateway associated with this VirtualHub.
          * `p2_s_vpn_gateway` (`pulumi.Input[dict]`) - The P2SVpnGateway associated with this VirtualHub.
          * `route_table` (`pulumi.Input[dict]`) - The routeTable associated with this virtual hub.
            * `routes` (`pulumi.Input[list]`) - List of all routes.
              * `address_prefixes` (`pulumi.Input[list]`) - List of all addressPrefixes.
              * `next_hop_ip_address` (`pulumi.Input[str]`) - NextHop ip address.

          * `security_partner_provider` (`pulumi.Input[dict]`) - The securityPartnerProvider associated with this VirtualHub.
          * `security_provider_name` (`pulumi.Input[str]`) - The Security Provider name.
          * `sku` (`pulumi.Input[str]`) - The sku of this VirtualHub.
          * `virtual_hub_route_table_v2s` (`pulumi.Input[list]`) - List of all virtual hub route table v2s associated with this VirtualHub.
            * `etag` (`pulumi.Input[str]`) - A unique read-only string that changes whenever the resource is updated.
            * `name` (`pulumi.Input[str]`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
            * `properties` (`pulumi.Input[dict]`) - Properties of the virtual hub route table v2.
              * `attached_connections` (`pulumi.Input[list]`) - List of all connections attached to this route table v2.
              * `provisioning_state` (`pulumi.Input[str]`) - The provisioning state of the virtual hub route table v2 resource.
              * `routes` (`pulumi.Input[list]`) - List of all routes.
                * `destination_type` (`pulumi.Input[str]`) - The type of destinations.
                * `destinations` (`pulumi.Input[list]`) - List of all destinations.
                * `next_hop_type` (`pulumi.Input[str]`) - The type of next hops.
                * `next_hops` (`pulumi.Input[list]`) - NextHops ip address.

          * `virtual_network_connections` (`pulumi.Input[list]`) - List of all vnet connections with this VirtualHub.
            * `id` (`pulumi.Input[str]`) - Resource ID.
            * `name` (`pulumi.Input[str]`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
            * `properties` (`pulumi.Input[dict]`) - Properties of the hub virtual network connection.
              * `allow_hub_to_remote_vnet_transit` (`pulumi.Input[bool]`) - VirtualHub to RemoteVnet transit to enabled or not.
              * `allow_remote_vnet_to_use_hub_vnet_gateways` (`pulumi.Input[bool]`) - Allow RemoteVnet to use Virtual Hub's gateways.
              * `enable_internet_security` (`pulumi.Input[bool]`) - Enable internet security.
              * `remote_virtual_network` (`pulumi.Input[dict]`) - Reference to the remote virtual network.
              * `routing_configuration` (`pulumi.Input[dict]`) - The Routing Configuration indicating the associated and propagated route tables on this connection.
                * `associated_route_table` (`pulumi.Input[dict]`) - The resource id RouteTable associated with this RoutingConfiguration.
                * `propagated_route_tables` (`pulumi.Input[dict]`) - The list of RouteTables to advertise the routes to.
                  * `ids` (`pulumi.Input[list]`) - The list of resource ids of all the RouteTables.
                  * `labels` (`pulumi.Input[list]`) - The list of labels.

                * `vnet_routes` (`pulumi.Input[dict]`) - List of routes that control routing from VirtualHub into a virtual network connection.
                  * `static_routes` (`pulumi.Input[list]`) - List of all Static Routes.
                    * `address_prefixes` (`pulumi.Input[list]`) - List of all address prefixes.
                    * `name` (`pulumi.Input[str]`) - The name of the StaticRoute that is unique within a VnetRoute.
                    * `next_hop_ip_address` (`pulumi.Input[str]`) - The ip address of the next hop.

          * `virtual_wan` (`pulumi.Input[dict]`) - The VirtualWAN to which the VirtualHub belongs.
          * `vpn_gateway` (`pulumi.Input[dict]`) - The VpnGateway associated with this VirtualHub.
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
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['etag'] = None
            __props__['type'] = None
        super(VirtualHub, __self__).__init__(
            'azurerm:network/v20200401:VirtualHub',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing VirtualHub resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return VirtualHub(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
