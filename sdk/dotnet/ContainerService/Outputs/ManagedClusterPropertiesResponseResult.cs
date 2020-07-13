// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ContainerService.Outputs
{

    [OutputType]
    public sealed class ManagedClusterPropertiesResponseResult
    {
        /// <summary>
        /// Profile of Azure Active Directory configuration.
        /// </summary>
        public readonly Outputs.ManagedClusterAADProfileResponseResult? AadProfile;
        /// <summary>
        /// Profile of managed cluster add-on.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? AddonProfiles;
        /// <summary>
        /// Properties of the agent pool.
        /// </summary>
        public readonly ImmutableArray<Outputs.ManagedClusterAgentPoolProfileResponseResult> AgentPoolProfiles;
        /// <summary>
        /// Access profile for managed cluster API server.
        /// </summary>
        public readonly Outputs.ManagedClusterAPIServerAccessProfileResponseResult? ApiServerAccessProfile;
        /// <summary>
        /// Parameters to be applied to the cluster-autoscaler when enabled
        /// </summary>
        public readonly Outputs.ManagedClusterPropertiesResponsePropertiesResult? AutoScalerProfile;
        /// <summary>
        /// ResourceId of the disk encryption set to use for enabling encryption at rest.
        /// </summary>
        public readonly string? DiskEncryptionSetID;
        /// <summary>
        /// DNS prefix specified when creating the managed cluster.
        /// </summary>
        public readonly string? DnsPrefix;
        /// <summary>
        /// (PREVIEW) Whether to enable Kubernetes Pod security policy.
        /// </summary>
        public readonly bool? EnablePodSecurityPolicy;
        /// <summary>
        /// Whether to enable Kubernetes Role-Based Access Control.
        /// </summary>
        public readonly bool? EnableRBAC;
        /// <summary>
        /// FQDN for the master pool.
        /// </summary>
        public readonly string Fqdn;
        /// <summary>
        /// Identities associated with the cluster.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? IdentityProfile;
        /// <summary>
        /// Version of Kubernetes specified when creating the managed cluster.
        /// </summary>
        public readonly string? KubernetesVersion;
        /// <summary>
        /// Profile for Linux VMs in the container service cluster.
        /// </summary>
        public readonly Outputs.ContainerServiceLinuxProfileResponseResult? LinuxProfile;
        /// <summary>
        /// The max number of agent pools for the managed cluster.
        /// </summary>
        public readonly int MaxAgentPools;
        /// <summary>
        /// Profile of network configuration.
        /// </summary>
        public readonly Outputs.ContainerServiceNetworkProfileResponseResult? NetworkProfile;
        /// <summary>
        /// Name of the resource group containing agent pool nodes.
        /// </summary>
        public readonly string? NodeResourceGroup;
        /// <summary>
        /// FQDN of private cluster.
        /// </summary>
        public readonly string PrivateFQDN;
        /// <summary>
        /// The current deployment or provisioning state, which only appears in the response.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Information about a service principal identity for the cluster to use for manipulating Azure APIs.
        /// </summary>
        public readonly Outputs.ManagedClusterServicePrincipalProfileResponseResult? ServicePrincipalProfile;
        /// <summary>
        /// Profile for Windows VMs in the container service cluster.
        /// </summary>
        public readonly Outputs.ManagedClusterWindowsProfileResponseResult? WindowsProfile;

        [OutputConstructor]
        private ManagedClusterPropertiesResponseResult(
            Outputs.ManagedClusterAADProfileResponseResult? aadProfile,

            ImmutableDictionary<string, string>? addonProfiles,

            ImmutableArray<Outputs.ManagedClusterAgentPoolProfileResponseResult> agentPoolProfiles,

            Outputs.ManagedClusterAPIServerAccessProfileResponseResult? apiServerAccessProfile,

            Outputs.ManagedClusterPropertiesResponsePropertiesResult? autoScalerProfile,

            string? diskEncryptionSetID,

            string? dnsPrefix,

            bool? enablePodSecurityPolicy,

            bool? enableRBAC,

            string fqdn,

            ImmutableDictionary<string, string>? identityProfile,

            string? kubernetesVersion,

            Outputs.ContainerServiceLinuxProfileResponseResult? linuxProfile,

            int maxAgentPools,

            Outputs.ContainerServiceNetworkProfileResponseResult? networkProfile,

            string? nodeResourceGroup,

            string privateFQDN,

            string provisioningState,

            Outputs.ManagedClusterServicePrincipalProfileResponseResult? servicePrincipalProfile,

            Outputs.ManagedClusterWindowsProfileResponseResult? windowsProfile)
        {
            AadProfile = aadProfile;
            AddonProfiles = addonProfiles;
            AgentPoolProfiles = agentPoolProfiles;
            ApiServerAccessProfile = apiServerAccessProfile;
            AutoScalerProfile = autoScalerProfile;
            DiskEncryptionSetID = diskEncryptionSetID;
            DnsPrefix = dnsPrefix;
            EnablePodSecurityPolicy = enablePodSecurityPolicy;
            EnableRBAC = enableRBAC;
            Fqdn = fqdn;
            IdentityProfile = identityProfile;
            KubernetesVersion = kubernetesVersion;
            LinuxProfile = linuxProfile;
            MaxAgentPools = maxAgentPools;
            NetworkProfile = networkProfile;
            NodeResourceGroup = nodeResourceGroup;
            PrivateFQDN = privateFQDN;
            ProvisioningState = provisioningState;
            ServicePrincipalProfile = servicePrincipalProfile;
            WindowsProfile = windowsProfile;
        }
    }
}