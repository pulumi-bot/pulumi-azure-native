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
    'GetNetworkProfileResult',
    'AwaitableGetNetworkProfileResult',
    'get_network_profile',
]

@pulumi.output_type
class GetNetworkProfileResult:
    """
    Network profile resource.
    """
    def __init__(__self__, container_network_interface_configurations=None, container_network_interfaces=None, etag=None, location=None, name=None, provisioning_state=None, resource_guid=None, tags=None, type=None):
        if container_network_interface_configurations and not isinstance(container_network_interface_configurations, list):
            raise TypeError("Expected argument 'container_network_interface_configurations' to be a list")
        pulumi.set(__self__, "container_network_interface_configurations", container_network_interface_configurations)
        if container_network_interfaces and not isinstance(container_network_interfaces, list):
            raise TypeError("Expected argument 'container_network_interfaces' to be a list")
        pulumi.set(__self__, "container_network_interfaces", container_network_interfaces)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if resource_guid and not isinstance(resource_guid, str):
            raise TypeError("Expected argument 'resource_guid' to be a str")
        pulumi.set(__self__, "resource_guid", resource_guid)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="containerNetworkInterfaceConfigurations")
    def container_network_interface_configurations(self) -> Optional[List['outputs.ContainerNetworkInterfaceConfigurationResponse']]:
        """
        List of chid container network interface configurations.
        """
        return pulumi.get(self, "container_network_interface_configurations")

    @property
    @pulumi.getter(name="containerNetworkInterfaces")
    def container_network_interfaces(self) -> Optional[List['outputs.ContainerNetworkInterfaceResponse']]:
        """
        List of child container network interfaces.
        """
        return pulumi.get(self, "container_network_interfaces")

    @property
    @pulumi.getter
    def etag(self) -> Optional[str]:
        """
        A unique read-only string that changes whenever the resource is updated.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def location(self) -> Optional[str]:
        """
        Resource location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        The provisioning state of the resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="resourceGuid")
    def resource_guid(self) -> str:
        """
        The resource GUID property of the network interface resource.
        """
        return pulumi.get(self, "resource_guid")

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
        Resource type.
        """
        return pulumi.get(self, "type")


class AwaitableGetNetworkProfileResult(GetNetworkProfileResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNetworkProfileResult(
            container_network_interface_configurations=self.container_network_interface_configurations,
            container_network_interfaces=self.container_network_interfaces,
            etag=self.etag,
            location=self.location,
            name=self.name,
            provisioning_state=self.provisioning_state,
            resource_guid=self.resource_guid,
            tags=self.tags,
            type=self.type)


def get_network_profile(expand: Optional[str] = None,
                        name: Optional[str] = None,
                        resource_group_name: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNetworkProfileResult:
    """
    Use this data source to access information about an existing resource.

    :param str expand: Expands referenced resources.
    :param str name: The name of the PublicIPPrefix.
    :param str resource_group_name: The name of the resource group.
    """
    __args__ = dict()
    __args__['expand'] = expand
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:network/v20181101:getNetworkProfile', __args__, opts=opts, typ=GetNetworkProfileResult).value

    return AwaitableGetNetworkProfileResult(
        container_network_interface_configurations=__ret__.container_network_interface_configurations,
        container_network_interfaces=__ret__.container_network_interfaces,
        etag=__ret__.etag,
        location=__ret__.location,
        name=__ret__.name,
        provisioning_state=__ret__.provisioning_state,
        resource_guid=__ret__.resource_guid,
        tags=__ret__.tags,
        type=__ret__.type)
