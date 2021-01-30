# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'AuthenticationMethod',
    'ProvisioningState',
    'ResourceIdentityType',
]


class AuthenticationMethod(str, Enum):
    """
    The mode of client authentication.
    """
    TOKEN = "Token"
    AAD = "AAD"


class ProvisioningState(str, Enum):
    """
    Provisioning state of the connected cluster resource.
    """
    SUCCEEDED = "Succeeded"
    FAILED = "Failed"
    CANCELED = "Canceled"
    PROVISIONING = "Provisioning"
    UPDATING = "Updating"
    DELETING = "Deleting"
    ACCEPTED = "Accepted"


class ResourceIdentityType(str, Enum):
    """
    The type of identity used for the connected cluster. The type 'SystemAssigned, includes a system created identity. The type 'None' means no identity is assigned to the connected cluster.
    """
    NONE = "None"
    SYSTEM_ASSIGNED = "SystemAssigned"
