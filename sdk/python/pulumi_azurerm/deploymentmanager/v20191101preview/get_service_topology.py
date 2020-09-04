# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'GetServiceTopologyResult',
    'AwaitableGetServiceTopologyResult',
    'get_service_topology',
]

@pulumi.output_type
class GetServiceTopologyResult:
    """
    The resource representation of a service topology.
    """
    def __init__(__self__, artifact_source_id=None, location=None, name=None, tags=None, type=None):
        if artifact_source_id and not isinstance(artifact_source_id, str):
            raise TypeError("Expected argument 'artifact_source_id' to be a str")
        pulumi.set(__self__, "artifact_source_id", artifact_source_id)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="artifactSourceId")
    def artifact_source_id(self) -> Optional[str]:
        """
        The resource Id of the artifact source that contains the artifacts that can be referenced in the service units.
        """
        return pulumi.get(self, "artifact_source_id")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        """
        return pulumi.get(self, "type")


class AwaitableGetServiceTopologyResult(GetServiceTopologyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServiceTopologyResult(
            artifact_source_id=self.artifact_source_id,
            location=self.location,
            name=self.name,
            tags=self.tags,
            type=self.type)


def get_service_topology(resource_group_name: Optional[str] = None,
                         service_topology_name: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServiceTopologyResult:
    """
    Use this data source to access information about an existing resource.

    :param str resource_group_name: The name of the resource group. The name is case insensitive.
    :param str service_topology_name: The name of the service topology .
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['serviceTopologyName'] = service_topology_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:deploymentmanager/v20191101preview:getServiceTopology', __args__, opts=opts, typ=GetServiceTopologyResult).value

    return AwaitableGetServiceTopologyResult(
        artifact_source_id=__ret__.artifact_source_id,
        location=__ret__.location,
        name=__ret__.name,
        tags=__ret__.tags,
        type=__ret__.type)
