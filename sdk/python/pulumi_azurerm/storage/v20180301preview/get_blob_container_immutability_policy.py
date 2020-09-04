# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'GetBlobContainerImmutabilityPolicyResult',
    'AwaitableGetBlobContainerImmutabilityPolicyResult',
    'get_blob_container_immutability_policy',
]

@pulumi.output_type
class GetBlobContainerImmutabilityPolicyResult:
    """
    The ImmutabilityPolicy property of a blob container, including Id, resource name, resource type, Etag.
    """
    def __init__(__self__, etag=None, immutability_period_since_creation_in_days=None, name=None, state=None, type=None):
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if immutability_period_since_creation_in_days and not isinstance(immutability_period_since_creation_in_days, float):
            raise TypeError("Expected argument 'immutability_period_since_creation_in_days' to be a float")
        pulumi.set(__self__, "immutability_period_since_creation_in_days", immutability_period_since_creation_in_days)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        Resource Etag.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="immutabilityPeriodSinceCreationInDays")
    def immutability_period_since_creation_in_days(self) -> float:
        """
        The immutability period for the blobs in the container since the policy creation, in days.
        """
        return pulumi.get(self, "immutability_period_since_creation_in_days")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The ImmutabilityPolicy state of a blob container, possible values include: Locked and Unlocked.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        """
        return pulumi.get(self, "type")


class AwaitableGetBlobContainerImmutabilityPolicyResult(GetBlobContainerImmutabilityPolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBlobContainerImmutabilityPolicyResult(
            etag=self.etag,
            immutability_period_since_creation_in_days=self.immutability_period_since_creation_in_days,
            name=self.name,
            state=self.state,
            type=self.type)


def get_blob_container_immutability_policy(account_name: Optional[str] = None,
                                           container_name: Optional[str] = None,
                                           immutability_policy_name: Optional[str] = None,
                                           resource_group_name: Optional[str] = None,
                                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBlobContainerImmutabilityPolicyResult:
    """
    Use this data source to access information about an existing resource.

    :param str account_name: The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
    :param str container_name: The name of the blob container within the specified storage account. Blob container names must be between 3 and 63 characters in length and use numbers, lower-case letters and dash (-) only. Every dash (-) character must be immediately preceded and followed by a letter or number.
    :param str immutability_policy_name: The name of the blob container immutabilityPolicy within the specified storage account. ImmutabilityPolicy Name must be 'default'
    :param str resource_group_name: The name of the resource group within the user's subscription. The name is case insensitive.
    """
    __args__ = dict()
    __args__['accountName'] = account_name
    __args__['containerName'] = container_name
    __args__['immutabilityPolicyName'] = immutability_policy_name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:storage/v20180301preview:getBlobContainerImmutabilityPolicy', __args__, opts=opts, typ=GetBlobContainerImmutabilityPolicyResult).value

    return AwaitableGetBlobContainerImmutabilityPolicyResult(
        etag=__ret__.etag,
        immutability_period_since_creation_in_days=__ret__.immutability_period_since_creation_in_days,
        name=__ret__.name,
        state=__ret__.state,
        type=__ret__.type)
