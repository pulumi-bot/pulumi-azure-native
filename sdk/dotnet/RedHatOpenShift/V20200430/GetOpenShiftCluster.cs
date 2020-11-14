// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.RedHatOpenShift.V20200430
{
    public static class GetOpenShiftCluster
    {
        public static Task<GetOpenShiftClusterResult> InvokeAsync(GetOpenShiftClusterArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetOpenShiftClusterResult>("azure-nextgen:redhatopenshift/v20200430:getOpenShiftCluster", args ?? new GetOpenShiftClusterArgs(), options.WithVersion());
    }


    public sealed class GetOpenShiftClusterArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the OpenShift cluster resource.
        /// </summary>
        [Input("resourceName", required: true)]
        public string ResourceName { get; set; } = null!;

        public GetOpenShiftClusterArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetOpenShiftClusterResult
    {
        /// <summary>
        /// The cluster API server profile.
        /// </summary>
        public readonly Outputs.APIServerProfileResponse? ApiserverProfile;
        /// <summary>
        /// The cluster profile.
        /// </summary>
        public readonly Outputs.ClusterProfileResponse? ClusterProfile;
        /// <summary>
        /// The console profile.
        /// </summary>
        public readonly Outputs.ConsoleProfileResponse? ConsoleProfile;
        /// <summary>
        /// The cluster ingress profiles.
        /// </summary>
        public readonly ImmutableArray<Outputs.IngressProfileResponse> IngressProfiles;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The cluster master profile.
        /// </summary>
        public readonly Outputs.MasterProfileResponse? MasterProfile;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The cluster network profile.
        /// </summary>
        public readonly Outputs.NetworkProfileResponse? NetworkProfile;
        /// <summary>
        /// The cluster provisioning state (immutable).
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// The cluster service principal profile.
        /// </summary>
        public readonly Outputs.ServicePrincipalProfileResponse? ServicePrincipalProfile;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The cluster worker profiles.
        /// </summary>
        public readonly ImmutableArray<Outputs.WorkerProfileResponse> WorkerProfiles;

        [OutputConstructor]
        private GetOpenShiftClusterResult(
            Outputs.APIServerProfileResponse? apiserverProfile,

            Outputs.ClusterProfileResponse? clusterProfile,

            Outputs.ConsoleProfileResponse? consoleProfile,

            ImmutableArray<Outputs.IngressProfileResponse> ingressProfiles,

            string location,

            Outputs.MasterProfileResponse? masterProfile,

            string name,

            Outputs.NetworkProfileResponse? networkProfile,

            string? provisioningState,

            Outputs.ServicePrincipalProfileResponse? servicePrincipalProfile,

            ImmutableDictionary<string, string>? tags,

            string type,

            ImmutableArray<Outputs.WorkerProfileResponse> workerProfiles)
        {
            ApiserverProfile = apiserverProfile;
            ClusterProfile = clusterProfile;
            ConsoleProfile = consoleProfile;
            IngressProfiles = ingressProfiles;
            Location = location;
            MasterProfile = masterProfile;
            Name = name;
            NetworkProfile = networkProfile;
            ProvisioningState = provisioningState;
            ServicePrincipalProfile = servicePrincipalProfile;
            Tags = tags;
            Type = type;
            WorkerProfiles = workerProfiles;
        }
    }
}
