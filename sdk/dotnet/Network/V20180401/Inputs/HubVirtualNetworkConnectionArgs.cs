// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20180401.Inputs
{

    /// <summary>
    /// HubVirtualNetworkConnection Resource.
    /// </summary>
    public sealed class HubVirtualNetworkConnectionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// VirtualHub to RemoteVnet transit to enabled or not.
        /// </summary>
        [Input("allowHubToRemoteVnetTransit")]
        public Input<bool>? AllowHubToRemoteVnetTransit { get; set; }

        /// <summary>
        /// Allow RemoteVnet to use Virtual Hub's gateways.
        /// </summary>
        [Input("allowRemoteVnetToUseHubVnetGateways")]
        public Input<bool>? AllowRemoteVnetToUseHubVnetGateways { get; set; }

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
        /// The provisioning state of the resource.
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        /// <summary>
        /// Reference to the remote virtual network.
        /// </summary>
        [Input("remoteVirtualNetwork")]
        public Input<Inputs.SubResourceArgs>? RemoteVirtualNetwork { get; set; }

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

        public HubVirtualNetworkConnectionArgs()
        {
        }
    }
}
