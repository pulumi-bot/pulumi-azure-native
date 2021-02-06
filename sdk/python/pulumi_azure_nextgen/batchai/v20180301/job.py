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

__all__ = ['Job']


class Job(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 caffe2_settings: Optional[pulumi.Input[pulumi.InputType['Caffe2SettingsArgs']]] = None,
                 caffe_settings: Optional[pulumi.Input[pulumi.InputType['CaffeSettingsArgs']]] = None,
                 chainer_settings: Optional[pulumi.Input[pulumi.InputType['ChainerSettingsArgs']]] = None,
                 cluster: Optional[pulumi.Input[pulumi.InputType['ResourceIdArgs']]] = None,
                 cntk_settings: Optional[pulumi.Input[pulumi.InputType['CNTKsettingsArgs']]] = None,
                 constraints: Optional[pulumi.Input[pulumi.InputType['JobBasePropertiesConstraintsArgs']]] = None,
                 container_settings: Optional[pulumi.Input[pulumi.InputType['ContainerSettingsArgs']]] = None,
                 custom_toolkit_settings: Optional[pulumi.Input[pulumi.InputType['CustomToolkitSettingsArgs']]] = None,
                 environment_variables: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EnvironmentVariableArgs']]]]] = None,
                 experiment_name: Optional[pulumi.Input[str]] = None,
                 input_directories: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InputDirectoryArgs']]]]] = None,
                 job_name: Optional[pulumi.Input[str]] = None,
                 job_preparation: Optional[pulumi.Input[pulumi.InputType['JobPreparationArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 mount_volumes: Optional[pulumi.Input[pulumi.InputType['MountVolumesArgs']]] = None,
                 node_count: Optional[pulumi.Input[int]] = None,
                 output_directories: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['OutputDirectoryArgs']]]]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 py_torch_settings: Optional[pulumi.Input[pulumi.InputType['PyTorchSettingsArgs']]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 secrets: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EnvironmentVariableWithSecretValueArgs']]]]] = None,
                 std_out_err_path_prefix: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tensor_flow_settings: Optional[pulumi.Input[pulumi.InputType['TensorFlowSettingsArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Contains information about the job.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['Caffe2SettingsArgs']] caffe2_settings: Specifies the settings for Caffe2 job.
        :param pulumi.Input[pulumi.InputType['CaffeSettingsArgs']] caffe_settings: Specifies the settings for Caffe job.
        :param pulumi.Input[pulumi.InputType['ChainerSettingsArgs']] chainer_settings: Specifies the settings for Chainer job.
        :param pulumi.Input[pulumi.InputType['ResourceIdArgs']] cluster: Represents a resource ID. For example, for a subnet, it is the resource URL for the subnet.
        :param pulumi.Input[pulumi.InputType['CNTKsettingsArgs']] cntk_settings: Specifies the settings for CNTK (aka Microsoft Cognitive Toolkit) job.
        :param pulumi.Input[pulumi.InputType['JobBasePropertiesConstraintsArgs']] constraints: Constraints associated with the Job.
        :param pulumi.Input[pulumi.InputType['ContainerSettingsArgs']] container_settings: If the container was downloaded as part of cluster setup then the same container image will be used. If not provided, the job will run on the VM.
        :param pulumi.Input[pulumi.InputType['CustomToolkitSettingsArgs']] custom_toolkit_settings: Specifies the settings for a custom tool kit job.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EnvironmentVariableArgs']]]] environment_variables: Batch AI will setup these additional environment variables for the job.
        :param pulumi.Input[str] experiment_name: Describe the experiment information of the job
        :param pulumi.Input[str] job_name: The name of the job within the specified resource group. Job names can only contain a combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
        :param pulumi.Input[pulumi.InputType['JobPreparationArgs']] job_preparation: The specified actions will run on all the nodes that are part of the job
        :param pulumi.Input[str] location: The region in which to create the job.
        :param pulumi.Input[pulumi.InputType['MountVolumesArgs']] mount_volumes: These volumes will be mounted before the job execution and will be unmounted after the job completion. The volumes will be mounted at location specified by $AZ_BATCHAI_JOB_MOUNT_ROOT environment variable.
        :param pulumi.Input[int] node_count: The job will be gang scheduled on that many compute nodes
        :param pulumi.Input[int] priority: Priority associated with the job. Priority values can range from -1000 to 1000, with -1000 being the lowest priority and 1000 being the highest priority. The default value is 0.
        :param pulumi.Input[pulumi.InputType['PyTorchSettingsArgs']] py_torch_settings: Specifies the settings for pyTorch job.
        :param pulumi.Input[str] resource_group_name: Name of the resource group to which the resource belongs.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EnvironmentVariableWithSecretValueArgs']]]] secrets: Batch AI will setup these additional environment variables for the job. Server will never report values of these variables back.
        :param pulumi.Input[str] std_out_err_path_prefix: The path where the Batch AI service will upload stdout and stderror of the job.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: The user specified tags associated with the job.
        :param pulumi.Input[pulumi.InputType['TensorFlowSettingsArgs']] tensor_flow_settings: Specifies the settings for TensorFlow job.
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

            __props__['caffe2_settings'] = caffe2_settings
            __props__['caffe_settings'] = caffe_settings
            __props__['chainer_settings'] = chainer_settings
            if cluster is None and not opts.urn:
                raise TypeError("Missing required property 'cluster'")
            __props__['cluster'] = cluster
            __props__['cntk_settings'] = cntk_settings
            __props__['constraints'] = constraints
            __props__['container_settings'] = container_settings
            __props__['custom_toolkit_settings'] = custom_toolkit_settings
            __props__['environment_variables'] = environment_variables
            __props__['experiment_name'] = experiment_name
            __props__['input_directories'] = input_directories
            if job_name is None and not opts.urn:
                raise TypeError("Missing required property 'job_name'")
            __props__['job_name'] = job_name
            __props__['job_preparation'] = job_preparation
            __props__['location'] = location
            __props__['mount_volumes'] = mount_volumes
            if node_count is None and not opts.urn:
                raise TypeError("Missing required property 'node_count'")
            __props__['node_count'] = node_count
            __props__['output_directories'] = output_directories
            if priority is None:
                priority = 0
            __props__['priority'] = priority
            __props__['py_torch_settings'] = py_torch_settings
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['secrets'] = secrets
            if std_out_err_path_prefix is None and not opts.urn:
                raise TypeError("Missing required property 'std_out_err_path_prefix'")
            __props__['std_out_err_path_prefix'] = std_out_err_path_prefix
            __props__['tags'] = tags
            __props__['tensor_flow_settings'] = tensor_flow_settings
            __props__['creation_time'] = None
            __props__['execution_info'] = None
            __props__['execution_state'] = None
            __props__['execution_state_transition_time'] = None
            __props__['job_output_directory_path_segment'] = None
            __props__['name'] = None
            __props__['provisioning_state'] = None
            __props__['provisioning_state_transition_time'] = None
            __props__['tool_type'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:batchai/v20170901preview:Job")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Job, __self__).__init__(
            'azure-nextgen:batchai/v20180301:Job',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Job':
        """
        Get an existing Job resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Job(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="caffeSettings")
    def caffe_settings(self) -> pulumi.Output[Optional['outputs.CaffeSettingsResponse']]:
        """
        Specifies the settings for Caffe job.
        """
        return pulumi.get(self, "caffe_settings")

    @property
    @pulumi.getter(name="chainerSettings")
    def chainer_settings(self) -> pulumi.Output[Optional['outputs.ChainerSettingsResponse']]:
        """
        Specifies the settings for Chainer job.
        """
        return pulumi.get(self, "chainer_settings")

    @property
    @pulumi.getter
    def cluster(self) -> pulumi.Output[Optional['outputs.ResourceIdResponse']]:
        """
        Represents a resource ID. For example, for a subnet, it is the resource URL for the subnet.
        """
        return pulumi.get(self, "cluster")

    @property
    @pulumi.getter(name="cntkSettings")
    def cntk_settings(self) -> pulumi.Output[Optional['outputs.CNTKsettingsResponse']]:
        """
        Specifies the settings for CNTK (aka Microsoft Cognitive Toolkit) job.
        """
        return pulumi.get(self, "cntk_settings")

    @property
    @pulumi.getter
    def constraints(self) -> pulumi.Output[Optional['outputs.JobPropertiesResponseConstraints']]:
        """
        Constraints associated with the Job.
        """
        return pulumi.get(self, "constraints")

    @property
    @pulumi.getter(name="containerSettings")
    def container_settings(self) -> pulumi.Output[Optional['outputs.ContainerSettingsResponse']]:
        """
        If the container was downloaded as part of cluster setup then the same container image will be used. If not provided, the job will run on the VM.
        """
        return pulumi.get(self, "container_settings")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> pulumi.Output[str]:
        """
        The creation time of the job.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="customToolkitSettings")
    def custom_toolkit_settings(self) -> pulumi.Output[Optional['outputs.CustomToolkitSettingsResponse']]:
        """
        Specifies the settings for a custom tool kit job.
        """
        return pulumi.get(self, "custom_toolkit_settings")

    @property
    @pulumi.getter(name="environmentVariables")
    def environment_variables(self) -> pulumi.Output[Optional[Sequence['outputs.EnvironmentVariableResponse']]]:
        """
        Batch AI will setup these additional environment variables for the job.
        """
        return pulumi.get(self, "environment_variables")

    @property
    @pulumi.getter(name="executionInfo")
    def execution_info(self) -> pulumi.Output[Optional['outputs.JobPropertiesResponseExecutionInfo']]:
        """
        Contains information about the execution of a job in the Azure Batch service.
        """
        return pulumi.get(self, "execution_info")

    @property
    @pulumi.getter(name="executionState")
    def execution_state(self) -> pulumi.Output[Optional[str]]:
        """
        The current state of the job. Possible values are: queued - The job is queued and able to run. A job enters this state when it is created, or when it is awaiting a retry after a failed run. running - The job is running on a compute cluster. This includes job-level preparation such as downloading resource files or set up container specified on the job - it does not necessarily mean that the job command line has started executing. terminating - The job is terminated by the user, the terminate operation is in progress. succeeded - The job has completed running successfully and exited with exit code 0. failed - The job has finished unsuccessfully (failed with a non-zero exit code) and has exhausted its retry limit. A job is also marked as failed if an error occurred launching the job.
        """
        return pulumi.get(self, "execution_state")

    @property
    @pulumi.getter(name="executionStateTransitionTime")
    def execution_state_transition_time(self) -> pulumi.Output[str]:
        """
        The time at which the job entered its current execution state.
        """
        return pulumi.get(self, "execution_state_transition_time")

    @property
    @pulumi.getter(name="experimentName")
    def experiment_name(self) -> pulumi.Output[Optional[str]]:
        """
        Describe the experiment information of the job
        """
        return pulumi.get(self, "experiment_name")

    @property
    @pulumi.getter(name="inputDirectories")
    def input_directories(self) -> pulumi.Output[Optional[Sequence['outputs.InputDirectoryResponse']]]:
        return pulumi.get(self, "input_directories")

    @property
    @pulumi.getter(name="jobOutputDirectoryPathSegment")
    def job_output_directory_path_segment(self) -> pulumi.Output[Optional[str]]:
        """
        Batch AI creates job's output directories under an unique path to avoid conflicts between jobs. This value contains a path segment generated by Batch AI to make the path unique and can be used to find the output directory on the node or mounted filesystem.
        """
        return pulumi.get(self, "job_output_directory_path_segment")

    @property
    @pulumi.getter(name="jobPreparation")
    def job_preparation(self) -> pulumi.Output[Optional['outputs.JobPreparationResponse']]:
        """
        The specified actions will run on all the nodes that are part of the job
        """
        return pulumi.get(self, "job_preparation")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The location of the resource
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="mountVolumes")
    def mount_volumes(self) -> pulumi.Output[Optional['outputs.MountVolumesResponse']]:
        """
        These volumes will be mounted before the job execution and will be unmounted after the job completion. The volumes will be mounted at location specified by $AZ_BATCHAI_JOB_MOUNT_ROOT environment variable.
        """
        return pulumi.get(self, "mount_volumes")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nodeCount")
    def node_count(self) -> pulumi.Output[Optional[int]]:
        """
        The job will be gang scheduled on that many compute nodes
        """
        return pulumi.get(self, "node_count")

    @property
    @pulumi.getter(name="outputDirectories")
    def output_directories(self) -> pulumi.Output[Optional[Sequence['outputs.OutputDirectoryResponse']]]:
        return pulumi.get(self, "output_directories")

    @property
    @pulumi.getter
    def priority(self) -> pulumi.Output[Optional[int]]:
        """
        Priority associated with the job. Priority values can range from -1000 to 1000, with -1000 being the lowest priority and 1000 being the highest priority. The default value is 0.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        The provisioned state of the Batch AI job
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="provisioningStateTransitionTime")
    def provisioning_state_transition_time(self) -> pulumi.Output[str]:
        """
        The time at which the job entered its current provisioning state.
        """
        return pulumi.get(self, "provisioning_state_transition_time")

    @property
    @pulumi.getter(name="pyTorchSettings")
    def py_torch_settings(self) -> pulumi.Output[Optional['outputs.PyTorchSettingsResponse']]:
        """
        Specifies the settings for pyTorch job.
        """
        return pulumi.get(self, "py_torch_settings")

    @property
    @pulumi.getter
    def secrets(self) -> pulumi.Output[Optional[Sequence['outputs.EnvironmentVariableWithSecretValueResponse']]]:
        """
        Batch AI will setup these additional environment variables for the job. Server will never report values of these variables back.
        """
        return pulumi.get(self, "secrets")

    @property
    @pulumi.getter(name="stdOutErrPathPrefix")
    def std_out_err_path_prefix(self) -> pulumi.Output[Optional[str]]:
        """
        The path where the Batch AI service will upload stdout and stderror of the job.
        """
        return pulumi.get(self, "std_out_err_path_prefix")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Mapping[str, str]]:
        """
        The tags of the resource
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tensorFlowSettings")
    def tensor_flow_settings(self) -> pulumi.Output[Optional['outputs.TensorFlowSettingsResponse']]:
        """
        Specifies the settings for TensorFlow job.
        """
        return pulumi.get(self, "tensor_flow_settings")

    @property
    @pulumi.getter(name="toolType")
    def tool_type(self) -> pulumi.Output[Optional[str]]:
        """
        Possible values are: cntk, tensorflow, caffe, caffe2, chainer, pytorch, custom.
        """
        return pulumi.get(self, "tool_type")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the resource
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

