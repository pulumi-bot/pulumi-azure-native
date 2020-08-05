// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20170501
{
    /// <summary>
    /// Class representing a Traffic Manager profile.
    /// </summary>
    public partial class Profile : Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure Region where the resource lives
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The properties of the Traffic Manager profile.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.ProfilePropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. Ex- Microsoft.Network/trafficManagerProfiles.
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
            : base("azurerm:network/v20170501:Profile", name, args ?? new ProfileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Profile(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:network/v20170501:Profile", name, null, MakeResourceOptions(options, id))
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
        /// The DNS settings of the Traffic Manager profile.
        /// </summary>
        [Input("dnsConfig")]
        public Input<Inputs.DnsConfigArgs>? DnsConfig { get; set; }

        [Input("endpoints")]
        private InputList<Inputs.EndpointArgs>? _endpoints;

        /// <summary>
        /// The list of endpoints in the Traffic Manager profile.
        /// </summary>
        public InputList<Inputs.EndpointArgs> Endpoints
        {
            get => _endpoints ?? (_endpoints = new InputList<Inputs.EndpointArgs>());
            set => _endpoints = value;
        }

        /// <summary>
        /// The Azure Region where the resource lives
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The endpoint monitoring settings of the Traffic Manager profile.
        /// </summary>
        [Input("monitorConfig")]
        public Input<Inputs.MonitorConfigArgs>? MonitorConfig { get; set; }

        /// <summary>
        /// The name of the Traffic Manager profile.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The status of the Traffic Manager profile.
        /// </summary>
        [Input("profileStatus")]
        public Input<string>? ProfileStatus { get; set; }

        /// <summary>
        /// The name of the resource group containing the Traffic Manager profile.
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

        /// <summary>
        /// The traffic routing method of the Traffic Manager profile.
        /// </summary>
        [Input("trafficRoutingMethod")]
        public Input<string>? TrafficRoutingMethod { get; set; }

        public ProfileArgs()
        {
        }
    }
}
