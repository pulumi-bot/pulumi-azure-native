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
    public sealed class MongoDBDatabaseGetPropertiesResponseResource
    {
        /// <summary>
        /// A system generated property representing the resource etag required for optimistic concurrency control.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Name of the Cosmos DB MongoDB database
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A system generated property. A unique identifier.
        /// </summary>
        public readonly string Rid;
        /// <summary>
        /// A system generated property that denotes the last updated timestamp of the resource.
        /// </summary>
        public readonly double Ts;

        [OutputConstructor]
        private MongoDBDatabaseGetPropertiesResponseResource(
            string etag,

            string id,

            string rid,

            double ts)
        {
            Etag = etag;
            Id = id;
            Rid = rid;
            Ts = ts;
        }
    }
}
