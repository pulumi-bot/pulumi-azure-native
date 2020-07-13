// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Migrate
{
    /// <summary>
    /// Azure Migrate Project.
    /// </summary>
    public partial class AssessmentProject : Pulumi.CustomResource
    {
        /// <summary>
        /// For optimistic concurrency control.
        /// </summary>
        [Output("eTag")]
        public Output<string?> ETag { get; private set; } = null!;

        /// <summary>
        /// Azure location in which project is created.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// Name of the project.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Properties of the project.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.ProjectPropertiesResponse> Properties { get; private set; } = null!;

        /// <summary>
        /// Tags provided by Azure Tagging service.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Type of the object = [Microsoft.Migrate/assessmentProjects].
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a AssessmentProject resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AssessmentProject(string name, AssessmentProjectArgs args, CustomResourceOptions? options = null)
            : base("azurerm:migrate:AssessmentProject", name, args ?? new AssessmentProjectArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AssessmentProject(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:migrate:AssessmentProject", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing AssessmentProject resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AssessmentProject Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AssessmentProject(name, id, options);
        }
    }

    public sealed class AssessmentProjectArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// For optimistic concurrency control.
        /// </summary>
        [Input("eTag")]
        public Input<string>? ETag { get; set; }

        /// <summary>
        /// Azure location in which project is created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Name of the Azure Migrate project.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Properties of the project.
        /// </summary>
        [Input("properties")]
        public Input<Inputs.ProjectPropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// Name of the Azure Resource Group that project is part of.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Tags provided by Azure Tagging service.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public AssessmentProjectArgs()
        {
        }
    }
}
