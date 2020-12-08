// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20170601.Inputs
{

    /// <summary>
    /// Backend address pool settings of an application gateway.
    /// </summary>
    public sealed class ApplicationGatewayBackendHttpSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cookie name to use for the affinity cookie.
        /// </summary>
        [Input("affinityCookieName")]
        public Input<string>? AffinityCookieName { get; set; }

        [Input("authenticationCertificates")]
        private InputList<Inputs.SubResourceArgs>? _authenticationCertificates;

        /// <summary>
        /// Array of references to application gateway authentication certificates.
        /// </summary>
        public InputList<Inputs.SubResourceArgs> AuthenticationCertificates
        {
            get => _authenticationCertificates ?? (_authenticationCertificates = new InputList<Inputs.SubResourceArgs>());
            set => _authenticationCertificates = value;
        }

        /// <summary>
        /// Connection draining of the backend http settings resource.
        /// </summary>
        [Input("connectionDraining")]
        public Input<Inputs.ApplicationGatewayConnectionDrainingArgs>? ConnectionDraining { get; set; }

        /// <summary>
        /// Cookie based affinity.
        /// </summary>
        [Input("cookieBasedAffinity")]
        public InputUnion<string, Pulumi.AzureNextGen.Network.V20170601.ApplicationGatewayCookieBasedAffinity>? CookieBasedAffinity { get; set; }

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Host header to be sent to the backend servers.
        /// </summary>
        [Input("hostName")]
        public Input<string>? HostName { get; set; }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Path which should be used as a prefix for all HTTP requests. Null means no path will be prefixed. Default value is null.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// Whether to pick host header should be picked from the host name of the backend server. Default value is false.
        /// </summary>
        [Input("pickHostNameFromBackendAddress")]
        public Input<bool>? PickHostNameFromBackendAddress { get; set; }

        /// <summary>
        /// Port
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Probe resource of an application gateway.
        /// </summary>
        [Input("probe")]
        public Input<Inputs.SubResourceArgs>? Probe { get; set; }

        /// <summary>
        /// Whether the probe is enabled. Default value is false.
        /// </summary>
        [Input("probeEnabled")]
        public Input<bool>? ProbeEnabled { get; set; }

        /// <summary>
        /// Protocol.
        /// </summary>
        [Input("protocol")]
        public InputUnion<string, Pulumi.AzureNextGen.Network.V20170601.ApplicationGatewayProtocol>? Protocol { get; set; }

        /// <summary>
        /// Provisioning state of the backend http settings resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        /// <summary>
        /// Request timeout in seconds. Application Gateway will fail the request if response is not received within RequestTimeout. Acceptable values are from 1 second to 86400 seconds.
        /// </summary>
        [Input("requestTimeout")]
        public Input<int>? RequestTimeout { get; set; }

        /// <summary>
        /// Type of the resource.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public ApplicationGatewayBackendHttpSettingsArgs()
        {
        }
    }
}
