// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.ManagedIdentity.V20181130
{
    public static class GetUserAssignedIdentity
    {
        public static Task<GetUserAssignedIdentityResult> InvokeAsync(GetUserAssignedIdentityArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetUserAssignedIdentityResult>("azure-nextgen:managedidentity/v20181130:getUserAssignedIdentity", args ?? new GetUserAssignedIdentityArgs(), options.WithVersion());
    }


    public sealed class GetUserAssignedIdentityArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Resource Group to which the identity belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the identity resource.
        /// </summary>
        [Input("resourceName", required: true)]
        public string ResourceName { get; set; } = null!;

        public GetUserAssignedIdentityArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetUserAssignedIdentityResult
    {
        /// <summary>
        /// The id of the app associated with the identity. This is a random generated UUID by MSI.
        /// </summary>
        public readonly string ClientId;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The id of the service principal object associated with the created identity.
        /// </summary>
        public readonly string PrincipalId;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The id of the tenant which the identity belongs to.
        /// </summary>
        public readonly string TenantId;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetUserAssignedIdentityResult(
            string clientId,

            string location,

            string name,

            string principalId,

            ImmutableDictionary<string, string>? tags,

            string tenantId,

            string type)
        {
            ClientId = clientId;
            Location = location;
            Name = name;
            PrincipalId = principalId;
            Tags = tags;
            TenantId = tenantId;
            Type = type;
        }
    }
}
