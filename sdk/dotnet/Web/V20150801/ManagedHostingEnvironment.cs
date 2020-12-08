// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Web.V20150801
{
    /// <summary>
    /// Description of an hostingEnvironment (App Service Environment)
    /// </summary>
    public partial class ManagedHostingEnvironment : Pulumi.CustomResource
    {
        /// <summary>
        /// List of comma separated strings describing which VM sizes are allowed for front-ends
        /// </summary>
        [Output("allowedMultiSizes")]
        public Output<string?> AllowedMultiSizes { get; private set; } = null!;

        /// <summary>
        /// List of comma separated strings describing which VM sizes are allowed for workers
        /// </summary>
        [Output("allowedWorkerSizes")]
        public Output<string?> AllowedWorkerSizes { get; private set; } = null!;

        /// <summary>
        /// Api Management Account associated with this Hosting Environment
        /// </summary>
        [Output("apiManagementAccountId")]
        public Output<string?> ApiManagementAccountId { get; private set; } = null!;

        /// <summary>
        /// Custom settings for changing the behavior of the hosting environment
        /// </summary>
        [Output("clusterSettings")]
        public Output<ImmutableArray<Outputs.NameValuePairResponse>> ClusterSettings { get; private set; } = null!;

        /// <summary>
        /// Edition of the metadata database for the hostingEnvironment (App Service Environment) e.g. "Standard"
        /// </summary>
        [Output("databaseEdition")]
        public Output<string?> DatabaseEdition { get; private set; } = null!;

        /// <summary>
        /// Service objective of the metadata database for the hostingEnvironment (App Service Environment) e.g. "S0"
        /// </summary>
        [Output("databaseServiceObjective")]
        public Output<string?> DatabaseServiceObjective { get; private set; } = null!;

        /// <summary>
        /// DNS suffix of the hostingEnvironment (App Service Environment)
        /// </summary>
        [Output("dnsSuffix")]
        public Output<string?> DnsSuffix { get; private set; } = null!;

        /// <summary>
        /// Current total, used, and available worker capacities
        /// </summary>
        [Output("environmentCapacities")]
        public Output<ImmutableArray<Outputs.StampCapacityResponse>> EnvironmentCapacities { get; private set; } = null!;

        /// <summary>
        /// True/false indicating whether the hostingEnvironment (App Service Environment) is healthy
        /// </summary>
        [Output("environmentIsHealthy")]
        public Output<bool?> EnvironmentIsHealthy { get; private set; } = null!;

        /// <summary>
        /// Detailed message about with results of the last check of the hostingEnvironment (App Service Environment)
        /// </summary>
        [Output("environmentStatus")]
        public Output<string?> EnvironmentStatus { get; private set; } = null!;

        /// <summary>
        /// Specifies which endpoints to serve internally in the hostingEnvironment's (App Service Environment) VNET
        /// </summary>
        [Output("internalLoadBalancingMode")]
        public Output<string?> InternalLoadBalancingMode { get; private set; } = null!;

        /// <summary>
        /// Number of IP SSL addresses reserved for this hostingEnvironment (App Service Environment)
        /// </summary>
        [Output("ipsslAddressCount")]
        public Output<int?> IpsslAddressCount { get; private set; } = null!;

        /// <summary>
        /// Kind of resource
        /// </summary>
        [Output("kind")]
        public Output<string?> Kind { get; private set; } = null!;

        /// <summary>
        /// Last deployment action on this hostingEnvironment (App Service Environment)
        /// </summary>
        [Output("lastAction")]
        public Output<string?> LastAction { get; private set; } = null!;

        /// <summary>
        /// Result of the last deployment action on this hostingEnvironment (App Service Environment)
        /// </summary>
        [Output("lastActionResult")]
        public Output<string?> LastActionResult { get; private set; } = null!;

        /// <summary>
        /// Resource Location
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Maximum number of VMs in this hostingEnvironment (App Service Environment)
        /// </summary>
        [Output("maximumNumberOfMachines")]
        public Output<int?> MaximumNumberOfMachines { get; private set; } = null!;

        /// <summary>
        /// Number of front-end instances
        /// </summary>
        [Output("multiRoleCount")]
        public Output<int?> MultiRoleCount { get; private set; } = null!;

        /// <summary>
        /// Front-end VM size, e.g. "Medium", "Large"
        /// </summary>
        [Output("multiSize")]
        public Output<string?> MultiSize { get; private set; } = null!;

        /// <summary>
        /// Resource Name
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// Access control list for controlling traffic to the hostingEnvironment (App Service Environment)
        /// </summary>
        [Output("networkAccessControlList")]
        public Output<ImmutableArray<Outputs.NetworkAccessControlEntryResponse>> NetworkAccessControlList { get; private set; } = null!;

        /// <summary>
        /// Provisioning state of the hostingEnvironment (App Service Environment)
        /// </summary>
        [Output("provisioningState")]
        public Output<string?> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Resource group of the hostingEnvironment (App Service Environment)
        /// </summary>
        [Output("resourceGroup")]
        public Output<string?> ResourceGroup { get; private set; } = null!;

        /// <summary>
        /// Current status of the hostingEnvironment (App Service Environment)
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Subscription of the hostingEnvironment (App Service Environment)
        /// </summary>
        [Output("subscriptionId")]
        public Output<string?> SubscriptionId { get; private set; } = null!;

        /// <summary>
        /// True/false indicating whether the hostingEnvironment is suspended. The environment can be suspended e.g. when the management endpoint is no longer available
        ///             (most likely because NSG blocked the incoming traffic)
        /// </summary>
        [Output("suspended")]
        public Output<bool?> Suspended { get; private set; } = null!;

        /// <summary>
        /// Resource tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Resource type
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;

        /// <summary>
        /// Number of upgrade domains of this hostingEnvironment (App Service Environment)
        /// </summary>
        [Output("upgradeDomains")]
        public Output<int?> UpgradeDomains { get; private set; } = null!;

        /// <summary>
        /// Description of IP SSL mapping for this hostingEnvironment (App Service Environment)
        /// </summary>
        [Output("vipMappings")]
        public Output<ImmutableArray<Outputs.VirtualIPMappingResponse>> VipMappings { get; private set; } = null!;

        /// <summary>
        /// Description of the hostingEnvironment's (App Service Environment) virtual network
        /// </summary>
        [Output("virtualNetwork")]
        public Output<Outputs.VirtualNetworkProfileResponse?> VirtualNetwork { get; private set; } = null!;

        /// <summary>
        /// Name of the hostingEnvironment's (App Service Environment) virtual network
        /// </summary>
        [Output("vnetName")]
        public Output<string?> VnetName { get; private set; } = null!;

        /// <summary>
        /// Resource group of the hostingEnvironment's (App Service Environment) virtual network
        /// </summary>
        [Output("vnetResourceGroupName")]
        public Output<string?> VnetResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// Subnet of the hostingEnvironment's (App Service Environment) virtual network
        /// </summary>
        [Output("vnetSubnetName")]
        public Output<string?> VnetSubnetName { get; private set; } = null!;

        /// <summary>
        /// Description of worker pools with worker size ids, VM sizes, and number of workers in each pool
        /// </summary>
        [Output("workerPools")]
        public Output<ImmutableArray<Outputs.WorkerPoolResponse>> WorkerPools { get; private set; } = null!;


        /// <summary>
        /// Create a ManagedHostingEnvironment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ManagedHostingEnvironment(string name, ManagedHostingEnvironmentArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:web/v20150801:ManagedHostingEnvironment", name, args ?? new ManagedHostingEnvironmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ManagedHostingEnvironment(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:web/v20150801:ManagedHostingEnvironment", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:web/latest:ManagedHostingEnvironment"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ManagedHostingEnvironment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ManagedHostingEnvironment Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ManagedHostingEnvironment(name, id, options);
        }
    }

    public sealed class ManagedHostingEnvironmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// List of comma separated strings describing which VM sizes are allowed for front-ends
        /// </summary>
        [Input("allowedMultiSizes")]
        public Input<string>? AllowedMultiSizes { get; set; }

        /// <summary>
        /// List of comma separated strings describing which VM sizes are allowed for workers
        /// </summary>
        [Input("allowedWorkerSizes")]
        public Input<string>? AllowedWorkerSizes { get; set; }

        /// <summary>
        /// Api Management Account associated with this Hosting Environment
        /// </summary>
        [Input("apiManagementAccountId")]
        public Input<string>? ApiManagementAccountId { get; set; }

        [Input("clusterSettings")]
        private InputList<Inputs.NameValuePairArgs>? _clusterSettings;

        /// <summary>
        /// Custom settings for changing the behavior of the hosting environment
        /// </summary>
        public InputList<Inputs.NameValuePairArgs> ClusterSettings
        {
            get => _clusterSettings ?? (_clusterSettings = new InputList<Inputs.NameValuePairArgs>());
            set => _clusterSettings = value;
        }

        /// <summary>
        /// Edition of the metadata database for the hostingEnvironment (App Service Environment) e.g. "Standard"
        /// </summary>
        [Input("databaseEdition")]
        public Input<string>? DatabaseEdition { get; set; }

        /// <summary>
        /// Service objective of the metadata database for the hostingEnvironment (App Service Environment) e.g. "S0"
        /// </summary>
        [Input("databaseServiceObjective")]
        public Input<string>? DatabaseServiceObjective { get; set; }

        /// <summary>
        /// DNS suffix of the hostingEnvironment (App Service Environment)
        /// </summary>
        [Input("dnsSuffix")]
        public Input<string>? DnsSuffix { get; set; }

        [Input("environmentCapacities")]
        private InputList<Inputs.StampCapacityArgs>? _environmentCapacities;

        /// <summary>
        /// Current total, used, and available worker capacities
        /// </summary>
        public InputList<Inputs.StampCapacityArgs> EnvironmentCapacities
        {
            get => _environmentCapacities ?? (_environmentCapacities = new InputList<Inputs.StampCapacityArgs>());
            set => _environmentCapacities = value;
        }

        /// <summary>
        /// True/false indicating whether the hostingEnvironment (App Service Environment) is healthy
        /// </summary>
        [Input("environmentIsHealthy")]
        public Input<bool>? EnvironmentIsHealthy { get; set; }

        /// <summary>
        /// Detailed message about with results of the last check of the hostingEnvironment (App Service Environment)
        /// </summary>
        [Input("environmentStatus")]
        public Input<string>? EnvironmentStatus { get; set; }

        /// <summary>
        /// Resource Id
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Specifies which endpoints to serve internally in the hostingEnvironment's (App Service Environment) VNET
        /// </summary>
        [Input("internalLoadBalancingMode")]
        public Input<Pulumi.AzureNextGen.Web.V20150801.InternalLoadBalancingMode>? InternalLoadBalancingMode { get; set; }

        /// <summary>
        /// Number of IP SSL addresses reserved for this hostingEnvironment (App Service Environment)
        /// </summary>
        [Input("ipsslAddressCount")]
        public Input<int>? IpsslAddressCount { get; set; }

        /// <summary>
        /// Kind of resource
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Last deployment action on this hostingEnvironment (App Service Environment)
        /// </summary>
        [Input("lastAction")]
        public Input<string>? LastAction { get; set; }

        /// <summary>
        /// Result of the last deployment action on this hostingEnvironment (App Service Environment)
        /// </summary>
        [Input("lastActionResult")]
        public Input<string>? LastActionResult { get; set; }

        /// <summary>
        /// Resource Location
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Maximum number of VMs in this hostingEnvironment (App Service Environment)
        /// </summary>
        [Input("maximumNumberOfMachines")]
        public Input<int>? MaximumNumberOfMachines { get; set; }

        /// <summary>
        /// Number of front-end instances
        /// </summary>
        [Input("multiRoleCount")]
        public Input<int>? MultiRoleCount { get; set; }

        /// <summary>
        /// Front-end VM size, e.g. "Medium", "Large"
        /// </summary>
        [Input("multiSize")]
        public Input<string>? MultiSize { get; set; }

        /// <summary>
        /// Resource Name
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("networkAccessControlList")]
        private InputList<Inputs.NetworkAccessControlEntryArgs>? _networkAccessControlList;

        /// <summary>
        /// Access control list for controlling traffic to the hostingEnvironment (App Service Environment)
        /// </summary>
        public InputList<Inputs.NetworkAccessControlEntryArgs> NetworkAccessControlList
        {
            get => _networkAccessControlList ?? (_networkAccessControlList = new InputList<Inputs.NetworkAccessControlEntryArgs>());
            set => _networkAccessControlList = value;
        }

        /// <summary>
        /// Provisioning state of the hostingEnvironment (App Service Environment)
        /// </summary>
        [Input("provisioningState")]
        public Input<Pulumi.AzureNextGen.Web.V20150801.ProvisioningState>? ProvisioningState { get; set; }

        /// <summary>
        /// Resource group of the hostingEnvironment (App Service Environment)
        /// </summary>
        [Input("resourceGroup")]
        public Input<string>? ResourceGroup { get; set; }

        /// <summary>
        /// Name of resource group
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Current status of the hostingEnvironment (App Service Environment)
        /// </summary>
        [Input("status", required: true)]
        public Input<Pulumi.AzureNextGen.Web.V20150801.HostingEnvironmentStatus> Status { get; set; } = null!;

        /// <summary>
        /// Subscription of the hostingEnvironment (App Service Environment)
        /// </summary>
        [Input("subscriptionId")]
        public Input<string>? SubscriptionId { get; set; }

        /// <summary>
        /// True/false indicating whether the hostingEnvironment is suspended. The environment can be suspended e.g. when the management endpoint is no longer available
        ///             (most likely because NSG blocked the incoming traffic)
        /// </summary>
        [Input("suspended")]
        public Input<bool>? Suspended { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Resource type
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Number of upgrade domains of this hostingEnvironment (App Service Environment)
        /// </summary>
        [Input("upgradeDomains")]
        public Input<int>? UpgradeDomains { get; set; }

        [Input("vipMappings")]
        private InputList<Inputs.VirtualIPMappingArgs>? _vipMappings;

        /// <summary>
        /// Description of IP SSL mapping for this hostingEnvironment (App Service Environment)
        /// </summary>
        public InputList<Inputs.VirtualIPMappingArgs> VipMappings
        {
            get => _vipMappings ?? (_vipMappings = new InputList<Inputs.VirtualIPMappingArgs>());
            set => _vipMappings = value;
        }

        /// <summary>
        /// Description of the hostingEnvironment's (App Service Environment) virtual network
        /// </summary>
        [Input("virtualNetwork")]
        public Input<Inputs.VirtualNetworkProfileArgs>? VirtualNetwork { get; set; }

        /// <summary>
        /// Name of the hostingEnvironment's (App Service Environment) virtual network
        /// </summary>
        [Input("vnetName")]
        public Input<string>? VnetName { get; set; }

        /// <summary>
        /// Resource group of the hostingEnvironment's (App Service Environment) virtual network
        /// </summary>
        [Input("vnetResourceGroupName")]
        public Input<string>? VnetResourceGroupName { get; set; }

        /// <summary>
        /// Subnet of the hostingEnvironment's (App Service Environment) virtual network
        /// </summary>
        [Input("vnetSubnetName")]
        public Input<string>? VnetSubnetName { get; set; }

        [Input("workerPools")]
        private InputList<Inputs.WorkerPoolArgs>? _workerPools;

        /// <summary>
        /// Description of worker pools with worker size ids, VM sizes, and number of workers in each pool
        /// </summary>
        public InputList<Inputs.WorkerPoolArgs> WorkerPools
        {
            get => _workerPools ?? (_workerPools = new InputList<Inputs.WorkerPoolArgs>());
            set => _workerPools = value;
        }

        public ManagedHostingEnvironmentArgs()
        {
        }
    }
}
