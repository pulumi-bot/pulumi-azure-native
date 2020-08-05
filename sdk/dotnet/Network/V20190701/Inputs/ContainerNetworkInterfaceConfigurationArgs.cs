// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20190701.Inputs
{

    /// <summary>
    /// Container network interface configuration child resource.
    /// </summary>
    public sealed class ContainerNetworkInterfaceConfigurationArgs : Pulumi.ResourceArgs
    {
        [Input("containerNetworkInterfaces")]
        private InputList<Inputs.SubResourceArgs>? _containerNetworkInterfaces;

        /// <summary>
        /// A list of container network interfaces created from this container network interface configuration.
        /// </summary>
        public InputList<Inputs.SubResourceArgs> ContainerNetworkInterfaces
        {
            get => _containerNetworkInterfaces ?? (_containerNetworkInterfaces = new InputList<Inputs.SubResourceArgs>());
            set => _containerNetworkInterfaces = value;
        }

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
        private InputList<Inputs.IPConfigurationProfileArgs>? _ipConfigurations;

        /// <summary>
        /// A list of ip configurations of the container network interface configuration.
        /// </summary>
        public InputList<Inputs.IPConfigurationProfileArgs> IpConfigurations
        {
            get => _ipConfigurations ?? (_ipConfigurations = new InputList<Inputs.IPConfigurationProfileArgs>());
            set => _ipConfigurations = value;
        }

        /// <summary>
        /// The name of the resource. This name can be used to access the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ContainerNetworkInterfaceConfigurationArgs()
        {
        }
    }
}
