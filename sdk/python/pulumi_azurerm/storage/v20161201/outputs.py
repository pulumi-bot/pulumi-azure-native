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
    'CustomDomainResponse',
    'EncryptionResponse',
    'EncryptionServiceResponse',
    'EncryptionServicesResponse',
    'EndpointsResponse',
    'SkuResponse',
    'StorageAccountKeyResponseResult',
]

@pulumi.output_type
class CustomDomainResponse(dict):
    """
    The custom domain assigned to this storage account. This can be set via Update.
    """
    def __init__(__self__, *,
                 name: str,
                 use_sub_domain_name: Optional[bool] = None):
        """
        The custom domain assigned to this storage account. This can be set via Update.
        :param str name: Gets or sets the custom domain name assigned to the storage account. Name is the CNAME source.
        :param bool use_sub_domain_name: Indicates whether indirect CName validation is enabled. Default value is false. This should only be set on updates.
        """
        pulumi.set(__self__, "name", name)
        if use_sub_domain_name is not None:
            pulumi.set(__self__, "use_sub_domain_name", use_sub_domain_name)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Gets or sets the custom domain name assigned to the storage account. Name is the CNAME source.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="useSubDomainName")
    def use_sub_domain_name(self) -> Optional[bool]:
        """
        Indicates whether indirect CName validation is enabled. Default value is false. This should only be set on updates.
        """
        return pulumi.get(self, "use_sub_domain_name")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class EncryptionResponse(dict):
    """
    The encryption settings on the storage account.
    """
    def __init__(__self__, *,
                 key_source: str,
                 services: Optional['outputs.EncryptionServicesResponse'] = None):
        """
        The encryption settings on the storage account.
        :param str key_source: The encryption keySource (provider). Possible values (case-insensitive):  Microsoft.Storage
        :param 'EncryptionServicesResponseArgs' services: List of services which support encryption.
        """
        pulumi.set(__self__, "key_source", key_source)
        if services is not None:
            pulumi.set(__self__, "services", services)

    @property
    @pulumi.getter(name="keySource")
    def key_source(self) -> str:
        """
        The encryption keySource (provider). Possible values (case-insensitive):  Microsoft.Storage
        """
        return pulumi.get(self, "key_source")

    @property
    @pulumi.getter
    def services(self) -> Optional['outputs.EncryptionServicesResponse']:
        """
        List of services which support encryption.
        """
        return pulumi.get(self, "services")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class EncryptionServiceResponse(dict):
    """
    A service that allows server-side encryption to be used.
    """
    def __init__(__self__, *,
                 last_enabled_time: str,
                 enabled: Optional[bool] = None):
        """
        A service that allows server-side encryption to be used.
        :param str last_enabled_time: Gets a rough estimate of the date/time when the encryption was last enabled by the user. Only returned when encryption is enabled. There might be some unencrypted blobs which were written after this time, as it is just a rough estimate.
        :param bool enabled: A boolean indicating whether or not the service encrypts the data as it is stored.
        """
        pulumi.set(__self__, "last_enabled_time", last_enabled_time)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)

    @property
    @pulumi.getter(name="lastEnabledTime")
    def last_enabled_time(self) -> str:
        """
        Gets a rough estimate of the date/time when the encryption was last enabled by the user. Only returned when encryption is enabled. There might be some unencrypted blobs which were written after this time, as it is just a rough estimate.
        """
        return pulumi.get(self, "last_enabled_time")

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        """
        A boolean indicating whether or not the service encrypts the data as it is stored.
        """
        return pulumi.get(self, "enabled")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class EncryptionServicesResponse(dict):
    """
    A list of services that support encryption.
    """
    def __init__(__self__, *,
                 queue: 'outputs.EncryptionServiceResponse',
                 table: 'outputs.EncryptionServiceResponse',
                 blob: Optional['outputs.EncryptionServiceResponse'] = None,
                 file: Optional['outputs.EncryptionServiceResponse'] = None):
        """
        A list of services that support encryption.
        :param 'EncryptionServiceResponseArgs' queue: The encryption function of the queue storage service.
        :param 'EncryptionServiceResponseArgs' table: The encryption function of the table storage service.
        :param 'EncryptionServiceResponseArgs' blob: The encryption function of the blob storage service.
        :param 'EncryptionServiceResponseArgs' file: The encryption function of the file storage service.
        """
        pulumi.set(__self__, "queue", queue)
        pulumi.set(__self__, "table", table)
        if blob is not None:
            pulumi.set(__self__, "blob", blob)
        if file is not None:
            pulumi.set(__self__, "file", file)

    @property
    @pulumi.getter
    def queue(self) -> 'outputs.EncryptionServiceResponse':
        """
        The encryption function of the queue storage service.
        """
        return pulumi.get(self, "queue")

    @property
    @pulumi.getter
    def table(self) -> 'outputs.EncryptionServiceResponse':
        """
        The encryption function of the table storage service.
        """
        return pulumi.get(self, "table")

    @property
    @pulumi.getter
    def blob(self) -> Optional['outputs.EncryptionServiceResponse']:
        """
        The encryption function of the blob storage service.
        """
        return pulumi.get(self, "blob")

    @property
    @pulumi.getter
    def file(self) -> Optional['outputs.EncryptionServiceResponse']:
        """
        The encryption function of the file storage service.
        """
        return pulumi.get(self, "file")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class EndpointsResponse(dict):
    """
    The URIs that are used to perform a retrieval of a public blob, queue, or table object.
    """
    def __init__(__self__, *,
                 blob: str,
                 file: str,
                 queue: str,
                 table: str):
        """
        The URIs that are used to perform a retrieval of a public blob, queue, or table object.
        :param str blob: Gets the blob endpoint.
        :param str file: Gets the file endpoint.
        :param str queue: Gets the queue endpoint.
        :param str table: Gets the table endpoint.
        """
        pulumi.set(__self__, "blob", blob)
        pulumi.set(__self__, "file", file)
        pulumi.set(__self__, "queue", queue)
        pulumi.set(__self__, "table", table)

    @property
    @pulumi.getter
    def blob(self) -> str:
        """
        Gets the blob endpoint.
        """
        return pulumi.get(self, "blob")

    @property
    @pulumi.getter
    def file(self) -> str:
        """
        Gets the file endpoint.
        """
        return pulumi.get(self, "file")

    @property
    @pulumi.getter
    def queue(self) -> str:
        """
        Gets the queue endpoint.
        """
        return pulumi.get(self, "queue")

    @property
    @pulumi.getter
    def table(self) -> str:
        """
        Gets the table endpoint.
        """
        return pulumi.get(self, "table")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class SkuResponse(dict):
    """
    The SKU of the storage account.
    """
    def __init__(__self__, *,
                 name: str,
                 tier: str):
        """
        The SKU of the storage account.
        :param str name: Gets or sets the sku name. Required for account creation; optional for update. Note that in older versions, sku name was called accountType.
        :param str tier: Gets the sku tier. This is based on the SKU name.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "tier", tier)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Gets or sets the sku name. Required for account creation; optional for update. Note that in older versions, sku name was called accountType.
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
class StorageAccountKeyResponseResult(dict):
    """
    An access key for the storage account.
    """
    def __init__(__self__, *,
                 key_name: str,
                 permissions: str,
                 value: str):
        """
        An access key for the storage account.
        :param str key_name: Name of the key.
        :param str permissions: Permissions for the key -- read-only or full permissions.
        :param str value: Base 64-encoded value of the key.
        """
        pulumi.set(__self__, "key_name", key_name)
        pulumi.set(__self__, "permissions", permissions)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter(name="keyName")
    def key_name(self) -> str:
        """
        Name of the key.
        """
        return pulumi.get(self, "key_name")

    @property
    @pulumi.getter
    def permissions(self) -> str:
        """
        Permissions for the key -- read-only or full permissions.
        """
        return pulumi.get(self, "permissions")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        Base 64-encoded value of the key.
        """
        return pulumi.get(self, "value")


