// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.ServiceBus.V20180101Preview
{
    /// <summary>
    /// Description of NetworkRuleSet resource.
    /// </summary>
    public partial class NamespaceNetworkRuleSet : Pulumi.CustomResource
    {
        /// <summary>
        /// Default Action for Network Rule Set
        /// </summary>
        [Output("defaultAction")]
        public Output<string?> DefaultAction { get; private set; } = null!;

        /// <summary>
        /// List of IpRules
        /// </summary>
        [Output("ipRules")]
        public Output<ImmutableArray<Outputs.NWRuleSetIpRulesResponse>> IpRules { get; private set; } = null!;

        /// <summary>
        /// Resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Resource type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// List VirtualNetwork Rules
        /// </summary>
        [Output("virtualNetworkRules")]
        public Output<ImmutableArray<Outputs.NWRuleSetVirtualNetworkRulesResponse>> VirtualNetworkRules { get; private set; } = null!;


        /// <summary>
        /// Create a NamespaceNetworkRuleSet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NamespaceNetworkRuleSet(string name, NamespaceNetworkRuleSetArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:servicebus/v20180101preview:NamespaceNetworkRuleSet", name, args ?? new NamespaceNetworkRuleSetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NamespaceNetworkRuleSet(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:servicebus/v20180101preview:NamespaceNetworkRuleSet", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:servicebus/latest:NamespaceNetworkRuleSet"},
                    new Pulumi.Alias { Type = "azure-nextgen:servicebus/v20170401:NamespaceNetworkRuleSet"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing NamespaceNetworkRuleSet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NamespaceNetworkRuleSet Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new NamespaceNetworkRuleSet(name, id, options);
        }
    }

    public sealed class NamespaceNetworkRuleSetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Default Action for Network Rule Set
        /// </summary>
        [Input("defaultAction")]
        public InputUnion<string, Pulumi.AzureNextGen.ServiceBus.V20180101Preview.DefaultAction>? DefaultAction { get; set; }

        [Input("ipRules")]
        private InputList<Inputs.NWRuleSetIpRulesArgs>? _ipRules;

        /// <summary>
        /// List of IpRules
        /// </summary>
        public InputList<Inputs.NWRuleSetIpRulesArgs> IpRules
        {
            get => _ipRules ?? (_ipRules = new InputList<Inputs.NWRuleSetIpRulesArgs>());
            set => _ipRules = value;
        }

        /// <summary>
        /// The namespace name
        /// </summary>
        [Input("namespaceName", required: true)]
        public Input<string> NamespaceName { get; set; } = null!;

        /// <summary>
        /// Name of the Resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("virtualNetworkRules")]
        private InputList<Inputs.NWRuleSetVirtualNetworkRulesArgs>? _virtualNetworkRules;

        /// <summary>
        /// List VirtualNetwork Rules
        /// </summary>
        public InputList<Inputs.NWRuleSetVirtualNetworkRulesArgs> VirtualNetworkRules
        {
            get => _virtualNetworkRules ?? (_virtualNetworkRules = new InputList<Inputs.NWRuleSetVirtualNetworkRulesArgs>());
            set => _virtualNetworkRules = value;
        }

        public NamespaceNetworkRuleSetArgs()
        {
        }
    }
}
