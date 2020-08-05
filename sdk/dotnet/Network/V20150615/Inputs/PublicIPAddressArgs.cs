// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20150615.Inputs
{

    /// <summary>
    /// Public IP address resource.
    /// </summary>
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
        /// Resource Identifier.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The idle timeout of the public IP address.
        /// </summary>
        [Input("idleTimeoutInMinutes")]
        public Input<int>? IdleTimeoutInMinutes { get; set; }

        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// IPConfiguration
        /// </summary>
        [Input("ipConfiguration")]
        public Input<Inputs.IPConfigurationArgs>? IpConfiguration { get; set; }

        /// <summary>
        /// Resource location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The provisioning state of the PublicIP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        /// <summary>
        /// The public IP allocation method. Possible values are: 'Static' and 'Dynamic'.
        /// </summary>
        [Input("publicIPAllocationMethod")]
        public Input<string>? PublicIPAllocationMethod { get; set; }

        /// <summary>
        /// The resource GUID property of the public IP resource.
        /// </summary>
        [Input("resourceGuid")]
        public Input<string>? ResourceGuid { get; set; }

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

        public PublicIPAddressArgs()
        {
        }
    }
}
