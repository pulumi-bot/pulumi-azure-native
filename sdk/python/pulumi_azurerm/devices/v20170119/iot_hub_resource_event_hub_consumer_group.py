# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = ['IotHubResourceEventHubConsumerGroup']


class IotHubResourceEventHubConsumerGroup(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 event_hub_endpoint_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 resource_name_: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The properties of the EventHubConsumerGroupInfo object.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] event_hub_endpoint_name: The name of the Event Hub-compatible endpoint in the IoT hub.
        :param pulumi.Input[str] name: The name of the consumer group to add.
        :param pulumi.Input[str] resource_group_name: The name of the resource group that contains the IoT hub.
        :param pulumi.Input[str] resource_name_: The name of the IoT hub.
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

            if event_hub_endpoint_name is None:
                raise TypeError("Missing required property 'event_hub_endpoint_name'")
            __props__['event_hub_endpoint_name'] = event_hub_endpoint_name
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if resource_name_ is None:
                raise TypeError("Missing required property 'resource_name_'")
            __props__['resource_name'] = resource_name_
            __props__['tags'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:devices/latest:IotHubResourceEventHubConsumerGroup"), pulumi.Alias(type_="azurerm:devices/v20160203:IotHubResourceEventHubConsumerGroup"), pulumi.Alias(type_="azurerm:devices/v20170701:IotHubResourceEventHubConsumerGroup"), pulumi.Alias(type_="azurerm:devices/v20180122:IotHubResourceEventHubConsumerGroup"), pulumi.Alias(type_="azurerm:devices/v20180401:IotHubResourceEventHubConsumerGroup"), pulumi.Alias(type_="azurerm:devices/v20181201preview:IotHubResourceEventHubConsumerGroup"), pulumi.Alias(type_="azurerm:devices/v20190322:IotHubResourceEventHubConsumerGroup"), pulumi.Alias(type_="azurerm:devices/v20190322preview:IotHubResourceEventHubConsumerGroup"), pulumi.Alias(type_="azurerm:devices/v20190701preview:IotHubResourceEventHubConsumerGroup"), pulumi.Alias(type_="azurerm:devices/v20191104:IotHubResourceEventHubConsumerGroup"), pulumi.Alias(type_="azurerm:devices/v20200301:IotHubResourceEventHubConsumerGroup"), pulumi.Alias(type_="azurerm:devices/v20200401:IotHubResourceEventHubConsumerGroup"), pulumi.Alias(type_="azurerm:devices/v20200615:IotHubResourceEventHubConsumerGroup"), pulumi.Alias(type_="azurerm:devices/v20200710preview:IotHubResourceEventHubConsumerGroup"), pulumi.Alias(type_="azurerm:devices/v20200801:IotHubResourceEventHubConsumerGroup")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(IotHubResourceEventHubConsumerGroup, __self__).__init__(
            'azurerm:devices/v20170119:IotHubResourceEventHubConsumerGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'IotHubResourceEventHubConsumerGroup':
        """
        Get an existing IotHubResourceEventHubConsumerGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return IotHubResourceEventHubConsumerGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        """
        The Event Hub-compatible consumer group name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        The tags.
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

