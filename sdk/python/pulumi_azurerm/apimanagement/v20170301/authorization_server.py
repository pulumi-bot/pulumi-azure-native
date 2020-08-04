# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class AuthorizationServer(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    Resource name.
    """
    properties: pulumi.Output[dict]
    """
    Properties of the External OAuth authorization server Contract.
      * `authorization_endpoint` (`str`) - OAuth authorization endpoint. See http://tools.ietf.org/html/rfc6749#section-3.2.
      * `authorization_methods` (`list`) - HTTP verbs supported by the authorization endpoint. GET must be always present. POST is optional.
      * `bearer_token_sending_methods` (`list`) - Specifies the mechanism by which access token is passed to the API. 
      * `client_authentication_method` (`list`) - Method of authentication supported by the token endpoint of this authorization server. Possible values are Basic and/or Body. When Body is specified, client credentials and other parameters are passed within the request body in the application/x-www-form-urlencoded format.
      * `client_id` (`str`) - Client or app id registered with this authorization server.
      * `client_registration_endpoint` (`str`) - Optional reference to a page where client or app registration for this authorization server is performed. Contains absolute URL to entity being referenced.
      * `client_secret` (`str`) - Client or app secret registered with this authorization server.
      * `default_scope` (`str`) - Access token scope that is going to be requested by default. Can be overridden at the API level. Should be provided in the form of a string containing space-delimited values.
      * `description` (`str`) - Description of the authorization server. Can contain HTML formatting tags.
      * `display_name` (`str`) - User-friendly authorization server name.
      * `grant_types` (`list`) - Form of an authorization grant, which the client uses to request the access token.
      * `resource_owner_password` (`str`) - Can be optionally specified when resource owner password grant type is supported by this authorization server. Default resource owner password.
      * `resource_owner_username` (`str`) - Can be optionally specified when resource owner password grant type is supported by this authorization server. Default resource owner username.
      * `support_state` (`bool`) - If true, authorization server will include state parameter from the authorization request to its response. Client may use state parameter to raise protocol security.
      * `token_body_parameters` (`list`) - Additional parameters required by the token endpoint of this authorization server represented as an array of JSON objects with name and value string properties, i.e. {"name" : "name value", "value": "a value"}.
        * `name` (`str`) - body parameter name.
        * `value` (`str`) - body parameter value.

      * `token_endpoint` (`str`) - OAuth token endpoint. Contains absolute URI to entity being referenced.
    """
    type: pulumi.Output[str]
    """
    Resource type for API Management resource.
    """
    def __init__(__self__, resource_name, opts=None, name=None, properties=None, resource_group_name=None, service_name=None, __props__=None, __name__=None, __opts__=None):
        """
        External OAuth authorization server settings.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Identifier of the authorization server.
        :param pulumi.Input[dict] properties: Properties of the External OAuth authorization server Contract.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] service_name: The name of the API Management service.

        The **properties** object supports the following:

          * `authorization_endpoint` (`pulumi.Input[str]`) - OAuth authorization endpoint. See http://tools.ietf.org/html/rfc6749#section-3.2.
          * `authorization_methods` (`pulumi.Input[list]`) - HTTP verbs supported by the authorization endpoint. GET must be always present. POST is optional.
          * `bearer_token_sending_methods` (`pulumi.Input[list]`) - Specifies the mechanism by which access token is passed to the API. 
          * `client_authentication_method` (`pulumi.Input[list]`) - Method of authentication supported by the token endpoint of this authorization server. Possible values are Basic and/or Body. When Body is specified, client credentials and other parameters are passed within the request body in the application/x-www-form-urlencoded format.
          * `client_id` (`pulumi.Input[str]`) - Client or app id registered with this authorization server.
          * `client_registration_endpoint` (`pulumi.Input[str]`) - Optional reference to a page where client or app registration for this authorization server is performed. Contains absolute URL to entity being referenced.
          * `client_secret` (`pulumi.Input[str]`) - Client or app secret registered with this authorization server.
          * `default_scope` (`pulumi.Input[str]`) - Access token scope that is going to be requested by default. Can be overridden at the API level. Should be provided in the form of a string containing space-delimited values.
          * `description` (`pulumi.Input[str]`) - Description of the authorization server. Can contain HTML formatting tags.
          * `display_name` (`pulumi.Input[str]`) - User-friendly authorization server name.
          * `grant_types` (`pulumi.Input[list]`) - Form of an authorization grant, which the client uses to request the access token.
          * `resource_owner_password` (`pulumi.Input[str]`) - Can be optionally specified when resource owner password grant type is supported by this authorization server. Default resource owner password.
          * `resource_owner_username` (`pulumi.Input[str]`) - Can be optionally specified when resource owner password grant type is supported by this authorization server. Default resource owner username.
          * `support_state` (`pulumi.Input[bool]`) - If true, authorization server will include state parameter from the authorization request to its response. Client may use state parameter to raise protocol security.
          * `token_body_parameters` (`pulumi.Input[list]`) - Additional parameters required by the token endpoint of this authorization server represented as an array of JSON objects with name and value string properties, i.e. {"name" : "name value", "value": "a value"}.
            * `name` (`pulumi.Input[str]`) - body parameter name.
            * `value` (`pulumi.Input[str]`) - body parameter value.

          * `token_endpoint` (`pulumi.Input[str]`) - OAuth token endpoint. Contains absolute URI to entity being referenced.
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

            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if service_name is None:
                raise TypeError("Missing required property 'service_name'")
            __props__['service_name'] = service_name
            __props__['type'] = None
        super(AuthorizationServer, __self__).__init__(
            'azurerm:apimanagement/v20170301:AuthorizationServer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing AuthorizationServer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return AuthorizationServer(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
