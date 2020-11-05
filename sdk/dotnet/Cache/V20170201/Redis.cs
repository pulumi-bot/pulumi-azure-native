// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Cache.V20170201
{
    /// <summary>
    /// A single Redis item in List or Get Operation.
    /// </summary>
    public partial class Redis : Pulumi.CustomResource
    {
        /// <summary>
        /// The keys of the Redis cache - not set if this object is not the response to Create or Update redis cache
        /// </summary>
        [Output("accessKeys")]
        public Output<Outputs.RedisAccessKeysResponse> AccessKeys { get; private set; } = null!;

        /// <summary>
        /// Specifies whether the non-ssl Redis server port (6379) is enabled.
        /// </summary>
        [Output("enableNonSslPort")]
        public Output<bool?> EnableNonSslPort { get; private set; } = null!;

        /// <summary>
        /// Redis host name.
        /// </summary>
        [Output("hostName")]
        public Output<string> HostName { get; private set; } = null!;

        /// <summary>
        /// List of the linked servers associated with the cache
        /// </summary>
        [Output("linkedServers")]
        public Output<Outputs.RedisLinkedServerListResponse> LinkedServers { get; private set; } = null!;

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Redis non-SSL port.
        /// </summary>
        [Output("port")]
        public Output<int> Port { get; private set; } = null!;

        /// <summary>
        /// Redis instance provisioning status.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// All Redis Settings. Few possible keys: rdb-backup-enabled,rdb-storage-connection-string,rdb-backup-frequency,maxmemory-delta,maxmemory-policy,notify-keyspace-events,maxmemory-samples,slowlog-log-slower-than,slowlog-max-len,list-max-ziplist-entries,list-max-ziplist-value,hash-max-ziplist-entries,hash-max-ziplist-value,set-max-intset-entries,zset-max-ziplist-entries,zset-max-ziplist-value etc.
        /// </summary>
        [Output("redisConfiguration")]
        public Output<ImmutableDictionary<string, string>?> RedisConfiguration { get; private set; } = null!;

        /// <summary>
        /// Redis version.
        /// </summary>
        [Output("redisVersion")]
        public Output<string> RedisVersion { get; private set; } = null!;

        /// <summary>
        /// The number of shards to be created on a Premium Cluster Cache.
        /// </summary>
        [Output("shardCount")]
        public Output<int?> ShardCount { get; private set; } = null!;

        /// <summary>
        /// The SKU of the Redis cache to deploy.
        /// </summary>
        [Output("sku")]
        public Output<Outputs.SkuResponse?> Sku { get; private set; } = null!;

        /// <summary>
        /// Redis SSL port.
        /// </summary>
        [Output("sslPort")]
        public Output<int> SslPort { get; private set; } = null!;

        /// <summary>
        /// Static IP address. Required when deploying a Redis cache inside an existing Azure Virtual Network.
        /// </summary>
        [Output("staticIP")]
        public Output<string?> StaticIP { get; private set; } = null!;

        /// <summary>
        /// The full resource ID of a subnet in a virtual network to deploy the Redis cache in. Example format: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/Microsoft.{Network|ClassicNetwork}/VirtualNetworks/vnet1/subnets/subnet1
        /// </summary>
        [Output("subnetId")]
        public Output<string?> SubnetId { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// tenantSettings
        /// </summary>
        [Output("tenantSettings")]
        public Output<ImmutableDictionary<string, string>?> TenantSettings { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Redis resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Redis(string name, RedisArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:cache/v20170201:Redis", name, args ?? new RedisArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Redis(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:cache/v20170201:Redis", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:cache/latest:Redis"},
                    new Pulumi.Alias { Type = "azure-nextgen:cache/v20150801:Redis"},
                    new Pulumi.Alias { Type = "azure-nextgen:cache/v20160401:Redis"},
                    new Pulumi.Alias { Type = "azure-nextgen:cache/v20171001:Redis"},
                    new Pulumi.Alias { Type = "azure-nextgen:cache/v20180301:Redis"},
                    new Pulumi.Alias { Type = "azure-nextgen:cache/v20190701:Redis"},
                    new Pulumi.Alias { Type = "azure-nextgen:cache/v20200601:Redis"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Redis resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Redis Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Redis(name, id, options);
        }
    }

    public sealed class RedisArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether the non-ssl Redis server port (6379) is enabled.
        /// </summary>
        [Input("enableNonSslPort")]
        public Input<bool>? EnableNonSslPort { get; set; }

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The name of the Redis cache.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("redisConfiguration")]
        private InputMap<string>? _redisConfiguration;

        /// <summary>
        /// All Redis Settings. Few possible keys: rdb-backup-enabled,rdb-storage-connection-string,rdb-backup-frequency,maxmemory-delta,maxmemory-policy,notify-keyspace-events,maxmemory-samples,slowlog-log-slower-than,slowlog-max-len,list-max-ziplist-entries,list-max-ziplist-value,hash-max-ziplist-entries,hash-max-ziplist-value,set-max-intset-entries,zset-max-ziplist-entries,zset-max-ziplist-value etc.
        /// </summary>
        public InputMap<string> RedisConfiguration
        {
            get => _redisConfiguration ?? (_redisConfiguration = new InputMap<string>());
            set => _redisConfiguration = value;
        }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The number of shards to be created on a Premium Cluster Cache.
        /// </summary>
        [Input("shardCount")]
        public Input<int>? ShardCount { get; set; }

        /// <summary>
        /// The SKU of the Redis cache to deploy.
        /// </summary>
        [Input("sku", required: true)]
        public Input<Inputs.SkuArgs> Sku { get; set; } = null!;

        /// <summary>
        /// Static IP address. Required when deploying a Redis cache inside an existing Azure Virtual Network.
        /// </summary>
        [Input("staticIP")]
        public Input<string>? StaticIP { get; set; }

        /// <summary>
        /// The full resource ID of a subnet in a virtual network to deploy the Redis cache in. Example format: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/Microsoft.{Network|ClassicNetwork}/VirtualNetworks/vnet1/subnets/subnet1
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tenantSettings")]
        private InputMap<string>? _tenantSettings;

        /// <summary>
        /// tenantSettings
        /// </summary>
        public InputMap<string> TenantSettings
        {
            get => _tenantSettings ?? (_tenantSettings = new InputMap<string>());
            set => _tenantSettings = value;
        }

        public RedisArgs()
        {
        }
    }
}
