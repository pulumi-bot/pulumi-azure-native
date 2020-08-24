# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'GetPolicyResult',
    'AwaitableGetPolicyResult',
    'get_policy',
]

@pulumi.output_type
class GetPolicyResult:
    """
    A Policy.
    """
    def __init__(__self__, created_date=None, description=None, evaluator_type=None, fact_data=None, fact_name=None, location=None, name=None, provisioning_state=None, status=None, tags=None, threshold=None, type=None, unique_identifier=None):
        if created_date and not isinstance(created_date, str):
            raise TypeError("Expected argument 'created_date' to be a str")
        pulumi.set(__self__, "created_date", created_date)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if evaluator_type and not isinstance(evaluator_type, str):
            raise TypeError("Expected argument 'evaluator_type' to be a str")
        pulumi.set(__self__, "evaluator_type", evaluator_type)
        if fact_data and not isinstance(fact_data, str):
            raise TypeError("Expected argument 'fact_data' to be a str")
        pulumi.set(__self__, "fact_data", fact_data)
        if fact_name and not isinstance(fact_name, str):
            raise TypeError("Expected argument 'fact_name' to be a str")
        pulumi.set(__self__, "fact_name", fact_name)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if threshold and not isinstance(threshold, str):
            raise TypeError("Expected argument 'threshold' to be a str")
        pulumi.set(__self__, "threshold", threshold)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if unique_identifier and not isinstance(unique_identifier, str):
            raise TypeError("Expected argument 'unique_identifier' to be a str")
        pulumi.set(__self__, "unique_identifier", unique_identifier)

    @property
    @pulumi.getter(name="createdDate")
    def created_date(self) -> str:
        """
        The creation date of the policy.
        """
        return pulumi.get(self, "created_date")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description of the policy.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="evaluatorType")
    def evaluator_type(self) -> Optional[str]:
        """
        The evaluator type of the policy (i.e. AllowedValuesPolicy, MaxValuePolicy).
        """
        return pulumi.get(self, "evaluator_type")

    @property
    @pulumi.getter(name="factData")
    def fact_data(self) -> Optional[str]:
        """
        The fact data of the policy.
        """
        return pulumi.get(self, "fact_data")

    @property
    @pulumi.getter(name="factName")
    def fact_name(self) -> Optional[str]:
        """
        The fact name of the policy (e.g. LabVmCount, LabVmSize, MaxVmsAllowedPerLab, etc.
        """
        return pulumi.get(self, "fact_name")

    @property
    @pulumi.getter
    def location(self) -> Optional[str]:
        """
        The location of the resource.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        The provisioning status of the resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        The status of the policy.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        The tags of the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def threshold(self) -> Optional[str]:
        """
        The threshold of the policy (i.e. a number for MaxValuePolicy, and a JSON array of values for AllowedValuesPolicy).
        """
        return pulumi.get(self, "threshold")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the resource.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="uniqueIdentifier")
    def unique_identifier(self) -> str:
        """
        The unique immutable identifier of a resource (Guid).
        """
        return pulumi.get(self, "unique_identifier")


class AwaitableGetPolicyResult(GetPolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPolicyResult(
            created_date=self.created_date,
            description=self.description,
            evaluator_type=self.evaluator_type,
            fact_data=self.fact_data,
            fact_name=self.fact_name,
            location=self.location,
            name=self.name,
            provisioning_state=self.provisioning_state,
            status=self.status,
            tags=self.tags,
            threshold=self.threshold,
            type=self.type,
            unique_identifier=self.unique_identifier)


def get_policy(expand: Optional[str] = None,
               lab_name: Optional[str] = None,
               name: Optional[str] = None,
               policy_set_name: Optional[str] = None,
               resource_group_name: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPolicyResult:
    """
    Use this data source to access information about an existing resource.

    :param str expand: Specify the $expand query. Example: 'properties($select=description)'
    :param str lab_name: The name of the lab.
    :param str name: The name of the policy.
    :param str policy_set_name: The name of the policy set.
    :param str resource_group_name: The name of the resource group.
    """
    __args__ = dict()
    __args__['expand'] = expand
    __args__['labName'] = lab_name
    __args__['name'] = name
    __args__['policySetName'] = policy_set_name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:devtestlab/v20180915:getPolicy', __args__, opts=opts, typ=GetPolicyResult).value

    return AwaitableGetPolicyResult(
        created_date=__ret__.created_date,
        description=__ret__.description,
        evaluator_type=__ret__.evaluator_type,
        fact_data=__ret__.fact_data,
        fact_name=__ret__.fact_name,
        location=__ret__.location,
        name=__ret__.name,
        provisioning_state=__ret__.provisioning_state,
        status=__ret__.status,
        tags=__ret__.tags,
        threshold=__ret__.threshold,
        type=__ret__.type,
        unique_identifier=__ret__.unique_identifier)
