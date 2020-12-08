// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.ContainerInstance.V20180401.Inputs
{

    /// <summary>
    /// The port exposed on the container instance.
    /// </summary>
    public sealed class ContainerPortArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The port number exposed within the container group.
        /// </summary>
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        /// <summary>
        /// The protocol associated with the port.
        /// </summary>
        [Input("protocol")]
        public InputUnion<string, Pulumi.AzureNextGen.ContainerInstance.V20180401.ContainerNetworkProtocol>? Protocol { get; set; }

        public ContainerPortArgs()
        {
        }
    }
}
