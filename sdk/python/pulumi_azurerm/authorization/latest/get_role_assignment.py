# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'GetRoleAssignmentResult',
    'AwaitableGetRoleAssignmentResult',
    'get_role_assignment',
]

@pulumi.output_type
class GetRoleAssignmentResult:
    """
    Role Assignments
    """
    def __init__(__self__, can_delegate=None, condition=None, condition_version=None, description=None, name=None, principal_id=None, principal_type=None, role_definition_id=None, scope=None, type=None):
        if can_delegate and not isinstance(can_delegate, bool):
            raise TypeError("Expected argument 'can_delegate' to be a bool")
        pulumi.set(__self__, "can_delegate", can_delegate)
        if condition and not isinstance(condition, str):
            raise TypeError("Expected argument 'condition' to be a str")
        pulumi.set(__self__, "condition", condition)
        if condition_version and not isinstance(condition_version, str):
            raise TypeError("Expected argument 'condition_version' to be a str")
        pulumi.set(__self__, "condition_version", condition_version)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if principal_id and not isinstance(principal_id, str):
            raise TypeError("Expected argument 'principal_id' to be a str")
        pulumi.set(__self__, "principal_id", principal_id)
        if principal_type and not isinstance(principal_type, str):
            raise TypeError("Expected argument 'principal_type' to be a str")
        pulumi.set(__self__, "principal_type", principal_type)
        if role_definition_id and not isinstance(role_definition_id, str):
            raise TypeError("Expected argument 'role_definition_id' to be a str")
        pulumi.set(__self__, "role_definition_id", role_definition_id)
        if scope and not isinstance(scope, str):
            raise TypeError("Expected argument 'scope' to be a str")
        pulumi.set(__self__, "scope", scope)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="canDelegate")
    def can_delegate(self) -> Optional[bool]:
        """
        The Delegation flag for the role assignment
        """
        return pulumi.get(self, "can_delegate")

    @property
    @pulumi.getter
    def condition(self) -> Optional[str]:
        """
        The conditions on the role assignment. This limits the resources it can be assigned to. e.g.: @Resource[Microsoft.Storage/storageAccounts/blobServices/containers:ContainerName] StringEqualsIgnoreCase 'foo_storage_container'
        """
        return pulumi.get(self, "condition")

    @property
    @pulumi.getter(name="conditionVersion")
    def condition_version(self) -> Optional[str]:
        """
        Version of the condition. Currently accepted value is '2.0'
        """
        return pulumi.get(self, "condition_version")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        Description of role assignment
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The role assignment name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> Optional[str]:
        """
        The principal ID.
        """
        return pulumi.get(self, "principal_id")

    @property
    @pulumi.getter(name="principalType")
    def principal_type(self) -> Optional[str]:
        """
        The principal type of the assigned principal ID.
        """
        return pulumi.get(self, "principal_type")

    @property
    @pulumi.getter(name="roleDefinitionId")
    def role_definition_id(self) -> Optional[str]:
        """
        The role definition ID.
        """
        return pulumi.get(self, "role_definition_id")

    @property
    @pulumi.getter
    def scope(self) -> Optional[str]:
        """
        The role assignment scope.
        """
        return pulumi.get(self, "scope")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The role assignment type.
        """
        return pulumi.get(self, "type")


class AwaitableGetRoleAssignmentResult(GetRoleAssignmentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRoleAssignmentResult(
            can_delegate=self.can_delegate,
            condition=self.condition,
            condition_version=self.condition_version,
            description=self.description,
            name=self.name,
            principal_id=self.principal_id,
            principal_type=self.principal_type,
            role_definition_id=self.role_definition_id,
            scope=self.scope,
            type=self.type)


def get_role_assignment(role_assignment_name: Optional[str] = None,
                        scope: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRoleAssignmentResult:
    """
    Use this data source to access information about an existing resource.

    :param str role_assignment_name: The name of the role assignment to get.
    :param str scope: The scope of the role assignment.
    """
    __args__ = dict()
    __args__['roleAssignmentName'] = role_assignment_name
    __args__['scope'] = scope
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:authorization/latest:getRoleAssignment', __args__, opts=opts, typ=GetRoleAssignmentResult).value

    return AwaitableGetRoleAssignmentResult(
        can_delegate=__ret__.can_delegate,
        condition=__ret__.condition,
        condition_version=__ret__.condition_version,
        description=__ret__.description,
        name=__ret__.name,
        principal_id=__ret__.principal_id,
        principal_type=__ret__.principal_type,
        role_definition_id=__ret__.role_definition_id,
        scope=__ret__.scope,
        type=__ret__.type)
