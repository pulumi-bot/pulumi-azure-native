// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20181201
{
    /// <summary>
    /// VpnSite Resource.
    /// </summary>
    public partial class VpnSite : Pulumi.CustomResource
    {
        /// <summary>
        /// Gets a unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

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
        /// Parameters for VpnSite
        /// </summary>
        [Output("properties")]
        public Output<Outputs.VpnSitePropertiesResponseResult> Properties { get; private set; } = null!;

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
        /// Create a VpnSite resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpnSite(string name, VpnSiteArgs args, CustomResourceOptions? options = null)
            : base("azurerm:network/v20181201:VpnSite", name, args ?? new VpnSiteArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpnSite(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:network/v20181201:VpnSite", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing VpnSite resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpnSite Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new VpnSite(name, id, options);
        }
    }

    public sealed class VpnSiteArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The AddressSpace that contains an array of IP address ranges.
        /// </summary>
        [Input("addressSpace")]
        public Input<Inputs.AddressSpaceArgs>? AddressSpace { get; set; }

        /// <summary>
        /// The set of bgp properties.
        /// </summary>
        [Input("bgpProperties")]
        public Input<Inputs.BgpSettingsArgs>? BgpProperties { get; set; }

        /// <summary>
        /// The device properties
        /// </summary>
        [Input("deviceProperties")]
        public Input<Inputs.DevicePropertiesArgs>? DeviceProperties { get; set; }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The ip-address for the vpn-site.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// IsSecuritySite flag
        /// </summary>
        [Input("isSecuritySite")]
        public Input<bool>? IsSecuritySite { get; set; }

        /// <summary>
        /// Resource location.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The name of the VpnSite being created or updated.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The provisioning state of the resource.
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        /// <summary>
        /// The resource group name of the VpnSite.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The key for vpn-site that can be used for connections.
        /// </summary>
        [Input("siteKey")]
        public Input<string>? SiteKey { get; set; }

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
        /// The VirtualWAN to which the vpnSite belongs
        /// </summary>
        [Input("virtualWan")]
        public Input<Inputs.SubResourceArgs>? VirtualWan { get; set; }

        public VpnSiteArgs()
        {
        }
    }
}
