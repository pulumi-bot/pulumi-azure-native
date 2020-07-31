// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20170601.Inputs
{

    /// <summary>
    /// Properties of Backend Address Pool of an application gateway.
    /// </summary>
    public sealed class ApplicationGatewayBackendAddressPoolPropertiesFormatArgs : Pulumi.ResourceArgs
    {
        [Input("backendAddresses")]
        private InputList<Inputs.ApplicationGatewayBackendAddressArgs>? _backendAddresses;

        /// <summary>
        /// Backend addresses
        /// </summary>
        public InputList<Inputs.ApplicationGatewayBackendAddressArgs> BackendAddresses
        {
            get => _backendAddresses ?? (_backendAddresses = new InputList<Inputs.ApplicationGatewayBackendAddressArgs>());
            set => _backendAddresses = value;
        }

        [Input("backendIPConfigurations")]
        private InputList<Inputs.NetworkInterfaceIPConfigurationArgs>? _backendIPConfigurations;

        /// <summary>
        /// Collection of references to IPs defined in network interfaces.
        /// </summary>
        public InputList<Inputs.NetworkInterfaceIPConfigurationArgs> BackendIPConfigurations
        {
            get => _backendIPConfigurations ?? (_backendIPConfigurations = new InputList<Inputs.NetworkInterfaceIPConfigurationArgs>());
            set => _backendIPConfigurations = value;
        }

        /// <summary>
        /// Provisioning state of the backend address pool resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        public ApplicationGatewayBackendAddressPoolPropertiesFormatArgs()
        {
        }
    }
}