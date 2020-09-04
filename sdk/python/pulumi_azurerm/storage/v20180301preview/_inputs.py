# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'CustomDomainArgs',
    'EncryptionArgs',
    'EncryptionServiceArgs',
    'EncryptionServicesArgs',
    'IPRuleArgs',
    'IdentityArgs',
    'KeyVaultPropertiesArgs',
    'NetworkRuleSetArgs',
    'RestrictionArgs',
    'SkuArgs',
    'VirtualNetworkRuleArgs',
]

@pulumi.input_type
class CustomDomainArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 use_sub_domain_name: Optional[pulumi.Input[bool]] = None):
        """
        The custom domain assigned to this storage account. This can be set via Update.
        :param pulumi.Input[str] name: Gets or sets the custom domain name assigned to the storage account. Name is the CNAME source.
        :param pulumi.Input[bool] use_sub_domain_name: Indicates whether indirect CName validation is enabled. Default value is false. This should only be set on updates.
        """
        pulumi.set(__self__, "name", name)
        if use_sub_domain_name is not None:
            pulumi.set(__self__, "use_sub_domain_name", use_sub_domain_name)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Gets or sets the custom domain name assigned to the storage account. Name is the CNAME source.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="useSubDomainName")
    def use_sub_domain_name(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether indirect CName validation is enabled. Default value is false. This should only be set on updates.
        """
        return pulumi.get(self, "use_sub_domain_name")

    @use_sub_domain_name.setter
    def use_sub_domain_name(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_sub_domain_name", value)


@pulumi.input_type
class EncryptionArgs:
    def __init__(__self__, *,
                 key_source: pulumi.Input[str],
                 key_vault_properties: Optional[pulumi.Input['KeyVaultPropertiesArgs']] = None,
                 services: Optional[pulumi.Input['EncryptionServicesArgs']] = None):
        """
        The encryption settings on the storage account.
        :param pulumi.Input[str] key_source: The encryption keySource (provider). Possible values (case-insensitive):  Microsoft.Storage, Microsoft.Keyvault
        :param pulumi.Input['KeyVaultPropertiesArgs'] key_vault_properties: Properties provided by key vault.
        :param pulumi.Input['EncryptionServicesArgs'] services: List of services which support encryption.
        """
        pulumi.set(__self__, "key_source", key_source)
        if key_vault_properties is not None:
            pulumi.set(__self__, "key_vault_properties", key_vault_properties)
        if services is not None:
            pulumi.set(__self__, "services", services)

    @property
    @pulumi.getter(name="keySource")
    def key_source(self) -> pulumi.Input[str]:
        """
        The encryption keySource (provider). Possible values (case-insensitive):  Microsoft.Storage, Microsoft.Keyvault
        """
        return pulumi.get(self, "key_source")

    @key_source.setter
    def key_source(self, value: pulumi.Input[str]):
        pulumi.set(self, "key_source", value)

    @property
    @pulumi.getter(name="keyVaultProperties")
    def key_vault_properties(self) -> Optional[pulumi.Input['KeyVaultPropertiesArgs']]:
        """
        Properties provided by key vault.
        """
        return pulumi.get(self, "key_vault_properties")

    @key_vault_properties.setter
    def key_vault_properties(self, value: Optional[pulumi.Input['KeyVaultPropertiesArgs']]):
        pulumi.set(self, "key_vault_properties", value)

    @property
    @pulumi.getter
    def services(self) -> Optional[pulumi.Input['EncryptionServicesArgs']]:
        """
        List of services which support encryption.
        """
        return pulumi.get(self, "services")

    @services.setter
    def services(self, value: Optional[pulumi.Input['EncryptionServicesArgs']]):
        pulumi.set(self, "services", value)


@pulumi.input_type
class EncryptionServiceArgs:
    def __init__(__self__, *,
                 enabled: Optional[pulumi.Input[bool]] = None):
        """
        A service that allows server-side encryption to be used.
        :param pulumi.Input[bool] enabled: A boolean indicating whether or not the service encrypts the data as it is stored.
        """
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        A boolean indicating whether or not the service encrypts the data as it is stored.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)


@pulumi.input_type
class EncryptionServicesArgs:
    def __init__(__self__, *,
                 blob: Optional[pulumi.Input['EncryptionServiceArgs']] = None,
                 file: Optional[pulumi.Input['EncryptionServiceArgs']] = None):
        """
        A list of services that support encryption.
        :param pulumi.Input['EncryptionServiceArgs'] blob: The encryption function of the blob storage service.
        :param pulumi.Input['EncryptionServiceArgs'] file: The encryption function of the file storage service.
        """
        if blob is not None:
            pulumi.set(__self__, "blob", blob)
        if file is not None:
            pulumi.set(__self__, "file", file)

    @property
    @pulumi.getter
    def blob(self) -> Optional[pulumi.Input['EncryptionServiceArgs']]:
        """
        The encryption function of the blob storage service.
        """
        return pulumi.get(self, "blob")

    @blob.setter
    def blob(self, value: Optional[pulumi.Input['EncryptionServiceArgs']]):
        pulumi.set(self, "blob", value)

    @property
    @pulumi.getter
    def file(self) -> Optional[pulumi.Input['EncryptionServiceArgs']]:
        """
        The encryption function of the file storage service.
        """
        return pulumi.get(self, "file")

    @file.setter
    def file(self, value: Optional[pulumi.Input['EncryptionServiceArgs']]):
        pulumi.set(self, "file", value)


@pulumi.input_type
class IPRuleArgs:
    def __init__(__self__, *,
                 i_p_address_or_range: pulumi.Input[str],
                 action: Optional[pulumi.Input[str]] = None):
        """
        IP rule with specific IP or IP range in CIDR format.
        :param pulumi.Input[str] i_p_address_or_range: Specifies the IP or IP range in CIDR format. Only IPV4 address is allowed.
        :param pulumi.Input[str] action: The action of IP ACL rule.
        """
        pulumi.set(__self__, "i_p_address_or_range", i_p_address_or_range)
        if action is not None:
            pulumi.set(__self__, "action", action)

    @property
    @pulumi.getter(name="iPAddressOrRange")
    def i_p_address_or_range(self) -> pulumi.Input[str]:
        """
        Specifies the IP or IP range in CIDR format. Only IPV4 address is allowed.
        """
        return pulumi.get(self, "i_p_address_or_range")

    @i_p_address_or_range.setter
    def i_p_address_or_range(self, value: pulumi.Input[str]):
        pulumi.set(self, "i_p_address_or_range", value)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        The action of IP ACL rule.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)


@pulumi.input_type
class IdentityArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str]):
        """
        Identity for the resource.
        :param pulumi.Input[str] type: The identity type.
        """
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The identity type.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class KeyVaultPropertiesArgs:
    def __init__(__self__, *,
                 key_name: Optional[pulumi.Input[str]] = None,
                 key_vault_uri: Optional[pulumi.Input[str]] = None,
                 key_version: Optional[pulumi.Input[str]] = None):
        """
        Properties of key vault.
        :param pulumi.Input[str] key_name: The name of KeyVault key.
        :param pulumi.Input[str] key_vault_uri: The Uri of KeyVault.
        :param pulumi.Input[str] key_version: The version of KeyVault key.
        """
        if key_name is not None:
            pulumi.set(__self__, "key_name", key_name)
        if key_vault_uri is not None:
            pulumi.set(__self__, "key_vault_uri", key_vault_uri)
        if key_version is not None:
            pulumi.set(__self__, "key_version", key_version)

    @property
    @pulumi.getter(name="keyName")
    def key_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of KeyVault key.
        """
        return pulumi.get(self, "key_name")

    @key_name.setter
    def key_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_name", value)

    @property
    @pulumi.getter(name="keyVaultUri")
    def key_vault_uri(self) -> Optional[pulumi.Input[str]]:
        """
        The Uri of KeyVault.
        """
        return pulumi.get(self, "key_vault_uri")

    @key_vault_uri.setter
    def key_vault_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_vault_uri", value)

    @property
    @pulumi.getter(name="keyVersion")
    def key_version(self) -> Optional[pulumi.Input[str]]:
        """
        The version of KeyVault key.
        """
        return pulumi.get(self, "key_version")

    @key_version.setter
    def key_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_version", value)


@pulumi.input_type
class NetworkRuleSetArgs:
    def __init__(__self__, *,
                 default_action: pulumi.Input[str],
                 bypass: Optional[pulumi.Input[str]] = None,
                 ip_rules: Optional[pulumi.Input[List[pulumi.Input['IPRuleArgs']]]] = None,
                 virtual_network_rules: Optional[pulumi.Input[List[pulumi.Input['VirtualNetworkRuleArgs']]]] = None):
        """
        Network rule set
        :param pulumi.Input[str] default_action: Specifies the default action of allow or deny when no other rules match.
        :param pulumi.Input[str] bypass: Specifies whether traffic is bypassed for Logging/Metrics/AzureServices. Possible values are any combination of Logging|Metrics|AzureServices (For example, "Logging, Metrics"), or None to bypass none of those traffics.
        :param pulumi.Input[List[pulumi.Input['IPRuleArgs']]] ip_rules: Sets the IP ACL rules
        :param pulumi.Input[List[pulumi.Input['VirtualNetworkRuleArgs']]] virtual_network_rules: Sets the virtual network rules
        """
        pulumi.set(__self__, "default_action", default_action)
        if bypass is not None:
            pulumi.set(__self__, "bypass", bypass)
        if ip_rules is not None:
            pulumi.set(__self__, "ip_rules", ip_rules)
        if virtual_network_rules is not None:
            pulumi.set(__self__, "virtual_network_rules", virtual_network_rules)

    @property
    @pulumi.getter(name="defaultAction")
    def default_action(self) -> pulumi.Input[str]:
        """
        Specifies the default action of allow or deny when no other rules match.
        """
        return pulumi.get(self, "default_action")

    @default_action.setter
    def default_action(self, value: pulumi.Input[str]):
        pulumi.set(self, "default_action", value)

    @property
    @pulumi.getter
    def bypass(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies whether traffic is bypassed for Logging/Metrics/AzureServices. Possible values are any combination of Logging|Metrics|AzureServices (For example, "Logging, Metrics"), or None to bypass none of those traffics.
        """
        return pulumi.get(self, "bypass")

    @bypass.setter
    def bypass(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bypass", value)

    @property
    @pulumi.getter(name="ipRules")
    def ip_rules(self) -> Optional[pulumi.Input[List[pulumi.Input['IPRuleArgs']]]]:
        """
        Sets the IP ACL rules
        """
        return pulumi.get(self, "ip_rules")

    @ip_rules.setter
    def ip_rules(self, value: Optional[pulumi.Input[List[pulumi.Input['IPRuleArgs']]]]):
        pulumi.set(self, "ip_rules", value)

    @property
    @pulumi.getter(name="virtualNetworkRules")
    def virtual_network_rules(self) -> Optional[pulumi.Input[List[pulumi.Input['VirtualNetworkRuleArgs']]]]:
        """
        Sets the virtual network rules
        """
        return pulumi.get(self, "virtual_network_rules")

    @virtual_network_rules.setter
    def virtual_network_rules(self, value: Optional[pulumi.Input[List[pulumi.Input['VirtualNetworkRuleArgs']]]]):
        pulumi.set(self, "virtual_network_rules", value)


@pulumi.input_type
class RestrictionArgs:
    def __init__(__self__, *,
                 reason_code: Optional[pulumi.Input[str]] = None):
        """
        The restriction because of which SKU cannot be used.
        :param pulumi.Input[str] reason_code: The reason for the restriction. As of now this can be "QuotaId" or "NotAvailableForSubscription". Quota Id is set when the SKU has requiredQuotas parameter as the subscription does not belong to that quota. The "NotAvailableForSubscription" is related to capacity at DC.
        """
        if reason_code is not None:
            pulumi.set(__self__, "reason_code", reason_code)

    @property
    @pulumi.getter(name="reasonCode")
    def reason_code(self) -> Optional[pulumi.Input[str]]:
        """
        The reason for the restriction. As of now this can be "QuotaId" or "NotAvailableForSubscription". Quota Id is set when the SKU has requiredQuotas parameter as the subscription does not belong to that quota. The "NotAvailableForSubscription" is related to capacity at DC.
        """
        return pulumi.get(self, "reason_code")

    @reason_code.setter
    def reason_code(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "reason_code", value)


@pulumi.input_type
class SkuArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 restrictions: Optional[pulumi.Input[List[pulumi.Input['RestrictionArgs']]]] = None):
        """
        The SKU of the storage account.
        :param pulumi.Input[str] name: Gets or sets the sku name. Required for account creation; optional for update. Note that in older versions, sku name was called accountType.
        :param pulumi.Input[List[pulumi.Input['RestrictionArgs']]] restrictions: The restrictions because of which SKU cannot be used. This is empty if there are no restrictions.
        """
        pulumi.set(__self__, "name", name)
        if restrictions is not None:
            pulumi.set(__self__, "restrictions", restrictions)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Gets or sets the sku name. Required for account creation; optional for update. Note that in older versions, sku name was called accountType.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def restrictions(self) -> Optional[pulumi.Input[List[pulumi.Input['RestrictionArgs']]]]:
        """
        The restrictions because of which SKU cannot be used. This is empty if there are no restrictions.
        """
        return pulumi.get(self, "restrictions")

    @restrictions.setter
    def restrictions(self, value: Optional[pulumi.Input[List[pulumi.Input['RestrictionArgs']]]]):
        pulumi.set(self, "restrictions", value)


@pulumi.input_type
class VirtualNetworkRuleArgs:
    def __init__(__self__, *,
                 virtual_network_resource_id: pulumi.Input[str],
                 action: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None):
        """
        Virtual Network rule.
        :param pulumi.Input[str] virtual_network_resource_id: Resource ID of a subnet, for example: /subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.Network/virtualNetworks/{vnetName}/subnets/{subnetName}.
        :param pulumi.Input[str] action: The action of virtual network rule.
        :param pulumi.Input[str] state: Gets the state of virtual network rule.
        """
        pulumi.set(__self__, "virtual_network_resource_id", virtual_network_resource_id)
        if action is not None:
            pulumi.set(__self__, "action", action)
        if state is not None:
            pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter(name="virtualNetworkResourceId")
    def virtual_network_resource_id(self) -> pulumi.Input[str]:
        """
        Resource ID of a subnet, for example: /subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.Network/virtualNetworks/{vnetName}/subnets/{subnetName}.
        """
        return pulumi.get(self, "virtual_network_resource_id")

    @virtual_network_resource_id.setter
    def virtual_network_resource_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "virtual_network_resource_id", value)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        The action of virtual network rule.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        Gets the state of virtual network rule.
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)


