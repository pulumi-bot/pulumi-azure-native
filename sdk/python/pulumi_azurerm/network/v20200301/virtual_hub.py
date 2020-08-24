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

__all__ = ['VirtualHub']


class VirtualHub(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address_prefix: Optional[pulumi.Input[str]] = None,
                 azure_firewall: Optional[pulumi.Input[pulumi.InputType['SubResourceArgs']]] = None,
                 express_route_gateway: Optional[pulumi.Input[pulumi.InputType['SubResourceArgs']]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 p2_s_vpn_gateway: Optional[pulumi.Input[pulumi.InputType['SubResourceArgs']]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 route_table: Optional[pulumi.Input[pulumi.InputType['VirtualHubRouteTableArgs']]] = None,
                 security_partner_provider: Optional[pulumi.Input[pulumi.InputType['SubResourceArgs']]] = None,
                 security_provider_name: Optional[pulumi.Input[str]] = None,
                 sku: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 virtual_hub_route_table_v2s: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['VirtualHubRouteTableV2Args']]]]] = None,
                 virtual_network_connections: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['HubVirtualNetworkConnectionArgs']]]]] = None,
                 virtual_wan: Optional[pulumi.Input[pulumi.InputType['SubResourceArgs']]] = None,
                 vpn_gateway: Optional[pulumi.Input[pulumi.InputType['SubResourceArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        VirtualHub Resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address_prefix: Address-prefix for this VirtualHub.
        :param pulumi.Input[pulumi.InputType['SubResourceArgs']] azure_firewall: The azureFirewall associated with this VirtualHub.
        :param pulumi.Input[pulumi.InputType['SubResourceArgs']] express_route_gateway: The expressRouteGateway associated with this VirtualHub.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[str] name: The name of the VirtualHub.
        :param pulumi.Input[pulumi.InputType['SubResourceArgs']] p2_s_vpn_gateway: The P2SVpnGateway associated with this VirtualHub.
        :param pulumi.Input[str] resource_group_name: The resource group name of the VirtualHub.
        :param pulumi.Input[pulumi.InputType['VirtualHubRouteTableArgs']] route_table: The routeTable associated with this virtual hub.
        :param pulumi.Input[pulumi.InputType['SubResourceArgs']] security_partner_provider: The securityPartnerProvider associated with this VirtualHub.
        :param pulumi.Input[str] security_provider_name: The Security Provider name.
        :param pulumi.Input[str] sku: The sku of this VirtualHub.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['VirtualHubRouteTableV2Args']]]] virtual_hub_route_table_v2s: List of all virtual hub route table v2s associated with this VirtualHub.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['HubVirtualNetworkConnectionArgs']]]] virtual_network_connections: List of all vnet connections with this VirtualHub.
        :param pulumi.Input[pulumi.InputType['SubResourceArgs']] virtual_wan: The VirtualWAN to which the VirtualHub belongs.
        :param pulumi.Input[pulumi.InputType['SubResourceArgs']] vpn_gateway: The VpnGateway associated with this VirtualHub.
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
            __props__['azure_firewall'] = azure_firewall
            __props__['express_route_gateway'] = express_route_gateway
            __props__['id'] = id
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['p2_s_vpn_gateway'] = p2_s_vpn_gateway
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['route_table'] = route_table
            __props__['security_partner_provider'] = security_partner_provider
            __props__['security_provider_name'] = security_provider_name
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['virtual_hub_route_table_v2s'] = virtual_hub_route_table_v2s
            __props__['virtual_network_connections'] = virtual_network_connections
            __props__['virtual_wan'] = virtual_wan
            __props__['vpn_gateway'] = vpn_gateway
            __props__['etag'] = None
            __props__['provisioning_state'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:network/v20180401:VirtualHub"), pulumi.Alias(type_="azurerm:network/v20180601:VirtualHub"), pulumi.Alias(type_="azurerm:network/v20180701:VirtualHub"), pulumi.Alias(type_="azurerm:network/v20180801:VirtualHub"), pulumi.Alias(type_="azurerm:network/v20181001:VirtualHub"), pulumi.Alias(type_="azurerm:network/v20181101:VirtualHub"), pulumi.Alias(type_="azurerm:network/v20181201:VirtualHub"), pulumi.Alias(type_="azurerm:network/v20190201:VirtualHub"), pulumi.Alias(type_="azurerm:network/v20190401:VirtualHub"), pulumi.Alias(type_="azurerm:network/v20190601:VirtualHub"), pulumi.Alias(type_="azurerm:network/v20190701:VirtualHub"), pulumi.Alias(type_="azurerm:network/v20190801:VirtualHub"), pulumi.Alias(type_="azurerm:network/v20190901:VirtualHub"), pulumi.Alias(type_="azurerm:network/v20191101:VirtualHub"), pulumi.Alias(type_="azurerm:network/v20191201:VirtualHub"), pulumi.Alias(type_="azurerm:network/v20200401:VirtualHub"), pulumi.Alias(type_="azurerm:network/v20200501:VirtualHub"), pulumi.Alias(type_="azurerm:network/v20200601:VirtualHub")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(VirtualHub, __self__).__init__(
            'azurerm:network/v20200301:VirtualHub',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'VirtualHub':
        """
        Get an existing VirtualHub resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return VirtualHub(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="addressPrefix")
    def address_prefix(self) -> Optional[str]:
        """
        Address-prefix for this VirtualHub.
        """
        return pulumi.get(self, "address_prefix")

    @property
    @pulumi.getter(name="azureFirewall")
    def azure_firewall(self) -> Optional['outputs.SubResourceResponse']:
        """
        The azureFirewall associated with this VirtualHub.
        """
        return pulumi.get(self, "azure_firewall")

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        A unique read-only string that changes whenever the resource is updated.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="expressRouteGateway")
    def express_route_gateway(self) -> Optional['outputs.SubResourceResponse']:
        """
        The expressRouteGateway associated with this VirtualHub.
        """
        return pulumi.get(self, "express_route_gateway")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        Resource location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="p2SVpnGateway")
    def p2_s_vpn_gateway(self) -> Optional['outputs.SubResourceResponse']:
        """
        The P2SVpnGateway associated with this VirtualHub.
        """
        return pulumi.get(self, "p2_s_vpn_gateway")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        The provisioning state of the virtual hub resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="routeTable")
    def route_table(self) -> Optional['outputs.VirtualHubRouteTableResponse']:
        """
        The routeTable associated with this virtual hub.
        """
        return pulumi.get(self, "route_table")

    @property
    @pulumi.getter(name="securityPartnerProvider")
    def security_partner_provider(self) -> Optional['outputs.SubResourceResponse']:
        """
        The securityPartnerProvider associated with this VirtualHub.
        """
        return pulumi.get(self, "security_partner_provider")

    @property
    @pulumi.getter(name="securityProviderName")
    def security_provider_name(self) -> Optional[str]:
        """
        The Security Provider name.
        """
        return pulumi.get(self, "security_provider_name")

    @property
    @pulumi.getter
    def sku(self) -> Optional[str]:
        """
        The sku of this VirtualHub.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="virtualHubRouteTableV2s")
    def virtual_hub_route_table_v2s(self) -> Optional[List['outputs.VirtualHubRouteTableV2Response']]:
        """
        List of all virtual hub route table v2s associated with this VirtualHub.
        """
        return pulumi.get(self, "virtual_hub_route_table_v2s")

    @property
    @pulumi.getter(name="virtualNetworkConnections")
    def virtual_network_connections(self) -> Optional[List['outputs.HubVirtualNetworkConnectionResponse']]:
        """
        List of all vnet connections with this VirtualHub.
        """
        return pulumi.get(self, "virtual_network_connections")

    @property
    @pulumi.getter(name="virtualWan")
    def virtual_wan(self) -> Optional['outputs.SubResourceResponse']:
        """
        The VirtualWAN to which the VirtualHub belongs.
        """
        return pulumi.get(self, "virtual_wan")

    @property
    @pulumi.getter(name="vpnGateway")
    def vpn_gateway(self) -> Optional['outputs.SubResourceResponse']:
        """
        The VpnGateway associated with this VirtualHub.
        """
        return pulumi.get(self, "vpn_gateway")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

