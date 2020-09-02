// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Media.V20180601Preview
{
    /// <summary>
    /// A Transform encapsulates the rules or instructions for generating desired outputs from input media, such as by transcoding or by extracting insights. After the Transform is created, it can be applied to input media by creating Jobs.
    /// </summary>
    public partial class Transform : Pulumi.CustomResource
    {
        /// <summary>
        /// The UTC date and time when the Transform was created, in 'YYYY-MM-DDThh:mm:ssZ' format.
        /// </summary>
        [Output("created")]
        public Output<string> Created { get; private set; } = null!;

        /// <summary>
        /// An optional verbose description of the Transform.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The UTC date and time when the Transform was last updated, in 'YYYY-MM-DDThh:mm:ssZ' format.
        /// </summary>
        [Output("lastModified")]
        public Output<string> LastModified { get; private set; } = null!;

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// An array of one or more TransformOutputs that the Transform should generate.
        /// </summary>
        [Output("outputs")]
        public Output<ImmutableArray<Outputs.TransformOutputResponseResult>> Outputs { get; private set; } = null!;

        /// <summary>
        /// The type of the resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Transform resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Transform(string name, TransformArgs args, CustomResourceOptions? options = null)
            : base("azurerm:media/v20180601preview:Transform", name, args ?? new TransformArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Transform(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:media/v20180601preview:Transform", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:media/latest:Transform"},
                    new Pulumi.Alias { Type = "azurerm:media/v20180330preview:Transform"},
                    new Pulumi.Alias { Type = "azurerm:media/v20180701:Transform"},
                    new Pulumi.Alias { Type = "azurerm:media/v20200501:Transform"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Transform resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Transform Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Transform(name, id, options);
        }
    }

    public sealed class TransformArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Media Services account name.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// An optional verbose description of the Transform.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("outputs", required: true)]
        private InputList<Inputs.TransformOutputArgs>? _outputs;

        /// <summary>
        /// An array of one or more TransformOutputs that the Transform should generate.
        /// </summary>
        public InputList<Inputs.TransformOutputArgs> Outputs
        {
            get => _outputs ?? (_outputs = new InputList<Inputs.TransformOutputArgs>());
            set => _outputs = value;
        }

        /// <summary>
        /// The name of the resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The Transform name.
        /// </summary>
        [Input("transformName", required: true)]
        public Input<string> TransformName { get; set; } = null!;

        public TransformArgs()
        {
        }
    }
}
