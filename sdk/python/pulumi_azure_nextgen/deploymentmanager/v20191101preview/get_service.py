# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables

__all__ = [
    'GetServiceResult',
    'AwaitableGetServiceResult',
    'get_service',
]

@pulumi.output_type
class GetServiceResult:
    """
    The resource representation of a service in a service topology.
    """
    def __init__(__self__, location=None, name=None, tags=None, target_location=None, target_subscription_id=None, type=None):
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if target_location and not isinstance(target_location, str):
            raise TypeError("Expected argument 'target_location' to be a str")
        pulumi.set(__self__, "target_location", target_location)
        if target_subscription_id and not isinstance(target_subscription_id, str):
            raise TypeError("Expected argument 'target_subscription_id' to be a str")
        pulumi.set(__self__, "target_subscription_id", target_subscription_id)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

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
    @pulumi.getter(name="targetLocation")
    def target_location(self) -> str:
        """
        The Azure location to which the resources in the service belong to or should be deployed to.
        """
        return pulumi.get(self, "target_location")

    @property
    @pulumi.getter(name="targetSubscriptionId")
    def target_subscription_id(self) -> str:
        """
        The subscription to which the resources in the service belong to or should be deployed to.
        """
        return pulumi.get(self, "target_subscription_id")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")


class AwaitableGetServiceResult(GetServiceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServiceResult(
            location=self.location,
            name=self.name,
            tags=self.tags,
            target_location=self.target_location,
            target_subscription_id=self.target_subscription_id,
            type=self.type)


def get_service(resource_group_name: Optional[str] = None,
                service_name: Optional[str] = None,
                service_topology_name: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServiceResult:
    """
    Use this data source to access information about an existing resource.

    :param str resource_group_name: The name of the resource group. The name is case insensitive.
    :param str service_name: The name of the service resource.
    :param str service_topology_name: The name of the service topology .
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['serviceName'] = service_name
    __args__['serviceTopologyName'] = service_topology_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure-nextgen:deploymentmanager/v20191101preview:getService', __args__, opts=opts, typ=GetServiceResult).value

    return AwaitableGetServiceResult(
        location=__ret__.location,
        name=__ret__.name,
        tags=__ret__.tags,
        target_location=__ret__.target_location,
        target_subscription_id=__ret__.target_subscription_id,
        type=__ret__.type)
