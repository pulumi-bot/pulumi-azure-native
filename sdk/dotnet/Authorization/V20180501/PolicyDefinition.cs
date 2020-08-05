// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Authorization.V20180501
{
    /// <summary>
    /// The policy definition.
    /// </summary>
    public partial class PolicyDefinition : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the policy definition.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The policy definition properties.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.PolicyDefinitionPropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// The type of the resource (Microsoft.Authorization/policyDefinitions).
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a PolicyDefinition resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PolicyDefinition(string name, PolicyDefinitionArgs args, CustomResourceOptions? options = null)
            : base("azurerm:authorization/v20180501:PolicyDefinition", name, args ?? new PolicyDefinitionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PolicyDefinition(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:authorization/v20180501:PolicyDefinition", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing PolicyDefinition resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PolicyDefinition Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PolicyDefinition(name, id, options);
        }
    }

    public sealed class PolicyDefinitionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The policy definition description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The display name of the policy definition.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("metadata")]
        private InputMap<object>? _metadata;

        /// <summary>
        /// The policy definition metadata.
        /// </summary>
        public InputMap<object> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<object>());
            set => _metadata = value;
        }

        /// <summary>
        /// The policy definition mode. Possible values are NotSpecified, Indexed, and All.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// The name of the policy definition to create.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("parameters")]
        private InputMap<object>? _parameters;

        /// <summary>
        /// Required if a parameter is used in policy rule.
        /// </summary>
        public InputMap<object> Parameters
        {
            get => _parameters ?? (_parameters = new InputMap<object>());
            set => _parameters = value;
        }

        [Input("policyRule")]
        private InputMap<object>? _policyRule;

        /// <summary>
        /// The policy rule.
        /// </summary>
        public InputMap<object> PolicyRule
        {
            get => _policyRule ?? (_policyRule = new InputMap<object>());
            set => _policyRule = value;
        }

        /// <summary>
        /// The type of policy definition. Possible values are NotSpecified, BuiltIn, and Custom.
        /// </summary>
        [Input("policyType")]
        public Input<string>? PolicyType { get; set; }

        public PolicyDefinitionArgs()
        {
        }
    }
}
