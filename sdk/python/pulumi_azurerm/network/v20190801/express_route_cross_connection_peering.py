# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class ExpressRouteCrossConnectionPeering(pulumi.CustomResource):
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
    Properties of the express route cross connection peering.
      * `azure_asn` (`float`) - The Azure ASN.
      * `gateway_manager_etag` (`str`) - The GatewayManager Etag.
      * `ipv6_peering_config` (`dict`) - The IPv6 peering configuration.
        * `microsoft_peering_config` (`dict`) - The Microsoft peering configuration.
          * `advertised_communities` (`list`) - The communities of bgp peering. Specified for microsoft peering.
          * `advertised_public_prefixes` (`list`) - The reference of AdvertisedPublicPrefixes.
          * `advertised_public_prefixes_state` (`str`) - The advertised public prefix state of the Peering resource.
          * `customer_asn` (`float`) - The CustomerASN of the peering.
          * `legacy_mode` (`float`) - The legacy mode of the peering.
          * `routing_registry_name` (`str`) - The RoutingRegistryName of the configuration.

        * `primary_peer_address_prefix` (`str`) - The primary address prefix.
        * `route_filter` (`dict`) - The reference of the RouteFilter resource.
          * `id` (`str`) - Resource ID.

        * `secondary_peer_address_prefix` (`str`) - The secondary address prefix.
        * `state` (`str`) - The state of peering.

      * `last_modified_by` (`str`) - Who was the last to modify the peering.
      * `microsoft_peering_config` (`dict`) - The Microsoft peering configuration.
      * `peer_asn` (`float`) - The peer ASN.
      * `peering_type` (`str`) - The peering type.
      * `primary_azure_port` (`str`) - The primary port.
      * `primary_peer_address_prefix` (`str`) - The primary address prefix.
      * `provisioning_state` (`str`) - The provisioning state of the express route cross connection peering resource.
      * `secondary_azure_port` (`str`) - The secondary port.
      * `secondary_peer_address_prefix` (`str`) - The secondary address prefix.
      * `shared_key` (`str`) - The shared key.
      * `state` (`str`) - The peering state.
      * `vlan_id` (`float`) - The VLAN ID.
    """
    def __init__(__self__, resource_name, opts=None, cross_connection_name=None, id=None, name=None, properties=None, resource_group_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Peering in an ExpressRoute Cross Connection resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cross_connection_name: The name of the ExpressRouteCrossConnection.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[str] name: The name of the peering.
        :param pulumi.Input[dict] properties: Properties of the express route cross connection peering.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.

        The **properties** object supports the following:

          * `gateway_manager_etag` (`pulumi.Input[str]`) - The GatewayManager Etag.
          * `ipv6_peering_config` (`pulumi.Input[dict]`) - The IPv6 peering configuration.
            * `microsoft_peering_config` (`pulumi.Input[dict]`) - The Microsoft peering configuration.
              * `advertised_communities` (`pulumi.Input[list]`) - The communities of bgp peering. Specified for microsoft peering.
              * `advertised_public_prefixes` (`pulumi.Input[list]`) - The reference of AdvertisedPublicPrefixes.
              * `advertised_public_prefixes_state` (`pulumi.Input[str]`) - The advertised public prefix state of the Peering resource.
              * `customer_asn` (`pulumi.Input[float]`) - The CustomerASN of the peering.
              * `legacy_mode` (`pulumi.Input[float]`) - The legacy mode of the peering.
              * `routing_registry_name` (`pulumi.Input[str]`) - The RoutingRegistryName of the configuration.

            * `primary_peer_address_prefix` (`pulumi.Input[str]`) - The primary address prefix.
            * `route_filter` (`pulumi.Input[dict]`) - The reference of the RouteFilter resource.
              * `id` (`pulumi.Input[str]`) - Resource ID.

            * `secondary_peer_address_prefix` (`pulumi.Input[str]`) - The secondary address prefix.
            * `state` (`pulumi.Input[str]`) - The state of peering.

          * `last_modified_by` (`pulumi.Input[str]`) - Who was the last to modify the peering.
          * `microsoft_peering_config` (`pulumi.Input[dict]`) - The Microsoft peering configuration.
          * `peer_asn` (`pulumi.Input[float]`) - The peer ASN.
          * `peering_type` (`pulumi.Input[str]`) - The peering type.
          * `primary_peer_address_prefix` (`pulumi.Input[str]`) - The primary address prefix.
          * `secondary_peer_address_prefix` (`pulumi.Input[str]`) - The secondary address prefix.
          * `shared_key` (`pulumi.Input[str]`) - The shared key.
          * `state` (`pulumi.Input[str]`) - The peering state.
          * `vlan_id` (`pulumi.Input[float]`) - The VLAN ID.
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

            if cross_connection_name is None:
                raise TypeError("Missing required property 'cross_connection_name'")
            __props__['cross_connection_name'] = cross_connection_name
            __props__['id'] = id
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['etag'] = None
        super(ExpressRouteCrossConnectionPeering, __self__).__init__(
            'azurerm:network/v20190801:ExpressRouteCrossConnectionPeering',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing ExpressRouteCrossConnectionPeering resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ExpressRouteCrossConnectionPeering(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
