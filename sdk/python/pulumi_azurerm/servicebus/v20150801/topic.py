# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class Topic(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    Resource location.
    """
    name: pulumi.Output[str]
    """
    Resource name
    """
    properties: pulumi.Output[dict]
    """
    The Topic Properties definition.
      * `accessed_at` (`str`) - Last time the message was sent, or a request was received, for this topic.
      * `auto_delete_on_idle` (`str`) - TimeSpan idle interval after which the topic is automatically deleted. The minimum duration is 5 minutes.
      * `count_details` (`dict`) - Message Count Details.
        * `active_message_count` (`float`) - Number of active messages in the queue, topic, or subscription.
        * `dead_letter_message_count` (`float`) - Number of messages that are dead lettered.
        * `scheduled_message_count` (`float`) - Number of scheduled messages.
        * `transfer_dead_letter_message_count` (`float`) - Number of messages transferred into dead letters.
        * `transfer_message_count` (`float`) - Number of messages transferred to another queue, topic, or subscription.

      * `created_at` (`str`) - Exact time the message was created.
      * `default_message_time_to_live` (`str`) - Default message time to live value. This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
      * `duplicate_detection_history_time_window` (`str`) - TimeSpan structure that defines the duration of the duplicate detection history. The default value is 10 minutes.
      * `enable_batched_operations` (`bool`) - Value that indicates whether server-side batched operations are enabled.
      * `enable_express` (`bool`) - Value that indicates whether Express Entities are enabled. An express topic holds a message in memory temporarily before writing it to persistent storage.
      * `enable_partitioning` (`bool`) - Value that indicates whether the topic to be partitioned across multiple message brokers is enabled.
      * `entity_availability_status` (`str`) - Entity availability status for the topic.
      * `filtering_messages_before_publishing` (`bool`) - Whether messages should be filtered before publishing.
      * `is_anonymous_accessible` (`bool`) - Value that indicates whether the message is accessible anonymously.
      * `is_express` (`bool`)
      * `max_size_in_megabytes` (`float`) - Maximum size of the topic in megabytes, which is the size of the memory allocated for the topic.
      * `requires_duplicate_detection` (`bool`) - Value indicating if this topic requires duplicate detection.
      * `size_in_bytes` (`float`) - Size of the topic, in bytes.
      * `status` (`str`) - Enumerates the possible values for the status of a messaging entity.
      * `subscription_count` (`float`) - Number of subscriptions.
      * `support_ordering` (`bool`) - Value that indicates whether the topic supports ordering.
      * `updated_at` (`str`) - The exact time the message was updated.
    """
    type: pulumi.Output[str]
    """
    Resource type
    """
    def __init__(__self__, resource_name, opts=None, location=None, name=None, namespace_name=None, properties=None, resource_group_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Description of topic resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: Location of the resource.
        :param pulumi.Input[str] name: The topic name.
        :param pulumi.Input[str] namespace_name: The namespace name
        :param pulumi.Input[dict] properties: The Topic Properties definition.
        :param pulumi.Input[str] resource_group_name: Name of the Resource group within the Azure subscription.

        The **properties** object supports the following:

          * `auto_delete_on_idle` (`pulumi.Input[str]`) - TimeSpan idle interval after which the topic is automatically deleted. The minimum duration is 5 minutes.
          * `default_message_time_to_live` (`pulumi.Input[str]`) - Default message time to live value. This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
          * `duplicate_detection_history_time_window` (`pulumi.Input[str]`) - TimeSpan structure that defines the duration of the duplicate detection history. The default value is 10 minutes.
          * `enable_batched_operations` (`pulumi.Input[bool]`) - Value that indicates whether server-side batched operations are enabled.
          * `enable_express` (`pulumi.Input[bool]`) - Value that indicates whether Express Entities are enabled. An express topic holds a message in memory temporarily before writing it to persistent storage.
          * `enable_partitioning` (`pulumi.Input[bool]`) - Value that indicates whether the topic to be partitioned across multiple message brokers is enabled.
          * `entity_availability_status` (`pulumi.Input[str]`) - Entity availability status for the topic.
          * `filtering_messages_before_publishing` (`pulumi.Input[bool]`) - Whether messages should be filtered before publishing.
          * `is_anonymous_accessible` (`pulumi.Input[bool]`) - Value that indicates whether the message is accessible anonymously.
          * `is_express` (`pulumi.Input[bool]`)
          * `max_size_in_megabytes` (`pulumi.Input[float]`) - Maximum size of the topic in megabytes, which is the size of the memory allocated for the topic.
          * `requires_duplicate_detection` (`pulumi.Input[bool]`) - Value indicating if this topic requires duplicate detection.
          * `status` (`pulumi.Input[str]`) - Enumerates the possible values for the status of a messaging entity.
          * `support_ordering` (`pulumi.Input[bool]`) - Value that indicates whether the topic supports ordering.
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
            if namespace_name is None:
                raise TypeError("Missing required property 'namespace_name'")
            __props__['namespace_name'] = namespace_name
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['type'] = None
        super(Topic, __self__).__init__(
            'azurerm:servicebus/v20150801:Topic',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Topic resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Topic(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
