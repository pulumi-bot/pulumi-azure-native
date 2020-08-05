// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20191101
{
    /// <summary>
    /// Defines web application firewall policy.
    /// </summary>
    public partial class WebApplicationFirewallPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// Resource location.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Properties of the web application firewall policy.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.WebApplicationFirewallPolicyPropertiesFormatResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a WebApplicationFirewallPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WebApplicationFirewallPolicy(string name, WebApplicationFirewallPolicyArgs args, CustomResourceOptions? options = null)
            : base("azurerm:network/v20191101:WebApplicationFirewallPolicy", name, args ?? new WebApplicationFirewallPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WebApplicationFirewallPolicy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:network/v20191101:WebApplicationFirewallPolicy", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing WebApplicationFirewallPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WebApplicationFirewallPolicy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new WebApplicationFirewallPolicy(name, id, options);
        }
    }

    public sealed class WebApplicationFirewallPolicyArgs : Pulumi.ResourceArgs
    {
        [Input("customRules")]
        private InputList<Inputs.WebApplicationFirewallCustomRuleArgs>? _customRules;

        /// <summary>
        /// The custom rules inside the policy.
        /// </summary>
        public InputList<Inputs.WebApplicationFirewallCustomRuleArgs> CustomRules
        {
            get => _customRules ?? (_customRules = new InputList<Inputs.WebApplicationFirewallCustomRuleArgs>());
            set => _customRules = value;
        }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Resource location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Describes the managedRules structure.
        /// </summary>
        [Input("managedRules", required: true)]
        public Input<Inputs.ManagedRulesDefinitionArgs> ManagedRules { get; set; } = null!;

        /// <summary>
        /// The name of the policy.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The PolicySettings for policy.
        /// </summary>
        [Input("policySettings")]
        public Input<Inputs.PolicySettingsArgs>? PolicySettings { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public WebApplicationFirewallPolicyArgs()
        {
        }
    }
}
