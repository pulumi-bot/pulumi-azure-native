// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Peering.V20200401
{
    /// <summary>
    /// Peering is a logical representation of a set of connections to the Microsoft Cloud Edge at a location.
    /// </summary>
    [AzureNextGenResourceType("azure-nextgen:peering/v20200401:Peering")]
    public partial class Peering : Pulumi.CustomResource
    {
        /// <summary>
        /// The properties that define a direct peering.
        /// </summary>
        [Output("direct")]
        public Output<Outputs.PeeringPropertiesDirectResponse?> Direct { get; private set; } = null!;

        /// <summary>
        /// The properties that define an exchange peering.
        /// </summary>
        [Output("exchange")]
        public Output<Outputs.PeeringPropertiesExchangeResponse?> Exchange { get; private set; } = null!;

        /// <summary>
        /// The kind of the peering.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

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
        /// The location of the peering.
        /// </summary>
        [Output("peeringLocation")]
        public Output<string?> PeeringLocation { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The SKU that defines the tier and kind of the peering.
        /// </summary>
        [Output("sku")]
        public Output<Outputs.PeeringSkuResponse> Sku { get; private set; } = null!;

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
        /// Create a Peering resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Peering(string name, PeeringArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:peering/v20200401:Peering", name, args ?? new PeeringArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Peering(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:peering/v20200401:Peering", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:peering/latest:Peering"},
                    new Pulumi.Alias { Type = "azure-nextgen:peering/v20190801preview:Peering"},
                    new Pulumi.Alias { Type = "azure-nextgen:peering/v20190901preview:Peering"},
                    new Pulumi.Alias { Type = "azure-nextgen:peering/v20200101preview:Peering"},
                    new Pulumi.Alias { Type = "azure-nextgen:peering/v20201001:Peering"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Peering resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Peering Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Peering(name, id, options);
        }
    }

    public sealed class PeeringArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The properties that define a direct peering.
        /// </summary>
        [Input("direct")]
        public Input<Inputs.PeeringPropertiesDirectArgs>? Direct { get; set; }

        /// <summary>
        /// The properties that define an exchange peering.
        /// </summary>
        [Input("exchange")]
        public Input<Inputs.PeeringPropertiesExchangeArgs>? Exchange { get; set; }

        /// <summary>
        /// The kind of the peering.
        /// </summary>
        [Input("kind", required: true)]
        public InputUnion<string, Pulumi.AzureNextGen.Peering.V20200401.Kind> Kind { get; set; } = null!;

        /// <summary>
        /// The location of the resource.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The location of the peering.
        /// </summary>
        [Input("peeringLocation")]
        public Input<string>? PeeringLocation { get; set; }

        /// <summary>
        /// The name of the peering.
        /// </summary>
        [Input("peeringName", required: true)]
        public Input<string> PeeringName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The SKU that defines the tier and kind of the peering.
        /// </summary>
        [Input("sku", required: true)]
        public Input<Inputs.PeeringSkuArgs> Sku { get; set; } = null!;

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

        public PeeringArgs()
        {
        }
    }
}
