// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Devices.V20200831Preview.Inputs
{

    /// <summary>
    /// The description of an X509 CA Certificate.
    /// </summary>
    public sealed class CertificatePropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The certificate content
        /// </summary>
        [Input("certificate")]
        public Input<string>? Certificate { get; set; }

        public CertificatePropertiesArgs()
        {
        }
    }
}
