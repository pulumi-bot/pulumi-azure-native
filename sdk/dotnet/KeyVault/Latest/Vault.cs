// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.KeyVault.Latest
{
    /// <summary>
    /// Resource information with extended details.
    /// Latest API Version: 2019-09-01.
    /// </summary>
    [AzureNextGenResourceType("azure-nextgen:keyvault/latest:Vault")]
    public partial class Vault : Pulumi.CustomResource
    {
        /// <summary>
        /// Azure location of the key vault resource.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// Name of the key vault resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Properties of the vault
        /// </summary>
        [Output("properties")]
        public Output<Outputs.VaultPropertiesResponse> Properties { get; private set; } = null!;

        /// <summary>
        /// Tags assigned to the key vault resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Resource type of the key vault resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Vault resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Vault(string name, VaultArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:keyvault/latest:Vault", name, args ?? new VaultArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Vault(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:keyvault/latest:Vault", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:keyvault/v20150601:Vault"},
                    new Pulumi.Alias { Type = "azure-nextgen:keyvault/v20161001:Vault"},
                    new Pulumi.Alias { Type = "azure-nextgen:keyvault/v20180214:Vault"},
                    new Pulumi.Alias { Type = "azure-nextgen:keyvault/v20180214preview:Vault"},
                    new Pulumi.Alias { Type = "azure-nextgen:keyvault/v20190901:Vault"},
                    new Pulumi.Alias { Type = "azure-nextgen:keyvault/v20200401preview:Vault"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Vault resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Vault Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Vault(name, id, options);
        }
    }

    public sealed class VaultArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The supported Azure location where the key vault should be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Properties of the vault
        /// </summary>
        [Input("properties", required: true)]
        public Input<Inputs.VaultPropertiesArgs> Properties { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group to which the server belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The tags that will be assigned to the key vault.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Name of the vault
        /// </summary>
        [Input("vaultName", required: true)]
        public Input<string> VaultName { get; set; } = null!;

        public VaultArgs()
        {
        }
    }
}
