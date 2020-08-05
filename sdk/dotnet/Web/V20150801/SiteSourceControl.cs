// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Web.V20150801
{
    /// <summary>
    /// Describes the source control configuration for web app
    /// </summary>
    public partial class SiteSourceControl : Pulumi.CustomResource
    {
        /// <summary>
        /// Kind of resource
        /// </summary>
        [Output("kind")]
        public Output<string?> Kind { get; private set; } = null!;

        /// <summary>
        /// Resource Location
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Resource Name
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        [Output("properties")]
        public Output<Outputs.SiteSourceControlResponsePropertiesResult> Properties { get; private set; } = null!;

        /// <summary>
        /// Resource tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Resource type
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;


        /// <summary>
        /// Create a SiteSourceControl resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SiteSourceControl(string name, SiteSourceControlArgs args, CustomResourceOptions? options = null)
            : base("azurerm:web/v20150801:SiteSourceControl", name, args ?? new SiteSourceControlArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SiteSourceControl(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:web/v20150801:SiteSourceControl", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing SiteSourceControl resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SiteSourceControl Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new SiteSourceControl(name, id, options);
        }
    }

    public sealed class SiteSourceControlArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of branch to use for deployment
        /// </summary>
        [Input("branch")]
        public Input<string>? Branch { get; set; }

        /// <summary>
        /// Whether to manual or continuous integration
        /// </summary>
        [Input("deploymentRollbackEnabled")]
        public Input<bool>? DeploymentRollbackEnabled { get; set; }

        /// <summary>
        /// Resource Id
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Whether to manual or continuous integration
        /// </summary>
        [Input("isManualIntegration")]
        public Input<bool>? IsManualIntegration { get; set; }

        /// <summary>
        /// Mercurial or Git repository type
        /// </summary>
        [Input("isMercurial")]
        public Input<bool>? IsMercurial { get; set; }

        /// <summary>
        /// Kind of resource
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Resource Location
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Resource Name
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Repository or source control url
        /// </summary>
        [Input("repoUrl")]
        public Input<string>? RepoUrl { get; set; }

        /// <summary>
        /// Name of resource group
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Resource type
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public SiteSourceControlArgs()
        {
        }
    }
}
