# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class Job(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    The location of the resource
    """
    name: pulumi.Output[str]
    """
    The name of the resource
    """
    properties: pulumi.Output[dict]
    """
    The properties associated with the job.
      * `caffe_settings` (`dict`) - Specifies the settings for Caffe job.
        * `command_line_args` (`str`)
        * `config_file_path` (`str`) - This property cannot be specified if pythonScriptFilePath is specified.
        * `process_count` (`float`) - The default value for this property is equal to nodeCount property
        * `python_interpreter_path` (`str`) - This property can be specified only if the pythonScriptFilePath is specified.
        * `python_script_file_path` (`str`) - This property cannot be specified if configFilePath is specified.

      * `chainer_settings` (`dict`) - Specifies the settings for Chainer job.
        * `command_line_args` (`str`)
        * `process_count` (`float`) - The default value for this property is equal to nodeCount property
        * `python_interpreter_path` (`str`)
        * `python_script_file_path` (`str`)

      * `cluster` (`dict`) - Represents a resource ID. For example, for a subnet, it is the resource URL for the subnet.
        * `id` (`str`) - The ID of the resource

      * `cntk_settings` (`dict`) - Specifies the settings for CNTK (aka Microsoft Cognitive Toolkit) job.
        * `command_line_args` (`str`)
        * `config_file_path` (`str`) - This property can be specified only if the languageType is 'BrainScript'.
        * `language_type` (`str`) - Valid values are 'BrainScript' or 'Python'.
        * `process_count` (`float`) - The default value for this property is equal to nodeCount property
        * `python_interpreter_path` (`str`) - This property can be specified only if the languageType is 'Python'.
        * `python_script_file_path` (`str`) - This property can be specified only if the languageType is 'Python'.

      * `constraints` (`dict`) - Constraints associated with the Job.
        * `max_wall_clock_time` (`str`) - Default Value = 1 week.

      * `container_settings` (`dict`) - If the container was downloaded as part of cluster setup then the same container image will be used. If not provided, the job will run on the VM.
        * `image_source_registry` (`dict`) - Details of the container image such as name, URL and credentials.
          * `credentials` (`dict`) - Credentials to access a container image in a private repository.
            * `password` (`str`) - One of password or passwordSecretReference must be specified.
            * `password_secret_reference` (`dict`) - Users can store their secrets in Azure KeyVault and pass it to the Batch AI Service to integrate with KeyVault. One of password or passwordSecretReference must be specified.
              * `secret_url` (`str`)
              * `source_vault` (`dict`) - Represents a resource ID. For example, for a subnet, it is the resource URL for the subnet.

            * `username` (`str`)

          * `image` (`str`)
          * `server_url` (`str`)

      * `creation_time` (`str`) - The creation time of the job.
      * `custom_toolkit_settings` (`dict`) - Specifies the settings for a custom tool kit job.
        * `command_line` (`str`)

      * `environment_variables` (`list`) - Batch AI will setup these additional environment variables for the job.
        * `name` (`str`)
        * `value` (`str`)

      * `execution_info` (`dict`) - Contains information about the execution of a job in the Azure Batch service.
      * `execution_state` (`str`) - The current state of the job. Possible values are: queued - The job is queued and able to run. A job enters this state when it is created, or when it is awaiting a retry after a failed run. running - The job is running on a compute cluster. This includes job-level preparation such as downloading resource files or set up container specified on the job - it does not necessarily mean that the job command line has started executing. terminating - The job is terminated by the user, the terminate operation is in progress. succeeded - The job has completed running successfully and exited with exit code 0. failed - The job has finished unsuccessfully (failed with a non-zero exit code) and has exhausted its retry limit. A job is also marked as failed if an error occurred launching the job.
      * `execution_state_transition_time` (`str`) - The time at which the job entered its current execution state.
      * `experiment_name` (`str`) - Describe the experiment information of the job
      * `input_directories` (`list`)
        * `id` (`str`) - The path of the input directory will be available as a value of an environment variable with AZ_BATCHAI_INPUT_<id> name, where <id> is the value of id attribute.
        * `path` (`str`)

      * `job_output_directory_path_segment` (`str`) - Batch AI creates job's output directories under an unique path to avoid conflicts between jobs. This value contains a path segment generated by Batch AI to make the path unique and can be used to find the output directory on the node or mounted filesystem.
      * `job_preparation` (`dict`) - The specified actions will run on all the nodes that are part of the job
        * `command_line` (`str`) - If containerSettings is specified on the job, this commandLine will be executed in the same container as job. Otherwise it will be executed on the node.

      * `mount_volumes` (`dict`) - These volumes will be mounted before the job execution and will be unmounted after the job completion. The volumes will be mounted at location specified by $AZ_BATCHAI_JOB_MOUNT_ROOT environment variable.
        * `azure_blob_file_systems` (`list`) - References to Azure Blob FUSE that are to be mounted to the cluster nodes.
          * `account_name` (`str`)
          * `container_name` (`str`)
          * `credentials` (`dict`) - Credentials to access Azure File Share.
            * `account_key` (`str`) - One of accountKey or accountKeySecretReference must be specified.
            * `account_key_secret_reference` (`dict`) - Users can store their secrets in Azure KeyVault and pass it to the Batch AI Service to integrate with KeyVault. One of accountKey or accountKeySecretReference must be specified.

          * `mount_options` (`str`)
          * `relative_mount_path` (`str`) - Note that all cluster level blob file systems will be mounted under $AZ_BATCHAI_MOUNT_ROOT location and all job level blob file systems will be mounted under $AZ_BATCHAI_JOB_MOUNT_ROOT.

        * `azure_file_shares` (`list`) - References to Azure File Shares that are to be mounted to the cluster nodes.
          * `account_name` (`str`)
          * `azure_file_url` (`str`)
          * `credentials` (`dict`) - Credentials to access Azure File Share.
          * `directory_mode` (`str`) - Default value is 0777. Valid only if OS is linux.
          * `file_mode` (`str`) - Default value is 0777. Valid only if OS is linux.
          * `relative_mount_path` (`str`) - Note that all cluster level file shares will be mounted under $AZ_BATCHAI_MOUNT_ROOT location and all job level file shares will be mounted under $AZ_BATCHAI_JOB_MOUNT_ROOT.

        * `file_servers` (`list`)
          * `file_server` (`dict`) - Represents a resource ID. For example, for a subnet, it is the resource URL for the subnet.
          * `mount_options` (`str`)
          * `relative_mount_path` (`str`) - Note that all cluster level file servers will be mounted under $AZ_BATCHAI_MOUNT_ROOT location and job level file servers will be mounted under $AZ_BATCHAI_JOB_MOUNT_ROOT.
          * `source_directory` (`str`) - If this property is not specified, the entire File Server will be mounted.

        * `unmanaged_file_systems` (`list`)
          * `mount_command` (`str`)
          * `relative_mount_path` (`str`) - Note that all cluster level unmanaged file system will be mounted under $AZ_BATCHAI_MOUNT_ROOT location and job level unmanaged file system will be mounted under $AZ_BATCHAI_JOB_MOUNT_ROOT.

      * `node_count` (`float`) - The job will be gang scheduled on that many compute nodes
      * `output_directories` (`list`)
        * `create_new` (`bool`) - Default is true. If false, then the directory is not created and can be any directory path that the user specifies.
        * `id` (`str`) - The path of the output directory will be available as a value of an environment variable with AZ_BATCHAI_OUTPUT_<id> name, where <id> is the value of id attribute.
        * `path_prefix` (`str`) - NOTE: This is an absolute path to prefix. E.g. $AZ_BATCHAI_MOUNT_ROOT/MyNFS/MyLogs. You can find the full path to the output directory by combining pathPrefix, jobOutputDirectoryPathSegment (reported by get job) and pathSuffix.
        * `path_suffix` (`str`) - The suffix path where the output directory will be created. E.g. models. You can find the full path to the output directory by combining pathPrefix, jobOutputDirectoryPathSegment (reported by get job) and pathSuffix.
        * `type` (`str`) - Default value is Custom. The possible values are Model, Logs, Summary, and Custom. Users can use multiple enums for a single directory. Eg. outPutType='Model,Logs, Summary'

      * `priority` (`float`) - Priority associated with the job. Priority values can range from -1000 to 1000, with -1000 being the lowest priority and 1000 being the highest priority. The default value is 0.
      * `provisioning_state` (`str`) - The provisioned state of the Batch AI job
      * `provisioning_state_transition_time` (`str`) - The time at which the job entered its current provisioning state.
      * `py_torch_settings` (`dict`) - Specifies the settings for pyTorch job.
        * `command_line_args` (`str`)
        * `communication_backend` (`str`) - Valid values are 'TCP', 'Gloo' or 'MPI'. Not required for non-distributed jobs.
        * `process_count` (`float`) - The default value for this property is equal to nodeCount property.
        * `python_interpreter_path` (`str`)
        * `python_script_file_path` (`str`)

      * `secrets` (`list`) - Batch AI will setup these additional environment variables for the job. Server will never report values of these variables back.
        * `name` (`str`)
        * `value` (`str`)
        * `value_secret_reference` (`dict`) - Specifies KeyVault Store and Secret which contains the value for the environment variable. One of value or valueSecretReference must be provided.

      * `std_out_err_path_prefix` (`str`) - The path where the Batch AI service will upload stdout and stderror of the job.
      * `tensor_flow_settings` (`dict`) - Specifies the settings for TensorFlow job.
        * `master_command_line_args` (`str`)
        * `parameter_server_command_line_args` (`str`) - This property is optional for single machine training.
        * `parameter_server_count` (`float`) - If specified, the value must be less than or equal to nodeCount. If not specified, the default value is equal to 1 for distributed TensorFlow training (This property is not applicable for single machine training). This property can be specified only for distributed TensorFlow training.
        * `python_interpreter_path` (`str`)
        * `python_script_file_path` (`str`)
        * `worker_command_line_args` (`str`) - This property is optional for single machine training.
        * `worker_count` (`float`) - If specified, the value must be less than or equal to (nodeCount * numberOfGPUs per VM). If not specified, the default value is equal to nodeCount. This property can be specified only for distributed TensorFlow training

      * `tool_type` (`str`) - Possible values are: cntk, tensorflow, caffe, caffe2, chainer, pytorch, custom.
    """
    tags: pulumi.Output[dict]
    """
    The tags of the resource
    """
    type: pulumi.Output[str]
    """
    The type of the resource
    """
    def __init__(__self__, resource_name, opts=None, location=None, name=None, properties=None, resource_group_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Contains information about the job.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: The region in which to create the job.
        :param pulumi.Input[str] name: The name of the job within the specified resource group. Job names can only contain a combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
        :param pulumi.Input[dict] properties: The properties of the Job.
        :param pulumi.Input[str] resource_group_name: Name of the resource group to which the resource belongs.
        :param pulumi.Input[dict] tags: The user specified tags associated with the job.

        The **properties** object supports the following:

          * `caffe2_settings` (`pulumi.Input[dict]`) - Specifies the settings for Caffe2 job.
            * `command_line_args` (`pulumi.Input[str]`)
            * `python_interpreter_path` (`pulumi.Input[str]`)
            * `python_script_file_path` (`pulumi.Input[str]`)

          * `caffe_settings` (`pulumi.Input[dict]`) - Specifies the settings for Caffe job.
            * `command_line_args` (`pulumi.Input[str]`)
            * `config_file_path` (`pulumi.Input[str]`) - This property cannot be specified if pythonScriptFilePath is specified.
            * `process_count` (`pulumi.Input[float]`) - The default value for this property is equal to nodeCount property
            * `python_interpreter_path` (`pulumi.Input[str]`) - This property can be specified only if the pythonScriptFilePath is specified.
            * `python_script_file_path` (`pulumi.Input[str]`) - This property cannot be specified if configFilePath is specified.

          * `chainer_settings` (`pulumi.Input[dict]`) - Specifies the settings for Chainer job.
            * `command_line_args` (`pulumi.Input[str]`)
            * `process_count` (`pulumi.Input[float]`) - The default value for this property is equal to nodeCount property
            * `python_interpreter_path` (`pulumi.Input[str]`)
            * `python_script_file_path` (`pulumi.Input[str]`)

          * `cluster` (`pulumi.Input[dict]`) - Represents a resource ID. For example, for a subnet, it is the resource URL for the subnet.
            * `id` (`pulumi.Input[str]`) - The ID of the resource

          * `cntk_settings` (`pulumi.Input[dict]`) - Specifies the settings for CNTK (aka Microsoft Cognitive Toolkit) job.
            * `command_line_args` (`pulumi.Input[str]`)
            * `config_file_path` (`pulumi.Input[str]`) - This property can be specified only if the languageType is 'BrainScript'.
            * `language_type` (`pulumi.Input[str]`) - Valid values are 'BrainScript' or 'Python'.
            * `process_count` (`pulumi.Input[float]`) - The default value for this property is equal to nodeCount property
            * `python_interpreter_path` (`pulumi.Input[str]`) - This property can be specified only if the languageType is 'Python'.
            * `python_script_file_path` (`pulumi.Input[str]`) - This property can be specified only if the languageType is 'Python'.

          * `constraints` (`pulumi.Input[dict]`) - Constraints associated with the Job.
            * `max_wall_clock_time` (`pulumi.Input[str]`) - Default Value = 1 week.

          * `container_settings` (`pulumi.Input[dict]`) - If the container was downloaded as part of cluster setup then the same container image will be used. If not provided, the job will run on the VM.
            * `image_source_registry` (`pulumi.Input[dict]`) - Details of the container image such as name, URL and credentials.
              * `credentials` (`pulumi.Input[dict]`) - Credentials to access a container image in a private repository.
                * `password` (`pulumi.Input[str]`) - One of password or passwordSecretReference must be specified.
                * `password_secret_reference` (`pulumi.Input[dict]`) - Users can store their secrets in Azure KeyVault and pass it to the Batch AI Service to integrate with KeyVault. One of password or passwordSecretReference must be specified.
                  * `secret_url` (`pulumi.Input[str]`)
                  * `source_vault` (`pulumi.Input[dict]`) - Represents a resource ID. For example, for a subnet, it is the resource URL for the subnet.

                * `username` (`pulumi.Input[str]`)

              * `image` (`pulumi.Input[str]`)
              * `server_url` (`pulumi.Input[str]`)

          * `custom_toolkit_settings` (`pulumi.Input[dict]`) - Specifies the settings for a custom tool kit job.
            * `command_line` (`pulumi.Input[str]`)

          * `environment_variables` (`pulumi.Input[list]`) - Batch AI will setup these additional environment variables for the job.
            * `name` (`pulumi.Input[str]`)
            * `value` (`pulumi.Input[str]`)

          * `experiment_name` (`pulumi.Input[str]`) - Describe the experiment information of the job
          * `input_directories` (`pulumi.Input[list]`)
            * `id` (`pulumi.Input[str]`) - The path of the input directory will be available as a value of an environment variable with AZ_BATCHAI_INPUT_<id> name, where <id> is the value of id attribute.
            * `path` (`pulumi.Input[str]`)

          * `job_preparation` (`pulumi.Input[dict]`) - The specified actions will run on all the nodes that are part of the job
            * `command_line` (`pulumi.Input[str]`) - If containerSettings is specified on the job, this commandLine will be executed in the same container as job. Otherwise it will be executed on the node.

          * `mount_volumes` (`pulumi.Input[dict]`) - These volumes will be mounted before the job execution and will be unmounted after the job completion. The volumes will be mounted at location specified by $AZ_BATCHAI_JOB_MOUNT_ROOT environment variable.
            * `azure_blob_file_systems` (`pulumi.Input[list]`) - References to Azure Blob FUSE that are to be mounted to the cluster nodes.
              * `account_name` (`pulumi.Input[str]`)
              * `container_name` (`pulumi.Input[str]`)
              * `credentials` (`pulumi.Input[dict]`) - Credentials to access Azure File Share.
                * `account_key` (`pulumi.Input[str]`) - One of accountKey or accountKeySecretReference must be specified.
                * `account_key_secret_reference` (`pulumi.Input[dict]`) - Users can store their secrets in Azure KeyVault and pass it to the Batch AI Service to integrate with KeyVault. One of accountKey or accountKeySecretReference must be specified.

              * `mount_options` (`pulumi.Input[str]`)
              * `relative_mount_path` (`pulumi.Input[str]`) - Note that all cluster level blob file systems will be mounted under $AZ_BATCHAI_MOUNT_ROOT location and all job level blob file systems will be mounted under $AZ_BATCHAI_JOB_MOUNT_ROOT.

            * `azure_file_shares` (`pulumi.Input[list]`) - References to Azure File Shares that are to be mounted to the cluster nodes.
              * `account_name` (`pulumi.Input[str]`)
              * `azure_file_url` (`pulumi.Input[str]`)
              * `credentials` (`pulumi.Input[dict]`) - Credentials to access Azure File Share.
              * `directory_mode` (`pulumi.Input[str]`) - Default value is 0777. Valid only if OS is linux.
              * `file_mode` (`pulumi.Input[str]`) - Default value is 0777. Valid only if OS is linux.
              * `relative_mount_path` (`pulumi.Input[str]`) - Note that all cluster level file shares will be mounted under $AZ_BATCHAI_MOUNT_ROOT location and all job level file shares will be mounted under $AZ_BATCHAI_JOB_MOUNT_ROOT.

            * `file_servers` (`pulumi.Input[list]`)
              * `file_server` (`pulumi.Input[dict]`) - Represents a resource ID. For example, for a subnet, it is the resource URL for the subnet.
              * `mount_options` (`pulumi.Input[str]`)
              * `relative_mount_path` (`pulumi.Input[str]`) - Note that all cluster level file servers will be mounted under $AZ_BATCHAI_MOUNT_ROOT location and job level file servers will be mounted under $AZ_BATCHAI_JOB_MOUNT_ROOT.
              * `source_directory` (`pulumi.Input[str]`) - If this property is not specified, the entire File Server will be mounted.

            * `unmanaged_file_systems` (`pulumi.Input[list]`)
              * `mount_command` (`pulumi.Input[str]`)
              * `relative_mount_path` (`pulumi.Input[str]`) - Note that all cluster level unmanaged file system will be mounted under $AZ_BATCHAI_MOUNT_ROOT location and job level unmanaged file system will be mounted under $AZ_BATCHAI_JOB_MOUNT_ROOT.

          * `node_count` (`pulumi.Input[float]`) - The job will be gang scheduled on that many compute nodes
          * `output_directories` (`pulumi.Input[list]`)
            * `create_new` (`pulumi.Input[bool]`) - Default is true. If false, then the directory is not created and can be any directory path that the user specifies.
            * `id` (`pulumi.Input[str]`) - The path of the output directory will be available as a value of an environment variable with AZ_BATCHAI_OUTPUT_<id> name, where <id> is the value of id attribute.
            * `path_prefix` (`pulumi.Input[str]`) - NOTE: This is an absolute path to prefix. E.g. $AZ_BATCHAI_MOUNT_ROOT/MyNFS/MyLogs. You can find the full path to the output directory by combining pathPrefix, jobOutputDirectoryPathSegment (reported by get job) and pathSuffix.
            * `path_suffix` (`pulumi.Input[str]`) - The suffix path where the output directory will be created. E.g. models. You can find the full path to the output directory by combining pathPrefix, jobOutputDirectoryPathSegment (reported by get job) and pathSuffix.
            * `type` (`pulumi.Input[str]`) - Default value is Custom. The possible values are Model, Logs, Summary, and Custom. Users can use multiple enums for a single directory. Eg. outPutType='Model,Logs, Summary'

          * `priority` (`pulumi.Input[float]`) - Priority associated with the job. Priority values can range from -1000 to 1000, with -1000 being the lowest priority and 1000 being the highest priority. The default value is 0.
          * `py_torch_settings` (`pulumi.Input[dict]`) - Specifies the settings for pyTorch job.
            * `command_line_args` (`pulumi.Input[str]`)
            * `communication_backend` (`pulumi.Input[str]`) - Valid values are 'TCP', 'Gloo' or 'MPI'. Not required for non-distributed jobs.
            * `process_count` (`pulumi.Input[float]`) - The default value for this property is equal to nodeCount property.
            * `python_interpreter_path` (`pulumi.Input[str]`)
            * `python_script_file_path` (`pulumi.Input[str]`)

          * `secrets` (`pulumi.Input[list]`) - Batch AI will setup these additional environment variables for the job. Server will never report values of these variables back.
            * `name` (`pulumi.Input[str]`)
            * `value` (`pulumi.Input[str]`)
            * `value_secret_reference` (`pulumi.Input[dict]`) - Specifies KeyVault Store and Secret which contains the value for the environment variable. One of value or valueSecretReference must be provided.

          * `std_out_err_path_prefix` (`pulumi.Input[str]`) - The path where the Batch AI service will upload stdout and stderror of the job.
          * `tensor_flow_settings` (`pulumi.Input[dict]`) - Specifies the settings for TensorFlow job.
            * `master_command_line_args` (`pulumi.Input[str]`)
            * `parameter_server_command_line_args` (`pulumi.Input[str]`) - This property is optional for single machine training.
            * `parameter_server_count` (`pulumi.Input[float]`) - If specified, the value must be less than or equal to nodeCount. If not specified, the default value is equal to 1 for distributed TensorFlow training (This property is not applicable for single machine training). This property can be specified only for distributed TensorFlow training.
            * `python_interpreter_path` (`pulumi.Input[str]`)
            * `python_script_file_path` (`pulumi.Input[str]`)
            * `worker_command_line_args` (`pulumi.Input[str]`) - This property is optional for single machine training.
            * `worker_count` (`pulumi.Input[float]`) - If specified, the value must be less than or equal to (nodeCount * numberOfGPUs per VM). If not specified, the default value is equal to nodeCount. This property can be specified only for distributed TensorFlow training
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

            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['type'] = None
        super(Job, __self__).__init__(
            'azurerm:batchai/v20180301:Job',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Job resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Job(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
