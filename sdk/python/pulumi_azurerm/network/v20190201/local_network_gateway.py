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

__all__ = ['LocalNetworkGateway']


class LocalNetworkGateway(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bgp_settings: Optional[pulumi.Input[pulumi.InputType['BgpSettingsArgs']]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 gateway_ip_address: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 local_network_address_space: Optional[pulumi.Input[pulumi.InputType['AddressSpaceArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 resource_guid: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        A common class for general resource information

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['BgpSettingsArgs']] bgp_settings: Local network gateway's BGP speaker settings.
        :param pulumi.Input[str] etag: A unique read-only string that changes whenever the resource is updated.
        :param pulumi.Input[str] gateway_ip_address: IP address of local network gateway.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[pulumi.InputType['AddressSpaceArgs']] local_network_address_space: Local network site address space.
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[str] name: The name of the local network gateway.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] resource_guid: The resource GUID property of the LocalNetworkGateway resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags.
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

            __props__['bgp_settings'] = bgp_settings
            __props__['etag'] = etag
            __props__['gateway_ip_address'] = gateway_ip_address
            __props__['id'] = id
            __props__['local_network_address_space'] = local_network_address_space
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['resource_guid'] = resource_guid
            __props__['tags'] = tags
            __props__['provisioning_state'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:network/v20150615:LocalNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20160330:LocalNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20160601:LocalNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20160901:LocalNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20161201:LocalNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20170301:LocalNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20170601:LocalNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20170801:LocalNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20170901:LocalNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20171001:LocalNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20171101:LocalNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20180101:LocalNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20180201:LocalNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20180401:LocalNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20180601:LocalNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20180701:LocalNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20180801:LocalNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20181001:LocalNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20181101:LocalNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20181201:LocalNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20190401:LocalNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20190601:LocalNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20190701:LocalNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20190801:LocalNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20190901:LocalNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20191101:LocalNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20191201:LocalNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20200301:LocalNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20200401:LocalNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20200501:LocalNetworkGateway"), pulumi.Alias(type_="azurerm:network/v20200601:LocalNetworkGateway")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(LocalNetworkGateway, __self__).__init__(
            'azurerm:network/v20190201:LocalNetworkGateway',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'LocalNetworkGateway':
        """
        Get an existing LocalNetworkGateway resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return LocalNetworkGateway(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="bgpSettings")
    def bgp_settings(self) -> Optional['outputs.BgpSettingsResponse']:
        """
        Local network gateway's BGP speaker settings.
        """
        return pulumi.get(self, "bgp_settings")

    @property
    @pulumi.getter
    def etag(self) -> Optional[str]:
        """
        A unique read-only string that changes whenever the resource is updated.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="gatewayIpAddress")
    def gateway_ip_address(self) -> Optional[str]:
        """
        IP address of local network gateway.
        """
        return pulumi.get(self, "gateway_ip_address")

    @property
    @pulumi.getter(name="localNetworkAddressSpace")
    def local_network_address_space(self) -> Optional['outputs.AddressSpaceResponse']:
        """
        Local network site address space.
        """
        return pulumi.get(self, "local_network_address_space")

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
        The provisioning state of the LocalNetworkGateway resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="resourceGuid")
    def resource_guid(self) -> Optional[str]:
        """
        The resource GUID property of the LocalNetworkGateway resource.
        """
        return pulumi.get(self, "resource_guid")

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

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

