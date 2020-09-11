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
    /// Describes a node type in the cluster, each node type represents sub set of nodes in the cluster.
    /// </summary>
    public partial class NodeType : Pulumi.CustomResource
    {
        /// <summary>
        /// The range of ports from which cluster assigned port to Service Fabric applications.
        /// </summary>
        [Output("applicationPorts")]
        public Output<Outputs.EndpointRangeDescriptionResponseResult?> ApplicationPorts { get; private set; } = null!;

        /// <summary>
        /// The capacity tags applied to the nodes in the node type, the cluster resource manager uses these tags to understand how much resource a node has.
        /// </summary>
        [Output("capacities")]
        public Output<ImmutableDictionary<string, string>?> Capacities { get; private set; } = null!;

        /// <summary>
        /// Disk size for each vm in the node type in GBs.
        /// </summary>
        [Output("dataDiskSizeGB")]
        public Output<int> DataDiskSizeGB { get; private set; } = null!;

        /// <summary>
        /// The range of ephemeral ports that nodes in this node type should be configured with.
        /// </summary>
        [Output("ephemeralPorts")]
        public Output<Outputs.EndpointRangeDescriptionResponseResult?> EphemeralPorts { get; private set; } = null!;

        /// <summary>
        /// The node type on which system services will run. Only one node type should be marked as primary. Primary node type cannot be deleted or changed for existing clusters.
        /// </summary>
        [Output("isPrimary")]
        public Output<bool> IsPrimary { get; private set; } = null!;

        /// <summary>
        /// Azure resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The placement tags applied to nodes in the node type, which can be used to indicate where certain services (workload) should run.
        /// </summary>
        [Output("placementProperties")]
        public Output<ImmutableDictionary<string, string>?> PlacementProperties { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the managed cluster resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

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
        /// Set of extensions that should be installed onto the virtual machines.
        /// </summary>
        [Output("vmExtensions")]
        public Output<ImmutableArray<Outputs.VMSSExtensionResponseResult>> VmExtensions { get; private set; } = null!;

        /// <summary>
        /// The offer type of the Azure Virtual Machines Marketplace image. For example, UbuntuServer or WindowsServer.
        /// </summary>
        [Output("vmImageOffer")]
        public Output<string?> VmImageOffer { get; private set; } = null!;

        /// <summary>
        /// The publisher of the Azure Virtual Machines Marketplace image. For example, Canonical or MicrosoftWindowsServer.
        /// </summary>
        [Output("vmImagePublisher")]
        public Output<string?> VmImagePublisher { get; private set; } = null!;

        /// <summary>
        /// The SKU of the Azure Virtual Machines Marketplace image. For example, 14.04.0-LTS or 2012-R2-Datacenter.
        /// </summary>
        [Output("vmImageSku")]
        public Output<string?> VmImageSku { get; private set; } = null!;

        /// <summary>
        /// The version of the Azure Virtual Machines Marketplace image. A value of 'latest' can be specified to select the latest version of an image. If omitted, the default is 'latest'.
        /// </summary>
        [Output("vmImageVersion")]
        public Output<string?> VmImageVersion { get; private set; } = null!;

        /// <summary>
        /// The number of nodes in the node type.
        /// </summary>
        [Output("vmInstanceCount")]
        public Output<int> VmInstanceCount { get; private set; } = null!;

        /// <summary>
        /// The secrets to install in the virtual machines.
        /// </summary>
        [Output("vmSecrets")]
        public Output<ImmutableArray<Outputs.VaultSecretGroupResponseResult>> VmSecrets { get; private set; } = null!;

        /// <summary>
        /// The size of virtual machines in the pool. All virtual machines in a pool are the same size. For example, Standard_D3.
        /// </summary>
        [Output("vmSize")]
        public Output<string?> VmSize { get; private set; } = null!;


        /// <summary>
        /// Create a NodeType resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NodeType(string name, NodeTypeArgs args, CustomResourceOptions? options = null)
            : base("azurerm:servicefabric/v20200101preview:NodeType", name, args ?? new NodeTypeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NodeType(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:servicefabric/v20200101preview:NodeType", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing NodeType resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NodeType Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new NodeType(name, id, options);
        }
    }

    public sealed class NodeTypeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The range of ports from which cluster assigned port to Service Fabric applications.
        /// </summary>
        [Input("applicationPorts")]
        public Input<Inputs.EndpointRangeDescriptionArgs>? ApplicationPorts { get; set; }

        [Input("capacities")]
        private InputMap<string>? _capacities;

        /// <summary>
        /// The capacity tags applied to the nodes in the node type, the cluster resource manager uses these tags to understand how much resource a node has.
        /// </summary>
        public InputMap<string> Capacities
        {
            get => _capacities ?? (_capacities = new InputMap<string>());
            set => _capacities = value;
        }

        /// <summary>
        /// The name of the cluster resource.
        /// </summary>
        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        /// <summary>
        /// Disk size for each vm in the node type in GBs.
        /// </summary>
        [Input("dataDiskSizeGB", required: true)]
        public Input<int> DataDiskSizeGB { get; set; } = null!;

        /// <summary>
        /// The range of ephemeral ports that nodes in this node type should be configured with.
        /// </summary>
        [Input("ephemeralPorts")]
        public Input<Inputs.EndpointRangeDescriptionArgs>? EphemeralPorts { get; set; }

        /// <summary>
        /// The node type on which system services will run. Only one node type should be marked as primary. Primary node type cannot be deleted or changed for existing clusters.
        /// </summary>
        [Input("isPrimary", required: true)]
        public Input<bool> IsPrimary { get; set; } = null!;

        /// <summary>
        /// The name of the node type.
        /// </summary>
        [Input("nodeTypeName", required: true)]
        public Input<string> NodeTypeName { get; set; } = null!;

        [Input("placementProperties")]
        private InputMap<string>? _placementProperties;

        /// <summary>
        /// The placement tags applied to nodes in the node type, which can be used to indicate where certain services (workload) should run.
        /// </summary>
        public InputMap<string> PlacementProperties
        {
            get => _placementProperties ?? (_placementProperties = new InputMap<string>());
            set => _placementProperties = value;
        }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

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

        [Input("vmExtensions")]
        private InputList<Inputs.VMSSExtensionArgs>? _vmExtensions;

        /// <summary>
        /// Set of extensions that should be installed onto the virtual machines.
        /// </summary>
        public InputList<Inputs.VMSSExtensionArgs> VmExtensions
        {
            get => _vmExtensions ?? (_vmExtensions = new InputList<Inputs.VMSSExtensionArgs>());
            set => _vmExtensions = value;
        }

        /// <summary>
        /// The offer type of the Azure Virtual Machines Marketplace image. For example, UbuntuServer or WindowsServer.
        /// </summary>
        [Input("vmImageOffer")]
        public Input<string>? VmImageOffer { get; set; }

        /// <summary>
        /// The publisher of the Azure Virtual Machines Marketplace image. For example, Canonical or MicrosoftWindowsServer.
        /// </summary>
        [Input("vmImagePublisher")]
        public Input<string>? VmImagePublisher { get; set; }

        /// <summary>
        /// The SKU of the Azure Virtual Machines Marketplace image. For example, 14.04.0-LTS or 2012-R2-Datacenter.
        /// </summary>
        [Input("vmImageSku")]
        public Input<string>? VmImageSku { get; set; }

        /// <summary>
        /// The version of the Azure Virtual Machines Marketplace image. A value of 'latest' can be specified to select the latest version of an image. If omitted, the default is 'latest'.
        /// </summary>
        [Input("vmImageVersion")]
        public Input<string>? VmImageVersion { get; set; }

        /// <summary>
        /// The number of nodes in the node type.
        /// </summary>
        [Input("vmInstanceCount", required: true)]
        public Input<int> VmInstanceCount { get; set; } = null!;

        [Input("vmSecrets")]
        private InputList<Inputs.VaultSecretGroupArgs>? _vmSecrets;

        /// <summary>
        /// The secrets to install in the virtual machines.
        /// </summary>
        public InputList<Inputs.VaultSecretGroupArgs> VmSecrets
        {
            get => _vmSecrets ?? (_vmSecrets = new InputList<Inputs.VaultSecretGroupArgs>());
            set => _vmSecrets = value;
        }

        /// <summary>
        /// The size of virtual machines in the pool. All virtual machines in a pool are the same size. For example, Standard_D3.
        /// </summary>
        [Input("vmSize")]
        public Input<string>? VmSize { get; set; }

        public NodeTypeArgs()
        {
        }
    }
}
