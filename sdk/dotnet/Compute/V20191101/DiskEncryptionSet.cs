// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.V20191101
{
    /// <summary>
    /// disk encryption set resource.
    /// 
    /// ## Example Usage
    /// ### Create a disk encryption set.
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var diskEncryptionSet = new AzureRM.Compute.V20191101.DiskEncryptionSet("diskEncryptionSet", new AzureRM.Compute.V20191101.DiskEncryptionSetArgs
    ///         {
    ///             ActiveKey = new AzureRM.Compute.V20191101.Inputs.KeyVaultAndKeyReferenceArgs
    ///             {
    ///                 KeyUrl = "https://myvmvault.vault-int.azure-int.net/keys/{key}",
    ///                 SourceVault = new AzureRM.Compute.V20191101.Inputs.SourceVaultArgs
    ///                 {
    ///                     Id = "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/myVMVault",
    ///                 },
    ///             },
    ///             DiskEncryptionSetName = "myDiskEncryptionSet",
    ///             Identity = new AzureRM.Compute.V20191101.Inputs.EncryptionSetIdentityArgs
    ///             {
    ///                 Type = "SystemAssigned",
    ///             },
    ///             Location = "West US",
    ///             ResourceGroupName = "myResourceGroup",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class DiskEncryptionSet : Pulumi.CustomResource
    {
        /// <summary>
        /// The key vault key which is currently used by this disk encryption set.
        /// </summary>
        [Output("activeKey")]
        public Output<Outputs.KeyVaultAndKeyReferenceResponseResult?> ActiveKey { get; private set; } = null!;

        /// <summary>
        /// The managed identity for the disk encryption set. It should be given permission on the key vault before it can be used to encrypt disks.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.EncryptionSetIdentityResponseResult?> Identity { get; private set; } = null!;

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
        /// A readonly collection of key vault keys previously used by this disk encryption set while a key rotation is in progress. It will be empty if there is no ongoing key rotation.
        /// </summary>
        [Output("previousKeys")]
        public Output<ImmutableArray<Outputs.KeyVaultAndKeyReferenceResponseResult>> PreviousKeys { get; private set; } = null!;

        /// <summary>
        /// The disk encryption set provisioning state.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

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
        /// Create a DiskEncryptionSet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DiskEncryptionSet(string name, DiskEncryptionSetArgs args, CustomResourceOptions? options = null)
            : base("azurerm:compute/v20191101:DiskEncryptionSet", name, args ?? new DiskEncryptionSetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DiskEncryptionSet(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:compute/v20191101:DiskEncryptionSet", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:compute/latest:DiskEncryptionSet"},
                    new Pulumi.Alias { Type = "azurerm:compute/v20190701:DiskEncryptionSet"},
                    new Pulumi.Alias { Type = "azurerm:compute/v20200501:DiskEncryptionSet"},
                    new Pulumi.Alias { Type = "azurerm:compute/v20200630:DiskEncryptionSet"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DiskEncryptionSet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DiskEncryptionSet Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DiskEncryptionSet(name, id, options);
        }
    }

    public sealed class DiskEncryptionSetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The key vault key which is currently used by this disk encryption set.
        /// </summary>
        [Input("activeKey")]
        public Input<Inputs.KeyVaultAndKeyReferenceArgs>? ActiveKey { get; set; }

        /// <summary>
        /// The name of the disk encryption set that is being created. The name can't be changed after the disk encryption set is created. Supported characters for the name are a-z, A-Z, 0-9 and _. The maximum name length is 80 characters.
        /// </summary>
        [Input("diskEncryptionSetName", required: true)]
        public Input<string> DiskEncryptionSetName { get; set; } = null!;

        /// <summary>
        /// The managed identity for the disk encryption set. It should be given permission on the key vault before it can be used to encrypt disks.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.EncryptionSetIdentityArgs>? Identity { get; set; }

        /// <summary>
        /// Resource location
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

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

        public DiskEncryptionSetArgs()
        {
        }
    }
}
