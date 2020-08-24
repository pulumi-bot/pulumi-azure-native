# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs

__all__ = [
    'GetDatabaseAccountCassandraTableResult',
    'AwaitableGetDatabaseAccountCassandraTableResult',
    'get_database_account_cassandra_table',
]

@pulumi.output_type
class GetDatabaseAccountCassandraTableResult:
    """
    An Azure Cosmos DB Cassandra table.
    """
    def __init__(__self__, default_ttl=None, location=None, name=None, schema=None, tags=None, type=None):
        if default_ttl and not isinstance(default_ttl, float):
            raise TypeError("Expected argument 'default_ttl' to be a float")
        pulumi.set(__self__, "default_ttl", default_ttl)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if schema and not isinstance(schema, dict):
            raise TypeError("Expected argument 'schema' to be a dict")
        pulumi.set(__self__, "schema", schema)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="defaultTtl")
    def default_ttl(self) -> Optional[float]:
        """
        Time to live of the Cosmos DB Cassandra table
        """
        return pulumi.get(self, "default_ttl")

    @property
    @pulumi.getter
    def location(self) -> Optional[str]:
        """
        The location of the resource group to which the resource belongs.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the database account.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def schema(self) -> Optional['outputs.CassandraSchemaResponse']:
        """
        Schema of the Cosmos DB Cassandra table
        """
        return pulumi.get(self, "schema")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of Azure resource.
        """
        return pulumi.get(self, "type")


class AwaitableGetDatabaseAccountCassandraTableResult(GetDatabaseAccountCassandraTableResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatabaseAccountCassandraTableResult(
            default_ttl=self.default_ttl,
            location=self.location,
            name=self.name,
            schema=self.schema,
            tags=self.tags,
            type=self.type)


def get_database_account_cassandra_table(account_name: Optional[str] = None,
                                         keyspace_name: Optional[str] = None,
                                         name: Optional[str] = None,
                                         resource_group_name: Optional[str] = None,
                                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatabaseAccountCassandraTableResult:
    """
    Use this data source to access information about an existing resource.

    :param str account_name: Cosmos DB database account name.
    :param str keyspace_name: Cosmos DB keyspace name.
    :param str name: Cosmos DB table name.
    :param str resource_group_name: Name of an Azure resource group.
    """
    __args__ = dict()
    __args__['accountName'] = account_name
    __args__['keyspaceName'] = keyspace_name
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:documentdb/v20160331:getDatabaseAccountCassandraTable', __args__, opts=opts, typ=GetDatabaseAccountCassandraTableResult).value

    return AwaitableGetDatabaseAccountCassandraTableResult(
        default_ttl=__ret__.default_ttl,
        location=__ret__.location,
        name=__ret__.name,
        schema=__ret__.schema,
        tags=__ret__.tags,
        type=__ret__.type)
