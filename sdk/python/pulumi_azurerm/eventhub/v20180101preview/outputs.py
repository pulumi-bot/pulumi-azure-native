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
    'ClusterSkuResponse',
    'ConnectionStateResponse',
    'EncryptionResponse',
    'IdentityResponse',
    'KeyVaultPropertiesResponse',
    'PrivateEndpointResponse',
    'SkuResponse',
]

@pulumi.output_type
class ClusterSkuResponse(dict):
    """
    SKU parameters particular to a cluster instance.
    """
    def __init__(__self__, *,
                 name: str,
                 capacity: Optional[float] = None):
        """
        SKU parameters particular to a cluster instance.
        :param str name: Name of this SKU.
        :param float capacity: The quantity of Event Hubs Cluster Capacity Units contained in this cluster.
        """
        pulumi.set(__self__, "name", name)
        if capacity is not None:
            pulumi.set(__self__, "capacity", capacity)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of this SKU.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def capacity(self) -> Optional[float]:
        """
        The quantity of Event Hubs Cluster Capacity Units contained in this cluster.
        """
        return pulumi.get(self, "capacity")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ConnectionStateResponse(dict):
    """
    ConnectionState information.
    """
    def __init__(__self__, *,
                 description: Optional[str] = None,
                 status: Optional[str] = None):
        """
        ConnectionState information.
        :param str description: Description of the connection state.
        :param str status: Status of the connection.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        Description of the connection state.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        Status of the connection.
        """
        return pulumi.get(self, "status")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class EncryptionResponse(dict):
    """
    Properties to configure Encryption
    """
    def __init__(__self__, *,
                 key_source: Optional[str] = None,
                 key_vault_properties: Optional[List['outputs.KeyVaultPropertiesResponse']] = None):
        """
        Properties to configure Encryption
        :param str key_source: Enumerates the possible value of keySource for Encryption
        :param List['KeyVaultPropertiesResponseArgs'] key_vault_properties: Properties of KeyVault
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
    def key_vault_properties(self) -> Optional[List['outputs.KeyVaultPropertiesResponse']]:
        """
        Properties of KeyVault
        """
        return pulumi.get(self, "key_vault_properties")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class IdentityResponse(dict):
    """
    Properties to configure Identity for Bring your Own Keys
    """
    def __init__(__self__, *,
                 principal_id: Optional[str] = None,
                 tenant_id: Optional[str] = None,
                 type: Optional[str] = None):
        """
        Properties to configure Identity for Bring your Own Keys
        :param str principal_id: ObjectId from the KeyVault
        :param str tenant_id: TenantId from the KeyVault
        :param str type: Enumerates the possible value Identity type, which currently supports only 'SystemAssigned'
        """
        if principal_id is not None:
            pulumi.set(__self__, "principal_id", principal_id)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> Optional[str]:
        """
        ObjectId from the KeyVault
        """
        return pulumi.get(self, "principal_id")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[str]:
        """
        TenantId from the KeyVault
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        Enumerates the possible value Identity type, which currently supports only 'SystemAssigned'
        """
        return pulumi.get(self, "type")

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
        :param str key_version: Key Version
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
        Key Version
        """
        return pulumi.get(self, "key_version")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PrivateEndpointResponse(dict):
    """
    PrivateEndpoint information.
    """
    def __init__(__self__, *,
                 id: Optional[str] = None):
        """
        PrivateEndpoint information.
        :param str id: The ARM identifier for Private Endpoint.
        """
        if id is not None:
            pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        The ARM identifier for Private Endpoint.
        """
        return pulumi.get(self, "id")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class SkuResponse(dict):
    """
    SKU parameters supplied to the create namespace operation
    """
    def __init__(__self__, *,
                 name: str,
                 capacity: Optional[float] = None,
                 tier: Optional[str] = None):
        """
        SKU parameters supplied to the create namespace operation
        :param str name: Name of this SKU.
        :param float capacity: The Event Hubs throughput units, value should be 0 to 20 throughput units.
        :param str tier: The billing tier of this particular SKU.
        """
        pulumi.set(__self__, "name", name)
        if capacity is not None:
            pulumi.set(__self__, "capacity", capacity)
        if tier is not None:
            pulumi.set(__self__, "tier", tier)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of this SKU.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def capacity(self) -> Optional[float]:
        """
        The Event Hubs throughput units, value should be 0 to 20 throughput units.
        """
        return pulumi.get(self, "capacity")

    @property
    @pulumi.getter
    def tier(self) -> Optional[str]:
        """
        The billing tier of this particular SKU.
        """
        return pulumi.get(self, "tier")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


