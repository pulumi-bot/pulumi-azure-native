// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ContainerInstance.V20171201Preview.Inputs
{

    /// <summary>
    /// IP address for the container group.
    /// </summary>
    public sealed class IpAddressArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The IP exposed to the public internet.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        [Input("ports", required: true)]
        private InputList<Inputs.PortArgs>? _ports;

        /// <summary>
        /// The list of ports exposed on the container group.
        /// </summary>
        public InputList<Inputs.PortArgs> Ports
        {
            get => _ports ?? (_ports = new InputList<Inputs.PortArgs>());
            set => _ports = value;
        }

        /// <summary>
        /// Specifies if the IP is exposed to the public internet.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public IpAddressArgs()
        {
        }
    }
}
