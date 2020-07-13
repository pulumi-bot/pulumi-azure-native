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
    /// A share subscription data transfer object.
    /// </summary>
    public partial class AccountShareSubscription : Pulumi.CustomResource
    {
        /// <summary>
        /// Name of the azure resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Properties on the share subscription
        /// </summary>
        [Output("properties")]
        public Output<Outputs.ShareSubscriptionPropertiesResponse> Properties { get; private set; } = null!;

        /// <summary>
        /// Type of the azure resource
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a AccountShareSubscription resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AccountShareSubscription(string name, AccountShareSubscriptionArgs args, CustomResourceOptions? options = null)
            : base("azurerm:datashare:AccountShareSubscription", name, args ?? new AccountShareSubscriptionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AccountShareSubscription(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:datashare:AccountShareSubscription", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing AccountShareSubscription resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AccountShareSubscription Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AccountShareSubscription(name, id, options);
        }
    }

    public sealed class AccountShareSubscriptionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the share account.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the shareSubscription.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Properties on the share subscription
        /// </summary>
        [Input("properties", required: true)]
        public Input<Inputs.ShareSubscriptionPropertiesArgs> Properties { get; set; } = null!;

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public AccountShareSubscriptionArgs()
        {
        }
    }
}
