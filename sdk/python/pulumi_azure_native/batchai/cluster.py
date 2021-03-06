# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['Cluster']


class Cluster(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 node_setup: Optional[pulumi.Input[pulumi.InputType['NodeSetupArgs']]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 scale_settings: Optional[pulumi.Input[pulumi.InputType['ScaleSettingsArgs']]] = None,
                 subnet: Optional[pulumi.Input[pulumi.InputType['ResourceIdArgs']]] = None,
                 user_account_settings: Optional[pulumi.Input[pulumi.InputType['UserAccountSettingsArgs']]] = None,
                 virtual_machine_configuration: Optional[pulumi.Input[pulumi.InputType['VirtualMachineConfigurationArgs']]] = None,
                 vm_priority: Optional[pulumi.Input['VmPriority']] = None,
                 vm_size: Optional[pulumi.Input[str]] = None,
                 workspace_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Information about a Cluster.
        API Version: 2018-05-01.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_name: The name of the cluster within the specified resource group. Cluster names can only contain a combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
        :param pulumi.Input[pulumi.InputType['NodeSetupArgs']] node_setup: Setup to be performed on each compute node in the cluster.
        :param pulumi.Input[str] resource_group_name: Name of the resource group to which the resource belongs.
        :param pulumi.Input[pulumi.InputType['ScaleSettingsArgs']] scale_settings: Scale settings for the cluster. Batch AI service supports manual and auto scale clusters.
        :param pulumi.Input[pulumi.InputType['ResourceIdArgs']] subnet: Existing virtual network subnet to put the cluster nodes in. Note, if a File Server mount configured in node setup, the File Server's subnet will be used automatically.
        :param pulumi.Input[pulumi.InputType['UserAccountSettingsArgs']] user_account_settings: Settings for an administrator user account that will be created on each compute node in the cluster.
        :param pulumi.Input[pulumi.InputType['VirtualMachineConfigurationArgs']] virtual_machine_configuration: OS image configuration for cluster nodes. All nodes in a cluster have the same OS image.
        :param pulumi.Input['VmPriority'] vm_priority: VM priority. Allowed values are: dedicated (default) and lowpriority.
        :param pulumi.Input[str] vm_size: The size of the virtual machines in the cluster. All nodes in a cluster have the same VM size. For information about available VM sizes for clusters using images from the Virtual Machines Marketplace see Sizes for Virtual Machines (Linux). Batch AI service supports all Azure VM sizes except STANDARD_A0 and those with premium storage (STANDARD_GS, STANDARD_DS, and STANDARD_DSV2 series).
        :param pulumi.Input[str] workspace_name: The name of the workspace. Workspace names can only contain a combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
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

            __props__['cluster_name'] = cluster_name
            __props__['node_setup'] = node_setup
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['scale_settings'] = scale_settings
            __props__['subnet'] = subnet
            if user_account_settings is None and not opts.urn:
                raise TypeError("Missing required property 'user_account_settings'")
            __props__['user_account_settings'] = user_account_settings
            __props__['virtual_machine_configuration'] = virtual_machine_configuration
            if vm_priority is None:
                vm_priority = 'dedicated'
            __props__['vm_priority'] = vm_priority
            if vm_size is None and not opts.urn:
                raise TypeError("Missing required property 'vm_size'")
            __props__['vm_size'] = vm_size
            if workspace_name is None and not opts.urn:
                raise TypeError("Missing required property 'workspace_name'")
            __props__['workspace_name'] = workspace_name
            __props__['allocation_state'] = None
            __props__['allocation_state_transition_time'] = None
            __props__['creation_time'] = None
            __props__['current_node_count'] = None
            __props__['errors'] = None
            __props__['name'] = None
            __props__['node_state_counts'] = None
            __props__['provisioning_state'] = None
            __props__['provisioning_state_transition_time'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:batchai:Cluster"), pulumi.Alias(type_="azure-native:batchai/latest:Cluster"), pulumi.Alias(type_="azure-nextgen:batchai/latest:Cluster"), pulumi.Alias(type_="azure-native:batchai/v20180501:Cluster"), pulumi.Alias(type_="azure-nextgen:batchai/v20180501:Cluster")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Cluster, __self__).__init__(
            'azure-native:batchai:Cluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Cluster':
        """
        Get an existing Cluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["allocation_state"] = None
        __props__["allocation_state_transition_time"] = None
        __props__["creation_time"] = None
        __props__["current_node_count"] = None
        __props__["errors"] = None
        __props__["name"] = None
        __props__["node_setup"] = None
        __props__["node_state_counts"] = None
        __props__["provisioning_state"] = None
        __props__["provisioning_state_transition_time"] = None
        __props__["scale_settings"] = None
        __props__["subnet"] = None
        __props__["type"] = None
        __props__["user_account_settings"] = None
        __props__["virtual_machine_configuration"] = None
        __props__["vm_priority"] = None
        __props__["vm_size"] = None
        return Cluster(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allocationState")
    def allocation_state(self) -> pulumi.Output[str]:
        """
        Allocation state of the cluster. Possible values are: steady - Indicates that the cluster is not resizing. There are no changes to the number of compute nodes in the cluster in progress. A cluster enters this state when it is created and when no operations are being performed on the cluster to change the number of compute nodes. resizing - Indicates that the cluster is resizing; that is, compute nodes are being added to or removed from the cluster.
        """
        return pulumi.get(self, "allocation_state")

    @property
    @pulumi.getter(name="allocationStateTransitionTime")
    def allocation_state_transition_time(self) -> pulumi.Output[str]:
        """
        The time at which the cluster entered its current allocation state.
        """
        return pulumi.get(self, "allocation_state_transition_time")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> pulumi.Output[str]:
        """
        The time when the cluster was created.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="currentNodeCount")
    def current_node_count(self) -> pulumi.Output[int]:
        """
        The number of compute nodes currently assigned to the cluster.
        """
        return pulumi.get(self, "current_node_count")

    @property
    @pulumi.getter
    def errors(self) -> pulumi.Output[Sequence['outputs.BatchAIErrorResponse']]:
        """
        Collection of errors encountered by various compute nodes during node setup.
        """
        return pulumi.get(self, "errors")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nodeSetup")
    def node_setup(self) -> pulumi.Output[Optional['outputs.NodeSetupResponse']]:
        """
        Setup (mount file systems, performance counters settings and custom setup task) to be performed on each compute node in the cluster.
        """
        return pulumi.get(self, "node_setup")

    @property
    @pulumi.getter(name="nodeStateCounts")
    def node_state_counts(self) -> pulumi.Output['outputs.NodeStateCountsResponse']:
        """
        Counts of various node states on the cluster.
        """
        return pulumi.get(self, "node_state_counts")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        Provisioning state of the cluster. Possible value are: creating - Specifies that the cluster is being created. succeeded - Specifies that the cluster has been created successfully. failed - Specifies that the cluster creation has failed. deleting - Specifies that the cluster is being deleted.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="provisioningStateTransitionTime")
    def provisioning_state_transition_time(self) -> pulumi.Output[str]:
        """
        Time when the provisioning state was changed.
        """
        return pulumi.get(self, "provisioning_state_transition_time")

    @property
    @pulumi.getter(name="scaleSettings")
    def scale_settings(self) -> pulumi.Output[Optional['outputs.ScaleSettingsResponse']]:
        """
        Scale settings of the cluster.
        """
        return pulumi.get(self, "scale_settings")

    @property
    @pulumi.getter
    def subnet(self) -> pulumi.Output[Optional['outputs.ResourceIdResponse']]:
        """
        Virtual network subnet resource ID the cluster nodes belong to.
        """
        return pulumi.get(self, "subnet")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the resource.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="userAccountSettings")
    def user_account_settings(self) -> pulumi.Output[Optional['outputs.UserAccountSettingsResponse']]:
        """
        Administrator user account settings which can be used to SSH to compute nodes.
        """
        return pulumi.get(self, "user_account_settings")

    @property
    @pulumi.getter(name="virtualMachineConfiguration")
    def virtual_machine_configuration(self) -> pulumi.Output[Optional['outputs.VirtualMachineConfigurationResponse']]:
        """
        Virtual machine configuration (OS image) of the compute nodes. All nodes in a cluster have the same OS image configuration.
        """
        return pulumi.get(self, "virtual_machine_configuration")

    @property
    @pulumi.getter(name="vmPriority")
    def vm_priority(self) -> pulumi.Output[Optional[str]]:
        """
        VM priority of cluster nodes.
        """
        return pulumi.get(self, "vm_priority")

    @property
    @pulumi.getter(name="vmSize")
    def vm_size(self) -> pulumi.Output[Optional[str]]:
        """
        The size of the virtual machines in the cluster. All nodes in a cluster have the same VM size.
        """
        return pulumi.get(self, "vm_size")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

