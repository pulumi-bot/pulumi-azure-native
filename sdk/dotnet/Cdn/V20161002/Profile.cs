// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Cdn.V20161002
{
    /// <summary>
    /// CDN profile represents the top level resource and the entry point into the CDN API. This allows users to set up a logical grouping of endpoints in addition to creating shared configuration settings and selecting pricing tiers and providers.
    /// </summary>
    [AzureNextGenResourceType("azure-nextgen:cdn/v20161002:Profile")]
    public partial class Profile : Pulumi.CustomResource
    {
        /// <summary>
        /// Resource location.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Provisioning status of the profile.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Resource status of the profile.
        /// </summary>
        [Output("resourceState")]
        public Output<string> ResourceState { get; private set; } = null!;

        /// <summary>
        /// The pricing tier (defines a CDN provider, feature list and rate) of the CDN profile.
        /// </summary>
        [Output("sku")]
        public Output<Outputs.SkuResponse> Sku { get; private set; } = null!;

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
        /// Create a Profile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Profile(string name, ProfileArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:cdn/v20161002:Profile", name, args ?? new ProfileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Profile(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:cdn/v20161002:Profile", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:cdn/latest:Profile"},
                    new Pulumi.Alias { Type = "azure-nextgen:cdn/v20150601:Profile"},
                    new Pulumi.Alias { Type = "azure-nextgen:cdn/v20160402:Profile"},
                    new Pulumi.Alias { Type = "azure-nextgen:cdn/v20170402:Profile"},
                    new Pulumi.Alias { Type = "azure-nextgen:cdn/v20171012:Profile"},
                    new Pulumi.Alias { Type = "azure-nextgen:cdn/v20190415:Profile"},
                    new Pulumi.Alias { Type = "azure-nextgen:cdn/v20190615:Profile"},
                    new Pulumi.Alias { Type = "azure-nextgen:cdn/v20190615preview:Profile"},
                    new Pulumi.Alias { Type = "azure-nextgen:cdn/v20191231:Profile"},
                    new Pulumi.Alias { Type = "azure-nextgen:cdn/v20200331:Profile"},
                    new Pulumi.Alias { Type = "azure-nextgen:cdn/v20200415:Profile"},
                    new Pulumi.Alias { Type = "azure-nextgen:cdn/v20200901:Profile"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Profile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Profile Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Profile(name, id, options);
        }
    }

    public sealed class ProfileArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Name of the CDN profile which is unique within the resource group.
        /// </summary>
        [Input("profileName", required: true)]
        public Input<string> ProfileName { get; set; } = null!;

        /// <summary>
        /// Name of the Resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The pricing tier (defines a CDN provider, feature list and rate) of the CDN profile.
        /// </summary>
        [Input("sku", required: true)]
        public Input<Inputs.SkuArgs> Sku { get; set; } = null!;

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

        public ProfileArgs()
        {
        }
    }
}
