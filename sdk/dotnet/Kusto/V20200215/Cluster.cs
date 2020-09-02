// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Kusto.V20200215
{
    /// <summary>
    /// Class representing a Kusto cluster.
    /// </summary>
    public partial class Cluster : Pulumi.CustomResource
    {
        /// <summary>
        /// The cluster data ingestion URI.
        /// </summary>
        [Output("dataIngestionUri")]
        public Output<string> DataIngestionUri { get; private set; } = null!;

        /// <summary>
        /// A boolean value that indicates if the cluster's disks are encrypted.
        /// </summary>
        [Output("enableDiskEncryption")]
        public Output<bool?> EnableDiskEncryption { get; private set; } = null!;

        /// <summary>
        /// A boolean value that indicates if the purge operations are enabled.
        /// </summary>
        [Output("enablePurge")]
        public Output<bool?> EnablePurge { get; private set; } = null!;

        /// <summary>
        /// A boolean value that indicates if the streaming ingest is enabled.
        /// </summary>
        [Output("enableStreamingIngest")]
        public Output<bool?> EnableStreamingIngest { get; private set; } = null!;

        /// <summary>
        /// The identity of the cluster, if configured.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.IdentityResponseResult?> Identity { get; private set; } = null!;

        /// <summary>
        /// KeyVault properties for the cluster encryption.
        /// </summary>
        [Output("keyVaultProperties")]
        public Output<Outputs.KeyVaultPropertiesResponseResult?> KeyVaultProperties { get; private set; } = null!;

        /// <summary>
        /// List of the cluster's language extensions.
        /// </summary>
        [Output("languageExtensions")]
        public Output<Outputs.LanguageExtensionsListResponseResult?> LanguageExtensions { get; private set; } = null!;

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Optimized auto scale definition.
        /// </summary>
        [Output("optimizedAutoscale")]
        public Output<Outputs.OptimizedAutoscaleResponseResult?> OptimizedAutoscale { get; private set; } = null!;

        /// <summary>
        /// The provisioned state of the resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The SKU of the cluster.
        /// </summary>
        [Output("sku")]
        public Output<Outputs.AzureSkuResponseResult> Sku { get; private set; } = null!;

        /// <summary>
        /// The state of the resource.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The reason for the cluster's current state.
        /// </summary>
        [Output("stateReason")]
        public Output<string> StateReason { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The cluster's external tenants.
        /// </summary>
        [Output("trustedExternalTenants")]
        public Output<ImmutableArray<Outputs.TrustedExternalTenantResponseResult>> TrustedExternalTenants { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The cluster URI.
        /// </summary>
        [Output("uri")]
        public Output<string> Uri { get; private set; } = null!;

        /// <summary>
        /// Virtual network definition.
        /// </summary>
        [Output("virtualNetworkConfiguration")]
        public Output<Outputs.VirtualNetworkConfigurationResponseResult?> VirtualNetworkConfiguration { get; private set; } = null!;

        /// <summary>
        /// The availability zones of the cluster.
        /// </summary>
        [Output("zones")]
        public Output<ImmutableArray<string>> Zones { get; private set; } = null!;


        /// <summary>
        /// Create a Cluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Cluster(string name, ClusterArgs args, CustomResourceOptions? options = null)
            : base("azurerm:kusto/v20200215:Cluster", name, args ?? new ClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Cluster(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:kusto/v20200215:Cluster", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:kusto/latest:Cluster"},
                    new Pulumi.Alias { Type = "azurerm:kusto/v20170907privatepreview:Cluster"},
                    new Pulumi.Alias { Type = "azurerm:kusto/v20180907preview:Cluster"},
                    new Pulumi.Alias { Type = "azurerm:kusto/v20190121:Cluster"},
                    new Pulumi.Alias { Type = "azurerm:kusto/v20190515:Cluster"},
                    new Pulumi.Alias { Type = "azurerm:kusto/v20190907:Cluster"},
                    new Pulumi.Alias { Type = "azurerm:kusto/v20191109:Cluster"},
                    new Pulumi.Alias { Type = "azurerm:kusto/v20200614:Cluster"},
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
        /// <summary>
        /// The name of the Kusto cluster.
        /// </summary>
        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        /// <summary>
        /// A boolean value that indicates if the cluster's disks are encrypted.
        /// </summary>
        [Input("enableDiskEncryption")]
        public Input<bool>? EnableDiskEncryption { get; set; }

        /// <summary>
        /// A boolean value that indicates if the purge operations are enabled.
        /// </summary>
        [Input("enablePurge")]
        public Input<bool>? EnablePurge { get; set; }

        /// <summary>
        /// A boolean value that indicates if the streaming ingest is enabled.
        /// </summary>
        [Input("enableStreamingIngest")]
        public Input<bool>? EnableStreamingIngest { get; set; }

        /// <summary>
        /// The identity of the cluster, if configured.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.IdentityArgs>? Identity { get; set; }

        /// <summary>
        /// KeyVault properties for the cluster encryption.
        /// </summary>
        [Input("keyVaultProperties")]
        public Input<Inputs.KeyVaultPropertiesArgs>? KeyVaultProperties { get; set; }

        /// <summary>
        /// List of the cluster's language extensions.
        /// </summary>
        [Input("languageExtensions")]
        public Input<Inputs.LanguageExtensionsListArgs>? LanguageExtensions { get; set; }

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Optimized auto scale definition.
        /// </summary>
        [Input("optimizedAutoscale")]
        public Input<Inputs.OptimizedAutoscaleArgs>? OptimizedAutoscale { get; set; }

        /// <summary>
        /// The name of the resource group containing the Kusto cluster.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The SKU of the cluster.
        /// </summary>
        [Input("sku", required: true)]
        public Input<Inputs.AzureSkuArgs> Sku { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("trustedExternalTenants")]
        private InputList<Inputs.TrustedExternalTenantArgs>? _trustedExternalTenants;

        /// <summary>
        /// The cluster's external tenants.
        /// </summary>
        public InputList<Inputs.TrustedExternalTenantArgs> TrustedExternalTenants
        {
            get => _trustedExternalTenants ?? (_trustedExternalTenants = new InputList<Inputs.TrustedExternalTenantArgs>());
            set => _trustedExternalTenants = value;
        }

        /// <summary>
        /// Virtual network definition.
        /// </summary>
        [Input("virtualNetworkConfiguration")]
        public Input<Inputs.VirtualNetworkConfigurationArgs>? VirtualNetworkConfiguration { get; set; }

        [Input("zones")]
        private InputList<string>? _zones;

        /// <summary>
        /// The availability zones of the cluster.
        /// </summary>
        public InputList<string> Zones
        {
            get => _zones ?? (_zones = new InputList<string>());
            set => _zones = value;
        }

        public ClusterArgs()
        {
        }
    }
}
