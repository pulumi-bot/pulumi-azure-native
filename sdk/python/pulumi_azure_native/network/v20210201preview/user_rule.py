# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['UserRule']


class UserRule(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 configuration_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 destination: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AddressPrefixItemArgs']]]]] = None,
                 destination_port_ranges: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 direction: Optional[pulumi.Input[Union[str, 'SecurityConfigurationRuleDirection']]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 network_manager_name: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[Union[str, 'SecurityConfigurationRuleProtocol']]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 rule_name: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AddressPrefixItemArgs']]]]] = None,
                 source_port_ranges: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Network security admin rule.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] configuration_name: The name of the network manager security Configuration.
        :param pulumi.Input[str] description: A description for this rule. Restricted to 140 chars.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AddressPrefixItemArgs']]]] destination: The destination address prefixes. CIDR or destination IP ranges.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] destination_port_ranges: The destination port ranges.
        :param pulumi.Input[Union[str, 'SecurityConfigurationRuleDirection']] direction: Indicates if the traffic matched against the rule in inbound or outbound.
        :param pulumi.Input[str] display_name: A friendly name for the rule.
        :param pulumi.Input[str] network_manager_name: The name of the network manager.
        :param pulumi.Input[Union[str, 'SecurityConfigurationRuleProtocol']] protocol: Network protocol this rule applies to.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] rule_name: The name of the rule.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AddressPrefixItemArgs']]]] source: The CIDR or source IP ranges.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] source_port_ranges: The source port ranges.
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

            if configuration_name is None and not opts.urn:
                raise TypeError("Missing required property 'configuration_name'")
            __props__['configuration_name'] = configuration_name
            __props__['description'] = description
            __props__['destination'] = destination
            __props__['destination_port_ranges'] = destination_port_ranges
            if direction is None and not opts.urn:
                raise TypeError("Missing required property 'direction'")
            __props__['direction'] = direction
            __props__['display_name'] = display_name
            if network_manager_name is None and not opts.urn:
                raise TypeError("Missing required property 'network_manager_name'")
            __props__['network_manager_name'] = network_manager_name
            if protocol is None and not opts.urn:
                raise TypeError("Missing required property 'protocol'")
            __props__['protocol'] = protocol
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['rule_name'] = rule_name
            __props__['source'] = source
            __props__['source_port_ranges'] = source_port_ranges
            __props__['etag'] = None
            __props__['name'] = None
            __props__['provisioning_state'] = None
            __props__['system_data'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:network/v20210201preview:UserRule"), pulumi.Alias(type_="azure-native:network:UserRule"), pulumi.Alias(type_="azure-nextgen:network:UserRule")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(UserRule, __self__).__init__(
            'azure-native:network/v20210201preview:UserRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'UserRule':
        """
        Get an existing UserRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = None
        __props__["destination"] = None
        __props__["destination_port_ranges"] = None
        __props__["direction"] = None
        __props__["display_name"] = None
        __props__["etag"] = None
        __props__["name"] = None
        __props__["protocol"] = None
        __props__["provisioning_state"] = None
        __props__["source"] = None
        __props__["source_port_ranges"] = None
        __props__["system_data"] = None
        __props__["type"] = None
        return UserRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description for this rule. Restricted to 140 chars.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def destination(self) -> pulumi.Output[Optional[Sequence['outputs.AddressPrefixItemResponse']]]:
        """
        The destination address prefixes. CIDR or destination IP ranges.
        """
        return pulumi.get(self, "destination")

    @property
    @pulumi.getter(name="destinationPortRanges")
    def destination_port_ranges(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The destination port ranges.
        """
        return pulumi.get(self, "destination_port_ranges")

    @property
    @pulumi.getter
    def direction(self) -> pulumi.Output[str]:
        """
        Indicates if the traffic matched against the rule in inbound or outbound.
        """
        return pulumi.get(self, "direction")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[Optional[str]]:
        """
        A friendly name for the rule.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        A unique read-only string that changes whenever the resource is updated.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Output[str]:
        """
        Network protocol this rule applies to.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        The provisioning state of the security Configuration resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def source(self) -> pulumi.Output[Optional[Sequence['outputs.AddressPrefixItemResponse']]]:
        """
        The CIDR or source IP ranges.
        """
        return pulumi.get(self, "source")

    @property
    @pulumi.getter(name="sourcePortRanges")
    def source_port_ranges(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The source port ranges.
        """
        return pulumi.get(self, "source_port_ranges")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> pulumi.Output['outputs.SystemDataResponse']:
        """
        The system metadata related to this resource.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Resource type.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

