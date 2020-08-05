# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class DatabaseAccountGremlinGraph(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    The location of the resource group to which the resource belongs.
    """
    name: pulumi.Output[str]
    """
    The name of the database account.
    """
    properties: pulumi.Output[dict]
    """
    The properties of an Azure Cosmos DB Gremlin graph
      * `_etag` (`str`) - A system generated property representing the resource etag required for optimistic concurrency control.
      * `_rid` (`str`) - A system generated property. A unique identifier.
      * `_ts` (`dict`) - A system generated property that denotes the last updated timestamp of the resource.
      * `conflict_resolution_policy` (`dict`) - The conflict resolution policy for the graph.
        * `conflict_resolution_path` (`str`) - The conflict resolution path in the case of LastWriterWins mode.
        * `conflict_resolution_procedure` (`str`) - The procedure to resolve conflicts in the case of custom mode.
        * `mode` (`str`) - Indicates the conflict resolution mode.

      * `default_ttl` (`float`) - Default time to live
      * `id` (`str`) - Name of the Cosmos DB Gremlin graph
      * `indexing_policy` (`dict`) - The configuration of the indexing policy. By default, the indexing is automatic for all document paths within the graph
        * `automatic` (`bool`) - Indicates if the indexing policy is automatic
        * `excluded_paths` (`list`) - List of paths to exclude from indexing
          * `path` (`str`) - The path for which the indexing behavior applies to. Index paths typically start with root and end with wildcard (/path/*)

        * `included_paths` (`list`) - List of paths to include in the indexing
          * `indexes` (`list`) - List of indexes for this path
            * `data_type` (`str`) - The datatype for which the indexing behavior is applied to.
            * `kind` (`str`) - Indicates the type of index.
            * `precision` (`float`) - The precision of the index. -1 is maximum precision.

          * `path` (`str`) - The path for which the indexing behavior applies to. Index paths typically start with root and end with wildcard (/path/*)

        * `indexing_mode` (`str`) - Indicates the indexing mode.

      * `partition_key` (`dict`) - The configuration of the partition key to be used for partitioning data into multiple partitions
        * `kind` (`str`) - Indicates the kind of algorithm used for partitioning
        * `paths` (`list`) - List of paths using which data within the container can be partitioned

      * `unique_key_policy` (`dict`) - The unique key policy configuration for specifying uniqueness constraints on documents in the collection in the Azure Cosmos DB service.
        * `unique_keys` (`list`) - List of unique keys on that enforces uniqueness constraint on documents in the collection in the Azure Cosmos DB service.
          * `paths` (`list`) - List of paths must be unique for each document in the Azure Cosmos DB service
    """
    tags: pulumi.Output[dict]
    """
    Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
    """
    type: pulumi.Output[str]
    """
    The type of Azure resource.
    """
    def __init__(__self__, resource_name, opts=None, account_name=None, database_name=None, name=None, options=None, resource=None, resource_group_name=None, __props__=None, __name__=None, __opts__=None):
        """
        An Azure Cosmos DB Gremlin graph.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: Cosmos DB database account name.
        :param pulumi.Input[str] database_name: Cosmos DB database name.
        :param pulumi.Input[str] name: Cosmos DB graph name.
        :param pulumi.Input[dict] options: A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
        :param pulumi.Input[dict] resource: The standard JSON format of a Gremlin graph
        :param pulumi.Input[str] resource_group_name: Name of an Azure resource group.

        The **resource** object supports the following:

          * `conflict_resolution_policy` (`pulumi.Input[dict]`) - The conflict resolution policy for the graph.
            * `conflict_resolution_path` (`pulumi.Input[str]`) - The conflict resolution path in the case of LastWriterWins mode.
            * `conflict_resolution_procedure` (`pulumi.Input[str]`) - The procedure to resolve conflicts in the case of custom mode.
            * `mode` (`pulumi.Input[str]`) - Indicates the conflict resolution mode.

          * `default_ttl` (`pulumi.Input[float]`) - Default time to live
          * `id` (`pulumi.Input[str]`) - Name of the Cosmos DB Gremlin graph
          * `indexing_policy` (`pulumi.Input[dict]`) - The configuration of the indexing policy. By default, the indexing is automatic for all document paths within the graph
            * `automatic` (`pulumi.Input[bool]`) - Indicates if the indexing policy is automatic
            * `excluded_paths` (`pulumi.Input[list]`) - List of paths to exclude from indexing
              * `path` (`pulumi.Input[str]`) - The path for which the indexing behavior applies to. Index paths typically start with root and end with wildcard (/path/*)

            * `included_paths` (`pulumi.Input[list]`) - List of paths to include in the indexing
              * `indexes` (`pulumi.Input[list]`) - List of indexes for this path
                * `data_type` (`pulumi.Input[str]`) - The datatype for which the indexing behavior is applied to.
                * `kind` (`pulumi.Input[str]`) - Indicates the type of index.
                * `precision` (`pulumi.Input[float]`) - The precision of the index. -1 is maximum precision.

              * `path` (`pulumi.Input[str]`) - The path for which the indexing behavior applies to. Index paths typically start with root and end with wildcard (/path/*)

            * `indexing_mode` (`pulumi.Input[str]`) - Indicates the indexing mode.

          * `partition_key` (`pulumi.Input[dict]`) - The configuration of the partition key to be used for partitioning data into multiple partitions
            * `kind` (`pulumi.Input[str]`) - Indicates the kind of algorithm used for partitioning
            * `paths` (`pulumi.Input[list]`) - List of paths using which data within the container can be partitioned

          * `unique_key_policy` (`pulumi.Input[dict]`) - The unique key policy configuration for specifying uniqueness constraints on documents in the collection in the Azure Cosmos DB service.
            * `unique_keys` (`pulumi.Input[list]`) - List of unique keys on that enforces uniqueness constraint on documents in the collection in the Azure Cosmos DB service.
              * `paths` (`pulumi.Input[list]`) - List of paths must be unique for each document in the Azure Cosmos DB service
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

            if account_name is None:
                raise TypeError("Missing required property 'account_name'")
            __props__['account_name'] = account_name
            if database_name is None:
                raise TypeError("Missing required property 'database_name'")
            __props__['database_name'] = database_name
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if options is None:
                raise TypeError("Missing required property 'options'")
            __props__['options'] = options
            if resource is None:
                raise TypeError("Missing required property 'resource'")
            __props__['resource'] = resource
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['location'] = None
            __props__['properties'] = None
            __props__['tags'] = None
            __props__['type'] = None
        super(DatabaseAccountGremlinGraph, __self__).__init__(
            'azurerm:documentdb/v20150408:DatabaseAccountGremlinGraph',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing DatabaseAccountGremlinGraph resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return DatabaseAccountGremlinGraph(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
