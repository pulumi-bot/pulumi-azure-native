// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ImportExport.Outputs
{

    [OutputType]
    public sealed class ExportResponse
    {
        /// <summary>
        /// A list of the blobs to be exported.
        /// </summary>
        public readonly Outputs.ExportResponseProperties? BlobList;
        /// <summary>
        /// The relative URI to the block blob that contains the list of blob paths or blob path prefixes as defined above, beginning with the container name. If the blob is in root container, the URI must begin with $root. 
        /// </summary>
        public readonly string? BlobListblobPath;

        [OutputConstructor]
        private ExportResponse(
            Outputs.ExportResponseProperties? blobList,

            string? blobListblobPath)
        {
            BlobList = blobList;
            BlobListblobPath = blobListblobPath;
        }
    }
}