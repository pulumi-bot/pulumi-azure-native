// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ContainerService.V20180331
{
    /// <summary>
    /// Managed cluster.
    /// 
    /// ## Example Usage
    /// ### Create/Update Managed Cluster
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var managedCluster = new AzureRM.ContainerService.V20180331.ManagedCluster("managedCluster", new AzureRM.ContainerService.V20180331.ManagedClusterArgs
    ///         {
    ///             AddonProfiles = ,
    ///             AgentPoolProfiles = 
    ///             {
    ///                 new AzureRM.ContainerService.V20180331.Inputs.ManagedClusterAgentPoolProfileArgs
    ///                 {
    ///                     Count = 3,
    ///                     Name = "nodepool1",
    ///                     OsType = "Linux",
    ///                     VmSize = "Standard_DS1_v2",
    ///                 },
    ///             },
    ///             DnsPrefix = "dnsprefix1",
    ///             EnableRBAC = false,
    ///             KubernetesVersion = "",
    ///             LinuxProfile = new AzureRM.ContainerService.V20180331.Inputs.ContainerServiceLinuxProfileArgs
    ///             {
    ///                 AdminUsername = "azureuser",
    ///                 Ssh = new AzureRM.ContainerService.V20180331.Inputs.ContainerServiceSshConfigurationArgs
    ///                 {
    ///                     PublicKeys = 
    ///                     {
    ///                         new AzureRM.ContainerService.V20180331.Inputs.ContainerServiceSshPublicKeyArgs
    ///                         {
    ///                             KeyData = "keydata",
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///             Location = "location1",
    ///             ResourceGroupName = "rg1",
    ///             ResourceName = "clustername1",
    ///             ServicePrincipalProfile = new AzureRM.ContainerService.V20180331.Inputs.ManagedClusterServicePrincipalProfileArgs
    ///             {
    ///                 ClientId = "clientid",
    ///                 Secret = "secret",
    ///             },
    ///             Tags = 
    ///             {
    ///                 { "archv2", "" },
    ///                 { "tier", "production" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class ManagedCluster : Pulumi.CustomResource
    {
        /// <summary>
        /// Profile of Azure Active Directory configuration.
        /// </summary>
        [Output("aadProfile")]
        public Output<Outputs.ManagedClusterAADProfileResponseResult?> AadProfile { get; private set; } = null!;

        /// <summary>
        /// Profile of managed cluster add-on.
        /// </summary>
        [Output("addonProfiles")]
        public Output<ImmutableDictionary<string, Outputs.ManagedClusterAddonProfileResponseResult>?> AddonProfiles { get; private set; } = null!;

        /// <summary>
        /// Properties of the agent pool. Currently only one agent pool can exist.
        /// </summary>
        [Output("agentPoolProfiles")]
        public Output<ImmutableArray<Outputs.ManagedClusterAgentPoolProfileResponseResult>> AgentPoolProfiles { get; private set; } = null!;

        /// <summary>
        /// DNS prefix specified when creating the managed cluster.
        /// </summary>
        [Output("dnsPrefix")]
        public Output<string?> DnsPrefix { get; private set; } = null!;

        /// <summary>
        /// Whether to enable Kubernetes Role-Based Access Control.
        /// </summary>
        [Output("enableRBAC")]
        public Output<bool?> EnableRBAC { get; private set; } = null!;

        /// <summary>
        /// FQDN for the master pool.
        /// </summary>
        [Output("fqdn")]
        public Output<string> Fqdn { get; private set; } = null!;

        /// <summary>
        /// Version of Kubernetes specified when creating the managed cluster.
        /// </summary>
        [Output("kubernetesVersion")]
        public Output<string?> KubernetesVersion { get; private set; } = null!;

        /// <summary>
        /// Profile for Linux VMs in the container service cluster.
        /// </summary>
        [Output("linuxProfile")]
        public Output<Outputs.ContainerServiceLinuxProfileResponseResult?> LinuxProfile { get; private set; } = null!;

        /// <summary>
        /// Resource location
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Profile of network configuration.
        /// </summary>
        [Output("networkProfile")]
        public Output<Outputs.ContainerServiceNetworkProfileResponseResult?> NetworkProfile { get; private set; } = null!;

        /// <summary>
        /// Name of the resource group containing agent pool nodes.
        /// </summary>
        [Output("nodeResourceGroup")]
        public Output<string> NodeResourceGroup { get; private set; } = null!;

        /// <summary>
        /// The current deployment or provisioning state, which only appears in the response.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Information about a service principal identity for the cluster to use for manipulating Azure APIs.
        /// </summary>
        [Output("servicePrincipalProfile")]
        public Output<Outputs.ManagedClusterServicePrincipalProfileResponseResult?> ServicePrincipalProfile { get; private set; } = null!;

        /// <summary>
        /// Resource tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Resource type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ManagedCluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ManagedCluster(string name, ManagedClusterArgs args, CustomResourceOptions? options = null)
            : base("azurerm:containerservice/v20180331:ManagedCluster", name, args ?? new ManagedClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ManagedCluster(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:containerservice/v20180331:ManagedCluster", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:containerservice/latest:ManagedCluster"},
                    new Pulumi.Alias { Type = "azurerm:containerservice/v20170831:ManagedCluster"},
                    new Pulumi.Alias { Type = "azurerm:containerservice/v20180801preview:ManagedCluster"},
                    new Pulumi.Alias { Type = "azurerm:containerservice/v20190201:ManagedCluster"},
                    new Pulumi.Alias { Type = "azurerm:containerservice/v20190401:ManagedCluster"},
                    new Pulumi.Alias { Type = "azurerm:containerservice/v20190601:ManagedCluster"},
                    new Pulumi.Alias { Type = "azurerm:containerservice/v20190801:ManagedCluster"},
                    new Pulumi.Alias { Type = "azurerm:containerservice/v20191001:ManagedCluster"},
                    new Pulumi.Alias { Type = "azurerm:containerservice/v20191101:ManagedCluster"},
                    new Pulumi.Alias { Type = "azurerm:containerservice/v20200101:ManagedCluster"},
                    new Pulumi.Alias { Type = "azurerm:containerservice/v20200201:ManagedCluster"},
                    new Pulumi.Alias { Type = "azurerm:containerservice/v20200301:ManagedCluster"},
                    new Pulumi.Alias { Type = "azurerm:containerservice/v20200401:ManagedCluster"},
                    new Pulumi.Alias { Type = "azurerm:containerservice/v20200601:ManagedCluster"},
                    new Pulumi.Alias { Type = "azurerm:containerservice/v20200701:ManagedCluster"},
                },
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
        /// Profile of Azure Active Directory configuration.
        /// </summary>
        [Input("aadProfile")]
        public Input<Inputs.ManagedClusterAADProfileArgs>? AadProfile { get; set; }

        [Input("addonProfiles")]
        private InputMap<Inputs.ManagedClusterAddonProfileArgs>? _addonProfiles;

        /// <summary>
        /// Profile of managed cluster add-on.
        /// </summary>
        public InputMap<Inputs.ManagedClusterAddonProfileArgs> AddonProfiles
        {
            get => _addonProfiles ?? (_addonProfiles = new InputMap<Inputs.ManagedClusterAddonProfileArgs>());
            set => _addonProfiles = value;
        }

        [Input("agentPoolProfiles")]
        private InputList<Inputs.ManagedClusterAgentPoolProfileArgs>? _agentPoolProfiles;

        /// <summary>
        /// Properties of the agent pool. Currently only one agent pool can exist.
        /// </summary>
        public InputList<Inputs.ManagedClusterAgentPoolProfileArgs> AgentPoolProfiles
        {
            get => _agentPoolProfiles ?? (_agentPoolProfiles = new InputList<Inputs.ManagedClusterAgentPoolProfileArgs>());
            set => _agentPoolProfiles = value;
        }

        /// <summary>
        /// DNS prefix specified when creating the managed cluster.
        /// </summary>
        [Input("dnsPrefix")]
        public Input<string>? DnsPrefix { get; set; }

        /// <summary>
        /// Whether to enable Kubernetes Role-Based Access Control.
        /// </summary>
        [Input("enableRBAC")]
        public Input<bool>? EnableRBAC { get; set; }

        /// <summary>
        /// Version of Kubernetes specified when creating the managed cluster.
        /// </summary>
        [Input("kubernetesVersion")]
        public Input<string>? KubernetesVersion { get; set; }

        /// <summary>
        /// Profile for Linux VMs in the container service cluster.
        /// </summary>
        [Input("linuxProfile")]
        public Input<Inputs.ContainerServiceLinuxProfileArgs>? LinuxProfile { get; set; }

        /// <summary>
        /// Resource location
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Profile of network configuration.
        /// </summary>
        [Input("networkProfile")]
        public Input<Inputs.ContainerServiceNetworkProfileArgs>? NetworkProfile { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the managed cluster resource.
        /// </summary>
        [Input("resourceName", required: true)]
        public Input<string> ResourceName { get; set; } = null!;

        /// <summary>
        /// Information about a service principal identity for the cluster to use for manipulating Azure APIs.
        /// </summary>
        [Input("servicePrincipalProfile")]
        public Input<Inputs.ManagedClusterServicePrincipalProfileArgs>? ServicePrincipalProfile { get; set; }

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

        public ManagedClusterArgs()
        {
        }
    }
}
