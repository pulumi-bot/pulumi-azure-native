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

__all__ = ['VirtualNetworkGateway']


class VirtualNetworkGateway(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 active_active: Optional[pulumi.Input[bool]] = None,
                 bgp_settings: Optional[pulumi.Input[pulumi.InputType['BgpSettingsArgs']]] = None,
                 enable_bgp: Optional[pulumi.Input[bool]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 gateway_default_site: Optional[pulumi.Input[pulumi.InputType['SubResourceArgs']]] = None,
                 gateway_type: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 ip_configurations: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['VirtualNetworkGatewayIPConfigurationArgs']]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 resource_guid: Optional[pulumi.Input[str]] = None,
                 sku: Optional[pulumi.Input[pulumi.InputType['VirtualNetworkGatewaySkuArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vpn_client_configuration: Optional[pulumi.Input[pulumi.InputType['VpnClientConfigurationArgs']]] = None,
                 vpn_type: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        A common class for general resource information

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] active_active: ActiveActive flag
        :param pulumi.Input[pulumi.InputType['BgpSettingsArgs']] bgp_settings: Virtual network gateway's BGP speaker settings.
        :param pulumi.Input[bool] enable_bgp: Whether BGP is enabled for this virtual network gateway or not.
        :param pulumi.Input[str] etag: Gets a unique read-only string that changes whenever the resource is updated.
        :param pulumi.Input[pulumi.InputType['SubResourceArgs']] gateway_default_site: The reference of the LocalNetworkGateway resource which represents local network site having default routes. Assign Null value in case of removing existing default site setting.
        :param pulumi.Input[str] gateway_type: The type of this virtual network gateway. Possible values are: 'Vpn' and 'ExpressRoute'.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['VirtualNetworkGatewayIPConfigurationArgs']]]] ip_configurations: IP configurations for virtual network gateway.
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[str] name: The name of the virtual network gateway.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] resource_guid: The resource GUID property of the VirtualNetworkGateway resource.
        :param pulumi.Input[pulumi.InputType['VirtualNetworkGatewaySkuArgs']] sku: The reference of the VirtualNetworkGatewaySku resource which represents the SKU selected for Virtual network gateway.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags.
        :param pulumi.Input[pulumi.InputType['VpnClientConfigurationArgs']] vpn_client_configuration: The reference of the VpnClientConfiguration resource which represents the P2S VpnClient configurations.
        :param pulumi.Input[str] vpn_type: The type of this virtual network gateway. Possible values are: 'PolicyBased' and 'RouteBased'.
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

            __props__['active_active'] = active_active
            __props__['bgp_settings'] = bgp_settings
            __props__['enable_bgp'] = enable_bgp
            __props__['etag'] = etag
            __props__['gateway_default_site'] = gateway_default_site
            __props__['gateway_type'] = gateway_type
            __props__['id'] = id
            __props__['ip_configurations'] = ip_configurations
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['resource_guid'] = resource_guid
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['vpn_client_configuration'] = vpn_client_configuration
            __props__['vpn_type'] = vpn_type
            __props__['provisioning_state'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:network/v20150615:VirtualNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20160330:VirtualNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20160601:VirtualNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20160901:VirtualNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20161201:VirtualNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20170601:VirtualNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20170801:VirtualNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20170901:VirtualNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20171001:VirtualNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20171101:VirtualNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20180101:VirtualNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20180201:VirtualNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20180401:VirtualNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20180601:VirtualNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20180701:VirtualNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20180801:VirtualNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20181001:VirtualNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20181101:VirtualNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20181201:VirtualNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20190201:VirtualNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20190401:VirtualNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20190601:VirtualNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20190701:VirtualNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20190801:VirtualNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20190901:VirtualNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20191101:VirtualNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20191201:VirtualNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20200301:VirtualNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20200401:VirtualNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20200501:VirtualNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20200601:VirtualNetworkGateway")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(VirtualNetworkGateway, __self__).__init__(
            'azurerm:network/v20170301:VirtualNetworkGateway',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'VirtualNetworkGateway':
        """
        Get an existing VirtualNetworkGateway resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return VirtualNetworkGateway(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="activeActive")
    def active_active(self) -> Optional[bool]:
        """
        ActiveActive flag
        """
        return pulumi.get(self, "active_active")

    @property
    @pulumi.getter(name="bgpSettings")
    def bgp_settings(self) -> Optional['outputs.BgpSettingsResponse']:
        """
        Virtual network gateway's BGP speaker settings.
        """
        return pulumi.get(self, "bgp_settings")

    @property
    @pulumi.getter(name="enableBgp")
    def enable_bgp(self) -> Optional[bool]:
        """
        Whether BGP is enabled for this virtual network gateway or not.
        """
        return pulumi.get(self, "enable_bgp")

    @property
    @pulumi.getter
    def etag(self) -> Optional[str]:
        """
        Gets a unique read-only string that changes whenever the resource is updated.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="gatewayDefaultSite")
    def gateway_default_site(self) -> Optional['outputs.SubResourceResponse']:
        """
        The reference of the LocalNetworkGateway resource which represents local network site having default routes. Assign Null value in case of removing existing default site setting.
        """
        return pulumi.get(self, "gateway_default_site")

    @property
    @pulumi.getter(name="gatewayType")
    def gateway_type(self) -> Optional[str]:
        """
        The type of this virtual network gateway. Possible values are: 'Vpn' and 'ExpressRoute'.
        """
        return pulumi.get(self, "gateway_type")

    @property
    @pulumi.getter(name="ipConfigurations")
    def ip_configurations(self) -> Optional[List['outputs.VirtualNetworkGatewayIPConfigurationResponse']]:
        """
        IP configurations for virtual network gateway.
        """
        return pulumi.get(self, "ip_configurations")

    @property
    @pulumi.getter
    def location(self) -> Optional[str]:
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
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        The provisioning state of the VirtualNetworkGateway resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="resourceGuid")
    def resource_guid(self) -> Optional[str]:
        """
        The resource GUID property of the VirtualNetworkGateway resource.
        """
        return pulumi.get(self, "resource_guid")

    @property
    @pulumi.getter
    def sku(self) -> Optional['outputs.VirtualNetworkGatewaySkuResponse']:
        """
        The reference of the VirtualNetworkGatewaySku resource which represents the SKU selected for Virtual network gateway.
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
    @pulumi.getter(name="vpnClientConfiguration")
    def vpn_client_configuration(self) -> Optional['outputs.VpnClientConfigurationResponse']:
        """
        The reference of the VpnClientConfiguration resource which represents the P2S VpnClient configurations.
        """
        return pulumi.get(self, "vpn_client_configuration")

    @property
    @pulumi.getter(name="vpnType")
    def vpn_type(self) -> Optional[str]:
        """
        The type of this virtual network gateway. Possible values are: 'PolicyBased' and 'RouteBased'.
        """
        return pulumi.get(self, "vpn_type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

