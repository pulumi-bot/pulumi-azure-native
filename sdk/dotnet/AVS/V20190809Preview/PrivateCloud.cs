// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.AVS.V20190809Preview
{
    /// <summary>
    /// A private cloud resource
    /// </summary>
    public partial class PrivateCloud : Pulumi.CustomResource
    {
        /// <summary>
        /// Resource location
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The properties of a private cloud resource
        /// </summary>
        [Output("properties")]
        public Output<Outputs.PrivateCloudPropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// The private cloud SKU
        /// </summary>
        [Output("sku")]
        public Output<Outputs.SkuResponseResult?> Sku { get; private set; } = null!;

        /// <summary>
        /// Resource tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a PrivateCloud resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PrivateCloud(string name, PrivateCloudArgs args, CustomResourceOptions? options = null)
            : base("azurerm:avs/v20190809preview:PrivateCloud", name, args ?? new PrivateCloudArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PrivateCloud(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:avs/v20190809preview:PrivateCloud", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:avs/latest:PrivateCloud"},
                    new Pulumi.Alias { Type = "azurerm:avs/v20200320:PrivateCloud"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PrivateCloud resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PrivateCloud Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PrivateCloud(name, id, options);
        }
    }

    public sealed class PrivateCloudArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource location
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Name of the private cloud
        /// </summary>
        [Input("privateCloudName", required: true)]
        public Input<string> PrivateCloudName { get; set; } = null!;

        /// <summary>
        /// The properties of a private cloud resource
        /// </summary>
        [Input("properties")]
        public Input<Inputs.PrivateCloudPropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// Name of the resource group within the Azure subscription
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The private cloud SKU
        /// </summary>
        [Input("sku")]
        public Input<Inputs.SkuArgs>? Sku { get; set; }

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

        public PrivateCloudArgs()
        {
        }
    }
}
