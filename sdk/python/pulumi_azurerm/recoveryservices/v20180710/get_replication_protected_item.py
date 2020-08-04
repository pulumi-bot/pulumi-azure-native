# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class GetReplicationProtectedItemResult:
    """
    Replication protected item.
    """
    def __init__(__self__, location=None, name=None, properties=None, type=None):
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        __self__.location = location
        """
        Resource Location
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        Resource Name
        """
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        __self__.properties = properties
        """
        The custom data.
        """
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        __self__.type = type
        """
        Resource Type
        """


class AwaitableGetReplicationProtectedItemResult(GetReplicationProtectedItemResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetReplicationProtectedItemResult(
            location=self.location,
            name=self.name,
            properties=self.properties,
            type=self.type)


def get_replication_protected_item(fabric_name=None, name=None, protection_container_name=None, resource_group_name=None, resource_name=None, opts=None):
    """
    Use this data source to access information about an existing resource.

    :param str fabric_name: Fabric unique name.
    :param str name: Replication protected item name.
    :param str protection_container_name: Protection container name.
    :param str resource_group_name: The name of the resource group where the recovery services vault is present.
    :param str resource_name: The name of the recovery services vault.
    """
    __args__ = dict()
    __args__['fabricName'] = fabric_name
    __args__['name'] = name
    __args__['protectionContainerName'] = protection_container_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['resourceName'] = resource_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:recoveryservices/v20180710:getReplicationProtectedItem', __args__, opts=opts).value

    return AwaitableGetReplicationProtectedItemResult(
        location=__ret__.get('location'),
        name=__ret__.get('name'),
        properties=__ret__.get('properties'),
        type=__ret__.get('type'))
