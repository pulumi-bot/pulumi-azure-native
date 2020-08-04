# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class ExpressRouteCircuit(pulumi.CustomResource):
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
    Properties of ExpressRouteCircuit.
      * `allow_classic_operations` (`bool`) - Allow classic operations
      * `allow_global_reach` (`bool`) - Flag to enable Global Reach on the circuit.
      * `authorizations` (`list`) - The list of authorizations.
        * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
        * `id` (`str`) - Resource ID.
        * `name` (`str`) - Gets name of the resource that is unique within a resource group. This name can be used to access the resource.
        * `properties` (`dict`)
          * `authorization_key` (`str`) - The authorization key.
          * `authorization_use_status` (`str`) - AuthorizationUseStatus. Possible values are: 'Available' and 'InUse'.
          * `provisioning_state` (`str`) - Gets the provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.

      * `bandwidth_in_gbps` (`float`) - The bandwidth of the circuit when the circuit is provisioned on an ExpressRoutePort resource.
      * `circuit_provisioning_state` (`str`) - The CircuitProvisioningState state of the resource.
      * `express_route_port` (`dict`) - The reference to the ExpressRoutePort resource when the circuit is provisioned on an ExpressRoutePort resource.
        * `id` (`str`) - Resource ID.

      * `gateway_manager_etag` (`str`) - The GatewayManager Etag.
      * `peerings` (`list`) - The list of peerings.
        * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
        * `id` (`str`) - Resource ID.
        * `name` (`str`) - Gets name of the resource that is unique within a resource group. This name can be used to access the resource.
        * `properties` (`dict`)
          * `azure_asn` (`float`) - The Azure ASN.
          * `connections` (`list`) - The list of circuit connections associated with Azure Private Peering for this circuit.
            * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
            * `id` (`str`) - Resource ID.
            * `name` (`str`) - Gets name of the resource that is unique within a resource group. This name can be used to access the resource.
            * `properties` (`dict`)
              * `address_prefix` (`str`) - /29 IP address space to carve out Customer addresses for tunnels.
              * `authorization_key` (`str`) - The authorization key.
              * `circuit_connection_status` (`str`) - Express Route Circuit Connection State. Possible values are: 'Connected' and 'Disconnected'.
              * `express_route_circuit_peering` (`dict`) - Reference to Express Route Circuit Private Peering Resource of the circuit initiating connection.
              * `peer_express_route_circuit_peering` (`dict`) - Reference to Express Route Circuit Private Peering Resource of the peered circuit.
              * `provisioning_state` (`str`) - Provisioning state of the circuit connection resource. Possible values are: 'Succeeded', 'Updating', 'Deleting', and 'Failed'.

          * `express_route_connection` (`dict`) - The ExpressRoute connection.
            * `id` (`str`) - The ID of the ExpressRouteConnection.

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
                * `peerings` (`list`) - A collection of references to express route circuit peerings.
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

              * `tags` (`dict`) - Resource tags.
              * `type` (`str`) - Resource type.

            * `secondary_peer_address_prefix` (`str`) - The secondary address prefix.
            * `state` (`str`) - The state of peering. Possible values are: 'Disabled' and 'Enabled'

          * `last_modified_by` (`str`) - Gets whether the provider or the customer last modified the peering.
          * `microsoft_peering_config` (`dict`) - The Microsoft peering configuration.
          * `peer_asn` (`float`) - The peer ASN.
          * `peering_type` (`str`) - The peering type.
          * `primary_azure_port` (`str`) - The primary port.
          * `primary_peer_address_prefix` (`str`) - The primary address prefix.
          * `provisioning_state` (`str`) - Gets the provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
          * `route_filter` (`dict`) - The reference of the RouteFilter resource.
          * `secondary_azure_port` (`str`) - The secondary port.
          * `secondary_peer_address_prefix` (`str`) - The secondary address prefix.
          * `shared_key` (`str`) - The shared key.
          * `state` (`str`) - The peering state.
          * `stats` (`dict`) - Gets peering stats.
            * `primarybytes_in` (`float`) - Gets BytesIn of the peering.
            * `primarybytes_out` (`float`) - Gets BytesOut of the peering.
            * `secondarybytes_in` (`float`) - Gets BytesIn of the peering.
            * `secondarybytes_out` (`float`) - Gets BytesOut of the peering.

          * `vlan_id` (`float`) - The VLAN ID.

      * `provisioning_state` (`str`) - Gets the provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
      * `service_key` (`str`) - The ServiceKey.
      * `service_provider_notes` (`str`) - The ServiceProviderNotes.
      * `service_provider_properties` (`dict`) - The ServiceProviderProperties.
        * `bandwidth_in_mbps` (`float`) - The BandwidthInMbps.
        * `peering_location` (`str`) - The peering location.
        * `service_provider_name` (`str`) - The serviceProviderName.

      * `service_provider_provisioning_state` (`str`) - The ServiceProviderProvisioningState state of the resource. Possible values are 'NotProvisioned', 'Provisioning', 'Provisioned', and 'Deprovisioning'.
      * `stag` (`float`) - The identifier of the circuit traffic. Outer tag for QinQ encapsulation.
    """
    sku: pulumi.Output[dict]
    """
    The SKU.
      * `family` (`str`) - The family of the SKU. Possible values are: 'UnlimitedData' and 'MeteredData'.
      * `name` (`str`) - The name of the SKU.
      * `tier` (`str`) - The tier of the SKU. Possible values are 'Standard', 'Premium' or 'Basic'.
    """
    tags: pulumi.Output[dict]
    """
    Resource tags.
    """
    type: pulumi.Output[str]
    """
    Resource type.
    """
    def __init__(__self__, resource_name, opts=None, id=None, location=None, name=None, properties=None, resource_group_name=None, sku=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        ExpressRouteCircuit resource

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[str] name: The name of the circuit.
        :param pulumi.Input[dict] properties: Properties of ExpressRouteCircuit.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[dict] sku: The SKU.
        :param pulumi.Input[dict] tags: Resource tags.

        The **properties** object supports the following:

          * `allow_classic_operations` (`pulumi.Input[bool]`) - Allow classic operations
          * `allow_global_reach` (`pulumi.Input[bool]`) - Flag to enable Global Reach on the circuit.
          * `authorizations` (`pulumi.Input[list]`) - The list of authorizations.
            * `etag` (`pulumi.Input[str]`) - A unique read-only string that changes whenever the resource is updated.
            * `name` (`pulumi.Input[str]`) - Gets name of the resource that is unique within a resource group. This name can be used to access the resource.
            * `properties` (`pulumi.Input[dict]`)
              * `authorization_key` (`pulumi.Input[str]`) - The authorization key.
              * `authorization_use_status` (`pulumi.Input[str]`) - AuthorizationUseStatus. Possible values are: 'Available' and 'InUse'.
              * `provisioning_state` (`pulumi.Input[str]`) - Gets the provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.

          * `bandwidth_in_gbps` (`pulumi.Input[float]`) - The bandwidth of the circuit when the circuit is provisioned on an ExpressRoutePort resource.
          * `circuit_provisioning_state` (`pulumi.Input[str]`) - The CircuitProvisioningState state of the resource.
          * `express_route_port` (`pulumi.Input[dict]`) - The reference to the ExpressRoutePort resource when the circuit is provisioned on an ExpressRoutePort resource.
            * `id` (`pulumi.Input[str]`) - Resource ID.

          * `gateway_manager_etag` (`pulumi.Input[str]`) - The GatewayManager Etag.
          * `peerings` (`pulumi.Input[list]`) - The list of peerings.
            * `id` (`pulumi.Input[str]`) - Resource ID.
            * `name` (`pulumi.Input[str]`) - Gets name of the resource that is unique within a resource group. This name can be used to access the resource.
            * `properties` (`pulumi.Input[dict]`)
              * `azure_asn` (`pulumi.Input[float]`) - The Azure ASN.
              * `connections` (`pulumi.Input[list]`) - The list of circuit connections associated with Azure Private Peering for this circuit.
                * `id` (`pulumi.Input[str]`) - Resource ID.
                * `name` (`pulumi.Input[str]`) - Gets name of the resource that is unique within a resource group. This name can be used to access the resource.
                * `properties` (`pulumi.Input[dict]`)
                  * `address_prefix` (`pulumi.Input[str]`) - /29 IP address space to carve out Customer addresses for tunnels.
                  * `authorization_key` (`pulumi.Input[str]`) - The authorization key.
                  * `express_route_circuit_peering` (`pulumi.Input[dict]`) - Reference to Express Route Circuit Private Peering Resource of the circuit initiating connection.
                  * `peer_express_route_circuit_peering` (`pulumi.Input[dict]`) - Reference to Express Route Circuit Private Peering Resource of the peered circuit.

              * `express_route_connection` (`pulumi.Input[dict]`) - The ExpressRoute connection.
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
                    * `peerings` (`pulumi.Input[list]`) - A collection of references to express route circuit peerings.
                    * `rules` (`pulumi.Input[list]`) - Collection of RouteFilterRules contained within a route filter.
                      * `etag` (`pulumi.Input[str]`) - A unique read-only string that changes whenever the resource is updated.
                      * `location` (`pulumi.Input[str]`) - Resource location.
                      * `name` (`pulumi.Input[str]`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
                      * `properties` (`pulumi.Input[dict]`) - Route Filter Rule Resource
                        * `access` (`pulumi.Input[str]`) - The access type of the rule. Valid values are: 'Allow', 'Deny'
                        * `communities` (`pulumi.Input[list]`) - The collection for bgp community values to filter on. e.g. ['12076:5010','12076:5020']
                        * `provisioning_state` (`pulumi.Input[str]`) - The provisioning state of the resource. Possible values are: 'Updating', 'Deleting', 'Succeeded' and 'Failed'.
                        * `route_filter_rule_type` (`pulumi.Input[str]`) - The rule type of the rule. Valid value is: 'Community'

                  * `tags` (`pulumi.Input[dict]`) - Resource tags.

                * `secondary_peer_address_prefix` (`pulumi.Input[str]`) - The secondary address prefix.
                * `state` (`pulumi.Input[str]`) - The state of peering. Possible values are: 'Disabled' and 'Enabled'

              * `last_modified_by` (`pulumi.Input[str]`) - Gets whether the provider or the customer last modified the peering.
              * `microsoft_peering_config` (`pulumi.Input[dict]`) - The Microsoft peering configuration.
              * `peer_asn` (`pulumi.Input[float]`) - The peer ASN.
              * `peering_type` (`pulumi.Input[str]`) - The peering type.
              * `primary_azure_port` (`pulumi.Input[str]`) - The primary port.
              * `primary_peer_address_prefix` (`pulumi.Input[str]`) - The primary address prefix.
              * `provisioning_state` (`pulumi.Input[str]`) - Gets the provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
              * `route_filter` (`pulumi.Input[dict]`) - The reference of the RouteFilter resource.
              * `secondary_azure_port` (`pulumi.Input[str]`) - The secondary port.
              * `secondary_peer_address_prefix` (`pulumi.Input[str]`) - The secondary address prefix.
              * `shared_key` (`pulumi.Input[str]`) - The shared key.
              * `state` (`pulumi.Input[str]`) - The peering state.
              * `stats` (`pulumi.Input[dict]`) - Gets peering stats.
                * `primarybytes_in` (`pulumi.Input[float]`) - Gets BytesIn of the peering.
                * `primarybytes_out` (`pulumi.Input[float]`) - Gets BytesOut of the peering.
                * `secondarybytes_in` (`pulumi.Input[float]`) - Gets BytesIn of the peering.
                * `secondarybytes_out` (`pulumi.Input[float]`) - Gets BytesOut of the peering.

              * `vlan_id` (`pulumi.Input[float]`) - The VLAN ID.

          * `provisioning_state` (`pulumi.Input[str]`) - Gets the provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
          * `service_key` (`pulumi.Input[str]`) - The ServiceKey.
          * `service_provider_notes` (`pulumi.Input[str]`) - The ServiceProviderNotes.
          * `service_provider_properties` (`pulumi.Input[dict]`) - The ServiceProviderProperties.
            * `bandwidth_in_mbps` (`pulumi.Input[float]`) - The BandwidthInMbps.
            * `peering_location` (`pulumi.Input[str]`) - The peering location.
            * `service_provider_name` (`pulumi.Input[str]`) - The serviceProviderName.

          * `service_provider_provisioning_state` (`pulumi.Input[str]`) - The ServiceProviderProvisioningState state of the resource. Possible values are 'NotProvisioned', 'Provisioning', 'Provisioned', and 'Deprovisioning'.

        The **sku** object supports the following:

          * `family` (`pulumi.Input[str]`) - The family of the SKU. Possible values are: 'UnlimitedData' and 'MeteredData'.
          * `name` (`pulumi.Input[str]`) - The name of the SKU.
          * `tier` (`pulumi.Input[str]`) - The tier of the SKU. Possible values are 'Standard', 'Premium' or 'Basic'.
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
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['etag'] = None
            __props__['type'] = None
        super(ExpressRouteCircuit, __self__).__init__(
            'azurerm:network/v20181001:ExpressRouteCircuit',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing ExpressRouteCircuit resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ExpressRouteCircuit(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
