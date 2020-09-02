// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Migrate.V20180901Preview
{
    /// <summary>
    /// Migrate Project REST Resource.
    /// </summary>
    public partial class MigrateProject : Pulumi.CustomResource
    {
        /// <summary>
        /// Gets or sets the eTag for concurrency control.
        /// </summary>
        [Output("eTag")]
        public Output<string?> ETag { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the Azure location in which migrate project is created.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// Gets the name of the migrate project.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the nested properties.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.MigrateProjectPropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the tags.
        /// </summary>
        [Output("tags")]
        public Output<Outputs.MigrateProjectResponseTagsResult?> Tags { get; private set; } = null!;

        /// <summary>
        /// Handled by resource provider. Type = Microsoft.Migrate/MigrateProject.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a MigrateProject resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MigrateProject(string name, MigrateProjectArgs args, CustomResourceOptions? options = null)
            : base("azurerm:migrate/v20180901preview:MigrateProject", name, args ?? new MigrateProjectArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MigrateProject(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:migrate/v20180901preview:MigrateProject", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing MigrateProject resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MigrateProject Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new MigrateProject(name, id, options);
        }
    }

    public sealed class MigrateProjectArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gets or sets the eTag for concurrency control.
        /// </summary>
        [Input("eTag")]
        public Input<string>? ETag { get; set; }

        /// <summary>
        /// Gets or sets the Azure location in which migrate project is created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Name of the Azure Migrate project.
        /// </summary>
        [Input("migrateProjectName", required: true)]
        public Input<string> MigrateProjectName { get; set; } = null!;

        /// <summary>
        /// Gets or sets the nested properties.
        /// </summary>
        [Input("properties")]
        public Input<Inputs.MigrateProjectPropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// Name of the Azure Resource Group that migrate project is part of.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Gets or sets the tags.
        /// </summary>
        [Input("tags")]
        public Input<Inputs.MigrateProjectTagsArgs>? Tags { get; set; }

        public MigrateProjectArgs()
        {
        }
    }
}
