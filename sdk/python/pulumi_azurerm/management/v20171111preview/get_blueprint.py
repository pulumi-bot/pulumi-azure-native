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
    'GetBlueprintResult',
    'AwaitableGetBlueprintResult',
    'get_blueprint',
]

@pulumi.output_type
class GetBlueprintResult:
    """
    Represents a Blueprint definition.
    """
    def __init__(__self__, description=None, display_name=None, layout=None, name=None, parameters=None, resource_groups=None, status=None, target_scope=None, type=None, versions=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if layout and not isinstance(layout, dict):
            raise TypeError("Expected argument 'layout' to be a dict")
        pulumi.set(__self__, "layout", layout)
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
        if versions and not isinstance(versions, dict):
            raise TypeError("Expected argument 'versions' to be a dict")
        pulumi.set(__self__, "versions", versions)

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
    def layout(self) -> Optional[Mapping[str, Any]]:
        """
        Layout view of the blueprint, for UI reference.
        """
        return pulumi.get(self, "layout")

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
    def target_scope(self) -> str:
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

    @property
    @pulumi.getter
    def versions(self) -> Optional[Mapping[str, Any]]:
        """
        Published versions of this blueprint.
        """
        return pulumi.get(self, "versions")


class AwaitableGetBlueprintResult(GetBlueprintResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBlueprintResult(
            description=self.description,
            display_name=self.display_name,
            layout=self.layout,
            name=self.name,
            parameters=self.parameters,
            resource_groups=self.resource_groups,
            status=self.status,
            target_scope=self.target_scope,
            type=self.type,
            versions=self.versions)


def get_blueprint(blueprint_name: Optional[str] = None,
                  management_group_name: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBlueprintResult:
    """
    Use this data source to access information about an existing resource.

    :param str blueprint_name: name of the blueprint.
    :param str management_group_name: ManagementGroup where blueprint stores.
    """
    __args__ = dict()
    __args__['blueprintName'] = blueprint_name
    __args__['managementGroupName'] = management_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:management/v20171111preview:getBlueprint', __args__, opts=opts, typ=GetBlueprintResult).value

    return AwaitableGetBlueprintResult(
        description=__ret__.description,
        display_name=__ret__.display_name,
        layout=__ret__.layout,
        name=__ret__.name,
        parameters=__ret__.parameters,
        resource_groups=__ret__.resource_groups,
        status=__ret__.status,
        target_scope=__ret__.target_scope,
        type=__ret__.type,
        versions=__ret__.versions)
