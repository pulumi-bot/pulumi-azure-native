// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceFabric.V20200101Preview
{
    /// <summary>
    /// The manged cluster resource
    /// </summary>
    public partial class ManagedCluster : Pulumi.CustomResource
    {
        /// <summary>
        /// vm admin user password.
        /// </summary>
        [Output("adminPassword")]
        public Output<string?> AdminPassword { get; private set; } = null!;

        /// <summary>
        /// vm admin user name.
        /// </summary>
        [Output("adminUserName")]
        public Output<string> AdminUserName { get; private set; } = null!;

        /// <summary>
        /// Azure active directory.
        /// </summary>
        [Output("azureActiveDirectory")]
        public Output<Outputs.AzureActiveDirectoryResponseResult?> AzureActiveDirectory { get; private set; } = null!;

        /// <summary>
        /// The port used for client connections to the cluster.
        /// </summary>
        [Output("clientConnectionPort")]
        public Output<int?> ClientConnectionPort { get; private set; } = null!;

        /// <summary>
        /// client certificates for the cluster.
        /// </summary>
        [Output("clients")]
        public Output<ImmutableArray<Outputs.ClientCertificateResponseResult>> Clients { get; private set; } = null!;

        /// <summary>
        /// The cluster certificate thumbprint used node to node communication.
        /// </summary>
        [Output("clusterCertificateThumbprint")]
        public Output<string> ClusterCertificateThumbprint { get; private set; } = null!;

        /// <summary>
        /// The Service Fabric runtime version of the cluster. This property can only by set the user when **upgradeMode** is set to 'Manual'. To get list of available Service Fabric versions for new clusters use [ClusterVersion API](./ClusterVersion.md). To get the list of available version for existing clusters use **availableClusterVersions**.
        /// </summary>
        [Output("clusterCodeVersion")]
        public Output<string?> ClusterCodeVersion { get; private set; } = null!;

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
        /// Describes the policy used when upgrading the cluster.
        /// </summary>
        [Output("clusterUpgradeDescription")]
        public Output<Outputs.ClusterUpgradePolicyResponseResult?> ClusterUpgradeDescription { get; private set; } = null!;

        /// <summary>
        /// The upgrade mode of the cluster when new Service Fabric runtime version is available.
        /// 
        ///   - Automatic - The cluster will be automatically upgraded to the latest Service Fabric runtime version as soon as it is available.
        ///   - Manual - The cluster will not be automatically upgraded to the latest Service Fabric runtime version. The cluster is upgraded by setting the **clusterCodeVersion** property in the cluster resource.
        /// </summary>
        [Output("clusterUpgradeMode")]
        public Output<string?> ClusterUpgradeMode { get; private set; } = null!;

        /// <summary>
        /// The cluster dns name.
        /// </summary>
        [Output("dnsName")]
        public Output<string> DnsName { get; private set; } = null!;

        /// <summary>
        /// Azure resource etag.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The list of custom fabric settings to configure the cluster.
        /// </summary>
        [Output("fabricSettings")]
        public Output<ImmutableArray<Outputs.SettingsSectionDescriptionResponseResult>> FabricSettings { get; private set; } = null!;

        /// <summary>
        /// the cluster Fully qualified domain name.
        /// </summary>
        [Output("fqdn")]
        public Output<string> Fqdn { get; private set; } = null!;

        /// <summary>
        /// The port used for http connections to the cluster.
        /// </summary>
        [Output("httpGatewayConnectionPort")]
        public Output<int?> HttpGatewayConnectionPort { get; private set; } = null!;

        /// <summary>
        /// Describes load balancing rules.
        /// </summary>
        [Output("loadBalancingRules")]
        public Output<ImmutableArray<Outputs.LoadBalancingRuleResponseResult>> LoadBalancingRules { get; private set; } = null!;

        /// <summary>
        /// Azure resource location.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Azure resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the managed cluster resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The endpoint used by reverse proxy.
        /// </summary>
        [Output("reverseProxyEndpointPort")]
        public Output<int?> ReverseProxyEndpointPort { get; private set; } = null!;

        /// <summary>
        /// The sku of the managed cluster
        /// </summary>
        [Output("sku")]
        public Output<Outputs.SkuResponseResult?> Sku { get; private set; } = null!;

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
        /// Use service fabric test vm extension, by default it's false.
        /// </summary>
        [Output("useTestExtension")]
        public Output<bool?> UseTestExtension { get; private set; } = null!;


        /// <summary>
        /// Create a ManagedCluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ManagedCluster(string name, ManagedClusterArgs args, CustomResourceOptions? options = null)
            : base("azurerm:servicefabric/v20200101preview:ManagedCluster", name, args ?? new ManagedClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ManagedCluster(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:servicefabric/v20200101preview:ManagedCluster", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ManagedCluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ManagedCluster Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ManagedCluster(name, id, options);
        }
    }

    public sealed class ManagedClusterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// vm admin user password.
        /// </summary>
        [Input("adminPassword")]
        public Input<string>? AdminPassword { get; set; }

        /// <summary>
        /// vm admin user name.
        /// </summary>
        [Input("adminUserName", required: true)]
        public Input<string> AdminUserName { get; set; } = null!;

        /// <summary>
        /// Azure active directory.
        /// </summary>
        [Input("azureActiveDirectory")]
        public Input<Inputs.AzureActiveDirectoryArgs>? AzureActiveDirectory { get; set; }

        /// <summary>
        /// The port used for client connections to the cluster.
        /// </summary>
        [Input("clientConnectionPort")]
        public Input<int>? ClientConnectionPort { get; set; }

        [Input("clients")]
        private InputList<Inputs.ClientCertificateArgs>? _clients;

        /// <summary>
        /// client certificates for the cluster.
        /// </summary>
        public InputList<Inputs.ClientCertificateArgs> Clients
        {
            get => _clients ?? (_clients = new InputList<Inputs.ClientCertificateArgs>());
            set => _clients = value;
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
        /// Describes the policy used when upgrading the cluster.
        /// </summary>
        [Input("clusterUpgradeDescription")]
        public Input<Inputs.ClusterUpgradePolicyArgs>? ClusterUpgradeDescription { get; set; }

        /// <summary>
        /// The upgrade mode of the cluster when new Service Fabric runtime version is available.
        /// 
        ///   - Automatic - The cluster will be automatically upgraded to the latest Service Fabric runtime version as soon as it is available.
        ///   - Manual - The cluster will not be automatically upgraded to the latest Service Fabric runtime version. The cluster is upgraded by setting the **clusterCodeVersion** property in the cluster resource.
        /// </summary>
        [Input("clusterUpgradeMode")]
        public Input<string>? ClusterUpgradeMode { get; set; }

        /// <summary>
        /// The cluster dns name.
        /// </summary>
        [Input("dnsName", required: true)]
        public Input<string> DnsName { get; set; } = null!;

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
        /// The port used for http connections to the cluster.
        /// </summary>
        [Input("httpGatewayConnectionPort")]
        public Input<int>? HttpGatewayConnectionPort { get; set; }

        [Input("loadBalancingRules")]
        private InputList<Inputs.LoadBalancingRuleArgs>? _loadBalancingRules;

        /// <summary>
        /// Describes load balancing rules.
        /// </summary>
        public InputList<Inputs.LoadBalancingRuleArgs> LoadBalancingRules
        {
            get => _loadBalancingRules ?? (_loadBalancingRules = new InputList<Inputs.LoadBalancingRuleArgs>());
            set => _loadBalancingRules = value;
        }

        /// <summary>
        /// Azure resource location.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The endpoint used by reverse proxy.
        /// </summary>
        [Input("reverseProxyEndpointPort")]
        public Input<int>? ReverseProxyEndpointPort { get; set; }

        /// <summary>
        /// The sku of the managed cluster
        /// </summary>
        [Input("sku")]
        public Input<Inputs.SkuArgs>? Sku { get; set; }

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
        /// Use service fabric test vm extension, by default it's false.
        /// </summary>
        [Input("useTestExtension")]
        public Input<bool>? UseTestExtension { get; set; }

        public ManagedClusterArgs()
        {
        }
    }
}
