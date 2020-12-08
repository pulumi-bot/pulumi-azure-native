// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Compute.V20170330.Inputs
{

    /// <summary>
    /// Describes Protocol and thumbprint of Windows Remote Management listener
    /// </summary>
    public sealed class WinRMListenerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// This is the URL of a certificate that has been uploaded to Key Vault as a secret. For adding a secret to the Key Vault, see [Add a key or secret to the key vault](https://docs.microsoft.com/azure/key-vault/key-vault-get-started/#add). In this case, your certificate needs to be It is the Base64 encoding of the following JSON Object which is encoded in UTF-8: &lt;br&gt;&lt;br&gt; {&lt;br&gt;  "data":"&lt;Base64-encoded-certificate&gt;",&lt;br&gt;  "dataType":"pfx",&lt;br&gt;  "password":"&lt;pfx-file-password&gt;"&lt;br&gt;}
        /// </summary>
        [Input("certificateUrl")]
        public Input<string>? CertificateUrl { get; set; }

        /// <summary>
        /// Specifies the protocol of listener. &lt;br&gt;&lt;br&gt; Possible values are: &lt;br&gt;**http** &lt;br&gt;&lt;br&gt; **https**
        /// </summary>
        [Input("protocol")]
        public Input<Pulumi.AzureNextGen.Compute.V20170330.ProtocolTypes>? Protocol { get; set; }

        public WinRMListenerArgs()
        {
        }
    }
}
