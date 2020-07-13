# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class DataManagerDataServiceJobDefinition(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    Name of the object.
    """
    properties: pulumi.Output[dict]
    """
    JobDefinition properties.
      * `customer_secrets` (`list`) - List of customer secrets containing a key identifier and key value. The key identifier is a way for the specific data source to understand the key. Value contains customer secret encrypted by the encryptionKeys.
        * `algorithm` (`str`) - The encryption algorithm used to encrypt data.
        * `key_identifier` (`str`) - The identifier to the data service input object which this secret corresponds to.
        * `key_value` (`str`) - It contains the encrypted customer secret.

      * `data_service_input` (`dict`) - A generic json used differently by each data service type.
      * `data_sink_id` (`str`) - Data Sink Id associated to the job definition.
      * `data_source_id` (`str`) - Data Source Id associated to the job definition.
      * `last_modified_time` (`str`) - Last modified time of the job definition.
      * `run_location` (`str`) - This is the preferred geo location for the job to run.
      * `schedules` (`list`) - Schedule for running the job definition
        * `name` (`str`) - Name of the schedule.
        * `policy_list` (`list`) - A list of repetition intervals in ISO 8601 format.

      * `state` (`str`) - State of the job definition.
      * `user_confirmation` (`str`) - Enum to detect if user confirmation is required. If not passed will default to NotRequired.
    """
    type: pulumi.Output[str]
    """
    Type of the object.
    """
    def __init__(__self__, resource_name, opts=None, data_manager_name=None, data_service_name=None, name=None, properties=None, resource_group_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Job Definition.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] data_manager_name: The name of the DataManager Resource within the specified resource group. DataManager names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
        :param pulumi.Input[str] data_service_name: The data service type of the job definition.
        :param pulumi.Input[str] name: The job definition name to be created or updated.
        :param pulumi.Input[dict] properties: JobDefinition properties.
        :param pulumi.Input[str] resource_group_name: The Resource Group Name

        The **properties** object supports the following:

          * `customer_secrets` (`pulumi.Input[list]`) - List of customer secrets containing a key identifier and key value. The key identifier is a way for the specific data source to understand the key. Value contains customer secret encrypted by the encryptionKeys.
            * `algorithm` (`pulumi.Input[str]`) - The encryption algorithm used to encrypt data.
            * `key_identifier` (`pulumi.Input[str]`) - The identifier to the data service input object which this secret corresponds to.
            * `key_value` (`pulumi.Input[str]`) - It contains the encrypted customer secret.

          * `data_service_input` (`pulumi.Input[dict]`) - A generic json used differently by each data service type.
          * `data_sink_id` (`pulumi.Input[str]`) - Data Sink Id associated to the job definition.
          * `data_source_id` (`pulumi.Input[str]`) - Data Source Id associated to the job definition.
          * `last_modified_time` (`pulumi.Input[str]`) - Last modified time of the job definition.
          * `run_location` (`pulumi.Input[str]`) - This is the preferred geo location for the job to run.
          * `schedules` (`pulumi.Input[list]`) - Schedule for running the job definition
            * `name` (`pulumi.Input[str]`) - Name of the schedule.
            * `policy_list` (`pulumi.Input[list]`) - A list of repetition intervals in ISO 8601 format.

          * `state` (`pulumi.Input[str]`) - State of the job definition.
          * `user_confirmation` (`pulumi.Input[str]`) - Enum to detect if user confirmation is required. If not passed will default to NotRequired.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if data_manager_name is None:
                raise TypeError("Missing required property 'data_manager_name'")
            __props__['data_manager_name'] = data_manager_name
            if data_service_name is None:
                raise TypeError("Missing required property 'data_service_name'")
            __props__['data_service_name'] = data_service_name
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if properties is None:
                raise TypeError("Missing required property 'properties'")
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['type'] = None
        super(DataManagerDataServiceJobDefinition, __self__).__init__(
            'azurerm:hybriddata:DataManagerDataServiceJobDefinition',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing DataManagerDataServiceJobDefinition resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return DataManagerDataServiceJobDefinition(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
