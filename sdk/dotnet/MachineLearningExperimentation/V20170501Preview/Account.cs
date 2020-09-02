// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.MachineLearningExperimentation.V20170501Preview
{
    /// <summary>
    /// An object that represents a machine learning team account.
    /// </summary>
    public partial class Account : Pulumi.CustomResource
    {
        /// <summary>
        /// The immutable id associated with this team account.
        /// </summary>
        [Output("accountId")]
        public Output<string> AccountId { get; private set; } = null!;

        /// <summary>
        /// The creation date of the machine learning team account in ISO8601 format.
        /// </summary>
        [Output("creationDate")]
        public Output<string> CreationDate { get; private set; } = null!;

        /// <summary>
        /// The description of this workspace.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The uri for this machine learning team account.
        /// </summary>
        [Output("discoveryUri")]
        public Output<string> DiscoveryUri { get; private set; } = null!;

        /// <summary>
        /// The friendly name for this workspace. This will be the workspace name in the arm id when the workspace object gets created
        /// </summary>
        [Output("friendlyName")]
        public Output<string?> FriendlyName { get; private set; } = null!;

        /// <summary>
        /// The fully qualified arm id of the user key vault.
        /// </summary>
        [Output("keyVaultId")]
        public Output<string> KeyVaultId { get; private set; } = null!;

        /// <summary>
        /// The location of the resource. This cannot be changed after the resource is created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The current deployment state of team account resource. The provisioningState is to indicate states for resource provisioning.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The no of users/seats who can access this team account. This property defines the charge on the team account.
        /// </summary>
        [Output("seats")]
        public Output<string?> Seats { get; private set; } = null!;

        /// <summary>
        /// The properties of the storage account for the machine learning team account.
        /// </summary>
        [Output("storageAccount")]
        public Output<Outputs.StorageAccountPropertiesResponseResult> StorageAccount { get; private set; } = null!;

        /// <summary>
        /// The tags of the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of the resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The fully qualified arm id of the vso account to be used for this team account.
        /// </summary>
        [Output("vsoAccountId")]
        public Output<string> VsoAccountId { get; private set; } = null!;


        /// <summary>
        /// Create a Account resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Account(string name, AccountArgs args, CustomResourceOptions? options = null)
            : base("azurerm:machinelearningexperimentation/v20170501preview:Account", name, args ?? new AccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Account(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:machinelearningexperimentation/v20170501preview:Account", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Account resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Account Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Account(name, id, options);
        }
    }

    public sealed class AccountArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the machine learning team account.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The description of this workspace.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The friendly name for this workspace. This will be the workspace name in the arm id when the workspace object gets created
        /// </summary>
        [Input("friendlyName")]
        public Input<string>? FriendlyName { get; set; }

        /// <summary>
        /// The fully qualified arm id of the user key vault.
        /// </summary>
        [Input("keyVaultId", required: true)]
        public Input<string> KeyVaultId { get; set; } = null!;

        /// <summary>
        /// The location of the resource. This cannot be changed after the resource is created.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The name of the resource group to which the machine learning team account belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The no of users/seats who can access this team account. This property defines the charge on the team account.
        /// </summary>
        [Input("seats")]
        public Input<string>? Seats { get; set; }

        /// <summary>
        /// The properties of the storage account for the machine learning team account.
        /// </summary>
        [Input("storageAccount", required: true)]
        public Input<Inputs.StorageAccountPropertiesArgs> StorageAccount { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The tags of the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The fully qualified arm id of the vso account to be used for this team account.
        /// </summary>
        [Input("vsoAccountId", required: true)]
        public Input<string> VsoAccountId { get; set; } = null!;

        public AccountArgs()
        {
        }
    }
}
