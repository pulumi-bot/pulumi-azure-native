// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Authorization.V20190101
{
    /// <summary>
    /// The policy assignment.
    /// </summary>
    public partial class PolicyAssignment : Pulumi.CustomResource
    {
        /// <summary>
        /// The managed identity associated with the policy assignment.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.IdentityResponseResult?> Identity { get; private set; } = null!;

        /// <summary>
        /// The location of the policy assignment. Only required when utilizing managed identity.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the policy assignment.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Properties for the policy assignment.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.PolicyAssignmentPropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// The policy sku. This property is optional, obsolete, and will be ignored.
        /// </summary>
        [Output("sku")]
        public Output<Outputs.PolicySkuResponseResult?> Sku { get; private set; } = null!;

        /// <summary>
        /// The type of the policy assignment.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a PolicyAssignment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PolicyAssignment(string name, PolicyAssignmentArgs args, CustomResourceOptions? options = null)
            : base("azurerm:authorization/v20190101:PolicyAssignment", name, args ?? new PolicyAssignmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PolicyAssignment(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:authorization/v20190101:PolicyAssignment", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing PolicyAssignment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PolicyAssignment Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PolicyAssignment(name, id, options);
        }
    }

    public sealed class PolicyAssignmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// This message will be part of response in case of policy violation.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The display name of the policy assignment.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The managed identity associated with the policy assignment.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.IdentityArgs>? Identity { get; set; }

        /// <summary>
        /// The location of the policy assignment. Only required when utilizing managed identity.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("metadata")]
        private InputMap<object>? _metadata;

        /// <summary>
        /// The policy assignment metadata.
        /// </summary>
        public InputMap<object> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<object>());
            set => _metadata = value;
        }

        /// <summary>
        /// The name of the policy assignment.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("notScopes")]
        private InputList<string>? _notScopes;

        /// <summary>
        /// The policy's excluded scopes.
        /// </summary>
        public InputList<string> NotScopes
        {
            get => _notScopes ?? (_notScopes = new InputList<string>());
            set => _notScopes = value;
        }

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

        /// <summary>
        /// The ID of the policy definition or policy set definition being assigned.
        /// </summary>
        [Input("policyDefinitionId")]
        public Input<string>? PolicyDefinitionId { get; set; }

        /// <summary>
        /// The scope for the policy assignment.
        /// </summary>
        [Input("scope", required: true)]
        public Input<string> Scope { get; set; } = null!;

        /// <summary>
        /// The policy sku. This property is optional, obsolete, and will be ignored.
        /// </summary>
        [Input("sku")]
        public Input<Inputs.PolicySkuArgs>? Sku { get; set; }

        public PolicyAssignmentArgs()
        {
        }
    }
}
