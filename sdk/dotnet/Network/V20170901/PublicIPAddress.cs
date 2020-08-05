// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20170901
{
    /// <summary>
    /// Public IP address resource.
    /// </summary>
    public partial class PublicIPAddress : Pulumi.CustomResource
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string?> Etag { get; private set; } = null!;

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
        /// Public IP address properties.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.PublicIPAddressPropertiesFormatResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// The public IP address SKU.
        /// </summary>
        [Output("sku")]
        public Output<Outputs.PublicIPAddressSkuResponseResult?> Sku { get; private set; } = null!;

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
        /// Create a PublicIPAddress resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PublicIPAddress(string name, PublicIPAddressArgs args, CustomResourceOptions? options = null)
            : base("azurerm:network/v20170901:PublicIPAddress", name, args ?? new PublicIPAddressArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PublicIPAddress(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:network/v20170901:PublicIPAddress", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing PublicIPAddress resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PublicIPAddress Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PublicIPAddress(name, id, options);
        }
    }

    public sealed class PublicIPAddressArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The FQDN of the DNS record associated with the public IP address.
        /// </summary>
        [Input("dnsSettings")]
        public Input<Inputs.PublicIPAddressDnsSettingsArgs>? DnsSettings { get; set; }

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The idle timeout of the public IP address.
        /// </summary>
        [Input("idleTimeoutInMinutes")]
        public Input<int>? IdleTimeoutInMinutes { get; set; }

        /// <summary>
        /// The IP address associated with the public IP address resource.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// Resource location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the public IP address.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The provisioning state of the PublicIP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        /// <summary>
        /// The public IP address version. Possible values are: 'IPv4' and 'IPv6'.
        /// </summary>
        [Input("publicIPAddressVersion")]
        public Input<string>? PublicIPAddressVersion { get; set; }

        /// <summary>
        /// The public IP allocation method. Possible values are: 'Static' and 'Dynamic'.
        /// </summary>
        [Input("publicIPAllocationMethod")]
        public Input<string>? PublicIPAllocationMethod { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The resource GUID property of the public IP resource.
        /// </summary>
        [Input("resourceGuid")]
        public Input<string>? ResourceGuid { get; set; }

        /// <summary>
        /// The public IP address SKU.
        /// </summary>
        [Input("sku")]
        public Input<Inputs.PublicIPAddressSkuArgs>? Sku { get; set; }

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

        public PublicIPAddressArgs()
        {
        }
    }
}
