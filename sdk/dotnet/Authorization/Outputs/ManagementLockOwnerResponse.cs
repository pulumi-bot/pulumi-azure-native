// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Authorization.Outputs
{

    [OutputType]
    public sealed class ManagementLockOwnerResponse
    {
        /// <summary>
        /// The application ID of the lock owner.
        /// </summary>
        public readonly string? ApplicationId;

        [OutputConstructor]
        private ManagementLockOwnerResponse(string? applicationId)
        {
            ApplicationId = applicationId;
        }
    }
}