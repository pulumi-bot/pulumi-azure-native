// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.VirtualMachineImages
{
    /// <summary>
    /// Image template is an ARM resource managed by Microsoft.VirtualMachineImages provider
    /// </summary>
    public partial class ImageTemplate : Pulumi.CustomResource
    {
        /// <summary>
        /// The identity of the image template, if configured.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.ImageTemplateIdentityResponse> Identity { get; private set; } = null!;

        /// <summary>
        /// Resource location
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The properties of the image template
        /// </summary>
        [Output("properties")]
        public Output<Outputs.ImageTemplatePropertiesResponse> Properties { get; private set; } = null!;

        /// <summary>
        /// Resource tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Resource type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ImageTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ImageTemplate(string name, ImageTemplateArgs args, CustomResourceOptions? options = null)
            : base("azurerm:virtualmachineimages:ImageTemplate", name, args ?? new ImageTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ImageTemplate(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:virtualmachineimages:ImageTemplate", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ImageTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ImageTemplate Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ImageTemplate(name, id, options);
        }
    }

    public sealed class ImageTemplateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The identity of the image template, if configured.
        /// </summary>
        [Input("identity", required: true)]
        public Input<Inputs.ImageTemplateIdentityArgs> Identity { get; set; } = null!;

        /// <summary>
        /// Resource location
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The name of the image Template
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The properties of the image template
        /// </summary>
        [Input("properties")]
        public Input<Inputs.ImageTemplatePropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// The name of the resource group.
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

        public ImageTemplateArgs()
        {
        }
    }
}
