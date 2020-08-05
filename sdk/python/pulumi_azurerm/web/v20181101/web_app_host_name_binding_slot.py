# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class WebAppHostNameBindingSlot(pulumi.CustomResource):
    kind: pulumi.Output[str]
    """
    Kind of resource.
    """
    name: pulumi.Output[str]
    """
    Resource Name.
    """
    properties: pulumi.Output[dict]
    """
    HostNameBinding resource specific properties
      * `azure_resource_name` (`str`) - Azure resource name.
      * `azure_resource_type` (`str`) - Azure resource type.
      * `custom_host_name_dns_record_type` (`str`) - Custom DNS record type.
      * `domain_id` (`str`) - Fully qualified ARM domain resource URI.
      * `host_name_type` (`str`) - Hostname type.
      * `site_name` (`str`) - App Service app name.
      * `ssl_state` (`str`) - SSL type
      * `thumbprint` (`str`) - SSL certificate thumbprint
      * `virtual_ip` (`str`) - Virtual IP address assigned to the hostname if IP based SSL is enabled.
    """
    type: pulumi.Output[str]
    """
    Resource type.
    """
    def __init__(__self__, resource_name, opts=None, azure_resource_name=None, azure_resource_type=None, custom_host_name_dns_record_type=None, domain_id=None, host_name_type=None, kind=None, name=None, resource_group_name=None, site_name=None, slot=None, ssl_state=None, thumbprint=None, __props__=None, __name__=None, __opts__=None):
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
            __props__['properties'] = None
            __props__['type'] = None
        super(WebAppHostNameBindingSlot, __self__).__init__(
            'azurerm:web/v20181101:WebAppHostNameBindingSlot',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing WebAppHostNameBindingSlot resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return WebAppHostNameBindingSlot(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
