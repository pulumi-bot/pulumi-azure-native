// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20190901.Inputs
{

    /// <summary>
    /// SSL certificates of an application gateway.
    /// </summary>
    public sealed class ApplicationGatewaySslCertificateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Base-64 encoded pfx certificate. Only applicable in PUT Request.
        /// </summary>
        [Input("data")]
        public Input<string>? Data { get; set; }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Secret Id of (base-64 encoded unencrypted pfx) 'Secret' or 'Certificate' object stored in KeyVault.
        /// </summary>
        [Input("keyVaultSecretId")]
        public Input<string>? KeyVaultSecretId { get; set; }

        /// <summary>
        /// Name of the SSL certificate that is unique within an Application Gateway.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Password for the pfx file specified in data. Only applicable in PUT request.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        public ApplicationGatewaySslCertificateArgs()
        {
        }
    }
}
