// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DocumentDB.Latest.Outputs
{

    [OutputType]
    public sealed class MongoDBCollectionGetPropertiesResponseResource
    {
        /// <summary>
        /// Analytical TTL.
        /// </summary>
        public readonly int? AnalyticalStorageTtl;
        /// <summary>
        /// A system generated property representing the resource etag required for optimistic concurrency control.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Name of the Cosmos DB MongoDB collection
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// List of index keys
        /// </summary>
        public readonly ImmutableArray<Outputs.MongoIndexResponse> Indexes;
        /// <summary>
        /// A system generated property. A unique identifier.
        /// </summary>
        public readonly string Rid;
        /// <summary>
        /// A key-value pair of shard keys to be applied for the request.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? ShardKey;
        /// <summary>
        /// A system generated property that denotes the last updated timestamp of the resource.
        /// </summary>
        public readonly double Ts;

        [OutputConstructor]
        private MongoDBCollectionGetPropertiesResponseResource(
            int? analyticalStorageTtl,

            string etag,

            string id,

            ImmutableArray<Outputs.MongoIndexResponse> indexes,

            string rid,

            ImmutableDictionary<string, string>? shardKey,

            double ts)
        {
            AnalyticalStorageTtl = analyticalStorageTtl;
            Etag = etag;
            Id = id;
            Indexes = indexes;
            Rid = rid;
            ShardKey = shardKey;
            Ts = ts;
        }
    }
}
