# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class RouteFilter(pulumi.CustomResource):
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
    Properties of the route filter.
      * `ipv6_peerings` (`list`) - A collection of references to express route circuit ipv6 peerings.
        * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
        * `id` (`str`) - Resource ID.
        * `name` (`str`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
        * `properties` (`dict`) - Properties of the express route circuit peering.
          * `azure_asn` (`float`) - The Azure ASN.
          * `connections` (`list`) - The list of circuit connections associated with Azure Private Peering for this circuit.
            * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
            * `id` (`str`) - Resource ID.
            * `name` (`str`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
            * `properties` (`dict`) - Properties of the express route circuit connection.
              * `address_prefix` (`str`) - /29 IP address space to carve out Customer addresses for tunnels.
              * `authorization_key` (`str`) - The authorization key.
              * `circuit_connection_status` (`str`) - Express Route Circuit connection state.
              * `express_route_circuit_peering` (`dict`) - Reference to Express Route Circuit Private Peering Resource of the circuit initiating connection.
                * `id` (`str`) - Resource ID.

              * `peer_express_route_circuit_peering` (`dict`) - Reference to Express Route Circuit Private Peering Resource of the peered circuit.
              * `provisioning_state` (`str`) - The provisioning state of the express route circuit connection resource.

            * `type` (`str`) - Type of the resource.

          * `express_route_connection` (`dict`) - The ExpressRoute connection.
            * `id` (`str`) - The ID of the ExpressRouteConnection.

          * `gateway_manager_etag` (`str`) - The GatewayManager Etag.
          * `ipv6_peering_config` (`dict`) - The IPv6 peering configuration.
            * `microsoft_peering_config` (`dict`) - The Microsoft peering configuration.
              * `advertised_communities` (`list`) - The communities of bgp peering. Specified for microsoft peering.
              * `advertised_public_prefixes` (`list`) - The reference to AdvertisedPublicPrefixes.
              * `advertised_public_prefixes_state` (`str`) - The advertised public prefix state of the Peering resource.
              * `customer_asn` (`float`) - The CustomerASN of the peering.
              * `legacy_mode` (`float`) - The legacy mode of the peering.
              * `routing_registry_name` (`str`) - The RoutingRegistryName of the configuration.

            * `primary_peer_address_prefix` (`str`) - The primary address prefix.
            * `route_filter` (`dict`) - The reference to the RouteFilter resource.
            * `secondary_peer_address_prefix` (`str`) - The secondary address prefix.
            * `state` (`str`) - The state of peering.

          * `last_modified_by` (`str`) - Who was the last to modify the peering.
          * `microsoft_peering_config` (`dict`) - The Microsoft peering configuration.
          * `peer_asn` (`float`) - The peer ASN.
          * `peered_connections` (`list`) - The list of peered circuit connections associated with Azure Private Peering for this circuit.
            * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
            * `id` (`str`) - Resource ID.
            * `name` (`str`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
            * `properties` (`dict`) - Properties of the peer express route circuit connection.
              * `address_prefix` (`str`) - /29 IP address space to carve out Customer addresses for tunnels.
              * `auth_resource_guid` (`str`) - The resource guid of the authorization used for the express route circuit connection.
              * `circuit_connection_status` (`str`) - Express Route Circuit connection state.
              * `connection_name` (`str`) - The name of the express route circuit connection resource.
              * `express_route_circuit_peering` (`dict`) - Reference to Express Route Circuit Private Peering Resource of the circuit.
              * `peer_express_route_circuit_peering` (`dict`) - Reference to Express Route Circuit Private Peering Resource of the peered circuit.
              * `provisioning_state` (`str`) - The provisioning state of the peer express route circuit connection resource.

            * `type` (`str`) - Type of the resource.

          * `peering_type` (`str`) - The peering type.
          * `primary_azure_port` (`str`) - The primary port.
          * `primary_peer_address_prefix` (`str`) - The primary address prefix.
          * `provisioning_state` (`str`) - The provisioning state of the express route circuit peering resource.
          * `route_filter` (`dict`) - The reference to the RouteFilter resource.
          * `secondary_azure_port` (`str`) - The secondary port.
          * `secondary_peer_address_prefix` (`str`) - The secondary address prefix.
          * `shared_key` (`str`) - The shared key.
          * `state` (`str`) - The peering state.
          * `stats` (`dict`) - The peering stats of express route circuit.
            * `primarybytes_in` (`float`) - The Primary BytesIn of the peering.
            * `primarybytes_out` (`float`) - The primary BytesOut of the peering.
            * `secondarybytes_in` (`float`) - The secondary BytesIn of the peering.
            * `secondarybytes_out` (`float`) - The secondary BytesOut of the peering.

          * `vlan_id` (`float`) - The VLAN ID.

        * `type` (`str`) - Type of the resource.

      * `peerings` (`list`) - A collection of references to express route circuit peerings.
      * `provisioning_state` (`str`) - The provisioning state of the route filter resource.
      * `rules` (`list`) - Collection of RouteFilterRules contained within a route filter.
        * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
        * `id` (`str`) - Resource ID.
        * `location` (`str`) - Resource location.
        * `name` (`str`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
        * `properties` (`dict`) - Properties of the route filter rule.
          * `access` (`str`) - The access type of the rule.
          * `communities` (`list`) - The collection for bgp community values to filter on. e.g. ['12076:5010','12076:5020'].
          * `provisioning_state` (`str`) - The provisioning state of the route filter rule resource.
          * `route_filter_rule_type` (`str`) - The rule type of the rule.
    """
    tags: pulumi.Output[dict]
    """
    Resource tags.
    """
    type: pulumi.Output[str]
    """
    Resource type.
    """
    def __init__(__self__, resource_name, opts=None, id=None, location=None, name=None, resource_group_name=None, rules=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Route Filter Resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[str] name: The name of the route filter.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[list] rules: Collection of RouteFilterRules contained within a route filter.
        :param pulumi.Input[dict] tags: Resource tags.

        The **rules** object supports the following:

          * `etag` (`pulumi.Input[str]`) - A unique read-only string that changes whenever the resource is updated.
          * `location` (`pulumi.Input[str]`) - Resource location.
          * `name` (`pulumi.Input[str]`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
          * `properties` (`pulumi.Input[dict]`) - Properties of the route filter rule.
            * `access` (`pulumi.Input[str]`) - The access type of the rule.
            * `communities` (`pulumi.Input[list]`) - The collection for bgp community values to filter on. e.g. ['12076:5010','12076:5020'].
            * `provisioning_state` (`pulumi.Input[str]`) - The provisioning state of the route filter rule resource.
            * `route_filter_rule_type` (`pulumi.Input[str]`) - The rule type of the rule.
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
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['rules'] = rules
            __props__['tags'] = tags
            __props__['etag'] = None
            __props__['properties'] = None
            __props__['type'] = None
        super(RouteFilter, __self__).__init__(
            'azurerm:network/v20191101:RouteFilter',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing RouteFilter resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return RouteFilter(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
