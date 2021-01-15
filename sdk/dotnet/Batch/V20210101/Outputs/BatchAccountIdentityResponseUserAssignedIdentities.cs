// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Batch.V20210101.Outputs
{

    [OutputType]
    public sealed class BatchAccountIdentityResponseUserAssignedIdentities
    {
        /// <summary>
        /// The client id of user assigned identity.
        /// </summary>
        public readonly string ClientId;
        /// <summary>
        /// The principal id of user assigned identity.
        /// </summary>
        public readonly string PrincipalId;

        [OutputConstructor]
        private BatchAccountIdentityResponseUserAssignedIdentities(
            string clientId,

            string principalId)
        {
            ClientId = clientId;
            PrincipalId = principalId;
        }
    }
}
