# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class GetCassandraResourceCassandraKeyspaceResult:
    """
    An Azure Cosmos DB Cassandra keyspace.
    """
    def __init__(__self__, location=None, name=None, properties=None, tags=None, type=None):
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        __self__.location = location
        """
        The location of the resource group to which the resource belongs.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        The name of the ARM resource.
        """
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        __self__.properties = properties
        """
        The properties of an Azure Cosmos DB Cassandra keyspace
        """
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        """
        Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
        """
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        __self__.type = type
        """
        The type of Azure resource.
        """


class AwaitableGetCassandraResourceCassandraKeyspaceResult(GetCassandraResourceCassandraKeyspaceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCassandraResourceCassandraKeyspaceResult(
            location=self.location,
            name=self.name,
            properties=self.properties,
            tags=self.tags,
            type=self.type)


def get_cassandra_resource_cassandra_keyspace(account_name=None, name=None, resource_group_name=None, opts=None):
    """
    Use this data source to access information about an existing resource.

    :param str account_name: Cosmos DB database account name.
    :param str name: Cosmos DB keyspace name.
    :param str resource_group_name: Name of an Azure resource group.
    """
    __args__ = dict()
    __args__['accountName'] = account_name
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:documentdb/v20191212:getCassandraResourceCassandraKeyspace', __args__, opts=opts).value

    return AwaitableGetCassandraResourceCassandraKeyspaceResult(
        location=__ret__.get('location'),
        name=__ret__.get('name'),
        properties=__ret__.get('properties'),
        tags=__ret__.get('tags'),
        type=__ret__.get('type'))
