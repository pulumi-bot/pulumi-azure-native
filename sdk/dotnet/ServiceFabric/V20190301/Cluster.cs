// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceFabric.V20190301
{
    /// <summary>
    /// The cluster resource
    /// </summary>
    public partial class Cluster : Pulumi.CustomResource
    {
        /// <summary>
        /// Azure resource etag.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

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
        /// The cluster resource properties
        /// </summary>
        [Output("properties")]
        public Output<Outputs.ClusterPropertiesResponseResult> Properties { get; private set; } = null!;

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
        /// Create a Cluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Cluster(string name, ClusterArgs args, CustomResourceOptions? options = null)
            : base("azurerm:servicefabric/v20190301:Cluster", name, args ?? new ClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Cluster(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:servicefabric/v20190301:Cluster", name, null, MakeResourceOptions(options, id))
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

        /// <summary>
        /// The name of the cluster resource.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

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
