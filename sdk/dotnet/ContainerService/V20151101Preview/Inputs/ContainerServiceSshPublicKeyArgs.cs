// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ContainerService.V20151101Preview.Inputs
{

    /// <summary>
    /// Contains information about SSH certificate public key data.
    /// </summary>
    public sealed class ContainerServiceSshPublicKeyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gets or sets Certificate public key used to authenticate with VM through SSH. The certificate must be in Pem format with or without headers.
        /// </summary>
        [Input("keyData", required: true)]
        public Input<string> KeyData { get; set; } = null!;

        public ContainerServiceSshPublicKeyArgs()
        {
        }
    }
}
