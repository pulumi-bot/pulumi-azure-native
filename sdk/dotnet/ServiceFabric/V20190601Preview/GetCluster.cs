// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceFabric.V20190601Preview
{
    public static class GetCluster
    {
        public static Task<GetClusterResult> InvokeAsync(GetClusterArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetClusterResult>("azurerm:servicefabric/v20190601preview:getCluster", args ?? new GetClusterArgs(), options.WithVersion());
    }


    public sealed class GetClusterArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the cluster resource.
        /// </summary>
        [Input("clusterName", required: true)]
        public string ClusterName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
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
        /// The list of add-on features to enable in the cluster.
        /// </summary>
        public readonly ImmutableArray<string> AddOnFeatures;
        /// <summary>
        /// The Service Fabric runtime versions available for this cluster.
        /// </summary>
        public readonly ImmutableArray<Outputs.ClusterVersionDetailsResponseResult> AvailableClusterVersions;
        /// <summary>
        /// The AAD authentication settings of the cluster.
        /// </summary>
        public readonly Outputs.AzureActiveDirectoryResponseResult? AzureActiveDirectory;
        /// <summary>
        /// The certificate to use for securing the cluster. The certificate provided will be used for node to node security within the cluster, SSL certificate for cluster management endpoint and default admin client.
        /// </summary>
        public readonly Outputs.CertificateDescriptionResponseResult? Certificate;
        /// <summary>
        /// Describes a list of server certificates referenced by common name that are used to secure the cluster.
        /// </summary>
        public readonly Outputs.ServerCertificateCommonNamesResponseResult? CertificateCommonNames;
        /// <summary>
        /// The list of client certificates referenced by common name that are allowed to manage the cluster.
        /// </summary>
        public readonly ImmutableArray<Outputs.ClientCertificateCommonNameResponseResult> ClientCertificateCommonNames;
        /// <summary>
        /// The list of client certificates referenced by thumbprint that are allowed to manage the cluster.
        /// </summary>
        public readonly ImmutableArray<Outputs.ClientCertificateThumbprintResponseResult> ClientCertificateThumbprints;
        /// <summary>
        /// The Service Fabric runtime version of the cluster. This property can only by set the user when **upgradeMode** is set to 'Manual'. To get list of available Service Fabric versions for new clusters use [ClusterVersion API](./ClusterVersion.md). To get the list of available version for existing clusters use **availableClusterVersions**.
        /// </summary>
        public readonly string? ClusterCodeVersion;
        /// <summary>
        /// The Azure Resource Provider endpoint. A system service in the cluster connects to this  endpoint.
        /// </summary>
        public readonly string ClusterEndpoint;
        /// <summary>
        /// A service generated unique identifier for the cluster resource.
        /// </summary>
        public readonly string ClusterId;
        /// <summary>
        /// The current state of the cluster.
        /// 
        ///   - WaitingForNodes - Indicates that the cluster resource is created and the resource provider is waiting for Service Fabric VM extension to boot up and report to it.
        ///   - Deploying - Indicates that the Service Fabric runtime is being installed on the VMs. Cluster resource will be in this state until the cluster boots up and system services are up.
        ///   - BaselineUpgrade - Indicates that the cluster is upgrading to establishes the cluster version. This upgrade is automatically initiated when the cluster boots up for the first time.
        ///   - UpdatingUserConfiguration - Indicates that the cluster is being upgraded with the user provided configuration.
        ///   - UpdatingUserCertificate - Indicates that the cluster is being upgraded with the user provided certificate.
        ///   - UpdatingInfrastructure - Indicates that the cluster is being upgraded with the latest Service Fabric runtime version. This happens only when the **upgradeMode** is set to 'Automatic'.
        ///   - EnforcingClusterVersion - Indicates that cluster is on a different version than expected and the cluster is being upgraded to the expected version.
        ///   - UpgradeServiceUnreachable - Indicates that the system service in the cluster is no longer polling the Resource Provider. Clusters in this state cannot be managed by the Resource Provider.
        ///   - AutoScale - Indicates that the ReliabilityLevel of the cluster is being adjusted.
        ///   - Ready - Indicates that the cluster is in a stable state.
        /// </summary>
        public readonly string ClusterState;
        /// <summary>
        /// The storage account information for storing Service Fabric diagnostic logs.
        /// </summary>
        public readonly Outputs.DiagnosticsStorageAccountConfigResponseResult? DiagnosticsStorageAccountConfig;
        /// <summary>
        /// Azure resource etag.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Indicates if the event store service is enabled.
        /// </summary>
        public readonly bool? EventStoreServiceEnabled;
        /// <summary>
        /// The list of custom fabric settings to configure the cluster.
        /// </summary>
        public readonly ImmutableArray<Outputs.SettingsSectionDescriptionResponseResult> FabricSettings;
        /// <summary>
        /// Azure resource location.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The http management endpoint of the cluster.
        /// </summary>
        public readonly string ManagementEndpoint;
        /// <summary>
        /// Azure resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The list of node types in the cluster.
        /// </summary>
        public readonly ImmutableArray<Outputs.NodeTypeDescriptionResponseResult> NodeTypes;
        /// <summary>
        /// The provisioning state of the cluster resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The reliability level sets the replica set size of system services. Learn about [ReliabilityLevel](https://docs.microsoft.com/azure/service-fabric/service-fabric-cluster-capacity).
        /// 
        ///   - None - Run the System services with a target replica set count of 1. This should only be used for test clusters.
        ///   - Bronze - Run the System services with a target replica set count of 3. This should only be used for test clusters.
        ///   - Silver - Run the System services with a target replica set count of 5.
        ///   - Gold - Run the System services with a target replica set count of 7.
        ///   - Platinum - Run the System services with a target replica set count of 9.
        /// </summary>
        public readonly string? ReliabilityLevel;
        /// <summary>
        /// The server certificate used by reverse proxy.
        /// </summary>
        public readonly Outputs.CertificateDescriptionResponseResult? ReverseProxyCertificate;
        /// <summary>
        /// Describes a list of server certificates referenced by common name that are used to secure the cluster.
        /// </summary>
        public readonly Outputs.ServerCertificateCommonNamesResponseResult? ReverseProxyCertificateCommonNames;
        /// <summary>
        /// Azure resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Azure resource type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The policy to use when upgrading the cluster.
        /// </summary>
        public readonly Outputs.ClusterUpgradePolicyResponseResult? UpgradeDescription;
        /// <summary>
        /// The upgrade mode of the cluster when new Service Fabric runtime version is available.
        /// 
        ///   - Automatic - The cluster will be automatically upgraded to the latest Service Fabric runtime version as soon as it is available.
        ///   - Manual - The cluster will not be automatically upgraded to the latest Service Fabric runtime version. The cluster is upgraded by setting the **clusterCodeVersion** property in the cluster resource.
        /// </summary>
        public readonly string? UpgradeMode;
        /// <summary>
        /// The VM image VMSS has been configured with. Generic names such as Windows or Linux can be used.
        /// </summary>
        public readonly string? VmImage;

        [OutputConstructor]
        private GetClusterResult(
            ImmutableArray<string> addOnFeatures,

            ImmutableArray<Outputs.ClusterVersionDetailsResponseResult> availableClusterVersions,

            Outputs.AzureActiveDirectoryResponseResult? azureActiveDirectory,

            Outputs.CertificateDescriptionResponseResult? certificate,

            Outputs.ServerCertificateCommonNamesResponseResult? certificateCommonNames,

            ImmutableArray<Outputs.ClientCertificateCommonNameResponseResult> clientCertificateCommonNames,

            ImmutableArray<Outputs.ClientCertificateThumbprintResponseResult> clientCertificateThumbprints,

            string? clusterCodeVersion,

            string clusterEndpoint,

            string clusterId,

            string clusterState,

            Outputs.DiagnosticsStorageAccountConfigResponseResult? diagnosticsStorageAccountConfig,

            string etag,

            bool? eventStoreServiceEnabled,

            ImmutableArray<Outputs.SettingsSectionDescriptionResponseResult> fabricSettings,

            string location,

            string managementEndpoint,

            string name,

            ImmutableArray<Outputs.NodeTypeDescriptionResponseResult> nodeTypes,

            string provisioningState,

            string? reliabilityLevel,

            Outputs.CertificateDescriptionResponseResult? reverseProxyCertificate,

            Outputs.ServerCertificateCommonNamesResponseResult? reverseProxyCertificateCommonNames,

            ImmutableDictionary<string, string>? tags,

            string type,

            Outputs.ClusterUpgradePolicyResponseResult? upgradeDescription,

            string? upgradeMode,

            string? vmImage)
        {
            AddOnFeatures = addOnFeatures;
            AvailableClusterVersions = availableClusterVersions;
            AzureActiveDirectory = azureActiveDirectory;
            Certificate = certificate;
            CertificateCommonNames = certificateCommonNames;
            ClientCertificateCommonNames = clientCertificateCommonNames;
            ClientCertificateThumbprints = clientCertificateThumbprints;
            ClusterCodeVersion = clusterCodeVersion;
            ClusterEndpoint = clusterEndpoint;
            ClusterId = clusterId;
            ClusterState = clusterState;
            DiagnosticsStorageAccountConfig = diagnosticsStorageAccountConfig;
            Etag = etag;
            EventStoreServiceEnabled = eventStoreServiceEnabled;
            FabricSettings = fabricSettings;
            Location = location;
            ManagementEndpoint = managementEndpoint;
            Name = name;
            NodeTypes = nodeTypes;
            ProvisioningState = provisioningState;
            ReliabilityLevel = reliabilityLevel;
            ReverseProxyCertificate = reverseProxyCertificate;
            ReverseProxyCertificateCommonNames = reverseProxyCertificateCommonNames;
            Tags = tags;
            Type = type;
            UpgradeDescription = upgradeDescription;
            UpgradeMode = upgradeMode;
            VmImage = vmImage;
        }
    }
}
