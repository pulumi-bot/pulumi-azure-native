// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Compute.V20201001Preview.Outputs
{

    [OutputType]
    public sealed class CloudServiceRoleProfileResponse
    {
        /// <summary>
        /// List of roles for the cloud service.
        /// </summary>
        public readonly ImmutableArray<Outputs.CloudServiceRoleProfilePropertiesResponse> Roles;

        [OutputConstructor]
        private CloudServiceRoleProfileResponse(ImmutableArray<Outputs.CloudServiceRoleProfilePropertiesResponse> roles)
        {
            Roles = roles;
        }
    }
}
