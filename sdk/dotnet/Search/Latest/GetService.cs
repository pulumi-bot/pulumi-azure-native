// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Search.Latest
{
    public static class GetService
    {
        public static Task<GetServiceResult> InvokeAsync(GetServiceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetServiceResult>("azurerm:search/latest:getService", args ?? new GetServiceArgs(), options.WithVersion());
    }


    public sealed class GetServiceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group within the current subscription. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the Azure Cognitive Search service associated with the specified resource group.
        /// </summary>
        [Input("searchServiceName", required: true)]
        public string SearchServiceName { get; set; } = null!;

        public GetServiceArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetServiceResult
    {
        /// <summary>
        /// Applicable only for the standard3 SKU. You can set this property to enable up to 3 high density partitions that allow up to 1000 indexes, which is much higher than the maximum indexes allowed for any other SKU. For the standard3 SKU, the value is either 'default' or 'highDensity'. For all other SKUs, this value must be 'default'.
        /// </summary>
        public readonly string? HostingMode;
        /// <summary>
        /// The identity of the resource.
        /// </summary>
        public readonly Outputs.IdentityResponseResult? Identity;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Network specific rules that determine how the Azure Cognitive Search service may be reached.
        /// </summary>
        public readonly Outputs.NetworkRuleSetResponseResult? NetworkRuleSet;
        /// <summary>
        /// The number of partitions in the search service; if specified, it can be 1, 2, 3, 4, 6, or 12. Values greater than 1 are only valid for standard SKUs. For 'standard3' services with hostingMode set to 'highDensity', the allowed values are between 1 and 3.
        /// </summary>
        public readonly int? PartitionCount;
        /// <summary>
        /// The list of private endpoint connections to the Azure Cognitive Search service.
        /// </summary>
        public readonly ImmutableArray<Outputs.PrivateEndpointConnectionResponseResult> PrivateEndpointConnections;
        /// <summary>
        /// The state of the last provisioning operation performed on the search service. Provisioning is an intermediate state that occurs while service capacity is being established. After capacity is set up, provisioningState changes to either 'succeeded' or 'failed'. Client applications can poll provisioning status (the recommended polling interval is from 30 seconds to one minute) by using the Get Search Service operation to see when an operation is completed. If you are using the free service, this value tends to come back as 'succeeded' directly in the call to Create search service. This is because the free service uses capacity that is already set up.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// This value can be set to 'enabled' to avoid breaking changes on existing customer resources and templates. If set to 'disabled', traffic over public interface is not allowed, and private endpoint connections would be the exclusive access method.
        /// </summary>
        public readonly string? PublicNetworkAccess;
        /// <summary>
        /// The number of replicas in the search service. If specified, it must be a value between 1 and 12 inclusive for standard SKUs or between 1 and 3 inclusive for basic SKU.
        /// </summary>
        public readonly int? ReplicaCount;
        /// <summary>
        /// The list of shared private link resources managed by the Azure Cognitive Search service.
        /// </summary>
        public readonly ImmutableArray<Outputs.SharedPrivateLinkResourceResponseResult> SharedPrivateLinkResources;
        /// <summary>
        /// The SKU of the Search Service, which determines price tier and capacity limits. This property is required when creating a new Search Service.
        /// </summary>
        public readonly Outputs.SkuResponseResult? Sku;
        /// <summary>
        /// The status of the search service. Possible values include: 'running': The search service is running and no provisioning operations are underway. 'provisioning': The search service is being provisioned or scaled up or down. 'deleting': The search service is being deleted. 'degraded': The search service is degraded. This can occur when the underlying search units are not healthy. The search service is most likely operational, but performance might be slow and some requests might be dropped. 'disabled': The search service is disabled. In this state, the service will reject all API requests. 'error': The search service is in an error state. If your service is in the degraded, disabled, or error states, it means the Azure Cognitive Search team is actively investigating the underlying issue. Dedicated services in these states are still chargeable based on the number of search units provisioned.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The details of the search service status.
        /// </summary>
        public readonly string StatusDetails;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetServiceResult(
            string? hostingMode,

            Outputs.IdentityResponseResult? identity,

            string location,

            string name,

            Outputs.NetworkRuleSetResponseResult? networkRuleSet,

            int? partitionCount,

            ImmutableArray<Outputs.PrivateEndpointConnectionResponseResult> privateEndpointConnections,

            string provisioningState,

            string? publicNetworkAccess,

            int? replicaCount,

            ImmutableArray<Outputs.SharedPrivateLinkResourceResponseResult> sharedPrivateLinkResources,

            Outputs.SkuResponseResult? sku,

            string status,

            string statusDetails,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            HostingMode = hostingMode;
            Identity = identity;
            Location = location;
            Name = name;
            NetworkRuleSet = networkRuleSet;
            PartitionCount = partitionCount;
            PrivateEndpointConnections = privateEndpointConnections;
            ProvisioningState = provisioningState;
            PublicNetworkAccess = publicNetworkAccess;
            ReplicaCount = replicaCount;
            SharedPrivateLinkResources = sharedPrivateLinkResources;
            Sku = sku;
            Status = status;
            StatusDetails = statusDetails;
            Tags = tags;
            Type = type;
        }
    }
}
