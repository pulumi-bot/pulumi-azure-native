// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.ContainerService.V20200301
{
    /// <summary>
    /// Managed cluster.
    /// </summary>
    [AzureNextGenResourceType("azure-nextgen:containerservice/v20200301:ManagedCluster")]
    public partial class ManagedCluster : Pulumi.CustomResource
    {
        /// <summary>
        /// Profile of Azure Active Directory configuration.
        /// </summary>
        [Output("aadProfile")]
        public Output<Outputs.ManagedClusterAADProfileResponse?> AadProfile { get; private set; } = null!;

        /// <summary>
        /// Profile of managed cluster add-on.
        /// </summary>
        [Output("addonProfiles")]
        public Output<ImmutableDictionary<string, Outputs.ManagedClusterAddonProfileResponse>?> AddonProfiles { get; private set; } = null!;

        /// <summary>
        /// Properties of the agent pool.
        /// </summary>
        [Output("agentPoolProfiles")]
        public Output<ImmutableArray<Outputs.ManagedClusterAgentPoolProfileResponse>> AgentPoolProfiles { get; private set; } = null!;

        /// <summary>
        /// Access profile for managed cluster API server.
        /// </summary>
        [Output("apiServerAccessProfile")]
        public Output<Outputs.ManagedClusterAPIServerAccessProfileResponse?> ApiServerAccessProfile { get; private set; } = null!;

        /// <summary>
        /// Parameters to be applied to the cluster-autoscaler when enabled
        /// </summary>
        [Output("autoScalerProfile")]
        public Output<Outputs.ManagedClusterPropertiesResponseAutoScalerProfile?> AutoScalerProfile { get; private set; } = null!;

        /// <summary>
        /// ResourceId of the disk encryption set to use for enabling encryption at rest.
        /// </summary>
        [Output("diskEncryptionSetID")]
        public Output<string?> DiskEncryptionSetID { get; private set; } = null!;

        /// <summary>
        /// DNS prefix specified when creating the managed cluster.
        /// </summary>
        [Output("dnsPrefix")]
        public Output<string?> DnsPrefix { get; private set; } = null!;

        /// <summary>
        /// (DEPRECATING) Whether to enable Kubernetes pod security policy (preview). This feature is set for removal on October 15th, 2020. Learn more at aka.ms/aks/azpodpolicy.
        /// </summary>
        [Output("enablePodSecurityPolicy")]
        public Output<bool?> EnablePodSecurityPolicy { get; private set; } = null!;

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
        /// The identity of the managed cluster, if configured.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.ManagedClusterIdentityResponse?> Identity { get; private set; } = null!;

        /// <summary>
        /// Identities associated with the cluster.
        /// </summary>
        [Output("identityProfile")]
        public Output<ImmutableDictionary<string, Outputs.ManagedClusterPropertiesResponseIdentityProfile>?> IdentityProfile { get; private set; } = null!;

        /// <summary>
        /// Version of Kubernetes specified when creating the managed cluster.
        /// </summary>
        [Output("kubernetesVersion")]
        public Output<string?> KubernetesVersion { get; private set; } = null!;

        /// <summary>
        /// Profile for Linux VMs in the container service cluster.
        /// </summary>
        [Output("linuxProfile")]
        public Output<Outputs.ContainerServiceLinuxProfileResponse?> LinuxProfile { get; private set; } = null!;

        /// <summary>
        /// Resource location
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The max number of agent pools for the managed cluster.
        /// </summary>
        [Output("maxAgentPools")]
        public Output<int> MaxAgentPools { get; private set; } = null!;

        /// <summary>
        /// Resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Profile of network configuration.
        /// </summary>
        [Output("networkProfile")]
        public Output<Outputs.ContainerServiceNetworkProfileResponse?> NetworkProfile { get; private set; } = null!;

        /// <summary>
        /// Name of the resource group containing agent pool nodes.
        /// </summary>
        [Output("nodeResourceGroup")]
        public Output<string?> NodeResourceGroup { get; private set; } = null!;

        /// <summary>
        /// FQDN of private cluster.
        /// </summary>
        [Output("privateFQDN")]
        public Output<string> PrivateFQDN { get; private set; } = null!;

        /// <summary>
        /// The current deployment or provisioning state, which only appears in the response.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Information about a service principal identity for the cluster to use for manipulating Azure APIs.
        /// </summary>
        [Output("servicePrincipalProfile")]
        public Output<Outputs.ManagedClusterServicePrincipalProfileResponse?> ServicePrincipalProfile { get; private set; } = null!;

        /// <summary>
        /// The managed cluster SKU.
        /// </summary>
        [Output("sku")]
        public Output<Outputs.ManagedClusterSKUResponse?> Sku { get; private set; } = null!;

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
        /// Profile for Windows VMs in the container service cluster.
        /// </summary>
        [Output("windowsProfile")]
        public Output<Outputs.ManagedClusterWindowsProfileResponse?> WindowsProfile { get; private set; } = null!;


        /// <summary>
        /// Create a ManagedCluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ManagedCluster(string name, ManagedClusterArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:containerservice/v20200301:ManagedCluster", name, args ?? new ManagedClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ManagedCluster(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:containerservice/v20200301:ManagedCluster", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:containerservice/latest:ManagedCluster"},
                    new Pulumi.Alias { Type = "azure-nextgen:containerservice/v20170831:ManagedCluster"},
                    new Pulumi.Alias { Type = "azure-nextgen:containerservice/v20180331:ManagedCluster"},
                    new Pulumi.Alias { Type = "azure-nextgen:containerservice/v20180801preview:ManagedCluster"},
                    new Pulumi.Alias { Type = "azure-nextgen:containerservice/v20190201:ManagedCluster"},
                    new Pulumi.Alias { Type = "azure-nextgen:containerservice/v20190401:ManagedCluster"},
                    new Pulumi.Alias { Type = "azure-nextgen:containerservice/v20190601:ManagedCluster"},
                    new Pulumi.Alias { Type = "azure-nextgen:containerservice/v20190801:ManagedCluster"},
                    new Pulumi.Alias { Type = "azure-nextgen:containerservice/v20191001:ManagedCluster"},
                    new Pulumi.Alias { Type = "azure-nextgen:containerservice/v20191101:ManagedCluster"},
                    new Pulumi.Alias { Type = "azure-nextgen:containerservice/v20200101:ManagedCluster"},
                    new Pulumi.Alias { Type = "azure-nextgen:containerservice/v20200201:ManagedCluster"},
                    new Pulumi.Alias { Type = "azure-nextgen:containerservice/v20200401:ManagedCluster"},
                    new Pulumi.Alias { Type = "azure-nextgen:containerservice/v20200601:ManagedCluster"},
                    new Pulumi.Alias { Type = "azure-nextgen:containerservice/v20200701:ManagedCluster"},
                    new Pulumi.Alias { Type = "azure-nextgen:containerservice/v20200901:ManagedCluster"},
                    new Pulumi.Alias { Type = "azure-nextgen:containerservice/v20201101:ManagedCluster"},
                    new Pulumi.Alias { Type = "azure-nextgen:containerservice/v20201201:ManagedCluster"},
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
        /// Properties of the agent pool.
        /// </summary>
        public InputList<Inputs.ManagedClusterAgentPoolProfileArgs> AgentPoolProfiles
        {
            get => _agentPoolProfiles ?? (_agentPoolProfiles = new InputList<Inputs.ManagedClusterAgentPoolProfileArgs>());
            set => _agentPoolProfiles = value;
        }

        /// <summary>
        /// Access profile for managed cluster API server.
        /// </summary>
        [Input("apiServerAccessProfile")]
        public Input<Inputs.ManagedClusterAPIServerAccessProfileArgs>? ApiServerAccessProfile { get; set; }

        /// <summary>
        /// Parameters to be applied to the cluster-autoscaler when enabled
        /// </summary>
        [Input("autoScalerProfile")]
        public Input<Inputs.ManagedClusterPropertiesAutoScalerProfileArgs>? AutoScalerProfile { get; set; }

        /// <summary>
        /// ResourceId of the disk encryption set to use for enabling encryption at rest.
        /// </summary>
        [Input("diskEncryptionSetID")]
        public Input<string>? DiskEncryptionSetID { get; set; }

        /// <summary>
        /// DNS prefix specified when creating the managed cluster.
        /// </summary>
        [Input("dnsPrefix")]
        public Input<string>? DnsPrefix { get; set; }

        /// <summary>
        /// (DEPRECATING) Whether to enable Kubernetes pod security policy (preview). This feature is set for removal on October 15th, 2020. Learn more at aka.ms/aks/azpodpolicy.
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

        [Input("identityProfile")]
        private InputMap<Inputs.ManagedClusterPropertiesIdentityProfileArgs>? _identityProfile;

        /// <summary>
        /// Identities associated with the cluster.
        /// </summary>
        public InputMap<Inputs.ManagedClusterPropertiesIdentityProfileArgs> IdentityProfile
        {
            get => _identityProfile ?? (_identityProfile = new InputMap<Inputs.ManagedClusterPropertiesIdentityProfileArgs>());
            set => _identityProfile = value;
        }

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
        [Input("location")]
        public Input<string>? Location { get; set; }

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
        /// The name of the managed cluster resource.
        /// </summary>
        [Input("resourceName", required: true)]
        public Input<string> ResourceName { get; set; } = null!;

        /// <summary>
        /// Information about a service principal identity for the cluster to use for manipulating Azure APIs.
        /// </summary>
        [Input("servicePrincipalProfile")]
        public Input<Inputs.ManagedClusterServicePrincipalProfileArgs>? ServicePrincipalProfile { get; set; }

        /// <summary>
        /// The managed cluster SKU.
        /// </summary>
        [Input("sku")]
        public Input<Inputs.ManagedClusterSKUArgs>? Sku { get; set; }

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
