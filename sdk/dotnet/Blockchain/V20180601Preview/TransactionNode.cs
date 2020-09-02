// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Blockchain.V20180601Preview
{
    /// <summary>
    /// Payload of the transaction node which is the request/response of the resource provider.
    /// </summary>
    public partial class TransactionNode : Pulumi.CustomResource
    {
        /// <summary>
        /// Gets or sets the transaction node dns endpoint.
        /// </summary>
        [Output("dns")]
        public Output<string> Dns { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the firewall rules.
        /// </summary>
        [Output("firewallRules")]
        public Output<ImmutableArray<Outputs.FirewallRuleResponseResult>> FirewallRules { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the transaction node location.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Sets the transaction node dns endpoint basic auth password.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the blockchain member provision state.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the transaction node public key.
        /// </summary>
        [Output("publicKey")]
        public Output<string> PublicKey { get; private set; } = null!;

        /// <summary>
        /// The type of the service - e.g. "Microsoft.Blockchain"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the transaction node dns endpoint basic auth user name.
        /// </summary>
        [Output("userName")]
        public Output<string> UserName { get; private set; } = null!;


        /// <summary>
        /// Create a TransactionNode resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TransactionNode(string name, TransactionNodeArgs args, CustomResourceOptions? options = null)
            : base("azurerm:blockchain/v20180601preview:TransactionNode", name, args ?? new TransactionNodeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TransactionNode(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:blockchain/v20180601preview:TransactionNode", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing TransactionNode resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TransactionNode Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new TransactionNode(name, id, options);
        }
    }

    public sealed class TransactionNodeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Blockchain member name.
        /// </summary>
        [Input("blockchainMemberName", required: true)]
        public Input<string> BlockchainMemberName { get; set; } = null!;

        [Input("firewallRules")]
        private InputList<Inputs.FirewallRuleArgs>? _firewallRules;

        /// <summary>
        /// Gets or sets the firewall rules.
        /// </summary>
        public InputList<Inputs.FirewallRuleArgs> FirewallRules
        {
            get => _firewallRules ?? (_firewallRules = new InputList<Inputs.FirewallRuleArgs>());
            set => _firewallRules = value;
        }

        /// <summary>
        /// Gets or sets the transaction node location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Sets the transaction node dns endpoint basic auth password.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Transaction node name.
        /// </summary>
        [Input("transactionNodeName", required: true)]
        public Input<string> TransactionNodeName { get; set; } = null!;

        public TransactionNodeArgs()
        {
        }
    }
}
