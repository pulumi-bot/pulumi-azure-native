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
    'EncryptionResponse',
    'KeyVaultPropertiesResponse',
    'SBSkuResponse',
]

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
class KeyVaultPropertiesResponse(dict):
    """
    Properties to configure keyVault Properties
    """
    def __init__(__self__, *,
                 key_name: Optional[str] = None,
                 key_vault_uri: Optional[str] = None):
        """
        Properties to configure keyVault Properties
        :param str key_name: Name of the Key from KeyVault
        :param str key_vault_uri: Uri of KeyVault
        """
        if key_name is not None:
            pulumi.set(__self__, "key_name", key_name)
        if key_vault_uri is not None:
            pulumi.set(__self__, "key_vault_uri", key_vault_uri)

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

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class SBSkuResponse(dict):
    """
    SKU of the namespace.
    """
    def __init__(__self__, *,
                 name: str,
                 capacity: Optional[float] = None,
                 tier: Optional[str] = None):
        """
        SKU of the namespace.
        :param str name: Name of this SKU.
        :param float capacity: The specified messaging units for the tier. For Premium tier, capacity are 1,2 and 4.
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
        The specified messaging units for the tier. For Premium tier, capacity are 1,2 and 4.
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


