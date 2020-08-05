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
    /// Http listener of an application gateway.
    /// </summary>
    public sealed class ApplicationGatewayHttpListenerArgs : Pulumi.ResourceArgs
    {
        [Input("customErrorConfigurations")]
        private InputList<Inputs.ApplicationGatewayCustomErrorArgs>? _customErrorConfigurations;

        /// <summary>
        /// Custom error configurations of the HTTP listener.
        /// </summary>
        public InputList<Inputs.ApplicationGatewayCustomErrorArgs> CustomErrorConfigurations
        {
            get => _customErrorConfigurations ?? (_customErrorConfigurations = new InputList<Inputs.ApplicationGatewayCustomErrorArgs>());
            set => _customErrorConfigurations = value;
        }

        /// <summary>
        /// Reference to the FirewallPolicy resource.
        /// </summary>
        [Input("firewallPolicy")]
        public Input<Inputs.SubResourceArgs>? FirewallPolicy { get; set; }

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

        [Input("hostNames")]
        private InputList<string>? _hostNames;

        /// <summary>
        /// List of Host names for HTTP Listener that allows special wildcard characters as well.
        /// </summary>
        public InputList<string> HostNames
        {
            get => _hostNames ?? (_hostNames = new InputList<string>());
            set => _hostNames = value;
        }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Name of the HTTP listener that is unique within an Application Gateway.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Protocol of the HTTP listener.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

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

        public ApplicationGatewayHttpListenerArgs()
        {
        }
    }
}
