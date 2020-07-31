// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceFabric.V20160901.Outputs
{

    [OutputType]
    public sealed class ClusterPropertiesResponseResult
    {
        /// <summary>
        /// The available cluster code version which the cluster can upgrade to, note that you must choose upgradeMode to manual to upgrade to
        /// </summary>
        public readonly ImmutableArray<Outputs.ClusterVersionDetailsResponseResult> AvailableClusterVersions;
        /// <summary>
        /// The settings to enable AAD authentication on the cluster
        /// </summary>
        public readonly Outputs.AzureActiveDirectoryResponseResult? AzureActiveDirectory;
        /// <summary>
        /// This primary certificate will be used as cluster node to node security, SSL certificate for cluster management endpoint and default admin client
        /// </summary>
        public readonly Outputs.CertificateDescriptionResponseResult? Certificate;
        /// <summary>
        ///  List of client certificates to whitelist based on common names
        /// </summary>
        public readonly ImmutableArray<Outputs.ClientCertificateCommonNameResponseResult> ClientCertificateCommonNames;
        /// <summary>
        /// The client thumbprint details ,it is used for client access for cluster operation
        /// </summary>
        public readonly ImmutableArray<Outputs.ClientCertificateThumbprintResponseResult> ClientCertificateThumbprints;
        /// <summary>
        /// The ServiceFabric code version running in your cluster
        /// </summary>
        public readonly string? ClusterCodeVersion;
        /// <summary>
        /// The endpoint for the cluster connecting to servicefabric resource provider
        /// </summary>
        public readonly string ClusterEndpoint;
        /// <summary>
        /// The unique identifier for the cluster resource
        /// </summary>
        public readonly string ClusterId;
        /// <summary>
        /// The state for the cluster
        /// </summary>
        public readonly string ClusterState;
        /// <summary>
        /// The storage diagnostics account configuration details
        /// </summary>
        public readonly Outputs.DiagnosticsStorageAccountConfigResponseResult? DiagnosticsStorageAccountConfig;
        /// <summary>
        /// List of custom fabric settings to configure the cluster.
        /// </summary>
        public readonly ImmutableArray<Outputs.SettingsSectionDescriptionResponseResult> FabricSettings;
        /// <summary>
        /// The http management endpoint of the cluster
        /// </summary>
        public readonly string ManagementEndpoint;
        /// <summary>
        /// The list of node types that make up the cluster
        /// </summary>
        public readonly ImmutableArray<Outputs.NodeTypeDescriptionResponseResult> NodeTypes;
        /// <summary>
        /// The provisioning state of the cluster resource
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Cluster reliability level indicates replica set size of system service
        /// </summary>
        public readonly string? ReliabilityLevel;
        /// <summary>
        /// The server certificate used by reverse proxy
        /// </summary>
        public readonly Outputs.CertificateDescriptionResponseResult? ReverseProxyCertificate;
        /// <summary>
        /// The policy to use when upgrading the cluster.
        /// </summary>
        public readonly Outputs.ClusterUpgradePolicyResponseResult? UpgradeDescription;
        /// <summary>
        /// Cluster upgrade mode indicates if fabric upgrade is initiated automatically by the system or not
        /// </summary>
        public readonly string? UpgradeMode;
        /// <summary>
        /// The name of VM image VMSS has been configured with. Generic names such as Windows or Linux can be used.
        /// </summary>
        public readonly string? VmImage;

        [OutputConstructor]
        private ClusterPropertiesResponseResult(
            ImmutableArray<Outputs.ClusterVersionDetailsResponseResult> availableClusterVersions,

            Outputs.AzureActiveDirectoryResponseResult? azureActiveDirectory,

            Outputs.CertificateDescriptionResponseResult? certificate,

            ImmutableArray<Outputs.ClientCertificateCommonNameResponseResult> clientCertificateCommonNames,

            ImmutableArray<Outputs.ClientCertificateThumbprintResponseResult> clientCertificateThumbprints,

            string? clusterCodeVersion,

            string clusterEndpoint,

            string clusterId,

            string clusterState,

            Outputs.DiagnosticsStorageAccountConfigResponseResult? diagnosticsStorageAccountConfig,

            ImmutableArray<Outputs.SettingsSectionDescriptionResponseResult> fabricSettings,

            string managementEndpoint,

            ImmutableArray<Outputs.NodeTypeDescriptionResponseResult> nodeTypes,

            string provisioningState,

            string? reliabilityLevel,

            Outputs.CertificateDescriptionResponseResult? reverseProxyCertificate,

            Outputs.ClusterUpgradePolicyResponseResult? upgradeDescription,

            string? upgradeMode,

            string? vmImage)
        {
            AvailableClusterVersions = availableClusterVersions;
            AzureActiveDirectory = azureActiveDirectory;
            Certificate = certificate;
            ClientCertificateCommonNames = clientCertificateCommonNames;
            ClientCertificateThumbprints = clientCertificateThumbprints;
            ClusterCodeVersion = clusterCodeVersion;
            ClusterEndpoint = clusterEndpoint;
            ClusterId = clusterId;
            ClusterState = clusterState;
            DiagnosticsStorageAccountConfig = diagnosticsStorageAccountConfig;
            FabricSettings = fabricSettings;
            ManagementEndpoint = managementEndpoint;
            NodeTypes = nodeTypes;
            ProvisioningState = provisioningState;
            ReliabilityLevel = reliabilityLevel;
            ReverseProxyCertificate = reverseProxyCertificate;
            UpgradeDescription = upgradeDescription;
            UpgradeMode = upgradeMode;
            VmImage = vmImage;
        }
    }
}