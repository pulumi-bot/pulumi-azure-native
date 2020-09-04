// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Management.V20171111Preview
{
    /// <summary>
    /// Represents a Blueprint definition.
    /// </summary>
    public partial class Blueprint : Pulumi.CustomResource
    {
        /// <summary>
        /// Multi-line explain this resource.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// One-liner string explain this resource.
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Layout view of the blueprint, for UI reference.
        /// </summary>
        [Output("layout")]
        public Output<ImmutableDictionary<string, object>?> Layout { get; private set; } = null!;

        /// <summary>
        /// Name of this resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Parameters required by this Blueprint definition.
        /// </summary>
        [Output("parameters")]
        public Output<ImmutableDictionary<string, Outputs.ParameterDefinitionResponseResult>?> Parameters { get; private set; } = null!;

        /// <summary>
        /// Resource group placeholders defined by this Blueprint definition.
        /// </summary>
        [Output("resourceGroups")]
        public Output<ImmutableDictionary<string, Outputs.ResourceGroupDefinitionResponseResult>?> ResourceGroups { get; private set; } = null!;

        /// <summary>
        /// Status of the Blueprint. This field is readonly.
        /// </summary>
        [Output("status")]
        public Output<Outputs.BlueprintStatusResponseResult> Status { get; private set; } = null!;

        /// <summary>
        /// The scope where this Blueprint can be applied.
        /// </summary>
        [Output("targetScope")]
        public Output<string> TargetScope { get; private set; } = null!;

        /// <summary>
        /// Type of this resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Published versions of this blueprint.
        /// </summary>
        [Output("versions")]
        public Output<ImmutableDictionary<string, object>?> Versions { get; private set; } = null!;


        /// <summary>
        /// Create a Blueprint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Blueprint(string name, BlueprintArgs args, CustomResourceOptions? options = null)
            : base("azurerm:management/v20171111preview:Blueprint", name, args ?? new BlueprintArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Blueprint(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:management/v20171111preview:Blueprint", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Blueprint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Blueprint Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Blueprint(name, id, options);
        }
    }

    public sealed class BlueprintArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// name of the blueprint.
        /// </summary>
        [Input("blueprintName", required: true)]
        public Input<string> BlueprintName { get; set; } = null!;

        /// <summary>
        /// Multi-line explain this resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// One-liner string explain this resource.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("layout")]
        private InputMap<object>? _layout;

        /// <summary>
        /// Layout view of the blueprint, for UI reference.
        /// </summary>
        public InputMap<object> Layout
        {
            get => _layout ?? (_layout = new InputMap<object>());
            set => _layout = value;
        }

        /// <summary>
        /// ManagementGroup where blueprint stores.
        /// </summary>
        [Input("managementGroupName", required: true)]
        public Input<string> ManagementGroupName { get; set; } = null!;

        [Input("parameters")]
        private InputMap<Inputs.ParameterDefinitionArgs>? _parameters;

        /// <summary>
        /// Parameters required by this Blueprint definition.
        /// </summary>
        public InputMap<Inputs.ParameterDefinitionArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputMap<Inputs.ParameterDefinitionArgs>());
            set => _parameters = value;
        }

        [Input("resourceGroups")]
        private InputMap<Inputs.ResourceGroupDefinitionArgs>? _resourceGroups;

        /// <summary>
        /// Resource group placeholders defined by this Blueprint definition.
        /// </summary>
        public InputMap<Inputs.ResourceGroupDefinitionArgs> ResourceGroups
        {
            get => _resourceGroups ?? (_resourceGroups = new InputMap<Inputs.ResourceGroupDefinitionArgs>());
            set => _resourceGroups = value;
        }

        /// <summary>
        /// The scope where this Blueprint can be applied.
        /// </summary>
        [Input("targetScope", required: true)]
        public Input<string> TargetScope { get; set; } = null!;

        [Input("versions")]
        private InputMap<object>? _versions;

        /// <summary>
        /// Published versions of this blueprint.
        /// </summary>
        public InputMap<object> Versions
        {
            get => _versions ?? (_versions = new InputMap<object>());
            set => _versions = value;
        }

        public BlueprintArgs()
        {
        }
    }
}
