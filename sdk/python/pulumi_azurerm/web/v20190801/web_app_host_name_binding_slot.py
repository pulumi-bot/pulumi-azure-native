# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = ['WebAppHostNameBindingSlot']


class WebAppHostNameBindingSlot(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 azure_resource_name: Optional[pulumi.Input[str]] = None,
                 azure_resource_type: Optional[pulumi.Input[str]] = None,
                 custom_host_name_dns_record_type: Optional[pulumi.Input[str]] = None,
                 domain_id: Optional[pulumi.Input[str]] = None,
                 host_name_type: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 site_name: Optional[pulumi.Input[str]] = None,
                 slot: Optional[pulumi.Input[str]] = None,
                 ssl_state: Optional[pulumi.Input[str]] = None,
                 thumbprint: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        A hostname binding object.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] azure_resource_name: Azure resource name.
        :param pulumi.Input[str] azure_resource_type: Azure resource type.
        :param pulumi.Input[str] custom_host_name_dns_record_type: Custom DNS record type.
        :param pulumi.Input[str] domain_id: Fully qualified ARM domain resource URI.
        :param pulumi.Input[str] host_name_type: Hostname type.
        :param pulumi.Input[str] kind: Kind of resource.
        :param pulumi.Input[str] name: Hostname in the hostname binding.
        :param pulumi.Input[str] resource_group_name: Name of the resource group to which the resource belongs.
        :param pulumi.Input[str] site_name: App Service app name.
        :param pulumi.Input[str] slot: Name of the deployment slot. If a slot is not specified, the API will create a binding for the production slot.
        :param pulumi.Input[str] ssl_state: SSL type
        :param pulumi.Input[str] thumbprint: SSL certificate thumbprint
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

            __props__['azure_resource_name'] = azure_resource_name
            __props__['azure_resource_type'] = azure_resource_type
            __props__['custom_host_name_dns_record_type'] = custom_host_name_dns_record_type
            __props__['domain_id'] = domain_id
            __props__['host_name_type'] = host_name_type
            __props__['kind'] = kind
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['site_name'] = site_name
            if slot is None:
                raise TypeError("Missing required property 'slot'")
            __props__['slot'] = slot
            __props__['ssl_state'] = ssl_state
            __props__['thumbprint'] = thumbprint
            __props__['type'] = None
            __props__['virtual_ip'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:web/v20150801:WebAppHostNameBindingSlot"), pulumi.Alias(type_="azurerm:web/v20160801:WebAppHostNameBindingSlot"), pulumi.Alias(type_="azurerm:web/v20180201:WebAppHostNameBindingSlot"), pulumi.Alias(type_="azurerm:web/v20181101:WebAppHostNameBindingSlot"), pulumi.Alias(type_="azurerm:web/v20200601:WebAppHostNameBindingSlot")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(WebAppHostNameBindingSlot, __self__).__init__(
            'azurerm:web/v20190801:WebAppHostNameBindingSlot',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'WebAppHostNameBindingSlot':
        """
        Get an existing WebAppHostNameBindingSlot resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return WebAppHostNameBindingSlot(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureResourceName")
    def azure_resource_name(self) -> Optional[str]:
        """
        Azure resource name.
        """
        return pulumi.get(self, "azure_resource_name")

    @property
    @pulumi.getter(name="azureResourceType")
    def azure_resource_type(self) -> Optional[str]:
        """
        Azure resource type.
        """
        return pulumi.get(self, "azure_resource_type")

    @property
    @pulumi.getter(name="customHostNameDnsRecordType")
    def custom_host_name_dns_record_type(self) -> Optional[str]:
        """
        Custom DNS record type.
        """
        return pulumi.get(self, "custom_host_name_dns_record_type")

    @property
    @pulumi.getter(name="domainId")
    def domain_id(self) -> Optional[str]:
        """
        Fully qualified ARM domain resource URI.
        """
        return pulumi.get(self, "domain_id")

    @property
    @pulumi.getter(name="hostNameType")
    def host_name_type(self) -> Optional[str]:
        """
        Hostname type.
        """
        return pulumi.get(self, "host_name_type")

    @property
    @pulumi.getter
    def kind(self) -> Optional[str]:
        """
        Kind of resource.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource Name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="siteName")
    def site_name(self) -> Optional[str]:
        """
        App Service app name.
        """
        return pulumi.get(self, "site_name")

    @property
    @pulumi.getter(name="sslState")
    def ssl_state(self) -> Optional[str]:
        """
        SSL type
        """
        return pulumi.get(self, "ssl_state")

    @property
    @pulumi.getter
    def thumbprint(self) -> Optional[str]:
        """
        SSL certificate thumbprint
        """
        return pulumi.get(self, "thumbprint")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="virtualIP")
    def virtual_ip(self) -> str:
        """
        Virtual IP address assigned to the hostname if IP based SSL is enabled.
        """
        return pulumi.get(self, "virtual_ip")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

