// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.BatchAI.V20180301.Outputs
{

    [OutputType]
    public sealed class MountVolumesResponse
    {
        /// <summary>
        /// References to Azure Blob FUSE that are to be mounted to the cluster nodes.
        /// </summary>
        public readonly ImmutableArray<Outputs.AzureBlobFileSystemReferenceResponse> AzureBlobFileSystems;
        /// <summary>
        /// References to Azure File Shares that are to be mounted to the cluster nodes.
        /// </summary>
        public readonly ImmutableArray<Outputs.AzureFileShareReferenceResponse> AzureFileShares;
        public readonly ImmutableArray<Outputs.FileServerReferenceResponse> FileServers;
        public readonly ImmutableArray<Outputs.UnmanagedFileSystemReferenceResponse> UnmanagedFileSystems;

        [OutputConstructor]
        private MountVolumesResponse(
            ImmutableArray<Outputs.AzureBlobFileSystemReferenceResponse> azureBlobFileSystems,

            ImmutableArray<Outputs.AzureFileShareReferenceResponse> azureFileShares,

            ImmutableArray<Outputs.FileServerReferenceResponse> fileServers,

            ImmutableArray<Outputs.UnmanagedFileSystemReferenceResponse> unmanagedFileSystems)
        {
            AzureBlobFileSystems = azureBlobFileSystems;
            AzureFileShares = azureFileShares;
            FileServers = fileServers;
            UnmanagedFileSystems = unmanagedFileSystems;
        }
    }
}
