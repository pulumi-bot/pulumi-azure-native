// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Web.V20180201.Inputs
{

    /// <summary>
    /// Azure Files or Blob Storage access information value for dictionary storage.
    /// </summary>
    public sealed class AzureStorageInfoValueArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Access key for the storage account.
        /// </summary>
        [Input("accessKey")]
        public Input<string>? AccessKey { get; set; }

        /// <summary>
        /// Name of the storage account.
        /// </summary>
        [Input("accountName")]
        public Input<string>? AccountName { get; set; }

        /// <summary>
        /// Path to mount the storage within the site's runtime environment.
        /// </summary>
        [Input("mountPath")]
        public Input<string>? MountPath { get; set; }

        /// <summary>
        /// Name of the file share (container name, for Blob storage).
        /// </summary>
        [Input("shareName")]
        public Input<string>? ShareName { get; set; }

        /// <summary>
        /// Type of storage.
        /// </summary>
        [Input("type")]
        public Input<Pulumi.AzureNextGen.Web.V20180201.AzureStorageType>? Type { get; set; }

        public AzureStorageInfoValueArgs()
        {
        }
    }
}
