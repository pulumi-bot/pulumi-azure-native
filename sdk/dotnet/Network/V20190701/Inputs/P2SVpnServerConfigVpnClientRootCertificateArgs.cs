// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20190701.Inputs
{

    /// <summary>
    /// VPN client root certificate of P2SVpnServerConfiguration.
    /// </summary>
    public sealed class P2SVpnServerConfigVpnClientRootCertificateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The certificate public data.
        /// </summary>
        [Input("publicCertData", required: true)]
        public Input<string> PublicCertData { get; set; } = null!;

        public P2SVpnServerConfigVpnClientRootCertificateArgs()
        {
        }
    }
}
