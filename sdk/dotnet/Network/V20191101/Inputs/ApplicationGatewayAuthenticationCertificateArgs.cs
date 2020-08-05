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
    /// Authentication certificates of an application gateway.
    /// </summary>
    public sealed class ApplicationGatewayAuthenticationCertificateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Certificate public data.
        /// </summary>
        [Input("data")]
        public Input<string>? Data { get; set; }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Name of the authentication certificate that is unique within an Application Gateway.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ApplicationGatewayAuthenticationCertificateArgs()
        {
        }
    }
}
