// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ContainerService.V20190401
{
    /// <summary>
    /// Managed cluster.
    /// </summary>
    public partial class ManagedCluster : Pulumi.CustomResource
    {
        /// <summary>
        /// The identity of the managed cluster, if configured.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.ManagedClusterIdentityResponseResult?> Identity { get; private set; } = null!;

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
        /// Properties of a managed cluster.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.ManagedClusterPropertiesResponseResult> Properties { get; private set; } = null!;

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
            : base("azurerm:containerservice/v20190401:ManagedCluster", name, args ?? new ManagedClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ManagedCluster(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:containerservice/v20190401:ManagedCluster", name, null, MakeResourceOptions(options, id))
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
        /// Properties of the agent pool.
        /// </summary>
        public InputList<Inputs.ManagedClusterAgentPoolProfileArgs> AgentPoolProfiles
        {
            get => _agentPoolProfiles ?? (_agentPoolProfiles = new InputList<Inputs.ManagedClusterAgentPoolProfileArgs>());
            set => _agentPoolProfiles = value;
        }

        [Input("apiServerAuthorizedIPRanges")]
        private InputList<string>? _apiServerAuthorizedIPRanges;

        /// <summary>
        /// (PREVIEW) Authorized IP Ranges to kubernetes API server.
        /// </summary>
        public InputList<string> ApiServerAuthorizedIPRanges
        {
            get => _apiServerAuthorizedIPRanges ?? (_apiServerAuthorizedIPRanges = new InputList<string>());
            set => _apiServerAuthorizedIPRanges = value;
        }

        /// <summary>
        /// DNS prefix specified when creating the managed cluster.
        /// </summary>
        [Input("dnsPrefix")]
        public Input<string>? DnsPrefix { get; set; }

        /// <summary>
        /// (PREVIEW) Whether to enable Kubernetes Pod security policy.
        /// </summary>
        [Input("enablePodSecurityPolicy")]
        public Input<bool>? EnablePodSecurityPolicy { get; set; }

        /// <summary>
        /// Whether to enable Kubernetes Role-Based Access Control.
        /// </summary>
        [Input("enableRBAC")]
        public Input<bool>? EnableRBAC { get; set; }

        /// <summary>
        /// The identity of the managed cluster, if configured.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.ManagedClusterIdentityArgs>? Identity { get; set; }

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
        /// The name of the managed cluster resource.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Profile of network configuration.
        /// </summary>
        [Input("networkProfile")]
        public Input<Inputs.ContainerServiceNetworkProfileArgs>? NetworkProfile { get; set; }

        /// <summary>
        /// Name of the resource group containing agent pool nodes.
        /// </summary>
        [Input("nodeResourceGroup")]
        public Input<string>? NodeResourceGroup { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

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

        /// <summary>
        /// Profile for Windows VMs in the container service cluster.
        /// </summary>
        [Input("windowsProfile")]
        public Input<Inputs.ManagedClusterWindowsProfileArgs>? WindowsProfile { get; set; }

        public ManagedClusterArgs()
        {
        }
    }
}
