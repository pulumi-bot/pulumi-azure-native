// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20190401.Inputs
{

    /// <summary>
    /// PrivateLinkServiceConnection resource.
    /// </summary>
    public sealed class PrivateLinkServiceConnectionArgs : Pulumi.ResourceArgs
    {
        [Input("groupIds")]
        private InputList<string>? _groupIds;

        /// <summary>
        /// The ID(s) of the group(s) obtained from the remote resource that this private endpoint should connect to.
        /// </summary>
        public InputList<string> GroupIds
        {
            get => _groupIds ?? (_groupIds = new InputList<string>());
            set => _groupIds = value;
        }

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
        /// A collection of read-only information about the state of the connection to the remote resource.
        /// </summary>
        [Input("privateLinkServiceConnectionState")]
        public Input<Inputs.PrivateLinkServiceConnectionStateArgs>? PrivateLinkServiceConnectionState { get; set; }

        /// <summary>
        /// The resource id of private link service.
        /// </summary>
        [Input("privateLinkServiceId")]
        public Input<string>? PrivateLinkServiceId { get; set; }

        /// <summary>
        /// A message passed to the owner of the remote resource with this connection request. Restricted to 140 chars.
        /// </summary>
        [Input("requestMessage")]
        public Input<string>? RequestMessage { get; set; }

        public PrivateLinkServiceConnectionArgs()
        {
        }
    }
}
