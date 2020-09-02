# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'PolicyAssignmentPropertiesArgs',
    'PolicyDefinitionPropertiesArgs',
]

@pulumi.input_type
class PolicyAssignmentPropertiesArgs:
    def __init__(__self__, *,
                 display_name: Optional[pulumi.Input[str]] = None,
                 policy_definition_id: Optional[pulumi.Input[str]] = None,
                 scope: Optional[pulumi.Input[str]] = None):
        """
        Policy Assignment properties.
        :param pulumi.Input[str] display_name: Gets or sets the policy assignment display name.
        :param pulumi.Input[str] policy_definition_id: Gets or sets the policy definition Id.
        :param pulumi.Input[str] scope: Gets or sets the policy assignment scope.
        """
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if policy_definition_id is not None:
            pulumi.set(__self__, "policy_definition_id", policy_definition_id)
        if scope is not None:
            pulumi.set(__self__, "scope", scope)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Gets or sets the policy assignment display name.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="policyDefinitionId")
    def policy_definition_id(self) -> Optional[pulumi.Input[str]]:
        """
        Gets or sets the policy definition Id.
        """
        return pulumi.get(self, "policy_definition_id")

    @policy_definition_id.setter
    def policy_definition_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_definition_id", value)

    @property
    @pulumi.getter
    def scope(self) -> Optional[pulumi.Input[str]]:
        """
        Gets or sets the policy assignment scope.
        """
        return pulumi.get(self, "scope")

    @scope.setter
    def scope(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope", value)


@pulumi.input_type
class PolicyDefinitionPropertiesArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 policy_rule: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        Policy definition properties.
        :param pulumi.Input[str] description: Gets or sets the policy definition description.
        :param pulumi.Input[str] display_name: Gets or sets the policy definition display name.
        :param pulumi.Input[Mapping[str, Any]] policy_rule: The policy rule json.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if policy_rule is not None:
            pulumi.set(__self__, "policy_rule", policy_rule)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Gets or sets the policy definition description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Gets or sets the policy definition display name.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="policyRule")
    def policy_rule(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        The policy rule json.
        """
        return pulumi.get(self, "policy_rule")

    @policy_rule.setter
    def policy_rule(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "policy_rule", value)

