// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20160601.Inputs
{

    /// <summary>
    /// Backend address pool settings of application gateway
    /// </summary>
    public sealed class ApplicationGatewayBackendHttpSettingsArgs : Pulumi.ResourceArgs
    {
        [Input("authenticationCertificates")]
        private InputList<Inputs.SubResourceArgs>? _authenticationCertificates;

        /// <summary>
        /// Array of references to Application Gateway Authentication Certificates
        /// </summary>
        public InputList<Inputs.SubResourceArgs> AuthenticationCertificates
        {
            get => _authenticationCertificates ?? (_authenticationCertificates = new InputList<Inputs.SubResourceArgs>());
            set => _authenticationCertificates = value;
        }

        /// <summary>
        /// Cookie affinity
        /// </summary>
        [Input("cookieBasedAffinity")]
        public InputUnion<string, Pulumi.AzureNextGen.Network.V20160601.ApplicationGatewayCookieBasedAffinity>? CookieBasedAffinity { get; set; }

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Resource Id
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Name of the resource that is unique within a resource group. This name can be used to access the resource
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Port
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Probe resource of application gateway 
        /// </summary>
        [Input("probe")]
        public Input<Inputs.SubResourceArgs>? Probe { get; set; }

        /// <summary>
        /// Protocol
        /// </summary>
        [Input("protocol")]
        public InputUnion<string, Pulumi.AzureNextGen.Network.V20160601.ApplicationGatewayProtocol>? Protocol { get; set; }

        /// <summary>
        /// Provisioning state of the backend http settings resource Updating/Deleting/Failed
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        /// <summary>
        /// Request timeout
        /// </summary>
        [Input("requestTimeout")]
        public Input<int>? RequestTimeout { get; set; }

        public ApplicationGatewayBackendHttpSettingsArgs()
        {
        }
    }
}
