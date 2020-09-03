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

__all__ = ['NotificationHub']


class NotificationHub(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 adm_credential: Optional[pulumi.Input[pulumi.InputType['AdmCredentialArgs']]] = None,
                 apns_credential: Optional[pulumi.Input[pulumi.InputType['ApnsCredentialArgs']]] = None,
                 authorization_rules: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['SharedAccessAuthorizationRulePropertiesArgs']]]]] = None,
                 baidu_credential: Optional[pulumi.Input[pulumi.InputType['BaiduCredentialArgs']]] = None,
                 gcm_credential: Optional[pulumi.Input[pulumi.InputType['GcmCredentialArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 mpns_credential: Optional[pulumi.Input[pulumi.InputType['MpnsCredentialArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace_name: Optional[pulumi.Input[str]] = None,
                 notification_hub_name: Optional[pulumi.Input[str]] = None,
                 registration_ttl: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sku: Optional[pulumi.Input[pulumi.InputType['SkuArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 wns_credential: Optional[pulumi.Input[pulumi.InputType['WnsCredentialArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Description of a NotificationHub Resource.

        ## Example Usage
        ### NotificationHubCreate

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        notification_hub = azurerm.notificationhubs.v20170401.NotificationHub("notificationHub",
            location="eastus",
            namespace_name="nh-sdk-ns",
            notification_hub_name="nh-sdk-hub",
            resource_group_name="5ktrial")

        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['AdmCredentialArgs']] adm_credential: The AdmCredential of the created NotificationHub
        :param pulumi.Input[pulumi.InputType['ApnsCredentialArgs']] apns_credential: The ApnsCredential of the created NotificationHub
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['SharedAccessAuthorizationRulePropertiesArgs']]]] authorization_rules: The AuthorizationRules of the created NotificationHub
        :param pulumi.Input[pulumi.InputType['BaiduCredentialArgs']] baidu_credential: The BaiduCredential of the created NotificationHub
        :param pulumi.Input[pulumi.InputType['GcmCredentialArgs']] gcm_credential: The GcmCredential of the created NotificationHub
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[pulumi.InputType['MpnsCredentialArgs']] mpns_credential: The MpnsCredential of the created NotificationHub
        :param pulumi.Input[str] name: The NotificationHub name.
        :param pulumi.Input[str] namespace_name: The namespace name.
        :param pulumi.Input[str] notification_hub_name: The notification hub name.
        :param pulumi.Input[str] registration_ttl: The RegistrationTtl of the created NotificationHub
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[pulumi.InputType['SkuArgs']] sku: The sku of the created namespace
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags
        :param pulumi.Input[pulumi.InputType['WnsCredentialArgs']] wns_credential: The WnsCredential of the created NotificationHub
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
            __props__['location'] = location
            __props__['mpns_credential'] = mpns_credential
            __props__['name'] = name
            if namespace_name is None:
                raise TypeError("Missing required property 'namespace_name'")
            __props__['namespace_name'] = namespace_name
            if notification_hub_name is None:
                raise TypeError("Missing required property 'notification_hub_name'")
            __props__['notification_hub_name'] = notification_hub_name
            __props__['registration_ttl'] = registration_ttl
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['wns_credential'] = wns_credential
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:notificationhubs/latest:NotificationHub"), pulumi.Alias(type_="azurerm:notificationhubs/v20140901:NotificationHub"), pulumi.Alias(type_="azurerm:notificationhubs/v20160301:NotificationHub")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(NotificationHub, __self__).__init__(
            'azurerm:notificationhubs/v20170401:NotificationHub',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'NotificationHub':
        """
        Get an existing NotificationHub resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return NotificationHub(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="admCredential")
    def adm_credential(self) -> pulumi.Output[Optional['outputs.AdmCredentialResponse']]:
        """
        The AdmCredential of the created NotificationHub
        """
        return pulumi.get(self, "adm_credential")

    @property
    @pulumi.getter(name="apnsCredential")
    def apns_credential(self) -> pulumi.Output[Optional['outputs.ApnsCredentialResponse']]:
        """
        The ApnsCredential of the created NotificationHub
        """
        return pulumi.get(self, "apns_credential")

    @property
    @pulumi.getter(name="authorizationRules")
    def authorization_rules(self) -> pulumi.Output[Optional[List['outputs.SharedAccessAuthorizationRulePropertiesResponse']]]:
        """
        The AuthorizationRules of the created NotificationHub
        """
        return pulumi.get(self, "authorization_rules")

    @property
    @pulumi.getter(name="baiduCredential")
    def baidu_credential(self) -> pulumi.Output[Optional['outputs.BaiduCredentialResponse']]:
        """
        The BaiduCredential of the created NotificationHub
        """
        return pulumi.get(self, "baidu_credential")

    @property
    @pulumi.getter(name="gcmCredential")
    def gcm_credential(self) -> pulumi.Output[Optional['outputs.GcmCredentialResponse']]:
        """
        The GcmCredential of the created NotificationHub
        """
        return pulumi.get(self, "gcm_credential")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[Optional[str]]:
        """
        Resource location
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="mpnsCredential")
    def mpns_credential(self) -> pulumi.Output[Optional['outputs.MpnsCredentialResponse']]:
        """
        The MpnsCredential of the created NotificationHub
        """
        return pulumi.get(self, "mpns_credential")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="registrationTtl")
    def registration_ttl(self) -> pulumi.Output[Optional[str]]:
        """
        The RegistrationTtl of the created NotificationHub
        """
        return pulumi.get(self, "registration_ttl")

    @property
    @pulumi.getter
    def sku(self) -> pulumi.Output[Optional['outputs.SkuResponse']]:
        """
        The sku of the created namespace
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Resource tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Resource type
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="wnsCredential")
    def wns_credential(self) -> pulumi.Output[Optional['outputs.WnsCredentialResponse']]:
        """
        The WnsCredential of the created NotificationHub
        """
        return pulumi.get(self, "wns_credential")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

