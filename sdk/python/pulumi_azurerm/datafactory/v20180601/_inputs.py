# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'ActivityArgs',
    'ActivityDependencyArgs',
    'DataFlowArgs',
    'DataFlowFolderArgs',
    'DatasetArgs',
    'DatasetFolderArgs',
    'FactoryIdentityArgs',
    'FactoryRepoConfigurationArgs',
    'GlobalParameterSpecificationArgs',
    'IntegrationRuntimeArgs',
    'IntegrationRuntimeReferenceArgs',
    'LinkedServiceArgs',
    'LinkedServiceReferenceArgs',
    'ManagedPrivateEndpointArgs',
    'ParameterSpecificationArgs',
    'PipelineFolderArgs',
    'TriggerArgs',
    'UserPropertyArgs',
    'VariableSpecificationArgs',
]

@pulumi.input_type
class ActivityArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 type: pulumi.Input[str],
                 depends_on: Optional[pulumi.Input[List[pulumi.Input['ActivityDependencyArgs']]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 user_properties: Optional[pulumi.Input[List[pulumi.Input['UserPropertyArgs']]]] = None):
        """
        A pipeline activity.
        :param pulumi.Input[str] name: Activity name.
        :param pulumi.Input[str] type: Type of activity.
        :param pulumi.Input[List[pulumi.Input['ActivityDependencyArgs']]] depends_on: Activity depends on condition.
        :param pulumi.Input[str] description: Activity description.
        :param pulumi.Input[List[pulumi.Input['UserPropertyArgs']]] user_properties: Activity user properties.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "type", type)
        if depends_on is not None:
            pulumi.set(__self__, "depends_on", depends_on)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if user_properties is not None:
            pulumi.set(__self__, "user_properties", user_properties)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Activity name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Type of activity.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="dependsOn")
    def depends_on(self) -> Optional[pulumi.Input[List[pulumi.Input['ActivityDependencyArgs']]]]:
        """
        Activity depends on condition.
        """
        return pulumi.get(self, "depends_on")

    @depends_on.setter
    def depends_on(self, value: Optional[pulumi.Input[List[pulumi.Input['ActivityDependencyArgs']]]]):
        pulumi.set(self, "depends_on", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Activity description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="userProperties")
    def user_properties(self) -> Optional[pulumi.Input[List[pulumi.Input['UserPropertyArgs']]]]:
        """
        Activity user properties.
        """
        return pulumi.get(self, "user_properties")

    @user_properties.setter
    def user_properties(self, value: Optional[pulumi.Input[List[pulumi.Input['UserPropertyArgs']]]]):
        pulumi.set(self, "user_properties", value)


@pulumi.input_type
class ActivityDependencyArgs:
    def __init__(__self__, *,
                 activity: pulumi.Input[str],
                 dependency_conditions: pulumi.Input[List[pulumi.Input[str]]]):
        """
        Activity dependency information.
        :param pulumi.Input[str] activity: Activity name.
        :param pulumi.Input[List[pulumi.Input[str]]] dependency_conditions: Match-Condition for the dependency.
        """
        pulumi.set(__self__, "activity", activity)
        pulumi.set(__self__, "dependency_conditions", dependency_conditions)

    @property
    @pulumi.getter
    def activity(self) -> pulumi.Input[str]:
        """
        Activity name.
        """
        return pulumi.get(self, "activity")

    @activity.setter
    def activity(self, value: pulumi.Input[str]):
        pulumi.set(self, "activity", value)

    @property
    @pulumi.getter(name="dependencyConditions")
    def dependency_conditions(self) -> pulumi.Input[List[pulumi.Input[str]]]:
        """
        Match-Condition for the dependency.
        """
        return pulumi.get(self, "dependency_conditions")

    @dependency_conditions.setter
    def dependency_conditions(self, value: pulumi.Input[List[pulumi.Input[str]]]):
        pulumi.set(self, "dependency_conditions", value)


@pulumi.input_type
class DataFlowArgs:
    def __init__(__self__, *,
                 annotations: Optional[pulumi.Input[List[pulumi.Input[Mapping[str, Any]]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 folder: Optional[pulumi.Input['DataFlowFolderArgs']] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Azure Data Factory nested object which contains a flow with data movements and transformations.
        :param pulumi.Input[List[pulumi.Input[Mapping[str, Any]]]] annotations: List of tags that can be used for describing the data flow.
        :param pulumi.Input[str] description: The description of the data flow.
        :param pulumi.Input['DataFlowFolderArgs'] folder: The folder that this data flow is in. If not specified, Data flow will appear at the root level.
        :param pulumi.Input[str] type: Type of data flow.
        """
        if annotations is not None:
            pulumi.set(__self__, "annotations", annotations)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if folder is not None:
            pulumi.set(__self__, "folder", folder)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def annotations(self) -> Optional[pulumi.Input[List[pulumi.Input[Mapping[str, Any]]]]]:
        """
        List of tags that can be used for describing the data flow.
        """
        return pulumi.get(self, "annotations")

    @annotations.setter
    def annotations(self, value: Optional[pulumi.Input[List[pulumi.Input[Mapping[str, Any]]]]]):
        pulumi.set(self, "annotations", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the data flow.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def folder(self) -> Optional[pulumi.Input['DataFlowFolderArgs']]:
        """
        The folder that this data flow is in. If not specified, Data flow will appear at the root level.
        """
        return pulumi.get(self, "folder")

    @folder.setter
    def folder(self, value: Optional[pulumi.Input['DataFlowFolderArgs']]):
        pulumi.set(self, "folder", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of data flow.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class DataFlowFolderArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The folder that this data flow is in. If not specified, Data flow will appear at the root level.
        :param pulumi.Input[str] name: The name of the folder that this data flow is in.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the folder that this data flow is in.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class DatasetArgs:
    def __init__(__self__, *,
                 linked_service_name: pulumi.Input['LinkedServiceReferenceArgs'],
                 type: pulumi.Input[str],
                 annotations: Optional[pulumi.Input[List[pulumi.Input[Mapping[str, Any]]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 folder: Optional[pulumi.Input['DatasetFolderArgs']] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input['ParameterSpecificationArgs']]]] = None,
                 schema: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 structure: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The Azure Data Factory nested object which identifies data within different data stores, such as tables, files, folders, and documents.
        :param pulumi.Input['LinkedServiceReferenceArgs'] linked_service_name: Linked service reference.
        :param pulumi.Input[str] type: Type of dataset.
        :param pulumi.Input[List[pulumi.Input[Mapping[str, Any]]]] annotations: List of tags that can be used for describing the Dataset.
        :param pulumi.Input[str] description: Dataset description.
        :param pulumi.Input['DatasetFolderArgs'] folder: The folder that this Dataset is in. If not specified, Dataset will appear at the root level.
        :param pulumi.Input[Mapping[str, pulumi.Input['ParameterSpecificationArgs']]] parameters: Parameters for dataset.
        :param pulumi.Input[Mapping[str, Any]] schema: Columns that define the physical type schema of the dataset. Type: array (or Expression with resultType array), itemType: DatasetSchemaDataElement.
        :param pulumi.Input[Mapping[str, Any]] structure: Columns that define the structure of the dataset. Type: array (or Expression with resultType array), itemType: DatasetDataElement.
        """
        pulumi.set(__self__, "linked_service_name", linked_service_name)
        pulumi.set(__self__, "type", type)
        if annotations is not None:
            pulumi.set(__self__, "annotations", annotations)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if folder is not None:
            pulumi.set(__self__, "folder", folder)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)
        if schema is not None:
            pulumi.set(__self__, "schema", schema)
        if structure is not None:
            pulumi.set(__self__, "structure", structure)

    @property
    @pulumi.getter(name="linkedServiceName")
    def linked_service_name(self) -> pulumi.Input['LinkedServiceReferenceArgs']:
        """
        Linked service reference.
        """
        return pulumi.get(self, "linked_service_name")

    @linked_service_name.setter
    def linked_service_name(self, value: pulumi.Input['LinkedServiceReferenceArgs']):
        pulumi.set(self, "linked_service_name", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Type of dataset.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def annotations(self) -> Optional[pulumi.Input[List[pulumi.Input[Mapping[str, Any]]]]]:
        """
        List of tags that can be used for describing the Dataset.
        """
        return pulumi.get(self, "annotations")

    @annotations.setter
    def annotations(self, value: Optional[pulumi.Input[List[pulumi.Input[Mapping[str, Any]]]]]):
        pulumi.set(self, "annotations", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Dataset description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def folder(self) -> Optional[pulumi.Input['DatasetFolderArgs']]:
        """
        The folder that this Dataset is in. If not specified, Dataset will appear at the root level.
        """
        return pulumi.get(self, "folder")

    @folder.setter
    def folder(self, value: Optional[pulumi.Input['DatasetFolderArgs']]):
        pulumi.set(self, "folder", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input['ParameterSpecificationArgs']]]]:
        """
        Parameters for dataset.
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input['ParameterSpecificationArgs']]]]):
        pulumi.set(self, "parameters", value)

    @property
    @pulumi.getter
    def schema(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Columns that define the physical type schema of the dataset. Type: array (or Expression with resultType array), itemType: DatasetSchemaDataElement.
        """
        return pulumi.get(self, "schema")

    @schema.setter
    def schema(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "schema", value)

    @property
    @pulumi.getter
    def structure(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Columns that define the structure of the dataset. Type: array (or Expression with resultType array), itemType: DatasetDataElement.
        """
        return pulumi.get(self, "structure")

    @structure.setter
    def structure(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "structure", value)


@pulumi.input_type
class DatasetFolderArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The folder that this Dataset is in. If not specified, Dataset will appear at the root level.
        :param pulumi.Input[str] name: The name of the folder that this Dataset is in.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the folder that this Dataset is in.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class FactoryIdentityArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str]):
        """
        Identity properties of the factory resource.
        :param pulumi.Input[str] type: The identity type. Currently the only supported type is 'SystemAssigned'.
        """
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The identity type. Currently the only supported type is 'SystemAssigned'.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class FactoryRepoConfigurationArgs:
    def __init__(__self__, *,
                 account_name: pulumi.Input[str],
                 collaboration_branch: pulumi.Input[str],
                 repository_name: pulumi.Input[str],
                 root_folder: pulumi.Input[str],
                 type: pulumi.Input[str],
                 last_commit_id: Optional[pulumi.Input[str]] = None):
        """
        Factory's git repo information.
        :param pulumi.Input[str] account_name: Account name.
        :param pulumi.Input[str] collaboration_branch: Collaboration branch.
        :param pulumi.Input[str] repository_name: Repository name.
        :param pulumi.Input[str] root_folder: Root folder.
        :param pulumi.Input[str] type: Type of repo configuration.
        :param pulumi.Input[str] last_commit_id: Last commit id.
        """
        pulumi.set(__self__, "account_name", account_name)
        pulumi.set(__self__, "collaboration_branch", collaboration_branch)
        pulumi.set(__self__, "repository_name", repository_name)
        pulumi.set(__self__, "root_folder", root_folder)
        pulumi.set(__self__, "type", type)
        if last_commit_id is not None:
            pulumi.set(__self__, "last_commit_id", last_commit_id)

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> pulumi.Input[str]:
        """
        Account name.
        """
        return pulumi.get(self, "account_name")

    @account_name.setter
    def account_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "account_name", value)

    @property
    @pulumi.getter(name="collaborationBranch")
    def collaboration_branch(self) -> pulumi.Input[str]:
        """
        Collaboration branch.
        """
        return pulumi.get(self, "collaboration_branch")

    @collaboration_branch.setter
    def collaboration_branch(self, value: pulumi.Input[str]):
        pulumi.set(self, "collaboration_branch", value)

    @property
    @pulumi.getter(name="repositoryName")
    def repository_name(self) -> pulumi.Input[str]:
        """
        Repository name.
        """
        return pulumi.get(self, "repository_name")

    @repository_name.setter
    def repository_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "repository_name", value)

    @property
    @pulumi.getter(name="rootFolder")
    def root_folder(self) -> pulumi.Input[str]:
        """
        Root folder.
        """
        return pulumi.get(self, "root_folder")

    @root_folder.setter
    def root_folder(self, value: pulumi.Input[str]):
        pulumi.set(self, "root_folder", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Type of repo configuration.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="lastCommitId")
    def last_commit_id(self) -> Optional[pulumi.Input[str]]:
        """
        Last commit id.
        """
        return pulumi.get(self, "last_commit_id")

    @last_commit_id.setter
    def last_commit_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_commit_id", value)


