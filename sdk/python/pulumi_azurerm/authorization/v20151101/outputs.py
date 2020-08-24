# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'PolicyAssignmentPropertiesResponse',
    'PolicyDefinitionPropertiesResponse',
]

@pulumi.output_type
class PolicyAssignmentPropertiesResponse(dict):
    """
    Policy Assignment properties.
    """
    def __init__(__self__, *,
                 display_name: Optional[str] = None,
                 policy_definition_id: Optional[str] = None,
                 scope: Optional[str] = None):
        """
        Policy Assignment properties.
        :param str display_name: Gets or sets the policy assignment display name.
        :param str policy_definition_id: Gets or sets the policy definition Id.
        :param str scope: Gets or sets the policy assignment scope.
        """
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if policy_definition_id is not None:
            pulumi.set(__self__, "policy_definition_id", policy_definition_id)
        if scope is not None:
            pulumi.set(__self__, "scope", scope)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[str]:
        """
        Gets or sets the policy assignment display name.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="policyDefinitionId")
    def policy_definition_id(self) -> Optional[str]:
        """
        Gets or sets the policy definition Id.
        """
        return pulumi.get(self, "policy_definition_id")

    @property
    @pulumi.getter
    def scope(self) -> Optional[str]:
        """
        Gets or sets the policy assignment scope.
        """
        return pulumi.get(self, "scope")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PolicyDefinitionPropertiesResponse(dict):
    """
    Policy definition properties.
    """
    def __init__(__self__, *,
                 description: Optional[str] = None,
                 display_name: Optional[str] = None,
                 policy_rule: Optional[Mapping[str, Any]] = None):
        """
        Policy definition properties.
        :param str description: Gets or sets the policy definition description.
        :param str display_name: Gets or sets the policy definition display name.
        :param Mapping[str, Any] policy_rule: The policy rule json.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if policy_rule is not None:
            pulumi.set(__self__, "policy_rule", policy_rule)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        Gets or sets the policy definition description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[str]:
        """
        Gets or sets the policy definition display name.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="policyRule")
    def policy_rule(self) -> Optional[Mapping[str, Any]]:
        """
        The policy rule json.
        """
        return pulumi.get(self, "policy_rule")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


