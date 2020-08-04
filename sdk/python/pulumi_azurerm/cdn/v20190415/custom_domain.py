# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class CustomDomain(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    Resource name.
    """
    properties: pulumi.Output[dict]
    """
    The JSON object that contains the properties of the custom domain to create.
      * `custom_https_parameters` (`dict`) - Certificate parameters for securing custom HTTPS
        * `certificate_source` (`str`) - Defines the source of the SSL certificate.
        * `minimum_tls_version` (`str`) - TLS protocol version that will be used for Https
        * `protocol_type` (`str`) - Defines the TLS extension protocol that is used for secure delivery.

      * `custom_https_provisioning_state` (`str`) - Provisioning status of Custom Https of the custom domain.
      * `custom_https_provisioning_substate` (`str`) - Provisioning substate shows the progress of custom HTTPS enabling/disabling process step by step.
      * `host_name` (`str`) - The host name of the custom domain. Must be a domain name.
      * `provisioning_state` (`str`) - Provisioning status of the custom domain.
      * `resource_state` (`str`) - Resource status of the custom domain.
      * `validation_data` (`str`) - Special validation or data may be required when delivering CDN to some regions due to local compliance reasons. E.g. ICP license number of a custom domain is required to deliver content in China.
    """
    type: pulumi.Output[str]
    """
    Resource type.
    """
    def __init__(__self__, resource_name, opts=None, endpoint_name=None, name=None, profile_name=None, properties=None, resource_group_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Friendly domain name mapping to the endpoint hostname that the customer provides for branding purposes, e.g. www.contoso.com.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] endpoint_name: Name of the endpoint under the profile which is unique globally.
        :param pulumi.Input[str] name: Name of the custom domain within an endpoint.
        :param pulumi.Input[str] profile_name: Name of the CDN profile which is unique within the resource group.
        :param pulumi.Input[dict] properties: The JSON object that contains the properties of the custom domain to create.
        :param pulumi.Input[str] resource_group_name: Name of the Resource group within the Azure subscription.

        The **properties** object supports the following:

          * `host_name` (`pulumi.Input[str]`) - The host name of the custom domain. Must be a domain name.
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

            if endpoint_name is None:
                raise TypeError("Missing required property 'endpoint_name'")
            __props__['endpoint_name'] = endpoint_name
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if profile_name is None:
                raise TypeError("Missing required property 'profile_name'")
            __props__['profile_name'] = profile_name
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['type'] = None
        super(CustomDomain, __self__).__init__(
            'azurerm:cdn/v20190415:CustomDomain',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing CustomDomain resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return CustomDomain(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
