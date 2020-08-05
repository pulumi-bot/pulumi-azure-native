# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class NotificationHub(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    Resource location
    """
    name: pulumi.Output[str]
    """
    Resource name
    """
    properties: pulumi.Output[dict]
    """
    Properties of the NotificationHub.
      * `adm_credential` (`dict`) - The AdmCredential of the created NotificationHub
        * `properties` (`dict`) - Properties of NotificationHub AdmCredential.
          * `auth_token_url` (`str`) - The URL of the authorization token.
          * `client_id` (`str`) - The client identifier.
          * `client_secret` (`str`) - The credential secret access key.

      * `apns_credential` (`dict`) - The ApnsCredential of the created NotificationHub
        * `properties` (`dict`) - Properties of NotificationHub ApnsCredential.
          * `apns_certificate` (`str`) - The APNS certificate.
          * `certificate_key` (`str`) - The certificate key.
          * `endpoint` (`str`) - The endpoint of this credential.
          * `thumbprint` (`str`) - The APNS certificate Thumbprint

      * `authorization_rules` (`list`) - The AuthorizationRules of the created NotificationHub
        * `rights` (`list`) - The rights associated with the rule.

      * `baidu_credential` (`dict`) - The BaiduCredential of the created NotificationHub
        * `properties` (`dict`) - Properties of NotificationHub BaiduCredential.
          * `baidu_api_key` (`str`) - Baidu Api Key.
          * `baidu_end_point` (`str`) - Baidu Endpoint.
          * `baidu_secret_key` (`str`) - Baidu Secret Key

      * `gcm_credential` (`dict`) - The GcmCredential of the created NotificationHub
        * `properties` (`dict`) - Properties of NotificationHub GcmCredential.
          * `gcm_endpoint` (`str`) - The GCM endpoint.
          * `google_api_key` (`str`) - The Google API key.

      * `mpns_credential` (`dict`) - The MpnsCredential of the created NotificationHub
        * `properties` (`dict`) - Properties of NotificationHub MpnsCredential.
          * `certificate_key` (`str`) - The certificate key for this credential.
          * `mpns_certificate` (`str`) - The MPNS certificate.
          * `thumbprint` (`str`) - The MPNS certificate Thumbprint

      * `name` (`str`) - The NotificationHub name.
      * `registration_ttl` (`str`) - The RegistrationTtl of the created NotificationHub
      * `wns_credential` (`dict`) - The WnsCredential of the created NotificationHub
        * `properties` (`dict`) - Properties of NotificationHub WnsCredential.
          * `package_sid` (`str`) - The package ID for this credential.
          * `secret_key` (`str`) - The secret key.
          * `windows_live_endpoint` (`str`) - The Windows Live endpoint.
    """
    sku: pulumi.Output[dict]
    """
    The sku of the created namespace
      * `capacity` (`float`) - The capacity of the resource
      * `family` (`str`) - The Sku Family
      * `name` (`str`) - Name of the notification hub sku
      * `size` (`str`) - The Sku size
      * `tier` (`str`) - The tier of particular sku
    """
    tags: pulumi.Output[dict]
    """
    Resource tags
    """
    type: pulumi.Output[str]
    """
    Resource type
    """
    def __init__(__self__, resource_name, opts=None, adm_credential=None, apns_credential=None, authorization_rules=None, baidu_credential=None, gcm_credential=None, location=None, mpns_credential=None, name=None, namespace_name=None, registration_ttl=None, resource_group_name=None, sku=None, tags=None, wns_credential=None, __props__=None, __name__=None, __opts__=None):
        """
        Description of a NotificationHub Resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] adm_credential: The AdmCredential of the created NotificationHub
        :param pulumi.Input[dict] apns_credential: The ApnsCredential of the created NotificationHub
        :param pulumi.Input[list] authorization_rules: The AuthorizationRules of the created NotificationHub
        :param pulumi.Input[dict] baidu_credential: The BaiduCredential of the created NotificationHub
        :param pulumi.Input[dict] gcm_credential: The GcmCredential of the created NotificationHub
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[dict] mpns_credential: The MpnsCredential of the created NotificationHub
        :param pulumi.Input[str] name: The notification hub name.
        :param pulumi.Input[str] namespace_name: The namespace name.
        :param pulumi.Input[str] registration_ttl: The RegistrationTtl of the created NotificationHub
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[dict] sku: The sku of the created namespace
        :param pulumi.Input[dict] tags: Resource tags
        :param pulumi.Input[dict] wns_credential: The WnsCredential of the created NotificationHub

        The **adm_credential** object supports the following:

          * `auth_token_url` (`pulumi.Input[str]`) - The URL of the authorization token.
          * `client_id` (`pulumi.Input[str]`) - The client identifier.
          * `client_secret` (`pulumi.Input[str]`) - The credential secret access key.

        The **apns_credential** object supports the following:

          * `apns_certificate` (`pulumi.Input[str]`) - The APNS certificate.
          * `certificate_key` (`pulumi.Input[str]`) - The certificate key.
          * `endpoint` (`pulumi.Input[str]`) - The endpoint of this credential.
          * `thumbprint` (`pulumi.Input[str]`) - The APNS certificate Thumbprint

        The **authorization_rules** object supports the following:

          * `rights` (`pulumi.Input[list]`) - The rights associated with the rule.

        The **baidu_credential** object supports the following:

          * `baidu_api_key` (`pulumi.Input[str]`) - Baidu Api Key.
          * `baidu_end_point` (`pulumi.Input[str]`) - Baidu Endpoint.
          * `baidu_secret_key` (`pulumi.Input[str]`) - Baidu Secret Key

        The **gcm_credential** object supports the following:

          * `gcm_endpoint` (`pulumi.Input[str]`) - The GCM endpoint.
          * `google_api_key` (`pulumi.Input[str]`) - The Google API key.

        The **mpns_credential** object supports the following:

          * `certificate_key` (`pulumi.Input[str]`) - The certificate key for this credential.
          * `mpns_certificate` (`pulumi.Input[str]`) - The MPNS certificate.
          * `thumbprint` (`pulumi.Input[str]`) - The MPNS certificate Thumbprint

        The **sku** object supports the following:

          * `capacity` (`pulumi.Input[float]`) - The capacity of the resource
          * `family` (`pulumi.Input[str]`) - The Sku Family
          * `name` (`pulumi.Input[str]`) - Name of the notification hub sku
          * `size` (`pulumi.Input[str]`) - The Sku size
          * `tier` (`pulumi.Input[str]`) - The tier of particular sku

        The **wns_credential** object supports the following:

          * `package_sid` (`pulumi.Input[str]`) - The package ID for this credential.
          * `secret_key` (`pulumi.Input[str]`) - The secret key.
          * `windows_live_endpoint` (`pulumi.Input[str]`) - The Windows Live endpoint.
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

            __props__['adm_credential'] = adm_credential
            __props__['apns_credential'] = apns_credential
            __props__['authorization_rules'] = authorization_rules
            __props__['baidu_credential'] = baidu_credential
            __props__['gcm_credential'] = gcm_credential
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            __props__['mpns_credential'] = mpns_credential
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if namespace_name is None:
                raise TypeError("Missing required property 'namespace_name'")
            __props__['namespace_name'] = namespace_name
            __props__['registration_ttl'] = registration_ttl
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['wns_credential'] = wns_credential
            __props__['properties'] = None
            __props__['type'] = None
        super(NotificationHub, __self__).__init__(
            'azurerm:notificationhubs/v20160301:NotificationHub',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing NotificationHub resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return NotificationHub(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
