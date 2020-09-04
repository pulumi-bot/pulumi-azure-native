# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'AuthenticationArgs',
    'IdentityArgs',
    'PrePostStepArgs',
    'ServiceUnitArtifactsArgs',
    'StepArgs',
    'StepPropertiesArgs',
]

@pulumi.input_type
class AuthenticationArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str]):
        """
        Defines the authentication method and properties to access the artifacts.
        :param pulumi.Input[str] type: The authentication type
        """
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The authentication type
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class IdentityArgs:
    def __init__(__self__, *,
                 identity_ids: pulumi.Input[List[pulumi.Input[str]]],
                 type: pulumi.Input[str]):
        """
        Identity for the resource.
        :param pulumi.Input[List[pulumi.Input[str]]] identity_ids: The list of identities.
        :param pulumi.Input[str] type: The identity type.
        """
        pulumi.set(__self__, "identity_ids", identity_ids)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="identityIds")
    def identity_ids(self) -> pulumi.Input[List[pulumi.Input[str]]]:
        """
        The list of identities.
        """
        return pulumi.get(self, "identity_ids")

    @identity_ids.setter
    def identity_ids(self, value: pulumi.Input[List[pulumi.Input[str]]]):
        pulumi.set(self, "identity_ids", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The identity type.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class PrePostStepArgs:
    def __init__(__self__, *,
                 step_id: pulumi.Input[str]):
        """
        The properties that define a step.
        :param pulumi.Input[str] step_id: The resource Id of the step to be run.
        """
        pulumi.set(__self__, "step_id", step_id)

    @property
    @pulumi.getter(name="stepId")
    def step_id(self) -> pulumi.Input[str]:
        """
        The resource Id of the step to be run.
        """
        return pulumi.get(self, "step_id")

    @step_id.setter
    def step_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "step_id", value)


@pulumi.input_type
class ServiceUnitArtifactsArgs:
    def __init__(__self__, *,
                 parameters_artifact_source_relative_path: Optional[pulumi.Input[str]] = None,
                 parameters_uri: Optional[pulumi.Input[str]] = None,
                 template_artifact_source_relative_path: Optional[pulumi.Input[str]] = None,
                 template_uri: Optional[pulumi.Input[str]] = None):
        """
        Defines the artifacts of a service unit.
        :param pulumi.Input[str] parameters_artifact_source_relative_path: The path to the ARM parameters file relative to the artifact source.
        :param pulumi.Input[str] parameters_uri: The full URI of the ARM parameters file with the SAS token.
        :param pulumi.Input[str] template_artifact_source_relative_path: The path to the ARM template file relative to the artifact source.
        :param pulumi.Input[str] template_uri: The full URI of the ARM template file with the SAS token.
        """
        if parameters_artifact_source_relative_path is not None:
            pulumi.set(__self__, "parameters_artifact_source_relative_path", parameters_artifact_source_relative_path)
        if parameters_uri is not None:
            pulumi.set(__self__, "parameters_uri", parameters_uri)
        if template_artifact_source_relative_path is not None:
            pulumi.set(__self__, "template_artifact_source_relative_path", template_artifact_source_relative_path)
        if template_uri is not None:
            pulumi.set(__self__, "template_uri", template_uri)

    @property
    @pulumi.getter(name="parametersArtifactSourceRelativePath")
    def parameters_artifact_source_relative_path(self) -> Optional[pulumi.Input[str]]:
        """
        The path to the ARM parameters file relative to the artifact source.
        """
        return pulumi.get(self, "parameters_artifact_source_relative_path")

    @parameters_artifact_source_relative_path.setter
    def parameters_artifact_source_relative_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parameters_artifact_source_relative_path", value)

    @property
    @pulumi.getter(name="parametersUri")
    def parameters_uri(self) -> Optional[pulumi.Input[str]]:
        """
        The full URI of the ARM parameters file with the SAS token.
        """
        return pulumi.get(self, "parameters_uri")

    @parameters_uri.setter
    def parameters_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parameters_uri", value)

    @property
    @pulumi.getter(name="templateArtifactSourceRelativePath")
    def template_artifact_source_relative_path(self) -> Optional[pulumi.Input[str]]:
        """
        The path to the ARM template file relative to the artifact source.
        """
        return pulumi.get(self, "template_artifact_source_relative_path")

    @template_artifact_source_relative_path.setter
    def template_artifact_source_relative_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "template_artifact_source_relative_path", value)

    @property
    @pulumi.getter(name="templateUri")
    def template_uri(self) -> Optional[pulumi.Input[str]]:
        """
        The full URI of the ARM template file with the SAS token.
        """
        return pulumi.get(self, "template_uri")

    @template_uri.setter
    def template_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "template_uri", value)


