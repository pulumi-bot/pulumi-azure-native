// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Management.V20170601Preview
{
    /// <summary>
    /// The policy set definition.
    /// </summary>
    public partial class PolicySetDefinitionAtManagementGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// The policy set definition description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The display name of the policy set definition.
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The policy set definition metadata.
        /// </summary>
        [Output("metadata")]
        public Output<ImmutableDictionary<string, object>?> Metadata { get; private set; } = null!;

        /// <summary>
        /// The name of the policy set definition.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The policy set definition parameters that can be used in policy definition references.
        /// </summary>
        [Output("parameters")]
        public Output<ImmutableDictionary<string, object>?> Parameters { get; private set; } = null!;

        /// <summary>
        /// An array of policy definition references.
        /// </summary>
        [Output("policyDefinitions")]
        public Output<ImmutableArray<Outputs.PolicyDefinitionReferenceResponseResult>> PolicyDefinitions { get; private set; } = null!;

        /// <summary>
        /// The type of policy definition. Possible values are NotSpecified, BuiltIn, and Custom.
        /// </summary>
        [Output("policyType")]
        public Output<string?> PolicyType { get; private set; } = null!;

        /// <summary>
        /// The type of the resource (Microsoft.Authorization/policySetDefinitions).
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a PolicySetDefinitionAtManagementGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PolicySetDefinitionAtManagementGroup(string name, PolicySetDefinitionAtManagementGroupArgs args, CustomResourceOptions? options = null)
            : base("azurerm:management/v20170601preview:PolicySetDefinitionAtManagementGroup", name, args ?? new PolicySetDefinitionAtManagementGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PolicySetDefinitionAtManagementGroup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:management/v20170601preview:PolicySetDefinitionAtManagementGroup", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:management/latest:PolicySetDefinitionAtManagementGroup"},
                    new Pulumi.Alias { Type = "azurerm:management/v20180301:PolicySetDefinitionAtManagementGroup"},
                    new Pulumi.Alias { Type = "azurerm:management/v20180501:PolicySetDefinitionAtManagementGroup"},
                    new Pulumi.Alias { Type = "azurerm:management/v20190101:PolicySetDefinitionAtManagementGroup"},
                    new Pulumi.Alias { Type = "azurerm:management/v20190601:PolicySetDefinitionAtManagementGroup"},
                    new Pulumi.Alias { Type = "azurerm:management/v20190901:PolicySetDefinitionAtManagementGroup"},
                    new Pulumi.Alias { Type = "azurerm:management/v20200301:PolicySetDefinitionAtManagementGroup"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PolicySetDefinitionAtManagementGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PolicySetDefinitionAtManagementGroup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PolicySetDefinitionAtManagementGroup(name, id, options);
        }
    }

    public sealed class PolicySetDefinitionAtManagementGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The policy set definition description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The display name of the policy set definition.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The ID of the management group.
        /// </summary>
        [Input("managementGroupId", required: true)]
        public Input<string> ManagementGroupId { get; set; } = null!;

        [Input("metadata")]
        private InputMap<object>? _metadata;

        /// <summary>
        /// The policy set definition metadata.
        /// </summary>
        public InputMap<object> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<object>());
            set => _metadata = value;
        }

        [Input("parameters")]
        private InputMap<object>? _parameters;

        /// <summary>
        /// The policy set definition parameters that can be used in policy definition references.
        /// </summary>
        public InputMap<object> Parameters
        {
            get => _parameters ?? (_parameters = new InputMap<object>());
            set => _parameters = value;
        }

        [Input("policyDefinitions", required: true)]
        private InputList<Inputs.PolicyDefinitionReferenceArgs>? _policyDefinitions;

        /// <summary>
        /// An array of policy definition references.
        /// </summary>
        public InputList<Inputs.PolicyDefinitionReferenceArgs> PolicyDefinitions
        {
            get => _policyDefinitions ?? (_policyDefinitions = new InputList<Inputs.PolicyDefinitionReferenceArgs>());
            set => _policyDefinitions = value;
        }

        /// <summary>
        /// The name of the policy set definition to create.
        /// </summary>
        [Input("policySetDefinitionName", required: true)]
        public Input<string> PolicySetDefinitionName { get; set; } = null!;

        /// <summary>
        /// The type of policy definition. Possible values are NotSpecified, BuiltIn, and Custom.
        /// </summary>
        [Input("policyType")]
        public Input<string>? PolicyType { get; set; }

        public PolicySetDefinitionAtManagementGroupArgs()
        {
        }
    }
}
