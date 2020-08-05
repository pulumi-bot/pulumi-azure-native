// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20181201.Inputs
{

    /// <summary>
    /// Container network interface child resource.
    /// </summary>
    public sealed class ContainerNetworkInterfaceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Reference to the container to which this container network interface is attached.
        /// </summary>
        [Input("container")]
        public Input<Inputs.ContainerArgs>? Container { get; set; }

        /// <summary>
        /// Container network interface configuration from which this container network interface is created.
        /// </summary>
        [Input("containerNetworkInterfaceConfiguration")]
        public Input<Inputs.ContainerNetworkInterfaceConfigurationArgs>? ContainerNetworkInterfaceConfiguration { get; set; }

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

        [Input("ipConfigurations")]
        private InputList<Inputs.ContainerNetworkInterfaceIpConfigurationArgs>? _ipConfigurations;

        /// <summary>
        /// Reference to the ip configuration on this container nic.
        /// </summary>
        public InputList<Inputs.ContainerNetworkInterfaceIpConfigurationArgs> IpConfigurations
        {
            get => _ipConfigurations ?? (_ipConfigurations = new InputList<Inputs.ContainerNetworkInterfaceIpConfigurationArgs>());
            set => _ipConfigurations = value;
        }

        /// <summary>
        /// The name of the resource. This name can be used to access the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ContainerNetworkInterfaceArgs()
        {
        }
    }
}
