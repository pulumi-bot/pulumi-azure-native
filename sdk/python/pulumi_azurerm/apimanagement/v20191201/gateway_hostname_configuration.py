# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class GatewayHostnameConfiguration(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    Resource name.
    """
    properties: pulumi.Output[dict]
    """
    Gateway hostname configuration details.
      * `certificate_id` (`str`) - Identifier of Certificate entity that will be used for TLS connection establishment
      * `hostname` (`str`) - Hostname value. Supports valid domain name, partial or full wildcard
      * `negotiate_client_certificate` (`bool`) - Determines whether gateway requests client certificate
    """
    type: pulumi.Output[str]
    """
    Resource type for API Management resource.
    """
    def __init__(__self__, resource_name, opts=None, certificate_id=None, gateway_id=None, hostname=None, name=None, negotiate_client_certificate=None, resource_group_name=None, service_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Gateway hostname configuration details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] certificate_id: Identifier of Certificate entity that will be used for TLS connection establishment
        :param pulumi.Input[str] gateway_id: Gateway entity identifier. Must be unique in the current API Management service instance. Must not have value 'managed'
        :param pulumi.Input[str] hostname: Hostname value. Supports valid domain name, partial or full wildcard
        :param pulumi.Input[str] name: Gateway hostname configuration identifier. Must be unique in the scope of parent Gateway entity.
        :param pulumi.Input[bool] negotiate_client_certificate: Determines whether gateway requests client certificate
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] service_name: The name of the API Management service.
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

            __props__['certificate_id'] = certificate_id
            if gateway_id is None:
                raise TypeError("Missing required property 'gateway_id'")
            __props__['gateway_id'] = gateway_id
            __props__['hostname'] = hostname
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['negotiate_client_certificate'] = negotiate_client_certificate
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if service_name is None:
                raise TypeError("Missing required property 'service_name'")
            __props__['service_name'] = service_name
            __props__['properties'] = None
            __props__['type'] = None
        super(GatewayHostnameConfiguration, __self__).__init__(
            'azurerm:apimanagement/v20191201:GatewayHostnameConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing GatewayHostnameConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return GatewayHostnameConfiguration(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
