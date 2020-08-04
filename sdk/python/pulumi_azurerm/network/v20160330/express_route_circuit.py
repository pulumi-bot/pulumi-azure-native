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
    Gets a unique read-only string that changes whenever the resource is updated
    """
    location: pulumi.Output[str]
    """
    Resource location
    """
    name: pulumi.Output[str]
    """
    Resource name
    """
    properties: pulumi.Output[dict]
    """
    Properties of ExpressRouteCircuit
      * `allow_classic_operations` (`bool`) - allow classic operations
      * `authorizations` (`list`) - Gets or sets list of authorizations
        * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated
        * `id` (`str`) - Resource Id
        * `name` (`str`) - Gets name of the resource that is unique within a resource group. This name can be used to access the resource
        * `properties` (`dict`)
          * `authorization_key` (`str`) - Gets or sets the authorization key
          * `authorization_use_status` (`str`) - Gets or sets AuthorizationUseStatus
          * `provisioning_state` (`str`) - Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed

      * `circuit_provisioning_state` (`str`) - Gets or sets CircuitProvisioningState state of the resource 
      * `peerings` (`list`) - Gets or sets list of peerings
        * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated
        * `id` (`str`) - Resource Id
        * `name` (`str`) - Gets name of the resource that is unique within a resource group. This name can be used to access the resource
        * `properties` (`dict`)
          * `azure_asn` (`float`) - Gets or sets the azure ASN
          * `microsoft_peering_config` (`dict`) - Gets or sets the Microsoft peering config
            * `advertised_public_prefixes` (`list`) - Gets or sets the reference of AdvertisedPublicPrefixes
            * `advertised_public_prefixes_state` (`str`) - Gets or sets AdvertisedPublicPrefixState of the Peering resource 
            * `customer_asn` (`float`) - Gets or Sets CustomerAsn of the peering.
            * `routing_registry_name` (`str`) - Gets or Sets RoutingRegistryName of the config.

          * `peer_asn` (`float`) - Gets or sets the peer ASN
          * `peering_type` (`str`) - Gets or sets PeeringType
          * `primary_azure_port` (`str`) - Gets or sets the primary port
          * `primary_peer_address_prefix` (`str`) - Gets or sets the primary address prefix
          * `provisioning_state` (`str`) - Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed
          * `secondary_azure_port` (`str`) - Gets or sets the secondary port
          * `secondary_peer_address_prefix` (`str`) - Gets or sets the secondary address prefix
          * `shared_key` (`str`) - Gets or sets the shared key
          * `state` (`str`) - Gets or sets state of Peering
          * `stats` (`dict`) - Gets or peering stats
            * `primarybytes_in` (`float`) - Gets BytesIn of the peering.
            * `primarybytes_out` (`float`) - Gets BytesOut of the peering.
            * `secondarybytes_in` (`float`) - Gets BytesIn of the peering.
            * `secondarybytes_out` (`float`) - Gets BytesOut of the peering.

          * `vlan_id` (`float`) - Gets or sets the vlan id

      * `provisioning_state` (`str`) - Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed
      * `service_key` (`str`) - Gets or sets ServiceKey
      * `service_provider_notes` (`str`) - Gets or sets ServiceProviderNotes
      * `service_provider_properties` (`dict`) - Gets or sets ServiceProviderProperties
        * `bandwidth_in_mbps` (`float`) - Gets or sets BandwidthInMbps.
        * `peering_location` (`str`) - Gets or sets peering location.
        * `service_provider_name` (`str`) - Gets or sets serviceProviderName.

      * `service_provider_provisioning_state` (`str`) - Gets or sets ServiceProviderProvisioningState state of the resource 
    """
    sku: pulumi.Output[dict]
    """
    Gets or sets sku
      * `family` (`str`) - Gets or sets family of the sku.
      * `name` (`str`) - Gets or sets name of the sku.
      * `tier` (`str`) - Gets or sets tier of the sku.
    """
    tags: pulumi.Output[dict]
    """
    Resource tags
    """
    type: pulumi.Output[str]
    """
    Resource type
    """
    def __init__(__self__, resource_name, opts=None, etag=None, id=None, location=None, name=None, properties=None, resource_group_name=None, sku=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        ExpressRouteCircuit resource

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] etag: Gets a unique read-only string that changes whenever the resource is updated
        :param pulumi.Input[str] id: Resource Id
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[str] name: The name of the circuit.
        :param pulumi.Input[dict] properties: Properties of ExpressRouteCircuit
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[dict] sku: Gets or sets sku
        :param pulumi.Input[dict] tags: Resource tags

        The **properties** object supports the following:

          * `allow_classic_operations` (`pulumi.Input[bool]`) - allow classic operations
          * `authorizations` (`pulumi.Input[list]`) - Gets or sets list of authorizations
            * `etag` (`pulumi.Input[str]`) - A unique read-only string that changes whenever the resource is updated
            * `name` (`pulumi.Input[str]`) - Gets name of the resource that is unique within a resource group. This name can be used to access the resource
            * `properties` (`pulumi.Input[dict]`)
              * `authorization_key` (`pulumi.Input[str]`) - Gets or sets the authorization key
              * `authorization_use_status` (`pulumi.Input[str]`) - Gets or sets AuthorizationUseStatus
              * `provisioning_state` (`pulumi.Input[str]`) - Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed

          * `circuit_provisioning_state` (`pulumi.Input[str]`) - Gets or sets CircuitProvisioningState state of the resource 
          * `peerings` (`pulumi.Input[list]`) - Gets or sets list of peerings
            * `etag` (`pulumi.Input[str]`) - A unique read-only string that changes whenever the resource is updated
            * `name` (`pulumi.Input[str]`) - Gets name of the resource that is unique within a resource group. This name can be used to access the resource
            * `properties` (`pulumi.Input[dict]`)
              * `azure_asn` (`pulumi.Input[float]`) - Gets or sets the azure ASN
              * `microsoft_peering_config` (`pulumi.Input[dict]`) - Gets or sets the Microsoft peering config
                * `advertised_public_prefixes` (`pulumi.Input[list]`) - Gets or sets the reference of AdvertisedPublicPrefixes
                * `advertised_public_prefixes_state` (`pulumi.Input[str]`) - Gets or sets AdvertisedPublicPrefixState of the Peering resource 
                * `customer_asn` (`pulumi.Input[float]`) - Gets or Sets CustomerAsn of the peering.
                * `routing_registry_name` (`pulumi.Input[str]`) - Gets or Sets RoutingRegistryName of the config.

              * `peer_asn` (`pulumi.Input[float]`) - Gets or sets the peer ASN
              * `peering_type` (`pulumi.Input[str]`) - Gets or sets PeeringType
              * `primary_azure_port` (`pulumi.Input[str]`) - Gets or sets the primary port
              * `primary_peer_address_prefix` (`pulumi.Input[str]`) - Gets or sets the primary address prefix
              * `provisioning_state` (`pulumi.Input[str]`) - Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed
              * `secondary_azure_port` (`pulumi.Input[str]`) - Gets or sets the secondary port
              * `secondary_peer_address_prefix` (`pulumi.Input[str]`) - Gets or sets the secondary address prefix
              * `shared_key` (`pulumi.Input[str]`) - Gets or sets the shared key
              * `state` (`pulumi.Input[str]`) - Gets or sets state of Peering
              * `stats` (`pulumi.Input[dict]`) - Gets or peering stats
                * `primarybytes_in` (`pulumi.Input[float]`) - Gets BytesIn of the peering.
                * `primarybytes_out` (`pulumi.Input[float]`) - Gets BytesOut of the peering.
                * `secondarybytes_in` (`pulumi.Input[float]`) - Gets BytesIn of the peering.
                * `secondarybytes_out` (`pulumi.Input[float]`) - Gets BytesOut of the peering.

              * `vlan_id` (`pulumi.Input[float]`) - Gets or sets the vlan id

          * `provisioning_state` (`pulumi.Input[str]`) - Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed
          * `service_key` (`pulumi.Input[str]`) - Gets or sets ServiceKey
          * `service_provider_notes` (`pulumi.Input[str]`) - Gets or sets ServiceProviderNotes
          * `service_provider_properties` (`pulumi.Input[dict]`) - Gets or sets ServiceProviderProperties
            * `bandwidth_in_mbps` (`pulumi.Input[float]`) - Gets or sets BandwidthInMbps.
            * `peering_location` (`pulumi.Input[str]`) - Gets or sets peering location.
            * `service_provider_name` (`pulumi.Input[str]`) - Gets or sets serviceProviderName.

          * `service_provider_provisioning_state` (`pulumi.Input[str]`) - Gets or sets ServiceProviderProvisioningState state of the resource 

        The **sku** object supports the following:

          * `family` (`pulumi.Input[str]`) - Gets or sets family of the sku.
          * `name` (`pulumi.Input[str]`) - Gets or sets name of the sku.
          * `tier` (`pulumi.Input[str]`) - Gets or sets tier of the sku.
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

            __props__['etag'] = etag
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
            __props__['type'] = None
        super(ExpressRouteCircuit, __self__).__init__(
            'azurerm:network/v20160330:ExpressRouteCircuit',
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
