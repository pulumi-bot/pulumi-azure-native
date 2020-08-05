# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class AgentPool(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    The name of the resource that is unique within a resource group. This name can be used to access the resource.
    """
    properties: pulumi.Output[dict]
    """
    Properties of an agent pool.
      * `availability_zones` (`list`) - (PREVIEW) Availability zones for nodes. Must use VirtualMachineScaleSets AgentPoolType.
      * `count` (`float`) - Number of agents (VMs) to host docker containers. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1. 
      * `enable_auto_scaling` (`bool`) - Whether to enable auto-scaler
      * `max_count` (`float`) - Maximum number of nodes for auto-scaling
      * `max_pods` (`float`) - Maximum number of pods that can run on a node.
      * `min_count` (`float`) - Minimum number of nodes for auto-scaling
      * `orchestrator_version` (`str`) - Version of orchestrator specified when creating the managed cluster.
      * `os_disk_size_gb` (`float`) - OS Disk Size in GB to be used to specify the disk size for every machine in this master/agent pool. If you specify 0, it will apply the default osDisk size according to the vmSize specified.
      * `os_type` (`str`) - OsType to be used to specify os type. Choose from Linux and Windows. Default to Linux.
      * `provisioning_state` (`str`) - The current deployment or provisioning state, which only appears in the response.
      * `type` (`str`) - AgentPoolType represents types of an agent pool
      * `vm_size` (`str`) - Size of agent VMs.
      * `vnet_subnet_id` (`str`) - VNet SubnetID specifies the VNet's subnet identifier.
    """
    type: pulumi.Output[str]
    """
    Resource type
    """
    def __init__(__self__, resource_name, opts=None, availability_zones=None, count=None, enable_auto_scaling=None, max_count=None, max_pods=None, min_count=None, name=None, orchestrator_version=None, os_disk_size_gb=None, os_type=None, resource_group_name=None, resource_name_=None, type=None, vm_size=None, vnet_subnet_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Agent Pool.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] availability_zones: (PREVIEW) Availability zones for nodes. Must use VirtualMachineScaleSets AgentPoolType.
        :param pulumi.Input[float] count: Number of agents (VMs) to host docker containers. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1. 
        :param pulumi.Input[bool] enable_auto_scaling: Whether to enable auto-scaler
        :param pulumi.Input[float] max_count: Maximum number of nodes for auto-scaling
        :param pulumi.Input[float] max_pods: Maximum number of pods that can run on a node.
        :param pulumi.Input[float] min_count: Minimum number of nodes for auto-scaling
        :param pulumi.Input[str] name: The name of the agent pool.
        :param pulumi.Input[str] orchestrator_version: Version of orchestrator specified when creating the managed cluster.
        :param pulumi.Input[float] os_disk_size_gb: OS Disk Size in GB to be used to specify the disk size for every machine in this master/agent pool. If you specify 0, it will apply the default osDisk size according to the vmSize specified.
        :param pulumi.Input[str] os_type: OsType to be used to specify os type. Choose from Linux and Windows. Default to Linux.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] resource_name_: The name of the managed cluster resource.
        :param pulumi.Input[str] type: AgentPoolType represents types of an agent pool
        :param pulumi.Input[str] vm_size: Size of agent VMs.
        :param pulumi.Input[str] vnet_subnet_id: VNet SubnetID specifies the VNet's subnet identifier.
        """
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

            __props__['availability_zones'] = availability_zones
            if count is None:
                raise TypeError("Missing required property 'count'")
            __props__['count'] = count
            __props__['enable_auto_scaling'] = enable_auto_scaling
            __props__['max_count'] = max_count
            __props__['max_pods'] = max_pods
            __props__['min_count'] = min_count
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['orchestrator_version'] = orchestrator_version
            __props__['os_disk_size_gb'] = os_disk_size_gb
            __props__['os_type'] = os_type
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if resource_name_ is None:
                raise TypeError("Missing required property 'resource_name_'")
            __props__['resource_name'] = resource_name_
            __props__['type'] = type
            if vm_size is None:
                raise TypeError("Missing required property 'vm_size'")
            __props__['vm_size'] = vm_size
            __props__['vnet_subnet_id'] = vnet_subnet_id
            __props__['properties'] = None
        super(AgentPool, __self__).__init__(
            'azurerm:containerservice/v20190401:AgentPool',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing AgentPool resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return AgentPool(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
