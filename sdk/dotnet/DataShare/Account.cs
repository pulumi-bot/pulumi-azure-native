// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataShare
{
    /// <summary>
    /// An account data transfer object.
    /// </summary>
    public partial class Account : Pulumi.CustomResource
    {
        /// <summary>
        /// Identity Info on the Account
        /// </summary>
        [Output("identity")]
        public Output<Outputs.IdentityResponse> Identity { get; private set; } = null!;

        /// <summary>
        /// Location of the azure resource.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// Name of the azure resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Properties on the account
        /// </summary>
        [Output("properties")]
        public Output<Outputs.AccountPropertiesResponse> Properties { get; private set; } = null!;

        /// <summary>
        /// Tags on the azure resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Type of the azure resource
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Account resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Account(string name, AccountArgs args, CustomResourceOptions? options = null)
            : base("azurerm:datashare:Account", name, args ?? new AccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Account(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:datashare:Account", name, null, MakeResourceOptions(options, id))
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
        /// Identity Info on the Account
        /// </summary>
        [Input("identity", required: true)]
        public Input<Inputs.IdentityArgs> Identity { get; set; } = null!;

        /// <summary>
        /// Location of the azure resource.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the share account.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Properties on the account
        /// </summary>
        [Input("properties")]
        public Input<Inputs.AccountPropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Tags on the azure resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public AccountArgs()
        {
        }
    }
}
