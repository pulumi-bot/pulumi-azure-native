# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'GetArtifactSourceResourceResult',
    'AwaitableGetArtifactSourceResourceResult',
    'get_artifact_source_resource',
]

@pulumi.output_type
class GetArtifactSourceResourceResult:
    """
    Properties of an artifact source.
    """
    def __init__(__self__, branch_ref=None, display_name=None, folder_path=None, location=None, name=None, provisioning_state=None, security_token=None, source_type=None, status=None, tags=None, type=None, uri=None):
        if branch_ref and not isinstance(branch_ref, str):
            raise TypeError("Expected argument 'branch_ref' to be a str")
        pulumi.set(__self__, "branch_ref", branch_ref)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if folder_path and not isinstance(folder_path, str):
            raise TypeError("Expected argument 'folder_path' to be a str")
        pulumi.set(__self__, "folder_path", folder_path)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if security_token and not isinstance(security_token, str):
            raise TypeError("Expected argument 'security_token' to be a str")
        pulumi.set(__self__, "security_token", security_token)
        if source_type and not isinstance(source_type, str):
            raise TypeError("Expected argument 'source_type' to be a str")
        pulumi.set(__self__, "source_type", source_type)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if uri and not isinstance(uri, str):
            raise TypeError("Expected argument 'uri' to be a str")
        pulumi.set(__self__, "uri", uri)

    @property
    @pulumi.getter(name="branchRef")
    def branch_ref(self) -> Optional[str]:
        """
        The branch reference of the artifact source.
        """
        return pulumi.get(self, "branch_ref")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[str]:
        """
        The display name of the artifact source.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="folderPath")
    def folder_path(self) -> Optional[str]:
        """
        The folder path of the artifact source.
        """
        return pulumi.get(self, "folder_path")

    @property
    @pulumi.getter
    def location(self) -> Optional[str]:
        """
        The location of the resource.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> Optional[str]:
        """
        The provisioning status of the resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="securityToken")
    def security_token(self) -> Optional[str]:
        """
        The security token of the artifact source.
        """
        return pulumi.get(self, "security_token")

    @property
    @pulumi.getter(name="sourceType")
    def source_type(self) -> Optional[str]:
        """
        The type of the artifact source.
        """
        return pulumi.get(self, "source_type")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        The status of the artifact source.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        The tags of the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        The type of the resource.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def uri(self) -> Optional[str]:
        """
        The URI of the artifact source.
        """
        return pulumi.get(self, "uri")


class AwaitableGetArtifactSourceResourceResult(GetArtifactSourceResourceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetArtifactSourceResourceResult(
            branch_ref=self.branch_ref,
            display_name=self.display_name,
            folder_path=self.folder_path,
            location=self.location,
            name=self.name,
            provisioning_state=self.provisioning_state,
            security_token=self.security_token,
            source_type=self.source_type,
            status=self.status,
            tags=self.tags,
            type=self.type,
            uri=self.uri)


def get_artifact_source_resource(lab_name: Optional[str] = None,
                                 name: Optional[str] = None,
                                 resource_group_name: Optional[str] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetArtifactSourceResourceResult:
    """
    Use this data source to access information about an existing resource.

    :param str lab_name: The name of the lab.
    :param str name: The name of the artifact source.
    :param str resource_group_name: The name of the resource group.
    """
    __args__ = dict()
    __args__['labName'] = lab_name
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:devtestlab/v20150521preview:getArtifactSourceResource', __args__, opts=opts, typ=GetArtifactSourceResourceResult).value

    return AwaitableGetArtifactSourceResourceResult(
        branch_ref=__ret__.branch_ref,
        display_name=__ret__.display_name,
        folder_path=__ret__.folder_path,
        location=__ret__.location,
        name=__ret__.name,
        provisioning_state=__ret__.provisioning_state,
        security_token=__ret__.security_token,
        source_type=__ret__.source_type,
        status=__ret__.status,
        tags=__ret__.tags,
        type=__ret__.type,
        uri=__ret__.uri)
