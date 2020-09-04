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
    'GetHybridUseBenefitResult',
    'AwaitableGetHybridUseBenefitResult',
    'get_hybrid_use_benefit',
]

@pulumi.output_type
class GetHybridUseBenefitResult:
    """
    Response on GET of a hybrid use benefit
    """
    def __init__(__self__, created_date=None, etag=None, last_updated_date=None, name=None, provisioning_state=None, sku=None, type=None):
        if created_date and not isinstance(created_date, str):
            raise TypeError("Expected argument 'created_date' to be a str")
        pulumi.set(__self__, "created_date", created_date)
        if etag and not isinstance(etag, float):
            raise TypeError("Expected argument 'etag' to be a float")
        pulumi.set(__self__, "etag", etag)
        if last_updated_date and not isinstance(last_updated_date, str):
            raise TypeError("Expected argument 'last_updated_date' to be a str")
        pulumi.set(__self__, "last_updated_date", last_updated_date)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if sku and not isinstance(sku, dict):
            raise TypeError("Expected argument 'sku' to be a dict")
        pulumi.set(__self__, "sku", sku)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="createdDate")
    def created_date(self) -> str:
        """
        Created date
        """
        return pulumi.get(self, "created_date")

    @property
    @pulumi.getter
    def etag(self) -> float:
        """
        Indicates the revision of the hybrid use benefit
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="lastUpdatedDate")
    def last_updated_date(self) -> str:
        """
        Last updated date
        """
        return pulumi.get(self, "last_updated_date")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        Provisioning state
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def sku(self) -> 'outputs.SkuResponse':
        """
        Hybrid use benefit SKU
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        """
        return pulumi.get(self, "type")


class AwaitableGetHybridUseBenefitResult(GetHybridUseBenefitResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetHybridUseBenefitResult(
            created_date=self.created_date,
            etag=self.etag,
            last_updated_date=self.last_updated_date,
            name=self.name,
            provisioning_state=self.provisioning_state,
            sku=self.sku,
            type=self.type)


def get_hybrid_use_benefit(plan_id: Optional[str] = None,
                           scope: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetHybridUseBenefitResult:
    """
    Use this data source to access information about an existing resource.

    :param str plan_id: This is a unique identifier for a plan. Should be a guid.
    :param str scope: The scope at which the operation is performed. This is limited to Microsoft.Compute/virtualMachines and Microsoft.Compute/hostGroups/hosts for now
    """
    __args__ = dict()
    __args__['planId'] = plan_id
    __args__['scope'] = scope
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:softwareplan/v20190601preview:getHybridUseBenefit', __args__, opts=opts, typ=GetHybridUseBenefitResult).value

    return AwaitableGetHybridUseBenefitResult(
        created_date=__ret__.created_date,
        etag=__ret__.etag,
        last_updated_date=__ret__.last_updated_date,
        name=__ret__.name,
        provisioning_state=__ret__.provisioning_state,
        sku=__ret__.sku,
        type=__ret__.type)
