// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20190901.Inputs
{

    /// <summary>
    /// Public IP address resource.
    /// </summary>
    public sealed class PublicIPAddressArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The DDoS protection custom policy associated with the public IP address.
        /// </summary>
        [Input("ddosSettings")]
        public Input<Inputs.DdosSettingsArgs>? DdosSettings { get; set; }

        /// <summary>
        /// The FQDN of the DNS record associated with the public IP address.
        /// </summary>
        [Input("dnsSettings")]
        public Input<Inputs.PublicIPAddressDnsSettingsArgs>? DnsSettings { get; set; }

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

        [Input("ipTags")]
        private InputList<Inputs.IpTagArgs>? _ipTags;

        /// <summary>
        /// The list of tags associated with the public IP address.
        /// </summary>
        public InputList<Inputs.IpTagArgs> IpTags
        {
            get => _ipTags ?? (_ipTags = new InputList<Inputs.IpTagArgs>());
            set => _ipTags = value;
        }

        /// <summary>
        /// Resource location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The public IP address version.
        /// </summary>
        [Input("publicIPAddressVersion")]
        public InputUnion<string, Pulumi.AzureNextGen.Network.V20190901.IPVersion>? PublicIPAddressVersion { get; set; }

        /// <summary>
        /// The public IP address allocation method.
        /// </summary>
        [Input("publicIPAllocationMethod")]
        public InputUnion<string, Pulumi.AzureNextGen.Network.V20190901.IPAllocationMethod>? PublicIPAllocationMethod { get; set; }

        /// <summary>
        /// The Public IP Prefix this Public IP Address should be allocated from.
        /// </summary>
        [Input("publicIPPrefix")]
        public Input<Inputs.SubResourceArgs>? PublicIPPrefix { get; set; }

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
