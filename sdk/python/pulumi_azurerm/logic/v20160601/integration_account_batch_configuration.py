# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class IntegrationAccountBatchConfiguration(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    The resource location.
    """
    name: pulumi.Output[str]
    """
    Gets the resource name.
    """
    properties: pulumi.Output[dict]
    """
    The batch configuration properties.
      * `batch_group_name` (`str`) - The name of the batch group.
      * `changed_time` (`str`) - The artifact changed time.
      * `created_time` (`str`) - The artifact creation time.
      * `metadata` (`dict`)
      * `release_criteria` (`dict`) - The batch release criteria.
        * `batch_size` (`float`) - The batch size in bytes.
        * `message_count` (`float`) - The message count.
        * `recurrence` (`dict`) - The recurrence.
          * `end_time` (`str`) - The end time.
          * `frequency` (`str`) - The frequency.
          * `interval` (`float`) - The interval.
          * `schedule` (`dict`) - The recurrence schedule.
            * `hours` (`list`) - The hours.
            * `minutes` (`list`) - The minutes.
            * `month_days` (`list`) - The month days.
            * `monthly_occurrences` (`list`) - The monthly occurrences.
              * `day` (`str`) - The day of the week.
              * `occurrence` (`float`) - The occurrence.

            * `week_days` (`list`) - The days of the week.

          * `start_time` (`str`) - The start time.
          * `time_zone` (`str`) - The time zone.
    """
    tags: pulumi.Output[dict]
    """
    The resource tags.
    """
    type: pulumi.Output[str]
    """
    Gets the resource type.
    """
    def __init__(__self__, resource_name, opts=None, integration_account_name=None, location=None, name=None, properties=None, resource_group_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        The batch configuration resource definition.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] integration_account_name: The integration account name.
        :param pulumi.Input[str] location: The resource location.
        :param pulumi.Input[str] name: The batch configuration name.
        :param pulumi.Input[dict] properties: The batch configuration properties.
        :param pulumi.Input[str] resource_group_name: The resource group name.
        :param pulumi.Input[dict] tags: The resource tags.

        The **properties** object supports the following:

          * `batch_group_name` (`pulumi.Input[str]`) - The name of the batch group.
          * `changed_time` (`pulumi.Input[str]`) - The artifact changed time.
          * `created_time` (`pulumi.Input[str]`) - The artifact creation time.
          * `metadata` (`pulumi.Input[dict]`)
          * `release_criteria` (`pulumi.Input[dict]`) - The batch release criteria.
            * `batch_size` (`pulumi.Input[float]`) - The batch size in bytes.
            * `message_count` (`pulumi.Input[float]`) - The message count.
            * `recurrence` (`pulumi.Input[dict]`) - The recurrence.
              * `end_time` (`pulumi.Input[str]`) - The end time.
              * `frequency` (`pulumi.Input[str]`) - The frequency.
              * `interval` (`pulumi.Input[float]`) - The interval.
              * `schedule` (`pulumi.Input[dict]`) - The recurrence schedule.
                * `hours` (`pulumi.Input[list]`) - The hours.
                * `minutes` (`pulumi.Input[list]`) - The minutes.
                * `month_days` (`pulumi.Input[list]`) - The month days.
                * `monthly_occurrences` (`pulumi.Input[list]`) - The monthly occurrences.
                  * `day` (`pulumi.Input[str]`) - The day of the week.
                  * `occurrence` (`pulumi.Input[float]`) - The occurrence.

                * `week_days` (`pulumi.Input[list]`) - The days of the week.

              * `start_time` (`pulumi.Input[str]`) - The start time.
              * `time_zone` (`pulumi.Input[str]`) - The time zone.
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

            if integration_account_name is None:
                raise TypeError("Missing required property 'integration_account_name'")
            __props__['integration_account_name'] = integration_account_name
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if properties is None:
                raise TypeError("Missing required property 'properties'")
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['type'] = None
        super(IntegrationAccountBatchConfiguration, __self__).__init__(
            'azurerm:logic/v20160601:IntegrationAccountBatchConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing IntegrationAccountBatchConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return IntegrationAccountBatchConfiguration(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
