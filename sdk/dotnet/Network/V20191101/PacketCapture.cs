// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20191101
{
    /// <summary>
    /// Information about packet capture session.
    /// </summary>
    [AzureNextGenResourceType("azure-nextgen:network/v20191101:PacketCapture")]
    public partial class PacketCapture : Pulumi.CustomResource
    {
        /// <summary>
        /// Number of bytes captured per packet, the remaining bytes are truncated.
        /// </summary>
        [Output("bytesToCapturePerPacket")]
        public Output<int?> BytesToCapturePerPacket { get; private set; } = null!;

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// A list of packet capture filters.
        /// </summary>
        [Output("filters")]
        public Output<ImmutableArray<Outputs.PacketCaptureFilterResponse>> Filters { get; private set; } = null!;

        /// <summary>
        /// Name of the packet capture session.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the packet capture session.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The storage location for a packet capture session.
        /// </summary>
        [Output("storageLocation")]
        public Output<Outputs.PacketCaptureStorageLocationResponse> StorageLocation { get; private set; } = null!;

        /// <summary>
        /// The ID of the targeted resource, only VM is currently supported.
        /// </summary>
        [Output("target")]
        public Output<string> Target { get; private set; } = null!;

        /// <summary>
        /// Maximum duration of the capture session in seconds.
        /// </summary>
        [Output("timeLimitInSeconds")]
        public Output<int?> TimeLimitInSeconds { get; private set; } = null!;

        /// <summary>
        /// Maximum size of the capture output.
        /// </summary>
        [Output("totalBytesPerSession")]
        public Output<int?> TotalBytesPerSession { get; private set; } = null!;


        /// <summary>
        /// Create a PacketCapture resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PacketCapture(string name, PacketCaptureArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:network/v20191101:PacketCapture", name, args ?? new PacketCaptureArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PacketCapture(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:network/v20191101:PacketCapture", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:network/latest:PacketCapture"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20160901:PacketCapture"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20161201:PacketCapture"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20170301:PacketCapture"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20170601:PacketCapture"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20170801:PacketCapture"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20170901:PacketCapture"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20171001:PacketCapture"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20171101:PacketCapture"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20180101:PacketCapture"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20180201:PacketCapture"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20180401:PacketCapture"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20180601:PacketCapture"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20180701:PacketCapture"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20180801:PacketCapture"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20181001:PacketCapture"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20181101:PacketCapture"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20181201:PacketCapture"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190201:PacketCapture"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190401:PacketCapture"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190601:PacketCapture"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190701:PacketCapture"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190801:PacketCapture"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190901:PacketCapture"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20191201:PacketCapture"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200301:PacketCapture"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200401:PacketCapture"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200501:PacketCapture"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200601:PacketCapture"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200701:PacketCapture"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200801:PacketCapture"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PacketCapture resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PacketCapture Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PacketCapture(name, id, options);
        }
    }

    public sealed class PacketCaptureArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Number of bytes captured per packet, the remaining bytes are truncated.
        /// </summary>
        [Input("bytesToCapturePerPacket")]
        public Input<int>? BytesToCapturePerPacket { get; set; }

        [Input("filters")]
        private InputList<Inputs.PacketCaptureFilterArgs>? _filters;

        /// <summary>
        /// A list of packet capture filters.
        /// </summary>
        public InputList<Inputs.PacketCaptureFilterArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.PacketCaptureFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// The name of the network watcher.
        /// </summary>
        [Input("networkWatcherName", required: true)]
        public Input<string> NetworkWatcherName { get; set; } = null!;

        /// <summary>
        /// The name of the packet capture session.
        /// </summary>
        [Input("packetCaptureName", required: true)]
        public Input<string> PacketCaptureName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The storage location for a packet capture session.
        /// </summary>
        [Input("storageLocation", required: true)]
        public Input<Inputs.PacketCaptureStorageLocationArgs> StorageLocation { get; set; } = null!;

        /// <summary>
        /// The ID of the targeted resource, only VM is currently supported.
        /// </summary>
        [Input("target", required: true)]
        public Input<string> Target { get; set; } = null!;

        /// <summary>
        /// Maximum duration of the capture session in seconds.
        /// </summary>
        [Input("timeLimitInSeconds")]
        public Input<int>? TimeLimitInSeconds { get; set; }

        /// <summary>
        /// Maximum size of the capture output.
        /// </summary>
        [Input("totalBytesPerSession")]
        public Input<int>? TotalBytesPerSession { get; set; }

        public PacketCaptureArgs()
        {
        }
    }
}
