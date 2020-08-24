# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['FirewallPolicyRuleGroup']


class FirewallPolicyRuleGroup(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 firewall_policy_name: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[float]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['FirewallPolicyRuleArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Rule Group resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] firewall_policy_name: The name of the Firewall Policy.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[str] name: The name of the FirewallPolicyRuleGroup.
        :param pulumi.Input[float] priority: Priority of the Firewall Policy Rule Group resource.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['FirewallPolicyRuleArgs']]]] rules: Group of Firewall Policy rules.
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

            if firewall_policy_name is None:
                raise TypeError("Missing required property 'firewall_policy_name'")
            __props__['firewall_policy_name'] = firewall_policy_name
            __props__['id'] = id
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['priority'] = priority
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['rules'] = rules
            __props__['etag'] = None
            __props__['provisioning_state'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:network/v20190601:FirewallPolicyRuleGroup"), pulumi.Alias(type_="azurerm:network/v20190701:FirewallPolicyRuleGroup"), pulumi.Alias(type_="azurerm:network/v20190801:FirewallPolicyRuleGroup"), pulumi.Alias(type_="azurerm:network/v20190901:FirewallPolicyRuleGroup"), pulumi.Alias(type_="azurerm:network/v20191201:FirewallPolicyRuleGroup"), pulumi.Alias(type_="azurerm:network/v20200301:FirewallPolicyRuleGroup"), pulumi.Alias(type_="azurerm:network/v20200401:FirewallPolicyRuleGroup")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(FirewallPolicyRuleGroup, __self__).__init__(
            'azurerm:network/v20191101:FirewallPolicyRuleGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'FirewallPolicyRuleGroup':
        """
        Get an existing FirewallPolicyRuleGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return FirewallPolicyRuleGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        A unique read-only string that changes whenever the resource is updated.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The name of the resource that is unique within a resource group. This name can be used to access the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def priority(self) -> Optional[float]:
        """
        Priority of the Firewall Policy Rule Group resource.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        The provisioning state of the firewall policy rule group resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def rules(self) -> Optional[List['outputs.FirewallPolicyRuleResponse']]:
        """
        Group of Firewall Policy rules.
        """
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Rule Group type.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

