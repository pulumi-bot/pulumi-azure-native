# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'CertificatePermissions',
    'CreateMode',
    'KeyPermissions',
    'NetworkRuleAction',
    'NetworkRuleBypassOptions',
    'PrivateEndpointServiceConnectionStatus',
    'SecretPermissions',
    'SkuFamily',
    'SkuName',
    'StoragePermissions',
]


class CertificatePermissions(str, Enum):
    GET = "get"
    LIST = "list"
    DELETE = "delete"
    CREATE = "create"
    IMPORT_ = "import"
    UPDATE = "update"
    MANAGECONTACTS = "managecontacts"
    GETISSUERS = "getissuers"
    LISTISSUERS = "listissuers"
    SETISSUERS = "setissuers"
    DELETEISSUERS = "deleteissuers"
    MANAGEISSUERS = "manageissuers"
    RECOVER = "recover"
    PURGE = "purge"
    BACKUP = "backup"
    RESTORE = "restore"


class CreateMode(str, Enum):
    """
    The vault's create mode to indicate whether the vault need to be recovered or not.
    """
    RECOVER = "recover"
    DEFAULT = "default"


class KeyPermissions(str, Enum):
    ENCRYPT = "encrypt"
    DECRYPT = "decrypt"
    WRAP_KEY = "wrapKey"
    UNWRAP_KEY = "unwrapKey"
    SIGN = "sign"
    VERIFY = "verify"
    GET = "get"
    LIST = "list"
    CREATE = "create"
    UPDATE = "update"
    IMPORT_ = "import"
    DELETE = "delete"
    BACKUP = "backup"
    RESTORE = "restore"
    RECOVER = "recover"
    PURGE = "purge"


class NetworkRuleAction(str, Enum):
    """
    The default action when no rule from ipRules and from virtualNetworkRules match. This is only used after the bypass property has been evaluated.
    """
    ALLOW = "Allow"
    DENY = "Deny"


class NetworkRuleBypassOptions(str, Enum):
    """
    Tells what traffic can bypass network rules. This can be 'AzureServices' or 'None'.  If not specified the default is 'AzureServices'.
    """
    AZURE_SERVICES = "AzureServices"
    NONE = "None"


class PrivateEndpointServiceConnectionStatus(str, Enum):
    """
    Indicates whether the connection has been approved, rejected or removed by the key vault owner.
    """
    PENDING = "Pending"
    APPROVED = "Approved"
    REJECTED = "Rejected"
    DISCONNECTED = "Disconnected"


class SecretPermissions(str, Enum):
    GET = "get"
    LIST = "list"
    SET = "set"
    DELETE = "delete"
    BACKUP = "backup"
    RESTORE = "restore"
    RECOVER = "recover"
    PURGE = "purge"


class SkuFamily(str, Enum):
    """
    SKU family name
    """
    A = "A"


class SkuName(str, Enum):
    """
    SKU name to specify whether the key vault is a standard vault or a premium vault.
    """
    STANDARD = "standard"
    PREMIUM = "premium"


class StoragePermissions(str, Enum):
    GET = "get"
    LIST = "list"
    DELETE = "delete"
    SET = "set"
    UPDATE = "update"
    REGENERATEKEY = "regeneratekey"
    RECOVER = "recover"
    PURGE = "purge"
    BACKUP = "backup"
    RESTORE = "restore"
    SETSAS = "setsas"
    LISTSAS = "listsas"
    GETSAS = "getsas"
    DELETESAS = "deletesas"
