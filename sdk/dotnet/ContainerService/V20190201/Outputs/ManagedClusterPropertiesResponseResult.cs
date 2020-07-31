// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ContainerService.V20190201.Outputs
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
        public readonly ImmutableDictionary<string, Outputs.ManagedClusterAddonProfileResponseResult>? AddonProfiles;
        /// <summary>
        /// Properties of the agent pool.
        /// </summary>
        public readonly ImmutableArray<Outputs.ManagedClusterAgentPoolProfileResponseResult> AgentPoolProfiles;
        /// <summary>
        /// (PREVIEW) Authorized IP Ranges to kubernetes API server.
        /// </summary>
        public readonly ImmutableArray<string> ApiServerAuthorizedIPRanges;
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
        /// Version of Kubernetes specified when creating the managed cluster.
        /// </summary>
        public readonly string? KubernetesVersion;
        /// <summary>
        /// Profile for Linux VMs in the container service cluster.
        /// </summary>
        public readonly Outputs.ContainerServiceLinuxProfileResponseResult? LinuxProfile;
        /// <summary>
        /// Profile of network configuration.
        /// </summary>
        public readonly Outputs.ContainerServiceNetworkProfileResponseResult? NetworkProfile;
        /// <summary>
        /// Name of the resource group containing agent pool nodes.
        /// </summary>
        public readonly string NodeResourceGroup;
        /// <summary>
        /// The current deployment or provisioning state, which only appears in the response.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Information about a service principal identity for the cluster to use for manipulating Azure APIs.
        /// </summary>
        public readonly Outputs.ManagedClusterServicePrincipalProfileResponseResult? ServicePrincipalProfile;

        [OutputConstructor]
        private ManagedClusterPropertiesResponseResult(
            Outputs.ManagedClusterAADProfileResponseResult? aadProfile,

            ImmutableDictionary<string, Outputs.ManagedClusterAddonProfileResponseResult>? addonProfiles,

            ImmutableArray<Outputs.ManagedClusterAgentPoolProfileResponseResult> agentPoolProfiles,

            ImmutableArray<string> apiServerAuthorizedIPRanges,

            string? dnsPrefix,

            bool? enablePodSecurityPolicy,

            bool? enableRBAC,

            string fqdn,

            string? kubernetesVersion,

            Outputs.ContainerServiceLinuxProfileResponseResult? linuxProfile,

            Outputs.ContainerServiceNetworkProfileResponseResult? networkProfile,

            string nodeResourceGroup,

            string provisioningState,

            Outputs.ManagedClusterServicePrincipalProfileResponseResult? servicePrincipalProfile)
        {
            AadProfile = aadProfile;
            AddonProfiles = addonProfiles;
            AgentPoolProfiles = agentPoolProfiles;
            ApiServerAuthorizedIPRanges = apiServerAuthorizedIPRanges;
            DnsPrefix = dnsPrefix;
            EnablePodSecurityPolicy = enablePodSecurityPolicy;
            EnableRBAC = enableRBAC;
            Fqdn = fqdn;
            KubernetesVersion = kubernetesVersion;
            LinuxProfile = linuxProfile;
            NetworkProfile = networkProfile;
            NodeResourceGroup = nodeResourceGroup;
            ProvisioningState = provisioningState;
            ServicePrincipalProfile = servicePrincipalProfile;
        }
    }
}