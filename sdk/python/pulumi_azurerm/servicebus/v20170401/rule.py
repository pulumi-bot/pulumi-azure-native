# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class Rule(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    Resource name
    """
    properties: pulumi.Output[dict]
    """
    Properties of Rule resource
      * `action` (`dict`) - Represents the filter actions which are allowed for the transformation of a message that have been matched by a filter expression.
        * `compatibility_level` (`float`) - This property is reserved for future use. An integer value showing the compatibility level, currently hard-coded to 20.
        * `requires_preprocessing` (`bool`) - Value that indicates whether the rule action requires preprocessing.
        * `sql_expression` (`str`) - SQL expression. e.g. MyProperty='ABC'

      * `correlation_filter` (`dict`) - Properties of correlationFilter
        * `content_type` (`str`) - Content type of the message.
        * `correlation_id` (`str`) - Identifier of the correlation.
        * `label` (`str`) - Application specific label.
        * `message_id` (`str`) - Identifier of the message.
        * `properties` (`dict`) - dictionary object for custom filters
        * `reply_to` (`str`) - Address of the queue to reply to.
        * `reply_to_session_id` (`str`) - Session identifier to reply to.
        * `requires_preprocessing` (`bool`) - Value that indicates whether the rule action requires preprocessing.
        * `session_id` (`str`) - Session identifier.
        * `to` (`str`) - Address to send to.

      * `filter_type` (`str`) - Filter type that is evaluated against a BrokeredMessage.
      * `sql_filter` (`dict`) - Properties of sqlFilter
        * `compatibility_level` (`float`) - This property is reserved for future use. An integer value showing the compatibility level, currently hard-coded to 20.
        * `requires_preprocessing` (`bool`) - Value that indicates whether the rule action requires preprocessing.
        * `sql_expression` (`str`) - The SQL expression. e.g. MyProperty='ABC'
    """
    type: pulumi.Output[str]
    """
    Resource type
    """
    def __init__(__self__, resource_name, opts=None, action=None, correlation_filter=None, filter_type=None, name=None, namespace_name=None, resource_group_name=None, sql_filter=None, subscription_name=None, topic_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Description of Rule Resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] action: Represents the filter actions which are allowed for the transformation of a message that have been matched by a filter expression.
        :param pulumi.Input[dict] correlation_filter: Properties of correlationFilter
        :param pulumi.Input[str] filter_type: Filter type that is evaluated against a BrokeredMessage.
        :param pulumi.Input[str] name: The rule name.
        :param pulumi.Input[str] namespace_name: The namespace name
        :param pulumi.Input[str] resource_group_name: Name of the Resource group within the Azure subscription.
        :param pulumi.Input[dict] sql_filter: Properties of sqlFilter
        :param pulumi.Input[str] subscription_name: The subscription name.
        :param pulumi.Input[str] topic_name: The topic name.

        The **action** object supports the following:

          * `compatibility_level` (`pulumi.Input[float]`) - This property is reserved for future use. An integer value showing the compatibility level, currently hard-coded to 20.
          * `requires_preprocessing` (`pulumi.Input[bool]`) - Value that indicates whether the rule action requires preprocessing.
          * `sql_expression` (`pulumi.Input[str]`) - SQL expression. e.g. MyProperty='ABC'

        The **correlation_filter** object supports the following:

          * `content_type` (`pulumi.Input[str]`) - Content type of the message.
          * `correlation_id` (`pulumi.Input[str]`) - Identifier of the correlation.
          * `label` (`pulumi.Input[str]`) - Application specific label.
          * `message_id` (`pulumi.Input[str]`) - Identifier of the message.
          * `properties` (`pulumi.Input[dict]`) - dictionary object for custom filters
          * `reply_to` (`pulumi.Input[str]`) - Address of the queue to reply to.
          * `reply_to_session_id` (`pulumi.Input[str]`) - Session identifier to reply to.
          * `requires_preprocessing` (`pulumi.Input[bool]`) - Value that indicates whether the rule action requires preprocessing.
          * `session_id` (`pulumi.Input[str]`) - Session identifier.
          * `to` (`pulumi.Input[str]`) - Address to send to.

        The **sql_filter** object supports the following:

          * `compatibility_level` (`pulumi.Input[float]`) - This property is reserved for future use. An integer value showing the compatibility level, currently hard-coded to 20.
          * `requires_preprocessing` (`pulumi.Input[bool]`) - Value that indicates whether the rule action requires preprocessing.
          * `sql_expression` (`pulumi.Input[str]`) - The SQL expression. e.g. MyProperty='ABC'
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

            __props__['action'] = action
            __props__['correlation_filter'] = correlation_filter
            __props__['filter_type'] = filter_type
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if namespace_name is None:
                raise TypeError("Missing required property 'namespace_name'")
            __props__['namespace_name'] = namespace_name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['sql_filter'] = sql_filter
            if subscription_name is None:
                raise TypeError("Missing required property 'subscription_name'")
            __props__['subscription_name'] = subscription_name
            if topic_name is None:
                raise TypeError("Missing required property 'topic_name'")
            __props__['topic_name'] = topic_name
            __props__['properties'] = None
            __props__['type'] = None
        super(Rule, __self__).__init__(
            'azurerm:servicebus/v20170401:Rule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Rule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Rule(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
