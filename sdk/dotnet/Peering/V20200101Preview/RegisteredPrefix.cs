// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Peering.V20200101Preview
{
    /// <summary>
    /// The customer's prefix that is registered by the peering service provider.
    /// </summary>
    public partial class RegisteredPrefix : Pulumi.CustomResource
    {
        /// <summary>
        /// The error message associated with the validation state, if any.
        /// </summary>
        [Output("errorMessage")]
        public Output<string> ErrorMessage { get; private set; } = null!;

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The peering service prefix key that is to be shared with the customer.
        /// </summary>
        [Output("peeringServicePrefixKey")]
        public Output<string> PeeringServicePrefixKey { get; private set; } = null!;

        /// <summary>
        /// The customer's prefix from which traffic originates.
        /// </summary>
        [Output("prefix")]
        public Output<string?> Prefix { get; private set; } = null!;

        /// <summary>
        /// The prefix validation state.
        /// </summary>
        [Output("prefixValidationState")]
        public Output<string> PrefixValidationState { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The type of the resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a RegisteredPrefix resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RegisteredPrefix(string name, RegisteredPrefixArgs args, CustomResourceOptions? options = null)
            : base("azurerm:peering/v20200101preview:RegisteredPrefix", name, args ?? new RegisteredPrefixArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RegisteredPrefix(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:peering/v20200101preview:RegisteredPrefix", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:peering/latest:RegisteredPrefix"},
                    new Pulumi.Alias { Type = "azurerm:peering/v20200401:RegisteredPrefix"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing RegisteredPrefix resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RegisteredPrefix Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new RegisteredPrefix(name, id, options);
        }
    }

    public sealed class RegisteredPrefixArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the peering.
        /// </summary>
        [Input("peeringName", required: true)]
        public Input<string> PeeringName { get; set; } = null!;

        /// <summary>
        /// The customer's prefix from which traffic originates.
        /// </summary>
        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        /// <summary>
        /// The name of the registered prefix.
        /// </summary>
        [Input("registeredPrefixName", required: true)]
        public Input<string> RegisteredPrefixName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public RegisteredPrefixArgs()
        {
        }
    }
}
