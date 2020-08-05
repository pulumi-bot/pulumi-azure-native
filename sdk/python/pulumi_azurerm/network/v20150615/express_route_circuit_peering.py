# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class ExpressRouteCircuitPeering(pulumi.CustomResource):
    etag: pulumi.Output[str]
    """
    A unique read-only string that changes whenever the resource is updated.
    """
    name: pulumi.Output[str]
    """
    Gets name of the resource that is unique within a resource group. This name can be used to access the resource.
    """
    properties: pulumi.Output[dict]
    def __init__(__self__, resource_name, opts=None, azure_asn=None, circuit_name=None, etag=None, id=None, microsoft_peering_config=None, name=None, peer_asn=None, peering_type=None, primary_azure_port=None, primary_peer_address_prefix=None, provisioning_state=None, resource_group_name=None, secondary_azure_port=None, secondary_peer_address_prefix=None, shared_key=None, state=None, stats=None, vlan_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Peering in an ExpressRouteCircuit resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] azure_asn: The Azure ASN.
        :param pulumi.Input[str] circuit_name: The name of the express route circuit.
        :param pulumi.Input[str] etag: A unique read-only string that changes whenever the resource is updated.
        :param pulumi.Input[str] id: Resource Identifier.
        :param pulumi.Input[dict] microsoft_peering_config: The Microsoft peering configuration.
        :param pulumi.Input[str] name: The name of the peering.
        :param pulumi.Input[float] peer_asn: The peer ASN.
        :param pulumi.Input[str] peering_type: The PeeringType. Possible values are: 'AzurePublicPeering', 'AzurePrivatePeering', and 'MicrosoftPeering'.
        :param pulumi.Input[str] primary_azure_port: The primary port.
        :param pulumi.Input[str] primary_peer_address_prefix: The primary address prefix.
        :param pulumi.Input[str] provisioning_state: Gets the provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] secondary_azure_port: The secondary port.
        :param pulumi.Input[str] secondary_peer_address_prefix: The secondary address prefix.
        :param pulumi.Input[str] shared_key: The shared key.
        :param pulumi.Input[str] state: The state of peering. Possible values are: 'Disabled' and 'Enabled'
        :param pulumi.Input[dict] stats: Gets peering stats.
        :param pulumi.Input[float] vlan_id: The VLAN ID.

        The **microsoft_peering_config** object supports the following:

          * `advertised_public_prefixes` (`pulumi.Input[list]`) - The reference of AdvertisedPublicPrefixes.
          * `advertised_public_prefixes_state` (`pulumi.Input[str]`) - AdvertisedPublicPrefixState of the Peering resource. Possible values are 'NotConfigured', 'Configuring', 'Configured', and 'ValidationNeeded'.
          * `customer_asn` (`pulumi.Input[float]`) - The CustomerASN of the peering.
          * `routing_registry_name` (`pulumi.Input[str]`) - The RoutingRegistryName of the configuration.

        The **stats** object supports the following:

          * `bytes_in` (`pulumi.Input[float]`) - Gets BytesIn of the peering.
          * `bytes_out` (`pulumi.Input[float]`) - Gets BytesOut of the peering.
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

            __props__['azure_asn'] = azure_asn
            if circuit_name is None:
                raise TypeError("Missing required property 'circuit_name'")
            __props__['circuit_name'] = circuit_name
            __props__['etag'] = etag
            __props__['id'] = id
            __props__['microsoft_peering_config'] = microsoft_peering_config
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['peer_asn'] = peer_asn
            __props__['peering_type'] = peering_type
            __props__['primary_azure_port'] = primary_azure_port
            __props__['primary_peer_address_prefix'] = primary_peer_address_prefix
            __props__['provisioning_state'] = provisioning_state
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['secondary_azure_port'] = secondary_azure_port
            __props__['secondary_peer_address_prefix'] = secondary_peer_address_prefix
            __props__['shared_key'] = shared_key
            __props__['state'] = state
            __props__['stats'] = stats
            __props__['vlan_id'] = vlan_id
            __props__['properties'] = None
        super(ExpressRouteCircuitPeering, __self__).__init__(
            'azurerm:network/v20150615:ExpressRouteCircuitPeering',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing ExpressRouteCircuitPeering resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ExpressRouteCircuitPeering(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
