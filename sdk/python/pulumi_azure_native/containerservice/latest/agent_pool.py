# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['AgentPool']

warnings.warn("""The 'latest' version is deprecated. Please migrate to the resource in the top-level module: 'azure-native:containerservice:AgentPool'.""", DeprecationWarning)


class AgentPool(pulumi.CustomResource):
    warnings.warn("""The 'latest' version is deprecated. Please migrate to the resource in the top-level module: 'azure-native:containerservice:AgentPool'.""", DeprecationWarning)

    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 agent_pool_name: Optional[pulumi.Input[str]] = None,
                 availability_zones: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 count: Optional[pulumi.Input[int]] = None,
                 enable_auto_scaling: Optional[pulumi.Input[bool]] = None,
                 enable_encryption_at_host: Optional[pulumi.Input[bool]] = None,
                 enable_node_public_ip: Optional[pulumi.Input[bool]] = None,
                 kubelet_config: Optional[pulumi.Input[pulumi.InputType['KubeletConfigArgs']]] = None,
                 kubelet_disk_type: Optional[pulumi.Input[Union[str, 'KubeletDiskType']]] = None,
                 linux_os_config: Optional[pulumi.Input[pulumi.InputType['LinuxOSConfigArgs']]] = None,
                 max_count: Optional[pulumi.Input[int]] = None,
                 max_pods: Optional[pulumi.Input[int]] = None,
                 min_count: Optional[pulumi.Input[int]] = None,
                 mode: Optional[pulumi.Input[Union[str, 'AgentPoolMode']]] = None,
                 node_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 node_public_ip_prefix_id: Optional[pulumi.Input[str]] = None,
                 node_taints: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 orchestrator_version: Optional[pulumi.Input[str]] = None,
                 os_disk_size_gb: Optional[pulumi.Input[int]] = None,
                 os_disk_type: Optional[pulumi.Input[Union[str, 'OSDiskType']]] = None,
                 os_type: Optional[pulumi.Input[Union[str, 'OSType']]] = None,
                 pod_subnet_id: Optional[pulumi.Input[str]] = None,
                 proximity_placement_group_id: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 resource_name_: Optional[pulumi.Input[str]] = None,
                 scale_set_eviction_policy: Optional[pulumi.Input[Union[str, 'ScaleSetEvictionPolicy']]] = None,
                 scale_set_priority: Optional[pulumi.Input[Union[str, 'ScaleSetPriority']]] = None,
                 spot_max_price: Optional[pulumi.Input[float]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[Union[str, 'AgentPoolType']]] = None,
                 upgrade_settings: Optional[pulumi.Input[pulumi.InputType['AgentPoolUpgradeSettingsArgs']]] = None,
                 vm_size: Optional[pulumi.Input[Union[str, 'ContainerServiceVMSizeTypes']]] = None,
                 vnet_subnet_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Agent Pool.
        Latest API Version: 2021-02-01.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] agent_pool_name: The name of the agent pool.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] availability_zones: Availability zones for nodes. Must use VirtualMachineScaleSets AgentPoolType.
        :param pulumi.Input[int] count: Number of agents (VMs) to host docker containers. Allowed values must be in the range of 0 to 100 (inclusive) for user pools and in the range of 1 to 100 (inclusive) for system pools. The default value is 1.
        :param pulumi.Input[bool] enable_auto_scaling: Whether to enable auto-scaler
        :param pulumi.Input[bool] enable_encryption_at_host: Whether to enable EncryptionAtHost
        :param pulumi.Input[bool] enable_node_public_ip: Enable public IP for nodes
        :param pulumi.Input[pulumi.InputType['KubeletConfigArgs']] kubelet_config: KubeletConfig specifies the configuration of kubelet on agent nodes.
        :param pulumi.Input[Union[str, 'KubeletDiskType']] kubelet_disk_type: KubeletDiskType determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage. Currently allows one value, OS, resulting in Kubelet using the OS disk for data.
        :param pulumi.Input[pulumi.InputType['LinuxOSConfigArgs']] linux_os_config: LinuxOSConfig specifies the OS configuration of linux agent nodes.
        :param pulumi.Input[int] max_count: Maximum number of nodes for auto-scaling
        :param pulumi.Input[int] max_pods: Maximum number of pods that can run on a node.
        :param pulumi.Input[int] min_count: Minimum number of nodes for auto-scaling
        :param pulumi.Input[Union[str, 'AgentPoolMode']] mode: AgentPoolMode represents mode of an agent pool
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] node_labels: Agent pool node labels to be persisted across all nodes in agent pool.
        :param pulumi.Input[str] node_public_ip_prefix_id: Public IP Prefix ID. VM nodes use IPs assigned from this Public IP Prefix.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] node_taints: Taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.
        :param pulumi.Input[str] orchestrator_version: Version of orchestrator specified when creating the managed cluster.
        :param pulumi.Input[int] os_disk_size_gb: OS Disk Size in GB to be used to specify the disk size for every machine in this master/agent pool. If you specify 0, it will apply the default osDisk size according to the vmSize specified.
        :param pulumi.Input[Union[str, 'OSDiskType']] os_disk_type: OS disk type to be used for machines in a given agent pool. Allowed values are 'Ephemeral' and 'Managed'. Defaults to 'Managed'. May not be changed after creation.
        :param pulumi.Input[Union[str, 'OSType']] os_type: OsType to be used to specify os type. Choose from Linux and Windows. Default to Linux.
        :param pulumi.Input[str] pod_subnet_id: Pod SubnetID specifies the VNet's subnet identifier for pods.
        :param pulumi.Input[str] proximity_placement_group_id: The ID for Proximity Placement Group.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] resource_name_: The name of the managed cluster resource.
        :param pulumi.Input[Union[str, 'ScaleSetEvictionPolicy']] scale_set_eviction_policy: ScaleSetEvictionPolicy to be used to specify eviction policy for Spot virtual machine scale set. Default to Delete.
        :param pulumi.Input[Union[str, 'ScaleSetPriority']] scale_set_priority: ScaleSetPriority to be used to specify virtual machine scale set priority. Default to regular.
        :param pulumi.Input[float] spot_max_price: SpotMaxPrice to be used to specify the maximum price you are willing to pay in US Dollars. Possible values are any decimal value greater than zero or -1 which indicates default price to be up-to on-demand.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Agent pool tags to be persisted on the agent pool virtual machine scale set.
        :param pulumi.Input[Union[str, 'AgentPoolType']] type: AgentPoolType represents types of an agent pool
        :param pulumi.Input[pulumi.InputType['AgentPoolUpgradeSettingsArgs']] upgrade_settings: Settings for upgrading the agentpool
        :param pulumi.Input[Union[str, 'ContainerServiceVMSizeTypes']] vm_size: Size of agent VMs.
        :param pulumi.Input[str] vnet_subnet_id: VNet SubnetID specifies the VNet's subnet identifier for nodes and maybe pods
        """
        pulumi.log.warn("""AgentPool is deprecated: The 'latest' version is deprecated. Please migrate to the resource in the top-level module: 'azure-native:containerservice:AgentPool'.""")
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['agent_pool_name'] = agent_pool_name
            __props__['availability_zones'] = availability_zones
            __props__['count'] = count
            __props__['enable_auto_scaling'] = enable_auto_scaling
            __props__['enable_encryption_at_host'] = enable_encryption_at_host
            __props__['enable_node_public_ip'] = enable_node_public_ip
            __props__['kubelet_config'] = kubelet_config
            __props__['kubelet_disk_type'] = kubelet_disk_type
            __props__['linux_os_config'] = linux_os_config
            __props__['max_count'] = max_count
            __props__['max_pods'] = max_pods
            __props__['min_count'] = min_count
            __props__['mode'] = mode
            __props__['node_labels'] = node_labels
            __props__['node_public_ip_prefix_id'] = node_public_ip_prefix_id
            __props__['node_taints'] = node_taints
            __props__['orchestrator_version'] = orchestrator_version
            __props__['os_disk_size_gb'] = os_disk_size_gb
            __props__['os_disk_type'] = os_disk_type
            __props__['os_type'] = os_type
            __props__['pod_subnet_id'] = pod_subnet_id
            __props__['proximity_placement_group_id'] = proximity_placement_group_id
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if resource_name_ is None and not opts.urn:
                raise TypeError("Missing required property 'resource_name_'")
            __props__['resource_name'] = resource_name_
            __props__['scale_set_eviction_policy'] = scale_set_eviction_policy
            __props__['scale_set_priority'] = scale_set_priority
            __props__['spot_max_price'] = spot_max_price
            __props__['tags'] = tags
            __props__['type'] = type
            __props__['upgrade_settings'] = upgrade_settings
            __props__['vm_size'] = vm_size
            __props__['vnet_subnet_id'] = vnet_subnet_id
            __props__['name'] = None
            __props__['node_image_version'] = None
            __props__['power_state'] = None
            __props__['provisioning_state'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:containerservice/latest:AgentPool"), pulumi.Alias(type_="azure-native:containerservice:AgentPool"), pulumi.Alias(type_="azure-nextgen:containerservice:AgentPool"), pulumi.Alias(type_="azure-native:containerservice/v20190201:AgentPool"), pulumi.Alias(type_="azure-nextgen:containerservice/v20190201:AgentPool"), pulumi.Alias(type_="azure-native:containerservice/v20190401:AgentPool"), pulumi.Alias(type_="azure-nextgen:containerservice/v20190401:AgentPool"), pulumi.Alias(type_="azure-native:containerservice/v20190601:AgentPool"), pulumi.Alias(type_="azure-nextgen:containerservice/v20190601:AgentPool"), pulumi.Alias(type_="azure-native:containerservice/v20190801:AgentPool"), pulumi.Alias(type_="azure-nextgen:containerservice/v20190801:AgentPool"), pulumi.Alias(type_="azure-native:containerservice/v20191001:AgentPool"), pulumi.Alias(type_="azure-nextgen:containerservice/v20191001:AgentPool"), pulumi.Alias(type_="azure-native:containerservice/v20191101:AgentPool"), pulumi.Alias(type_="azure-nextgen:containerservice/v20191101:AgentPool"), pulumi.Alias(type_="azure-native:containerservice/v20200101:AgentPool"), pulumi.Alias(type_="azure-nextgen:containerservice/v20200101:AgentPool"), pulumi.Alias(type_="azure-native:containerservice/v20200201:AgentPool"), pulumi.Alias(type_="azure-nextgen:containerservice/v20200201:AgentPool"), pulumi.Alias(type_="azure-native:containerservice/v20200301:AgentPool"), pulumi.Alias(type_="azure-nextgen:containerservice/v20200301:AgentPool"), pulumi.Alias(type_="azure-native:containerservice/v20200401:AgentPool"), pulumi.Alias(type_="azure-nextgen:containerservice/v20200401:AgentPool"), pulumi.Alias(type_="azure-native:containerservice/v20200601:AgentPool"), pulumi.Alias(type_="azure-nextgen:containerservice/v20200601:AgentPool"), pulumi.Alias(type_="azure-native:containerservice/v20200701:AgentPool"), pulumi.Alias(type_="azure-nextgen:containerservice/v20200701:AgentPool"), pulumi.Alias(type_="azure-native:containerservice/v20200901:AgentPool"), pulumi.Alias(type_="azure-nextgen:containerservice/v20200901:AgentPool"), pulumi.Alias(type_="azure-native:containerservice/v20201101:AgentPool"), pulumi.Alias(type_="azure-nextgen:containerservice/v20201101:AgentPool"), pulumi.Alias(type_="azure-native:containerservice/v20201201:AgentPool"), pulumi.Alias(type_="azure-nextgen:containerservice/v20201201:AgentPool"), pulumi.Alias(type_="azure-native:containerservice/v20210201:AgentPool"), pulumi.Alias(type_="azure-nextgen:containerservice/v20210201:AgentPool")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(AgentPool, __self__).__init__(
            'azure-native:containerservice/latest:AgentPool',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'AgentPool':
        """
        Get an existing AgentPool resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["availability_zones"] = None
        __props__["count"] = None
        __props__["enable_auto_scaling"] = None
        __props__["enable_encryption_at_host"] = None
        __props__["enable_node_public_ip"] = None
        __props__["kubelet_config"] = None
        __props__["kubelet_disk_type"] = None
        __props__["linux_os_config"] = None
        __props__["max_count"] = None
        __props__["max_pods"] = None
        __props__["min_count"] = None
        __props__["mode"] = None
        __props__["name"] = None
        __props__["node_image_version"] = None
        __props__["node_labels"] = None
        __props__["node_public_ip_prefix_id"] = None
        __props__["node_taints"] = None
        __props__["orchestrator_version"] = None
        __props__["os_disk_size_gb"] = None
        __props__["os_disk_type"] = None
        __props__["os_type"] = None
        __props__["pod_subnet_id"] = None
        __props__["power_state"] = None
        __props__["provisioning_state"] = None
        __props__["proximity_placement_group_id"] = None
        __props__["scale_set_eviction_policy"] = None
        __props__["scale_set_priority"] = None
        __props__["spot_max_price"] = None
        __props__["tags"] = None
        __props__["type"] = None
        __props__["upgrade_settings"] = None
        __props__["vm_size"] = None
        __props__["vnet_subnet_id"] = None
        return AgentPool(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="availabilityZones")
    def availability_zones(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Availability zones for nodes. Must use VirtualMachineScaleSets AgentPoolType.
        """
        return pulumi.get(self, "availability_zones")

    @property
    @pulumi.getter
    def count(self) -> pulumi.Output[Optional[int]]:
        """
        Number of agents (VMs) to host docker containers. Allowed values must be in the range of 0 to 100 (inclusive) for user pools and in the range of 1 to 100 (inclusive) for system pools. The default value is 1.
        """
        return pulumi.get(self, "count")

    @property
    @pulumi.getter(name="enableAutoScaling")
    def enable_auto_scaling(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to enable auto-scaler
        """
        return pulumi.get(self, "enable_auto_scaling")

    @property
    @pulumi.getter(name="enableEncryptionAtHost")
    def enable_encryption_at_host(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to enable EncryptionAtHost
        """
        return pulumi.get(self, "enable_encryption_at_host")

    @property
    @pulumi.getter(name="enableNodePublicIP")
    def enable_node_public_ip(self) -> pulumi.Output[Optional[bool]]:
        """
        Enable public IP for nodes
        """
        return pulumi.get(self, "enable_node_public_ip")

    @property
    @pulumi.getter(name="kubeletConfig")
    def kubelet_config(self) -> pulumi.Output[Optional['outputs.KubeletConfigResponse']]:
        """
        KubeletConfig specifies the configuration of kubelet on agent nodes.
        """
        return pulumi.get(self, "kubelet_config")

    @property
    @pulumi.getter(name="kubeletDiskType")
    def kubelet_disk_type(self) -> pulumi.Output[Optional[str]]:
        """
        KubeletDiskType determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage. Currently allows one value, OS, resulting in Kubelet using the OS disk for data.
        """
        return pulumi.get(self, "kubelet_disk_type")

    @property
    @pulumi.getter(name="linuxOSConfig")
    def linux_os_config(self) -> pulumi.Output[Optional['outputs.LinuxOSConfigResponse']]:
        """
        LinuxOSConfig specifies the OS configuration of linux agent nodes.
        """
        return pulumi.get(self, "linux_os_config")

    @property
    @pulumi.getter(name="maxCount")
    def max_count(self) -> pulumi.Output[Optional[int]]:
        """
        Maximum number of nodes for auto-scaling
        """
        return pulumi.get(self, "max_count")

    @property
    @pulumi.getter(name="maxPods")
    def max_pods(self) -> pulumi.Output[Optional[int]]:
        """
        Maximum number of pods that can run on a node.
        """
        return pulumi.get(self, "max_pods")

    @property
    @pulumi.getter(name="minCount")
    def min_count(self) -> pulumi.Output[Optional[int]]:
        """
        Minimum number of nodes for auto-scaling
        """
        return pulumi.get(self, "min_count")

    @property
    @pulumi.getter
    def mode(self) -> pulumi.Output[Optional[str]]:
        """
        AgentPoolMode represents mode of an agent pool
        """
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the resource that is unique within a resource group. This name can be used to access the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nodeImageVersion")
    def node_image_version(self) -> pulumi.Output[str]:
        """
        Version of node image
        """
        return pulumi.get(self, "node_image_version")

    @property
    @pulumi.getter(name="nodeLabels")
    def node_labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Agent pool node labels to be persisted across all nodes in agent pool.
        """
        return pulumi.get(self, "node_labels")

    @property
    @pulumi.getter(name="nodePublicIPPrefixID")
    def node_public_ip_prefix_id(self) -> pulumi.Output[Optional[str]]:
        """
        Public IP Prefix ID. VM nodes use IPs assigned from this Public IP Prefix.
        """
        return pulumi.get(self, "node_public_ip_prefix_id")

    @property
    @pulumi.getter(name="nodeTaints")
    def node_taints(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.
        """
        return pulumi.get(self, "node_taints")

    @property
    @pulumi.getter(name="orchestratorVersion")
    def orchestrator_version(self) -> pulumi.Output[Optional[str]]:
        """
        Version of orchestrator specified when creating the managed cluster.
        """
        return pulumi.get(self, "orchestrator_version")

    @property
    @pulumi.getter(name="osDiskSizeGB")
    def os_disk_size_gb(self) -> pulumi.Output[Optional[int]]:
        """
        OS Disk Size in GB to be used to specify the disk size for every machine in this master/agent pool. If you specify 0, it will apply the default osDisk size according to the vmSize specified.
        """
        return pulumi.get(self, "os_disk_size_gb")

    @property
    @pulumi.getter(name="osDiskType")
    def os_disk_type(self) -> pulumi.Output[Optional[str]]:
        """
        OS disk type to be used for machines in a given agent pool. Allowed values are 'Ephemeral' and 'Managed'. Defaults to 'Managed'. May not be changed after creation.
        """
        return pulumi.get(self, "os_disk_type")

    @property
    @pulumi.getter(name="osType")
    def os_type(self) -> pulumi.Output[Optional[str]]:
        """
        OsType to be used to specify os type. Choose from Linux and Windows. Default to Linux.
        """
        return pulumi.get(self, "os_type")

    @property
    @pulumi.getter(name="podSubnetID")
    def pod_subnet_id(self) -> pulumi.Output[Optional[str]]:
        """
        Pod SubnetID specifies the VNet's subnet identifier for pods.
        """
        return pulumi.get(self, "pod_subnet_id")

    @property
    @pulumi.getter(name="powerState")
    def power_state(self) -> pulumi.Output['outputs.PowerStateResponse']:
        """
        Describes whether the Agent Pool is Running or Stopped
        """
        return pulumi.get(self, "power_state")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        The current deployment or provisioning state, which only appears in the response.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="proximityPlacementGroupID")
    def proximity_placement_group_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID for Proximity Placement Group.
        """
        return pulumi.get(self, "proximity_placement_group_id")

    @property
    @pulumi.getter(name="scaleSetEvictionPolicy")
    def scale_set_eviction_policy(self) -> pulumi.Output[Optional[str]]:
        """
        ScaleSetEvictionPolicy to be used to specify eviction policy for Spot virtual machine scale set. Default to Delete.
        """
        return pulumi.get(self, "scale_set_eviction_policy")

    @property
    @pulumi.getter(name="scaleSetPriority")
    def scale_set_priority(self) -> pulumi.Output[Optional[str]]:
        """
        ScaleSetPriority to be used to specify virtual machine scale set priority. Default to regular.
        """
        return pulumi.get(self, "scale_set_priority")

    @property
    @pulumi.getter(name="spotMaxPrice")
    def spot_max_price(self) -> pulumi.Output[Optional[float]]:
        """
        SpotMaxPrice to be used to specify the maximum price you are willing to pay in US Dollars. Possible values are any decimal value greater than zero or -1 which indicates default price to be up-to on-demand.
        """
        return pulumi.get(self, "spot_max_price")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Agent pool tags to be persisted on the agent pool virtual machine scale set.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        AgentPoolType represents types of an agent pool
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="upgradeSettings")
    def upgrade_settings(self) -> pulumi.Output[Optional['outputs.AgentPoolUpgradeSettingsResponse']]:
        """
        Settings for upgrading the agentpool
        """
        return pulumi.get(self, "upgrade_settings")

    @property
    @pulumi.getter(name="vmSize")
    def vm_size(self) -> pulumi.Output[Optional[str]]:
        """
        Size of agent VMs.
        """
        return pulumi.get(self, "vm_size")

    @property
    @pulumi.getter(name="vnetSubnetID")
    def vnet_subnet_id(self) -> pulumi.Output[Optional[str]]:
        """
        VNet SubnetID specifies the VNet's subnet identifier for nodes and maybe pods
        """
        return pulumi.get(self, "vnet_subnet_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

