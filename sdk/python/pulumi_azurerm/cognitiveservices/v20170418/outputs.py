# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs

__all__ = [
    'CognitiveServicesAccountApiPropertiesResponse',
    'CognitiveServicesAccountPropertiesResponse',
    'EncryptionResponse',
    'IdentityResponse',
    'IpRuleResponse',
    'KeyVaultPropertiesResponse',
    'NetworkRuleSetResponse',
    'PrivateEndpointConnectionPropertiesResponse',
    'PrivateEndpointConnectionResponse',
    'PrivateEndpointResponse',
    'PrivateLinkServiceConnectionStateResponse',
    'SkuCapabilityResponse',
    'SkuResponse',
    'UserAssignedIdentityResponse',
    'UserOwnedStorageResponse',
    'VirtualNetworkRuleResponse',
]

@pulumi.output_type
class CognitiveServicesAccountApiPropertiesResponse(dict):
    """
    The api properties for special APIs.
    """
    def __init__(__self__, *,
                 event_hub_connection_string: Optional[str] = None,
                 qna_runtime_endpoint: Optional[str] = None,
                 statistics_enabled: Optional[bool] = None,
                 storage_account_connection_string: Optional[str] = None):
        """
        The api properties for special APIs.
        :param str event_hub_connection_string: (Personalization Only) The flag to enable statistics of Bing Search.
        :param str qna_runtime_endpoint: (QnAMaker Only) The runtime endpoint of QnAMaker.
        :param bool statistics_enabled: (Bing Search Only) The flag to enable statistics of Bing Search.
        :param str storage_account_connection_string: (Personalization Only) The storage account connection string.
        """
        if event_hub_connection_string is not None:
            pulumi.set(__self__, "event_hub_connection_string", event_hub_connection_string)
        if qna_runtime_endpoint is not None:
            pulumi.set(__self__, "qna_runtime_endpoint", qna_runtime_endpoint)
        if statistics_enabled is not None:
            pulumi.set(__self__, "statistics_enabled", statistics_enabled)
        if storage_account_connection_string is not None:
            pulumi.set(__self__, "storage_account_connection_string", storage_account_connection_string)

    @property
    @pulumi.getter(name="eventHubConnectionString")
    def event_hub_connection_string(self) -> Optional[str]:
        """
        (Personalization Only) The flag to enable statistics of Bing Search.
        """
        return pulumi.get(self, "event_hub_connection_string")

    @property
    @pulumi.getter(name="qnaRuntimeEndpoint")
    def qna_runtime_endpoint(self) -> Optional[str]:
        """
        (QnAMaker Only) The runtime endpoint of QnAMaker.
        """
        return pulumi.get(self, "qna_runtime_endpoint")

    @property
    @pulumi.getter(name="statisticsEnabled")
    def statistics_enabled(self) -> Optional[bool]:
        """
        (Bing Search Only) The flag to enable statistics of Bing Search.
        """
        return pulumi.get(self, "statistics_enabled")

    @property
    @pulumi.getter(name="storageAccountConnectionString")
    def storage_account_connection_string(self) -> Optional[str]:
        """
        (Personalization Only) The storage account connection string.
        """
        return pulumi.get(self, "storage_account_connection_string")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class CognitiveServicesAccountPropertiesResponse(dict):
    """
    Properties of Cognitive Services account.
    """
    def __init__(__self__, *,
                 capabilities: List['outputs.SkuCapabilityResponse'],
                 endpoint: str,
                 internal_id: str,
                 provisioning_state: str,
                 api_properties: Optional['outputs.CognitiveServicesAccountApiPropertiesResponse'] = None,
                 custom_sub_domain_name: Optional[str] = None,
                 encryption: Optional['outputs.EncryptionResponse'] = None,
                 network_acls: Optional['outputs.NetworkRuleSetResponse'] = None,
                 private_endpoint_connections: Optional[List['outputs.PrivateEndpointConnectionResponse']] = None,
                 public_network_access: Optional[str] = None,
                 user_owned_storage: Optional[List['outputs.UserOwnedStorageResponse']] = None):
        """
        Properties of Cognitive Services account.
        :param List['SkuCapabilityResponseArgs'] capabilities: Gets the capabilities of the cognitive services account. Each item indicates the capability of a specific feature. The values are read-only and for reference only.
        :param str endpoint: Endpoint of the created account.
        :param str internal_id: The internal identifier.
        :param str provisioning_state: Gets the status of the cognitive services account at the time the operation was called.
        :param 'CognitiveServicesAccountApiPropertiesResponseArgs' api_properties: The api properties for special APIs.
        :param str custom_sub_domain_name: Optional subdomain name used for token-based authentication.
        :param 'EncryptionResponseArgs' encryption: The encryption properties for this resource.
        :param 'NetworkRuleSetResponseArgs' network_acls: A collection of rules governing the accessibility from specific network locations.
        :param List['PrivateEndpointConnectionResponseArgs'] private_endpoint_connections: The private endpoint connection associated with the Cognitive Services account.
        :param str public_network_access: Whether or not public endpoint access is allowed for this account. Value is optional but if passed in, must be 'Enabled' or 'Disabled'
        :param List['UserOwnedStorageResponseArgs'] user_owned_storage: The storage accounts for this resource.
        """
        pulumi.set(__self__, "capabilities", capabilities)
        pulumi.set(__self__, "endpoint", endpoint)
        pulumi.set(__self__, "internal_id", internal_id)
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if api_properties is not None:
            pulumi.set(__self__, "api_properties", api_properties)
        if custom_sub_domain_name is not None:
            pulumi.set(__self__, "custom_sub_domain_name", custom_sub_domain_name)
        if encryption is not None:
            pulumi.set(__self__, "encryption", encryption)
        if network_acls is not None:
            pulumi.set(__self__, "network_acls", network_acls)
        if private_endpoint_connections is not None:
            pulumi.set(__self__, "private_endpoint_connections", private_endpoint_connections)
        if public_network_access is not None:
            pulumi.set(__self__, "public_network_access", public_network_access)
        if user_owned_storage is not None:
            pulumi.set(__self__, "user_owned_storage", user_owned_storage)

    @property
    @pulumi.getter
    def capabilities(self) -> List['outputs.SkuCapabilityResponse']:
        """
        Gets the capabilities of the cognitive services account. Each item indicates the capability of a specific feature. The values are read-only and for reference only.
        """
        return pulumi.get(self, "capabilities")

    @property
    @pulumi.getter
    def endpoint(self) -> str:
        """
        Endpoint of the created account.
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter(name="internalId")
    def internal_id(self) -> str:
        """
        The internal identifier.
        """
        return pulumi.get(self, "internal_id")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        Gets the status of the cognitive services account at the time the operation was called.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="apiProperties")
    def api_properties(self) -> Optional['outputs.CognitiveServicesAccountApiPropertiesResponse']:
        """
        The api properties for special APIs.
        """
        return pulumi.get(self, "api_properties")

    @property
    @pulumi.getter(name="customSubDomainName")
    def custom_sub_domain_name(self) -> Optional[str]:
        """
        Optional subdomain name used for token-based authentication.
        """
        return pulumi.get(self, "custom_sub_domain_name")

    @property
    @pulumi.getter
    def encryption(self) -> Optional['outputs.EncryptionResponse']:
        """
        The encryption properties for this resource.
        """
        return pulumi.get(self, "encryption")

    @property
    @pulumi.getter(name="networkAcls")
    def network_acls(self) -> Optional['outputs.NetworkRuleSetResponse']:
        """
        A collection of rules governing the accessibility from specific network locations.
        """
        return pulumi.get(self, "network_acls")

    @property
    @pulumi.getter(name="privateEndpointConnections")
    def private_endpoint_connections(self) -> Optional[List['outputs.PrivateEndpointConnectionResponse']]:
        """
        The private endpoint connection associated with the Cognitive Services account.
        """
        return pulumi.get(self, "private_endpoint_connections")

    @property
    @pulumi.getter(name="publicNetworkAccess")
    def public_network_access(self) -> Optional[str]:
        """
        Whether or not public endpoint access is allowed for this account. Value is optional but if passed in, must be 'Enabled' or 'Disabled'
        """
        return pulumi.get(self, "public_network_access")

    @property
    @pulumi.getter(name="userOwnedStorage")
    def user_owned_storage(self) -> Optional[List['outputs.UserOwnedStorageResponse']]:
        """
        The storage accounts for this resource.
        """
        return pulumi.get(self, "user_owned_storage")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class EncryptionResponse(dict):
    """
    Properties to configure Encryption
    """
    def __init__(__self__, *,
                 key_source: Optional[str] = None,
                 key_vault_properties: Optional['outputs.KeyVaultPropertiesResponse'] = None):
        """
        Properties to configure Encryption
        :param str key_source: Enumerates the possible value of keySource for Encryption
        :param 'KeyVaultPropertiesResponseArgs' key_vault_properties: Properties of KeyVault
        """
        if key_source is not None:
            pulumi.set(__self__, "key_source", key_source)
        if key_vault_properties is not None:
            pulumi.set(__self__, "key_vault_properties", key_vault_properties)

    @property
    @pulumi.getter(name="keySource")
    def key_source(self) -> Optional[str]:
        """
        Enumerates the possible value of keySource for Encryption
        """
        return pulumi.get(self, "key_source")

    @property
    @pulumi.getter(name="keyVaultProperties")
    def key_vault_properties(self) -> Optional['outputs.KeyVaultPropertiesResponse']:
        """
        Properties of KeyVault
        """
        return pulumi.get(self, "key_vault_properties")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class IdentityResponse(dict):
    """
    Managed service identity.
    """
    def __init__(__self__, *,
                 principal_id: str,
                 tenant_id: str,
                 type: Optional[str] = None,
                 user_assigned_identities: Optional[Mapping[str, 'outputs.UserAssignedIdentityResponse']] = None):
        """
        Managed service identity.
        :param str principal_id: Principal Id of managed service identity.
        :param str tenant_id: Tenant of managed service identity.
        :param str type: Type of managed service identity.
        :param Mapping[str, 'UserAssignedIdentityResponseArgs'] user_assigned_identities: The list of user assigned identities associated with the resource. The user identity dictionary key references will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}
        """
        pulumi.set(__self__, "principal_id", principal_id)
        pulumi.set(__self__, "tenant_id", tenant_id)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if user_assigned_identities is not None:
            pulumi.set(__self__, "user_assigned_identities", user_assigned_identities)

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> str:
        """
        Principal Id of managed service identity.
        """
        return pulumi.get(self, "principal_id")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> str:
        """
        Tenant of managed service identity.
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        Type of managed service identity.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="userAssignedIdentities")
    def user_assigned_identities(self) -> Optional[Mapping[str, 'outputs.UserAssignedIdentityResponse']]:
        """
        The list of user assigned identities associated with the resource. The user identity dictionary key references will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}
        """
        return pulumi.get(self, "user_assigned_identities")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class IpRuleResponse(dict):
    """
    A rule governing the accessibility from a specific ip address or ip range.
    """
    def __init__(__self__, *,
                 value: str):
        """
        A rule governing the accessibility from a specific ip address or ip range.
        :param str value: An IPv4 address range in CIDR notation, such as '124.56.78.91' (simple IP address) or '124.56.78.0/24' (all addresses that start with 124.56.78).
        """
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        An IPv4 address range in CIDR notation, such as '124.56.78.91' (simple IP address) or '124.56.78.0/24' (all addresses that start with 124.56.78).
        """
        return pulumi.get(self, "value")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class KeyVaultPropertiesResponse(dict):
    """
    Properties to configure keyVault Properties
    """
    def __init__(__self__, *,
                 key_name: Optional[str] = None,
                 key_vault_uri: Optional[str] = None,
                 key_version: Optional[str] = None):
        """
        Properties to configure keyVault Properties
        :param str key_name: Name of the Key from KeyVault
        :param str key_vault_uri: Uri of KeyVault
        :param str key_version: Version of the Key from KeyVault
        """
        if key_name is not None:
            pulumi.set(__self__, "key_name", key_name)
        if key_vault_uri is not None:
            pulumi.set(__self__, "key_vault_uri", key_vault_uri)
        if key_version is not None:
            pulumi.set(__self__, "key_version", key_version)

    @property
    @pulumi.getter(name="keyName")
    def key_name(self) -> Optional[str]:
        """
        Name of the Key from KeyVault
        """
        return pulumi.get(self, "key_name")

    @property
    @pulumi.getter(name="keyVaultUri")
    def key_vault_uri(self) -> Optional[str]:
        """
        Uri of KeyVault
        """
        return pulumi.get(self, "key_vault_uri")

    @property
    @pulumi.getter(name="keyVersion")
    def key_version(self) -> Optional[str]:
        """
        Version of the Key from KeyVault
        """
        return pulumi.get(self, "key_version")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class NetworkRuleSetResponse(dict):
    """
    A set of rules governing the network accessibility.
    """
    def __init__(__self__, *,
                 default_action: Optional[str] = None,
                 ip_rules: Optional[List['outputs.IpRuleResponse']] = None,
                 virtual_network_rules: Optional[List['outputs.VirtualNetworkRuleResponse']] = None):
        """
        A set of rules governing the network accessibility.
        :param str default_action: The default action when no rule from ipRules and from virtualNetworkRules match. This is only used after the bypass property has been evaluated.
        :param List['IpRuleResponseArgs'] ip_rules: The list of IP address rules.
        :param List['VirtualNetworkRuleResponseArgs'] virtual_network_rules: The list of virtual network rules.
        """
        if default_action is not None:
            pulumi.set(__self__, "default_action", default_action)
        if ip_rules is not None:
            pulumi.set(__self__, "ip_rules", ip_rules)
        if virtual_network_rules is not None:
            pulumi.set(__self__, "virtual_network_rules", virtual_network_rules)

    @property
    @pulumi.getter(name="defaultAction")
    def default_action(self) -> Optional[str]:
        """
        The default action when no rule from ipRules and from virtualNetworkRules match. This is only used after the bypass property has been evaluated.
        """
        return pulumi.get(self, "default_action")

    @property
    @pulumi.getter(name="ipRules")
    def ip_rules(self) -> Optional[List['outputs.IpRuleResponse']]:
        """
        The list of IP address rules.
        """
        return pulumi.get(self, "ip_rules")

    @property
    @pulumi.getter(name="virtualNetworkRules")
    def virtual_network_rules(self) -> Optional[List['outputs.VirtualNetworkRuleResponse']]:
        """
        The list of virtual network rules.
        """
        return pulumi.get(self, "virtual_network_rules")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PrivateEndpointConnectionPropertiesResponse(dict):
    """
    Properties of the PrivateEndpointConnectProperties.
    """
    def __init__(__self__, *,
                 private_link_service_connection_state: 'outputs.PrivateLinkServiceConnectionStateResponse',
                 group_ids: Optional[List[str]] = None,
                 private_endpoint: Optional['outputs.PrivateEndpointResponse'] = None):
        """
        Properties of the PrivateEndpointConnectProperties.
        :param 'PrivateLinkServiceConnectionStateResponseArgs' private_link_service_connection_state: A collection of information about the state of the connection between service consumer and provider.
        :param List[str] group_ids: The private link resource group ids.
        :param 'PrivateEndpointResponseArgs' private_endpoint: The resource of private end point.
        """
        pulumi.set(__self__, "private_link_service_connection_state", private_link_service_connection_state)
        if group_ids is not None:
            pulumi.set(__self__, "group_ids", group_ids)
        if private_endpoint is not None:
            pulumi.set(__self__, "private_endpoint", private_endpoint)

    @property
    @pulumi.getter(name="privateLinkServiceConnectionState")
    def private_link_service_connection_state(self) -> 'outputs.PrivateLinkServiceConnectionStateResponse':
        """
        A collection of information about the state of the connection between service consumer and provider.
        """
        return pulumi.get(self, "private_link_service_connection_state")

    @property
    @pulumi.getter(name="groupIds")
    def group_ids(self) -> Optional[List[str]]:
        """
        The private link resource group ids.
        """
        return pulumi.get(self, "group_ids")

    @property
    @pulumi.getter(name="privateEndpoint")
    def private_endpoint(self) -> Optional['outputs.PrivateEndpointResponse']:
        """
        The resource of private end point.
        """
        return pulumi.get(self, "private_endpoint")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PrivateEndpointConnectionResponse(dict):
    """
    The Private Endpoint Connection resource.
    """
    def __init__(__self__, *,
                 id: str,
                 name: str,
                 type: str,
                 properties: Optional['outputs.PrivateEndpointConnectionPropertiesResponse'] = None):
        """
        The Private Endpoint Connection resource.
        :param str id: Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        :param str name: The name of the resource
        :param str type: The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        :param 'PrivateEndpointConnectionPropertiesResponseArgs' properties: Resource properties.
        """
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "type", type)
        if properties is not None:
            pulumi.set(__self__, "properties", properties)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def properties(self) -> Optional['outputs.PrivateEndpointConnectionPropertiesResponse']:
        """
        Resource properties.
        """
        return pulumi.get(self, "properties")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PrivateEndpointResponse(dict):
    """
    The Private Endpoint resource.
    """
    def __init__(__self__, *,
                 id: str):
        """
        The Private Endpoint resource.
        :param str id: The ARM identifier for Private Endpoint
        """
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ARM identifier for Private Endpoint
        """
        return pulumi.get(self, "id")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PrivateLinkServiceConnectionStateResponse(dict):
    """
    A collection of information about the state of the connection between service consumer and provider.
    """
    def __init__(__self__, *,
                 action_required: Optional[str] = None,
                 description: Optional[str] = None,
                 status: Optional[str] = None):
        """
        A collection of information about the state of the connection between service consumer and provider.
        :param str action_required: A message indicating if changes on the service provider require any updates on the consumer.
        :param str description: The reason for approval/rejection of the connection.
        :param str status: Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
        """
        if action_required is not None:
            pulumi.set(__self__, "action_required", action_required)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="actionRequired")
    def action_required(self) -> Optional[str]:
        """
        A message indicating if changes on the service provider require any updates on the consumer.
        """
        return pulumi.get(self, "action_required")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The reason for approval/rejection of the connection.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
        """
        return pulumi.get(self, "status")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class SkuCapabilityResponse(dict):
    """
    SkuCapability indicates the capability of a certain feature.
    """
    def __init__(__self__, *,
                 name: Optional[str] = None,
                 value: Optional[str] = None):
        """
        SkuCapability indicates the capability of a certain feature.
        :param str name: The name of the SkuCapability.
        :param str value: The value of the SkuCapability.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The name of the SkuCapability.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def value(self) -> Optional[str]:
        """
        The value of the SkuCapability.
        """
        return pulumi.get(self, "value")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class SkuResponse(dict):
    """
    The SKU of the cognitive services account.
    """
    def __init__(__self__, *,
                 name: str,
                 tier: str):
        """
        The SKU of the cognitive services account.
        :param str name: Gets or sets the sku name. Required for account creation, optional for update.
        :param str tier: Gets the sku tier. This is based on the SKU name.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "tier", tier)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Gets or sets the sku name. Required for account creation, optional for update.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tier(self) -> str:
        """
        Gets the sku tier. This is based on the SKU name.
        """
        return pulumi.get(self, "tier")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class UserAssignedIdentityResponse(dict):
    """
    User-assigned managed identity.
    """
    def __init__(__self__, *,
                 client_id: Optional[str] = None,
                 principal_id: Optional[str] = None):
        """
        User-assigned managed identity.
        :param str client_id: Client App Id associated with this identity.
        :param str principal_id: Azure Active Directory principal ID associated with this Identity.
        """
        if client_id is not None:
            pulumi.set(__self__, "client_id", client_id)
        if principal_id is not None:
            pulumi.set(__self__, "principal_id", principal_id)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> Optional[str]:
        """
        Client App Id associated with this identity.
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> Optional[str]:
        """
        Azure Active Directory principal ID associated with this Identity.
        """
        return pulumi.get(self, "principal_id")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class UserOwnedStorageResponse(dict):
    """
    The user owned storage for Cognitive Services account.
    """
    def __init__(__self__, *,
                 resource_id: Optional[str] = None):
        """
        The user owned storage for Cognitive Services account.
        :param str resource_id: Full resource id of a Microsoft.Storage resource.
        """
        if resource_id is not None:
            pulumi.set(__self__, "resource_id", resource_id)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> Optional[str]:
        """
        Full resource id of a Microsoft.Storage resource.
        """
        return pulumi.get(self, "resource_id")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class VirtualNetworkRuleResponse(dict):
    """
    A rule governing the accessibility from a specific virtual network.
    """
    def __init__(__self__, *,
                 id: str,
                 ignore_missing_vnet_service_endpoint: Optional[bool] = None,
                 state: Optional[str] = None):
        """
        A rule governing the accessibility from a specific virtual network.
        :param str id: Full resource id of a vnet subnet, such as '/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/subnet1'.
        :param bool ignore_missing_vnet_service_endpoint: Ignore missing vnet service endpoint or not.
        :param str state: Gets the state of virtual network rule.
        """
        pulumi.set(__self__, "id", id)
        if ignore_missing_vnet_service_endpoint is not None:
            pulumi.set(__self__, "ignore_missing_vnet_service_endpoint", ignore_missing_vnet_service_endpoint)
        if state is not None:
            pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Full resource id of a vnet subnet, such as '/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/subnet1'.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ignoreMissingVnetServiceEndpoint")
    def ignore_missing_vnet_service_endpoint(self) -> Optional[bool]:
        """
        Ignore missing vnet service endpoint or not.
        """
        return pulumi.get(self, "ignore_missing_vnet_service_endpoint")

    @property
    @pulumi.getter
    def state(self) -> Optional[str]:
        """
        Gets the state of virtual network rule.
        """
        return pulumi.get(self, "state")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


