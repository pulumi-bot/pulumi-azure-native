# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = ['FirewallRule']


class FirewallRule(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cache_name: Optional[pulumi.Input[str]] = None,
                 end_ip: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 rule_name: Optional[pulumi.Input[str]] = None,
                 start_ip: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        A firewall rule on a redis cache has a name, and describes a contiguous range of IP addresses permitted to connect

        ## Example Usage
        ### RedisCacheFirewallRuleCreate

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        firewall_rule = azurerm.cache.v20170201.FirewallRule("firewallRule",
            cache_name="cache1",
            end_ip="192.168.1.4",
            resource_group_name="rg1",
            rule_name="rule1",
            start_ip="192.168.1.1")

        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cache_name: The name of the Redis cache.
        :param pulumi.Input[str] end_ip: highest IP address included in the range
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] rule_name: The name of the firewall rule.
        :param pulumi.Input[str] start_ip: lowest IP address included in the range
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

            if cache_name is None:
                raise TypeError("Missing required property 'cache_name'")
            __props__['cache_name'] = cache_name
            if end_ip is None:
                raise TypeError("Missing required property 'end_ip'")
            __props__['end_ip'] = end_ip
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if rule_name is None:
                raise TypeError("Missing required property 'rule_name'")
            __props__['rule_name'] = rule_name
            if start_ip is None:
                raise TypeError("Missing required property 'start_ip'")
            __props__['start_ip'] = start_ip
            __props__['name'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:cache/latest:FirewallRule"), pulumi.Alias(type_="azurerm:cache/v20160401:FirewallRule"), pulumi.Alias(type_="azurerm:cache/v20171001:FirewallRule"), pulumi.Alias(type_="azurerm:cache/v20180301:FirewallRule"), pulumi.Alias(type_="azurerm:cache/v20190701:FirewallRule")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(FirewallRule, __self__).__init__(
            'azurerm:cache/v20170201:FirewallRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'FirewallRule':
        """
        Get an existing FirewallRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return FirewallRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="endIP")
    def end_ip(self) -> pulumi.Output[str]:
        """
        highest IP address included in the range
        """
        return pulumi.get(self, "end_ip")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="startIP")
    def start_ip(self) -> pulumi.Output[str]:
        """
        lowest IP address included in the range
        """
        return pulumi.get(self, "start_ip")

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

