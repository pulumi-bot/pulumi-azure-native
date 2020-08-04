# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class GetServerKeyResult:
    """
    A PostgreSQL Server key.
    """
    def __init__(__self__, kind=None, name=None, properties=None, type=None):
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        __self__.kind = kind
        """
        Kind of encryption protector used to protect the key.
        """
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
        Properties of the ServerKey Resource.
        """
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        __self__.type = type
        """
        The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        """


class AwaitableGetServerKeyResult(GetServerKeyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServerKeyResult(
            kind=self.kind,
            name=self.name,
            properties=self.properties,
            type=self.type)


def get_server_key(name=None, resource_group_name=None, server_name=None, opts=None):
    """
    Use this data source to access information about an existing resource.

    :param str name: The name of the PostgreSQL Server key to be retrieved.
    :param str resource_group_name: The name of the resource group. The name is case insensitive.
    :param str server_name: The name of the server.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    __args__['serverName'] = server_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:dbforpostgresql/v20200101:getServerKey', __args__, opts=opts).value

    return AwaitableGetServerKeyResult(
        kind=__ret__.get('kind'),
        name=__ret__.get('name'),
        properties=__ret__.get('properties'),
        type=__ret__.get('type'))
