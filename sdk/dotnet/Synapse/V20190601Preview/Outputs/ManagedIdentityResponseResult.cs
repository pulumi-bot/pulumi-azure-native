// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Synapse.V20190601Preview.Outputs
{

    [OutputType]
    public sealed class ManagedIdentityResponseResult
    {
        /// <summary>
        /// The principal ID of the workspace managed identity
        /// </summary>
        public readonly string PrincipalId;
        /// <summary>
        /// The tenant ID of the workspace managed identity
        /// </summary>
        public readonly string TenantId;
        /// <summary>
        /// The type of managed identity for the workspace
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private ManagedIdentityResponseResult(
            string principalId,

            string tenantId,

            string? type)
        {
            PrincipalId = principalId;
            TenantId = tenantId;
            Type = type;
        }
    }
}
