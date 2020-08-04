# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class GetHybridConnectionAuthorizationRuleResult:
    """
    Description of a namespace authorization rule.
    """
    def __init__(__self__, name=None, properties=None, type=None):
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        Resource name.
        """
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        __self__.properties = properties
        """
        Authorization rule properties.
        """
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        __self__.type = type
        """
        Resource type.
        """


class AwaitableGetHybridConnectionAuthorizationRuleResult(GetHybridConnectionAuthorizationRuleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetHybridConnectionAuthorizationRuleResult(
            name=self.name,
            properties=self.properties,
            type=self.type)


def get_hybrid_connection_authorization_rule(hybrid_connection_name=None, name=None, namespace_name=None, resource_group_name=None, opts=None):
    """
    Use this data source to access information about an existing resource.

    :param str hybrid_connection_name: The hybrid connection name.
    :param str name: The authorization rule name.
    :param str namespace_name: The namespace name
    :param str resource_group_name: Name of the Resource group within the Azure subscription.
    """
    __args__ = dict()
    __args__['hybridConnectionName'] = hybrid_connection_name
    __args__['name'] = name
    __args__['namespaceName'] = namespace_name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:relay/v20170401:getHybridConnectionAuthorizationRule', __args__, opts=opts).value

    return AwaitableGetHybridConnectionAuthorizationRuleResult(
        name=__ret__.get('name'),
        properties=__ret__.get('properties'),
        type=__ret__.get('type'))
