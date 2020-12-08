// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20171001.Inputs
{

    /// <summary>
    /// Http listener of an application gateway.
    /// </summary>
    public sealed class ApplicationGatewayHttpListenerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Frontend IP configuration resource of an application gateway.
        /// </summary>
        [Input("frontendIPConfiguration")]
        public Input<Inputs.SubResourceArgs>? FrontendIPConfiguration { get; set; }

        /// <summary>
        /// Frontend port resource of an application gateway.
        /// </summary>
        [Input("frontendPort")]
        public Input<Inputs.SubResourceArgs>? FrontendPort { get; set; }

        /// <summary>
        /// Host name of HTTP listener.
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
        /// Protocol.
        /// </summary>
        [Input("protocol")]
        public InputUnion<string, Pulumi.AzureNextGen.Network.V20171001.ApplicationGatewayProtocol>? Protocol { get; set; }

        /// <summary>
        /// Provisioning state of the HTTP listener resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        /// <summary>
        /// Applicable only if protocol is https. Enables SNI for multi-hosting.
        /// </summary>
        [Input("requireServerNameIndication")]
        public Input<bool>? RequireServerNameIndication { get; set; }

        /// <summary>
        /// SSL certificate resource of an application gateway.
        /// </summary>
        [Input("sslCertificate")]
        public Input<Inputs.SubResourceArgs>? SslCertificate { get; set; }

        /// <summary>
        /// Type of the resource.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public ApplicationGatewayHttpListenerArgs()
        {
        }
    }
}
