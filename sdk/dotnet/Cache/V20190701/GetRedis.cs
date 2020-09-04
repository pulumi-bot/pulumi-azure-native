// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Cache.V20190701
{
    public static class GetRedis
    {
        public static Task<GetRedisResult> InvokeAsync(GetRedisArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRedisResult>("azurerm:cache/v20190701:getRedis", args ?? new GetRedisArgs(), options.WithVersion());
    }


    public sealed class GetRedisArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Redis cache.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetRedisArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRedisResult
    {
        /// <summary>
        /// The keys of the Redis cache - not set if this object is not the response to Create or Update redis cache
        /// </summary>
        public readonly Outputs.RedisAccessKeysResponseResult AccessKeys;
        /// <summary>
        /// Specifies whether the non-ssl Redis server port (6379) is enabled.
        /// </summary>
        public readonly bool? EnableNonSslPort;
        /// <summary>
        /// Redis host name.
        /// </summary>
        public readonly string HostName;
        /// <summary>
        /// List of the Redis instances associated with the cache
        /// </summary>
        public readonly ImmutableArray<Outputs.RedisInstanceDetailsResponseResult> Instances;
        /// <summary>
        /// List of the linked servers associated with the cache
        /// </summary>
        public readonly ImmutableArray<Outputs.RedisLinkedServerResponseResult> LinkedServers;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Optional: requires clients to use a specified TLS version (or higher) to connect (e,g, '1.0', '1.1', '1.2')
        /// </summary>
        public readonly string? MinimumTlsVersion;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Redis non-SSL port.
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// Redis instance provisioning status.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// All Redis Settings. Few possible keys: rdb-backup-enabled,rdb-storage-connection-string,rdb-backup-frequency,maxmemory-delta,maxmemory-policy,notify-keyspace-events,maxmemory-samples,slowlog-log-slower-than,slowlog-max-len,list-max-ziplist-entries,list-max-ziplist-value,hash-max-ziplist-entries,hash-max-ziplist-value,set-max-intset-entries,zset-max-ziplist-entries,zset-max-ziplist-value etc.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? RedisConfiguration;
        /// <summary>
        /// Redis version.
        /// </summary>
        public readonly string RedisVersion;
        /// <summary>
        /// The number of replicas to be created per master.
        /// </summary>
        public readonly int? ReplicasPerMaster;
        /// <summary>
        /// The number of shards to be created on a Premium Cluster Cache.
        /// </summary>
        public readonly int? ShardCount;
        /// <summary>
        /// The SKU of the Redis cache to deploy.
        /// </summary>
        public readonly Outputs.SkuResponseResult Sku;
        /// <summary>
        /// Redis SSL port.
        /// </summary>
        public readonly int SslPort;
        /// <summary>
        /// Static IP address. Required when deploying a Redis cache inside an existing Azure Virtual Network.
        /// </summary>
        public readonly string? StaticIP;
        /// <summary>
        /// The full resource ID of a subnet in a virtual network to deploy the Redis cache in. Example format: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/Microsoft.{Network|ClassicNetwork}/VirtualNetworks/vnet1/subnets/subnet1
        /// </summary>
        public readonly string? SubnetId;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// A dictionary of tenant settings
        /// </summary>
        public readonly ImmutableDictionary<string, string>? TenantSettings;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// A list of availability zones denoting where the resource needs to come from.
        /// </summary>
        public readonly ImmutableArray<string> Zones;

        [OutputConstructor]
        private GetRedisResult(
            Outputs.RedisAccessKeysResponseResult accessKeys,

            bool? enableNonSslPort,

            string hostName,

            ImmutableArray<Outputs.RedisInstanceDetailsResponseResult> instances,

            ImmutableArray<Outputs.RedisLinkedServerResponseResult> linkedServers,

            string location,

            string? minimumTlsVersion,

            string name,

            int port,

            string provisioningState,

            ImmutableDictionary<string, string>? redisConfiguration,

            string redisVersion,

            int? replicasPerMaster,

            int? shardCount,

            Outputs.SkuResponseResult sku,

            int sslPort,

            string? staticIP,

            string? subnetId,

            ImmutableDictionary<string, string>? tags,

            ImmutableDictionary<string, string>? tenantSettings,

            string type,

            ImmutableArray<string> zones)
        {
            AccessKeys = accessKeys;
            EnableNonSslPort = enableNonSslPort;
            HostName = hostName;
            Instances = instances;
            LinkedServers = linkedServers;
            Location = location;
            MinimumTlsVersion = minimumTlsVersion;
            Name = name;
            Port = port;
            ProvisioningState = provisioningState;
            RedisConfiguration = redisConfiguration;
            RedisVersion = redisVersion;
            ReplicasPerMaster = replicasPerMaster;
            ShardCount = shardCount;
            Sku = sku;
            SslPort = sslPort;
            StaticIP = staticIP;
            SubnetId = subnetId;
            Tags = tags;
            TenantSettings = tenantSettings;
            Type = type;
            Zones = zones;
        }
    }
}
