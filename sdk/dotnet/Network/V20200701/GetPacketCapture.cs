// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20200701
{
    public static class GetPacketCapture
    {
        public static Task<GetPacketCaptureResult> InvokeAsync(GetPacketCaptureArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPacketCaptureResult>("azure-nextgen:network/v20200701:getPacketCapture", args ?? new GetPacketCaptureArgs(), options.WithVersion());
    }


    public sealed class GetPacketCaptureArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the network watcher.
        /// </summary>
        [Input("networkWatcherName", required: true)]
        public string NetworkWatcherName { get; set; } = null!;

        /// <summary>
        /// The name of the packet capture session.
        /// </summary>
        [Input("packetCaptureName", required: true)]
        public string PacketCaptureName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetPacketCaptureArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPacketCaptureResult
    {
        /// <summary>
        /// Number of bytes captured per packet, the remaining bytes are truncated.
        /// </summary>
        public readonly int? BytesToCapturePerPacket;
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// A list of packet capture filters.
        /// </summary>
        public readonly ImmutableArray<Outputs.PacketCaptureFilterResponse> Filters;
        /// <summary>
        /// Name of the packet capture session.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The provisioning state of the packet capture session.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The storage location for a packet capture session.
        /// </summary>
        public readonly Outputs.PacketCaptureStorageLocationResponse StorageLocation;
        /// <summary>
        /// The ID of the targeted resource, only VM is currently supported.
        /// </summary>
        public readonly string Target;
        /// <summary>
        /// Maximum duration of the capture session in seconds.
        /// </summary>
        public readonly int? TimeLimitInSeconds;
        /// <summary>
        /// Maximum size of the capture output.
        /// </summary>
        public readonly int? TotalBytesPerSession;

        [OutputConstructor]
        private GetPacketCaptureResult(
            int? bytesToCapturePerPacket,

            string etag,

            ImmutableArray<Outputs.PacketCaptureFilterResponse> filters,

            string name,

            string provisioningState,

            Outputs.PacketCaptureStorageLocationResponse storageLocation,

            string target,

            int? timeLimitInSeconds,

            int? totalBytesPerSession)
        {
            BytesToCapturePerPacket = bytesToCapturePerPacket;
            Etag = etag;
            Filters = filters;
            Name = name;
            ProvisioningState = provisioningState;
            StorageLocation = storageLocation;
            Target = target;
            TimeLimitInSeconds = timeLimitInSeconds;
            TotalBytesPerSession = totalBytesPerSession;
        }
    }
}
