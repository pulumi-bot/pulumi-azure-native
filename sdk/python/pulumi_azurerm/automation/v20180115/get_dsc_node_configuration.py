# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class GetDscNodeConfigurationResult:
    """
    Definition of the dsc node configuration.
    """
    def __init__(__self__, name=None, properties=None, type=None):
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        The name of the resource
        """
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        __self__.properties = properties
        """
        Gets or sets the configuration properties.
        """
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        __self__.type = type
        """
        The type of the resource.
        """


class AwaitableGetDscNodeConfigurationResult(GetDscNodeConfigurationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDscNodeConfigurationResult(
            name=self.name,
            properties=self.properties,
            type=self.type)


def get_dsc_node_configuration(automation_account_name=None, name=None, resource_group_name=None, opts=None):
    """
    Use this data source to access information about an existing resource.

    :param str automation_account_name: The name of the automation account.
    :param str name: The Dsc node configuration name.
    :param str resource_group_name: Name of an Azure Resource group.
    """
    __args__ = dict()
    __args__['automationAccountName'] = automation_account_name
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:automation/v20180115:getDscNodeConfiguration', __args__, opts=opts).value

    return AwaitableGetDscNodeConfigurationResult(
        name=__ret__.get('name'),
        properties=__ret__.get('properties'),
        type=__ret__.get('type'))
