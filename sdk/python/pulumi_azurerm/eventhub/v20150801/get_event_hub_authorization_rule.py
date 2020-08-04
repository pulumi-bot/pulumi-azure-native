# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class GetEventHubAuthorizationRuleResult:
    """
    Single item in a List or Get AuthorizationRule operation
    """
    def __init__(__self__, location=None, name=None, properties=None, type=None):
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        __self__.location = location
        """
        Resource location
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        Resource name
        """
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        __self__.properties = properties
        """
        Properties supplied to create or update SharedAccessAuthorizationRule
        """
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        __self__.type = type
        """
        Resource type
        """


class AwaitableGetEventHubAuthorizationRuleResult(GetEventHubAuthorizationRuleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEventHubAuthorizationRuleResult(
            location=self.location,
            name=self.name,
            properties=self.properties,
            type=self.type)


def get_event_hub_authorization_rule(event_hub_name=None, name=None, namespace_name=None, resource_group_name=None, opts=None):
    """
    Use this data source to access information about an existing resource.

    :param str event_hub_name: The Event Hub name
    :param str name: The authorization rule name.
    :param str namespace_name: The Namespace name
    :param str resource_group_name: Name of the resource group within the azure subscription.
    """
    __args__ = dict()
    __args__['eventHubName'] = event_hub_name
    __args__['name'] = name
    __args__['namespaceName'] = namespace_name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:eventhub/v20150801:getEventHubAuthorizationRule', __args__, opts=opts).value

    return AwaitableGetEventHubAuthorizationRuleResult(
        location=__ret__.get('location'),
        name=__ret__.get('name'),
        properties=__ret__.get('properties'),
        type=__ret__.get('type'))
