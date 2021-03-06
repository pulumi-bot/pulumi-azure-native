# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs

__all__ = [
    'GetRoleManagementPolicyResult',
    'AwaitableGetRoleManagementPolicyResult',
    'get_role_management_policy',
]

@pulumi.output_type
class GetRoleManagementPolicyResult:
    """
    Role management policy
    """
    def __init__(__self__, description=None, display_name=None, effective_rules=None, id=None, is_organization_default=None, last_updated_date_time=None, name=None, rules=None, scope=None, type=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if effective_rules and not isinstance(effective_rules, list):
            raise TypeError("Expected argument 'effective_rules' to be a list")
        pulumi.set(__self__, "effective_rules", effective_rules)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_organization_default and not isinstance(is_organization_default, bool):
            raise TypeError("Expected argument 'is_organization_default' to be a bool")
        pulumi.set(__self__, "is_organization_default", is_organization_default)
        if last_updated_date_time and not isinstance(last_updated_date_time, str):
            raise TypeError("Expected argument 'last_updated_date_time' to be a str")
        pulumi.set(__self__, "last_updated_date_time", last_updated_date_time)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if rules and not isinstance(rules, list):
            raise TypeError("Expected argument 'rules' to be a list")
        pulumi.set(__self__, "rules", rules)
        if scope and not isinstance(scope, str):
            raise TypeError("Expected argument 'scope' to be a str")
        pulumi.set(__self__, "scope", scope)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The role management policy description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[str]:
        """
        The role management policy display name.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="effectiveRules")
    def effective_rules(self) -> Sequence[Any]:
        """
        The readonly computed rule applied to the policy.
        """
        return pulumi.get(self, "effective_rules")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The role management policy Id.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isOrganizationDefault")
    def is_organization_default(self) -> Optional[bool]:
        """
        The role management policy is default policy.
        """
        return pulumi.get(self, "is_organization_default")

    @property
    @pulumi.getter(name="lastUpdatedDateTime")
    def last_updated_date_time(self) -> str:
        """
        The last updated date time.
        """
        return pulumi.get(self, "last_updated_date_time")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The role management policy name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def rules(self) -> Optional[Sequence[Any]]:
        """
        The rule applied to the policy.
        """
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter
    def scope(self) -> Optional[str]:
        """
        The role management policy scope.
        """
        return pulumi.get(self, "scope")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The role management policy type.
        """
        return pulumi.get(self, "type")


class AwaitableGetRoleManagementPolicyResult(GetRoleManagementPolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRoleManagementPolicyResult(
            description=self.description,
            display_name=self.display_name,
            effective_rules=self.effective_rules,
            id=self.id,
            is_organization_default=self.is_organization_default,
            last_updated_date_time=self.last_updated_date_time,
            name=self.name,
            rules=self.rules,
            scope=self.scope,
            type=self.type)


def get_role_management_policy(role_management_policy_name: Optional[str] = None,
                               scope: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRoleManagementPolicyResult:
    """
    Role management policy


    :param str role_management_policy_name: The name (guid) of the role management policy to get.
    :param str scope: The scope of the role management policy.
    """
    __args__ = dict()
    __args__['roleManagementPolicyName'] = role_management_policy_name
    __args__['scope'] = scope
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure-native:authorization/v20201001preview:getRoleManagementPolicy', __args__, opts=opts, typ=GetRoleManagementPolicyResult).value

    return AwaitableGetRoleManagementPolicyResult(
        description=__ret__.description,
        display_name=__ret__.display_name,
        effective_rules=__ret__.effective_rules,
        id=__ret__.id,
        is_organization_default=__ret__.is_organization_default,
        last_updated_date_time=__ret__.last_updated_date_time,
        name=__ret__.name,
        rules=__ret__.rules,
        scope=__ret__.scope,
        type=__ret__.type)
