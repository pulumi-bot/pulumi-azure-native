# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = ['VirtualRouterPeering']


class VirtualRouterPeering(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 peer_asn: Optional[pulumi.Input[float]] = None,
                 peer_ip: Optional[pulumi.Input[str]] = None,
                 peering_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 virtual_router_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Virtual Router Peering resource.

        ## Example Usage
        ### Create Virtual Router Peering

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        virtual_router_peering = azurerm.network.v20200601.VirtualRouterPeering("virtualRouterPeering",
            peer_asn=20000,
            peer_ip="192.168.1.5",
            peering_name="peering1",
            resource_group_name="rg1",
            virtual_router_name="virtualRouter")

        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[str] name: Name of the virtual router peering that is unique within a virtual router.
        :param pulumi.Input[float] peer_asn: Peer ASN.
        :param pulumi.Input[str] peer_ip: Peer IP.
        :param pulumi.Input[str] peering_name: The name of the Virtual Router Peering.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] virtual_router_name: The name of the Virtual Router.
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
            __props__['name'] = name
            __props__['peer_asn'] = peer_asn
            __props__['peer_ip'] = peer_ip
            if peering_name is None:
                raise TypeError("Missing required property 'peering_name'")
            __props__['peering_name'] = peering_name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if virtual_router_name is None:
                raise TypeError("Missing required property 'virtual_router_name'")
            __props__['virtual_router_name'] = virtual_router_name
            __props__['etag'] = None
            __props__['provisioning_state'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:network/latest:VirtualRouterPeering"), pulumi.Alias(type_="azurerm:network/v20190701:VirtualRouterPeering"), pulumi.Alias(type_="azurerm:network/v20190801:VirtualRouterPeering"), pulumi.Alias(type_="azurerm:network/v20190901:VirtualRouterPeering"), pulumi.Alias(type_="azurerm:network/v20191101:VirtualRouterPeering"), pulumi.Alias(type_="azurerm:network/v20191201:VirtualRouterPeering"), pulumi.Alias(type_="azurerm:network/v20200301:VirtualRouterPeering"), pulumi.Alias(type_="azurerm:network/v20200401:VirtualRouterPeering"), pulumi.Alias(type_="azurerm:network/v20200501:VirtualRouterPeering")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(VirtualRouterPeering, __self__).__init__(
            'azurerm:network/v20200601:VirtualRouterPeering',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'VirtualRouterPeering':
        """
        Get an existing VirtualRouterPeering resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return VirtualRouterPeering(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        A unique read-only string that changes whenever the resource is updated.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        """
        Name of the virtual router peering that is unique within a virtual router.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="peerAsn")
    def peer_asn(self) -> pulumi.Output[Optional[float]]:
        """
        Peer ASN.
        """
        return pulumi.get(self, "peer_asn")

    @property
    @pulumi.getter(name="peerIp")
    def peer_ip(self) -> pulumi.Output[Optional[str]]:
        """
        Peer IP.
        """
        return pulumi.get(self, "peer_ip")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        The provisioning state of the resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Peering type.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

