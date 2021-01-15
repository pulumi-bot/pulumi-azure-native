# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'DiskCreateOption',
    'DiskEncryptionSetIdentityType',
    'DiskEncryptionSetType',
    'DiskStorageAccountTypes',
    'EncryptionType',
    'ExtendedLocationTypes',
    'GallerySharingPermissionTypes',
    'HostCaching',
    'HyperVGeneration',
    'NetworkAccessPolicy',
    'OperatingSystemStateTypes',
    'OperatingSystemTypes',
    'PrivateEndpointServiceConnectionStatus',
    'SnapshotStorageAccountTypes',
    'StorageAccountType',
]


class DiskCreateOption(str, Enum):
    """
    This enumerates the possible sources of a disk's creation.
    """
    EMPTY = "Empty"
    ATTACH = "Attach"
    FROM_IMAGE = "FromImage"
    IMPORT_ = "Import"
    COPY = "Copy"
    RESTORE = "Restore"
    UPLOAD = "Upload"


class DiskEncryptionSetIdentityType(str, Enum):
    """
    The type of Managed Identity used by the DiskEncryptionSet. Only SystemAssigned is supported for new creations. Disk Encryption Sets can be updated with Identity type None during migration of subscription to a new Azure Active Directory tenant; it will cause the encrypted resources to lose access to the keys.
    """
    SYSTEM_ASSIGNED = "SystemAssigned"
    NONE = "None"


class DiskEncryptionSetType(str, Enum):
    """
    The type of key used to encrypt the data of the disk.
    """
    ENCRYPTION_AT_REST_WITH_CUSTOMER_KEY = "EncryptionAtRestWithCustomerKey"
    ENCRYPTION_AT_REST_WITH_PLATFORM_AND_CUSTOMER_KEYS = "EncryptionAtRestWithPlatformAndCustomerKeys"


class DiskStorageAccountTypes(str, Enum):
    """
    The sku name.
    """
    STANDARD_LRS = "Standard_LRS"
    PREMIUM_LRS = "Premium_LRS"
    STANDARD_SS_D_LRS = "StandardSSD_LRS"
    ULTRA_SS_D_LRS = "UltraSSD_LRS"


class EncryptionType(str, Enum):
    """
    The type of key used to encrypt the data of the disk.
    """
    ENCRYPTION_AT_REST_WITH_PLATFORM_KEY = "EncryptionAtRestWithPlatformKey"
    ENCRYPTION_AT_REST_WITH_CUSTOMER_KEY = "EncryptionAtRestWithCustomerKey"
    ENCRYPTION_AT_REST_WITH_PLATFORM_AND_CUSTOMER_KEYS = "EncryptionAtRestWithPlatformAndCustomerKeys"


class ExtendedLocationTypes(str, Enum):
    """
    The type of the extended location.
    """
    EDGE_ZONE = "EdgeZone"


class GallerySharingPermissionTypes(str, Enum):
    """
    This property allows you to specify the permission of sharing gallery. <br><br> Possible values are: <br><br> **Private** <br><br> **Groups**
    """
    PRIVATE = "Private"
    GROUPS = "Groups"


class HostCaching(str, Enum):
    """
    The host caching of the disk. Valid values are 'None', 'ReadOnly', and 'ReadWrite'
    """
    NONE = "None"
    READ_ONLY = "ReadOnly"
    READ_WRITE = "ReadWrite"


class HyperVGeneration(str, Enum):
    """
    The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
    """
    V1 = "V1"
    V2 = "V2"


class NetworkAccessPolicy(str, Enum):
    """
    Policy for accessing the disk via network.
    """
    ALLOW_ALL = "AllowAll"
    ALLOW_PRIVATE = "AllowPrivate"
    DENY_ALL = "DenyAll"


class OperatingSystemStateTypes(str, Enum):
    """
    This property allows the user to specify whether the virtual machines created under this image are 'Generalized' or 'Specialized'.
    """
    GENERALIZED = "Generalized"
    SPECIALIZED = "Specialized"


class OperatingSystemTypes(str, Enum):
    """
    The Operating System type.
    """
    WINDOWS = "Windows"
    LINUX = "Linux"


class PrivateEndpointServiceConnectionStatus(str, Enum):
    """
    Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
    """
    PENDING = "Pending"
    APPROVED = "Approved"
    REJECTED = "Rejected"


class SnapshotStorageAccountTypes(str, Enum):
    """
    The sku name.
    """
    STANDARD_LRS = "Standard_LRS"
    PREMIUM_LRS = "Premium_LRS"
    STANDARD_ZRS = "Standard_ZRS"


class StorageAccountType(str, Enum):
    """
    Specifies the storage account type to be used to store the image. This property is not updatable.
    """
    STANDARD_LRS = "Standard_LRS"
    STANDARD_ZRS = "Standard_ZRS"
    PREMIUM_LRS = "Premium_LRS"
