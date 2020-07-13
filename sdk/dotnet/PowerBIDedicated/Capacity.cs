// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.PowerBIDedicated
{
    /// <summary>
    /// Represents an instance of a Dedicated Capacity resource.
    /// </summary>
    public partial class Capacity : Pulumi.CustomResource
    {
        /// <summary>
        /// Location of the PowerBI Dedicated resource.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the PowerBI Dedicated resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Properties of the provision operation request.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.DedicatedCapacityPropertiesResponse> Properties { get; private set; } = null!;

        /// <summary>
        /// The SKU of the PowerBI Dedicated resource.
        /// </summary>
        [Output("sku")]
        public Output<Outputs.ResourceSkuResponse> Sku { get; private set; } = null!;

        /// <summary>
        /// Key-value pairs of additional resource provisioning properties.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of the PowerBI Dedicated resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Capacity resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Capacity(string name, CapacityArgs args, CustomResourceOptions? options = null)
            : base("azurerm:powerbidedicated:Capacity", name, args ?? new CapacityArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Capacity(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:powerbidedicated:Capacity", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Capacity resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Capacity Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Capacity(name, id, options);
        }
    }

    public sealed class CapacityArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Location of the PowerBI Dedicated resource.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The name of the Dedicated capacity. It must be a minimum of 3 characters, and a maximum of 63.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Properties of the provision operation request.
        /// </summary>
        [Input("properties")]
        public Input<Inputs.DedicatedCapacityPropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// The name of the Azure Resource group of which a given PowerBIDedicated capacity is part. This name must be at least 1 character in length, and no more than 90.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The SKU of the PowerBI Dedicated resource.
        /// </summary>
        [Input("sku", required: true)]
        public Input<Inputs.ResourceSkuArgs> Sku { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value pairs of additional resource provisioning properties.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public CapacityArgs()
        {
        }
    }
}
