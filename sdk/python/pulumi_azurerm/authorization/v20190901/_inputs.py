# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'IdentityArgs',
    'ParameterDefinitionsValueArgs',
    'ParameterDefinitionsValueMetadataArgs',
    'ParameterValuesValueArgs',
    'PolicyDefinitionGroupArgs',
    'PolicyDefinitionReferenceArgs',
    'PolicySkuArgs',
]

@pulumi.input_type
class IdentityArgs:
    def __init__(__self__, *,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Identity for the resource.
        :param pulumi.Input[str] type: The identity type. This is the only required field when adding a system assigned identity to a resource.
        """
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The identity type. This is the only required field when adding a system assigned identity to a resource.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class ParameterDefinitionsValueArgs:
    def __init__(__self__, *,
                 allowed_values: Optional[pulumi.Input[List[pulumi.Input[Mapping[str, Any]]]]] = None,
                 default_value: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 metadata: Optional[pulumi.Input['ParameterDefinitionsValueMetadataArgs']] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        The definition of a parameter that can be provided to the policy.
        :param pulumi.Input[List[pulumi.Input[Mapping[str, Any]]]] allowed_values: The allowed values for the parameter.
        :param pulumi.Input[Mapping[str, Any]] default_value: The default value for the parameter if no value is provided.
        :param pulumi.Input['ParameterDefinitionsValueMetadataArgs'] metadata: General metadata for the parameter.
        :param pulumi.Input[str] type: The data type of the parameter.
        """
        if allowed_values is not None:
            pulumi.set(__self__, "allowed_values", allowed_values)
        if default_value is not None:
            pulumi.set(__self__, "default_value", default_value)
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="allowedValues")
    def allowed_values(self) -> Optional[pulumi.Input[List[pulumi.Input[Mapping[str, Any]]]]]:
        """
        The allowed values for the parameter.
        """
        return pulumi.get(self, "allowed_values")

    @allowed_values.setter
    def allowed_values(self, value: Optional[pulumi.Input[List[pulumi.Input[Mapping[str, Any]]]]]):
        pulumi.set(self, "allowed_values", value)

    @property
    @pulumi.getter(name="defaultValue")
    def default_value(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        The default value for the parameter if no value is provided.
        """
        return pulumi.get(self, "default_value")

    @default_value.setter
    def default_value(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "default_value", value)

    @property
    @pulumi.getter
    def metadata(self) -> Optional[pulumi.Input['ParameterDefinitionsValueMetadataArgs']]:
        """
        General metadata for the parameter.
        """
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input['ParameterDefinitionsValueMetadataArgs']]):
        pulumi.set(self, "metadata", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The data type of the parameter.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class ParameterDefinitionsValueMetadataArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None):
        """
        General metadata for the parameter.
        :param pulumi.Input[str] description: The description of the parameter.
        :param pulumi.Input[str] display_name: The display name for the parameter.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the parameter.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        The display name for the parameter.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)


@pulumi.input_type
class ParameterValuesValueArgs:
    def __init__(__self__, *,
                 value: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The value of a parameter.
        :param pulumi.Input[Mapping[str, Any]] value: The value of the parameter.
        """
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        The value of the parameter.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class PolicyDefinitionGroupArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 additional_metadata_id: Optional[pulumi.Input[str]] = None,
                 category: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None):
        """
        The policy definition group.
        :param pulumi.Input[str] name: The name of the group.
        :param pulumi.Input[str] additional_metadata_id: A resource ID of a resource that contains additional metadata about the group.
        :param pulumi.Input[str] category: The group's category.
        :param pulumi.Input[str] description: The group's description.
        :param pulumi.Input[str] display_name: The group's display name.
        """
        pulumi.set(__self__, "name", name)
        if additional_metadata_id is not None:
            pulumi.set(__self__, "additional_metadata_id", additional_metadata_id)
        if category is not None:
            pulumi.set(__self__, "category", category)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The name of the group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="additionalMetadataId")
    def additional_metadata_id(self) -> Optional[pulumi.Input[str]]:
        """
        A resource ID of a resource that contains additional metadata about the group.
        """
        return pulumi.get(self, "additional_metadata_id")

    @additional_metadata_id.setter
    def additional_metadata_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "additional_metadata_id", value)

    @property
    @pulumi.getter
    def category(self) -> Optional[pulumi.Input[str]]:
        """
        The group's category.
        """
        return pulumi.get(self, "category")

    @category.setter
    def category(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "category", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The group's description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        The group's display name.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)


@pulumi.input_type
class PolicyDefinitionReferenceArgs:
    def __init__(__self__, *,
                 policy_definition_id: pulumi.Input[str],
                 group_names: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input['ParameterValuesValueArgs']]]] = None,
                 policy_definition_reference_id: Optional[pulumi.Input[str]] = None):
        """
        The policy definition reference.
        :param pulumi.Input[str] policy_definition_id: The ID of the policy definition or policy set definition.
        :param pulumi.Input[List[pulumi.Input[str]]] group_names: The name of the groups that this policy definition reference belongs to.
        :param pulumi.Input[Mapping[str, pulumi.Input['ParameterValuesValueArgs']]] parameters: The parameter values for the referenced policy rule. The keys are the parameter names.
        :param pulumi.Input[str] policy_definition_reference_id: A unique id (within the policy set definition) for this policy definition reference.
        """
        pulumi.set(__self__, "policy_definition_id", policy_definition_id)
        if group_names is not None:
            pulumi.set(__self__, "group_names", group_names)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)
        if policy_definition_reference_id is not None:
            pulumi.set(__self__, "policy_definition_reference_id", policy_definition_reference_id)

    @property
    @pulumi.getter(name="policyDefinitionId")
    def policy_definition_id(self) -> pulumi.Input[str]:
        """
        The ID of the policy definition or policy set definition.
        """
        return pulumi.get(self, "policy_definition_id")

    @policy_definition_id.setter
    def policy_definition_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy_definition_id", value)

    @property
    @pulumi.getter(name="groupNames")
    def group_names(self) -> Optional[pulumi.Input[List[pulumi.Input[str]]]]:
        """
        The name of the groups that this policy definition reference belongs to.
        """
        return pulumi.get(self, "group_names")

    @group_names.setter
    def group_names(self, value: Optional[pulumi.Input[List[pulumi.Input[str]]]]):
        pulumi.set(self, "group_names", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input['ParameterValuesValueArgs']]]]:
        """
        The parameter values for the referenced policy rule. The keys are the parameter names.
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input['ParameterValuesValueArgs']]]]):
        pulumi.set(self, "parameters", value)

    @property
    @pulumi.getter(name="policyDefinitionReferenceId")
    def policy_definition_reference_id(self) -> Optional[pulumi.Input[str]]:
        """
        A unique id (within the policy set definition) for this policy definition reference.
        """
        return pulumi.get(self, "policy_definition_reference_id")

    @policy_definition_reference_id.setter
    def policy_definition_reference_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_definition_reference_id", value)


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


