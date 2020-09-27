// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DocumentDB.V20200301.Outputs
{

    [OutputType]
    public sealed class GremlinDatabaseGetPropertiesResponseResource
    {
        /// <summary>
        /// A system generated property representing the resource etag required for optimistic concurrency control.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Name of the Cosmos DB Gremlin database
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A system generated property. A unique identifier.
        /// </summary>
        public readonly string Rid;

        [OutputConstructor]
        private GremlinDatabaseGetPropertiesResponseResource(
            string etag,

            string id,

            string rid)
        {
            Etag = etag;
            Id = id;
            Rid = rid;
        }
    }
}
