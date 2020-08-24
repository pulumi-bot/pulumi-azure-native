# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'AccessPolicyEntryArgs',
    'IPRuleArgs',
    'NetworkRuleSetArgs',
    'PermissionsArgs',
    'PrivateLinkServiceConnectionStateArgs',
    'SkuArgs',
    'VaultPropertiesArgs',
    'VirtualNetworkRuleArgs',
]

@pulumi.input_type
class AccessPolicyEntryArgs:
    def __init__(__self__, *,
                 object_id: pulumi.Input[str],
                 permissions: pulumi.Input['PermissionsArgs'],
                 tenant_id: pulumi.Input[str],
                 application_id: Optional[pulumi.Input[str]] = None):
        """
        An identity that have access to the key vault. All identities in the array must use the same tenant ID as the key vault's tenant ID.
        :param pulumi.Input[str] object_id: The object ID of a user, service principal or security group in the Azure Active Directory tenant for the vault. The object ID must be unique for the list of access policies.
        :param pulumi.Input['PermissionsArgs'] permissions: Permissions the identity has for keys, secrets and certificates.
        :param pulumi.Input[str] tenant_id: The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.
        :param pulumi.Input[str] application_id:  Application ID of the client making request on behalf of a principal
        """
        pulumi.set(__self__, "object_id", object_id)
        pulumi.set(__self__, "permissions", permissions)
        pulumi.set(__self__, "tenant_id", tenant_id)
        if application_id is not None:
            pulumi.set(__self__, "application_id", application_id)

    @property
    @pulumi.getter(name="objectId")
    def object_id(self) -> pulumi.Input[str]:
        """
        The object ID of a user, service principal or security group in the Azure Active Directory tenant for the vault. The object ID must be unique for the list of access policies.
        """
        return pulumi.get(self, "object_id")

    @object_id.setter
    def object_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "object_id", value)

    @property
    @pulumi.getter
    def permissions(self) -> pulumi.Input['PermissionsArgs']:
        """
        Permissions the identity has for keys, secrets and certificates.
        """
        return pulumi.get(self, "permissions")

    @permissions.setter
    def permissions(self, value: pulumi.Input['PermissionsArgs']):
        pulumi.set(self, "permissions", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Input[str]:
        """
        The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> Optional[pulumi.Input[str]]:
        """
         Application ID of the client making request on behalf of a principal
        """
        return pulumi.get(self, "application_id")

    @application_id.setter
    def application_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "application_id", value)


@pulumi.input_type
class IPRuleArgs:
    def __init__(__self__, *,
                 value: pulumi.Input[str]):
        """
        A rule governing the accessibility of a vault from a specific ip address or ip range.
        :param pulumi.Input[str] value: An IPv4 address range in CIDR notation, such as '124.56.78.91' (simple IP address) or '124.56.78.0/24' (all addresses that start with 124.56.78).
        """
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        An IPv4 address range in CIDR notation, such as '124.56.78.91' (simple IP address) or '124.56.78.0/24' (all addresses that start with 124.56.78).
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class NetworkRuleSetArgs:
    def __init__(__self__, *,
                 bypass: Optional[pulumi.Input[str]] = None,
                 default_action: Optional[pulumi.Input[str]] = None,
                 ip_rules: Optional[pulumi.Input[List[pulumi.Input['IPRuleArgs']]]] = None,
                 virtual_network_rules: Optional[pulumi.Input[List[pulumi.Input['VirtualNetworkRuleArgs']]]] = None):
        """
        A set of rules governing the network accessibility of a vault.
        :param pulumi.Input[str] bypass: Tells what traffic can bypass network rules. This can be 'AzureServices' or 'None'.  If not specified the default is 'AzureServices'.
        :param pulumi.Input[str] default_action: The default action when no rule from ipRules and from virtualNetworkRules match. This is only used after the bypass property has been evaluated.
        :param pulumi.Input[List[pulumi.Input['IPRuleArgs']]] ip_rules: The list of IP address rules.
        :param pulumi.Input[List[pulumi.Input['VirtualNetworkRuleArgs']]] virtual_network_rules: The list of virtual network rules.
        """
        if bypass is not None:
            pulumi.set(__self__, "bypass", bypass)
        if default_action is not None:
            pulumi.set(__self__, "default_action", default_action)
        if ip_rules is not None:
            pulumi.set(__self__, "ip_rules", ip_rules)
        if virtual_network_rules is not None:
            pulumi.set(__self__, "virtual_network_rules", virtual_network_rules)

    @property
    @pulumi.getter
    def bypass(self) -> Optional[pulumi.Input[str]]:
        """
        Tells what traffic can bypass network rules. This can be 'AzureServices' or 'None'.  If not specified the default is 'AzureServices'.
        """
        return pulumi.get(self, "bypass")

    @bypass.setter
    def bypass(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bypass", value)

    @property
    @pulumi.getter(name="defaultAction")
    def default_action(self) -> Optional[pulumi.Input[str]]:
        """
        The default action when no rule from ipRules and from virtualNetworkRules match. This is only used after the bypass property has been evaluated.
        """
        return pulumi.get(self, "default_action")

    @default_action.setter
    def default_action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_action", value)

    @property
    @pulumi.getter(name="ipRules")
    def ip_rules(self) -> Optional[pulumi.Input[List[pulumi.Input['IPRuleArgs']]]]:
        """
        The list of IP address rules.
        """
        return pulumi.get(self, "ip_rules")

    @ip_rules.setter
    def ip_rules(self, value: Optional[pulumi.Input[List[pulumi.Input['IPRuleArgs']]]]):
        pulumi.set(self, "ip_rules", value)

    @property
    @pulumi.getter(name="virtualNetworkRules")
    def virtual_network_rules(self) -> Optional[pulumi.Input[List[pulumi.Input['VirtualNetworkRuleArgs']]]]:
        """
        The list of virtual network rules.
        """
        return pulumi.get(self, "virtual_network_rules")

    @virtual_network_rules.setter
    def virtual_network_rules(self, value: Optional[pulumi.Input[List[pulumi.Input['VirtualNetworkRuleArgs']]]]):
        pulumi.set(self, "virtual_network_rules", value)


@pulumi.input_type
class PermissionsArgs:
    def __init__(__self__, *,
                 certificates: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 keys: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 secrets: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 storage: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None):
        """
        Permissions the identity has for keys, secrets, certificates and storage.
        :param pulumi.Input[List[pulumi.Input[str]]] certificates: Permissions to certificates
        :param pulumi.Input[List[pulumi.Input[str]]] keys: Permissions to keys
        :param pulumi.Input[List[pulumi.Input[str]]] secrets: Permissions to secrets
        :param pulumi.Input[List[pulumi.Input[str]]] storage: Permissions to storage accounts
        """
        if certificates is not None:
            pulumi.set(__self__, "certificates", certificates)
        if keys is not None:
            pulumi.set(__self__, "keys", keys)
        if secrets is not None:
            pulumi.set(__self__, "secrets", secrets)
        if storage is not None:
            pulumi.set(__self__, "storage", storage)

    @property
    @pulumi.getter
    def certificates(self) -> Optional[pulumi.Input[List[pulumi.Input[str]]]]:
        """
        Permissions to certificates
        """
        return pulumi.get(self, "certificates")

    @certificates.setter
    def certificates(self, value: Optional[pulumi.Input[List[pulumi.Input[str]]]]):
        pulumi.set(self, "certificates", value)

    @property
    @pulumi.getter
    def keys(self) -> Optional[pulumi.Input[List[pulumi.Input[str]]]]:
        """
        Permissions to keys
        """
        return pulumi.get(self, "keys")

    @keys.setter
    def keys(self, value: Optional[pulumi.Input[List[pulumi.Input[str]]]]):
        pulumi.set(self, "keys", value)

    @property
    @pulumi.getter
    def secrets(self) -> Optional[pulumi.Input[List[pulumi.Input[str]]]]:
        """
        Permissions to secrets
        """
        return pulumi.get(self, "secrets")

    @secrets.setter
    def secrets(self, value: Optional[pulumi.Input[List[pulumi.Input[str]]]]):
        pulumi.set(self, "secrets", value)

    @property
    @pulumi.getter
    def storage(self) -> Optional[pulumi.Input[List[pulumi.Input[str]]]]:
        """
        Permissions to storage accounts
        """
        return pulumi.get(self, "storage")

    @storage.setter
    def storage(self, value: Optional[pulumi.Input[List[pulumi.Input[str]]]]):
        pulumi.set(self, "storage", value)


@pulumi.input_type
class PrivateLinkServiceConnectionStateArgs:
    def __init__(__self__, *,
                 action_required: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        An object that represents the approval state of the private link connection.
        :param pulumi.Input[str] action_required: A message indicating if changes on the service provider require any updates on the consumer.
        :param pulumi.Input[str] description: The reason for approval or rejection.
        :param pulumi.Input[str] status: Indicates whether the connection has been approved, rejected or removed by the key vault owner.
        """
        if action_required is not None:
            pulumi.set(__self__, "action_required", action_required)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="actionRequired")
    def action_required(self) -> Optional[pulumi.Input[str]]:
        """
        A message indicating if changes on the service provider require any updates on the consumer.
        """
        return pulumi.get(self, "action_required")

    @action_required.setter
    def action_required(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action_required", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The reason for approval or rejection.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Indicates whether the connection has been approved, rejected or removed by the key vault owner.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


@pulumi.input_type
class SkuArgs:
    def __init__(__self__, *,
                 family: pulumi.Input[str],
                 name: pulumi.Input[str]):
        """
        SKU details
        :param pulumi.Input[str] family: SKU family name
        :param pulumi.Input[str] name: SKU name to specify whether the key vault is a standard vault or a premium vault.
        """
        pulumi.set(__self__, "family", family)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def family(self) -> pulumi.Input[str]:
        """
        SKU family name
        """
        return pulumi.get(self, "family")

    @family.setter
    def family(self, value: pulumi.Input[str]):
        pulumi.set(self, "family", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        SKU name to specify whether the key vault is a standard vault or a premium vault.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class VaultPropertiesArgs:
    def __init__(__self__, *,
                 sku: pulumi.Input['SkuArgs'],
                 tenant_id: pulumi.Input[str],
                 access_policies: Optional[pulumi.Input[List[pulumi.Input['AccessPolicyEntryArgs']]]] = None,
                 create_mode: Optional[pulumi.Input[str]] = None,
                 enable_purge_protection: Optional[pulumi.Input[bool]] = None,
                 enable_rbac_authorization: Optional[pulumi.Input[bool]] = None,
                 enable_soft_delete: Optional[pulumi.Input[bool]] = None,
                 enabled_for_deployment: Optional[pulumi.Input[bool]] = None,
                 enabled_for_disk_encryption: Optional[pulumi.Input[bool]] = None,
                 enabled_for_template_deployment: Optional[pulumi.Input[bool]] = None,
                 network_acls: Optional[pulumi.Input['NetworkRuleSetArgs']] = None,
                 soft_delete_retention_in_days: Optional[pulumi.Input[float]] = None,
                 vault_uri: Optional[pulumi.Input[str]] = None):
        """
        Properties of the vault
        :param pulumi.Input['SkuArgs'] sku: SKU details
        :param pulumi.Input[str] tenant_id: The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.
        :param pulumi.Input[List[pulumi.Input['AccessPolicyEntryArgs']]] access_policies: An array of 0 to 1024 identities that have access to the key vault. All identities in the array must use the same tenant ID as the key vault's tenant ID. When `createMode` is set to `recover`, access policies are not required. Otherwise, access policies are required.
        :param pulumi.Input[str] create_mode: The vault's create mode to indicate whether the vault need to be recovered or not.
        :param pulumi.Input[bool] enable_purge_protection: Property specifying whether protection against purge is enabled for this vault. Setting this property to true activates protection against purge for this vault and its content - only the Key Vault service may initiate a hard, irrecoverable deletion. The setting is effective only if soft delete is also enabled. Enabling this functionality is irreversible - that is, the property does not accept false as its value.
        :param pulumi.Input[bool] enable_rbac_authorization: Property that controls how data actions are authorized. When true, the key vault will use Role Based Access Control (RBAC) for authorization of data actions, and the access policies specified in vault properties will be  ignored (warning: this is a preview feature). When false, the key vault will use the access policies specified in vault properties, and any policy stored on Azure Resource Manager will be ignored. If null or not specified, the vault is created with the default value of false. Note that management actions are always authorized with RBAC.
        :param pulumi.Input[bool] enable_soft_delete: Property to specify whether the 'soft delete' functionality is enabled for this key vault. If it's not set to any value(true or false) when creating new key vault, it will be set to true by default. Once set to true, it cannot be reverted to false.
        :param pulumi.Input[bool] enabled_for_deployment: Property to specify whether Azure Virtual Machines are permitted to retrieve certificates stored as secrets from the key vault.
        :param pulumi.Input[bool] enabled_for_disk_encryption: Property to specify whether Azure Disk Encryption is permitted to retrieve secrets from the vault and unwrap keys.
        :param pulumi.Input[bool] enabled_for_template_deployment: Property to specify whether Azure Resource Manager is permitted to retrieve secrets from the key vault.
        :param pulumi.Input['NetworkRuleSetArgs'] network_acls: Rules governing the accessibility of the key vault from specific network locations.
        :param pulumi.Input[float] soft_delete_retention_in_days: softDelete data retention days. It accepts >=7 and <=90.
        :param pulumi.Input[str] vault_uri: The URI of the vault for performing operations on keys and secrets.
        """
        pulumi.set(__self__, "sku", sku)
        pulumi.set(__self__, "tenant_id", tenant_id)
        if access_policies is not None:
            pulumi.set(__self__, "access_policies", access_policies)
        if create_mode is not None:
            pulumi.set(__self__, "create_mode", create_mode)
        if enable_purge_protection is not None:
            pulumi.set(__self__, "enable_purge_protection", enable_purge_protection)
        if enable_rbac_authorization is not None:
            pulumi.set(__self__, "enable_rbac_authorization", enable_rbac_authorization)
        if enable_soft_delete is not None:
            pulumi.set(__self__, "enable_soft_delete", enable_soft_delete)
        if enabled_for_deployment is not None:
            pulumi.set(__self__, "enabled_for_deployment", enabled_for_deployment)
        if enabled_for_disk_encryption is not None:
            pulumi.set(__self__, "enabled_for_disk_encryption", enabled_for_disk_encryption)
        if enabled_for_template_deployment is not None:
            pulumi.set(__self__, "enabled_for_template_deployment", enabled_for_template_deployment)
        if network_acls is not None:
            pulumi.set(__self__, "network_acls", network_acls)
        if soft_delete_retention_in_days is not None:
            pulumi.set(__self__, "soft_delete_retention_in_days", soft_delete_retention_in_days)
        if vault_uri is not None:
            pulumi.set(__self__, "vault_uri", vault_uri)

    @property
    @pulumi.getter
    def sku(self) -> pulumi.Input['SkuArgs']:
        """
        SKU details
        """
        return pulumi.get(self, "sku")

    @sku.setter
    def sku(self, value: pulumi.Input['SkuArgs']):
        pulumi.set(self, "sku", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Input[str]:
        """
        The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter(name="accessPolicies")
    def access_policies(self) -> Optional[pulumi.Input[List[pulumi.Input['AccessPolicyEntryArgs']]]]:
        """
        An array of 0 to 1024 identities that have access to the key vault. All identities in the array must use the same tenant ID as the key vault's tenant ID. When `createMode` is set to `recover`, access policies are not required. Otherwise, access policies are required.
        """
        return pulumi.get(self, "access_policies")

    @access_policies.setter
    def access_policies(self, value: Optional[pulumi.Input[List[pulumi.Input['AccessPolicyEntryArgs']]]]):
        pulumi.set(self, "access_policies", value)

    @property
    @pulumi.getter(name="createMode")
    def create_mode(self) -> Optional[pulumi.Input[str]]:
        """
        The vault's create mode to indicate whether the vault need to be recovered or not.
        """
        return pulumi.get(self, "create_mode")

    @create_mode.setter
    def create_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_mode", value)

    @property
    @pulumi.getter(name="enablePurgeProtection")
    def enable_purge_protection(self) -> Optional[pulumi.Input[bool]]:
        """
        Property specifying whether protection against purge is enabled for this vault. Setting this property to true activates protection against purge for this vault and its content - only the Key Vault service may initiate a hard, irrecoverable deletion. The setting is effective only if soft delete is also enabled. Enabling this functionality is irreversible - that is, the property does not accept false as its value.
        """
        return pulumi.get(self, "enable_purge_protection")

    @enable_purge_protection.setter
    def enable_purge_protection(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_purge_protection", value)

    @property
    @pulumi.getter(name="enableRbacAuthorization")
    def enable_rbac_authorization(self) -> Optional[pulumi.Input[bool]]:
        """
        Property that controls how data actions are authorized. When true, the key vault will use Role Based Access Control (RBAC) for authorization of data actions, and the access policies specified in vault properties will be  ignored (warning: this is a preview feature). When false, the key vault will use the access policies specified in vault properties, and any policy stored on Azure Resource Manager will be ignored. If null or not specified, the vault is created with the default value of false. Note that management actions are always authorized with RBAC.
        """
        return pulumi.get(self, "enable_rbac_authorization")

    @enable_rbac_authorization.setter
    def enable_rbac_authorization(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_rbac_authorization", value)

    @property
    @pulumi.getter(name="enableSoftDelete")
    def enable_soft_delete(self) -> Optional[pulumi.Input[bool]]:
        """
        Property to specify whether the 'soft delete' functionality is enabled for this key vault. If it's not set to any value(true or false) when creating new key vault, it will be set to true by default. Once set to true, it cannot be reverted to false.
        """
        return pulumi.get(self, "enable_soft_delete")

    @enable_soft_delete.setter
    def enable_soft_delete(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_soft_delete", value)

    @property
    @pulumi.getter(name="enabledForDeployment")
    def enabled_for_deployment(self) -> Optional[pulumi.Input[bool]]:
        """
        Property to specify whether Azure Virtual Machines are permitted to retrieve certificates stored as secrets from the key vault.
        """
        return pulumi.get(self, "enabled_for_deployment")

    @enabled_for_deployment.setter
    def enabled_for_deployment(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled_for_deployment", value)

    @property
    @pulumi.getter(name="enabledForDiskEncryption")
    def enabled_for_disk_encryption(self) -> Optional[pulumi.Input[bool]]:
        """
        Property to specify whether Azure Disk Encryption is permitted to retrieve secrets from the vault and unwrap keys.
        """
        return pulumi.get(self, "enabled_for_disk_encryption")

    @enabled_for_disk_encryption.setter
    def enabled_for_disk_encryption(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled_for_disk_encryption", value)

    @property
    @pulumi.getter(name="enabledForTemplateDeployment")
    def enabled_for_template_deployment(self) -> Optional[pulumi.Input[bool]]:
        """
        Property to specify whether Azure Resource Manager is permitted to retrieve secrets from the key vault.
        """
        return pulumi.get(self, "enabled_for_template_deployment")

    @enabled_for_template_deployment.setter
    def enabled_for_template_deployment(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled_for_template_deployment", value)

    @property
    @pulumi.getter(name="networkAcls")
    def network_acls(self) -> Optional[pulumi.Input['NetworkRuleSetArgs']]:
        """
        Rules governing the accessibility of the key vault from specific network locations.
        """
        return pulumi.get(self, "network_acls")

    @network_acls.setter
    def network_acls(self, value: Optional[pulumi.Input['NetworkRuleSetArgs']]):
        pulumi.set(self, "network_acls", value)

    @property
    @pulumi.getter(name="softDeleteRetentionInDays")
    def soft_delete_retention_in_days(self) -> Optional[pulumi.Input[float]]:
        """
        softDelete data retention days. It accepts >=7 and <=90.
        """
        return pulumi.get(self, "soft_delete_retention_in_days")

    @soft_delete_retention_in_days.setter
    def soft_delete_retention_in_days(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "soft_delete_retention_in_days", value)

    @property
    @pulumi.getter(name="vaultUri")
    def vault_uri(self) -> Optional[pulumi.Input[str]]:
        """
        The URI of the vault for performing operations on keys and secrets.
        """
        return pulumi.get(self, "vault_uri")

    @vault_uri.setter
    def vault_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vault_uri", value)


@pulumi.input_type
class VirtualNetworkRuleArgs:
    def __init__(__self__, *,
                 id: pulumi.Input[str]):
        """
        A rule governing the accessibility of a vault from a specific virtual network.
        :param pulumi.Input[str] id: Full resource id of a vnet subnet, such as '/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/subnet1'.
        """
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> pulumi.Input[str]:
        """
        Full resource id of a vnet subnet, such as '/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/subnet1'.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: pulumi.Input[str]):
        pulumi.set(self, "id", value)


