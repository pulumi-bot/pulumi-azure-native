// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20180201.Inputs
{

    /// <summary>
    /// NetworkInterface properties. 
    /// </summary>
    public sealed class NetworkInterfacePropertiesFormatArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The DNS settings in network interface.
        /// </summary>
        [Input("dnsSettings")]
        public Input<Inputs.NetworkInterfaceDnsSettingsArgs>? DnsSettings { get; set; }

        /// <summary>
        /// If the network interface is accelerated networking enabled.
        /// </summary>
        [Input("enableAcceleratedNetworking")]
        public Input<bool>? EnableAcceleratedNetworking { get; set; }

        /// <summary>
        /// Indicates whether IP forwarding is enabled on this network interface.
        /// </summary>
        [Input("enableIPForwarding")]
        public Input<bool>? EnableIPForwarding { get; set; }

        [Input("ipConfigurations")]
        private InputList<Inputs.NetworkInterfaceIPConfigurationArgs>? _ipConfigurations;

        /// <summary>
        /// A list of IPConfigurations of the network interface.
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
        /// The reference of the NetworkSecurityGroup resource.
        /// </summary>
        [Input("networkSecurityGroup")]
        public Input<Inputs.NetworkSecurityGroupArgs>? NetworkSecurityGroup { get; set; }

        /// <summary>
        /// Gets whether this is a primary network interface on a virtual machine.
        /// </summary>
        [Input("primary")]
        public Input<bool>? Primary { get; set; }

        /// <summary>
        /// The provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        /// <summary>
        /// The resource GUID property of the network interface resource.
        /// </summary>
        [Input("resourceGuid")]
        public Input<string>? ResourceGuid { get; set; }

        /// <summary>
        /// The reference of a virtual machine.
        /// </summary>
        [Input("virtualMachine")]
        public Input<Inputs.SubResourceArgs>? VirtualMachine { get; set; }

        public NetworkInterfacePropertiesFormatArgs()
        {
        }
    }
}