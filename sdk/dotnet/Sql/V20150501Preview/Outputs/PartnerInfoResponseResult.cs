// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Sql.V20150501Preview.Outputs
{

    [OutputType]
    public sealed class PartnerInfoResponseResult
    {
        /// <summary>
        /// Resource identifier of the partner server.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Geo location of the partner server.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Replication role of the partner server.
        /// </summary>
        public readonly string ReplicationRole;

        [OutputConstructor]
        private PartnerInfoResponseResult(
            string id,

            string location,

            string replicationRole)
        {
            Id = id;
            Location = location;
            ReplicationRole = replicationRole;
        }
    }
}
