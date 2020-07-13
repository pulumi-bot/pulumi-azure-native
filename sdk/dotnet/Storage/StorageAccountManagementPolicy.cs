// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Storage
{
    /// <summary>
    /// The Get Storage Account ManagementPolicies operation response.
    /// </summary>
    public partial class StorageAccountManagementPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Returns the Storage Account Data Policies Rules.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.ManagementPolicyPropertiesResponse> Properties { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a StorageAccountManagementPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StorageAccountManagementPolicy(string name, StorageAccountManagementPolicyArgs args, CustomResourceOptions? options = null)
            : base("azurerm:storage:StorageAccountManagementPolicy", name, args ?? new StorageAccountManagementPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StorageAccountManagementPolicy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:storage:StorageAccountManagementPolicy", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing StorageAccountManagementPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StorageAccountManagementPolicy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new StorageAccountManagementPolicy(name, id, options);
        }
    }

    public sealed class StorageAccountManagementPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the Storage Account Management Policy. It should always be 'default'
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Returns the Storage Account Data Policies Rules.
        /// </summary>
        [Input("properties")]
        public Input<Inputs.ManagementPolicyPropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// The name of the resource group within the user's subscription. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public StorageAccountManagementPolicyArgs()
        {
        }
    }
}
