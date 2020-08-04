# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class Redis(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    The geo-location where the resource lives
    """
    name: pulumi.Output[str]
    """
    Resource name.
    """
    properties: pulumi.Output[dict]
    """
    Redis cache properties.
      * `access_keys` (`dict`) - The keys of the Redis cache - not set if this object is not the response to Create or Update redis cache
        * `primary_key` (`str`) - The current primary key that clients can use to authenticate with Redis cache.
        * `secondary_key` (`str`) - The current secondary key that clients can use to authenticate with Redis cache.

      * `enable_non_ssl_port` (`bool`) - Specifies whether the non-ssl Redis server port (6379) is enabled.
      * `host_name` (`str`) - Redis host name.
      * `linked_servers` (`dict`) - List of the linked servers associated with the cache
        * `value` (`list`) - List of linked server Ids of a Redis cache.
          * `id` (`str`) - Linked server Id.

      * `port` (`float`) - Redis non-SSL port.
      * `provisioning_state` (`str`) - Redis instance provisioning status.
      * `redis_configuration` (`dict`) - All Redis Settings. Few possible keys: rdb-backup-enabled,rdb-storage-connection-string,rdb-backup-frequency,maxmemory-delta,maxmemory-policy,notify-keyspace-events,maxmemory-samples,slowlog-log-slower-than,slowlog-max-len,list-max-ziplist-entries,list-max-ziplist-value,hash-max-ziplist-entries,hash-max-ziplist-value,set-max-intset-entries,zset-max-ziplist-entries,zset-max-ziplist-value etc.
      * `redis_version` (`str`) - Redis version.
      * `shard_count` (`float`) - The number of shards to be created on a Premium Cluster Cache.
      * `sku` (`dict`) - The SKU of the Redis cache to deploy.
        * `capacity` (`float`) - The size of the Redis cache to deploy. Valid values: for C (Basic/Standard) family (0, 1, 2, 3, 4, 5, 6), for P (Premium) family (1, 2, 3, 4).
        * `family` (`str`) - The SKU family to use. Valid values: (C, P). (C = Basic/Standard, P = Premium).
        * `name` (`str`) - The type of Redis cache to deploy. Valid values: (Basic, Standard, Premium)

      * `ssl_port` (`float`) - Redis SSL port.
      * `static_ip` (`str`) - Static IP address. Required when deploying a Redis cache inside an existing Azure Virtual Network.
      * `subnet_id` (`str`) - The full resource ID of a subnet in a virtual network to deploy the Redis cache in. Example format: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/Microsoft.{Network|ClassicNetwork}/VirtualNetworks/vnet1/subnets/subnet1
      * `tenant_settings` (`dict`) - tenantSettings
    """
    tags: pulumi.Output[dict]
    """
    Resource tags.
    """
    type: pulumi.Output[str]
    """
    Resource type.
    """
    def __init__(__self__, resource_name, opts=None, location=None, name=None, properties=None, resource_group_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        A single Redis item in List or Get Operation.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: The geo-location where the resource lives
        :param pulumi.Input[str] name: The name of the Redis cache.
        :param pulumi.Input[dict] properties: Redis cache properties.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[dict] tags: Resource tags.

        The **properties** object supports the following:

          * `enable_non_ssl_port` (`pulumi.Input[bool]`) - Specifies whether the non-ssl Redis server port (6379) is enabled.
          * `redis_configuration` (`pulumi.Input[dict]`) - All Redis Settings. Few possible keys: rdb-backup-enabled,rdb-storage-connection-string,rdb-backup-frequency,maxmemory-delta,maxmemory-policy,notify-keyspace-events,maxmemory-samples,slowlog-log-slower-than,slowlog-max-len,list-max-ziplist-entries,list-max-ziplist-value,hash-max-ziplist-entries,hash-max-ziplist-value,set-max-intset-entries,zset-max-ziplist-entries,zset-max-ziplist-value etc.
          * `shard_count` (`pulumi.Input[float]`) - The number of shards to be created on a Premium Cluster Cache.
          * `sku` (`pulumi.Input[dict]`) - The SKU of the Redis cache to deploy.
            * `capacity` (`pulumi.Input[float]`) - The size of the Redis cache to deploy. Valid values: for C (Basic/Standard) family (0, 1, 2, 3, 4, 5, 6), for P (Premium) family (1, 2, 3, 4).
            * `family` (`pulumi.Input[str]`) - The SKU family to use. Valid values: (C, P). (C = Basic/Standard, P = Premium).
            * `name` (`pulumi.Input[str]`) - The type of Redis cache to deploy. Valid values: (Basic, Standard, Premium)

          * `static_ip` (`pulumi.Input[str]`) - Static IP address. Required when deploying a Redis cache inside an existing Azure Virtual Network.
          * `subnet_id` (`pulumi.Input[str]`) - The full resource ID of a subnet in a virtual network to deploy the Redis cache in. Example format: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/Microsoft.{Network|ClassicNetwork}/VirtualNetworks/vnet1/subnets/subnet1
          * `tenant_settings` (`pulumi.Input[dict]`) - tenantSettings
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

            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if properties is None:
                raise TypeError("Missing required property 'properties'")
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['type'] = None
        super(Redis, __self__).__init__(
            'azurerm:cache/v20170201:Redis',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Redis resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Redis(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
