# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class FirewallPolicyRuleGroup(pulumi.CustomResource):
    etag: pulumi.Output[str]
    """
    A unique read-only string that changes whenever the resource is updated.
    """
    name: pulumi.Output[str]
    """
    The name of the resource that is unique within a resource group. This name can be used to access the resource.
    """
    properties: pulumi.Output[dict]
    """
    The properties of the firewall policy rule group.
      * `priority` (`float`) - Priority of the Firewall Policy Rule Group resource.
      * `provisioning_state` (`str`) - The provisioning state of the firewall policy rule group resource.
      * `rules` (`list`) - Group of Firewall Policy rules.
        * `name` (`str`) - The name of the rule.
        * `priority` (`float`) - Priority of the Firewall Policy Rule resource.
        * `rule_type` (`str`) - The type of the rule.
    """
    type: pulumi.Output[str]
    """
    Rule Group type.
    """
    def __init__(__self__, resource_name, opts=None, firewall_policy_name=None, id=None, name=None, properties=None, resource_group_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Rule Group resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] firewall_policy_name: The name of the Firewall Policy.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[str] name: The name of the FirewallPolicyRuleGroup.
        :param pulumi.Input[dict] properties: The properties of the firewall policy rule group.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.

        The **properties** object supports the following:

          * `priority` (`pulumi.Input[float]`) - Priority of the Firewall Policy Rule Group resource.
          * `rules` (`pulumi.Input[list]`) - Group of Firewall Policy rules.
            * `name` (`pulumi.Input[str]`) - The name of the rule.
            * `priority` (`pulumi.Input[float]`) - Priority of the Firewall Policy Rule resource.
            * `rule_type` (`pulumi.Input[str]`) - The type of the rule.
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
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['etag'] = None
            __props__['type'] = None
        super(FirewallPolicyRuleGroup, __self__).__init__(
            'azurerm:network/v20200401:FirewallPolicyRuleGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing FirewallPolicyRuleGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return FirewallPolicyRuleGroup(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
