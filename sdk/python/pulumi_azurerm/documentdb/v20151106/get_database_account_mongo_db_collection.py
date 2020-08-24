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
    'GetDatabaseAccountMongoDBCollectionResult',
    'AwaitableGetDatabaseAccountMongoDBCollectionResult',
    'get_database_account_mongo_db_collection',
]

@pulumi.output_type
class GetDatabaseAccountMongoDBCollectionResult:
    """
    An Azure Cosmos DB MongoDB collection.
    """
    def __init__(__self__, indexes=None, location=None, name=None, shard_key=None, tags=None, type=None):
        if indexes and not isinstance(indexes, list):
            raise TypeError("Expected argument 'indexes' to be a list")
        pulumi.set(__self__, "indexes", indexes)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if shard_key and not isinstance(shard_key, dict):
            raise TypeError("Expected argument 'shard_key' to be a dict")
        pulumi.set(__self__, "shard_key", shard_key)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def indexes(self) -> Optional[List['outputs.MongoIndexResponse']]:
        """
        List of index keys
        """
        return pulumi.get(self, "indexes")

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
    @pulumi.getter(name="shardKey")
    def shard_key(self) -> Optional[Mapping[str, str]]:
        """
        A key-value pair of shard keys to be applied for the request.
        """
        return pulumi.get(self, "shard_key")

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


class AwaitableGetDatabaseAccountMongoDBCollectionResult(GetDatabaseAccountMongoDBCollectionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatabaseAccountMongoDBCollectionResult(
            indexes=self.indexes,
            location=self.location,
            name=self.name,
            shard_key=self.shard_key,
            tags=self.tags,
            type=self.type)


def get_database_account_mongo_db_collection(account_name: Optional[str] = None,
                                             database_name: Optional[str] = None,
                                             name: Optional[str] = None,
                                             resource_group_name: Optional[str] = None,
                                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatabaseAccountMongoDBCollectionResult:
    """
    Use this data source to access information about an existing resource.

    :param str account_name: Cosmos DB database account name.
    :param str database_name: Cosmos DB database name.
    :param str name: Cosmos DB collection name.
    :param str resource_group_name: Name of an Azure resource group.
    """
    __args__ = dict()
    __args__['accountName'] = account_name
    __args__['databaseName'] = database_name
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:documentdb/v20151106:getDatabaseAccountMongoDBCollection', __args__, opts=opts, typ=GetDatabaseAccountMongoDBCollectionResult).value

    return AwaitableGetDatabaseAccountMongoDBCollectionResult(
        indexes=__ret__.indexes,
        location=__ret__.location,
        name=__ret__.name,
        shard_key=__ret__.shard_key,
        tags=__ret__.tags,
        type=__ret__.type)
