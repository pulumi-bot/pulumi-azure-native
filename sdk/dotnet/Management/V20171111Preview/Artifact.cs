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
    /// Represents a Blueprint artifact.
    /// </summary>
    public partial class Artifact : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the kind of Blueprint artifact.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Name of this resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Type of this resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Artifact resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Artifact(string name, ArtifactArgs args, CustomResourceOptions? options = null)
            : base("azurerm:management/v20171111preview:Artifact", name, args ?? new ArtifactArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Artifact(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:management/v20171111preview:Artifact", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Artifact resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Artifact Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Artifact(name, id, options);
        }
    }

    public sealed class ArtifactArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// name of the artifact.
        /// </summary>
        [Input("artifactName", required: true)]
        public Input<string> ArtifactName { get; set; } = null!;

        /// <summary>
        /// name of the blueprint.
        /// </summary>
        [Input("blueprintName", required: true)]
        public Input<string> BlueprintName { get; set; } = null!;

        /// <summary>
        /// Specifies the kind of Blueprint artifact.
        /// </summary>
        [Input("kind", required: true)]
        public Input<string> Kind { get; set; } = null!;

        /// <summary>
        /// ManagementGroup where blueprint stores.
        /// </summary>
        [Input("managementGroupName", required: true)]
        public Input<string> ManagementGroupName { get; set; } = null!;

        public ArtifactArgs()
        {
        }
    }
}
