// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.V20160430Preview.Outputs
{

    [OutputType]
    public sealed class VirtualMachineIdentityResponseResult
    {
        /// <summary>
        /// The principal id of virtual machine identity.
        /// </summary>
        public readonly string PrincipalId;
        /// <summary>
        /// The tenant id associated with the virtual machine.
        /// </summary>
        public readonly string TenantId;
        /// <summary>
        /// The type of identity used for the virtual machine. Currently, the only supported type is 'SystemAssigned', which implicitly creates an identity.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private VirtualMachineIdentityResponseResult(
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
