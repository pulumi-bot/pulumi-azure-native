// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Cache.Latest.Outputs
{

    [OutputType]
    public sealed class RedisInstanceDetailsResponseResult
    {
        /// <summary>
        /// Specifies whether the instance is a master node.
        /// </summary>
        public readonly bool IsMaster;
        /// <summary>
        /// If enableNonSslPort is true, provides Redis instance Non-SSL port.
        /// </summary>
        public readonly int NonSslPort;
        /// <summary>
        /// If clustering is enabled, the Shard ID of Redis Instance
        /// </summary>
        public readonly int ShardId;
        /// <summary>
        /// Redis instance SSL port.
        /// </summary>
        public readonly int SslPort;
        /// <summary>
        /// If the Cache uses availability zones, specifies availability zone where this instance is located.
        /// </summary>
        public readonly string Zone;

        [OutputConstructor]
        private RedisInstanceDetailsResponseResult(
            bool isMaster,

            int nonSslPort,

            int shardId,

            int sslPort,

            string zone)
        {
            IsMaster = isMaster;
            NonSslPort = nonSslPort;
            ShardId = shardId;
            SslPort = sslPort;
            Zone = zone;
        }
    }
}
