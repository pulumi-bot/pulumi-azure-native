// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.HDInsight.Outputs
{

    [OutputType]
    public sealed class ComputeProfileResponse
    {
        /// <summary>
        /// The list of roles in the cluster.
        /// </summary>
        public readonly ImmutableArray<Outputs.RoleResponse> Roles;

        [OutputConstructor]
        private ComputeProfileResponse(ImmutableArray<Outputs.RoleResponse> roles)
        {
            Roles = roles;
        }
    }
}