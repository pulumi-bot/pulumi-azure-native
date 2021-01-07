// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20200801
{
    public static class GetDdosProtectionPlan
    {
        public static Task<GetDdosProtectionPlanResult> InvokeAsync(GetDdosProtectionPlanArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDdosProtectionPlanResult>("azure-nextgen:network/v20200801:getDdosProtectionPlan", args ?? new GetDdosProtectionPlanArgs(), options.WithVersion());
    }


    public sealed class GetDdosProtectionPlanArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the DDoS protection plan.
        /// </summary>
        [Input("ddosProtectionPlanName", required: true)]
        public string DdosProtectionPlanName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetDdosProtectionPlanArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDdosProtectionPlanResult
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Resource location.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The provisioning state of the DDoS protection plan resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The resource GUID property of the DDoS protection plan resource. It uniquely identifies the resource, even if the user changes its name or migrate the resource across subscriptions or resource groups.
        /// </summary>
        public readonly string ResourceGuid;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The list of virtual networks associated with the DDoS protection plan resource. This list is read-only.
        /// </summary>
        public readonly ImmutableArray<Outputs.SubResourceResponse> VirtualNetworks;

        [OutputConstructor]
        private GetDdosProtectionPlanResult(
            string etag,

            string id,

            string? location,

            string name,

            string provisioningState,

            string resourceGuid,

            ImmutableDictionary<string, string>? tags,

            string type,

            ImmutableArray<Outputs.SubResourceResponse> virtualNetworks)
        {
            Etag = etag;
            Id = id;
            Location = location;
            Name = name;
            ProvisioningState = provisioningState;
            ResourceGuid = resourceGuid;
            Tags = tags;
            Type = type;
            VirtualNetworks = virtualNetworks;
        }
    }
}