@pulumi.input_type
class GlobalParameterSpecificationArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 value: pulumi.Input[Mapping[str, Any]]):
        """
        Definition of a single parameter for an entity.
        :param pulumi.Input[str] type: Global Parameter type.
        :param pulumi.Input[Mapping[str, Any]] value: Value of parameter.
        """
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Global Parameter type.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[Mapping[str, Any]]:
        """
        Value of parameter.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[Mapping[str, Any]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class IntegrationRuntimeArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None):
        """
        Azure Data Factory nested object which serves as a compute resource for activities.
        :param pulumi.Input[str] type: Type of integration runtime.
        :param pulumi.Input[str] description: Integration runtime description.
        """
        pulumi.set(__self__, "type", type)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Type of integration runtime.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Integration runtime description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


@pulumi.input_type
class IntegrationRuntimeReferenceArgs:
    def __init__(__self__, *,
                 reference_name: pulumi.Input[str],
                 type: pulumi.Input[str],
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[Mapping[str, Any]]]]] = None):
        """
        Integration runtime reference type.
        :param pulumi.Input[str] reference_name: Reference integration runtime name.
        :param pulumi.Input[str] type: Type of integration runtime.
        :param pulumi.Input[Mapping[str, pulumi.Input[Mapping[str, Any]]]] parameters: Arguments for integration runtime.
        """
        pulumi.set(__self__, "reference_name", reference_name)
        pulumi.set(__self__, "type", type)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)

    @property
    @pulumi.getter(name="referenceName")
    def reference_name(self) -> pulumi.Input[str]:
        """
        Reference integration runtime name.
        """
        return pulumi.get(self, "reference_name")

    @reference_name.setter
    def reference_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "reference_name", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Type of integration runtime.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[Mapping[str, Any]]]]]:
        """
        Arguments for integration runtime.
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[Mapping[str, Any]]]]]):
        pulumi.set(self, "parameters", value)


