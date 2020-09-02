# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'GetAgentPoolResult',
    'AwaitableGetAgentPoolResult',
    'get_agent_pool',
]

@pulumi.output_type
class GetAgentPoolResult:
    """
    The agentpool that has the ARM resource and properties. 
    The agentpool will have all information to create an agent pool.
    """
    def __init__(__self__, count=None, location=None, name=None, os=None, provisioning_state=None, tags=None, tier=None, type=None, virtual_network_subnet_resource_id=None):
        if count and not isinstance(count, float):
            raise TypeError("Expected argument 'count' to be a float")
        pulumi.set(__self__, "count", count)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if os and not isinstance(os, str):
            raise TypeError("Expected argument 'os' to be a str")
        pulumi.set(__self__, "os", os)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if tier and not isinstance(tier, str):
            raise TypeError("Expected argument 'tier' to be a str")
        pulumi.set(__self__, "tier", tier)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if virtual_network_subnet_resource_id and not isinstance(virtual_network_subnet_resource_id, str):
            raise TypeError("Expected argument 'virtual_network_subnet_resource_id' to be a str")
        pulumi.set(__self__, "virtual_network_subnet_resource_id", virtual_network_subnet_resource_id)

    @property
    @pulumi.getter
    def count(self) -> Optional[float]:
        """
        The count of agent machine
        """
        return pulumi.get(self, "count")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        The location of the resource. This cannot be changed after the resource is created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def os(self) -> Optional[str]:
        """
        The OS of agent machine
        """
        return pulumi.get(self, "os")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        The provisioning state of this agent pool
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        The tags of the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def tier(self) -> Optional[str]:
        """
        The Tier of agent machine
        """
        return pulumi.get(self, "tier")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the resource.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="virtualNetworkSubnetResourceId")
    def virtual_network_subnet_resource_id(self) -> Optional[str]:
        """
        The Virtual Network Subnet Resource Id of the agent machine
        """
        return pulumi.get(self, "virtual_network_subnet_resource_id")


class AwaitableGetAgentPoolResult(GetAgentPoolResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAgentPoolResult(
            count=self.count,
            location=self.location,
            name=self.name,
            os=self.os,
            provisioning_state=self.provisioning_state,
            tags=self.tags,
            tier=self.tier,
            type=self.type,
            virtual_network_subnet_resource_id=self.virtual_network_subnet_resource_id)


def get_agent_pool(agent_pool_name: Optional[str] = None,
                   registry_name: Optional[str] = None,
                   resource_group_name: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAgentPoolResult:
    """
    Use this data source to access information about an existing resource.

    :param str agent_pool_name: The name of the agent pool.
    :param str registry_name: The name of the container registry.
    :param str resource_group_name: The name of the resource group to which the container registry belongs.
    """
    __args__ = dict()
    __args__['agentPoolName'] = agent_pool_name
    __args__['registryName'] = registry_name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:containerregistry/v20190601preview:getAgentPool', __args__, opts=opts, typ=GetAgentPoolResult).value

    return AwaitableGetAgentPoolResult(
        count=__ret__.count,
        location=__ret__.location,
        name=__ret__.name,
        os=__ret__.os,
        provisioning_state=__ret__.provisioning_state,
        tags=__ret__.tags,
        tier=__ret__.tier,
        type=__ret__.type,
        virtual_network_subnet_resource_id=__ret__.virtual_network_subnet_resource_id)
