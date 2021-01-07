// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20200801
{
    /// <summary>
    /// Custom IP prefix resource.
    /// </summary>
    [AzureNextGenResourceType("azure-nextgen:network/v20200801:CustomIPPrefix")]
    public partial class CustomIPPrefix : Pulumi.CustomResource
    {
        /// <summary>
        /// The prefix range in CIDR notation. Should include the start address and the prefix length.
        /// </summary>
        [Output("cidr")]
        public Output<string?> Cidr { get; private set; } = null!;

        /// <summary>
        /// The commissioned state of the Custom IP Prefix.
        /// </summary>
        [Output("commissionedState")]
        public Output<string?> CommissionedState { get; private set; } = null!;

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The extended location of the custom IP prefix.
        /// </summary>
        [Output("extendedLocation")]
        public Output<Outputs.ExtendedLocationResponse?> ExtendedLocation { get; private set; } = null!;

        /// <summary>
        /// Resource location.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the custom IP prefix resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The list of all referenced PublicIpPrefixes.
        /// </summary>
        [Output("publicIpPrefixes")]
        public Output<ImmutableArray<Outputs.SubResourceResponse>> PublicIpPrefixes { get; private set; } = null!;

        /// <summary>
        /// The resource GUID property of the custom IP prefix resource.
        /// </summary>
        [Output("resourceGuid")]
        public Output<string> ResourceGuid { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// A list of availability zones denoting the IP allocated for the resource needs to come from.
        /// </summary>
        [Output("zones")]
        public Output<ImmutableArray<string>> Zones { get; private set; } = null!;


        /// <summary>
        /// Create a CustomIPPrefix resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CustomIPPrefix(string name, CustomIPPrefixArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:network/v20200801:CustomIPPrefix", name, args ?? new CustomIPPrefixArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CustomIPPrefix(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:network/v20200801:CustomIPPrefix", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:network/latest:CustomIPPrefix"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200601:CustomIPPrefix"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200701:CustomIPPrefix"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CustomIPPrefix resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CustomIPPrefix Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new CustomIPPrefix(name, id, options);
        }
    }

    public sealed class CustomIPPrefixArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The prefix range in CIDR notation. Should include the start address and the prefix length.
        /// </summary>
        [Input("cidr")]
        public Input<string>? Cidr { get; set; }

        /// <summary>
        /// The commissioned state of the Custom IP Prefix.
        /// </summary>
        [Input("commissionedState")]
        public InputUnion<string, Pulumi.AzureNextGen.Network.V20200801.CommissionedState>? CommissionedState { get; set; }

        /// <summary>
        /// The name of the custom IP prefix.
        /// </summary>
        [Input("customIpPrefixName", required: true)]
        public Input<string> CustomIpPrefixName { get; set; } = null!;

        /// <summary>
        /// The extended location of the custom IP prefix.
        /// </summary>
        [Input("extendedLocation")]
        public Input<Inputs.ExtendedLocationArgs>? ExtendedLocation { get; set; }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Resource location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("zones")]
        private InputList<string>? _zones;

        /// <summary>
        /// A list of availability zones denoting the IP allocated for the resource needs to come from.
        /// </summary>
        public InputList<string> Zones
        {
            get => _zones ?? (_zones = new InputList<string>());
            set => _zones = value;
        }

        public CustomIPPrefixArgs()
        {
        }
    }
}
