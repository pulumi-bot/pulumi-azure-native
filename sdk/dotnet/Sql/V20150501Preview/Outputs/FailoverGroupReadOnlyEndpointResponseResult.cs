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
    public sealed class FailoverGroupReadOnlyEndpointResponseResult
    {
        /// <summary>
        /// Failover policy of the read-only endpoint for the failover group.
        /// </summary>
        public readonly string? FailoverPolicy;

        [OutputConstructor]
        private FailoverGroupReadOnlyEndpointResponseResult(string? failoverPolicy)
        {
            FailoverPolicy = failoverPolicy;
        }
    }
}
