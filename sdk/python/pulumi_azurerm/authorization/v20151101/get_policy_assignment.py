# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class GetPolicyAssignmentResult:
    """
    Policy assignment.
    """
    def __init__(__self__, name=None, properties=None):
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        Gets or sets the policy assignment name.
        """
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        __self__.properties = properties
        """
        Gets or sets the policy assignment properties.
        """


class AwaitableGetPolicyAssignmentResult(GetPolicyAssignmentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPolicyAssignmentResult(
            name=self.name,
            properties=self.properties)


def get_policy_assignment(name=None, scope=None, opts=None):
    """
    Use this data source to access information about an existing resource.

    :param str name: Policy assignment name.
    :param str scope: Scope.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['scope'] = scope
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:authorization/v20151101:getPolicyAssignment', __args__, opts=opts).value

    return AwaitableGetPolicyAssignmentResult(
        name=__ret__.get('name'),
        properties=__ret__.get('properties'))
