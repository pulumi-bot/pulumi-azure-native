// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceFabric.Latest
{
    /// <summary>
    /// The cluster resource
    /// 
    /// ## Example Usage
    /// ### Put a cluster with maximum parameters
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var cluster = new AzureRM.ServiceFabric.Latest.Cluster("cluster", new AzureRM.ServiceFabric.Latest.ClusterArgs
    ///         {
    ///             AddOnFeatures = 
    ///             {
    ///                 "RepairManager",
    ///                 "DnsService",
    ///                 "BackupRestoreService",
    ///                 "ResourceMonitorService",
    ///             },
    ///             ApplicationTypeVersionsCleanupPolicy = new AzureRM.ServiceFabric.Latest.Inputs.ApplicationTypeVersionsCleanupPolicyArgs
    ///             {
    ///                 MaxUnusedVersionsToKeep = 2,
    ///             },
    ///             AzureActiveDirectory = new AzureRM.ServiceFabric.Latest.Inputs.AzureActiveDirectoryArgs
    ///             {
    ///                 ClientApplication = "d151ad89-4bce-4ae8-b3d1-1dc79679fa75",
    ///                 ClusterApplication = "5886372e-7bf4-4878-a497-8098aba608ae",
    ///                 TenantId = "6abcc6a0-8666-43f1-87b8-172cf86a9f9c",
    ///             },
    ///             CertificateCommonNames = new AzureRM.ServiceFabric.Latest.Inputs.ServerCertificateCommonNamesArgs
    ///             {
    ///                 CommonNames = 
    ///                 {
    ///                     new AzureRM.ServiceFabric.Latest.Inputs.ServerCertificateCommonNameArgs
    ///                     {
    ///                         CertificateCommonName = "abc.com",
    ///                         CertificateIssuerThumbprint = "12599211F8F14C90AFA9532AD79A6F2CA1C00622",
    ///                     },
    ///                 },
    ///                 X509StoreName = "My",
    ///             },
    ///             ClientCertificateCommonNames = 
    ///             {
    ///                 new AzureRM.ServiceFabric.Latest.Inputs.ClientCertificateCommonNameArgs
    ///                 {
    ///                     CertificateCommonName = "abc.com",
    ///                     CertificateIssuerThumbprint = "5F3660C715EBBDA31DB1FFDCF508302348DE8E7A",
    ///                     IsAdmin = true,
    ///                 },
    ///             },
    ///             ClientCertificateThumbprints = 
    ///             {
    ///                 new AzureRM.ServiceFabric.Latest.Inputs.ClientCertificateThumbprintArgs
    ///                 {
    ///                     CertificateThumbprint = "5F3660C715EBBDA31DB1FFDCF508302348DE8E7A",
    ///                     IsAdmin = true,
    ///                 },
    ///             },
    ///             ClusterCodeVersion = "7.0.470.9590",
    ///             ClusterName = "myCluster",
    ///             DiagnosticsStorageAccountConfig = new AzureRM.ServiceFabric.Latest.Inputs.DiagnosticsStorageAccountConfigArgs
    ///             {
    ///                 BlobEndpoint = "https://diag.blob.core.windows.net/",
    ///                 ProtectedAccountKeyName = "StorageAccountKey1",
    ///                 QueueEndpoint = "https://diag.queue.core.windows.net/",
    ///                 StorageAccountName = "diag",
    ///                 TableEndpoint = "https://diag.table.core.windows.net/",
    ///             },
    ///             EventStoreServiceEnabled = true,
    ///             FabricSettings = 
    ///             {
    ///                 new AzureRM.ServiceFabric.Latest.Inputs.SettingsSectionDescriptionArgs
    ///                 {
    ///                     Name = "UpgradeService",
    ///                     Parameters = 
    ///                     {
    ///                         new AzureRM.ServiceFabric.Latest.Inputs.SettingsParameterDescriptionArgs
    ///                         {
    ///                             Name = "AppPollIntervalInSeconds",
    ///                             Value = "60",
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///             Location = "eastus",
    ///             ManagementEndpoint = "https://myCluster.eastus.cloudapp.azure.com:19080",
    ///             NodeTypes = 
    ///             {
    ///                 new AzureRM.ServiceFabric.Latest.Inputs.NodeTypeDescriptionArgs
    ///                 {
    ///                     ApplicationPorts = new AzureRM.ServiceFabric.Latest.Inputs.EndpointRangeDescriptionArgs
    ///                     {
    ///                         EndPort = 30000,
    ///                         StartPort = 20000,
    ///                     },
    ///                     ClientConnectionEndpointPort = 19000,
    ///                     DurabilityLevel = "Bronze",
    ///                     EphemeralPorts = new AzureRM.ServiceFabric.Latest.Inputs.EndpointRangeDescriptionArgs
    ///                     {
    ///                         EndPort = 64000,
    ///                         StartPort = 49000,
    ///                     },
    ///                     HttpGatewayEndpointPort = 19007,
    ///                     IsPrimary = true,
    ///                     Name = "nt1vm",
    ///                     VmInstanceCount = 5,
    ///                 },
    ///             },
    ///             ReliabilityLevel = "Silver",
    ///             ResourceGroupName = "resRg",
    ///             ReverseProxyCertificateCommonNames = new AzureRM.ServiceFabric.Latest.Inputs.ServerCertificateCommonNamesArgs
    ///             {
    ///                 CommonNames = 
    ///                 {
    ///                     new AzureRM.ServiceFabric.Latest.Inputs.ServerCertificateCommonNameArgs
    ///                     {
    ///                         CertificateCommonName = "abc.com",
    ///                         CertificateIssuerThumbprint = "12599211F8F14C90AFA9532AD79A6F2CA1C00622",
    ///                     },
    ///                 },
    ///                 X509StoreName = "My",
    ///             },
    ///             Tags = ,
    ///             UpgradeDescription = new AzureRM.ServiceFabric.Latest.Inputs.ClusterUpgradePolicyArgs
    ///             {
    ///                 DeltaHealthPolicy = new AzureRM.ServiceFabric.Latest.Inputs.ClusterUpgradeDeltaHealthPolicyArgs
    ///                 {
    ///                     ApplicationDeltaHealthPolicies = 
    ///                     {
    ///                         { "fabric:/myApp1", new AzureRM.ServiceFabric.Latest.Inputs.ApplicationDeltaHealthPolicyArgs
    ///                         {
    ///                             DefaultServiceTypeDeltaHealthPolicy = new AzureRM.ServiceFabric.Latest.Inputs.ServiceTypeDeltaHealthPolicyArgs
    ///                             {
    ///                                 MaxPercentDeltaUnhealthyServices = 0,
    ///                             },
    ///                             ServiceTypeDeltaHealthPolicies = 
    ///                             {
    ///                                 { "myServiceType1", new AzureRM.ServiceFabric.Latest.Inputs.ServiceTypeDeltaHealthPolicyArgs
    ///                                 {
    ///                                     MaxPercentDeltaUnhealthyServices = 0,
    ///                                 } },
    ///                             },
    ///                         } },
    ///                     },
    ///                     MaxPercentDeltaUnhealthyApplications = 0,
    ///                     MaxPercentDeltaUnhealthyNodes = 0,
    ///                     MaxPercentUpgradeDomainDeltaUnhealthyNodes = 0,
    ///                 },
    ///                 ForceRestart = false,
    ///                 HealthCheckRetryTimeout = "00:05:00",
    ///                 HealthCheckStableDuration = "00:00:30",
    ///                 HealthCheckWaitDuration = "00:00:30",
    ///                 HealthPolicy = new AzureRM.ServiceFabric.Latest.Inputs.ClusterHealthPolicyArgs
    ///                 {
    ///                     ApplicationHealthPolicies = 
    ///                     {
    ///                         { "fabric:/myApp1", new AzureRM.ServiceFabric.Latest.Inputs.ApplicationHealthPolicyArgs
    ///                         {
    ///                             DefaultServiceTypeHealthPolicy = new AzureRM.ServiceFabric.Latest.Inputs.ServiceTypeHealthPolicyArgs
    ///                             {
    ///                                 MaxPercentUnhealthyServices = 0,
    ///                             },
    ///                             ServiceTypeHealthPolicies = 
    ///                             {
    ///                                 { "myServiceType1", new AzureRM.ServiceFabric.Latest.Inputs.ServiceTypeHealthPolicyArgs
    ///                                 {
    ///                                     MaxPercentUnhealthyServices = 100,
    ///                                 } },
    ///                             },
    ///                         } },
    ///                     },
    ///                     MaxPercentUnhealthyApplications = 0,
    ///                     MaxPercentUnhealthyNodes = 0,
    ///                 },
    ///                 UpgradeDomainTimeout = "00:15:00",
    ///                 UpgradeReplicaSetCheckTimeout = "00:10:00",
    ///                 UpgradeTimeout = "01:00:00",
    ///             },
    ///             UpgradeMode = "Manual",
    ///             VmImage = "Windows",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// ### Put a cluster with minimum parameters
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var cluster = new AzureRM.ServiceFabric.Latest.Cluster("cluster", new AzureRM.ServiceFabric.Latest.ClusterArgs
    ///         {
    ///             ClusterName = "myCluster",
    ///             DiagnosticsStorageAccountConfig = new AzureRM.ServiceFabric.Latest.Inputs.DiagnosticsStorageAccountConfigArgs
    ///             {
    ///                 BlobEndpoint = "https://diag.blob.core.windows.net/",
    ///                 ProtectedAccountKeyName = "StorageAccountKey1",
    ///                 QueueEndpoint = "https://diag.queue.core.windows.net/",
    ///                 StorageAccountName = "diag",
    ///                 TableEndpoint = "https://diag.table.core.windows.net/",
    ///             },
    ///             FabricSettings = 
    ///             {
    ///                 new AzureRM.ServiceFabric.Latest.Inputs.SettingsSectionDescriptionArgs
    ///                 {
    ///                     Name = "UpgradeService",
    ///                     Parameters = 
    ///                     {
    ///                         new AzureRM.ServiceFabric.Latest.Inputs.SettingsParameterDescriptionArgs
    ///                         {
    ///                             Name = "AppPollIntervalInSeconds",
    ///                             Value = "60",
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///             Location = "eastus",
    ///             ManagementEndpoint = "http://myCluster.eastus.cloudapp.azure.com:19080",
    ///             NodeTypes = 
    ///             {
    ///                 new AzureRM.ServiceFabric.Latest.Inputs.NodeTypeDescriptionArgs
    ///                 {
    ///                     ApplicationPorts = new AzureRM.ServiceFabric.Latest.Inputs.EndpointRangeDescriptionArgs
    ///                     {
    ///                         EndPort = 30000,
    ///                         StartPort = 20000,
    ///                     },
    ///                     ClientConnectionEndpointPort = 19000,
    ///                     DurabilityLevel = "Bronze",
    ///                     EphemeralPorts = new AzureRM.ServiceFabric.Latest.Inputs.EndpointRangeDescriptionArgs
    ///                     {
    ///                         EndPort = 64000,
    ///                         StartPort = 49000,
    ///                     },
    ///                     HttpGatewayEndpointPort = 19007,
    ///                     IsPrimary = true,
    ///                     Name = "nt1vm",
    ///                     VmInstanceCount = 5,
    ///                 },
    ///             },
    ///             ReliabilityLevel = "Silver",
    ///             ResourceGroupName = "resRg",
    ///             Tags = ,
    ///             UpgradeMode = "Automatic",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class Cluster : Pulumi.CustomResource
    {
        /// <summary>
        /// The list of add-on features to enable in the cluster.
        /// </summary>
        [Output("addOnFeatures")]
        public Output<ImmutableArray<string>> AddOnFeatures { get; private set; } = null!;

        /// <summary>
        /// The policy used to clean up unused versions.
        /// </summary>
        [Output("applicationTypeVersionsCleanupPolicy")]
        public Output<Outputs.ApplicationTypeVersionsCleanupPolicyResponseResult?> ApplicationTypeVersionsCleanupPolicy { get; private set; } = null!;

        /// <summary>
        /// The Service Fabric runtime versions available for this cluster.
        /// </summary>
        [Output("availableClusterVersions")]
        public Output<ImmutableArray<Outputs.ClusterVersionDetailsResponseResult>> AvailableClusterVersions { get; private set; } = null!;

        /// <summary>
        /// The AAD authentication settings of the cluster.
        /// </summary>
        [Output("azureActiveDirectory")]
        public Output<Outputs.AzureActiveDirectoryResponseResult?> AzureActiveDirectory { get; private set; } = null!;

        /// <summary>
        /// The certificate to use for securing the cluster. The certificate provided will be used for node to node security within the cluster, SSL certificate for cluster management endpoint and default admin client.
        /// </summary>
        [Output("certificate")]
        public Output<Outputs.CertificateDescriptionResponseResult?> Certificate { get; private set; } = null!;

        /// <summary>
        /// Describes a list of server certificates referenced by common name that are used to secure the cluster.
        /// </summary>
        [Output("certificateCommonNames")]
        public Output<Outputs.ServerCertificateCommonNamesResponseResult?> CertificateCommonNames { get; private set; } = null!;

        /// <summary>
        /// The list of client certificates referenced by common name that are allowed to manage the cluster.
        /// </summary>
        [Output("clientCertificateCommonNames")]
        public Output<ImmutableArray<Outputs.ClientCertificateCommonNameResponseResult>> ClientCertificateCommonNames { get; private set; } = null!;

        /// <summary>
        /// The list of client certificates referenced by thumbprint that are allowed to manage the cluster.
        /// </summary>
        [Output("clientCertificateThumbprints")]
        public Output<ImmutableArray<Outputs.ClientCertificateThumbprintResponseResult>> ClientCertificateThumbprints { get; private set; } = null!;

        /// <summary>
        /// The Service Fabric runtime version of the cluster. This property can only by set the user when **upgradeMode** is set to 'Manual'. To get list of available Service Fabric versions for new clusters use [ClusterVersion API](./ClusterVersion.md). To get the list of available version for existing clusters use **availableClusterVersions**.
        /// </summary>
        [Output("clusterCodeVersion")]
        public Output<string?> ClusterCodeVersion { get; private set; } = null!;

        /// <summary>
        /// The Azure Resource Provider endpoint. A system service in the cluster connects to this  endpoint.
        /// </summary>
        [Output("clusterEndpoint")]
        public Output<string> ClusterEndpoint { get; private set; } = null!;

        /// <summary>
        /// A service generated unique identifier for the cluster resource.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

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
        [Output("clusterState")]
        public Output<string> ClusterState { get; private set; } = null!;

        /// <summary>
        /// The storage account information for storing Service Fabric diagnostic logs.
        /// </summary>
        [Output("diagnosticsStorageAccountConfig")]
        public Output<Outputs.DiagnosticsStorageAccountConfigResponseResult?> DiagnosticsStorageAccountConfig { get; private set; } = null!;

        /// <summary>
        /// Azure resource etag.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// Indicates if the event store service is enabled.
        /// </summary>
        [Output("eventStoreServiceEnabled")]
        public Output<bool?> EventStoreServiceEnabled { get; private set; } = null!;

        /// <summary>
        /// The list of custom fabric settings to configure the cluster.
        /// </summary>
        [Output("fabricSettings")]
        public Output<ImmutableArray<Outputs.SettingsSectionDescriptionResponseResult>> FabricSettings { get; private set; } = null!;

        /// <summary>
        /// Azure resource location.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The http management endpoint of the cluster.
        /// </summary>
        [Output("managementEndpoint")]
        public Output<string> ManagementEndpoint { get; private set; } = null!;

        /// <summary>
        /// Azure resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The list of node types in the cluster.
        /// </summary>
        [Output("nodeTypes")]
        public Output<ImmutableArray<Outputs.NodeTypeDescriptionResponseResult>> NodeTypes { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the cluster resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The reliability level sets the replica set size of system services. Learn about [ReliabilityLevel](https://docs.microsoft.com/azure/service-fabric/service-fabric-cluster-capacity).
        /// 
        ///   - None - Run the System services with a target replica set count of 1. This should only be used for test clusters.
        ///   - Bronze - Run the System services with a target replica set count of 3. This should only be used for test clusters.
        ///   - Silver - Run the System services with a target replica set count of 5.
        ///   - Gold - Run the System services with a target replica set count of 7.
        ///   - Platinum - Run the System services with a target replica set count of 9.
        /// </summary>
        [Output("reliabilityLevel")]
        public Output<string?> ReliabilityLevel { get; private set; } = null!;

        /// <summary>
        /// The server certificate used by reverse proxy.
        /// </summary>
        [Output("reverseProxyCertificate")]
        public Output<Outputs.CertificateDescriptionResponseResult?> ReverseProxyCertificate { get; private set; } = null!;

        /// <summary>
        /// Describes a list of server certificates referenced by common name that are used to secure the cluster.
        /// </summary>
        [Output("reverseProxyCertificateCommonNames")]
        public Output<Outputs.ServerCertificateCommonNamesResponseResult?> ReverseProxyCertificateCommonNames { get; private set; } = null!;

        /// <summary>
        /// Azure resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Azure resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The policy to use when upgrading the cluster.
        /// </summary>
        [Output("upgradeDescription")]
        public Output<Outputs.ClusterUpgradePolicyResponseResult?> UpgradeDescription { get; private set; } = null!;

        /// <summary>
        /// The upgrade mode of the cluster when new Service Fabric runtime version is available.
        /// 
        ///   - Automatic - The cluster will be automatically upgraded to the latest Service Fabric runtime version as soon as it is available.
        ///   - Manual - The cluster will not be automatically upgraded to the latest Service Fabric runtime version. The cluster is upgraded by setting the **clusterCodeVersion** property in the cluster resource.
        /// </summary>
        [Output("upgradeMode")]
        public Output<string?> UpgradeMode { get; private set; } = null!;

        /// <summary>
        /// The VM image VMSS has been configured with. Generic names such as Windows or Linux can be used.
        /// </summary>
        [Output("vmImage")]
        public Output<string?> VmImage { get; private set; } = null!;


        /// <summary>
        /// Create a Cluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Cluster(string name, ClusterArgs args, CustomResourceOptions? options = null)
            : base("azurerm:servicefabric/latest:Cluster", name, args ?? new ClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Cluster(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:servicefabric/latest:Cluster", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:servicefabric/v20160901:Cluster"},
                    new Pulumi.Alias { Type = "azurerm:servicefabric/v20170701preview:Cluster"},
                    new Pulumi.Alias { Type = "azurerm:servicefabric/v20180201:Cluster"},
                    new Pulumi.Alias { Type = "azurerm:servicefabric/v20190301:Cluster"},
                    new Pulumi.Alias { Type = "azurerm:servicefabric/v20190301preview:Cluster"},
                    new Pulumi.Alias { Type = "azurerm:servicefabric/v20190601preview:Cluster"},
                    new Pulumi.Alias { Type = "azurerm:servicefabric/v20191101preview:Cluster"},
                    new Pulumi.Alias { Type = "azurerm:servicefabric/v20200301:Cluster"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Cluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Cluster Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Cluster(name, id, options);
        }
    }

    public sealed class ClusterArgs : Pulumi.ResourceArgs
    {
        [Input("addOnFeatures")]
        private InputList<string>? _addOnFeatures;

        /// <summary>
        /// The list of add-on features to enable in the cluster.
        /// </summary>
        public InputList<string> AddOnFeatures
        {
            get => _addOnFeatures ?? (_addOnFeatures = new InputList<string>());
            set => _addOnFeatures = value;
        }

        /// <summary>
        /// The policy used to clean up unused versions.
        /// </summary>
        [Input("applicationTypeVersionsCleanupPolicy")]
        public Input<Inputs.ApplicationTypeVersionsCleanupPolicyArgs>? ApplicationTypeVersionsCleanupPolicy { get; set; }

        /// <summary>
        /// The AAD authentication settings of the cluster.
        /// </summary>
        [Input("azureActiveDirectory")]
        public Input<Inputs.AzureActiveDirectoryArgs>? AzureActiveDirectory { get; set; }

        /// <summary>
        /// The certificate to use for securing the cluster. The certificate provided will be used for node to node security within the cluster, SSL certificate for cluster management endpoint and default admin client.
        /// </summary>
        [Input("certificate")]
        public Input<Inputs.CertificateDescriptionArgs>? Certificate { get; set; }

        /// <summary>
        /// Describes a list of server certificates referenced by common name that are used to secure the cluster.
        /// </summary>
        [Input("certificateCommonNames")]
        public Input<Inputs.ServerCertificateCommonNamesArgs>? CertificateCommonNames { get; set; }

        [Input("clientCertificateCommonNames")]
        private InputList<Inputs.ClientCertificateCommonNameArgs>? _clientCertificateCommonNames;

        /// <summary>
        /// The list of client certificates referenced by common name that are allowed to manage the cluster.
        /// </summary>
        public InputList<Inputs.ClientCertificateCommonNameArgs> ClientCertificateCommonNames
        {
            get => _clientCertificateCommonNames ?? (_clientCertificateCommonNames = new InputList<Inputs.ClientCertificateCommonNameArgs>());
            set => _clientCertificateCommonNames = value;
        }

        [Input("clientCertificateThumbprints")]
        private InputList<Inputs.ClientCertificateThumbprintArgs>? _clientCertificateThumbprints;

        /// <summary>
        /// The list of client certificates referenced by thumbprint that are allowed to manage the cluster.
        /// </summary>
        public InputList<Inputs.ClientCertificateThumbprintArgs> ClientCertificateThumbprints
        {
            get => _clientCertificateThumbprints ?? (_clientCertificateThumbprints = new InputList<Inputs.ClientCertificateThumbprintArgs>());
            set => _clientCertificateThumbprints = value;
        }

        /// <summary>
        /// The Service Fabric runtime version of the cluster. This property can only by set the user when **upgradeMode** is set to 'Manual'. To get list of available Service Fabric versions for new clusters use [ClusterVersion API](./ClusterVersion.md). To get the list of available version for existing clusters use **availableClusterVersions**.
        /// </summary>
        [Input("clusterCodeVersion")]
        public Input<string>? ClusterCodeVersion { get; set; }

        /// <summary>
        /// The name of the cluster resource.
        /// </summary>
        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        /// <summary>
        /// The storage account information for storing Service Fabric diagnostic logs.
        /// </summary>
        [Input("diagnosticsStorageAccountConfig")]
        public Input<Inputs.DiagnosticsStorageAccountConfigArgs>? DiagnosticsStorageAccountConfig { get; set; }

        /// <summary>
        /// Indicates if the event store service is enabled.
        /// </summary>
        [Input("eventStoreServiceEnabled")]
        public Input<bool>? EventStoreServiceEnabled { get; set; }

        [Input("fabricSettings")]
        private InputList<Inputs.SettingsSectionDescriptionArgs>? _fabricSettings;

        /// <summary>
        /// The list of custom fabric settings to configure the cluster.
        /// </summary>
        public InputList<Inputs.SettingsSectionDescriptionArgs> FabricSettings
        {
            get => _fabricSettings ?? (_fabricSettings = new InputList<Inputs.SettingsSectionDescriptionArgs>());
            set => _fabricSettings = value;
        }

        /// <summary>
        /// Azure resource location.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The http management endpoint of the cluster.
        /// </summary>
        [Input("managementEndpoint", required: true)]
        public Input<string> ManagementEndpoint { get; set; } = null!;

        [Input("nodeTypes", required: true)]
        private InputList<Inputs.NodeTypeDescriptionArgs>? _nodeTypes;

        /// <summary>
        /// The list of node types in the cluster.
        /// </summary>
        public InputList<Inputs.NodeTypeDescriptionArgs> NodeTypes
        {
            get => _nodeTypes ?? (_nodeTypes = new InputList<Inputs.NodeTypeDescriptionArgs>());
            set => _nodeTypes = value;
        }

        /// <summary>
        /// The reliability level sets the replica set size of system services. Learn about [ReliabilityLevel](https://docs.microsoft.com/azure/service-fabric/service-fabric-cluster-capacity).
        /// 
        ///   - None - Run the System services with a target replica set count of 1. This should only be used for test clusters.
        ///   - Bronze - Run the System services with a target replica set count of 3. This should only be used for test clusters.
        ///   - Silver - Run the System services with a target replica set count of 5.
        ///   - Gold - Run the System services with a target replica set count of 7.
        ///   - Platinum - Run the System services with a target replica set count of 9.
        /// </summary>
        [Input("reliabilityLevel")]
        public Input<string>? ReliabilityLevel { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The server certificate used by reverse proxy.
        /// </summary>
        [Input("reverseProxyCertificate")]
        public Input<Inputs.CertificateDescriptionArgs>? ReverseProxyCertificate { get; set; }

        /// <summary>
        /// Describes a list of server certificates referenced by common name that are used to secure the cluster.
        /// </summary>
        [Input("reverseProxyCertificateCommonNames")]
        public Input<Inputs.ServerCertificateCommonNamesArgs>? ReverseProxyCertificateCommonNames { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Azure resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The policy to use when upgrading the cluster.
        /// </summary>
        [Input("upgradeDescription")]
        public Input<Inputs.ClusterUpgradePolicyArgs>? UpgradeDescription { get; set; }

        /// <summary>
        /// The upgrade mode of the cluster when new Service Fabric runtime version is available.
        /// 
        ///   - Automatic - The cluster will be automatically upgraded to the latest Service Fabric runtime version as soon as it is available.
        ///   - Manual - The cluster will not be automatically upgraded to the latest Service Fabric runtime version. The cluster is upgraded by setting the **clusterCodeVersion** property in the cluster resource.
        /// </summary>
        [Input("upgradeMode")]
        public Input<string>? UpgradeMode { get; set; }

        /// <summary>
        /// The VM image VMSS has been configured with. Generic names such as Windows or Linux can be used.
        /// </summary>
        [Input("vmImage")]
        public Input<string>? VmImage { get; set; }

        public ClusterArgs()
        {
        }
    }
}
