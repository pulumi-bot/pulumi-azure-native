# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'GetScopeMapResult',
    'AwaitableGetScopeMapResult',
    'get_scope_map',
]

@pulumi.output_type
class GetScopeMapResult:
    """
    An object that represents a scope map for a container registry.
    """
    def __init__(__self__, actions=None, creation_date=None, description=None, name=None, provisioning_state=None, type=None):
        if actions and not isinstance(actions, list):
            raise TypeError("Expected argument 'actions' to be a list")
        pulumi.set(__self__, "actions", actions)
        if creation_date and not isinstance(creation_date, str):
            raise TypeError("Expected argument 'creation_date' to be a str")
        pulumi.set(__self__, "creation_date", creation_date)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def actions(self) -> List[str]:
        """
        The list of scoped permissions for registry artifacts.
        E.g. repositories/repository-name/content/read,
        repositories/repository-name/metadata/write
        """
        return pulumi.get(self, "actions")

    @property
    @pulumi.getter(name="creationDate")
    def creation_date(self) -> str:
        """
        The creation date of scope map.
        """
        return pulumi.get(self, "creation_date")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The user friendly description of the scope map.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        Provisioning state of the resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the resource.
        """
        return pulumi.get(self, "type")


class AwaitableGetScopeMapResult(GetScopeMapResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetScopeMapResult(
            actions=self.actions,
            creation_date=self.creation_date,
            description=self.description,
            name=self.name,
            provisioning_state=self.provisioning_state,
            type=self.type)


def get_scope_map(registry_name: Optional[str] = None,
                  resource_group_name: Optional[str] = None,
                  scope_map_name: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetScopeMapResult:
    """
    Use this data source to access information about an existing resource.

    :param str registry_name: The name of the container registry.
    :param str resource_group_name: The name of the resource group to which the container registry belongs.
    :param str scope_map_name: The name of the scope map.
    """
    __args__ = dict()
    __args__['registryName'] = registry_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['scopeMapName'] = scope_map_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:containerregistry/v20190501preview:getScopeMap', __args__, opts=opts, typ=GetScopeMapResult).value

    return AwaitableGetScopeMapResult(
        actions=__ret__.actions,
        creation_date=__ret__.creation_date,
        description=__ret__.description,
        name=__ret__.name,
        provisioning_state=__ret__.provisioning_state,
        type=__ret__.type)
