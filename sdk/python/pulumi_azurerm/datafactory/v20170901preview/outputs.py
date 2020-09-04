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
    'ActivityDependencyResponse',
    'ActivityResponse',
    'DatasetResponse',
    'FactoryIdentityResponse',
    'FactoryVSTSConfigurationResponse',
    'IntegrationRuntimeReferenceResponse',
    'IntegrationRuntimeResponse',
    'LinkedServiceReferenceResponse',
    'LinkedServiceResponse',
    'ParameterSpecificationResponse',
    'TriggerResponse',
]

@pulumi.output_type
class ActivityDependencyResponse(dict):
    """
    Activity dependency information.
    """
    def __init__(__self__, *,
                 activity: str,
                 dependency_conditions: List[str]):
        """
        Activity dependency information.
        :param str activity: Activity name.
        :param List[str] dependency_conditions: Match-Condition for the dependency.
        """
        pulumi.set(__self__, "activity", activity)
        pulumi.set(__self__, "dependency_conditions", dependency_conditions)

    @property
    @pulumi.getter
    def activity(self) -> str:
        """
        Activity name.
        """
        return pulumi.get(self, "activity")

    @property
    @pulumi.getter(name="dependencyConditions")
    def dependency_conditions(self) -> List[str]:
        """
        Match-Condition for the dependency.
        """
        return pulumi.get(self, "dependency_conditions")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ActivityResponse(dict):
    """
    A pipeline activity.
    """
    def __init__(__self__, *,
                 name: str,
                 type: str,
                 depends_on: Optional[List['outputs.ActivityDependencyResponse']] = None,
                 description: Optional[str] = None):
        """
        A pipeline activity.
        :param str name: Activity name.
        :param str type: Type of activity.
        :param List['ActivityDependencyResponseArgs'] depends_on: Activity depends on condition.
        :param str description: Activity description.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "type", type)
        if depends_on is not None:
            pulumi.set(__self__, "depends_on", depends_on)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Activity name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Type of activity.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="dependsOn")
    def depends_on(self) -> Optional[List['outputs.ActivityDependencyResponse']]:
        """
        Activity depends on condition.
        """
        return pulumi.get(self, "depends_on")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        Activity description.
        """
        return pulumi.get(self, "description")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class DatasetResponse(dict):
    """
    The Azure Data Factory nested object which identifies data within different data stores, such as tables, files, folders, and documents.
    """
    def __init__(__self__, *,
                 linked_service_name: 'outputs.LinkedServiceReferenceResponse',
                 type: str,
                 annotations: Optional[List[Mapping[str, Any]]] = None,
                 description: Optional[str] = None,
                 parameters: Optional[Mapping[str, 'outputs.ParameterSpecificationResponse']] = None,
                 structure: Optional[Mapping[str, Any]] = None):
        """
        The Azure Data Factory nested object which identifies data within different data stores, such as tables, files, folders, and documents.
        :param 'LinkedServiceReferenceResponseArgs' linked_service_name: Linked service reference.
        :param str type: Type of dataset.
        :param List[Mapping[str, Any]] annotations: List of tags that can be used for describing the Dataset.
        :param str description: Dataset description.
        :param Mapping[str, 'ParameterSpecificationResponseArgs'] parameters: Parameters for dataset.
        :param Mapping[str, Any] structure: Columns that define the structure of the dataset. Type: array (or Expression with resultType array), itemType: DatasetDataElement.
        """
        pulumi.set(__self__, "linked_service_name", linked_service_name)
        pulumi.set(__self__, "type", type)
        if annotations is not None:
            pulumi.set(__self__, "annotations", annotations)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)
        if structure is not None:
            pulumi.set(__self__, "structure", structure)

    @property
    @pulumi.getter(name="linkedServiceName")
    def linked_service_name(self) -> 'outputs.LinkedServiceReferenceResponse':
        """
        Linked service reference.
        """
        return pulumi.get(self, "linked_service_name")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Type of dataset.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def annotations(self) -> Optional[List[Mapping[str, Any]]]:
        """
        List of tags that can be used for describing the Dataset.
        """
        return pulumi.get(self, "annotations")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        Dataset description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def parameters(self) -> Optional[Mapping[str, 'outputs.ParameterSpecificationResponse']]:
        """
        Parameters for dataset.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter
    def structure(self) -> Optional[Mapping[str, Any]]:
        """
        Columns that define the structure of the dataset. Type: array (or Expression with resultType array), itemType: DatasetDataElement.
        """
        return pulumi.get(self, "structure")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class FactoryIdentityResponse(dict):
    """
    Identity properties of the factory resource.
    """
    def __init__(__self__, *,
                 principal_id: str,
                 tenant_id: str,
                 type: str):
        """
        Identity properties of the factory resource.
        :param str principal_id: The principal id of the identity.
        :param str tenant_id: The client tenant id of the identity.
        :param str type: The identity type. Currently the only supported type is 'SystemAssigned'.
        """
        pulumi.set(__self__, "principal_id", principal_id)
        pulumi.set(__self__, "tenant_id", tenant_id)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> str:
        """
        The principal id of the identity.
        """
        return pulumi.get(self, "principal_id")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> str:
        """
        The client tenant id of the identity.
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The identity type. Currently the only supported type is 'SystemAssigned'.
        """
        return pulumi.get(self, "type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class FactoryVSTSConfigurationResponse(dict):
    """
    Factory's VSTS repo information.
    """
    def __init__(__self__, *,
                 account_name: Optional[str] = None,
                 collaboration_branch: Optional[str] = None,
                 last_commit_id: Optional[str] = None,
                 project_name: Optional[str] = None,
                 repository_name: Optional[str] = None,
                 root_folder: Optional[str] = None,
                 tenant_id: Optional[str] = None):
        """
        Factory's VSTS repo information.
        :param str account_name: VSTS account name.
        :param str collaboration_branch: VSTS collaboration branch.
        :param str last_commit_id: VSTS last commit id.
        :param str project_name: VSTS project name.
        :param str repository_name: VSTS repository name.
        :param str root_folder: VSTS root folder.
        :param str tenant_id: VSTS tenant id.
        """
        if account_name is not None:
            pulumi.set(__self__, "account_name", account_name)
        if collaboration_branch is not None:
            pulumi.set(__self__, "collaboration_branch", collaboration_branch)
        if last_commit_id is not None:
            pulumi.set(__self__, "last_commit_id", last_commit_id)
        if project_name is not None:
            pulumi.set(__self__, "project_name", project_name)
        if repository_name is not None:
            pulumi.set(__self__, "repository_name", repository_name)
        if root_folder is not None:
            pulumi.set(__self__, "root_folder", root_folder)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> Optional[str]:
        """
        VSTS account name.
        """
        return pulumi.get(self, "account_name")

    @property
    @pulumi.getter(name="collaborationBranch")
    def collaboration_branch(self) -> Optional[str]:
        """
        VSTS collaboration branch.
        """
        return pulumi.get(self, "collaboration_branch")

    @property
    @pulumi.getter(name="lastCommitId")
    def last_commit_id(self) -> Optional[str]:
        """
        VSTS last commit id.
        """
        return pulumi.get(self, "last_commit_id")

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> Optional[str]:
        """
        VSTS project name.
        """
        return pulumi.get(self, "project_name")

    @property
    @pulumi.getter(name="repositoryName")
    def repository_name(self) -> Optional[str]:
        """
        VSTS repository name.
        """
        return pulumi.get(self, "repository_name")

    @property
    @pulumi.getter(name="rootFolder")
    def root_folder(self) -> Optional[str]:
        """
        VSTS root folder.
        """
        return pulumi.get(self, "root_folder")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[str]:
        """
        VSTS tenant id.
        """
        return pulumi.get(self, "tenant_id")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class IntegrationRuntimeReferenceResponse(dict):
    """
    Integration runtime reference type.
    """
    def __init__(__self__, *,
                 reference_name: str,
                 type: str,
                 parameters: Optional[Mapping[str, Mapping[str, Any]]] = None):
        """
        Integration runtime reference type.
        :param str reference_name: Reference integration runtime name.
        :param str type: Type of integration runtime.
        :param Mapping[str, Mapping[str, Any]] parameters: Arguments for integration runtime.
        """
        pulumi.set(__self__, "reference_name", reference_name)
        pulumi.set(__self__, "type", type)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)

    @property
    @pulumi.getter(name="referenceName")
    def reference_name(self) -> str:
        """
        Reference integration runtime name.
        """
        return pulumi.get(self, "reference_name")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Type of integration runtime.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def parameters(self) -> Optional[Mapping[str, Mapping[str, Any]]]:
        """
        Arguments for integration runtime.
        """
        return pulumi.get(self, "parameters")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class IntegrationRuntimeResponse(dict):
    """
    Azure Data Factory nested object which serves as a compute resource for activities.
    """
    def __init__(__self__, *,
                 type: str,
                 description: Optional[str] = None):
        """
        Azure Data Factory nested object which serves as a compute resource for activities.
        :param str type: Type of integration runtime.
        :param str description: Integration runtime description.
        """
        pulumi.set(__self__, "type", type)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Type of integration runtime.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        Integration runtime description.
        """
        return pulumi.get(self, "description")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class LinkedServiceReferenceResponse(dict):
    """
    Linked service reference type.
    """
    def __init__(__self__, *,
                 reference_name: str,
                 type: str,
                 parameters: Optional[Mapping[str, Mapping[str, Any]]] = None):
        """
        Linked service reference type.
        :param str reference_name: Reference LinkedService name.
        :param str type: Linked service reference type.
        :param Mapping[str, Mapping[str, Any]] parameters: Arguments for LinkedService.
        """
        pulumi.set(__self__, "reference_name", reference_name)
        pulumi.set(__self__, "type", type)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)

    @property
    @pulumi.getter(name="referenceName")
    def reference_name(self) -> str:
        """
        Reference LinkedService name.
        """
        return pulumi.get(self, "reference_name")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Linked service reference type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def parameters(self) -> Optional[Mapping[str, Mapping[str, Any]]]:
        """
        Arguments for LinkedService.
        """
        return pulumi.get(self, "parameters")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class LinkedServiceResponse(dict):
    """
    The Azure Data Factory nested object which contains the information and credential which can be used to connect with related store or compute resource.
    """
    def __init__(__self__, *,
                 type: str,
                 annotations: Optional[List[Mapping[str, Any]]] = None,
                 connect_via: Optional['outputs.IntegrationRuntimeReferenceResponse'] = None,
                 description: Optional[str] = None,
                 parameters: Optional[Mapping[str, 'outputs.ParameterSpecificationResponse']] = None):
        """
        The Azure Data Factory nested object which contains the information and credential which can be used to connect with related store or compute resource.
        :param str type: Type of linked service.
        :param List[Mapping[str, Any]] annotations: List of tags that can be used for describing the Dataset.
        :param 'IntegrationRuntimeReferenceResponseArgs' connect_via: The integration runtime reference.
        :param str description: Linked service description.
        :param Mapping[str, 'ParameterSpecificationResponseArgs'] parameters: Parameters for linked service.
        """
        pulumi.set(__self__, "type", type)
        if annotations is not None:
            pulumi.set(__self__, "annotations", annotations)
        if connect_via is not None:
            pulumi.set(__self__, "connect_via", connect_via)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Type of linked service.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def annotations(self) -> Optional[List[Mapping[str, Any]]]:
        """
        List of tags that can be used for describing the Dataset.
        """
        return pulumi.get(self, "annotations")

    @property
    @pulumi.getter(name="connectVia")
    def connect_via(self) -> Optional['outputs.IntegrationRuntimeReferenceResponse']:
        """
        The integration runtime reference.
        """
        return pulumi.get(self, "connect_via")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        Linked service description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def parameters(self) -> Optional[Mapping[str, 'outputs.ParameterSpecificationResponse']]:
        """
        Parameters for linked service.
        """
        return pulumi.get(self, "parameters")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ParameterSpecificationResponse(dict):
    """
    Definition of a single parameter for an entity.
    """
    def __init__(__self__, *,
                 type: str,
                 default_value: Optional[Mapping[str, Any]] = None):
        """
        Definition of a single parameter for an entity.
        :param str type: Parameter type.
        :param Mapping[str, Any] default_value: Default value of parameter.
        """
        pulumi.set(__self__, "type", type)
        if default_value is not None:
            pulumi.set(__self__, "default_value", default_value)

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Parameter type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="defaultValue")
    def default_value(self) -> Optional[Mapping[str, Any]]:
        """
        Default value of parameter.
        """
        return pulumi.get(self, "default_value")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class TriggerResponse(dict):
    """
    Azure data factory nested object which contains information about creating pipeline run
    """
    def __init__(__self__, *,
                 runtime_state: str,
                 type: str,
                 description: Optional[str] = None):
        """
        Azure data factory nested object which contains information about creating pipeline run
        :param str runtime_state: Indicates if trigger is running or not. Updated when Start/Stop APIs are called on the Trigger.
        :param str type: Trigger type.
        :param str description: Trigger description.
        """
        pulumi.set(__self__, "runtime_state", runtime_state)
        pulumi.set(__self__, "type", type)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter(name="runtimeState")
    def runtime_state(self) -> str:
        """
        Indicates if trigger is running or not. Updated when Start/Stop APIs are called on the Trigger.
        """
        return pulumi.get(self, "runtime_state")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Trigger type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        Trigger description.
        """
        return pulumi.get(self, "description")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


