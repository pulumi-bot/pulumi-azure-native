# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from ._inputs import *

__all__ = ['DatabaseAccountMongoDBDatabase']


class DatabaseAccountMongoDBDatabase(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_name: Optional[pulumi.Input[str]] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 options: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 resource: Optional[pulumi.Input[pulumi.InputType['MongoDBDatabaseResourceArgs']]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        An Azure Cosmos DB MongoDB database.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: Cosmos DB database account name.
        :param pulumi.Input[str] database_name: Cosmos DB database name.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] options: A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
        :param pulumi.Input[pulumi.InputType['MongoDBDatabaseResourceArgs']] resource: The standard JSON format of a MongoDB database
        :param pulumi.Input[str] resource_group_name: Name of an Azure resource group.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if account_name is None and not opts.urn:
                raise TypeError("Missing required property 'account_name'")
            __props__['account_name'] = account_name
            __props__['database_name'] = database_name
            if options is None and not opts.urn:
                raise TypeError("Missing required property 'options'")
            __props__['options'] = options
            if resource is None and not opts.urn:
                raise TypeError("Missing required property 'resource'")
            __props__['resource'] = resource
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['location'] = None
            __props__['name'] = None
            __props__['tags'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:documentdb/v20151106:DatabaseAccountMongoDBDatabase"), pulumi.Alias(type_="azure-native:documentdb:DatabaseAccountMongoDBDatabase"), pulumi.Alias(type_="azure-nextgen:documentdb:DatabaseAccountMongoDBDatabase"), pulumi.Alias(type_="azure-native:documentdb/latest:DatabaseAccountMongoDBDatabase"), pulumi.Alias(type_="azure-nextgen:documentdb/latest:DatabaseAccountMongoDBDatabase"), pulumi.Alias(type_="azure-native:documentdb/v20150401:DatabaseAccountMongoDBDatabase"), pulumi.Alias(type_="azure-nextgen:documentdb/v20150401:DatabaseAccountMongoDBDatabase"), pulumi.Alias(type_="azure-native:documentdb/v20150408:DatabaseAccountMongoDBDatabase"), pulumi.Alias(type_="azure-nextgen:documentdb/v20150408:DatabaseAccountMongoDBDatabase"), pulumi.Alias(type_="azure-native:documentdb/v20160319:DatabaseAccountMongoDBDatabase"), pulumi.Alias(type_="azure-nextgen:documentdb/v20160319:DatabaseAccountMongoDBDatabase"), pulumi.Alias(type_="azure-native:documentdb/v20160331:DatabaseAccountMongoDBDatabase"), pulumi.Alias(type_="azure-nextgen:documentdb/v20160331:DatabaseAccountMongoDBDatabase")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(DatabaseAccountMongoDBDatabase, __self__).__init__(
            'azure-native:documentdb/v20151106:DatabaseAccountMongoDBDatabase',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DatabaseAccountMongoDBDatabase':
        """
        Get an existing DatabaseAccountMongoDBDatabase resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["location"] = None
        __props__["name"] = None
        __props__["tags"] = None
        __props__["type"] = None
        return DatabaseAccountMongoDBDatabase(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[Optional[str]]:
        """
        The location of the resource group to which the resource belongs.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the database account.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of Azure resource.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

