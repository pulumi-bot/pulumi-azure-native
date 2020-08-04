# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class GetDatabaseResult:
    """
    Class representing a Kusto database.
    """
    def __init__(__self__, kind=None, location=None, name=None, type=None):
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        __self__.kind = kind
        """
        Kind of the database
        """
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        __self__.location = location
        """
        Resource location.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        The name of the resource
        """
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        __self__.type = type
        """
        The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        """


class AwaitableGetDatabaseResult(GetDatabaseResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatabaseResult(
            kind=self.kind,
            location=self.location,
            name=self.name,
            type=self.type)


def get_database(cluster_name=None, name=None, resource_group_name=None, opts=None):
    """
    Use this data source to access information about an existing resource.

    :param str cluster_name: The name of the Kusto cluster.
    :param str name: The name of the database in the Kusto cluster.
    :param str resource_group_name: The name of the resource group containing the Kusto cluster.
    """
    __args__ = dict()
    __args__['clusterName'] = cluster_name
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:kusto/v20191109:getDatabase', __args__, opts=opts).value

    return AwaitableGetDatabaseResult(
        kind=__ret__.get('kind'),
        location=__ret__.get('location'),
        name=__ret__.get('name'),
        type=__ret__.get('type'))
