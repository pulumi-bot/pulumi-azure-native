// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ManagedIdentity
{
    public static class GetUserAssignedIdentity
    {
        public static Task<GetUserAssignedIdentityResult> InvokeAsync(GetUserAssignedIdentityArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetUserAssignedIdentityResult>("azurerm:managedidentity:getUserAssignedIdentity", args ?? new GetUserAssignedIdentityArgs(), options.WithVersion());
    }


    public sealed class GetUserAssignedIdentityArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the identity resource.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group to which the identity belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetUserAssignedIdentityArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetUserAssignedIdentityResult
    {
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The properties associated with the identity.
        /// </summary>
        public readonly Outputs.UserAssignedIdentityPropertiesResponseResult Properties;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetUserAssignedIdentityResult(
            string location,

            string name,

            Outputs.UserAssignedIdentityPropertiesResponseResult properties,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Location = location;
            Name = name;
            Properties = properties;
            Tags = tags;
            Type = type;
        }
    }
}