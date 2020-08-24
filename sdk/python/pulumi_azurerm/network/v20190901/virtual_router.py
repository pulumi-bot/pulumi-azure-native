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

__all__ = ['VirtualRouter']


class VirtualRouter(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 hosted_gateway: Optional[pulumi.Input[pulumi.InputType['SubResourceArgs']]] = None,
                 hosted_subnet: Optional[pulumi.Input[pulumi.InputType['SubResourceArgs']]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 virtual_router_asn: Optional[pulumi.Input[float]] = None,
                 virtual_router_ips: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        VirtualRouter Resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['SubResourceArgs']] hosted_gateway: The Gateway on which VirtualRouter is hosted.
        :param pulumi.Input[pulumi.InputType['SubResourceArgs']] hosted_subnet: The Subnet on which VirtualRouter is hosted.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[str] name: The name of the Virtual Router.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags.
        :param pulumi.Input[float] virtual_router_asn: VirtualRouter ASN.
        :param pulumi.Input[List[pulumi.Input[str]]] virtual_router_ips: VirtualRouter IPs.
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

            __props__['hosted_gateway'] = hosted_gateway
            __props__['hosted_subnet'] = hosted_subnet
            __props__['id'] = id
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['virtual_router_asn'] = virtual_router_asn
            __props__['virtual_router_ips'] = virtual_router_ips
            __props__['etag'] = None
            __props__['peerings'] = None
            __props__['provisioning_state'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:network/v20190701:VirtualRouter"), pulumi.Alias(type_="azurerm:network/v20190801:VirtualRouter"), pulumi.Alias(type_="azurerm:network/v20191101:VirtualRouter"), pulumi.Alias(type_="azurerm:network/v20191201:VirtualRouter"), pulumi.Alias(type_="azurerm:network/v20200301:VirtualRouter"), pulumi.Alias(type_="azurerm:network/v20200401:VirtualRouter"), pulumi.Alias(type_="azurerm:network/v20200501:VirtualRouter"), pulumi.Alias(type_="azurerm:network/v20200601:VirtualRouter")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(VirtualRouter, __self__).__init__(
            'azurerm:network/v20190901:VirtualRouter',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'VirtualRouter':
        """
        Get an existing VirtualRouter resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return VirtualRouter(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        A unique read-only string that changes whenever the resource is updated.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="hostedGateway")
    def hosted_gateway(self) -> Optional['outputs.SubResourceResponse']:
        """
        The Gateway on which VirtualRouter is hosted.
        """
        return pulumi.get(self, "hosted_gateway")

    @property
    @pulumi.getter(name="hostedSubnet")
    def hosted_subnet(self) -> Optional['outputs.SubResourceResponse']:
        """
        The Subnet on which VirtualRouter is hosted.
        """
        return pulumi.get(self, "hosted_subnet")

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
    @pulumi.getter
    def peerings(self) -> List['outputs.SubResourceResponse']:
        """
        List of references to VirtualRouterPeerings.
        """
        return pulumi.get(self, "peerings")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        The provisioning state of the resource.
        """
        return pulumi.get(self, "provisioning_state")

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
    @pulumi.getter(name="virtualRouterAsn")
    def virtual_router_asn(self) -> Optional[float]:
        """
        VirtualRouter ASN.
        """
        return pulumi.get(self, "virtual_router_asn")

    @property
    @pulumi.getter(name="virtualRouterIps")
    def virtual_router_ips(self) -> Optional[List[str]]:
        """
        VirtualRouter IPs.
        """
        return pulumi.get(self, "virtual_router_ips")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

