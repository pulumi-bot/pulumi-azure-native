// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20181001.Inputs
{

    /// <summary>
    /// NetworkInterface properties. 
    /// </summary>
    public sealed class NetworkInterfacePropertiesFormatResponseArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The DNS settings in network interface.
        /// </summary>
        [Input("dnsSettings")]
        public Inputs.NetworkInterfaceDnsSettingsResponseArgs? DnsSettings { get; set; }

        /// <summary>
        /// If the network interface is accelerated networking enabled.
        /// </summary>
        [Input("enableAcceleratedNetworking")]
        public bool? EnableAcceleratedNetworking { get; set; }

        /// <summary>
        /// Indicates whether IP forwarding is enabled on this network interface.
        /// </summary>
        [Input("enableIPForwarding")]
        public bool? EnableIPForwarding { get; set; }

        [Input("hostedWorkloads", required: true)]
        private List<string>? _hostedWorkloads;

        /// <summary>
        /// A list of references to linked BareMetal resources
        /// </summary>
        public List<string> HostedWorkloads
        {
            get => _hostedWorkloads ?? (_hostedWorkloads = new List<string>());
            set => _hostedWorkloads = value;
        }

        /// <summary>
        /// A reference to the interface endpoint to which the network interface is linked.
        /// </summary>
        [Input("interfaceEndpoint", required: true)]
        public Inputs.InterfaceEndpointResponseArgs InterfaceEndpoint { get; set; } = null!;

        [Input("ipConfigurations")]
        private List<Inputs.NetworkInterfaceIPConfigurationResponseArgs>? _ipConfigurations;

        /// <summary>
        /// A list of IPConfigurations of the network interface.
        /// </summary>
        public List<Inputs.NetworkInterfaceIPConfigurationResponseArgs> IpConfigurations
        {
            get => _ipConfigurations ?? (_ipConfigurations = new List<Inputs.NetworkInterfaceIPConfigurationResponseArgs>());
            set => _ipConfigurations = value;
        }

        /// <summary>
        /// The MAC address of the network interface.
        /// </summary>
        [Input("macAddress")]
        public string? MacAddress { get; set; }

        /// <summary>
        /// The reference of the NetworkSecurityGroup resource.
        /// </summary>
        [Input("networkSecurityGroup")]
        public Inputs.NetworkSecurityGroupResponseArgs? NetworkSecurityGroup { get; set; }

        /// <summary>
        /// Gets whether this is a primary network interface on a virtual machine.
        /// </summary>
        [Input("primary")]
        public bool? Primary { get; set; }

        /// <summary>
        /// The provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        [Input("provisioningState")]
        public string? ProvisioningState { get; set; }

        /// <summary>
        /// The resource GUID property of the network interface resource.
        /// </summary>
        [Input("resourceGuid")]
        public string? ResourceGuid { get; set; }

        [Input("tapConfigurations")]
        private List<Inputs.NetworkInterfaceTapConfigurationResponseArgs>? _tapConfigurations;

        /// <summary>
        /// A list of TapConfigurations of the network interface.
        /// </summary>
        public List<Inputs.NetworkInterfaceTapConfigurationResponseArgs> TapConfigurations
        {
            get => _tapConfigurations ?? (_tapConfigurations = new List<Inputs.NetworkInterfaceTapConfigurationResponseArgs>());
            set => _tapConfigurations = value;
        }

        /// <summary>
        /// The reference of a virtual machine.
        /// </summary>
        [Input("virtualMachine", required: true)]
        public Inputs.SubResourceResponseArgs VirtualMachine { get; set; } = null!;

        public NetworkInterfacePropertiesFormatResponseArgs()
        {
        }
    }
}
