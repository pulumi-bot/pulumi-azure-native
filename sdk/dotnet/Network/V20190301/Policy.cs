// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20190301
{
    /// <summary>
    /// Defines web application firewall policy.
    /// 
    /// ## Example Usage
    /// ### Creates specific policy
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var policy = new AzureRM.Network.V20190301.Policy("policy", new AzureRM.Network.V20190301.PolicyArgs
    ///         {
    ///             CustomRules = new AzureRM.Network.V20190301.Inputs.CustomRuleListArgs
    ///             {
    ///                 Rules = 
    ///                 {
    ///                     new AzureRM.Network.V20190301.Inputs.CustomRuleArgs
    ///                     {
    ///                         Action = "Block",
    ///                         MatchConditions = 
    ///                         {
    ///                             new AzureRM.Network.V20190301.Inputs.MatchConditionArgs
    ///                             {
    ///                                 MatchValue = 
    ///                                 {
    ///                                     "192.168.1.0/24",
    ///                                     "10.0.0.0/24",
    ///                                 },
    ///                                 MatchVariable = "RemoteAddr",
    ///                                 Operator = "IPMatch",
    ///                             },
    ///                         },
    ///                         Name = "Rule1",
    ///                         Priority = 1,
    ///                         RateLimitThreshold = 1000,
    ///                         RuleType = "RateLimitRule",
    ///                     },
    ///                     new AzureRM.Network.V20190301.Inputs.CustomRuleArgs
    ///                     {
    ///                         Action = "Block",
    ///                         MatchConditions = 
    ///                         {
    ///                             new AzureRM.Network.V20190301.Inputs.MatchConditionArgs
    ///                             {
    ///                                 MatchValue = 
    ///                                 {
    ///                                     "CH",
    ///                                 },
    ///                                 MatchVariable = "RemoteAddr",
    ///                                 Operator = "GeoMatch",
    ///                             },
    ///                             new AzureRM.Network.V20190301.Inputs.MatchConditionArgs
    ///                             {
    ///                                 MatchValue = 
    ///                                 {
    ///                                     "windows",
    ///                                 },
    ///                                 MatchVariable = "RequestHeader",
    ///                                 Operator = "Contains",
    ///                                 Selector = "UserAgent",
    ///                                 Transforms = 
    ///                                 {
    ///                                     "Lowercase",
    ///                                 },
    ///                             },
    ///                         },
    ///                         Name = "Rule2",
    ///                         Priority = 2,
    ///                         RuleType = "MatchRule",
    ///                     },
    ///                 },
    ///             },
    ///             ManagedRules = new AzureRM.Network.V20190301.Inputs.ManagedRuleSetListArgs
    ///             {
    ///                 ManagedRuleSets = 
    ///                 {
    ///                     new AzureRM.Network.V20190301.Inputs.ManagedRuleSetArgs
    ///                     {
    ///                         RuleGroupOverrides = 
    ///                         {
    ///                             new AzureRM.Network.V20190301.Inputs.ManagedRuleGroupOverrideArgs
    ///                             {
    ///                                 RuleGroupName = "Group1",
    ///                                 Rules = 
    ///                                 {
    ///                                     new AzureRM.Network.V20190301.Inputs.ManagedRuleOverrideArgs
    ///                                     {
    ///                                         Action = "Redirect",
    ///                                         EnabledState = "Enabled",
    ///                                         RuleId = "GROUP1-0001",
    ///                                     },
    ///                                     new AzureRM.Network.V20190301.Inputs.ManagedRuleOverrideArgs
    ///                                     {
    ///                                         RuleId = "GROUP1-0002",
    ///                                     },
    ///                                 },
    ///                             },
    ///                         },
    ///                         RuleSetType = "DefaultRuleSet",
    ///                         RuleSetVersion = "preview-1.0",
    ///                     },
    ///                 },
    ///             },
    ///             PolicyName = "Policy1",
    ///             PolicySettings = new AzureRM.Network.V20190301.Inputs.PolicySettingsArgs
    ///             {
    ///                 CustomBlockResponseBody = "PGh0bWw+CjxoZWFkZXI+PHRpdGxlPkhlbGxvPC90aXRsZT48L2hlYWRlcj4KPGJvZHk+CkhlbGxvIHdvcmxkCjwvYm9keT4KPC9odG1sPg==",
    ///                 CustomBlockResponseStatusCode = 499,
    ///                 RedirectUrl = "http://www.bing.com",
    ///             },
    ///             ResourceGroupName = "rg1",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class Policy : Pulumi.CustomResource
    {
        /// <summary>
        /// Describes custom rules inside the policy.
        /// </summary>
        [Output("customRules")]
        public Output<Outputs.CustomRuleListResponseResult?> CustomRules { get; private set; } = null!;

        /// <summary>
        /// Gets a unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string?> Etag { get; private set; } = null!;

        /// <summary>
        /// Describes Frontend Endpoints associated with this Web Application Firewall policy.
        /// </summary>
        [Output("frontendEndpointLinks")]
        public Output<ImmutableArray<Outputs.FrontendEndpointLinkResponseResult>> FrontendEndpointLinks { get; private set; } = null!;

        /// <summary>
        /// Resource location.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// Describes managed rules inside the policy.
        /// </summary>
        [Output("managedRules")]
        public Output<Outputs.ManagedRuleSetListResponseResult?> ManagedRules { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Describes settings for the policy.
        /// </summary>
        [Output("policySettings")]
        public Output<Outputs.PolicySettingsResponseResult?> PolicySettings { get; private set; } = null!;

        /// <summary>
        /// Provisioning state of the policy.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        [Output("resourceState")]
        public Output<string> ResourceState { get; private set; } = null!;

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
        /// Create a Policy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Policy(string name, PolicyArgs args, CustomResourceOptions? options = null)
            : base("azurerm:network/v20190301:Policy", name, args ?? new PolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Policy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:network/v20190301:Policy", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:network/latest:Policy"},
                    new Pulumi.Alias { Type = "azurerm:network/v20180801:Policy"},
                    new Pulumi.Alias { Type = "azurerm:network/v20191001:Policy"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200401:Policy"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Policy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Policy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Policy(name, id, options);
        }
    }

    public sealed class PolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Describes custom rules inside the policy.
        /// </summary>
        [Input("customRules")]
        public Input<Inputs.CustomRuleListArgs>? CustomRules { get; set; }

        /// <summary>
        /// Gets a unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Resource location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Describes managed rules inside the policy.
        /// </summary>
        [Input("managedRules")]
        public Input<Inputs.ManagedRuleSetListArgs>? ManagedRules { get; set; }

        /// <summary>
        /// The name of the Web Application Firewall Policy.
        /// </summary>
        [Input("policyName", required: true)]
        public Input<string> PolicyName { get; set; } = null!;

        /// <summary>
        /// Describes settings for the policy.
        /// </summary>
        [Input("policySettings")]
        public Input<Inputs.PolicySettingsArgs>? PolicySettings { get; set; }

        /// <summary>
        /// Name of the Resource group within the Azure subscription.
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

        public PolicyArgs()
        {
        }
    }
}
