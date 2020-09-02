// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.AzureStackHCI.V20200301Preview
{
    public static class GetCluster
    {
        public static Task<GetClusterResult> InvokeAsync(GetClusterArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetClusterResult>("azurerm:azurestackhci/v20200301preview:getCluster", args ?? new GetClusterArgs(), options.WithVersion());
    }


    public sealed class GetClusterArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the cluster.
        /// </summary>
        [Input("clusterName", required: true)]
        public string ClusterName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
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
        /// App id of cluster AAD identity.
        /// </summary>
        public readonly string AadClientId;
        /// <summary>
        /// Tenant id of cluster AAD identity.
        /// </summary>
        public readonly string AadTenantId;
        /// <summary>
        /// Type of billing applied to the resource.
        /// </summary>
        public readonly string BillingModel;
        /// <summary>
        /// Unique, immutable resource id.
        /// </summary>
        public readonly string CloudId;
        /// <summary>
        /// Most recent billing meter timestamp.
        /// </summary>
        public readonly string LastBillingTimestamp;
        /// <summary>
        /// Most recent cluster sync timestamp.
        /// </summary>
        public readonly string LastSyncTimestamp;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Provisioning state.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// First cluster sync timestamp.
        /// </summary>
        public readonly string RegistrationTimestamp;
        /// <summary>
        /// Properties reported by cluster agent.
        /// </summary>
        public readonly Outputs.ClusterReportedPropertiesResponseResult? ReportedProperties;
        /// <summary>
        /// Status of the cluster agent.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Number of days remaining in the trial period.
        /// </summary>
        public readonly double TrialDaysRemaining;
        /// <summary>
        /// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetClusterResult(
            string aadClientId,

            string aadTenantId,

            string billingModel,

            string cloudId,

            string lastBillingTimestamp,

            string lastSyncTimestamp,

            string location,

            string name,

            string provisioningState,

            string registrationTimestamp,

            Outputs.ClusterReportedPropertiesResponseResult? reportedProperties,

            string status,

            ImmutableDictionary<string, string>? tags,

            double trialDaysRemaining,

            string type)
        {
            AadClientId = aadClientId;
            AadTenantId = aadTenantId;
            BillingModel = billingModel;
            CloudId = cloudId;
            LastBillingTimestamp = lastBillingTimestamp;
            LastSyncTimestamp = lastSyncTimestamp;
            Location = location;
            Name = name;
            ProvisioningState = provisioningState;
            RegistrationTimestamp = registrationTimestamp;
            ReportedProperties = reportedProperties;
            Status = status;
            Tags = tags;
            TrialDaysRemaining = trialDaysRemaining;
            Type = type;
        }
    }
}
