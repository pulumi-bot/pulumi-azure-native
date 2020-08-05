# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class Job(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    The name of the resource
    """
    properties: pulumi.Output[dict]
    """
    The resource properties.
      * `correlation_data` (`dict`) - Customer provided key, value pairs that will be returned in Job and JobOutput state events.
      * `created` (`str`) - The UTC date and time when the Job was created, in 'YYYY-MM-DDThh:mm:ssZ' format.
      * `description` (`str`) - Optional customer supplied description of the Job.
      * `end_time` (`str`) - The UTC date and time at which this Job finished processing.
      * `input` (`dict`) - The inputs for the Job.
      * `last_modified` (`str`) - The UTC date and time when the Job was last updated, in 'YYYY-MM-DDThh:mm:ssZ' format.
      * `outputs` (`list`) - The outputs for the Job.
        * `end_time` (`str`) - The UTC date and time at which this Job Output finished processing.
        * `error` (`dict`) - If the JobOutput is in the Error state, it contains the details of the error.
          * `category` (`str`) - Helps with categorization of errors.
          * `code` (`str`) - Error code describing the error.
          * `details` (`list`) - An array of details about specific errors that led to this reported error.
            * `code` (`str`) - Code describing the error detail.
            * `message` (`str`) - A human-readable representation of the error.

          * `message` (`str`) - A human-readable language-dependent representation of the error.
          * `retry` (`str`) - Indicates that it may be possible to retry the Job. If retry is unsuccessful, please contact Azure support via Azure Portal.

        * `label` (`str`) - A label that is assigned to a JobOutput in order to help uniquely identify it. This is useful when your Transform has more than one TransformOutput, whereby your Job has more than one JobOutput. In such cases, when you submit the Job, you will add two or more JobOutputs, in the same order as TransformOutputs in the Transform. Subsequently, when you retrieve the Job, either through events or on a GET request, you can use the label to easily identify the JobOutput. If a label is not provided, a default value of '{presetName}_{outputIndex}' will be used, where the preset name is the name of the preset in the corresponding TransformOutput and the output index is the relative index of the this JobOutput within the Job. Note that this index is the same as the relative index of the corresponding TransformOutput within its Transform.
        * `progress` (`float`) - If the JobOutput is in a Processing state, this contains the Job completion percentage. The value is an estimate and not intended to be used to predict Job completion times. To determine if the JobOutput is complete, use the State property.
        * `start_time` (`str`) - The UTC date and time at which this Job Output began processing.
        * `state` (`str`) - Describes the state of the JobOutput.

      * `priority` (`str`) - Priority with which the job should be processed. Higher priority jobs are processed before lower priority jobs. If not set, the default is normal.
      * `start_time` (`str`) - The UTC date and time at which this Job began processing.
      * `state` (`str`) - The current state of the job.
    """
    type: pulumi.Output[str]
    """
    The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
    """
    def __init__(__self__, resource_name, opts=None, account_name=None, correlation_data=None, description=None, input=None, name=None, outputs=None, priority=None, resource_group_name=None, transform_name=None, __props__=None, __name__=None, __opts__=None):
        """
        A Job resource type. The progress and state can be obtained by polling a Job or subscribing to events using EventGrid.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: The Media Services account name.
        :param pulumi.Input[dict] correlation_data: Customer provided key, value pairs that will be returned in Job and JobOutput state events.
        :param pulumi.Input[str] description: Optional customer supplied description of the Job.
        :param pulumi.Input[dict] input: The inputs for the Job.
        :param pulumi.Input[str] name: The Job name.
        :param pulumi.Input[list] outputs: The outputs for the Job.
        :param pulumi.Input[str] priority: Priority with which the job should be processed. Higher priority jobs are processed before lower priority jobs. If not set, the default is normal.
        :param pulumi.Input[str] resource_group_name: The name of the resource group within the Azure subscription.
        :param pulumi.Input[str] transform_name: The Transform name.

        The **outputs** object supports the following:

          * `label` (`pulumi.Input[str]`) - A label that is assigned to a JobOutput in order to help uniquely identify it. This is useful when your Transform has more than one TransformOutput, whereby your Job has more than one JobOutput. In such cases, when you submit the Job, you will add two or more JobOutputs, in the same order as TransformOutputs in the Transform. Subsequently, when you retrieve the Job, either through events or on a GET request, you can use the label to easily identify the JobOutput. If a label is not provided, a default value of '{presetName}_{outputIndex}' will be used, where the preset name is the name of the preset in the corresponding TransformOutput and the output index is the relative index of the this JobOutput within the Job. Note that this index is the same as the relative index of the corresponding TransformOutput within its Transform.
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
            __props__['correlation_data'] = correlation_data
            __props__['description'] = description
            if input is None:
                raise TypeError("Missing required property 'input'")
            __props__['input'] = input
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if outputs is None:
                raise TypeError("Missing required property 'outputs'")
            __props__['outputs'] = outputs
            __props__['priority'] = priority
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if transform_name is None:
                raise TypeError("Missing required property 'transform_name'")
            __props__['transform_name'] = transform_name
            __props__['properties'] = None
            __props__['type'] = None
        super(Job, __self__).__init__(
            'azurerm:media/v20200501:Job',
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
