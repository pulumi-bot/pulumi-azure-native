# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'DataStoreTypes',
    'StorageSettingStoreTypes',
    'StorageSettingTypes',
]


class DataStoreTypes(str, Enum):
    """
    type of datastore; Operational/Vault/Archive
    """
    OPERATIONAL_STORE = "OperationalStore"
    VAULT_STORE = "VaultStore"
    ARCHIVE_STORE = "ArchiveStore"


class StorageSettingStoreTypes(str, Enum):
    """
    Gets or sets the type of the datastore.
    """
    ARCHIVE_STORE = "ArchiveStore"
    SNAPSHOT_STORE = "SnapshotStore"
    VAULT_STORE = "VaultStore"


class StorageSettingTypes(str, Enum):
    """
    Gets or sets the type.
    """
    GEO_REDUNDANT = "GeoRedundant"
    LOCALLY_REDUNDANT = "LocallyRedundant"
