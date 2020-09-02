# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = ['OpenIdConnectProvider']


class OpenIdConnectProvider(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_secret: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 metadata_endpoint: Optional[pulumi.Input[str]] = None,
                 opid: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        OpenId Connect Provider details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] client_id: Client ID of developer console which is the client application.
        :param pulumi.Input[str] client_secret: Client Secret of developer console which is the client application.
        :param pulumi.Input[str] description: User-friendly description of OpenID Connect Provider.
        :param pulumi.Input[str] display_name: User-friendly OpenID Connect Provider name.
        :param pulumi.Input[str] metadata_endpoint: Metadata endpoint URI.
        :param pulumi.Input[str] opid: Identifier of the OpenID Connect Provider.
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

            if client_id is None:
                raise TypeError("Missing required property 'client_id'")
            __props__['client_id'] = client_id
            __props__['client_secret'] = client_secret
            __props__['description'] = description
            if display_name is None:
                raise TypeError("Missing required property 'display_name'")
            __props__['display_name'] = display_name
            if metadata_endpoint is None:
                raise TypeError("Missing required property 'metadata_endpoint'")
            __props__['metadata_endpoint'] = metadata_endpoint
            if opid is None:
                raise TypeError("Missing required property 'opid'")
            __props__['opid'] = opid
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if service_name is None:
                raise TypeError("Missing required property 'service_name'")
            __props__['service_name'] = service_name
            __props__['name'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:apimanagement/v20160707:OpenIdConnectProvider"), pulumi.Alias(type_="azurerm:apimanagement/v20161010:OpenIdConnectProvider"), pulumi.Alias(type_="azurerm:apimanagement/v20170301:OpenIdConnectProvider"), pulumi.Alias(type_="azurerm:apimanagement/v20180101:OpenIdConnectProvider"), pulumi.Alias(type_="azurerm:apimanagement/v20180601preview:OpenIdConnectProvider"), pulumi.Alias(type_="azurerm:apimanagement/v20190101:OpenIdConnectProvider"), pulumi.Alias(type_="azurerm:apimanagement/v20191201:OpenIdConnectProvider"), pulumi.Alias(type_="azurerm:apimanagement/v20191201preview:OpenIdConnectProvider")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(OpenIdConnectProvider, __self__).__init__(
            'azurerm:apimanagement/latest:OpenIdConnectProvider',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'OpenIdConnectProvider':
        """
        Get an existing OpenIdConnectProvider resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return OpenIdConnectProvider(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Output[str]:
        """
        Client ID of developer console which is the client application.
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> pulumi.Output[Optional[str]]:
        """
        Client Secret of developer console which is the client application.
        """
        return pulumi.get(self, "client_secret")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        User-friendly description of OpenID Connect Provider.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        User-friendly OpenID Connect Provider name.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="metadataEndpoint")
    def metadata_endpoint(self) -> pulumi.Output[str]:
        """
        Metadata endpoint URI.
        """
        return pulumi.get(self, "metadata_endpoint")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Resource type for API Management resource.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

