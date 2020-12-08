// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20200401.Inputs
{

    /// <summary>
    /// Backend address of a frontDoor load balancer.
    /// </summary>
    public sealed class BackendArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Location of the backend (IP address or FQDN)
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// The value to use as the host header sent to the backend. If blank or unspecified, this defaults to the incoming host.
        /// </summary>
        [Input("backendHostHeader")]
        public Input<string>? BackendHostHeader { get; set; }

        /// <summary>
        /// Whether to enable use of this backend. Permitted values are 'Enabled' or 'Disabled'
        /// </summary>
        [Input("enabledState")]
        public InputUnion<string, Pulumi.AzureNextGen.Network.V20200401.BackendEnabledState>? EnabledState { get; set; }

        /// <summary>
        /// The HTTP TCP port number. Must be between 1 and 65535.
        /// </summary>
        [Input("httpPort")]
        public Input<int>? HttpPort { get; set; }

        /// <summary>
        /// The HTTPS TCP port number. Must be between 1 and 65535.
        /// </summary>
        [Input("httpsPort")]
        public Input<int>? HttpsPort { get; set; }

        /// <summary>
        /// Priority to use for load balancing. Higher priorities will not be used for load balancing if any lower priority backend is healthy.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// The Alias of the Private Link resource. Populating this optional field indicates that this backend is 'Private'
        /// </summary>
        [Input("privateLinkAlias")]
        public Input<string>? PrivateLinkAlias { get; set; }

        /// <summary>
        /// A custom message to be included in the approval request to connect to the Private Link
        /// </summary>
        [Input("privateLinkApprovalMessage")]
        public Input<string>? PrivateLinkApprovalMessage { get; set; }

        /// <summary>
        /// Weight of this endpoint for load balancing purposes.
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public BackendArgs()
        {
        }
    }
}
