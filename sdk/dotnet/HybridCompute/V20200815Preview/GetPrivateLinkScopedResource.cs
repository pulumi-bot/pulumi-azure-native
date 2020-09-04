// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.HybridCompute.V20200815Preview
{
    public static class GetPrivateLinkScopedResource
    {
        public static Task<GetPrivateLinkScopedResourceResult> InvokeAsync(GetPrivateLinkScopedResourceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPrivateLinkScopedResourceResult>("azurerm:hybridcompute/v20200815preview:getPrivateLinkScopedResource", args ?? new GetPrivateLinkScopedResourceArgs(), options.WithVersion());
    }


    public sealed class GetPrivateLinkScopedResourceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the scoped resource object.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the Azure Arc PrivateLinkScope resource.
        /// </summary>
        [Input("scopeName", required: true)]
        public string ScopeName { get; set; } = null!;

        public GetPrivateLinkScopedResourceArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPrivateLinkScopedResourceResult
    {
        /// <summary>
        /// The resource id of the scoped Azure monitor resource.
        /// </summary>
        public readonly string? LinkedResourceId;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// State of the private endpoint connection.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetPrivateLinkScopedResourceResult(
            string? linkedResourceId,

            string name,

            string provisioningState,

            string type)
        {
            LinkedResourceId = linkedResourceId;
            Name = name;
            ProvisioningState = provisioningState;
            Type = type;
        }
    }
}
