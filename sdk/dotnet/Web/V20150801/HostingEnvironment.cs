// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Web.V20150801
{
    /// <summary>
    /// Description of an hostingEnvironment (App Service Environment)
    /// </summary>
    public partial class HostingEnvironment : Pulumi.CustomResource
    {
        /// <summary>
        /// Kind of resource
        /// </summary>
        [Output("kind")]
        public Output<string?> Kind { get; private set; } = null!;

        /// <summary>
        /// Resource Location
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Resource Name
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        [Output("properties")]
        public Output<Outputs.HostingEnvironmentResponsePropertiesResult> Properties { get; private set; } = null!;

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
        /// Create a HostingEnvironment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HostingEnvironment(string name, HostingEnvironmentArgs args, CustomResourceOptions? options = null)
            : base("azurerm:web/v20150801:HostingEnvironment", name, args ?? new HostingEnvironmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private HostingEnvironment(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:web/v20150801:HostingEnvironment", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing HostingEnvironment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HostingEnvironment Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new HostingEnvironment(name, id, options);
        }
    }

    public sealed class HostingEnvironmentArgs : Pulumi.ResourceArgs
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
        public Input<string>? InternalLoadBalancingMode { get; set; }

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
        public Input<string>? ProvisioningState { get; set; }

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
        public Input<string> Status { get; set; } = null!;

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

        public HostingEnvironmentArgs()
        {
        }
    }
}
