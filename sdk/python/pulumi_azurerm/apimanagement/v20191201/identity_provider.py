# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class IdentityProvider(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    Resource name.
    """
    properties: pulumi.Output[dict]
    """
    Identity Provider contract properties.
      * `allowed_tenants` (`list`) - List of Allowed Tenants when configuring Azure Active Directory login.
      * `authority` (`str`) - OpenID Connect discovery endpoint hostname for AAD or AAD B2C.
      * `client_id` (`str`) - Client Id of the Application in the external Identity Provider. It is App ID for Facebook login, Client ID for Google login, App ID for Microsoft.
      * `client_secret` (`str`) - Client secret of the Application in external Identity Provider, used to authenticate login request. For example, it is App Secret for Facebook login, API Key for Google login, Public Key for Microsoft. This property will not be filled on 'GET' operations! Use '/listSecrets' POST request to get the value.
      * `password_reset_policy_name` (`str`) - Password Reset Policy Name. Only applies to AAD B2C Identity Provider.
      * `profile_editing_policy_name` (`str`) - Profile Editing Policy Name. Only applies to AAD B2C Identity Provider.
      * `signin_policy_name` (`str`) - Signin Policy Name. Only applies to AAD B2C Identity Provider.
      * `signin_tenant` (`str`) - The TenantId to use instead of Common when logging into Active Directory
      * `signup_policy_name` (`str`) - Signup Policy Name. Only applies to AAD B2C Identity Provider.
      * `type` (`str`) - Identity Provider Type identifier.
    """
    type: pulumi.Output[str]
    """
    Resource type for API Management resource.
    """
    def __init__(__self__, resource_name, opts=None, name=None, properties=None, resource_group_name=None, service_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Identity Provider details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Identity Provider Type identifier.
        :param pulumi.Input[dict] properties: Identity Provider contract properties.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] service_name: The name of the API Management service.

        The **properties** object supports the following:

          * `allowed_tenants` (`pulumi.Input[list]`) - List of Allowed Tenants when configuring Azure Active Directory login.
          * `authority` (`pulumi.Input[str]`) - OpenID Connect discovery endpoint hostname for AAD or AAD B2C.
          * `client_id` (`pulumi.Input[str]`) - Client Id of the Application in the external Identity Provider. It is App ID for Facebook login, Client ID for Google login, App ID for Microsoft.
          * `client_secret` (`pulumi.Input[str]`) - Client secret of the Application in external Identity Provider, used to authenticate login request. For example, it is App Secret for Facebook login, API Key for Google login, Public Key for Microsoft. This property will not be filled on 'GET' operations! Use '/listSecrets' POST request to get the value.
          * `password_reset_policy_name` (`pulumi.Input[str]`) - Password Reset Policy Name. Only applies to AAD B2C Identity Provider.
          * `profile_editing_policy_name` (`pulumi.Input[str]`) - Profile Editing Policy Name. Only applies to AAD B2C Identity Provider.
          * `signin_policy_name` (`pulumi.Input[str]`) - Signin Policy Name. Only applies to AAD B2C Identity Provider.
          * `signin_tenant` (`pulumi.Input[str]`) - The TenantId to use instead of Common when logging into Active Directory
          * `signup_policy_name` (`pulumi.Input[str]`) - Signup Policy Name. Only applies to AAD B2C Identity Provider.
          * `type` (`pulumi.Input[str]`) - Identity Provider Type identifier.
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
        super(IdentityProvider, __self__).__init__(
            'azurerm:apimanagement/v20191201:IdentityProvider',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing IdentityProvider resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return IdentityProvider(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
