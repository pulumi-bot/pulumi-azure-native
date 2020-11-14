// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20200701.Outputs
{

    [OutputType]
    public sealed class ProbeResponse
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The interval, in seconds, for how frequently to probe the endpoint for health status. Typically, the interval is slightly less than half the allocated timeout period (in seconds) which allows two full probes before taking the instance out of rotation. The default value is 15, the minimum value is 5.
        /// </summary>
        public readonly int? IntervalInSeconds;
        /// <summary>
        /// The load balancer rules that use this probe.
        /// </summary>
        public readonly ImmutableArray<Outputs.SubResourceResponse> LoadBalancingRules;
        /// <summary>
        /// The name of the resource that is unique within the set of probes used by the load balancer. This name can be used to access the resource.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The number of probes where if no response, will result in stopping further traffic from being delivered to the endpoint. This values allows endpoints to be taken out of rotation faster or slower than the typical times used in Azure.
        /// </summary>
        public readonly int? NumberOfProbes;
        /// <summary>
        /// The port for communicating the probe. Possible values range from 1 to 65535, inclusive.
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// The protocol of the end point. If 'Tcp' is specified, a received ACK is required for the probe to be successful. If 'Http' or 'Https' is specified, a 200 OK response from the specifies URI is required for the probe to be successful.
        /// </summary>
        public readonly string Protocol;
        /// <summary>
        /// The provisioning state of the probe resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The URI used for requesting health status from the VM. Path is required if a protocol is set to http. Otherwise, it is not allowed. There is no default value.
        /// </summary>
        public readonly string? RequestPath;
        /// <summary>
        /// Type of the resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ProbeResponse(
            string etag,

            string? id,

            int? intervalInSeconds,

            ImmutableArray<Outputs.SubResourceResponse> loadBalancingRules,

            string? name,

            int? numberOfProbes,

            int port,

            string protocol,

            string provisioningState,

            string? requestPath,

            string type)
        {
            Etag = etag;
            Id = id;
            IntervalInSeconds = intervalInSeconds;
            LoadBalancingRules = loadBalancingRules;
            Name = name;
            NumberOfProbes = numberOfProbes;
            Port = port;
            Protocol = protocol;
            ProvisioningState = provisioningState;
            RequestPath = requestPath;
            Type = type;
        }
    }
}
