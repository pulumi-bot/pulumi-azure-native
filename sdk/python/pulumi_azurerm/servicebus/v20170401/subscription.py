# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class Subscription(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    Resource name
    """
    properties: pulumi.Output[dict]
    """
    Properties of subscriptions resource.
      * `accessed_at` (`str`) - Last time there was a receive request to this subscription.
      * `auto_delete_on_idle` (`str`) - ISO 8061 timeSpan idle interval after which the topic is automatically deleted. The minimum duration is 5 minutes.
      * `count_details` (`dict`) - Message count details
        * `active_message_count` (`float`) - Number of active messages in the queue, topic, or subscription.
        * `dead_letter_message_count` (`float`) - Number of messages that are dead lettered.
        * `scheduled_message_count` (`float`) - Number of scheduled messages.
        * `transfer_dead_letter_message_count` (`float`) - Number of messages transferred into dead letters.
        * `transfer_message_count` (`float`) - Number of messages transferred to another queue, topic, or subscription.

      * `created_at` (`str`) - Exact time the message was created.
      * `dead_lettering_on_filter_evaluation_exceptions` (`bool`) - Value that indicates whether a subscription has dead letter support on filter evaluation exceptions.
      * `dead_lettering_on_message_expiration` (`bool`) - Value that indicates whether a subscription has dead letter support when a message expires.
      * `default_message_time_to_live` (`str`) - ISO 8061 Default message timespan to live value. This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
      * `duplicate_detection_history_time_window` (`str`) - ISO 8601 timeSpan structure that defines the duration of the duplicate detection history. The default value is 10 minutes.
      * `enable_batched_operations` (`bool`) - Value that indicates whether server-side batched operations are enabled.
      * `forward_dead_lettered_messages_to` (`str`) - Queue/Topic name to forward the Dead Letter message
      * `forward_to` (`str`) - Queue/Topic name to forward the messages
      * `lock_duration` (`str`) - ISO 8061 lock duration timespan for the subscription. The default value is 1 minute.
      * `max_delivery_count` (`float`) - Number of maximum deliveries.
      * `message_count` (`float`) - Number of messages.
      * `requires_session` (`bool`) - Value indicating if a subscription supports the concept of sessions.
      * `status` (`str`) - Enumerates the possible values for the status of a messaging entity.
      * `updated_at` (`str`) - The exact time the message was updated.
    """
    type: pulumi.Output[str]
    """
    Resource type
    """
    def __init__(__self__, resource_name, opts=None, name=None, namespace_name=None, properties=None, resource_group_name=None, topic_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Description of subscription resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The subscription name.
        :param pulumi.Input[str] namespace_name: The namespace name
        :param pulumi.Input[dict] properties: Properties of subscriptions resource.
        :param pulumi.Input[str] resource_group_name: Name of the Resource group within the Azure subscription.
        :param pulumi.Input[str] topic_name: The topic name.

        The **properties** object supports the following:

          * `auto_delete_on_idle` (`pulumi.Input[str]`) - ISO 8061 timeSpan idle interval after which the topic is automatically deleted. The minimum duration is 5 minutes.
          * `dead_lettering_on_filter_evaluation_exceptions` (`pulumi.Input[bool]`) - Value that indicates whether a subscription has dead letter support on filter evaluation exceptions.
          * `dead_lettering_on_message_expiration` (`pulumi.Input[bool]`) - Value that indicates whether a subscription has dead letter support when a message expires.
          * `default_message_time_to_live` (`pulumi.Input[str]`) - ISO 8061 Default message timespan to live value. This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
          * `duplicate_detection_history_time_window` (`pulumi.Input[str]`) - ISO 8601 timeSpan structure that defines the duration of the duplicate detection history. The default value is 10 minutes.
          * `enable_batched_operations` (`pulumi.Input[bool]`) - Value that indicates whether server-side batched operations are enabled.
          * `forward_dead_lettered_messages_to` (`pulumi.Input[str]`) - Queue/Topic name to forward the Dead Letter message
          * `forward_to` (`pulumi.Input[str]`) - Queue/Topic name to forward the messages
          * `lock_duration` (`pulumi.Input[str]`) - ISO 8061 lock duration timespan for the subscription. The default value is 1 minute.
          * `max_delivery_count` (`pulumi.Input[float]`) - Number of maximum deliveries.
          * `requires_session` (`pulumi.Input[bool]`) - Value indicating if a subscription supports the concept of sessions.
          * `status` (`pulumi.Input[str]`) - Enumerates the possible values for the status of a messaging entity.
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
            if topic_name is None:
                raise TypeError("Missing required property 'topic_name'")
            __props__['topic_name'] = topic_name
            __props__['type'] = None
        super(Subscription, __self__).__init__(
            'azurerm:servicebus/v20170401:Subscription',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Subscription resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Subscription(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
