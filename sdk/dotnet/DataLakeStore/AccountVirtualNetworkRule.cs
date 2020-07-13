// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataLakeStore
{
    /// <summary>
    /// Data Lake Store virtual network rule information.
    /// </summary>
    public partial class AccountVirtualNetworkRule : Pulumi.CustomResource
    {
        /// <summary>
        /// The resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The virtual network rule properties.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.VirtualNetworkRulePropertiesResponse> Properties { get; private set; } = null!;

        /// <summary>
        /// The resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a AccountVirtualNetworkRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AccountVirtualNetworkRule(string name, AccountVirtualNetworkRuleArgs args, CustomResourceOptions? options = null)
            : base("azurerm:datalakestore:AccountVirtualNetworkRule", name, args ?? new AccountVirtualNetworkRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AccountVirtualNetworkRule(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:datalakestore:AccountVirtualNetworkRule", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing AccountVirtualNetworkRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AccountVirtualNetworkRule Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AccountVirtualNetworkRule(name, id, options);
        }
    }

    public sealed class AccountVirtualNetworkRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Data Lake Store account.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the virtual network rule to create or update.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The virtual network rule properties to use when creating a new virtual network rule.
        /// </summary>
        [Input("properties", required: true)]
        public Input<Inputs.CreateOrUpdateVirtualNetworkRulePropertiesArgs> Properties { get; set; } = null!;

        /// <summary>
        /// The name of the Azure resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public AccountVirtualNetworkRuleArgs()
        {
        }
    }
}
