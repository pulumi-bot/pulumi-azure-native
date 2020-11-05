// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Resources.V20190601Preview
{
    /// <summary>
    /// Template Spec object.
    /// </summary>
    public partial class TemplateSpec : Pulumi.CustomResource
    {
        /// <summary>
        /// Template Spec description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Template Spec display name.
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The location of the Template Spec. It cannot be changed after Template Spec creation. It must be one of the supported Azure locations.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Name of this resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Type of this resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// High-level information about the versions within this Template Spec. The keys are the version names. Only populated if the $expand query parameter is set to 'versions'.
        /// </summary>
        [Output("versions")]
        public Output<ImmutableDictionary<string, Outputs.TemplateSpecVersionInfoResponse>> Versions { get; private set; } = null!;


        /// <summary>
        /// Create a TemplateSpec resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TemplateSpec(string name, TemplateSpecArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:resources/v20190601preview:TemplateSpec", name, args ?? new TemplateSpecArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TemplateSpec(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:resources/v20190601preview:TemplateSpec", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:resources/v20201001preview:TemplateSpec"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing TemplateSpec resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TemplateSpec Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new TemplateSpec(name, id, options);
        }
    }

    public sealed class TemplateSpecArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Template Spec description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Template Spec display name.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The location of the Template Spec. It cannot be changed after Template Spec creation. It must be one of the supported Azure locations.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
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

        /// <summary>
        /// Name of the Template Spec.
        /// </summary>
        [Input("templateSpecName", required: true)]
        public Input<string> TemplateSpecName { get; set; } = null!;

        public TemplateSpecArgs()
        {
        }
    }
}
