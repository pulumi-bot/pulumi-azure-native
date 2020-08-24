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
    'AliasPathResponse',
    'AliasPatternResponse',
    'AliasResponse',
    'BasicDependencyResponse',
    'DebugSettingResponse',
    'DependencyResponse',
    'DeploymentPropertiesExtendedResponse',
    'ErrorAdditionalInfoResponse',
    'ErrorResponseResponse',
    'OnErrorDeploymentExtendedResponse',
    'ParametersLinkResponse',
    'ProviderResourceTypeResponse',
    'ProviderResponse',
    'ResourceReferenceResponse',
    'TemplateLinkResponse',
]

@pulumi.output_type
class AliasPathResponse(dict):
    """
    The type of the paths for alias.
    """
    def __init__(__self__, *,
                 api_versions: Optional[List[str]] = None,
                 path: Optional[str] = None,
                 pattern: Optional['outputs.AliasPatternResponse'] = None):
        """
        The type of the paths for alias.
        :param List[str] api_versions: The API versions.
        :param str path: The path of an alias.
        :param 'AliasPatternResponseArgs' pattern: The pattern for an alias path.
        """
        if api_versions is not None:
            pulumi.set(__self__, "api_versions", api_versions)
        if path is not None:
            pulumi.set(__self__, "path", path)
        if pattern is not None:
            pulumi.set(__self__, "pattern", pattern)

    @property
    @pulumi.getter(name="apiVersions")
    def api_versions(self) -> Optional[List[str]]:
        """
        The API versions.
        """
        return pulumi.get(self, "api_versions")

    @property
    @pulumi.getter
    def path(self) -> Optional[str]:
        """
        The path of an alias.
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter
    def pattern(self) -> Optional['outputs.AliasPatternResponse']:
        """
        The pattern for an alias path.
        """
        return pulumi.get(self, "pattern")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class AliasPatternResponse(dict):
    """
    The type of the pattern for an alias path.
    """
    def __init__(__self__, *,
                 phrase: Optional[str] = None,
                 type: Optional[str] = None,
                 variable: Optional[str] = None):
        """
        The type of the pattern for an alias path.
        :param str phrase: The alias pattern phrase.
        :param str type: The type of alias pattern
        :param str variable: The alias pattern variable.
        """
        if phrase is not None:
            pulumi.set(__self__, "phrase", phrase)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if variable is not None:
            pulumi.set(__self__, "variable", variable)

    @property
    @pulumi.getter
    def phrase(self) -> Optional[str]:
        """
        The alias pattern phrase.
        """
        return pulumi.get(self, "phrase")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        The type of alias pattern
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def variable(self) -> Optional[str]:
        """
        The alias pattern variable.
        """
        return pulumi.get(self, "variable")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class AliasResponse(dict):
    """
    The alias type. 
    """
    def __init__(__self__, *,
                 default_path: Optional[str] = None,
                 default_pattern: Optional['outputs.AliasPatternResponse'] = None,
                 name: Optional[str] = None,
                 paths: Optional[List['outputs.AliasPathResponse']] = None,
                 type: Optional[str] = None):
        """
        The alias type. 
        :param str default_path: The default path for an alias.
        :param 'AliasPatternResponseArgs' default_pattern: The default pattern for an alias.
        :param str name: The alias name.
        :param List['AliasPathResponseArgs'] paths: The paths for an alias.
        :param str type: The type of the alias.
        """
        if default_path is not None:
            pulumi.set(__self__, "default_path", default_path)
        if default_pattern is not None:
            pulumi.set(__self__, "default_pattern", default_pattern)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if paths is not None:
            pulumi.set(__self__, "paths", paths)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="defaultPath")
    def default_path(self) -> Optional[str]:
        """
        The default path for an alias.
        """
        return pulumi.get(self, "default_path")

    @property
    @pulumi.getter(name="defaultPattern")
    def default_pattern(self) -> Optional['outputs.AliasPatternResponse']:
        """
        The default pattern for an alias.
        """
        return pulumi.get(self, "default_pattern")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The alias name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def paths(self) -> Optional[List['outputs.AliasPathResponse']]:
        """
        The paths for an alias.
        """
        return pulumi.get(self, "paths")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        The type of the alias.
        """
        return pulumi.get(self, "type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class BasicDependencyResponse(dict):
    """
    Deployment dependency information.
    """
    def __init__(__self__, *,
                 id: Optional[str] = None,
                 resource_name: Optional[str] = None,
                 resource_type: Optional[str] = None):
        """
        Deployment dependency information.
        :param str id: The ID of the dependency.
        :param str resource_name: The dependency resource name.
        :param str resource_type: The dependency resource type.
        """
        if id is not None:
            pulumi.set(__self__, "id", id)
        if resource_name is not None:
            pulumi.set(__self__, "resource_name", resource_name)
        if resource_type is not None:
            pulumi.set(__self__, "resource_type", resource_type)

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        The ID of the dependency.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="resourceName")
    def resource_name(self) -> Optional[str]:
        """
        The dependency resource name.
        """
        return pulumi.get(self, "resource_name")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> Optional[str]:
        """
        The dependency resource type.
        """
        return pulumi.get(self, "resource_type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class DebugSettingResponse(dict):
    """
    The debug setting.
    """
    def __init__(__self__, *,
                 detail_level: Optional[str] = None):
        """
        The debug setting.
        :param str detail_level: Specifies the type of information to log for debugging. The permitted values are none, requestContent, responseContent, or both requestContent and responseContent separated by a comma. The default is none. When setting this value, carefully consider the type of information you are passing in during deployment. By logging information about the request or response, you could potentially expose sensitive data that is retrieved through the deployment operations.
        """
        if detail_level is not None:
            pulumi.set(__self__, "detail_level", detail_level)

    @property
    @pulumi.getter(name="detailLevel")
    def detail_level(self) -> Optional[str]:
        """
        Specifies the type of information to log for debugging. The permitted values are none, requestContent, responseContent, or both requestContent and responseContent separated by a comma. The default is none. When setting this value, carefully consider the type of information you are passing in during deployment. By logging information about the request or response, you could potentially expose sensitive data that is retrieved through the deployment operations.
        """
        return pulumi.get(self, "detail_level")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class DependencyResponse(dict):
    """
    Deployment dependency information.
    """
    def __init__(__self__, *,
                 depends_on: Optional[List['outputs.BasicDependencyResponse']] = None,
                 id: Optional[str] = None,
                 resource_name: Optional[str] = None,
                 resource_type: Optional[str] = None):
        """
        Deployment dependency information.
        :param List['BasicDependencyResponseArgs'] depends_on: The list of dependencies.
        :param str id: The ID of the dependency.
        :param str resource_name: The dependency resource name.
        :param str resource_type: The dependency resource type.
        """
        if depends_on is not None:
            pulumi.set(__self__, "depends_on", depends_on)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if resource_name is not None:
            pulumi.set(__self__, "resource_name", resource_name)
        if resource_type is not None:
            pulumi.set(__self__, "resource_type", resource_type)

    @property
    @pulumi.getter(name="dependsOn")
    def depends_on(self) -> Optional[List['outputs.BasicDependencyResponse']]:
        """
        The list of dependencies.
        """
        return pulumi.get(self, "depends_on")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        The ID of the dependency.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="resourceName")
    def resource_name(self) -> Optional[str]:
        """
        The dependency resource name.
        """
        return pulumi.get(self, "resource_name")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> Optional[str]:
        """
        The dependency resource type.
        """
        return pulumi.get(self, "resource_type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class DeploymentPropertiesExtendedResponse(dict):
    """
    Deployment properties with additional details.
    """
    def __init__(__self__, *,
                 correlation_id: str,
                 debug_setting: 'outputs.DebugSettingResponse',
                 dependencies: List['outputs.DependencyResponse'],
                 duration: str,
                 error: 'outputs.ErrorResponseResponse',
                 mode: str,
                 on_error_deployment: 'outputs.OnErrorDeploymentExtendedResponse',
                 output_resources: List['outputs.ResourceReferenceResponse'],
                 outputs: Mapping[str, Any],
                 parameters: Mapping[str, Any],
                 parameters_link: 'outputs.ParametersLinkResponse',
                 providers: List['outputs.ProviderResponse'],
                 provisioning_state: str,
                 template_hash: str,
                 template_link: 'outputs.TemplateLinkResponse',
                 timestamp: str,
                 validated_resources: List['outputs.ResourceReferenceResponse']):
        """
        Deployment properties with additional details.
        :param str correlation_id: The correlation ID of the deployment.
        :param 'DebugSettingResponseArgs' debug_setting: The debug setting of the deployment.
        :param List['DependencyResponseArgs'] dependencies: The list of deployment dependencies.
        :param str duration: The duration of the template deployment.
        :param 'ErrorResponseResponseArgs' error: The deployment error.
        :param str mode: The deployment mode. Possible values are Incremental and Complete.
        :param 'OnErrorDeploymentExtendedResponseArgs' on_error_deployment: The deployment on error behavior.
        :param List['ResourceReferenceResponseArgs'] output_resources: Array of provisioned resources.
        :param Mapping[str, Any] outputs: Key/value pairs that represent deployment output.
        :param Mapping[str, Any] parameters: Deployment parameters. 
        :param 'ParametersLinkResponseArgs' parameters_link: The URI referencing the parameters. 
        :param List['ProviderResponseArgs'] providers: The list of resource providers needed for the deployment.
        :param str provisioning_state: The state of the provisioning.
        :param str template_hash: The hash produced for the template.
        :param 'TemplateLinkResponseArgs' template_link: The URI referencing the template.
        :param str timestamp: The timestamp of the template deployment.
        :param List['ResourceReferenceResponseArgs'] validated_resources: Array of validated resources.
        """
        pulumi.set(__self__, "correlation_id", correlation_id)
        pulumi.set(__self__, "debug_setting", debug_setting)
        pulumi.set(__self__, "dependencies", dependencies)
        pulumi.set(__self__, "duration", duration)
        pulumi.set(__self__, "error", error)
        pulumi.set(__self__, "mode", mode)
        pulumi.set(__self__, "on_error_deployment", on_error_deployment)
        pulumi.set(__self__, "output_resources", output_resources)
        pulumi.set(__self__, "outputs", outputs)
        pulumi.set(__self__, "parameters", parameters)
        pulumi.set(__self__, "parameters_link", parameters_link)
        pulumi.set(__self__, "providers", providers)
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        pulumi.set(__self__, "template_hash", template_hash)
        pulumi.set(__self__, "template_link", template_link)
        pulumi.set(__self__, "timestamp", timestamp)
        pulumi.set(__self__, "validated_resources", validated_resources)

    @property
    @pulumi.getter(name="correlationId")
    def correlation_id(self) -> str:
        """
        The correlation ID of the deployment.
        """
        return pulumi.get(self, "correlation_id")

    @property
    @pulumi.getter(name="debugSetting")
    def debug_setting(self) -> 'outputs.DebugSettingResponse':
        """
        The debug setting of the deployment.
        """
        return pulumi.get(self, "debug_setting")

    @property
    @pulumi.getter
    def dependencies(self) -> List['outputs.DependencyResponse']:
        """
        The list of deployment dependencies.
        """
        return pulumi.get(self, "dependencies")

    @property
    @pulumi.getter
    def duration(self) -> str:
        """
        The duration of the template deployment.
        """
        return pulumi.get(self, "duration")

    @property
    @pulumi.getter
    def error(self) -> 'outputs.ErrorResponseResponse':
        """
        The deployment error.
        """
        return pulumi.get(self, "error")

    @property
    @pulumi.getter
    def mode(self) -> str:
        """
        The deployment mode. Possible values are Incremental and Complete.
        """
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter(name="onErrorDeployment")
    def on_error_deployment(self) -> 'outputs.OnErrorDeploymentExtendedResponse':
        """
        The deployment on error behavior.
        """
        return pulumi.get(self, "on_error_deployment")

    @property
    @pulumi.getter(name="outputResources")
    def output_resources(self) -> List['outputs.ResourceReferenceResponse']:
        """
        Array of provisioned resources.
        """
        return pulumi.get(self, "output_resources")

    @property
    @pulumi.getter
    def outputs(self) -> Mapping[str, Any]:
        """
        Key/value pairs that represent deployment output.
        """
        return pulumi.get(self, "outputs")

    @property
    @pulumi.getter
    def parameters(self) -> Mapping[str, Any]:
        """
        Deployment parameters. 
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter(name="parametersLink")
    def parameters_link(self) -> 'outputs.ParametersLinkResponse':
        """
        The URI referencing the parameters. 
        """
        return pulumi.get(self, "parameters_link")

    @property
    @pulumi.getter
    def providers(self) -> List['outputs.ProviderResponse']:
        """
        The list of resource providers needed for the deployment.
        """
        return pulumi.get(self, "providers")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        The state of the provisioning.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="templateHash")
    def template_hash(self) -> str:
        """
        The hash produced for the template.
        """
        return pulumi.get(self, "template_hash")

    @property
    @pulumi.getter(name="templateLink")
    def template_link(self) -> 'outputs.TemplateLinkResponse':
        """
        The URI referencing the template.
        """
        return pulumi.get(self, "template_link")

    @property
    @pulumi.getter
    def timestamp(self) -> str:
        """
        The timestamp of the template deployment.
        """
        return pulumi.get(self, "timestamp")

    @property
    @pulumi.getter(name="validatedResources")
    def validated_resources(self) -> List['outputs.ResourceReferenceResponse']:
        """
        Array of validated resources.
        """
        return pulumi.get(self, "validated_resources")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ErrorAdditionalInfoResponse(dict):
    """
    The resource management error additional info.
    """
    def __init__(__self__, *,
                 info: Mapping[str, Any],
                 type: str):
        """
        The resource management error additional info.
        :param Mapping[str, Any] info: The additional info.
        :param str type: The additional info type.
        """
        pulumi.set(__self__, "info", info)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def info(self) -> Mapping[str, Any]:
        """
        The additional info.
        """
        return pulumi.get(self, "info")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The additional info type.
        """
        return pulumi.get(self, "type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ErrorResponseResponse(dict):
    """
    The resource management error response.
    """
    def __init__(__self__, *,
                 additional_info: List['outputs.ErrorAdditionalInfoResponse'],
                 code: str,
                 details: List['outputs.ErrorResponseResponse'],
                 message: str,
                 target: str):
        """
        The resource management error response.
        :param List['ErrorAdditionalInfoResponseArgs'] additional_info: The error additional info.
        :param str code: The error code.
        :param List['ErrorResponseResponseArgs'] details: The error details.
        :param str message: The error message.
        :param str target: The error target.
        """
        pulumi.set(__self__, "additional_info", additional_info)
        pulumi.set(__self__, "code", code)
        pulumi.set(__self__, "details", details)
        pulumi.set(__self__, "message", message)
        pulumi.set(__self__, "target", target)

    @property
    @pulumi.getter(name="additionalInfo")
    def additional_info(self) -> List['outputs.ErrorAdditionalInfoResponse']:
        """
        The error additional info.
        """
        return pulumi.get(self, "additional_info")

    @property
    @pulumi.getter
    def code(self) -> str:
        """
        The error code.
        """
        return pulumi.get(self, "code")

    @property
    @pulumi.getter
    def details(self) -> List['outputs.ErrorResponseResponse']:
        """
        The error details.
        """
        return pulumi.get(self, "details")

    @property
    @pulumi.getter
    def message(self) -> str:
        """
        The error message.
        """
        return pulumi.get(self, "message")

    @property
    @pulumi.getter
    def target(self) -> str:
        """
        The error target.
        """
        return pulumi.get(self, "target")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class OnErrorDeploymentExtendedResponse(dict):
    """
    Deployment on error behavior with additional details.
    """
    def __init__(__self__, *,
                 provisioning_state: str,
                 deployment_name: Optional[str] = None,
                 type: Optional[str] = None):
        """
        Deployment on error behavior with additional details.
        :param str provisioning_state: The state of the provisioning for the on error deployment.
        :param str deployment_name: The deployment to be used on error case.
        :param str type: The deployment on error behavior type. Possible values are LastSuccessful and SpecificDeployment.
        """
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if deployment_name is not None:
            pulumi.set(__self__, "deployment_name", deployment_name)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        The state of the provisioning for the on error deployment.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="deploymentName")
    def deployment_name(self) -> Optional[str]:
        """
        The deployment to be used on error case.
        """
        return pulumi.get(self, "deployment_name")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        The deployment on error behavior type. Possible values are LastSuccessful and SpecificDeployment.
        """
        return pulumi.get(self, "type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ParametersLinkResponse(dict):
    """
    Entity representing the reference to the deployment parameters.
    """
    def __init__(__self__, *,
                 uri: str,
                 content_version: Optional[str] = None):
        """
        Entity representing the reference to the deployment parameters.
        :param str uri: The URI of the parameters file.
        :param str content_version: If included, must match the ContentVersion in the template.
        """
        pulumi.set(__self__, "uri", uri)
        if content_version is not None:
            pulumi.set(__self__, "content_version", content_version)

    @property
    @pulumi.getter
    def uri(self) -> str:
        """
        The URI of the parameters file.
        """
        return pulumi.get(self, "uri")

    @property
    @pulumi.getter(name="contentVersion")
    def content_version(self) -> Optional[str]:
        """
        If included, must match the ContentVersion in the template.
        """
        return pulumi.get(self, "content_version")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ProviderResourceTypeResponse(dict):
    """
    Resource type managed by the resource provider.
    """
    def __init__(__self__, *,
                 aliases: Optional[List['outputs.AliasResponse']] = None,
                 api_versions: Optional[List[str]] = None,
                 capabilities: Optional[str] = None,
                 locations: Optional[List[str]] = None,
                 properties: Optional[Mapping[str, str]] = None,
                 resource_type: Optional[str] = None):
        """
        Resource type managed by the resource provider.
        :param List['AliasResponseArgs'] aliases: The aliases that are supported by this resource type.
        :param List[str] api_versions: The API version.
        :param str capabilities: The additional capabilities offered by this resource type.
        :param List[str] locations: The collection of locations where this resource type can be created.
        :param Mapping[str, str] properties: The properties.
        :param str resource_type: The resource type.
        """
        if aliases is not None:
            pulumi.set(__self__, "aliases", aliases)
        if api_versions is not None:
            pulumi.set(__self__, "api_versions", api_versions)
        if capabilities is not None:
            pulumi.set(__self__, "capabilities", capabilities)
        if locations is not None:
            pulumi.set(__self__, "locations", locations)
        if properties is not None:
            pulumi.set(__self__, "properties", properties)
        if resource_type is not None:
            pulumi.set(__self__, "resource_type", resource_type)

    @property
    @pulumi.getter
    def aliases(self) -> Optional[List['outputs.AliasResponse']]:
        """
        The aliases that are supported by this resource type.
        """
        return pulumi.get(self, "aliases")

    @property
    @pulumi.getter(name="apiVersions")
    def api_versions(self) -> Optional[List[str]]:
        """
        The API version.
        """
        return pulumi.get(self, "api_versions")

    @property
    @pulumi.getter
    def capabilities(self) -> Optional[str]:
        """
        The additional capabilities offered by this resource type.
        """
        return pulumi.get(self, "capabilities")

    @property
    @pulumi.getter
    def locations(self) -> Optional[List[str]]:
        """
        The collection of locations where this resource type can be created.
        """
        return pulumi.get(self, "locations")

    @property
    @pulumi.getter
    def properties(self) -> Optional[Mapping[str, str]]:
        """
        The properties.
        """
        return pulumi.get(self, "properties")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> Optional[str]:
        """
        The resource type.
        """
        return pulumi.get(self, "resource_type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ProviderResponse(dict):
    """
    Resource provider information.
    """
    def __init__(__self__, *,
                 id: str,
                 registration_policy: str,
                 registration_state: str,
                 resource_types: List['outputs.ProviderResourceTypeResponse'],
                 namespace: Optional[str] = None):
        """
        Resource provider information.
        :param str id: The provider ID.
        :param str registration_policy: The registration policy of the resource provider.
        :param str registration_state: The registration state of the resource provider.
        :param List['ProviderResourceTypeResponseArgs'] resource_types: The collection of provider resource types.
        :param str namespace: The namespace of the resource provider.
        """
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "registration_policy", registration_policy)
        pulumi.set(__self__, "registration_state", registration_state)
        pulumi.set(__self__, "resource_types", resource_types)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="registrationPolicy")
    def registration_policy(self) -> str:
        """
        The registration policy of the resource provider.
        """
        return pulumi.get(self, "registration_policy")

    @property
    @pulumi.getter(name="registrationState")
    def registration_state(self) -> str:
        """
        The registration state of the resource provider.
        """
        return pulumi.get(self, "registration_state")

    @property
    @pulumi.getter(name="resourceTypes")
    def resource_types(self) -> List['outputs.ProviderResourceTypeResponse']:
        """
        The collection of provider resource types.
        """
        return pulumi.get(self, "resource_types")

    @property
    @pulumi.getter
    def namespace(self) -> Optional[str]:
        """
        The namespace of the resource provider.
        """
        return pulumi.get(self, "namespace")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ResourceReferenceResponse(dict):
    """
    The resource Id model.
    """
    def __init__(__self__, *,
                 id: str):
        """
        The resource Id model.
        :param str id: The fully qualified resource Id.
        """
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The fully qualified resource Id.
        """
        return pulumi.get(self, "id")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class TemplateLinkResponse(dict):
    """
    Entity representing the reference to the template.
    """
    def __init__(__self__, *,
                 uri: str,
                 content_version: Optional[str] = None):
        """
        Entity representing the reference to the template.
        :param str uri: The URI of the template to deploy.
        :param str content_version: If included, must match the ContentVersion in the template.
        """
        pulumi.set(__self__, "uri", uri)
        if content_version is not None:
            pulumi.set(__self__, "content_version", content_version)

    @property
    @pulumi.getter
    def uri(self) -> str:
        """
        The URI of the template to deploy.
        """
        return pulumi.get(self, "uri")

    @property
    @pulumi.getter(name="contentVersion")
    def content_version(self) -> Optional[str]:
        """
        If included, must match the ContentVersion in the template.
        """
        return pulumi.get(self, "content_version")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


