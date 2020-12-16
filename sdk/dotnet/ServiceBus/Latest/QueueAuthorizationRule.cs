// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.ServiceBus.Latest
{
    /// <summary>
    /// Description of a namespace authorization rule.
    /// Latest API Version: 2017-04-01.
    /// </summary>
    public partial class QueueAuthorizationRule : Pulumi.CustomResource
    {
        /// <summary>
        /// Resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The rights associated with the rule.
        /// </summary>
        [Output("rights")]
        public Output<ImmutableArray<string>> Rights { get; private set; } = null!;

        /// <summary>
        /// Resource type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a QueueAuthorizationRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public QueueAuthorizationRule(string name, QueueAuthorizationRuleArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:servicebus/latest:QueueAuthorizationRule", name, args ?? new QueueAuthorizationRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private QueueAuthorizationRule(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:servicebus/latest:QueueAuthorizationRule", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:servicebus/v20140901:QueueAuthorizationRule"},
                    new Pulumi.Alias { Type = "azure-nextgen:servicebus/v20150801:QueueAuthorizationRule"},
                    new Pulumi.Alias { Type = "azure-nextgen:servicebus/v20170401:QueueAuthorizationRule"},
                    new Pulumi.Alias { Type = "azure-nextgen:servicebus/v20180101preview:QueueAuthorizationRule"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing QueueAuthorizationRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static QueueAuthorizationRule Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new QueueAuthorizationRule(name, id, options);
        }
    }

    public sealed class QueueAuthorizationRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The authorization rule name.
        /// </summary>
        [Input("authorizationRuleName", required: true)]
        public Input<string> AuthorizationRuleName { get; set; } = null!;

        /// <summary>
        /// The namespace name
        /// </summary>
        [Input("namespaceName", required: true)]
        public Input<string> NamespaceName { get; set; } = null!;

        /// <summary>
        /// The queue name.
        /// </summary>
        [Input("queueName", required: true)]
        public Input<string> QueueName { get; set; } = null!;

        /// <summary>
        /// Name of the Resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("rights", required: true)]
        private InputList<Pulumi.AzureNextGen.ServiceBus.Latest.AccessRights>? _rights;

        /// <summary>
        /// The rights associated with the rule.
        /// </summary>
        public InputList<Pulumi.AzureNextGen.ServiceBus.Latest.AccessRights> Rights
        {
            get => _rights ?? (_rights = new InputList<Pulumi.AzureNextGen.ServiceBus.Latest.AccessRights>());
            set => _rights = value;
        }

        public QueueAuthorizationRuleArgs()
        {
        }
    }
}
