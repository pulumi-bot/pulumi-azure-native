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
    'CertificatePropertiesResponse',
    'EncryptionPropertiesDescriptionResponse',
    'IotDpsPropertiesDescriptionResponse',
    'IotDpsSkuInfoResponse',
    'IotHubDefinitionDescriptionResponse',
    'IpFilterRuleResponse',
    'KeyVaultKeyPropertiesResponse',
    'PrivateEndpointConnectionPropertiesResponse',
    'PrivateEndpointConnectionResponse',
    'PrivateEndpointResponse',
    'PrivateLinkServiceConnectionStateResponse',
    'SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponse',
]

@pulumi.output_type
class CertificatePropertiesResponse(dict):
    """
    The description of an X509 CA Certificate.
    """
    def __init__(__self__, *,
                 created: str,
                 expiry: str,
                 is_verified: bool,
                 subject: str,
                 thumbprint: str,
                 updated: str):
        """
        The description of an X509 CA Certificate.
        :param str created: The certificate's creation date and time.
        :param str expiry: The certificate's expiration date and time.
        :param bool is_verified: Determines whether certificate has been verified.
        :param str subject: The certificate's subject name.
        :param str thumbprint: The certificate's thumbprint.
        :param str updated: The certificate's last update date and time.
        """
        pulumi.set(__self__, "created", created)
        pulumi.set(__self__, "expiry", expiry)
        pulumi.set(__self__, "is_verified", is_verified)
        pulumi.set(__self__, "subject", subject)
        pulumi.set(__self__, "thumbprint", thumbprint)
        pulumi.set(__self__, "updated", updated)

    @property
    @pulumi.getter
    def created(self) -> str:
        """
        The certificate's creation date and time.
        """
        return pulumi.get(self, "created")

    @property
    @pulumi.getter
    def expiry(self) -> str:
        """
        The certificate's expiration date and time.
        """
        return pulumi.get(self, "expiry")

    @property
    @pulumi.getter(name="isVerified")
    def is_verified(self) -> bool:
        """
        Determines whether certificate has been verified.
        """
        return pulumi.get(self, "is_verified")

    @property
    @pulumi.getter
    def subject(self) -> str:
        """
        The certificate's subject name.
        """
        return pulumi.get(self, "subject")

    @property
    @pulumi.getter
    def thumbprint(self) -> str:
        """
        The certificate's thumbprint.
        """
        return pulumi.get(self, "thumbprint")

    @property
    @pulumi.getter
    def updated(self) -> str:
        """
        The certificate's last update date and time.
        """
        return pulumi.get(self, "updated")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class EncryptionPropertiesDescriptionResponse(dict):
    """
    The encryption properties for the IoT DPS instance.
    """
    def __init__(__self__, *,
                 key_source: Optional[str] = None,
                 key_vault_properties: Optional[List['outputs.KeyVaultKeyPropertiesResponse']] = None):
        """
        The encryption properties for the IoT DPS instance.
        :param str key_source: The source of the key.
        :param List['KeyVaultKeyPropertiesResponseArgs'] key_vault_properties: The properties of the KeyVault key.
        """
        if key_source is not None:
            pulumi.set(__self__, "key_source", key_source)
        if key_vault_properties is not None:
            pulumi.set(__self__, "key_vault_properties", key_vault_properties)

    @property
    @pulumi.getter(name="keySource")
    def key_source(self) -> Optional[str]:
        """
        The source of the key.
        """
        return pulumi.get(self, "key_source")

    @property
    @pulumi.getter(name="keyVaultProperties")
    def key_vault_properties(self) -> Optional[List['outputs.KeyVaultKeyPropertiesResponse']]:
        """
        The properties of the KeyVault key.
        """
        return pulumi.get(self, "key_vault_properties")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class IotDpsPropertiesDescriptionResponse(dict):
    """
    the service specific properties of a provisioning service, including keys, linked iot hubs, current state, and system generated properties such as hostname and idScope
    """
    def __init__(__self__, *,
                 device_provisioning_host_name: str,
                 id_scope: str,
                 service_operations_host_name: str,
                 allocation_policy: Optional[str] = None,
                 authorization_policies: Optional[List['outputs.SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponse']] = None,
                 encryption: Optional['outputs.EncryptionPropertiesDescriptionResponse'] = None,
                 iot_hubs: Optional[List['outputs.IotHubDefinitionDescriptionResponse']] = None,
                 ip_filter_rules: Optional[List['outputs.IpFilterRuleResponse']] = None,
                 private_endpoint_connections: Optional[List['outputs.PrivateEndpointConnectionResponse']] = None,
                 provisioning_state: Optional[str] = None,
                 public_network_access: Optional[str] = None,
                 state: Optional[str] = None):
        """
        the service specific properties of a provisioning service, including keys, linked iot hubs, current state, and system generated properties such as hostname and idScope
        :param str device_provisioning_host_name: Device endpoint for this provisioning service.
        :param str id_scope: Unique identifier of this provisioning service.
        :param str service_operations_host_name: Service endpoint for provisioning service.
        :param str allocation_policy: Allocation policy to be used by this provisioning service.
        :param List['SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArgs'] authorization_policies: List of authorization keys for a provisioning service.
        :param 'EncryptionPropertiesDescriptionResponseArgs' encryption: The encryption properties for the IoT DPS instance.
        :param List['IotHubDefinitionDescriptionResponseArgs'] iot_hubs: List of IoT hubs associated with this provisioning service.
        :param List['IpFilterRuleResponseArgs'] ip_filter_rules: The IP filter rules.
        :param List['PrivateEndpointConnectionResponseArgs'] private_endpoint_connections: Private endpoint connections created on this IotHub
        :param str provisioning_state: The ARM provisioning state of the provisioning service.
        :param str public_network_access: Whether requests from Public Network are allowed
        :param str state: Current state of the provisioning service.
        """
        pulumi.set(__self__, "device_provisioning_host_name", device_provisioning_host_name)
        pulumi.set(__self__, "id_scope", id_scope)
        pulumi.set(__self__, "service_operations_host_name", service_operations_host_name)
        if allocation_policy is not None:
            pulumi.set(__self__, "allocation_policy", allocation_policy)
        if authorization_policies is not None:
            pulumi.set(__self__, "authorization_policies", authorization_policies)
        if encryption is not None:
            pulumi.set(__self__, "encryption", encryption)
        if iot_hubs is not None:
            pulumi.set(__self__, "iot_hubs", iot_hubs)
        if ip_filter_rules is not None:
            pulumi.set(__self__, "ip_filter_rules", ip_filter_rules)
        if private_endpoint_connections is not None:
            pulumi.set(__self__, "private_endpoint_connections", private_endpoint_connections)
        if provisioning_state is not None:
            pulumi.set(__self__, "provisioning_state", provisioning_state)
        if public_network_access is not None:
            pulumi.set(__self__, "public_network_access", public_network_access)
        if state is not None:
            pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter(name="deviceProvisioningHostName")
    def device_provisioning_host_name(self) -> str:
        """
        Device endpoint for this provisioning service.
        """
        return pulumi.get(self, "device_provisioning_host_name")

    @property
    @pulumi.getter(name="idScope")
    def id_scope(self) -> str:
        """
        Unique identifier of this provisioning service.
        """
        return pulumi.get(self, "id_scope")

    @property
    @pulumi.getter(name="serviceOperationsHostName")
    def service_operations_host_name(self) -> str:
        """
        Service endpoint for provisioning service.
        """
        return pulumi.get(self, "service_operations_host_name")

    @property
    @pulumi.getter(name="allocationPolicy")
    def allocation_policy(self) -> Optional[str]:
        """
        Allocation policy to be used by this provisioning service.
        """
        return pulumi.get(self, "allocation_policy")

    @property
    @pulumi.getter(name="authorizationPolicies")
    def authorization_policies(self) -> Optional[List['outputs.SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponse']]:
        """
        List of authorization keys for a provisioning service.
        """
        return pulumi.get(self, "authorization_policies")

    @property
    @pulumi.getter
    def encryption(self) -> Optional['outputs.EncryptionPropertiesDescriptionResponse']:
        """
        The encryption properties for the IoT DPS instance.
        """
        return pulumi.get(self, "encryption")

    @property
    @pulumi.getter(name="iotHubs")
    def iot_hubs(self) -> Optional[List['outputs.IotHubDefinitionDescriptionResponse']]:
        """
        List of IoT hubs associated with this provisioning service.
        """
        return pulumi.get(self, "iot_hubs")

    @property
    @pulumi.getter(name="ipFilterRules")
    def ip_filter_rules(self) -> Optional[List['outputs.IpFilterRuleResponse']]:
        """
        The IP filter rules.
        """
        return pulumi.get(self, "ip_filter_rules")

    @property
    @pulumi.getter(name="privateEndpointConnections")
    def private_endpoint_connections(self) -> Optional[List['outputs.PrivateEndpointConnectionResponse']]:
        """
        Private endpoint connections created on this IotHub
        """
        return pulumi.get(self, "private_endpoint_connections")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> Optional[str]:
        """
        The ARM provisioning state of the provisioning service.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="publicNetworkAccess")
    def public_network_access(self) -> Optional[str]:
        """
        Whether requests from Public Network are allowed
        """
        return pulumi.get(self, "public_network_access")

    @property
    @pulumi.getter
    def state(self) -> Optional[str]:
        """
        Current state of the provisioning service.
        """
        return pulumi.get(self, "state")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class IotDpsSkuInfoResponse(dict):
    """
    List of possible provisioning service SKUs.
    """
    def __init__(__self__, *,
                 tier: str,
                 capacity: Optional[float] = None,
                 name: Optional[str] = None):
        """
        List of possible provisioning service SKUs.
        :param str tier: Pricing tier name of the provisioning service.
        :param float capacity: The number of units to provision
        :param str name: Sku name.
        """
        pulumi.set(__self__, "tier", tier)
        if capacity is not None:
            pulumi.set(__self__, "capacity", capacity)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def tier(self) -> str:
        """
        Pricing tier name of the provisioning service.
        """
        return pulumi.get(self, "tier")

    @property
    @pulumi.getter
    def capacity(self) -> Optional[float]:
        """
        The number of units to provision
        """
        return pulumi.get(self, "capacity")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Sku name.
        """
        return pulumi.get(self, "name")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class IotHubDefinitionDescriptionResponse(dict):
    """
    Description of the IoT hub.
    """
    def __init__(__self__, *,
                 connection_string: str,
                 location: str,
                 name: str,
                 allocation_weight: Optional[float] = None,
                 apply_allocation_policy: Optional[bool] = None):
        """
        Description of the IoT hub.
        :param str connection_string: Connection string of the IoT hub.
        :param str location: ARM region of the IoT hub.
        :param str name: Host name of the IoT hub.
        :param float allocation_weight: weight to apply for a given iot h.
        :param bool apply_allocation_policy: flag for applying allocationPolicy or not for a given iot hub.
        """
        pulumi.set(__self__, "connection_string", connection_string)
        pulumi.set(__self__, "location", location)
        pulumi.set(__self__, "name", name)
        if allocation_weight is not None:
            pulumi.set(__self__, "allocation_weight", allocation_weight)
        if apply_allocation_policy is not None:
            pulumi.set(__self__, "apply_allocation_policy", apply_allocation_policy)

    @property
    @pulumi.getter(name="connectionString")
    def connection_string(self) -> str:
        """
        Connection string of the IoT hub.
        """
        return pulumi.get(self, "connection_string")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        ARM region of the IoT hub.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Host name of the IoT hub.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="allocationWeight")
    def allocation_weight(self) -> Optional[float]:
        """
        weight to apply for a given iot h.
        """
        return pulumi.get(self, "allocation_weight")

    @property
    @pulumi.getter(name="applyAllocationPolicy")
    def apply_allocation_policy(self) -> Optional[bool]:
        """
        flag for applying allocationPolicy or not for a given iot hub.
        """
        return pulumi.get(self, "apply_allocation_policy")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class IpFilterRuleResponse(dict):
    """
    The IP filter rules for a provisioning Service.
    """
    def __init__(__self__, *,
                 action: str,
                 filter_name: str,
                 ip_mask: str,
                 target: Optional[str] = None):
        """
        The IP filter rules for a provisioning Service.
        :param str action: The desired action for requests captured by this rule.
        :param str filter_name: The name of the IP filter rule.
        :param str ip_mask: A string that contains the IP address range in CIDR notation for the rule.
        :param str target: Target for requests captured by this rule.
        """
        pulumi.set(__self__, "action", action)
        pulumi.set(__self__, "filter_name", filter_name)
        pulumi.set(__self__, "ip_mask", ip_mask)
        if target is not None:
            pulumi.set(__self__, "target", target)

    @property
    @pulumi.getter
    def action(self) -> str:
        """
        The desired action for requests captured by this rule.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="filterName")
    def filter_name(self) -> str:
        """
        The name of the IP filter rule.
        """
        return pulumi.get(self, "filter_name")

    @property
    @pulumi.getter(name="ipMask")
    def ip_mask(self) -> str:
        """
        A string that contains the IP address range in CIDR notation for the rule.
        """
        return pulumi.get(self, "ip_mask")

    @property
    @pulumi.getter
    def target(self) -> Optional[str]:
        """
        Target for requests captured by this rule.
        """
        return pulumi.get(self, "target")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class KeyVaultKeyPropertiesResponse(dict):
    """
    The properties of the KeyVault key.
    """
    def __init__(__self__, *,
                 key_identifier: Optional[str] = None):
        """
        The properties of the KeyVault key.
        :param str key_identifier: The identifier of the key.
        """
        if key_identifier is not None:
            pulumi.set(__self__, "key_identifier", key_identifier)

    @property
    @pulumi.getter(name="keyIdentifier")
    def key_identifier(self) -> Optional[str]:
        """
        The identifier of the key.
        """
        return pulumi.get(self, "key_identifier")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PrivateEndpointConnectionPropertiesResponse(dict):
    """
    The properties of a private endpoint connection
    """
    def __init__(__self__, *,
                 private_link_service_connection_state: 'outputs.PrivateLinkServiceConnectionStateResponse',
                 private_endpoint: Optional['outputs.PrivateEndpointResponse'] = None):
        """
        The properties of a private endpoint connection
        :param 'PrivateLinkServiceConnectionStateResponseArgs' private_link_service_connection_state: The current state of a private endpoint connection
        :param 'PrivateEndpointResponseArgs' private_endpoint: The private endpoint property of a private endpoint connection
        """
        pulumi.set(__self__, "private_link_service_connection_state", private_link_service_connection_state)
        if private_endpoint is not None:
            pulumi.set(__self__, "private_endpoint", private_endpoint)

    @property
    @pulumi.getter(name="privateLinkServiceConnectionState")
    def private_link_service_connection_state(self) -> 'outputs.PrivateLinkServiceConnectionStateResponse':
        """
        The current state of a private endpoint connection
        """
        return pulumi.get(self, "private_link_service_connection_state")

    @property
    @pulumi.getter(name="privateEndpoint")
    def private_endpoint(self) -> Optional['outputs.PrivateEndpointResponse']:
        """
        The private endpoint property of a private endpoint connection
        """
        return pulumi.get(self, "private_endpoint")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PrivateEndpointConnectionResponse(dict):
    """
    The private endpoint connection of a provisioning service
    """
    def __init__(__self__, *,
                 id: str,
                 name: str,
                 properties: 'outputs.PrivateEndpointConnectionPropertiesResponse',
                 type: str):
        """
        The private endpoint connection of a provisioning service
        :param str id: The resource identifier.
        :param str name: The resource name.
        :param 'PrivateEndpointConnectionPropertiesResponseArgs' properties: The properties of a private endpoint connection
        :param str type: The resource type.
        """
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "properties", properties)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The resource identifier.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def properties(self) -> 'outputs.PrivateEndpointConnectionPropertiesResponse':
        """
        The properties of a private endpoint connection
        """
        return pulumi.get(self, "properties")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The resource type.
        """
        return pulumi.get(self, "type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PrivateEndpointResponse(dict):
    """
    The private endpoint property of a private endpoint connection
    """
    def __init__(__self__, *,
                 id: str):
        """
        The private endpoint property of a private endpoint connection
        :param str id: The resource identifier.
        """
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The resource identifier.
        """
        return pulumi.get(self, "id")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PrivateLinkServiceConnectionStateResponse(dict):
    """
    The current state of a private endpoint connection
    """
    def __init__(__self__, *,
                 description: str,
                 status: str,
                 actions_required: Optional[str] = None):
        """
        The current state of a private endpoint connection
        :param str description: The description for the current state of a private endpoint connection
        :param str status: The status of a private endpoint connection
        :param str actions_required: Actions required for a private endpoint connection
        """
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "status", status)
        if actions_required is not None:
            pulumi.set(__self__, "actions_required", actions_required)

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description for the current state of a private endpoint connection
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The status of a private endpoint connection
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="actionsRequired")
    def actions_required(self) -> Optional[str]:
        """
        Actions required for a private endpoint connection
        """
        return pulumi.get(self, "actions_required")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponse(dict):
    """
    Description of the shared access key.
    """
    def __init__(__self__, *,
                 key_name: str,
                 rights: str,
                 primary_key: Optional[str] = None,
                 secondary_key: Optional[str] = None):
        """
        Description of the shared access key.
        :param str key_name: Name of the key.
        :param str rights: Rights that this key has.
        :param str primary_key: Primary SAS key value.
        :param str secondary_key: Secondary SAS key value.
        """
        pulumi.set(__self__, "key_name", key_name)
        pulumi.set(__self__, "rights", rights)
        if primary_key is not None:
            pulumi.set(__self__, "primary_key", primary_key)
        if secondary_key is not None:
            pulumi.set(__self__, "secondary_key", secondary_key)

    @property
    @pulumi.getter(name="keyName")
    def key_name(self) -> str:
        """
        Name of the key.
        """
        return pulumi.get(self, "key_name")

    @property
    @pulumi.getter
    def rights(self) -> str:
        """
        Rights that this key has.
        """
        return pulumi.get(self, "rights")

    @property
    @pulumi.getter(name="primaryKey")
    def primary_key(self) -> Optional[str]:
        """
        Primary SAS key value.
        """
        return pulumi.get(self, "primary_key")

    @property
    @pulumi.getter(name="secondaryKey")
    def secondary_key(self) -> Optional[str]:
        """
        Secondary SAS key value.
        """
        return pulumi.get(self, "secondary_key")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


