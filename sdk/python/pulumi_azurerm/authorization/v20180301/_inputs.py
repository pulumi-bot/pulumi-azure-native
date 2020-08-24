# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'PolicyDefinitionReferenceArgs',
    'PolicySkuArgs',
]

@pulumi.input_type
class PolicyDefinitionReferenceArgs:
    def __init__(__self__, *,
                 parameters: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 policy_definition_id: Optional[pulumi.Input[str]] = None):
        """
        The policy definition reference.
        :param pulumi.Input[Mapping[str, Any]] parameters: Required if a parameter is used in policy rule.
        :param pulumi.Input[str] policy_definition_id: The ID of the policy definition or policy set definition.
        """
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)
        if policy_definition_id is not None:
            pulumi.set(__self__, "policy_definition_id", policy_definition_id)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Required if a parameter is used in policy rule.
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "parameters", value)

    @property
    @pulumi.getter(name="policyDefinitionId")
    def policy_definition_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the policy definition or policy set definition.
        """
        return pulumi.get(self, "policy_definition_id")

    @policy_definition_id.setter
    def policy_definition_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_definition_id", value)


@pulumi.input_type
class PolicySkuArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 tier: Optional[pulumi.Input[str]] = None):
        """
        The policy sku. This property is optional, obsolete, and will be ignored.
        :param pulumi.Input[str] name: The name of the policy sku. Possible values are A0 and A1.
        :param pulumi.Input[str] tier: The policy sku tier. Possible values are Free and Standard.
        """
        pulumi.set(__self__, "name", name)
        if tier is not None:
            pulumi.set(__self__, "tier", tier)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The name of the policy sku. Possible values are A0 and A1.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tier(self) -> Optional[pulumi.Input[str]]:
        """
        The policy sku tier. Possible values are Free and Standard.
        """
        return pulumi.get(self, "tier")

    @tier.setter
    def tier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tier", value)


