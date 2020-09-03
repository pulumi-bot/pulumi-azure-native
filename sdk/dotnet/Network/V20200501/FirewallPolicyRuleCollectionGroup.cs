// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20200501
{
    /// <summary>
    /// Rule Collection Group resource.
    /// 
    /// ## Example Usage
    /// ### Create FirewallPolicyRuleCollectionGroup
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var firewallPolicyRuleCollectionGroup = new AzureRM.Network.V20200501.FirewallPolicyRuleCollectionGroup("firewallPolicyRuleCollectionGroup", new AzureRM.Network.V20200501.FirewallPolicyRuleCollectionGroupArgs
    ///         {
    ///             FirewallPolicyName = "firewallPolicy",
    ///             Priority = 110,
    ///             ResourceGroupName = "rg1",
    ///             RuleCollectionGroupName = "ruleCollectionGroup1",
    ///             RuleCollections = 
    ///             {
    ///                 new AzureRM.Network.V20200501.Inputs.FirewallPolicyRuleCollectionArgs
    ///                 {
    ///                     Name = "Example-Filter-Rule-Collection",
    ///                     RuleCollectionType = "FirewallPolicyFilterRuleCollection",
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// ### Create FirewallPolicyRuleCollectionGroup With IpGroups
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var firewallPolicyRuleCollectionGroup = new AzureRM.Network.V20200501.FirewallPolicyRuleCollectionGroup("firewallPolicyRuleCollectionGroup", new AzureRM.Network.V20200501.FirewallPolicyRuleCollectionGroupArgs
    ///         {
    ///             FirewallPolicyName = "firewallPolicy",
    ///             Priority = 110,
    ///             ResourceGroupName = "rg1",
    ///             RuleCollectionGroupName = "ruleCollectionGroup1",
    ///             RuleCollections = 
    ///             {
    ///                 new AzureRM.Network.V20200501.Inputs.FirewallPolicyRuleCollectionArgs
    ///                 {
    ///                     Name = "Example-Filter-Rule-Collection",
    ///                     RuleCollectionType = "FirewallPolicyFilterRuleCollection",
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class FirewallPolicyRuleCollectionGroup : Pulumi.CustomResource
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
        /// Priority of the Firewall Policy Rule Collection Group resource.
        /// </summary>
        [Output("priority")]
        public Output<int?> Priority { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the firewall policy rule collection group resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Group of Firewall Policy rule collections.
        /// </summary>
        [Output("ruleCollections")]
        public Output<ImmutableArray<Outputs.FirewallPolicyRuleCollectionResponseResult>> RuleCollections { get; private set; } = null!;

        /// <summary>
        /// Rule Group type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a FirewallPolicyRuleCollectionGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FirewallPolicyRuleCollectionGroup(string name, FirewallPolicyRuleCollectionGroupArgs args, CustomResourceOptions? options = null)
            : base("azurerm:network/v20200501:FirewallPolicyRuleCollectionGroup", name, args ?? new FirewallPolicyRuleCollectionGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FirewallPolicyRuleCollectionGroup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:network/v20200501:FirewallPolicyRuleCollectionGroup", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:network/latest:FirewallPolicyRuleCollectionGroup"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200601:FirewallPolicyRuleCollectionGroup"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing FirewallPolicyRuleCollectionGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FirewallPolicyRuleCollectionGroup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new FirewallPolicyRuleCollectionGroup(name, id, options);
        }
    }

    public sealed class FirewallPolicyRuleCollectionGroupArgs : Pulumi.ResourceArgs
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
        /// Priority of the Firewall Policy Rule Collection Group resource.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the FirewallPolicyRuleCollectionGroup.
        /// </summary>
        [Input("ruleCollectionGroupName", required: true)]
        public Input<string> RuleCollectionGroupName { get; set; } = null!;

        [Input("ruleCollections")]
        private InputList<Inputs.FirewallPolicyRuleCollectionArgs>? _ruleCollections;

        /// <summary>
        /// Group of Firewall Policy rule collections.
        /// </summary>
        public InputList<Inputs.FirewallPolicyRuleCollectionArgs> RuleCollections
        {
            get => _ruleCollections ?? (_ruleCollections = new InputList<Inputs.FirewallPolicyRuleCollectionArgs>());
            set => _ruleCollections = value;
        }

        public FirewallPolicyRuleCollectionGroupArgs()
        {
        }
    }
}
