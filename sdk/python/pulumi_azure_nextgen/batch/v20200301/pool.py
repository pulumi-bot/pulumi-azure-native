# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Pool']


class Pool(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_name: Optional[pulumi.Input[str]] = None,
                 application_licenses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 application_packages: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationPackageReferenceArgs']]]]] = None,
                 certificates: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CertificateReferenceArgs']]]]] = None,
                 deployment_configuration: Optional[pulumi.Input[pulumi.InputType['DeploymentConfigurationArgs']]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 inter_node_communication: Optional[pulumi.Input[str]] = None,
                 max_tasks_per_node: Optional[pulumi.Input[int]] = None,
                 metadata: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MetadataItemArgs']]]]] = None,
                 mount_configuration: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MountConfigurationArgs']]]]] = None,
                 network_configuration: Optional[pulumi.Input[pulumi.InputType['NetworkConfigurationArgs']]] = None,
                 pool_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 scale_settings: Optional[pulumi.Input[pulumi.InputType['ScaleSettingsArgs']]] = None,
                 start_task: Optional[pulumi.Input[pulumi.InputType['StartTaskArgs']]] = None,
                 task_scheduling_policy: Optional[pulumi.Input[pulumi.InputType['TaskSchedulingPolicyArgs']]] = None,
                 user_accounts: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['UserAccountArgs']]]]] = None,
                 vm_size: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Contains information about a pool.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: The name of the Batch account.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] application_licenses: The list of application licenses must be a subset of available Batch service application licenses. If a license is requested which is not supported, pool creation will fail.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationPackageReferenceArgs']]]] application_packages: Changes to application package references affect all new compute nodes joining the pool, but do not affect compute nodes that are already in the pool until they are rebooted or reimaged. There is a maximum of 10 application package references on any given pool.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CertificateReferenceArgs']]]] certificates: For Windows compute nodes, the Batch service installs the certificates to the specified certificate store and location. For Linux compute nodes, the certificates are stored in a directory inside the task working directory and an environment variable AZ_BATCH_CERTIFICATES_DIR is supplied to the task to query for this location. For certificates with visibility of 'remoteUser', a 'certs' directory is created in the user's home directory (e.g., /home/{user-name}/certs) and certificates are placed in that directory.
        :param pulumi.Input[pulumi.InputType['DeploymentConfigurationArgs']] deployment_configuration: Using CloudServiceConfiguration specifies that the nodes should be creating using Azure Cloud Services (PaaS), while VirtualMachineConfiguration uses Azure Virtual Machines (IaaS).
        :param pulumi.Input[str] display_name: The display name need not be unique and can contain any Unicode characters up to a maximum length of 1024.
        :param pulumi.Input[str] inter_node_communication: This imposes restrictions on which nodes can be assigned to the pool. Enabling this value can reduce the chance of the requested number of nodes to be allocated in the pool. If not specified, this value defaults to 'Disabled'.
        :param pulumi.Input[int] max_tasks_per_node: The default value is 1. The maximum value is the smaller of 4 times the number of cores of the vmSize of the pool or 256.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MetadataItemArgs']]]] metadata: The Batch service does not assign any meaning to metadata; it is solely for the use of user code.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MountConfigurationArgs']]]] mount_configuration: This supports Azure Files, NFS, CIFS/SMB, and Blobfuse.
        :param pulumi.Input[pulumi.InputType['NetworkConfigurationArgs']] network_configuration: The network configuration for a pool.
        :param pulumi.Input[str] pool_name: The pool name. This must be unique within the account.
        :param pulumi.Input[str] resource_group_name: The name of the resource group that contains the Batch account.
        :param pulumi.Input[pulumi.InputType['ScaleSettingsArgs']] scale_settings: Defines the desired size of the pool. This can either be 'fixedScale' where the requested targetDedicatedNodes is specified, or 'autoScale' which defines a formula which is periodically reevaluated. If this property is not specified, the pool will have a fixed scale with 0 targetDedicatedNodes.
        :param pulumi.Input[pulumi.InputType['StartTaskArgs']] start_task: In an PATCH (update) operation, this property can be set to an empty object to remove the start task from the pool.
        :param pulumi.Input[pulumi.InputType['TaskSchedulingPolicyArgs']] task_scheduling_policy: If not specified, the default is spread.
        :param pulumi.Input[str] vm_size: For information about available sizes of virtual machines for Cloud Services pools (pools created with cloudServiceConfiguration), see Sizes for Cloud Services (https://azure.microsoft.com/documentation/articles/cloud-services-sizes-specs/). Batch supports all Cloud Services VM sizes except ExtraSmall. For information about available VM sizes for pools using images from the Virtual Machines Marketplace (pools created with virtualMachineConfiguration) see Sizes for Virtual Machines (Linux) (https://azure.microsoft.com/documentation/articles/virtual-machines-linux-sizes/) or Sizes for Virtual Machines (Windows) (https://azure.microsoft.com/documentation/articles/virtual-machines-windows-sizes/). Batch supports all Azure VM sizes except STANDARD_A0 and those with premium storage (STANDARD_GS, STANDARD_DS, and STANDARD_DSV2 series).
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

            if account_name is None:
                raise TypeError("Missing required property 'account_name'")
            __props__['account_name'] = account_name
            __props__['application_licenses'] = application_licenses
            __props__['application_packages'] = application_packages
            __props__['certificates'] = certificates
            __props__['deployment_configuration'] = deployment_configuration
            __props__['display_name'] = display_name
            __props__['inter_node_communication'] = inter_node_communication
            __props__['max_tasks_per_node'] = max_tasks_per_node
            __props__['metadata'] = metadata
            __props__['mount_configuration'] = mount_configuration
            __props__['network_configuration'] = network_configuration
            if pool_name is None:
                raise TypeError("Missing required property 'pool_name'")
            __props__['pool_name'] = pool_name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['scale_settings'] = scale_settings
            __props__['start_task'] = start_task
            __props__['task_scheduling_policy'] = task_scheduling_policy
            __props__['user_accounts'] = user_accounts
            __props__['vm_size'] = vm_size
            __props__['allocation_state'] = None
            __props__['allocation_state_transition_time'] = None
            __props__['auto_scale_run'] = None
            __props__['creation_time'] = None
            __props__['current_dedicated_nodes'] = None
            __props__['current_low_priority_nodes'] = None
            __props__['etag'] = None
            __props__['last_modified'] = None
            __props__['name'] = None
            __props__['provisioning_state'] = None
            __props__['provisioning_state_transition_time'] = None
            __props__['resize_operation_status'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:batch/latest:Pool"), pulumi.Alias(type_="azure-nextgen:batch/v20170901:Pool"), pulumi.Alias(type_="azure-nextgen:batch/v20181201:Pool"), pulumi.Alias(type_="azure-nextgen:batch/v20190401:Pool"), pulumi.Alias(type_="azure-nextgen:batch/v20190801:Pool"), pulumi.Alias(type_="azure-nextgen:batch/v20200501:Pool"), pulumi.Alias(type_="azure-nextgen:batch/v20200901:Pool")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Pool, __self__).__init__(
            'azure-nextgen:batch/v20200301:Pool',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Pool':
        """
        Get an existing Pool resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Pool(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allocationState")
    def allocation_state(self) -> pulumi.Output[str]:
        return pulumi.get(self, "allocation_state")

    @property
    @pulumi.getter(name="allocationStateTransitionTime")
    def allocation_state_transition_time(self) -> pulumi.Output[str]:
        return pulumi.get(self, "allocation_state_transition_time")

    @property
    @pulumi.getter(name="applicationLicenses")
    def application_licenses(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The list of application licenses must be a subset of available Batch service application licenses. If a license is requested which is not supported, pool creation will fail.
        """
        return pulumi.get(self, "application_licenses")

    @property
    @pulumi.getter(name="applicationPackages")
    def application_packages(self) -> pulumi.Output[Optional[Sequence['outputs.ApplicationPackageReferenceResponse']]]:
        """
        Changes to application package references affect all new compute nodes joining the pool, but do not affect compute nodes that are already in the pool until they are rebooted or reimaged. There is a maximum of 10 application package references on any given pool.
        """
        return pulumi.get(self, "application_packages")

    @property
    @pulumi.getter(name="autoScaleRun")
    def auto_scale_run(self) -> pulumi.Output['outputs.AutoScaleRunResponse']:
        """
        This property is set only if the pool automatically scales, i.e. autoScaleSettings are used.
        """
        return pulumi.get(self, "auto_scale_run")

    @property
    @pulumi.getter
    def certificates(self) -> pulumi.Output[Optional[Sequence['outputs.CertificateReferenceResponse']]]:
        """
        For Windows compute nodes, the Batch service installs the certificates to the specified certificate store and location. For Linux compute nodes, the certificates are stored in a directory inside the task working directory and an environment variable AZ_BATCH_CERTIFICATES_DIR is supplied to the task to query for this location. For certificates with visibility of 'remoteUser', a 'certs' directory is created in the user's home directory (e.g., /home/{user-name}/certs) and certificates are placed in that directory.
        """
        return pulumi.get(self, "certificates")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> pulumi.Output[str]:
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="currentDedicatedNodes")
    def current_dedicated_nodes(self) -> pulumi.Output[int]:
        return pulumi.get(self, "current_dedicated_nodes")

    @property
    @pulumi.getter(name="currentLowPriorityNodes")
    def current_low_priority_nodes(self) -> pulumi.Output[int]:
        return pulumi.get(self, "current_low_priority_nodes")

    @property
    @pulumi.getter(name="deploymentConfiguration")
    def deployment_configuration(self) -> pulumi.Output[Optional['outputs.DeploymentConfigurationResponse']]:
        """
        Using CloudServiceConfiguration specifies that the nodes should be creating using Azure Cloud Services (PaaS), while VirtualMachineConfiguration uses Azure Virtual Machines (IaaS).
        """
        return pulumi.get(self, "deployment_configuration")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[Optional[str]]:
        """
        The display name need not be unique and can contain any Unicode characters up to a maximum length of 1024.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        The ETag of the resource, used for concurrency statements.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="interNodeCommunication")
    def inter_node_communication(self) -> pulumi.Output[Optional[str]]:
        """
        This imposes restrictions on which nodes can be assigned to the pool. Enabling this value can reduce the chance of the requested number of nodes to be allocated in the pool. If not specified, this value defaults to 'Disabled'.
        """
        return pulumi.get(self, "inter_node_communication")

    @property
    @pulumi.getter(name="lastModified")
    def last_modified(self) -> pulumi.Output[str]:
        """
        This is the last time at which the pool level data, such as the targetDedicatedNodes or autoScaleSettings, changed. It does not factor in node-level changes such as a compute node changing state.
        """
        return pulumi.get(self, "last_modified")

    @property
    @pulumi.getter(name="maxTasksPerNode")
    def max_tasks_per_node(self) -> pulumi.Output[Optional[int]]:
        """
        The default value is 1. The maximum value is the smaller of 4 times the number of cores of the vmSize of the pool or 256.
        """
        return pulumi.get(self, "max_tasks_per_node")

    @property
    @pulumi.getter
    def metadata(self) -> pulumi.Output[Optional[Sequence['outputs.MetadataItemResponse']]]:
        """
        The Batch service does not assign any meaning to metadata; it is solely for the use of user code.
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter(name="mountConfiguration")
    def mount_configuration(self) -> pulumi.Output[Optional[Sequence['outputs.MountConfigurationResponse']]]:
        """
        This supports Azure Files, NFS, CIFS/SMB, and Blobfuse.
        """
        return pulumi.get(self, "mount_configuration")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkConfiguration")
    def network_configuration(self) -> pulumi.Output[Optional['outputs.NetworkConfigurationResponse']]:
        """
        The network configuration for a pool.
        """
        return pulumi.get(self, "network_configuration")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="provisioningStateTransitionTime")
    def provisioning_state_transition_time(self) -> pulumi.Output[str]:
        return pulumi.get(self, "provisioning_state_transition_time")

    @property
    @pulumi.getter(name="resizeOperationStatus")
    def resize_operation_status(self) -> pulumi.Output['outputs.ResizeOperationStatusResponse']:
        """
        Describes either the current operation (if the pool AllocationState is Resizing) or the previously completed operation (if the AllocationState is Steady).
        """
        return pulumi.get(self, "resize_operation_status")

    @property
    @pulumi.getter(name="scaleSettings")
    def scale_settings(self) -> pulumi.Output[Optional['outputs.ScaleSettingsResponse']]:
        """
        Defines the desired size of the pool. This can either be 'fixedScale' where the requested targetDedicatedNodes is specified, or 'autoScale' which defines a formula which is periodically reevaluated. If this property is not specified, the pool will have a fixed scale with 0 targetDedicatedNodes.
        """
        return pulumi.get(self, "scale_settings")

    @property
    @pulumi.getter(name="startTask")
    def start_task(self) -> pulumi.Output[Optional['outputs.StartTaskResponse']]:
        """
        In an PATCH (update) operation, this property can be set to an empty object to remove the start task from the pool.
        """
        return pulumi.get(self, "start_task")

    @property
    @pulumi.getter(name="taskSchedulingPolicy")
    def task_scheduling_policy(self) -> pulumi.Output[Optional['outputs.TaskSchedulingPolicyResponse']]:
        """
        If not specified, the default is spread.
        """
        return pulumi.get(self, "task_scheduling_policy")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the resource.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="userAccounts")
    def user_accounts(self) -> pulumi.Output[Optional[Sequence['outputs.UserAccountResponse']]]:
        return pulumi.get(self, "user_accounts")

    @property
    @pulumi.getter(name="vmSize")
    def vm_size(self) -> pulumi.Output[Optional[str]]:
        """
        For information about available sizes of virtual machines for Cloud Services pools (pools created with cloudServiceConfiguration), see Sizes for Cloud Services (https://azure.microsoft.com/documentation/articles/cloud-services-sizes-specs/). Batch supports all Cloud Services VM sizes except ExtraSmall. For information about available VM sizes for pools using images from the Virtual Machines Marketplace (pools created with virtualMachineConfiguration) see Sizes for Virtual Machines (Linux) (https://azure.microsoft.com/documentation/articles/virtual-machines-linux-sizes/) or Sizes for Virtual Machines (Windows) (https://azure.microsoft.com/documentation/articles/virtual-machines-windows-sizes/). Batch supports all Azure VM sizes except STANDARD_A0 and those with premium storage (STANDARD_GS, STANDARD_DS, and STANDARD_DSV2 series).
        """
        return pulumi.get(self, "vm_size")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

