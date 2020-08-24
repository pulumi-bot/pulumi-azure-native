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
    'IdentityResponse',
    'ParameterDefinitionsValueResponse',
    'ParameterDefinitionsValueResponseMetadata',
    'ParameterValuesValueResponse',
    'PolicyDefinitionGroupResponse',
    'PolicyDefinitionReferenceResponse',
    'PolicySkuResponse',
]

@pulumi.output_type
class IdentityResponse(dict):
    """
    Identity for the resource.
    """
    def __init__(__self__, *,
                 principal_id: str,
                 tenant_id: str,
                 type: Optional[str] = None):
        """
        Identity for the resource.
        :param str principal_id: The principal ID of the resource identity.
        :param str tenant_id: The tenant ID of the resource identity.
        :param str type: The identity type. This is the only required field when adding a system assigned identity to a resource.
        """
        pulumi.set(__self__, "principal_id", principal_id)
        pulumi.set(__self__, "tenant_id", tenant_id)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> str:
        """
        The principal ID of the resource identity.
        """
        return pulumi.get(self, "principal_id")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> str:
        """
        The tenant ID of the resource identity.
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        The identity type. This is the only required field when adding a system assigned identity to a resource.
        """
        return pulumi.get(self, "type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ParameterDefinitionsValueResponse(dict):
    """
    The definition of a parameter that can be provided to the policy.
    """
    def __init__(__self__, *,
                 allowed_values: Optional[List[Mapping[str, Any]]] = None,
                 default_value: Optional[Mapping[str, Any]] = None,
                 metadata: Optional['outputs.ParameterDefinitionsValueResponseMetadata'] = None,
                 type: Optional[str] = None):
        """
        The definition of a parameter that can be provided to the policy.
        :param List[Mapping[str, Any]] allowed_values: The allowed values for the parameter.
        :param Mapping[str, Any] default_value: The default value for the parameter if no value is provided.
        :param 'ParameterDefinitionsValueResponseMetadataArgs' metadata: General metadata for the parameter.
        :param str type: The data type of the parameter.
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
    def allowed_values(self) -> Optional[List[Mapping[str, Any]]]:
        """
        The allowed values for the parameter.
        """
        return pulumi.get(self, "allowed_values")

    @property
    @pulumi.getter(name="defaultValue")
    def default_value(self) -> Optional[Mapping[str, Any]]:
        """
        The default value for the parameter if no value is provided.
        """
        return pulumi.get(self, "default_value")

    @property
    @pulumi.getter
    def metadata(self) -> Optional['outputs.ParameterDefinitionsValueResponseMetadata']:
        """
        General metadata for the parameter.
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        The data type of the parameter.
        """
        return pulumi.get(self, "type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ParameterDefinitionsValueResponseMetadata(dict):
    """
    General metadata for the parameter.
    """
    def __init__(__self__, *,
                 description: Optional[str] = None,
                 display_name: Optional[str] = None):
        """
        General metadata for the parameter.
        :param str description: The description of the parameter.
        :param str display_name: The display name for the parameter.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description of the parameter.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[str]:
        """
        The display name for the parameter.
        """
        return pulumi.get(self, "display_name")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ParameterValuesValueResponse(dict):
    """
    The value of a parameter.
    """
    def __init__(__self__, *,
                 value: Optional[Mapping[str, Any]] = None):
        """
        The value of a parameter.
        :param Mapping[str, Any] value: The value of the parameter.
        """
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[Mapping[str, Any]]:
        """
        The value of the parameter.
        """
        return pulumi.get(self, "value")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PolicyDefinitionGroupResponse(dict):
    """
    The policy definition group.
    """
    def __init__(__self__, *,
                 name: str,
                 additional_metadata_id: Optional[str] = None,
                 category: Optional[str] = None,
                 description: Optional[str] = None,
                 display_name: Optional[str] = None):
        """
        The policy definition group.
        :param str name: The name of the group.
        :param str additional_metadata_id: A resource ID of a resource that contains additional metadata about the group.
        :param str category: The group's category.
        :param str description: The group's description.
        :param str display_name: The group's display name.
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
    def name(self) -> str:
        """
        The name of the group.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="additionalMetadataId")
    def additional_metadata_id(self) -> Optional[str]:
        """
        A resource ID of a resource that contains additional metadata about the group.
        """
        return pulumi.get(self, "additional_metadata_id")

    @property
    @pulumi.getter
    def category(self) -> Optional[str]:
        """
        The group's category.
        """
        return pulumi.get(self, "category")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The group's description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[str]:
        """
        The group's display name.
        """
        return pulumi.get(self, "display_name")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PolicyDefinitionReferenceResponse(dict):
    """
    The policy definition reference.
    """
    def __init__(__self__, *,
                 policy_definition_id: str,
                 group_names: Optional[List[str]] = None,
                 parameters: Optional[Mapping[str, 'outputs.ParameterValuesValueResponse']] = None,
                 policy_definition_reference_id: Optional[str] = None):
        """
        The policy definition reference.
        :param str policy_definition_id: The ID of the policy definition or policy set definition.
        :param List[str] group_names: The name of the groups that this policy definition reference belongs to.
        :param Mapping[str, 'ParameterValuesValueResponseArgs'] parameters: The parameter values for the referenced policy rule. The keys are the parameter names.
        :param str policy_definition_reference_id: A unique id (within the policy set definition) for this policy definition reference.
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
    def policy_definition_id(self) -> str:
        """
        The ID of the policy definition or policy set definition.
        """
        return pulumi.get(self, "policy_definition_id")

    @property
    @pulumi.getter(name="groupNames")
    def group_names(self) -> Optional[List[str]]:
        """
        The name of the groups that this policy definition reference belongs to.
        """
        return pulumi.get(self, "group_names")

    @property
    @pulumi.getter
    def parameters(self) -> Optional[Mapping[str, 'outputs.ParameterValuesValueResponse']]:
        """
        The parameter values for the referenced policy rule. The keys are the parameter names.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter(name="policyDefinitionReferenceId")
    def policy_definition_reference_id(self) -> Optional[str]:
        """
        A unique id (within the policy set definition) for this policy definition reference.
        """
        return pulumi.get(self, "policy_definition_reference_id")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PolicySkuResponse(dict):
    """
    The policy sku. This property is optional, obsolete, and will be ignored.
    """
    def __init__(__self__, *,
                 name: str,
                 tier: Optional[str] = None):
        """
        The policy sku. This property is optional, obsolete, and will be ignored.
        :param str name: The name of the policy sku. Possible values are A0 and A1.
        :param str tier: The policy sku tier. Possible values are Free and Standard.
        """
        pulumi.set(__self__, "name", name)
        if tier is not None:
            pulumi.set(__self__, "tier", tier)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the policy sku. Possible values are A0 and A1.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tier(self) -> Optional[str]:
        """
        The policy sku tier. Possible values are Free and Standard.
        """
        return pulumi.get(self, "tier")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


