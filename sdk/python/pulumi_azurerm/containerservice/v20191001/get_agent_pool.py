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
    Agent Pool.
    """
    def __init__(__self__, availability_zones=None, count=None, enable_auto_scaling=None, enable_node_public_ip=None, max_count=None, max_pods=None, min_count=None, name=None, node_taints=None, orchestrator_version=None, os_disk_size_gb=None, os_type=None, provisioning_state=None, scale_set_eviction_policy=None, scale_set_priority=None, type=None, vm_size=None, vnet_subnet_id=None):
        if availability_zones and not isinstance(availability_zones, list):
            raise TypeError("Expected argument 'availability_zones' to be a list")
        pulumi.set(__self__, "availability_zones", availability_zones)
        if count and not isinstance(count, float):
            raise TypeError("Expected argument 'count' to be a float")
        pulumi.set(__self__, "count", count)
        if enable_auto_scaling and not isinstance(enable_auto_scaling, bool):
            raise TypeError("Expected argument 'enable_auto_scaling' to be a bool")
        pulumi.set(__self__, "enable_auto_scaling", enable_auto_scaling)
        if enable_node_public_ip and not isinstance(enable_node_public_ip, bool):
            raise TypeError("Expected argument 'enable_node_public_ip' to be a bool")
        pulumi.set(__self__, "enable_node_public_ip", enable_node_public_ip)
        if max_count and not isinstance(max_count, float):
            raise TypeError("Expected argument 'max_count' to be a float")
        pulumi.set(__self__, "max_count", max_count)
        if max_pods and not isinstance(max_pods, float):
            raise TypeError("Expected argument 'max_pods' to be a float")
        pulumi.set(__self__, "max_pods", max_pods)
        if min_count and not isinstance(min_count, float):
            raise TypeError("Expected argument 'min_count' to be a float")
        pulumi.set(__self__, "min_count", min_count)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if node_taints and not isinstance(node_taints, list):
            raise TypeError("Expected argument 'node_taints' to be a list")
        pulumi.set(__self__, "node_taints", node_taints)
        if orchestrator_version and not isinstance(orchestrator_version, str):
            raise TypeError("Expected argument 'orchestrator_version' to be a str")
        pulumi.set(__self__, "orchestrator_version", orchestrator_version)
        if os_disk_size_gb and not isinstance(os_disk_size_gb, float):
            raise TypeError("Expected argument 'os_disk_size_gb' to be a float")
        pulumi.set(__self__, "os_disk_size_gb", os_disk_size_gb)
        if os_type and not isinstance(os_type, str):
            raise TypeError("Expected argument 'os_type' to be a str")
        pulumi.set(__self__, "os_type", os_type)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if scale_set_eviction_policy and not isinstance(scale_set_eviction_policy, str):
            raise TypeError("Expected argument 'scale_set_eviction_policy' to be a str")
        pulumi.set(__self__, "scale_set_eviction_policy", scale_set_eviction_policy)
        if scale_set_priority and not isinstance(scale_set_priority, str):
            raise TypeError("Expected argument 'scale_set_priority' to be a str")
        pulumi.set(__self__, "scale_set_priority", scale_set_priority)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if vm_size and not isinstance(vm_size, str):
            raise TypeError("Expected argument 'vm_size' to be a str")
        pulumi.set(__self__, "vm_size", vm_size)
        if vnet_subnet_id and not isinstance(vnet_subnet_id, str):
            raise TypeError("Expected argument 'vnet_subnet_id' to be a str")
        pulumi.set(__self__, "vnet_subnet_id", vnet_subnet_id)

    @property
    @pulumi.getter(name="availabilityZones")
    def availability_zones(self) -> Optional[List[str]]:
        """
        (PREVIEW) Availability zones for nodes. Must use VirtualMachineScaleSets AgentPoolType.
        """
        return pulumi.get(self, "availability_zones")

    @property
    @pulumi.getter
    def count(self) -> Optional[float]:
        """
        Number of agents (VMs) to host docker containers. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1.
        """
        return pulumi.get(self, "count")

    @property
    @pulumi.getter(name="enableAutoScaling")
    def enable_auto_scaling(self) -> Optional[bool]:
        """
        Whether to enable auto-scaler
        """
        return pulumi.get(self, "enable_auto_scaling")

    @property
    @pulumi.getter(name="enableNodePublicIP")
    def enable_node_public_ip(self) -> Optional[bool]:
        """
        Enable public IP for nodes
        """
        return pulumi.get(self, "enable_node_public_ip")

    @property
    @pulumi.getter(name="maxCount")
    def max_count(self) -> Optional[float]:
        """
        Maximum number of nodes for auto-scaling
        """
        return pulumi.get(self, "max_count")

    @property
    @pulumi.getter(name="maxPods")
    def max_pods(self) -> Optional[float]:
        """
        Maximum number of pods that can run on a node.
        """
        return pulumi.get(self, "max_pods")

    @property
    @pulumi.getter(name="minCount")
    def min_count(self) -> Optional[float]:
        """
        Minimum number of nodes for auto-scaling
        """
        return pulumi.get(self, "min_count")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource that is unique within a resource group. This name can be used to access the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nodeTaints")
    def node_taints(self) -> Optional[List[str]]:
        """
        Taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.
        """
        return pulumi.get(self, "node_taints")

    @property
    @pulumi.getter(name="orchestratorVersion")
    def orchestrator_version(self) -> Optional[str]:
        """
        Version of orchestrator specified when creating the managed cluster.
        """
        return pulumi.get(self, "orchestrator_version")

    @property
    @pulumi.getter(name="osDiskSizeGB")
    def os_disk_size_gb(self) -> Optional[float]:
        """
        OS Disk Size in GB to be used to specify the disk size for every machine in this master/agent pool. If you specify 0, it will apply the default osDisk size according to the vmSize specified.
        """
        return pulumi.get(self, "os_disk_size_gb")

    @property
    @pulumi.getter(name="osType")
    def os_type(self) -> Optional[str]:
        """
        OsType to be used to specify os type. Choose from Linux and Windows. Default to Linux.
        """
        return pulumi.get(self, "os_type")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        The current deployment or provisioning state, which only appears in the response.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="scaleSetEvictionPolicy")
    def scale_set_eviction_policy(self) -> Optional[str]:
        """
        ScaleSetEvictionPolicy to be used to specify eviction policy for low priority virtual machine scale set. Default to Delete.
        """
        return pulumi.get(self, "scale_set_eviction_policy")

    @property
    @pulumi.getter(name="scaleSetPriority")
    def scale_set_priority(self) -> Optional[str]:
        """
        ScaleSetPriority to be used to specify virtual machine scale set priority. Default to regular.
        """
        return pulumi.get(self, "scale_set_priority")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        AgentPoolType represents types of an agent pool
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="vmSize")
    def vm_size(self) -> Optional[str]:
        """
        Size of agent VMs.
        """
        return pulumi.get(self, "vm_size")

    @property
    @pulumi.getter(name="vnetSubnetID")
    def vnet_subnet_id(self) -> Optional[str]:
        """
        VNet SubnetID specifies the VNet's subnet identifier.
        """
        return pulumi.get(self, "vnet_subnet_id")


class AwaitableGetAgentPoolResult(GetAgentPoolResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAgentPoolResult(
            availability_zones=self.availability_zones,
            count=self.count,
            enable_auto_scaling=self.enable_auto_scaling,
            enable_node_public_ip=self.enable_node_public_ip,
            max_count=self.max_count,
            max_pods=self.max_pods,
            min_count=self.min_count,
            name=self.name,
            node_taints=self.node_taints,
            orchestrator_version=self.orchestrator_version,
            os_disk_size_gb=self.os_disk_size_gb,
            os_type=self.os_type,
            provisioning_state=self.provisioning_state,
            scale_set_eviction_policy=self.scale_set_eviction_policy,
            scale_set_priority=self.scale_set_priority,
            type=self.type,
            vm_size=self.vm_size,
            vnet_subnet_id=self.vnet_subnet_id)


def get_agent_pool(name: Optional[str] = None,
                   resource_group_name: Optional[str] = None,
                   resource_name: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAgentPoolResult:
    """
    Use this data source to access information about an existing resource.

    :param str name: The name of the agent pool.
    :param str resource_group_name: The name of the resource group.
    :param str resource_name: The name of the managed cluster resource.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    __args__['resourceName'] = resource_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:containerservice/v20191001:getAgentPool', __args__, opts=opts, typ=GetAgentPoolResult).value

    return AwaitableGetAgentPoolResult(
        availability_zones=__ret__.availability_zones,
        count=__ret__.count,
        enable_auto_scaling=__ret__.enable_auto_scaling,
        enable_node_public_ip=__ret__.enable_node_public_ip,
        max_count=__ret__.max_count,
        max_pods=__ret__.max_pods,
        min_count=__ret__.min_count,
        name=__ret__.name,
        node_taints=__ret__.node_taints,
        orchestrator_version=__ret__.orchestrator_version,
        os_disk_size_gb=__ret__.os_disk_size_gb,
        os_type=__ret__.os_type,
        provisioning_state=__ret__.provisioning_state,
        scale_set_eviction_policy=__ret__.scale_set_eviction_policy,
        scale_set_priority=__ret__.scale_set_priority,
        type=__ret__.type,
        vm_size=__ret__.vm_size,
        vnet_subnet_id=__ret__.vnet_subnet_id)
