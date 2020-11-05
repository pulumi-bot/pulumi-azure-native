// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.HybridNetwork.V20200101Preview.Inputs
{

    /// <summary>
    /// Network interface properties.
    /// </summary>
    public sealed class NetworkInterfaceArgs : Pulumi.ResourceArgs
    {
        [Input("ipConfigurations")]
        private InputList<Inputs.NetworkInterfaceIPConfigurationArgs>? _ipConfigurations;

        /// <summary>
        /// A list of IP configurations of the network interface.
        /// </summary>
        public InputList<Inputs.NetworkInterfaceIPConfigurationArgs> IpConfigurations
        {
            get => _ipConfigurations ?? (_ipConfigurations = new InputList<Inputs.NetworkInterfaceIPConfigurationArgs>());
            set => _ipConfigurations = value;
        }

        /// <summary>
        /// The MAC address of the network interface.
        /// </summary>
        [Input("macAddress")]
        public Input<string>? MacAddress { get; set; }

        /// <summary>
        /// The name of the network interface.
        /// </summary>
        [Input("networkInterfaceName")]
        public Input<string>? NetworkInterfaceName { get; set; }

        /// <summary>
        /// The type of the VM switch.
        /// </summary>
        [Input("vmSwitchType")]
        public Input<string>? VmSwitchType { get; set; }

        public NetworkInterfaceArgs()
        {
        }
    }
}
