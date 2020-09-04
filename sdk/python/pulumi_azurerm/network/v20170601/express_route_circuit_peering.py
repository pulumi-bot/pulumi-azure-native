# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['ExpressRouteCircuitPeering']


class ExpressRouteCircuitPeering(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 azure_asn: Optional[pulumi.Input[float]] = None,
                 circuit_name: Optional[pulumi.Input[str]] = None,
                 gateway_manager_etag: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 ipv6_peering_config: Optional[pulumi.Input[pulumi.InputType['Ipv6ExpressRouteCircuitPeeringConfigArgs']]] = None,
                 last_modified_by: Optional[pulumi.Input[str]] = None,
                 microsoft_peering_config: Optional[pulumi.Input[pulumi.InputType['ExpressRouteCircuitPeeringConfigArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 peer_asn: Optional[pulumi.Input[float]] = None,
                 peering_name: Optional[pulumi.Input[str]] = None,
                 peering_type: Optional[pulumi.Input[str]] = None,
                 primary_azure_port: Optional[pulumi.Input[str]] = None,
                 primary_peer_address_prefix: Optional[pulumi.Input[str]] = None,
                 provisioning_state: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 route_filter: Optional[pulumi.Input[pulumi.InputType['RouteFilterArgs']]] = None,
                 secondary_azure_port: Optional[pulumi.Input[str]] = None,
                 secondary_peer_address_prefix: Optional[pulumi.Input[str]] = None,
                 shared_key: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 stats: Optional[pulumi.Input[pulumi.InputType['ExpressRouteCircuitStatsArgs']]] = None,
                 vlan_id: Optional[pulumi.Input[float]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Peering in an ExpressRouteCircuit resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] azure_asn: The Azure ASN.
        :param pulumi.Input[str] circuit_name: The name of the express route circuit.
        :param pulumi.Input[str] gateway_manager_etag: The GatewayManager Etag.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[pulumi.InputType['Ipv6ExpressRouteCircuitPeeringConfigArgs']] ipv6_peering_config: The IPv6 peering configuration.
        :param pulumi.Input[str] last_modified_by: Gets whether the provider or the customer last modified the peering.
        :param pulumi.Input[pulumi.InputType['ExpressRouteCircuitPeeringConfigArgs']] microsoft_peering_config: The Microsoft peering configuration.
        :param pulumi.Input[str] name: Gets name of the resource that is unique within a resource group. This name can be used to access the resource.
        :param pulumi.Input[float] peer_asn: The peer ASN.
        :param pulumi.Input[str] peering_name: The name of the peering.
        :param pulumi.Input[str] peering_type: The PeeringType. Possible values are: 'AzurePublicPeering', 'AzurePrivatePeering', and 'MicrosoftPeering'.
        :param pulumi.Input[str] primary_azure_port: The primary port.
        :param pulumi.Input[str] primary_peer_address_prefix: The primary address prefix.
        :param pulumi.Input[str] provisioning_state: Gets the provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[pulumi.InputType['RouteFilterArgs']] route_filter: The reference of the RouteFilter resource.
        :param pulumi.Input[str] secondary_azure_port: The secondary port.
        :param pulumi.Input[str] secondary_peer_address_prefix: The secondary address prefix.
        :param pulumi.Input[str] shared_key: The shared key.
        :param pulumi.Input[str] state: The state of peering. Possible values are: 'Disabled' and 'Enabled'
        :param pulumi.Input[pulumi.InputType['ExpressRouteCircuitStatsArgs']] stats: Gets peering stats.
        :param pulumi.Input[float] vlan_id: The VLAN ID.
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
            __props__['gateway_manager_etag'] = gateway_manager_etag
            __props__['id'] = id
            __props__['ipv6_peering_config'] = ipv6_peering_config
            __props__['last_modified_by'] = last_modified_by
            __props__['microsoft_peering_config'] = microsoft_peering_config
            __props__['name'] = name
            __props__['peer_asn'] = peer_asn
            if peering_name is None:
                raise TypeError("Missing required property 'peering_name'")
            __props__['peering_name'] = peering_name
            __props__['peering_type'] = peering_type
            __props__['primary_azure_port'] = primary_azure_port
            __props__['primary_peer_address_prefix'] = primary_peer_address_prefix
            __props__['provisioning_state'] = provisioning_state
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['route_filter'] = route_filter
            __props__['secondary_azure_port'] = secondary_azure_port
            __props__['secondary_peer_address_prefix'] = secondary_peer_address_prefix
            __props__['shared_key'] = shared_key
            __props__['state'] = state
            __props__['stats'] = stats
            __props__['vlan_id'] = vlan_id
            __props__['etag'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:network/latest:ExpressRouteCircuitPeering"), pulumi.Alias(type_="azurerm:network/v20150501preview:ExpressRouteCircuitPeering"), pulumi.Alias(type_="azurerm:network/v20150615:ExpressRouteCircuitPeering"), pulumi.Alias(type_="azurerm:network/v20160330:ExpressRouteCircuitPeering"), pulumi.Alias(type_="azurerm:network/v20160601:ExpressRouteCircuitPeering"), pulumi.Alias(type_="azurerm:network/v20160901:ExpressRouteCircuitPeering"), pulumi.Alias(type_="azurerm:network/v20161201:ExpressRouteCircuitPeering"), pulumi.Alias(type_="azurerm:network/v20170301:ExpressRouteCircuitPeering"), pulumi.Alias(type_="azurerm:network/v20170801:ExpressRouteCircuitPeering"), pulumi.Alias(type_="azurerm:network/v20170901:ExpressRouteCircuitPeering"), pulumi.Alias(type_="azurerm:network/v20171001:ExpressRouteCircuitPeering"), pulumi.Alias(type_="azurerm:network/v20171101:ExpressRouteCircuitPeering"), pulumi.Alias(type_="azurerm:network/v20180101:ExpressRouteCircuitPeering"), pulumi.Alias(type_="azurerm:network/v20180201:ExpressRouteCircuitPeering"), pulumi.Alias(type_="azurerm:network/v20180401:ExpressRouteCircuitPeering"), pulumi.Alias(type_="azurerm:network/v20180601:ExpressRouteCircuitPeering"), pulumi.Alias(type_="azurerm:network/v20180701:ExpressRouteCircuitPeering"), pulumi.Alias(type_="azurerm:network/v20180801:ExpressRouteCircuitPeering"), pulumi.Alias(type_="azurerm:network/v20181001:ExpressRouteCircuitPeering"), pulumi.Alias(type_="azurerm:network/v20181101:ExpressRouteCircuitPeering"), pulumi.Alias(type_="azurerm:network/v20181201:ExpressRouteCircuitPeering"), pulumi.Alias(type_="azurerm:network/v20190201:ExpressRouteCircuitPeering"), pulumi.Alias(type_="azurerm:network/v20190401:ExpressRouteCircuitPeering"), pulumi.Alias(type_="azurerm:network/v20190601:ExpressRouteCircuitPeering"), pulumi.Alias(type_="azurerm:network/v20190701:ExpressRouteCircuitPeering"), pulumi.Alias(type_="azurerm:network/v20190801:ExpressRouteCircuitPeering"), pulumi.Alias(type_="azurerm:network/v20190901:ExpressRouteCircuitPeering"), pulumi.Alias(type_="azurerm:network/v20191101:ExpressRouteCircuitPeering"), pulumi.Alias(type_="azurerm:network/v20191201:ExpressRouteCircuitPeering"), pulumi.Alias(type_="azurerm:network/v20200301:ExpressRouteCircuitPeering"), pulumi.Alias(type_="azurerm:network/v20200401:ExpressRouteCircuitPeering"), pulumi.Alias(type_="azurerm:network/v20200501:ExpressRouteCircuitPeering"), pulumi.Alias(type_="azurerm:network/v20200601:ExpressRouteCircuitPeering")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(ExpressRouteCircuitPeering, __self__).__init__(
            'azurerm:network/v20170601:ExpressRouteCircuitPeering',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ExpressRouteCircuitPeering':
        """
        Get an existing ExpressRouteCircuitPeering resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ExpressRouteCircuitPeering(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureASN")
    def azure_asn(self) -> pulumi.Output[Optional[float]]:
        """
        The Azure ASN.
        """
        return pulumi.get(self, "azure_asn")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        A unique read-only string that changes whenever the resource is updated.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="gatewayManagerEtag")
    def gateway_manager_etag(self) -> pulumi.Output[Optional[str]]:
        """
        The GatewayManager Etag.
        """
        return pulumi.get(self, "gateway_manager_etag")

    @property
    @pulumi.getter(name="ipv6PeeringConfig")
    def ipv6_peering_config(self) -> pulumi.Output[Optional['outputs.Ipv6ExpressRouteCircuitPeeringConfigResponse']]:
        """
        The IPv6 peering configuration.
        """
        return pulumi.get(self, "ipv6_peering_config")

    @property
    @pulumi.getter(name="lastModifiedBy")
    def last_modified_by(self) -> pulumi.Output[Optional[str]]:
        """
        Gets whether the provider or the customer last modified the peering.
        """
        return pulumi.get(self, "last_modified_by")

    @property
    @pulumi.getter(name="microsoftPeeringConfig")
    def microsoft_peering_config(self) -> pulumi.Output[Optional['outputs.ExpressRouteCircuitPeeringConfigResponse']]:
        """
        The Microsoft peering configuration.
        """
        return pulumi.get(self, "microsoft_peering_config")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        """
        Gets name of the resource that is unique within a resource group. This name can be used to access the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="peerASN")
    def peer_asn(self) -> pulumi.Output[Optional[float]]:
        """
        The peer ASN.
        """
        return pulumi.get(self, "peer_asn")

    @property
    @pulumi.getter(name="peeringType")
    def peering_type(self) -> pulumi.Output[Optional[str]]:
        """
        The PeeringType. Possible values are: 'AzurePublicPeering', 'AzurePrivatePeering', and 'MicrosoftPeering'.
        """
        return pulumi.get(self, "peering_type")

    @property
    @pulumi.getter(name="primaryAzurePort")
    def primary_azure_port(self) -> pulumi.Output[Optional[str]]:
        """
        The primary port.
        """
        return pulumi.get(self, "primary_azure_port")

    @property
    @pulumi.getter(name="primaryPeerAddressPrefix")
    def primary_peer_address_prefix(self) -> pulumi.Output[Optional[str]]:
        """
        The primary address prefix.
        """
        return pulumi.get(self, "primary_peer_address_prefix")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[Optional[str]]:
        """
        Gets the provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="routeFilter")
    def route_filter(self) -> pulumi.Output[Optional['outputs.RouteFilterResponse']]:
        """
        The reference of the RouteFilter resource.
        """
        return pulumi.get(self, "route_filter")

    @property
    @pulumi.getter(name="secondaryAzurePort")
    def secondary_azure_port(self) -> pulumi.Output[Optional[str]]:
        """
        The secondary port.
        """
        return pulumi.get(self, "secondary_azure_port")

    @property
    @pulumi.getter(name="secondaryPeerAddressPrefix")
    def secondary_peer_address_prefix(self) -> pulumi.Output[Optional[str]]:
        """
        The secondary address prefix.
        """
        return pulumi.get(self, "secondary_peer_address_prefix")

    @property
    @pulumi.getter(name="sharedKey")
    def shared_key(self) -> pulumi.Output[Optional[str]]:
        """
        The shared key.
        """
        return pulumi.get(self, "shared_key")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[Optional[str]]:
        """
        The state of peering. Possible values are: 'Disabled' and 'Enabled'
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def stats(self) -> pulumi.Output[Optional['outputs.ExpressRouteCircuitStatsResponse']]:
        """
        Gets peering stats.
        """
        return pulumi.get(self, "stats")

    @property
    @pulumi.getter(name="vlanId")
    def vlan_id(self) -> pulumi.Output[Optional[float]]:
        """
        The VLAN ID.
        """
        return pulumi.get(self, "vlan_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

