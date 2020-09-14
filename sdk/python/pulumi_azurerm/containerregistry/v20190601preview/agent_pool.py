# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = ['AgentPool']


class AgentPool(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 agent_pool_name: Optional[pulumi.Input[str]] = None,
                 count: Optional[pulumi.Input[float]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 os: Optional[pulumi.Input[str]] = None,
                 registry_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tier: Optional[pulumi.Input[str]] = None,
                 virtual_network_subnet_resource_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The agentpool that has the ARM resource and properties.
        The agentpool will have all information to create an agent pool.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] agent_pool_name: The name of the agent pool.
        :param pulumi.Input[float] count: The count of agent machine
        :param pulumi.Input[str] location: The location of the resource. This cannot be changed after the resource is created.
        :param pulumi.Input[str] os: The OS of agent machine
        :param pulumi.Input[str] registry_name: The name of the container registry.
        :param pulumi.Input[str] resource_group_name: The name of the resource group to which the container registry belongs.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: The tags of the resource.
        :param pulumi.Input[str] tier: The Tier of agent machine
        :param pulumi.Input[str] virtual_network_subnet_resource_id: The Virtual Network Subnet Resource Id of the agent machine
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

            if agent_pool_name is None:
                raise TypeError("Missing required property 'agent_pool_name'")
            __props__['agent_pool_name'] = agent_pool_name
            __props__['count'] = count
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            __props__['os'] = os
            if registry_name is None:
                raise TypeError("Missing required property 'registry_name'")
            __props__['registry_name'] = registry_name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['tier'] = tier
            __props__['virtual_network_subnet_resource_id'] = virtual_network_subnet_resource_id
            __props__['name'] = None
            __props__['provisioning_state'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:containerregistry/latest:AgentPool")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(AgentPool, __self__).__init__(
            'azurerm:containerregistry/v20190601preview:AgentPool',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'AgentPool':
        """
        Get an existing AgentPool resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return AgentPool(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def count(self) -> pulumi.Output[Optional[float]]:
        """
        The count of agent machine
        """
        return pulumi.get(self, "count")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The location of the resource. This cannot be changed after the resource is created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def os(self) -> pulumi.Output[Optional[str]]:
        """
        The OS of agent machine
        """
        return pulumi.get(self, "os")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        The provisioning state of this agent pool
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        The tags of the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def tier(self) -> pulumi.Output[Optional[str]]:
        """
        The Tier of agent machine
        """
        return pulumi.get(self, "tier")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the resource.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="virtualNetworkSubnetResourceId")
    def virtual_network_subnet_resource_id(self) -> pulumi.Output[Optional[str]]:
        """
        The Virtual Network Subnet Resource Id of the agent machine
        """
        return pulumi.get(self, "virtual_network_subnet_resource_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

