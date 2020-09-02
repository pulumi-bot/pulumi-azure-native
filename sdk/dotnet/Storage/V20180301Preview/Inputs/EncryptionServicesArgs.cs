// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Storage.V20180301Preview.Inputs
{

    /// <summary>
    /// A list of services that support encryption.
    /// </summary>
    public sealed class EncryptionServicesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The encryption function of the blob storage service.
        /// </summary>
        [Input("blob")]
        public Input<Inputs.EncryptionServiceArgs>? Blob { get; set; }

        /// <summary>
        /// The encryption function of the file storage service.
        /// </summary>
        [Input("file")]
        public Input<Inputs.EncryptionServiceArgs>? File { get; set; }

        public EncryptionServicesArgs()
        {
        }
    }
}
