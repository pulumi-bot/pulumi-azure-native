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

__all__ = ['EventChannel']


class EventChannel(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 destination: Optional[pulumi.Input[pulumi.InputType['EventChannelDestinationArgs']]] = None,
                 event_channel_name: Optional[pulumi.Input[str]] = None,
                 expiration_time_if_not_activated_utc: Optional[pulumi.Input[str]] = None,
                 filter: Optional[pulumi.Input[pulumi.InputType['EventChannelFilterArgs']]] = None,
                 partner_namespace_name: Optional[pulumi.Input[str]] = None,
                 partner_topic_friendly_description: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[pulumi.InputType['EventChannelSourceArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Event Channel.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['EventChannelDestinationArgs']] destination: Represents the destination of an event channel.
        :param pulumi.Input[str] event_channel_name: Name of the event channel.
        :param pulumi.Input[str] expiration_time_if_not_activated_utc: Expiration time of the event channel. If this timer expires while the corresponding partner topic is never activated,
               the event channel and corresponding partner topic are deleted.
        :param pulumi.Input[pulumi.InputType['EventChannelFilterArgs']] filter: Information about the filter for the event channel.
        :param pulumi.Input[str] partner_namespace_name: Name of the partner namespace.
        :param pulumi.Input[str] partner_topic_friendly_description: Friendly description about the topic. This can be set by the publisher/partner to show custom description for the customer partner topic.
               This will be helpful to remove any ambiguity of the origin of creation of the partner topic for the customer.
        :param pulumi.Input[str] resource_group_name: The name of the resource group within the user's subscription.
        :param pulumi.Input[pulumi.InputType['EventChannelSourceArgs']] source: Source of the event channel. This represents a unique resource in the partner's resource model.
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

            __props__['destination'] = destination
            if event_channel_name is None and not opts.urn:
                raise TypeError("Missing required property 'event_channel_name'")
            __props__['event_channel_name'] = event_channel_name
            __props__['expiration_time_if_not_activated_utc'] = expiration_time_if_not_activated_utc
            __props__['filter'] = filter
            if partner_namespace_name is None and not opts.urn:
                raise TypeError("Missing required property 'partner_namespace_name'")
            __props__['partner_namespace_name'] = partner_namespace_name
            __props__['partner_topic_friendly_description'] = partner_topic_friendly_description
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['source'] = source
            __props__['name'] = None
            __props__['partner_topic_readiness_state'] = None
            __props__['provisioning_state'] = None
            __props__['system_data'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:eventgrid/v20200401preview:EventChannel")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(EventChannel, __self__).__init__(
            'azure-nextgen:eventgrid/v20201015preview:EventChannel',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'EventChannel':
        """
        Get an existing EventChannel resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return EventChannel(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def destination(self) -> pulumi.Output[Optional['outputs.EventChannelDestinationResponse']]:
        """
        Represents the destination of an event channel.
        """
        return pulumi.get(self, "destination")

    @property
    @pulumi.getter(name="expirationTimeIfNotActivatedUtc")
    def expiration_time_if_not_activated_utc(self) -> pulumi.Output[Optional[str]]:
        """
        Expiration time of the event channel. If this timer expires while the corresponding partner topic is never activated,
        the event channel and corresponding partner topic are deleted.
        """
        return pulumi.get(self, "expiration_time_if_not_activated_utc")

    @property
    @pulumi.getter
    def filter(self) -> pulumi.Output[Optional['outputs.EventChannelFilterResponse']]:
        """
        Information about the filter for the event channel.
        """
        return pulumi.get(self, "filter")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="partnerTopicFriendlyDescription")
    def partner_topic_friendly_description(self) -> pulumi.Output[Optional[str]]:
        """
        Friendly description about the topic. This can be set by the publisher/partner to show custom description for the customer partner topic.
        This will be helpful to remove any ambiguity of the origin of creation of the partner topic for the customer.
        """
        return pulumi.get(self, "partner_topic_friendly_description")

    @property
    @pulumi.getter(name="partnerTopicReadinessState")
    def partner_topic_readiness_state(self) -> pulumi.Output[str]:
        """
        The readiness state of the corresponding partner topic.
        """
        return pulumi.get(self, "partner_topic_readiness_state")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        Provisioning state of the event channel.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def source(self) -> pulumi.Output[Optional['outputs.EventChannelSourceResponse']]:
        """
        Source of the event channel. This represents a unique resource in the partner's resource model.
        """
        return pulumi.get(self, "source")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> pulumi.Output['outputs.SystemDataResponse']:
        """
        The system metadata relating to this resource.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Type of the resource.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

