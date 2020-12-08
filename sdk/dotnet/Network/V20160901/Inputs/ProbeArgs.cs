// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20160901.Inputs
{

    /// <summary>
    /// A load balancer probe.
    /// </summary>
    public sealed class ProbeArgs : Pulumi.ResourceArgs
    {
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
        /// The interval, in seconds, for how frequently to probe the endpoint for health status. Typically, the interval is slightly less than half the allocated timeout period (in seconds) which allows two full probes before taking the instance out of rotation. The default value is 15, the minimum value is 5.
        /// </summary>
        [Input("intervalInSeconds")]
        public Input<int>? IntervalInSeconds { get; set; }

        /// <summary>
        /// Gets name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The number of probes where if no response, will result in stopping further traffic from being delivered to the endpoint. This values allows endpoints to be taken out of rotation faster or slower than the typical times used in Azure.
        /// </summary>
        [Input("numberOfProbes")]
        public Input<int>? NumberOfProbes { get; set; }

        /// <summary>
        /// The port for communicating the probe. Possible values range from 1 to 65535, inclusive.
        /// </summary>
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        /// <summary>
        /// The protocol of the end point. Possible values are: 'Http' or 'Tcp'. If 'Tcp' is specified, a received ACK is required for the probe to be successful. If 'Http' is specified, a 200 OK response from the specifies URI is required for the probe to be successful.
        /// </summary>
        [Input("protocol", required: true)]
        public InputUnion<string, Pulumi.AzureNextGen.Network.V20160901.ProbeProtocol> Protocol { get; set; } = null!;

        /// <summary>
        /// Gets the provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        /// <summary>
        /// The URI used for requesting health status from the VM. Path is required if a protocol is set to http. Otherwise, it is not allowed. There is no default value.
        /// </summary>
        [Input("requestPath")]
        public Input<string>? RequestPath { get; set; }

        public ProbeArgs()
        {
        }
    }
}
