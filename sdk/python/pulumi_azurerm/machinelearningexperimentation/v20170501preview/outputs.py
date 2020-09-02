# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'StorageAccountPropertiesResponse',
]

@pulumi.output_type
class StorageAccountPropertiesResponse(dict):
    """
    The properties of a storage account for a machine learning team account.
    """
    def __init__(__self__, *,
                 access_key: str,
                 storage_account_id: str):
        """
        The properties of a storage account for a machine learning team account.
        :param str access_key: The access key to the storage account.
        :param str storage_account_id: The fully qualified arm Id of the storage account.
        """
        pulumi.set(__self__, "access_key", access_key)
        pulumi.set(__self__, "storage_account_id", storage_account_id)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> str:
        """
        The access key to the storage account.
        """
        return pulumi.get(self, "access_key")

    @property
    @pulumi.getter(name="storageAccountId")
    def storage_account_id(self) -> str:
        """
        The fully qualified arm Id of the storage account.
        """
        return pulumi.get(self, "storage_account_id")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


