// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.RedHatOpenShift.Outputs
{

    [OutputType]
    public sealed class OpenShiftClusterPropertiesResponseResult
    {
        /// <summary>
        /// The cluster API server profile.
        /// </summary>
        public readonly Outputs.APIServerProfileResponseResult? ApiserverProfile;
        /// <summary>
        /// The cluster profile.
        /// </summary>
        public readonly Outputs.ClusterProfileResponseResult? ClusterProfile;
        /// <summary>
        /// The console profile.
        /// </summary>
        public readonly Outputs.ConsoleProfileResponseResult? ConsoleProfile;
        /// <summary>
        /// The cluster ingress profiles.
        /// </summary>
        public readonly ImmutableArray<Outputs.IngressProfileResponseResult> IngressProfiles;
        /// <summary>
        /// The cluster master profile.
        /// </summary>
        public readonly Outputs.MasterProfileResponseResult? MasterProfile;
        /// <summary>
        /// The cluster network profile.
        /// </summary>
        public readonly Outputs.NetworkProfileResponseResult? NetworkProfile;
        /// <summary>
        /// The cluster provisioning state (immutable).
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// The cluster service principal profile.
        /// </summary>
        public readonly Outputs.ServicePrincipalProfileResponseResult? ServicePrincipalProfile;
        /// <summary>
        /// The cluster worker profiles.
        /// </summary>
        public readonly ImmutableArray<Outputs.WorkerProfileResponseResult> WorkerProfiles;

        [OutputConstructor]
        private OpenShiftClusterPropertiesResponseResult(
            Outputs.APIServerProfileResponseResult? apiserverProfile,

            Outputs.ClusterProfileResponseResult? clusterProfile,

            Outputs.ConsoleProfileResponseResult? consoleProfile,

            ImmutableArray<Outputs.IngressProfileResponseResult> ingressProfiles,

            Outputs.MasterProfileResponseResult? masterProfile,

            Outputs.NetworkProfileResponseResult? networkProfile,

            string? provisioningState,

            Outputs.ServicePrincipalProfileResponseResult? servicePrincipalProfile,

            ImmutableArray<Outputs.WorkerProfileResponseResult> workerProfiles)
        {
            ApiserverProfile = apiserverProfile;
            ClusterProfile = clusterProfile;
            ConsoleProfile = consoleProfile;
            IngressProfiles = ingressProfiles;
            MasterProfile = masterProfile;
            NetworkProfile = networkProfile;
            ProvisioningState = provisioningState;
            ServicePrincipalProfile = servicePrincipalProfile;
            WorkerProfiles = workerProfiles;
        }
    }
}