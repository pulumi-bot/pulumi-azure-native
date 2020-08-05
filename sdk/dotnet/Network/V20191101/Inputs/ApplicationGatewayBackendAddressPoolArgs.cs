// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20191101.Inputs
{

    /// <summary>
    /// Backend Address Pool of an application gateway.
    /// </summary>
    public sealed class ApplicationGatewayBackendAddressPoolArgs : Pulumi.ResourceArgs
    {
        [Input("backendAddresses")]
        private InputList<Inputs.ApplicationGatewayBackendAddressArgs>? _backendAddresses;

        /// <summary>
        /// Backend addresses.
        /// </summary>
        public InputList<Inputs.ApplicationGatewayBackendAddressArgs> BackendAddresses
        {
            get => _backendAddresses ?? (_backendAddresses = new InputList<Inputs.ApplicationGatewayBackendAddressArgs>());
            set => _backendAddresses = value;
        }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Name of the backend address pool that is unique within an Application Gateway.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ApplicationGatewayBackendAddressPoolArgs()
        {
        }
    }
}
