// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DevTestLab.V20160515
{
    /// <summary>
    /// Properties of an artifact source.
    /// </summary>
    public partial class ArtifactSource : Pulumi.CustomResource
    {
        /// <summary>
        /// The location of the resource.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The properties of the resource.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.ArtifactSourcePropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// The tags of the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of the resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ArtifactSource resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ArtifactSource(string name, ArtifactSourceArgs args, CustomResourceOptions? options = null)
            : base("azurerm:devtestlab/v20160515:ArtifactSource", name, args ?? new ArtifactSourceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ArtifactSource(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:devtestlab/v20160515:ArtifactSource", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ArtifactSource resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ArtifactSource Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ArtifactSource(name, id, options);
        }
    }

    public sealed class ArtifactSourceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The folder containing Azure Resource Manager templates.
        /// </summary>
        [Input("armTemplateFolderPath")]
        public Input<string>? ArmTemplateFolderPath { get; set; }

        /// <summary>
        /// The artifact source's branch reference.
        /// </summary>
        [Input("branchRef")]
        public Input<string>? BranchRef { get; set; }

        /// <summary>
        /// The artifact source's display name.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The folder containing artifacts.
        /// </summary>
        [Input("folderPath")]
        public Input<string>? FolderPath { get; set; }

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
        /// The name of the artifact source.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

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
        /// The security token to authenticate to the artifact source.
        /// </summary>
        [Input("securityToken")]
        public Input<string>? SecurityToken { get; set; }

        /// <summary>
        /// The artifact source's type.
        /// </summary>
        [Input("sourceType")]
        public Input<string>? SourceType { get; set; }

        /// <summary>
        /// Indicates if the artifact source is enabled (values: Enabled, Disabled).
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
        /// The unique immutable identifier of a resource (Guid).
        /// </summary>
        [Input("uniqueIdentifier")]
        public Input<string>? UniqueIdentifier { get; set; }

        /// <summary>
        /// The artifact source's URI.
        /// </summary>
        [Input("uri")]
        public Input<string>? Uri { get; set; }

        public ArtifactSourceArgs()
        {
        }
    }
}
