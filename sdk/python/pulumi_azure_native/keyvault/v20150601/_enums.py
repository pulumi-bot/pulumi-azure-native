# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'CertificatePermissions',
    'KeyPermissions',
    'SecretPermissions',
    'SkuFamily',
    'SkuName',
]


class CertificatePermissions(str, Enum):
    ALL = "all"
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


class KeyPermissions(str, Enum):
    ALL = "all"
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


class SecretPermissions(str, Enum):
    ALL = "all"
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
