// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DevTestLab.V20150521Preview
{
    /// <summary>
    /// A Policy.
    /// </summary>
    public partial class PolicyResource : Pulumi.CustomResource
    {
        /// <summary>
        /// The description of the policy.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The evaluator type of the policy.
        /// </summary>
        [Output("evaluatorType")]
        public Output<string?> EvaluatorType { get; private set; } = null!;

        /// <summary>
        /// The fact data of the policy.
        /// </summary>
        [Output("factData")]
        public Output<string?> FactData { get; private set; } = null!;

        /// <summary>
        /// The fact name of the policy.
        /// </summary>
        [Output("factName")]
        public Output<string?> FactName { get; private set; } = null!;

        /// <summary>
        /// The location of the resource.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// The provisioning status of the resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string?> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The status of the policy.
        /// </summary>
        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;

        /// <summary>
        /// The tags of the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The threshold of the policy.
        /// </summary>
        [Output("threshold")]
        public Output<string?> Threshold { get; private set; } = null!;

        /// <summary>
        /// The type of the resource.
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;


        /// <summary>
        /// Create a PolicyResource resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PolicyResource(string name, PolicyResourceArgs args, CustomResourceOptions? options = null)
            : base("azurerm:devtestlab/v20150521preview:PolicyResource", name, args ?? new PolicyResourceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PolicyResource(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:devtestlab/v20150521preview:PolicyResource", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:devtestlab/latest:PolicyResource"},
                    new Pulumi.Alias { Type = "azurerm:devtestlab/v20160515:PolicyResource"},
                    new Pulumi.Alias { Type = "azurerm:devtestlab/v20180915:PolicyResource"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PolicyResource resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PolicyResource Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PolicyResource(name, id, options);
        }
    }

    public sealed class PolicyResourceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The evaluator type of the policy.
        /// </summary>
        [Input("evaluatorType")]
        public Input<string>? EvaluatorType { get; set; }

        /// <summary>
        /// The fact data of the policy.
        /// </summary>
        [Input("factData")]
        public Input<string>? FactData { get; set; }

        /// <summary>
        /// The fact name of the policy.
        /// </summary>
        [Input("factName")]
        public Input<string>? FactName { get; set; }

        /// <summary>
        /// The identifier of the resource.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The name of the lab.
        /// </summary>
        [Input("labName", required: true)]
        public Input<string> LabName { get; set; } = null!;

        /// <summary>
        /// The location of the resource.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the policy set.
        /// </summary>
        [Input("policySetName", required: true)]
        public Input<string> PolicySetName { get; set; } = null!;

        /// <summary>
        /// The provisioning status of the resource.
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The status of the policy.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The tags of the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The threshold of the policy.
        /// </summary>
        [Input("threshold")]
        public Input<string>? Threshold { get; set; }

        /// <summary>
        /// The type of the resource.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public PolicyResourceArgs()
        {
        }
    }
}
