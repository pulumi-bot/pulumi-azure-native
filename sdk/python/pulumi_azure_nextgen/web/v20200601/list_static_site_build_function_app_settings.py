# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables

__all__ = [
    'ListStaticSiteBuildFunctionAppSettingsResult',
    'AwaitableListStaticSiteBuildFunctionAppSettingsResult',
    'list_static_site_build_function_app_settings',
]

@pulumi.output_type
class ListStaticSiteBuildFunctionAppSettingsResult:
    """
    String dictionary resource.
    """
    def __init__(__self__, kind=None, name=None, properties=None, type=None):
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        pulumi.set(__self__, "properties", properties)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def kind(self) -> Optional[str]:
        """
        Kind of resource.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource Name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def properties(self) -> Mapping[str, str]:
        """
        Settings.
        """
        return pulumi.get(self, "properties")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type.
        """
        return pulumi.get(self, "type")


class AwaitableListStaticSiteBuildFunctionAppSettingsResult(ListStaticSiteBuildFunctionAppSettingsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListStaticSiteBuildFunctionAppSettingsResult(
            kind=self.kind,
            name=self.name,
            properties=self.properties,
            type=self.type)


def list_static_site_build_function_app_settings(name: Optional[str] = None,
                                                 pr_id: Optional[str] = None,
                                                 resource_group_name: Optional[str] = None,
                                                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListStaticSiteBuildFunctionAppSettingsResult:
    """
    Use this data source to access information about an existing resource.

    :param str name: Name of the static site.
    :param str pr_id: The stage site identifier.
    :param str resource_group_name: Name of the resource group to which the resource belongs.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['prId'] = pr_id
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure-nextgen:web/v20200601:listStaticSiteBuildFunctionAppSettings', __args__, opts=opts, typ=ListStaticSiteBuildFunctionAppSettingsResult).value

    return AwaitableListStaticSiteBuildFunctionAppSettingsResult(
        kind=__ret__.kind,
        name=__ret__.name,
        properties=__ret__.properties,
        type=__ret__.type)
