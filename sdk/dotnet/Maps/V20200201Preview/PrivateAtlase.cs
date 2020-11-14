// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Maps.V20200201Preview
{
    /// <summary>
    /// An Azure resource which represents which will provision the ability to create private location data.
    /// </summary>
    public partial class PrivateAtlase : Pulumi.CustomResource
    {
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The Private Atlas resource properties.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.PrivateAtlasPropertiesResponse> Properties { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a PrivateAtlase resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PrivateAtlase(string name, PrivateAtlaseArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:maps/v20200201preview:PrivateAtlase", name, args ?? new PrivateAtlaseArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PrivateAtlase(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:maps/v20200201preview:PrivateAtlase", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing PrivateAtlase resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PrivateAtlase Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PrivateAtlase(name, id, options);
        }
    }

    public sealed class PrivateAtlaseArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Maps Account.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The location of the resource.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The name of the Private Atlas instance.
        /// </summary>
        [Input("privateAtlasName", required: true)]
        public Input<string> PrivateAtlasName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Gets or sets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public PrivateAtlaseArgs()
        {
        }
    }
}
