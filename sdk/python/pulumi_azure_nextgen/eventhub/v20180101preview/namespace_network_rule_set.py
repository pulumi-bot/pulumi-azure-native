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

__all__ = ['NamespaceNetworkRuleSet']


class NamespaceNetworkRuleSet(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_action: Optional[pulumi.Input[str]] = None,
                 ip_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NWRuleSetIpRulesArgs']]]]] = None,
                 namespace_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Description of topic resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] default_action: Default Action for Network Rule Set
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NWRuleSetIpRulesArgs']]]] ip_rules: List of IpRules
        :param pulumi.Input[str] namespace_name: The namespace name
        :param pulumi.Input[str] resource_group_name: Name of the Resource group within the Azure subscription.
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

            __props__['default_action'] = default_action
            __props__['ip_rules'] = ip_rules
            if namespace_name is None:
                raise TypeError("Missing required property 'namespace_name'")
            __props__['namespace_name'] = namespace_name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['name'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:eventhub/latest:NamespaceNetworkRuleSet"), pulumi.Alias(type_="azure-nextgen:eventhub/v20170401:NamespaceNetworkRuleSet")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(NamespaceNetworkRuleSet, __self__).__init__(
            'azure-nextgen:eventhub/v20180101preview:NamespaceNetworkRuleSet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'NamespaceNetworkRuleSet':
        """
        Get an existing NamespaceNetworkRuleSet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return NamespaceNetworkRuleSet(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="defaultAction")
    def default_action(self) -> pulumi.Output[Optional[str]]:
        """
        Default Action for Network Rule Set
        """
        return pulumi.get(self, "default_action")

    @property
    @pulumi.getter(name="ipRules")
    def ip_rules(self) -> pulumi.Output[Optional[Sequence['outputs.NWRuleSetIpRulesResponse']]]:
        """
        List of IpRules
        """
        return pulumi.get(self, "ip_rules")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

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

