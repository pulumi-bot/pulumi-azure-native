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

__all__ = ['Service']


class Service(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 group_name: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 public_key: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 sku: Optional[pulumi.Input[pulumi.InputType['ServiceSkuArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 virtual_subnet_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        A Database Migration Service resource

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] etag: HTTP strong entity tag value. Ignored if submitted
        :param pulumi.Input[str] group_name: Name of the resource group
        :param pulumi.Input[str] kind: The resource kind. Only 'vm' (the default) is supported.
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[str] public_key: The public key of the service, used to encrypt secrets sent to the service
        :param pulumi.Input[str] service_name: Name of the service
        :param pulumi.Input[pulumi.InputType['ServiceSkuArgs']] sku: Service SKU
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags.
        :param pulumi.Input[str] virtual_subnet_id: The ID of the Microsoft.Network/virtualNetworks/subnets resource to which the service should be joined
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
            if group_name is None:
                raise TypeError("Missing required property 'group_name'")
            __props__['group_name'] = group_name
            __props__['kind'] = kind
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            __props__['public_key'] = public_key
            if service_name is None:
                raise TypeError("Missing required property 'service_name'")
            __props__['service_name'] = service_name
            __props__['sku'] = sku
            __props__['tags'] = tags
            if virtual_subnet_id is None:
                raise TypeError("Missing required property 'virtual_subnet_id'")
            __props__['virtual_subnet_id'] = virtual_subnet_id
            __props__['name'] = None
            __props__['provisioning_state'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:datamigration/latest:Service"), pulumi.Alias(type_="azurerm:datamigration/v20171115preview:Service"), pulumi.Alias(type_="azurerm:datamigration/v20180315preview:Service"), pulumi.Alias(type_="azurerm:datamigration/v20180331preview:Service"), pulumi.Alias(type_="azurerm:datamigration/v20180419:Service")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Service, __self__).__init__(
            'azurerm:datamigration/v20180715preview:Service',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Service':
        """
        Get an existing Service resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Service(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[Optional[str]]:
        """
        HTTP strong entity tag value. Ignored if submitted
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[Optional[str]]:
        """
        The resource kind. Only 'vm' (the default) is supported.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
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
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        The resource's provisioning state
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> pulumi.Output[Optional[str]]:
        """
        The public key of the service, used to encrypt secrets sent to the service
        """
        return pulumi.get(self, "public_key")

    @property
    @pulumi.getter
    def sku(self) -> pulumi.Output[Optional['outputs.ServiceSkuResponse']]:
        """
        Service SKU
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
    @pulumi.getter(name="virtualSubnetId")
    def virtual_subnet_id(self) -> pulumi.Output[str]:
        """
        The ID of the Microsoft.Network/virtualNetworks/subnets resource to which the service should be joined
        """
        return pulumi.get(self, "virtual_subnet_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

