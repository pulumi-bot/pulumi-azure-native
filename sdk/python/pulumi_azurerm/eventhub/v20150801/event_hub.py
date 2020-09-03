# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = ['EventHub']


class EventHub(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 event_hub_name: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 message_retention_in_days: Optional[pulumi.Input[float]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace_name: Optional[pulumi.Input[str]] = None,
                 partition_count: Optional[pulumi.Input[float]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Single item in List or Get Event Hub operation

        ## Example Usage
        ### EventHubCreate

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        event_hub = azurerm.eventhub.v20150801.EventHub("eventHub",
            event_hub_name="sdk-EventHub6448",
            location="West US",
            message_retention_in_days=7,
            namespace_name="sdk-Namespace7834",
            partition_count=4,
            resource_group_name="Default-ServiceBus-WestUS",
            status="Active")

        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] event_hub_name: The Event Hub name
        :param pulumi.Input[str] location: Location of the resource.
        :param pulumi.Input[float] message_retention_in_days: Number of days to retain the events for this Event Hub.
        :param pulumi.Input[str] name: Name of the Event Hub.
        :param pulumi.Input[str] namespace_name: The Namespace name
        :param pulumi.Input[float] partition_count: Number of partitions created for the Event Hub.
        :param pulumi.Input[str] resource_group_name: Name of the resource group within the azure subscription.
        :param pulumi.Input[str] status: Enumerates the possible values for the status of the Event Hub.
        :param pulumi.Input[str] type: ARM type of the Namespace.
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

            if event_hub_name is None:
                raise TypeError("Missing required property 'event_hub_name'")
            __props__['event_hub_name'] = event_hub_name
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            __props__['message_retention_in_days'] = message_retention_in_days
            __props__['name'] = name
            if namespace_name is None:
                raise TypeError("Missing required property 'namespace_name'")
            __props__['namespace_name'] = namespace_name
            __props__['partition_count'] = partition_count
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['status'] = status
            __props__['type'] = type
            __props__['created_at'] = None
            __props__['partition_ids'] = None
            __props__['updated_at'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:eventhub/latest:EventHub"), pulumi.Alias(type_="azurerm:eventhub/v20140901:EventHub"), pulumi.Alias(type_="azurerm:eventhub/v20170401:EventHub")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(EventHub, __self__).__init__(
            'azurerm:eventhub/v20150801:EventHub',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'EventHub':
        """
        Get an existing EventHub resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return EventHub(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        Exact time the Event Hub was created.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[Optional[str]]:
        """
        Resource location
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="messageRetentionInDays")
    def message_retention_in_days(self) -> pulumi.Output[Optional[float]]:
        """
        Number of days to retain the events for this Event Hub.
        """
        return pulumi.get(self, "message_retention_in_days")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="partitionCount")
    def partition_count(self) -> pulumi.Output[Optional[float]]:
        """
        Number of partitions created for the Event Hub.
        """
        return pulumi.get(self, "partition_count")

    @property
    @pulumi.getter(name="partitionIds")
    def partition_ids(self) -> pulumi.Output[List[str]]:
        """
        Current number of shards on the Event Hub.
        """
        return pulumi.get(self, "partition_ids")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[Optional[str]]:
        """
        Enumerates the possible values for the status of the Event Hub.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Resource type
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        """
        The exact time the message was updated.
        """
        return pulumi.get(self, "updated_at")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

