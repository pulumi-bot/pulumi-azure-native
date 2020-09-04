// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Resources.V20191001Preview.Outputs
{

    [OutputType]
    public sealed class ManagedServiceIdentityResponseResult
    {
        /// <summary>
        /// Type of the managed identity.
        /// </summary>
        public readonly string? Type;
        /// <summary>
        /// The list of user-assigned managed identities associated with the resource. Key is the Azure resource Id of the managed identity.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.UserAssignedIdentityResponseResult>? UserAssignedIdentities;

        [OutputConstructor]
        private ManagedServiceIdentityResponseResult(
            string? type,

            ImmutableDictionary<string, Outputs.UserAssignedIdentityResponseResult>? userAssignedIdentities)
        {
            Type = type;
            UserAssignedIdentities = userAssignedIdentities;
        }
    }
}
