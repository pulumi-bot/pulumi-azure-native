# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class GetFirewallPolicyRuleGroupResult:
    """
    Rule Group resource.
    """
    def __init__(__self__, etag=None, name=None, properties=None, type=None):
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        __self__.etag = etag
        """
        A unique read-only string that changes whenever the resource is updated.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        The name of the resource that is unique within a resource group. This name can be used to access the resource.
        """
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        __self__.properties = properties
        """
        The properties of the firewall policy rule group.
        """
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        __self__.type = type
        """
        Rule Group type.
        """


class AwaitableGetFirewallPolicyRuleGroupResult(GetFirewallPolicyRuleGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFirewallPolicyRuleGroupResult(
            etag=self.etag,
            name=self.name,
            properties=self.properties,
            type=self.type)


def get_firewall_policy_rule_group(firewall_policy_name=None, name=None, resource_group_name=None, opts=None):
    """
    Use this data source to access information about an existing resource.

    :param str firewall_policy_name: The name of the Firewall Policy.
    :param str name: The name of the FirewallPolicyRuleGroup.
    :param str resource_group_name: The name of the resource group.
    """
    __args__ = dict()
    __args__['firewallPolicyName'] = firewall_policy_name
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:network/v20190801:getFirewallPolicyRuleGroup', __args__, opts=opts).value

    return AwaitableGetFirewallPolicyRuleGroupResult(
        etag=__ret__.get('etag'),
        name=__ret__.get('name'),
        properties=__ret__.get('properties'),
        type=__ret__.get('type'))
