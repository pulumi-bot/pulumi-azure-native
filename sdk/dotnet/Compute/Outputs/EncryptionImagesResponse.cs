// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.Outputs
{

    [OutputType]
    public sealed class EncryptionImagesResponse
    {
        /// <summary>
        /// A list of encryption specifications for data disk images.
        /// </summary>
        public readonly ImmutableArray<Outputs.DataDiskImageEncryptionResponse> DataDiskImages;
        /// <summary>
        /// Contains encryption settings for an OS disk image.
        /// </summary>
        public readonly Outputs.OSDiskImageEncryptionResponse? OsDiskImage;

        [OutputConstructor]
        private EncryptionImagesResponse(
            ImmutableArray<Outputs.DataDiskImageEncryptionResponse> dataDiskImages,

            Outputs.OSDiskImageEncryptionResponse? osDiskImage)
        {
            DataDiskImages = dataDiskImages;
            OsDiskImage = osDiskImage;
        }
    }
}