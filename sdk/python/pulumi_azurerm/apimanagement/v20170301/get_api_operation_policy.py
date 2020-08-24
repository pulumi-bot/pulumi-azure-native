# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'GetApiOperationPolicyResult',
    'AwaitableGetApiOperationPolicyResult',
    'get_api_operation_policy',
]

@pulumi.output_type
class GetApiOperationPolicyResult:
    """
    Policy Contract details.
    """
    def __init__(__self__, name=None, policy_content=None, type=None):
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if policy_content and not isinstance(policy_content, str):
            raise TypeError("Expected argument 'policy_content' to be a str")
        pulumi.set(__self__, "policy_content", policy_content)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="policyContent")
    def policy_content(self) -> str:
        """
        Json escaped Xml Encoded contents of the Policy.
        """
        return pulumi.get(self, "policy_content")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type for API Management resource.
        """
        return pulumi.get(self, "type")


class AwaitableGetApiOperationPolicyResult(GetApiOperationPolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetApiOperationPolicyResult(
            name=self.name,
            policy_content=self.policy_content,
            type=self.type)


def get_api_operation_policy(api_id: Optional[str] = None,
                             name: Optional[str] = None,
                             operation_id: Optional[str] = None,
                             resource_group_name: Optional[str] = None,
                             service_name: Optional[str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetApiOperationPolicyResult:
    """
    Use this data source to access information about an existing resource.

    :param str api_id: API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
    :param str name: The identifier of the Policy.
    :param str operation_id: Operation identifier within an API. Must be unique in the current API Management service instance.
    :param str resource_group_name: The name of the resource group.
    :param str service_name: The name of the API Management service.
    """
    __args__ = dict()
    __args__['apiId'] = api_id
    __args__['name'] = name
    __args__['operationId'] = operation_id
    __args__['resourceGroupName'] = resource_group_name
    __args__['serviceName'] = service_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:apimanagement/v20170301:getApiOperationPolicy', __args__, opts=opts, typ=GetApiOperationPolicyResult).value

    return AwaitableGetApiOperationPolicyResult(
        name=__ret__.name,
        policy_content=__ret__.policy_content,
        type=__ret__.type)