@pulumi.input_type
class LinkedServiceArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 annotations: Optional[pulumi.Input[List[pulumi.Input[Mapping[str, Any]]]]] = None,
                 connect_via: Optional[pulumi.Input['IntegrationRuntimeReferenceArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input['ParameterSpecificationArgs']]]] = None):
        """
        The Azure Data Factory nested object which contains the information and credential which can be used to connect with related store or compute resource.
        :param pulumi.Input[str] type: Type of linked service.
        :param pulumi.Input[List[pulumi.Input[Mapping[str, Any]]]] annotations: List of tags that can be used for describing the linked service.
        :param pulumi.Input['IntegrationRuntimeReferenceArgs'] connect_via: The integration runtime reference.
        :param pulumi.Input[str] description: Linked service description.
        :param pulumi.Input[Mapping[str, pulumi.Input['ParameterSpecificationArgs']]] parameters: Parameters for linked service.
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
    def type(self) -> pulumi.Input[str]:
        """
        Type of linked service.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def annotations(self) -> Optional[pulumi.Input[List[pulumi.Input[Mapping[str, Any]]]]]:
        """
        List of tags that can be used for describing the linked service.
        """
        return pulumi.get(self, "annotations")

    @annotations.setter
    def annotations(self, value: Optional[pulumi.Input[List[pulumi.Input[Mapping[str, Any]]]]]):
        pulumi.set(self, "annotations", value)

    @property
    @pulumi.getter(name="connectVia")
    def connect_via(self) -> Optional[pulumi.Input['IntegrationRuntimeReferenceArgs']]:
        """
        The integration runtime reference.
        """
        return pulumi.get(self, "connect_via")

    @connect_via.setter
    def connect_via(self, value: Optional[pulumi.Input['IntegrationRuntimeReferenceArgs']]):
        pulumi.set(self, "connect_via", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Linked service description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input['ParameterSpecificationArgs']]]]:
        """
        Parameters for linked service.
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input['ParameterSpecificationArgs']]]]):
        pulumi.set(self, "parameters", value)


@pulumi.input_type
class LinkedServiceReferenceArgs:
    def __init__(__self__, *,
                 reference_name: pulumi.Input[str],
                 type: pulumi.Input[str],
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[Mapping[str, Any]]]]] = None):
        """
        Linked service reference type.
        :param pulumi.Input[str] reference_name: Reference LinkedService name.
        :param pulumi.Input[str] type: Linked service reference type.
        :param pulumi.Input[Mapping[str, pulumi.Input[Mapping[str, Any]]]] parameters: Arguments for LinkedService.
        """
        pulumi.set(__self__, "reference_name", reference_name)
        pulumi.set(__self__, "type", type)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)

    @property
    @pulumi.getter(name="referenceName")
    def reference_name(self) -> pulumi.Input[str]:
        """
        Reference LinkedService name.
        """
        return pulumi.get(self, "reference_name")

    @reference_name.setter
    def reference_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "reference_name", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Linked service reference type.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[Mapping[str, Any]]]]]:
        """
        Arguments for LinkedService.
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[Mapping[str, Any]]]]]):
        pulumi.set(self, "parameters", value)


