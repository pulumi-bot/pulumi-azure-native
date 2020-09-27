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
    public sealed class SqlStoredProcedureGetPropertiesResponseResource
    {
        /// <summary>
        /// Body of the Stored Procedure
        /// </summary>
        public readonly string? Body;
        /// <summary>
        /// A system generated property representing the resource etag required for optimistic concurrency control.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Name of the Cosmos DB SQL storedProcedure
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A system generated property. A unique identifier.
        /// </summary>
        public readonly string Rid;

        [OutputConstructor]
        private SqlStoredProcedureGetPropertiesResponseResource(
            string? body,

            string etag,

            string id,

            string rid)
        {
            Body = body;
            Etag = etag;
            Id = id;
            Rid = rid;
        }
    }
}
