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

__all__ = ['DomainService']


class DomainService(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 domain_security_settings: Optional[pulumi.Input[pulumi.InputType['DomainSecuritySettingsArgs']]] = None,
                 domain_service_name: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 filtered_sync: Optional[pulumi.Input[str]] = None,
                 ldaps_settings: Optional[pulumi.Input[pulumi.InputType['LdapsSettingsArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 notification_settings: Optional[pulumi.Input[pulumi.InputType['NotificationSettingsArgs']]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Domain service.

        ## Example Usage
        ### Create Domain Service

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        domain_service = azurerm.aad.v20170601.DomainService("domainService",
            domain_name="zdomain.zforest.com",
            domain_security_settings={
                "ntlmV1": "Enabled",
                "syncNtlmPasswords": "Enabled",
                "tlsV1": "Disabled",
            },
            domain_service_name="zdomain.zforest.com",
            filtered_sync="Enabled",
            ldaps_settings={
                "externalAccess": "Enabled",
                "ldaps": "Enabled",
                "pfxCertificate": "MIIDPDCCAiSgAwIBAgIQQUI9P6tq2p9OFIJa7DLNvTANBgkqhkiG9w0BAQsFADAgMR4w...",
                "pfxCertificatePassword": "Password01",
            },
            location="westus",
            notification_settings={
                "additionalRecipients": [
                    "jicha@microsoft.com",
                    "caalmont@microsoft.com",
                ],
                "notifyDcAdmins": "Enabled",
                "notifyGlobalAdmins": "Enabled",
            },
            resource_group_name="sva-tt-WUS",
            subnet_id="/subscriptions/1639790a-76a2-4ac4-98d9-8562f5dfcb4d/resourceGroups/Default-Networking/providers/Microsoft.Network/virtualNetworks/DCIaasTmpWusNet/subnets/Subnet-1")

        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain_name: The name of the Azure domain that the user would like to deploy Domain Services to.
        :param pulumi.Input[pulumi.InputType['DomainSecuritySettingsArgs']] domain_security_settings: DomainSecurity Settings
        :param pulumi.Input[str] domain_service_name: The name of the domain service.
        :param pulumi.Input[str] etag: Resource etag
        :param pulumi.Input[str] filtered_sync: Enabled or Disabled flag to turn on Group-based filtered sync
        :param pulumi.Input[pulumi.InputType['LdapsSettingsArgs']] ldaps_settings: Secure LDAP Settings
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[pulumi.InputType['NotificationSettingsArgs']] notification_settings: Notification Settings
        :param pulumi.Input[str] resource_group_name: The name of the resource group within the user's subscription. The name is case insensitive.
        :param pulumi.Input[str] subnet_id: The name of the virtual network that Domain Services will be deployed on. The id of the subnet that Domain Services will be deployed on. /virtualNetwork/vnetName/subnets/subnetName.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags
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

            __props__['domain_name'] = domain_name
            __props__['domain_security_settings'] = domain_security_settings
            if domain_service_name is None:
                raise TypeError("Missing required property 'domain_service_name'")
            __props__['domain_service_name'] = domain_service_name
            __props__['etag'] = etag
            __props__['filtered_sync'] = filtered_sync
            __props__['ldaps_settings'] = ldaps_settings
            __props__['location'] = location
            __props__['notification_settings'] = notification_settings
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['subnet_id'] = subnet_id
            __props__['tags'] = tags
            __props__['domain_controller_ip_address'] = None
            __props__['health_alerts'] = None
            __props__['health_last_evaluated'] = None
            __props__['health_monitors'] = None
            __props__['name'] = None
            __props__['provisioning_state'] = None
            __props__['service_status'] = None
            __props__['tenant_id'] = None
            __props__['type'] = None
            __props__['vnet_site_id'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:aad/latest:DomainService"), pulumi.Alias(type_="azurerm:aad/v20170101:DomainService"), pulumi.Alias(type_="azurerm:aad/v20200101:DomainService")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(DomainService, __self__).__init__(
            'azurerm:aad/v20170601:DomainService',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DomainService':
        """
        Get an existing DomainService resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return DomainService(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="domainControllerIpAddress")
    def domain_controller_ip_address(self) -> pulumi.Output[List[str]]:
        """
        List of Domain Controller IP Address
        """
        return pulumi.get(self, "domain_controller_ip_address")

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the Azure domain that the user would like to deploy Domain Services to.
        """
        return pulumi.get(self, "domain_name")

    @property
    @pulumi.getter(name="domainSecuritySettings")
    def domain_security_settings(self) -> pulumi.Output[Optional['outputs.DomainSecuritySettingsResponse']]:
        """
        DomainSecurity Settings
        """
        return pulumi.get(self, "domain_security_settings")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[Optional[str]]:
        """
        Resource etag
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="filteredSync")
    def filtered_sync(self) -> pulumi.Output[Optional[str]]:
        """
        Enabled or Disabled flag to turn on Group-based filtered sync
        """
        return pulumi.get(self, "filtered_sync")

    @property
    @pulumi.getter(name="healthAlerts")
    def health_alerts(self) -> pulumi.Output[List['outputs.HealthAlertResponse']]:
        """
        List of Domain Health Alerts
        """
        return pulumi.get(self, "health_alerts")

    @property
    @pulumi.getter(name="healthLastEvaluated")
    def health_last_evaluated(self) -> pulumi.Output[str]:
        """
        Last domain evaluation run DateTime
        """
        return pulumi.get(self, "health_last_evaluated")

    @property
    @pulumi.getter(name="healthMonitors")
    def health_monitors(self) -> pulumi.Output[List['outputs.HealthMonitorResponse']]:
        """
        List of Domain Health Monitors
        """
        return pulumi.get(self, "health_monitors")

    @property
    @pulumi.getter(name="ldapsSettings")
    def ldaps_settings(self) -> pulumi.Output[Optional['outputs.LdapsSettingsResponse']]:
        """
        Secure LDAP Settings
        """
        return pulumi.get(self, "ldaps_settings")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[Optional[str]]:
        """
        Resource location
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="notificationSettings")
    def notification_settings(self) -> pulumi.Output[Optional['outputs.NotificationSettingsResponse']]:
        """
        Notification Settings
        """
        return pulumi.get(self, "notification_settings")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        the current deployment or provisioning state, which only appears in the response.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="serviceStatus")
    def service_status(self) -> pulumi.Output[str]:
        """
        Status of Domain Service instance
        """
        return pulumi.get(self, "service_status")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the virtual network that Domain Services will be deployed on. The id of the subnet that Domain Services will be deployed on. /virtualNetwork/vnetName/subnets/subnetName.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Resource tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[str]:
        """
        Azure Active Directory tenant id
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Resource type
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="vnetSiteId")
    def vnet_site_id(self) -> pulumi.Output[str]:
        """
        Virtual network site id
        """
        return pulumi.get(self, "vnet_site_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

