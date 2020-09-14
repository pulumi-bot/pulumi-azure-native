# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = ['Channel']


class Channel(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_name: Optional[pulumi.Input[str]] = None,
                 channel_functions: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 channel_name: Optional[pulumi.Input[str]] = None,
                 channel_type: Optional[pulumi.Input[str]] = None,
                 credentials: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The EngagementFabric channel

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: Account Name
        :param pulumi.Input[List[pulumi.Input[str]]] channel_functions: The functions to be enabled for the channel
        :param pulumi.Input[str] channel_name: Channel Name
        :param pulumi.Input[str] channel_type: The channel type
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] credentials: The channel credentials
        :param pulumi.Input[str] resource_group_name: Resource Group Name
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
            __props__['channel_functions'] = channel_functions
            if channel_name is None:
                raise TypeError("Missing required property 'channel_name'")
            __props__['channel_name'] = channel_name
            if channel_type is None:
                raise TypeError("Missing required property 'channel_type'")
            __props__['channel_type'] = channel_type
            __props__['credentials'] = credentials
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['name'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:engagementfabric/latest:Channel")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Channel, __self__).__init__(
            'azurerm:engagementfabric/v20180901preview:Channel',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Channel':
        """
        Get an existing Channel resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Channel(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="channelFunctions")
    def channel_functions(self) -> pulumi.Output[Optional[List[str]]]:
        """
        The functions to be enabled for the channel
        """
        return pulumi.get(self, "channel_functions")

    @property
    @pulumi.getter(name="channelType")
    def channel_type(self) -> pulumi.Output[str]:
        """
        The channel type
        """
        return pulumi.get(self, "channel_type")

    @property
    @pulumi.getter
    def credentials(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        The channel credentials
        """
        return pulumi.get(self, "credentials")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The fully qualified type of the resource
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

