// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20191201.Inputs
{

    /// <summary>
    /// P2SConnectionConfiguration Resource.
    /// </summary>
    public sealed class P2SConnectionConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The reference to the address space resource which represents Address space for P2S VpnClient.
        /// </summary>
        [Input("vpnClientAddressPool")]
        public Input<Inputs.AddressSpaceArgs>? VpnClientAddressPool { get; set; }

        public P2SConnectionConfigurationArgs()
        {
        }
    }
}
