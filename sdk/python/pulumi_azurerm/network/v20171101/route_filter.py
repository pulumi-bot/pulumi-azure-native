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
    Route Filter Resource
      * `peerings` (`list`) - A collection of references to express route circuit peerings.
        * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
        * `id` (`str`) - Resource ID.
        * `name` (`str`) - Gets name of the resource that is unique within a resource group. This name can be used to access the resource.
        * `properties` (`dict`)
          * `azure_asn` (`float`) - The Azure ASN.
          * `gateway_manager_etag` (`str`) - The GatewayManager Etag.
          * `ipv6_peering_config` (`dict`) - The IPv6 peering configuration.
            * `microsoft_peering_config` (`dict`) - The Microsoft peering configuration.
              * `advertised_communities` (`list`) - The communities of bgp peering. Specified for microsoft peering
              * `advertised_public_prefixes` (`list`) - The reference of AdvertisedPublicPrefixes.
              * `advertised_public_prefixes_state` (`str`) - AdvertisedPublicPrefixState of the Peering resource. Possible values are 'NotConfigured', 'Configuring', 'Configured', and 'ValidationNeeded'.
              * `customer_asn` (`float`) - The CustomerASN of the peering.
              * `legacy_mode` (`float`) - The legacy mode of the peering.
              * `routing_registry_name` (`str`) - The RoutingRegistryName of the configuration.

            * `primary_peer_address_prefix` (`str`) - The primary address prefix.
            * `route_filter` (`dict`) - The reference of the RouteFilter resource.
              * `etag` (`str`) - Gets a unique read-only string that changes whenever the resource is updated.
              * `id` (`str`) - Resource ID.
              * `location` (`str`) - Resource location.
              * `name` (`str`) - Resource name.
              * `properties` (`dict`) - Route Filter Resource
              * `tags` (`dict`) - Resource tags.
              * `type` (`str`) - Resource type.

            * `secondary_peer_address_prefix` (`str`) - The secondary address prefix.
            * `state` (`str`) - The state of peering. Possible values are: 'Disabled' and 'Enabled'

          * `last_modified_by` (`str`) - Gets whether the provider or the customer last modified the peering.
          * `microsoft_peering_config` (`dict`) - The Microsoft peering configuration.
          * `peer_asn` (`float`) - The peer ASN.
          * `peering_type` (`str`) - The PeeringType. Possible values are: 'AzurePublicPeering', 'AzurePrivatePeering', and 'MicrosoftPeering'.
          * `primary_azure_port` (`str`) - The primary port.
          * `primary_peer_address_prefix` (`str`) - The primary address prefix.
          * `provisioning_state` (`str`) - Gets the provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
          * `route_filter` (`dict`) - The reference of the RouteFilter resource.
          * `secondary_azure_port` (`str`) - The secondary port.
          * `secondary_peer_address_prefix` (`str`) - The secondary address prefix.
          * `shared_key` (`str`) - The shared key.
          * `state` (`str`) - The state of peering. Possible values are: 'Disabled' and 'Enabled'
          * `stats` (`dict`) - Gets peering stats.
            * `primarybytes_in` (`float`) - Gets BytesIn of the peering.
            * `primarybytes_out` (`float`) - Gets BytesOut of the peering.
            * `secondarybytes_in` (`float`) - Gets BytesIn of the peering.
            * `secondarybytes_out` (`float`) - Gets BytesOut of the peering.

          * `vlan_id` (`float`) - The VLAN ID.

      * `provisioning_state` (`str`) - The provisioning state of the resource. Possible values are: 'Updating', 'Deleting', 'Succeeded' and 'Failed'.
      * `rules` (`list`) - Collection of RouteFilterRules contained within a route filter.
        * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
        * `id` (`str`) - Resource ID.
        * `location` (`str`) - Resource location.
        * `name` (`str`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
        * `properties` (`dict`) - Route Filter Rule Resource
          * `access` (`str`) - The access type of the rule. Valid values are: 'Allow', 'Deny'
          * `communities` (`list`) - The collection for bgp community values to filter on. e.g. ['12076:5010','12076:5020']
          * `provisioning_state` (`str`) - The provisioning state of the resource. Possible values are: 'Updating', 'Deleting', 'Succeeded' and 'Failed'.
          * `route_filter_rule_type` (`str`) - The rule type of the rule. Valid value is: 'Community'
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
        Route Filter Resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[str] name: The name of the route filter.
        :param pulumi.Input[dict] properties: Route Filter Resource
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[dict] tags: Resource tags.

        The **properties** object supports the following:

          * `peerings` (`pulumi.Input[list]`) - A collection of references to express route circuit peerings.
            * `id` (`pulumi.Input[str]`) - Resource ID.
            * `name` (`pulumi.Input[str]`) - Gets name of the resource that is unique within a resource group. This name can be used to access the resource.
            * `properties` (`pulumi.Input[dict]`)
              * `azure_asn` (`pulumi.Input[float]`) - The Azure ASN.
              * `gateway_manager_etag` (`pulumi.Input[str]`) - The GatewayManager Etag.
              * `ipv6_peering_config` (`pulumi.Input[dict]`) - The IPv6 peering configuration.
                * `microsoft_peering_config` (`pulumi.Input[dict]`) - The Microsoft peering configuration.
                  * `advertised_communities` (`pulumi.Input[list]`) - The communities of bgp peering. Specified for microsoft peering
                  * `advertised_public_prefixes` (`pulumi.Input[list]`) - The reference of AdvertisedPublicPrefixes.
                  * `advertised_public_prefixes_state` (`pulumi.Input[str]`) - AdvertisedPublicPrefixState of the Peering resource. Possible values are 'NotConfigured', 'Configuring', 'Configured', and 'ValidationNeeded'.
                  * `customer_asn` (`pulumi.Input[float]`) - The CustomerASN of the peering.
                  * `legacy_mode` (`pulumi.Input[float]`) - The legacy mode of the peering.
                  * `routing_registry_name` (`pulumi.Input[str]`) - The RoutingRegistryName of the configuration.

                * `primary_peer_address_prefix` (`pulumi.Input[str]`) - The primary address prefix.
                * `route_filter` (`pulumi.Input[dict]`) - The reference of the RouteFilter resource.
                  * `id` (`pulumi.Input[str]`) - Resource ID.
                  * `location` (`pulumi.Input[str]`) - Resource location.
                  * `properties` (`pulumi.Input[dict]`) - Route Filter Resource
                  * `tags` (`pulumi.Input[dict]`) - Resource tags.

                * `secondary_peer_address_prefix` (`pulumi.Input[str]`) - The secondary address prefix.
                * `state` (`pulumi.Input[str]`) - The state of peering. Possible values are: 'Disabled' and 'Enabled'

              * `last_modified_by` (`pulumi.Input[str]`) - Gets whether the provider or the customer last modified the peering.
              * `microsoft_peering_config` (`pulumi.Input[dict]`) - The Microsoft peering configuration.
              * `peer_asn` (`pulumi.Input[float]`) - The peer ASN.
              * `peering_type` (`pulumi.Input[str]`) - The PeeringType. Possible values are: 'AzurePublicPeering', 'AzurePrivatePeering', and 'MicrosoftPeering'.
              * `primary_azure_port` (`pulumi.Input[str]`) - The primary port.
              * `primary_peer_address_prefix` (`pulumi.Input[str]`) - The primary address prefix.
              * `provisioning_state` (`pulumi.Input[str]`) - Gets the provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
              * `route_filter` (`pulumi.Input[dict]`) - The reference of the RouteFilter resource.
              * `secondary_azure_port` (`pulumi.Input[str]`) - The secondary port.
              * `secondary_peer_address_prefix` (`pulumi.Input[str]`) - The secondary address prefix.
              * `shared_key` (`pulumi.Input[str]`) - The shared key.
              * `state` (`pulumi.Input[str]`) - The state of peering. Possible values are: 'Disabled' and 'Enabled'
              * `stats` (`pulumi.Input[dict]`) - Gets peering stats.
                * `primarybytes_in` (`pulumi.Input[float]`) - Gets BytesIn of the peering.
                * `primarybytes_out` (`pulumi.Input[float]`) - Gets BytesOut of the peering.
                * `secondarybytes_in` (`pulumi.Input[float]`) - Gets BytesIn of the peering.
                * `secondarybytes_out` (`pulumi.Input[float]`) - Gets BytesOut of the peering.

              * `vlan_id` (`pulumi.Input[float]`) - The VLAN ID.

          * `rules` (`pulumi.Input[list]`) - Collection of RouteFilterRules contained within a route filter.
            * `etag` (`pulumi.Input[str]`) - A unique read-only string that changes whenever the resource is updated.
            * `location` (`pulumi.Input[str]`) - Resource location.
            * `name` (`pulumi.Input[str]`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
            * `properties` (`pulumi.Input[dict]`) - Route Filter Rule Resource
              * `access` (`pulumi.Input[str]`) - The access type of the rule. Valid values are: 'Allow', 'Deny'
              * `communities` (`pulumi.Input[list]`) - The collection for bgp community values to filter on. e.g. ['12076:5010','12076:5020']
              * `provisioning_state` (`pulumi.Input[str]`) - The provisioning state of the resource. Possible values are: 'Updating', 'Deleting', 'Succeeded' and 'Failed'.
              * `route_filter_rule_type` (`pulumi.Input[str]`) - The rule type of the rule. Valid value is: 'Community'
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
        super(RouteFilter, __self__).__init__(
            'azurerm:network/v20171101:RouteFilter',
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
