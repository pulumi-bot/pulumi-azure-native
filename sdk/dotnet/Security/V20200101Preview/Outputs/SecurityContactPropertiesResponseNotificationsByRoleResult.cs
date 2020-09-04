// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Security.V20200101Preview.Outputs
{

    [OutputType]
    public sealed class SecurityContactPropertiesResponseNotificationsByRoleResult
    {
        /// <summary>
        /// Defines which RBAC roles will get email notifications from Azure Security Center. List of allowed RBAC roles: 
        /// </summary>
        public readonly ImmutableArray<string> Roles;
        /// <summary>
        /// Defines whether to send email notifications from Azure Security Center to persons with specific RBAC roles on the subscription.
        /// </summary>
        public readonly string? State;

        [OutputConstructor]
        private SecurityContactPropertiesResponseNotificationsByRoleResult(
            ImmutableArray<string> roles,

            string? state)
        {
            Roles = roles;
            State = state;
        }
    }
}
