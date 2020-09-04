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
    'GetPublishedBlueprintResult',
    'AwaitableGetPublishedBlueprintResult',
    'get_published_blueprint',
]

@pulumi.output_type
class GetPublishedBlueprintResult:
    """
    Represents a published Blueprint.
    """
    def __init__(__self__, blueprint_name=None, change_notes=None, description=None, display_name=None, name=None, parameters=None, resource_groups=None, status=None, target_scope=None, type=None):
        if blueprint_name and not isinstance(blueprint_name, str):
            raise TypeError("Expected argument 'blueprint_name' to be a str")
        pulumi.set(__self__, "blueprint_name", blueprint_name)
        if change_notes and not isinstance(change_notes, str):
            raise TypeError("Expected argument 'change_notes' to be a str")
        pulumi.set(__self__, "change_notes", change_notes)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if parameters and not isinstance(parameters, dict):
            raise TypeError("Expected argument 'parameters' to be a dict")
        pulumi.set(__self__, "parameters", parameters)
        if resource_groups and not isinstance(resource_groups, dict):
            raise TypeError("Expected argument 'resource_groups' to be a dict")
        pulumi.set(__self__, "resource_groups", resource_groups)
        if status and not isinstance(status, dict):
            raise TypeError("Expected argument 'status' to be a dict")
        pulumi.set(__self__, "status", status)
        if target_scope and not isinstance(target_scope, str):
            raise TypeError("Expected argument 'target_scope' to be a str")
        pulumi.set(__self__, "target_scope", target_scope)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="blueprintName")
    def blueprint_name(self) -> Optional[str]:
        """
        Name of the Blueprint definition.
        """
        return pulumi.get(self, "blueprint_name")

    @property
    @pulumi.getter(name="changeNotes")
    def change_notes(self) -> Optional[str]:
        """
        Version-specific change notes
        """
        return pulumi.get(self, "change_notes")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        Multi-line explain this resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[str]:
        """
        One-liner string explain this resource.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of this resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def parameters(self) -> Optional[Mapping[str, 'outputs.ParameterDefinitionResponse']]:
        """
        Parameters required by this Blueprint definition.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter(name="resourceGroups")
    def resource_groups(self) -> Optional[Mapping[str, 'outputs.ResourceGroupDefinitionResponse']]:
        """
        Resource group placeholders defined by this Blueprint definition.
        """
        return pulumi.get(self, "resource_groups")

    @property
    @pulumi.getter
    def status(self) -> 'outputs.BlueprintStatusResponse':
        """
        Status of the Blueprint. This field is readonly.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="targetScope")
    def target_scope(self) -> Optional[str]:
        """
        The scope where this Blueprint can be applied.
        """
        return pulumi.get(self, "target_scope")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Type of this resource.
        """
        return pulumi.get(self, "type")


class AwaitableGetPublishedBlueprintResult(GetPublishedBlueprintResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPublishedBlueprintResult(
            blueprint_name=self.blueprint_name,
            change_notes=self.change_notes,
            description=self.description,
            display_name=self.display_name,
            name=self.name,
            parameters=self.parameters,
            resource_groups=self.resource_groups,
            status=self.status,
            target_scope=self.target_scope,
            type=self.type)


def get_published_blueprint(blueprint_name: Optional[str] = None,
                            management_group_name: Optional[str] = None,
                            version_id: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPublishedBlueprintResult:
    """
    Use this data source to access information about an existing resource.

    :param str blueprint_name: name of the blueprint.
    :param str management_group_name: ManagementGroup where blueprint stores.
    :param str version_id: version of the published blueprint.
    """
    __args__ = dict()
    __args__['blueprintName'] = blueprint_name
    __args__['managementGroupName'] = management_group_name
    __args__['versionId'] = version_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:management/v20171111preview:getPublishedBlueprint', __args__, opts=opts, typ=GetPublishedBlueprintResult).value

    return AwaitableGetPublishedBlueprintResult(
        blueprint_name=__ret__.blueprint_name,
        change_notes=__ret__.change_notes,
        description=__ret__.description,
        display_name=__ret__.display_name,
        name=__ret__.name,
        parameters=__ret__.parameters,
        resource_groups=__ret__.resource_groups,
        status=__ret__.status,
        target_scope=__ret__.target_scope,
        type=__ret__.type)
