// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Peering.V20200401
{
    /// <summary>
    /// Peering Service
    /// </summary>
    public partial class PeeringService : Pulumi.CustomResource
    {
        /// <summary>
        /// The location of the resource.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The properties that define a peering service.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.PeeringServicePropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// The SKU that defines the type of the peering service.
        /// </summary>
        [Output("sku")]
        public Output<Outputs.PeeringServiceSkuResponseResult?> Sku { get; private set; } = null!;

        /// <summary>
        /// The resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of the resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a PeeringService resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PeeringService(string name, PeeringServiceArgs args, CustomResourceOptions? options = null)
            : base("azurerm:peering/v20200401:PeeringService", name, args ?? new PeeringServiceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PeeringService(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:peering/v20200401:PeeringService", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing PeeringService resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PeeringService Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PeeringService(name, id, options);
        }
    }

    public sealed class PeeringServiceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The location of the resource.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The name of the peering service.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The PeeringServiceLocation of the Customer.
        /// </summary>
        [Input("peeringServiceLocation")]
        public Input<string>? PeeringServiceLocation { get; set; }

        /// <summary>
        /// The MAPS Provider Name.
        /// </summary>
        [Input("peeringServiceProvider")]
        public Input<string>? PeeringServiceProvider { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The SKU that defines the type of the peering service.
        /// </summary>
        [Input("sku")]
        public Input<Inputs.PeeringServiceSkuArgs>? Sku { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public PeeringServiceArgs()
        {
        }
    }
}
