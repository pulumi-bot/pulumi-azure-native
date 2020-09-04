// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Kusto.V20170907privatePreview
{
    public static class GetCluster
    {
        public static Task<GetClusterResult> InvokeAsync(GetClusterArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetClusterResult>("azurerm:kusto/v20170907privatepreview:getCluster", args ?? new GetClusterArgs(), options.WithVersion());
    }


    public sealed class GetClusterArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Kusto cluster.
        /// </summary>
        [Input("clusterName", required: true)]
        public string ClusterName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group containing the Kusto cluster.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetClusterArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetClusterResult
    {
        /// <summary>
        /// The cluster data ingestion URI.
        /// </summary>
        public readonly string DataIngestionUri;
        /// <summary>
        /// An ETag of the resource created.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The provisioned state of the resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The SKU of the cluster.
        /// </summary>
        public readonly Outputs.AzureSkuResponseResult Sku;
        /// <summary>
        /// The state of the resource.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The cluster's external tenants.
        /// </summary>
        public readonly ImmutableArray<Outputs.TrustedExternalTenantResponseResult> TrustedExternalTenants;
        /// <summary>
        /// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The cluster URI.
        /// </summary>
        public readonly string Uri;

        [OutputConstructor]
        private GetClusterResult(
            string dataIngestionUri,

            string etag,

            string location,

            string name,

            string provisioningState,

            Outputs.AzureSkuResponseResult sku,

            string state,

            ImmutableDictionary<string, string>? tags,

            ImmutableArray<Outputs.TrustedExternalTenantResponseResult> trustedExternalTenants,

            string type,

            string uri)
        {
            DataIngestionUri = dataIngestionUri;
            Etag = etag;
            Location = location;
            Name = name;
            ProvisioningState = provisioningState;
            Sku = sku;
            State = state;
            Tags = tags;
            TrustedExternalTenants = trustedExternalTenants;
            Type = type;
            Uri = uri;
        }
    }
}
