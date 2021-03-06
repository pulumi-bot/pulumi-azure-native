# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['SiteVNETConnectionSlot']


class SiteVNETConnectionSlot(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cert_blob: Optional[pulumi.Input[str]] = None,
                 cert_thumbprint: Optional[pulumi.Input[str]] = None,
                 dns_servers: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 resync_required: Optional[pulumi.Input[bool]] = None,
                 routes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VnetRouteArgs']]]]] = None,
                 slot: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 vnet_name: Optional[pulumi.Input[str]] = None,
                 vnet_resource_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        VNETInfo contract. This contract is public and is a stripped down version of VNETInfoInternal

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cert_blob: A certificate file (.cer) blob containing the public key of the private key used to authenticate a 
                           Point-To-Site VPN connection.
        :param pulumi.Input[str] cert_thumbprint: The client certificate thumbprint
        :param pulumi.Input[str] dns_servers: Dns servers to be used by this VNET. This should be a comma-separated list of IP addresses.
        :param pulumi.Input[str] id: Resource Id
        :param pulumi.Input[str] kind: Kind of resource
        :param pulumi.Input[str] location: Resource Location
        :param pulumi.Input[str] name: Resource Name
        :param pulumi.Input[str] resource_group_name: The resource group name
        :param pulumi.Input[bool] resync_required: Flag to determine if a resync is required
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VnetRouteArgs']]]] routes: The routes that this virtual network connection uses.
        :param pulumi.Input[str] slot: The name of the slot for this web app.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags
        :param pulumi.Input[str] type: Resource type
        :param pulumi.Input[str] vnet_name: The name of the Virtual Network
        :param pulumi.Input[str] vnet_resource_id: The vnet resource id
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

            __props__['cert_blob'] = cert_blob
            __props__['cert_thumbprint'] = cert_thumbprint
            __props__['dns_servers'] = dns_servers
            __props__['id'] = id
            __props__['kind'] = kind
            __props__['location'] = location
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['resync_required'] = resync_required
            __props__['routes'] = routes
            if slot is None and not opts.urn:
                raise TypeError("Missing required property 'slot'")
            __props__['slot'] = slot
            __props__['tags'] = tags
            __props__['type'] = type
            __props__['vnet_name'] = vnet_name
            __props__['vnet_resource_id'] = vnet_resource_id
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:web/v20150801:SiteVNETConnectionSlot"), pulumi.Alias(type_="azure-native:web:SiteVNETConnectionSlot"), pulumi.Alias(type_="azure-nextgen:web:SiteVNETConnectionSlot"), pulumi.Alias(type_="azure-native:web/latest:SiteVNETConnectionSlot"), pulumi.Alias(type_="azure-nextgen:web/latest:SiteVNETConnectionSlot"), pulumi.Alias(type_="azure-native:web/v20160801:SiteVNETConnectionSlot"), pulumi.Alias(type_="azure-nextgen:web/v20160801:SiteVNETConnectionSlot"), pulumi.Alias(type_="azure-native:web/v20180201:SiteVNETConnectionSlot"), pulumi.Alias(type_="azure-nextgen:web/v20180201:SiteVNETConnectionSlot"), pulumi.Alias(type_="azure-native:web/v20181101:SiteVNETConnectionSlot"), pulumi.Alias(type_="azure-nextgen:web/v20181101:SiteVNETConnectionSlot"), pulumi.Alias(type_="azure-native:web/v20190801:SiteVNETConnectionSlot"), pulumi.Alias(type_="azure-nextgen:web/v20190801:SiteVNETConnectionSlot"), pulumi.Alias(type_="azure-native:web/v20200601:SiteVNETConnectionSlot"), pulumi.Alias(type_="azure-nextgen:web/v20200601:SiteVNETConnectionSlot"), pulumi.Alias(type_="azure-native:web/v20200901:SiteVNETConnectionSlot"), pulumi.Alias(type_="azure-nextgen:web/v20200901:SiteVNETConnectionSlot"), pulumi.Alias(type_="azure-native:web/v20201001:SiteVNETConnectionSlot"), pulumi.Alias(type_="azure-nextgen:web/v20201001:SiteVNETConnectionSlot")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(SiteVNETConnectionSlot, __self__).__init__(
            'azure-native:web/v20150801:SiteVNETConnectionSlot',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'SiteVNETConnectionSlot':
        """
        Get an existing SiteVNETConnectionSlot resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["cert_blob"] = None
        __props__["cert_thumbprint"] = None
        __props__["dns_servers"] = None
        __props__["kind"] = None
        __props__["location"] = None
        __props__["name"] = None
        __props__["resync_required"] = None
        __props__["routes"] = None
        __props__["tags"] = None
        __props__["type"] = None
        __props__["vnet_resource_id"] = None
        return SiteVNETConnectionSlot(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="certBlob")
    def cert_blob(self) -> pulumi.Output[Optional[str]]:
        """
        A certificate file (.cer) blob containing the public key of the private key used to authenticate a 
                    Point-To-Site VPN connection.
        """
        return pulumi.get(self, "cert_blob")

    @property
    @pulumi.getter(name="certThumbprint")
    def cert_thumbprint(self) -> pulumi.Output[Optional[str]]:
        """
        The client certificate thumbprint
        """
        return pulumi.get(self, "cert_thumbprint")

    @property
    @pulumi.getter(name="dnsServers")
    def dns_servers(self) -> pulumi.Output[Optional[str]]:
        """
        Dns servers to be used by this VNET. This should be a comma-separated list of IP addresses.
        """
        return pulumi.get(self, "dns_servers")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[Optional[str]]:
        """
        Kind of resource
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Resource Location
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        """
        Resource Name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resyncRequired")
    def resync_required(self) -> pulumi.Output[Optional[bool]]:
        """
        Flag to determine if a resync is required
        """
        return pulumi.get(self, "resync_required")

    @property
    @pulumi.getter
    def routes(self) -> pulumi.Output[Optional[Sequence['outputs.VnetRouteResponse']]]:
        """
        The routes that this virtual network connection uses.
        """
        return pulumi.get(self, "routes")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Resource tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[str]]:
        """
        Resource type
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="vnetResourceId")
    def vnet_resource_id(self) -> pulumi.Output[Optional[str]]:
        """
        The vnet resource id
        """
        return pulumi.get(self, "vnet_resource_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

