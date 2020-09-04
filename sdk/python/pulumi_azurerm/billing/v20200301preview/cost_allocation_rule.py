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

__all__ = ['CostAllocationRule']


class CostAllocationRule(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 billing_account_id: Optional[pulumi.Input[str]] = None,
                 properties: Optional[pulumi.Input[pulumi.InputType['CostAllocationRulePropertiesArgs']]] = None,
                 rule_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The cost allocation rule model definition

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] billing_account_id: BillingAccount ID
        :param pulumi.Input[pulumi.InputType['CostAllocationRulePropertiesArgs']] properties: Cost allocation rule properties
        :param pulumi.Input[str] rule_name: Cost allocation rule name. The name cannot include spaces or any non alphanumeric characters other than '_' and '-'. The max length is 260 characters.
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

            if billing_account_id is None:
                raise TypeError("Missing required property 'billing_account_id'")
            __props__['billing_account_id'] = billing_account_id
            __props__['properties'] = properties
            if rule_name is None:
                raise TypeError("Missing required property 'rule_name'")
            __props__['rule_name'] = rule_name
            __props__['name'] = None
            __props__['type'] = None
        super(CostAllocationRule, __self__).__init__(
            'azurerm:billing/v20200301preview:CostAllocationRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'CostAllocationRule':
        """
        Get an existing CostAllocationRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return CostAllocationRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the rule. This is a read only value.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def properties(self) -> pulumi.Output['outputs.CostAllocationRulePropertiesResponse']:
        """
        Cost allocation rule properties
        """
        return pulumi.get(self, "properties")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Resource type of the rule. This is a read only value of Microsoft.CostManagement/CostAllocationRule.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

