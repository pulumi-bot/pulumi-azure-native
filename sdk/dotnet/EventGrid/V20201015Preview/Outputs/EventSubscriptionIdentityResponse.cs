// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.EventGrid.V20201015Preview.Outputs
{

    [OutputType]
    public sealed class EventSubscriptionIdentityResponse
    {
        /// <summary>
        /// The type of managed identity used. The type 'SystemAssigned, UserAssigned' includes both an implicitly created identity and a set of user-assigned identities. The type 'None' will remove any identity.
        /// </summary>
        public readonly string? Type;
        /// <summary>
        /// The user identity associated with the resource.
        /// </summary>
        public readonly string? UserAssignedIdentity;

        [OutputConstructor]
        private EventSubscriptionIdentityResponse(
            string? type,

            string? userAssignedIdentity)
        {
            Type = type;
            UserAssignedIdentity = userAssignedIdentity;
        }
    }
}
