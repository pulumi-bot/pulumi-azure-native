// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Cdn.V20171012.Inputs
{

    /// <summary>
    /// The main origin of CDN content which is added when creating a CDN endpoint.
    /// </summary>
    public sealed class DeepCreatedOriginArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The address of the origin. It can be a domain name, IPv4 address, or IPv6 address.
        /// </summary>
        [Input("hostName", required: true)]
        public Input<string> HostName { get; set; } = null!;

        /// <summary>
        /// The value of the HTTP port. Must be between 1 and 65535
        /// </summary>
        [Input("httpPort")]
        public Input<int>? HttpPort { get; set; }

        /// <summary>
        /// The value of the HTTPS port. Must be between 1 and 65535
        /// </summary>
        [Input("httpsPort")]
        public Input<int>? HttpsPort { get; set; }

        /// <summary>
        /// Origin name
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public DeepCreatedOriginArgs()
        {
        }
    }
}