@pulumi.input_type
class ManagedPrivateEndpointArgs:
    def __init__(__self__, *,
                 fqdns: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 private_link_resource_id: Optional[pulumi.Input[str]] = None):
        """
        Properties of a managed private endpoint
        :param pulumi.Input[List[pulumi.Input[str]]] fqdns: Fully qualified domain names
        :param pulumi.Input[str] group_id: The groupId to which the managed private endpoint is created
        :param pulumi.Input[str] private_link_resource_id: The ARM resource ID of the resource to which the managed private endpoint is created
        """
        if fqdns is not None:
            pulumi.set(__self__, "fqdns", fqdns)
        if group_id is not None:
            pulumi.set(__self__, "group_id", group_id)
        if private_link_resource_id is not None:
            pulumi.set(__self__, "private_link_resource_id", private_link_resource_id)

    @property
    @pulumi.getter
    def fqdns(self) -> Optional[pulumi.Input[List[pulumi.Input[str]]]]:
        """
        Fully qualified domain names
        """
        return pulumi.get(self, "fqdns")

    @fqdns.setter
    def fqdns(self, value: Optional[pulumi.Input[List[pulumi.Input[str]]]]):
        pulumi.set(self, "fqdns", value)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The groupId to which the managed private endpoint is created
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter(name="privateLinkResourceId")
    def private_link_resource_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ARM resource ID of the resource to which the managed private endpoint is created
        """
        return pulumi.get(self, "private_link_resource_id")

    @private_link_resource_id.setter
    def private_link_resource_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_link_resource_id", value)


@pulumi.input_type
class ParameterSpecificationArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 default_value: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        Definition of a single parameter for an entity.
        :param pulumi.Input[str] type: Parameter type.
        :param pulumi.Input[Mapping[str, Any]] default_value: Default value of parameter.
        """
        pulumi.set(__self__, "type", type)
        if default_value is not None:
            pulumi.set(__self__, "default_value", default_value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Parameter type.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="defaultValue")
    def default_value(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Default value of parameter.
        """
        return pulumi.get(self, "default_value")

    @default_value.setter
    def default_value(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "default_value", value)


@pulumi.input_type
class PipelineFolderArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The folder that this Pipeline is in. If not specified, Pipeline will appear at the root level.
        :param pulumi.Input[str] name: The name of the folder that this Pipeline is in.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the folder that this Pipeline is in.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class TriggerArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 annotations: Optional[pulumi.Input[List[pulumi.Input[Mapping[str, Any]]]]] = None,
                 description: Optional[pulumi.Input[str]] = None):
        """
        Azure data factory nested object which contains information about creating pipeline run
        :param pulumi.Input[str] type: Trigger type.
        :param pulumi.Input[List[pulumi.Input[Mapping[str, Any]]]] annotations: List of tags that can be used for describing the trigger.
        :param pulumi.Input[str] description: Trigger description.
        """
        pulumi.set(__self__, "type", type)
        if annotations is not None:
            pulumi.set(__self__, "annotations", annotations)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Trigger type.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def annotations(self) -> Optional[pulumi.Input[List[pulumi.Input[Mapping[str, Any]]]]]:
        """
        List of tags that can be used for describing the trigger.
        """
        return pulumi.get(self, "annotations")

    @annotations.setter
    def annotations(self, value: Optional[pulumi.Input[List[pulumi.Input[Mapping[str, Any]]]]]):
        pulumi.set(self, "annotations", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Trigger description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


@pulumi.input_type
class UserPropertyArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 value: pulumi.Input[Mapping[str, Any]]):
        """
        User property.
        :param pulumi.Input[str] name: User property name.
        :param pulumi.Input[Mapping[str, Any]] value: User property value. Type: string (or Expression with resultType string).
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        User property name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[Mapping[str, Any]]:
        """
        User property value. Type: string (or Expression with resultType string).
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[Mapping[str, Any]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class VariableSpecificationArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 default_value: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        Definition of a single variable for a Pipeline.
        :param pulumi.Input[str] type: Variable type.
        :param pulumi.Input[Mapping[str, Any]] default_value: Default value of variable.
        """
        pulumi.set(__self__, "type", type)
        if default_value is not None:
            pulumi.set(__self__, "default_value", default_value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Variable type.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="defaultValue")
    def default_value(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Default value of variable.
        """
        return pulumi.get(self, "default_value")

    @default_value.setter
    def default_value(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "default_value", value)

