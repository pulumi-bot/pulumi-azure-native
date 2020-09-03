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
    /// Rule Group resource.
    /// 
    /// ## Example Usage
    /// ### Create FirewallPolicyRuleGroup
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var firewallPolicyRuleGroup = new AzureRM.Network.V20191101.FirewallPolicyRuleGroup("firewallPolicyRuleGroup", new AzureRM.Network.V20191101.FirewallPolicyRuleGroupArgs
    ///         {
    ///             FirewallPolicyName = "firewallPolicy",
    ///             Priority = 110,
    ///             ResourceGroupName = "rg1",
    ///             RuleGroupName = "ruleGroup1",
    ///             Rules = 
    ///             {
    ///                 new AzureRM.Network.V20191101.Inputs.FirewallPolicyRuleArgs
    ///                 {
    ///                     Name = "Example-Filter-Rule",
    ///                     RuleType = "FirewallPolicyFilterRule",
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class FirewallPolicyRuleGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// Priority of the Firewall Policy Rule Group resource.
        /// </summary>
        [Output("priority")]
        public Output<int?> Priority { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the firewall policy rule group resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Group of Firewall Policy rules.
        /// </summary>
        [Output("rules")]
        public Output<ImmutableArray<Outputs.FirewallPolicyRuleResponseResult>> Rules { get; private set; } = null!;

        /// <summary>
        /// Rule Group type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a FirewallPolicyRuleGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FirewallPolicyRuleGroup(string name, FirewallPolicyRuleGroupArgs args, CustomResourceOptions? options = null)
            : base("azurerm:network/v20191101:FirewallPolicyRuleGroup", name, args ?? new FirewallPolicyRuleGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FirewallPolicyRuleGroup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:network/v20191101:FirewallPolicyRuleGroup", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:network/latest:FirewallPolicyRuleGroup"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190601:FirewallPolicyRuleGroup"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190701:FirewallPolicyRuleGroup"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190801:FirewallPolicyRuleGroup"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190901:FirewallPolicyRuleGroup"},
                    new Pulumi.Alias { Type = "azurerm:network/v20191201:FirewallPolicyRuleGroup"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200301:FirewallPolicyRuleGroup"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200401:FirewallPolicyRuleGroup"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing FirewallPolicyRuleGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FirewallPolicyRuleGroup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new FirewallPolicyRuleGroup(name, id, options);
        }
    }

    public sealed class FirewallPolicyRuleGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Firewall Policy.
        /// </summary>
        [Input("firewallPolicyName", required: true)]
        public Input<string> FirewallPolicyName { get; set; } = null!;

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Priority of the Firewall Policy Rule Group resource.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the FirewallPolicyRuleGroup.
        /// </summary>
        [Input("ruleGroupName", required: true)]
        public Input<string> RuleGroupName { get; set; } = null!;

        [Input("rules")]
        private InputList<Inputs.FirewallPolicyRuleArgs>? _rules;

        /// <summary>
        /// Group of Firewall Policy rules.
        /// </summary>
        public InputList<Inputs.FirewallPolicyRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.FirewallPolicyRuleArgs>());
            set => _rules = value;
        }

        public FirewallPolicyRuleGroupArgs()
        {
        }
    }
}