@pulumi.input_type
class StepArgs:
    def __init__(__self__, *,
                 deployment_target_id: pulumi.Input[str],
                 name: pulumi.Input[str],
                 depends_on_step_groups: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 post_deployment_steps: Optional[pulumi.Input[List[pulumi.Input['PrePostStepArgs']]]] = None,
                 pre_deployment_steps: Optional[pulumi.Input[List[pulumi.Input['PrePostStepArgs']]]] = None):
        """
        The properties that define an Azure Deployment Manager step.
        :param pulumi.Input[str] deployment_target_id: The resource Id of service unit to be deployed. The service unit should be from the service topology referenced in targetServiceTopologyId
        :param pulumi.Input[str] name: The name of the step group.
        :param pulumi.Input[List[pulumi.Input[str]]] depends_on_step_groups: The list of step group names on which this step group depends on.
        :param pulumi.Input[List[pulumi.Input['PrePostStepArgs']]] post_deployment_steps: The list of steps to be run after deploying the target.
        :param pulumi.Input[List[pulumi.Input['PrePostStepArgs']]] pre_deployment_steps: The list of steps to be run before deploying the target.
        """
        pulumi.set(__self__, "deployment_target_id", deployment_target_id)
        pulumi.set(__self__, "name", name)
        if depends_on_step_groups is not None:
            pulumi.set(__self__, "depends_on_step_groups", depends_on_step_groups)
        if post_deployment_steps is not None:
            pulumi.set(__self__, "post_deployment_steps", post_deployment_steps)
        if pre_deployment_steps is not None:
            pulumi.set(__self__, "pre_deployment_steps", pre_deployment_steps)

    @property
    @pulumi.getter(name="deploymentTargetId")
    def deployment_target_id(self) -> pulumi.Input[str]:
        """
        The resource Id of service unit to be deployed. The service unit should be from the service topology referenced in targetServiceTopologyId
        """
        return pulumi.get(self, "deployment_target_id")

    @deployment_target_id.setter
    def deployment_target_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "deployment_target_id", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The name of the step group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="dependsOnStepGroups")
    def depends_on_step_groups(self) -> Optional[pulumi.Input[List[pulumi.Input[str]]]]:
        """
        The list of step group names on which this step group depends on.
        """
        return pulumi.get(self, "depends_on_step_groups")

    @depends_on_step_groups.setter
    def depends_on_step_groups(self, value: Optional[pulumi.Input[List[pulumi.Input[str]]]]):
        pulumi.set(self, "depends_on_step_groups", value)

    @property
    @pulumi.getter(name="postDeploymentSteps")
    def post_deployment_steps(self) -> Optional[pulumi.Input[List[pulumi.Input['PrePostStepArgs']]]]:
        """
        The list of steps to be run after deploying the target.
        """
        return pulumi.get(self, "post_deployment_steps")

    @post_deployment_steps.setter
    def post_deployment_steps(self, value: Optional[pulumi.Input[List[pulumi.Input['PrePostStepArgs']]]]):
        pulumi.set(self, "post_deployment_steps", value)

    @property
    @pulumi.getter(name="preDeploymentSteps")
    def pre_deployment_steps(self) -> Optional[pulumi.Input[List[pulumi.Input['PrePostStepArgs']]]]:
        """
        The list of steps to be run before deploying the target.
        """
        return pulumi.get(self, "pre_deployment_steps")

    @pre_deployment_steps.setter
    def pre_deployment_steps(self, value: Optional[pulumi.Input[List[pulumi.Input['PrePostStepArgs']]]]):
        pulumi.set(self, "pre_deployment_steps", value)


@pulumi.input_type
class StepPropertiesArgs:
    def __init__(__self__, *,
                 step_type: pulumi.Input[str]):
        """
        The properties of a step resource.
        :param pulumi.Input[str] step_type: The type of step.
        """
        pulumi.set(__self__, "step_type", step_type)

    @property
    @pulumi.getter(name="stepType")
    def step_type(self) -> pulumi.Input[str]:
        """
        The type of step.
        """
        return pulumi.get(self, "step_type")

    @step_type.setter
    def step_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "step_type", value)


