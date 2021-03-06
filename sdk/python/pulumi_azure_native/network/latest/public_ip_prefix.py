# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['PublicIPPrefix']

warnings.warn("""The 'latest' version is deprecated. Please migrate to the resource in the top-level module: 'azure-native:network:PublicIPPrefix'.""", DeprecationWarning)


class PublicIPPrefix(pulumi.CustomResource):
    warnings.warn("""The 'latest' version is deprecated. Please migrate to the resource in the top-level module: 'azure-native:network:PublicIPPrefix'.""", DeprecationWarning)

    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 custom_ip_prefix: Optional[pulumi.Input[pulumi.InputType['SubResourceArgs']]] = None,
                 extended_location: Optional[pulumi.Input[pulumi.InputType['ExtendedLocationArgs']]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 ip_tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IpTagArgs']]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 nat_gateway: Optional[pulumi.Input[pulumi.InputType['NatGatewayArgs']]] = None,
                 prefix_length: Optional[pulumi.Input[int]] = None,
                 public_ip_address_version: Optional[pulumi.Input[Union[str, 'IPVersion']]] = None,
                 public_ip_prefix_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sku: Optional[pulumi.Input[pulumi.InputType['PublicIPPrefixSkuArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 zones: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Public IP prefix resource.
        Latest API Version: 2020-08-01.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['SubResourceArgs']] custom_ip_prefix: The customIpPrefix that this prefix is associated with.
        :param pulumi.Input[pulumi.InputType['ExtendedLocationArgs']] extended_location: The extended location of the public ip address.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IpTagArgs']]]] ip_tags: The list of tags associated with the public IP prefix.
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[pulumi.InputType['NatGatewayArgs']] nat_gateway: NatGateway of Public IP Prefix.
        :param pulumi.Input[int] prefix_length: The Length of the Public IP Prefix.
        :param pulumi.Input[Union[str, 'IPVersion']] public_ip_address_version: The public IP address version.
        :param pulumi.Input[str] public_ip_prefix_name: The name of the public IP prefix.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[pulumi.InputType['PublicIPPrefixSkuArgs']] sku: The public IP prefix SKU.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] zones: A list of availability zones denoting the IP allocated for the resource needs to come from.
        """
        pulumi.log.warn("""PublicIPPrefix is deprecated: The 'latest' version is deprecated. Please migrate to the resource in the top-level module: 'azure-native:network:PublicIPPrefix'.""")
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

            __props__['custom_ip_prefix'] = custom_ip_prefix
            __props__['extended_location'] = extended_location
            __props__['id'] = id
            __props__['ip_tags'] = ip_tags
            __props__['location'] = location
            __props__['nat_gateway'] = nat_gateway
            __props__['prefix_length'] = prefix_length
            __props__['public_ip_address_version'] = public_ip_address_version
            __props__['public_ip_prefix_name'] = public_ip_prefix_name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['zones'] = zones
            __props__['etag'] = None
            __props__['ip_prefix'] = None
            __props__['load_balancer_frontend_ip_configuration'] = None
            __props__['name'] = None
            __props__['provisioning_state'] = None
            __props__['public_ip_addresses'] = None
            __props__['resource_guid'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:network/latest:PublicIPPrefix"), pulumi.Alias(type_="azure-native:network:PublicIPPrefix"), pulumi.Alias(type_="azure-nextgen:network:PublicIPPrefix"), pulumi.Alias(type_="azure-native:network/v20180701:PublicIPPrefix"), pulumi.Alias(type_="azure-nextgen:network/v20180701:PublicIPPrefix"), pulumi.Alias(type_="azure-native:network/v20180801:PublicIPPrefix"), pulumi.Alias(type_="azure-nextgen:network/v20180801:PublicIPPrefix"), pulumi.Alias(type_="azure-native:network/v20181001:PublicIPPrefix"), pulumi.Alias(type_="azure-nextgen:network/v20181001:PublicIPPrefix"), pulumi.Alias(type_="azure-native:network/v20181101:PublicIPPrefix"), pulumi.Alias(type_="azure-nextgen:network/v20181101:PublicIPPrefix"), pulumi.Alias(type_="azure-native:network/v20181201:PublicIPPrefix"), pulumi.Alias(type_="azure-nextgen:network/v20181201:PublicIPPrefix"), pulumi.Alias(type_="azure-native:network/v20190201:PublicIPPrefix"), pulumi.Alias(type_="azure-nextgen:network/v20190201:PublicIPPrefix"), pulumi.Alias(type_="azure-native:network/v20190401:PublicIPPrefix"), pulumi.Alias(type_="azure-nextgen:network/v20190401:PublicIPPrefix"), pulumi.Alias(type_="azure-native:network/v20190601:PublicIPPrefix"), pulumi.Alias(type_="azure-nextgen:network/v20190601:PublicIPPrefix"), pulumi.Alias(type_="azure-native:network/v20190701:PublicIPPrefix"), pulumi.Alias(type_="azure-nextgen:network/v20190701:PublicIPPrefix"), pulumi.Alias(type_="azure-native:network/v20190801:PublicIPPrefix"), pulumi.Alias(type_="azure-nextgen:network/v20190801:PublicIPPrefix"), pulumi.Alias(type_="azure-native:network/v20190901:PublicIPPrefix"), pulumi.Alias(type_="azure-nextgen:network/v20190901:PublicIPPrefix"), pulumi.Alias(type_="azure-native:network/v20191101:PublicIPPrefix"), pulumi.Alias(type_="azure-nextgen:network/v20191101:PublicIPPrefix"), pulumi.Alias(type_="azure-native:network/v20191201:PublicIPPrefix"), pulumi.Alias(type_="azure-nextgen:network/v20191201:PublicIPPrefix"), pulumi.Alias(type_="azure-native:network/v20200301:PublicIPPrefix"), pulumi.Alias(type_="azure-nextgen:network/v20200301:PublicIPPrefix"), pulumi.Alias(type_="azure-native:network/v20200401:PublicIPPrefix"), pulumi.Alias(type_="azure-nextgen:network/v20200401:PublicIPPrefix"), pulumi.Alias(type_="azure-native:network/v20200501:PublicIPPrefix"), pulumi.Alias(type_="azure-nextgen:network/v20200501:PublicIPPrefix"), pulumi.Alias(type_="azure-native:network/v20200601:PublicIPPrefix"), pulumi.Alias(type_="azure-nextgen:network/v20200601:PublicIPPrefix"), pulumi.Alias(type_="azure-native:network/v20200701:PublicIPPrefix"), pulumi.Alias(type_="azure-nextgen:network/v20200701:PublicIPPrefix"), pulumi.Alias(type_="azure-native:network/v20200801:PublicIPPrefix"), pulumi.Alias(type_="azure-nextgen:network/v20200801:PublicIPPrefix")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(PublicIPPrefix, __self__).__init__(
            'azure-native:network/latest:PublicIPPrefix',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'PublicIPPrefix':
        """
        Get an existing PublicIPPrefix resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["custom_ip_prefix"] = None
        __props__["etag"] = None
        __props__["extended_location"] = None
        __props__["ip_prefix"] = None
        __props__["ip_tags"] = None
        __props__["load_balancer_frontend_ip_configuration"] = None
        __props__["location"] = None
        __props__["name"] = None
        __props__["nat_gateway"] = None
        __props__["prefix_length"] = None
        __props__["provisioning_state"] = None
        __props__["public_ip_address_version"] = None
        __props__["public_ip_addresses"] = None
        __props__["resource_guid"] = None
        __props__["sku"] = None
        __props__["tags"] = None
        __props__["type"] = None
        __props__["zones"] = None
        return PublicIPPrefix(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="customIPPrefix")
    def custom_ip_prefix(self) -> pulumi.Output[Optional['outputs.SubResourceResponse']]:
        """
        The customIpPrefix that this prefix is associated with.
        """
        return pulumi.get(self, "custom_ip_prefix")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        A unique read-only string that changes whenever the resource is updated.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="extendedLocation")
    def extended_location(self) -> pulumi.Output[Optional['outputs.ExtendedLocationResponse']]:
        """
        The extended location of the public ip address.
        """
        return pulumi.get(self, "extended_location")

    @property
    @pulumi.getter(name="ipPrefix")
    def ip_prefix(self) -> pulumi.Output[str]:
        """
        The allocated Prefix.
        """
        return pulumi.get(self, "ip_prefix")

    @property
    @pulumi.getter(name="ipTags")
    def ip_tags(self) -> pulumi.Output[Optional[Sequence['outputs.IpTagResponse']]]:
        """
        The list of tags associated with the public IP prefix.
        """
        return pulumi.get(self, "ip_tags")

    @property
    @pulumi.getter(name="loadBalancerFrontendIpConfiguration")
    def load_balancer_frontend_ip_configuration(self) -> pulumi.Output['outputs.SubResourceResponse']:
        """
        The reference to load balancer frontend IP configuration associated with the public IP prefix.
        """
        return pulumi.get(self, "load_balancer_frontend_ip_configuration")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[Optional[str]]:
        """
        Resource location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="natGateway")
    def nat_gateway(self) -> pulumi.Output[Optional['outputs.NatGatewayResponse']]:
        """
        NatGateway of Public IP Prefix.
        """
        return pulumi.get(self, "nat_gateway")

    @property
    @pulumi.getter(name="prefixLength")
    def prefix_length(self) -> pulumi.Output[Optional[int]]:
        """
        The Length of the Public IP Prefix.
        """
        return pulumi.get(self, "prefix_length")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        The provisioning state of the public IP prefix resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="publicIPAddressVersion")
    def public_ip_address_version(self) -> pulumi.Output[Optional[str]]:
        """
        The public IP address version.
        """
        return pulumi.get(self, "public_ip_address_version")

    @property
    @pulumi.getter(name="publicIPAddresses")
    def public_ip_addresses(self) -> pulumi.Output[Sequence['outputs.ReferencedPublicIpAddressResponse']]:
        """
        The list of all referenced PublicIPAddresses.
        """
        return pulumi.get(self, "public_ip_addresses")

    @property
    @pulumi.getter(name="resourceGuid")
    def resource_guid(self) -> pulumi.Output[str]:
        """
        The resource GUID property of the public IP prefix resource.
        """
        return pulumi.get(self, "resource_guid")

    @property
    @pulumi.getter
    def sku(self) -> pulumi.Output[Optional['outputs.PublicIPPrefixSkuResponse']]:
        """
        The public IP prefix SKU.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Resource type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def zones(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of availability zones denoting the IP allocated for the resource needs to come from.
        """
        return pulumi.get(self, "zones")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

