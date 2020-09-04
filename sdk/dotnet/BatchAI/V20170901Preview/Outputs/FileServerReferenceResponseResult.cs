// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.BatchAI.V20170901Preview.Outputs
{

    [OutputType]
    public sealed class FileServerReferenceResponseResult
    {
        /// <summary>
        /// Represents a resource ID. For example, for a subnet, it is the resource URL for the subnet.
        /// </summary>
        public readonly Outputs.ResourceIdResponseResult FileServer;
        public readonly string? MountOptions;
        /// <summary>
        /// Note that all file shares will be mounted under $AZ_BATCHAI_MOUNT_ROOT location.
        /// </summary>
        public readonly string RelativeMountPath;
        /// <summary>
        /// If this property is not specified, the entire File Server will be mounted.
        /// </summary>
        public readonly string? SourceDirectory;

        [OutputConstructor]
        private FileServerReferenceResponseResult(
            Outputs.ResourceIdResponseResult fileServer,

            string? mountOptions,

            string relativeMountPath,

            string? sourceDirectory)
        {
            FileServer = fileServer;
            MountOptions = mountOptions;
            RelativeMountPath = relativeMountPath;
            SourceDirectory = sourceDirectory;
        }
    }
}
