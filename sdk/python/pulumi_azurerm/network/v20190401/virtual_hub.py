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
    Gets a unique read-only string that changes whenever the resource is updated.
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
      * `express_route_gateway` (`dict`) - The expressRouteGateway associated with this VirtualHub.
        * `id` (`str`) - Resource ID.

      * `p2_s_vpn_gateway` (`dict`) - The P2SVpnGateway associated with this VirtualHub.
      * `provisioning_state` (`str`) - The provisioning state of the resource.
      * `route_table` (`dict`) - The routeTable associated with this virtual hub.
        * `routes` (`list`) - List of all routes.
          * `address_prefixes` (`list`) - List of all addressPrefixes.
          * `next_hop_ip_address` (`str`) - NextHop ip address.

      * `virtual_network_connections` (`list`) - List of all vnet connections with this VirtualHub.
        * `etag` (`str`) - Gets a unique read-only string that changes whenever the resource is updated.
        * `id` (`str`) - Resource ID.
        * `name` (`str`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
        * `properties` (`dict`) - Properties of the hub virtual network connection.
          * `allow_hub_to_remote_vnet_transit` (`bool`) - VirtualHub to RemoteVnet transit to enabled or not.
          * `allow_remote_vnet_to_use_hub_vnet_gateways` (`bool`) - Allow RemoteVnet to use Virtual Hub's gateways.
          * `enable_internet_security` (`bool`) - Enable internet security.
          * `provisioning_state` (`str`) - The provisioning state of the resource.
          * `remote_virtual_network` (`dict`) - Reference to the remote virtual network.

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
    def __init__(__self__, resource_name, opts=None, address_prefix=None, express_route_gateway=None, id=None, location=None, name=None, p2_s_vpn_gateway=None, provisioning_state=None, resource_group_name=None, route_table=None, tags=None, virtual_network_connections=None, virtual_wan=None, vpn_gateway=None, __props__=None, __name__=None, __opts__=None):
        """
        VirtualHub Resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address_prefix: Address-prefix for this VirtualHub.
        :param pulumi.Input[dict] express_route_gateway: The expressRouteGateway associated with this VirtualHub.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[str] name: The name of the VirtualHub.
        :param pulumi.Input[dict] p2_s_vpn_gateway: The P2SVpnGateway associated with this VirtualHub.
        :param pulumi.Input[str] provisioning_state: The provisioning state of the resource.
        :param pulumi.Input[str] resource_group_name: The resource group name of the VirtualHub.
        :param pulumi.Input[dict] route_table: The routeTable associated with this virtual hub.
        :param pulumi.Input[dict] tags: Resource tags.
        :param pulumi.Input[list] virtual_network_connections: List of all vnet connections with this VirtualHub.
        :param pulumi.Input[dict] virtual_wan: The VirtualWAN to which the VirtualHub belongs.
        :param pulumi.Input[dict] vpn_gateway: The VpnGateway associated with this VirtualHub.

        The **express_route_gateway** object supports the following:

          * `id` (`pulumi.Input[str]`) - Resource ID.

        The **route_table** object supports the following:

          * `routes` (`pulumi.Input[list]`) - List of all routes.
            * `address_prefixes` (`pulumi.Input[list]`) - List of all addressPrefixes.
            * `next_hop_ip_address` (`pulumi.Input[str]`) - NextHop ip address.

        The **virtual_network_connections** object supports the following:

          * `allow_hub_to_remote_vnet_transit` (`pulumi.Input[bool]`) - VirtualHub to RemoteVnet transit to enabled or not.
          * `allow_remote_vnet_to_use_hub_vnet_gateways` (`pulumi.Input[bool]`) - Allow RemoteVnet to use Virtual Hub's gateways.
          * `enable_internet_security` (`pulumi.Input[bool]`) - Enable internet security.
          * `id` (`pulumi.Input[str]`) - Resource ID.
          * `name` (`pulumi.Input[str]`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
          * `provisioning_state` (`pulumi.Input[str]`) - The provisioning state of the resource.
          * `remote_virtual_network` (`pulumi.Input[dict]`) - Reference to the remote virtual network.
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
            __props__['express_route_gateway'] = express_route_gateway
            __props__['id'] = id
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['p2_s_vpn_gateway'] = p2_s_vpn_gateway
            __props__['provisioning_state'] = provisioning_state
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['route_table'] = route_table
            __props__['tags'] = tags
            __props__['virtual_network_connections'] = virtual_network_connections
            __props__['virtual_wan'] = virtual_wan
            __props__['vpn_gateway'] = vpn_gateway
            __props__['etag'] = None
            __props__['properties'] = None
            __props__['type'] = None
        super(VirtualHub, __self__).__init__(
            'azurerm:network/v20190401:VirtualHub',
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
