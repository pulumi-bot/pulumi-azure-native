// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20170301
{
    /// <summary>
    /// Information about packet capture session.
    /// </summary>
    public partial class PacketCapture : Pulumi.CustomResource
    {
        [Output("etag")]
        public Output<string?> Etag { get; private set; } = null!;

        /// <summary>
        /// Name of the packet capture session.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Describes the properties of a packet capture session.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.PacketCaptureResultPropertiesResponseResult> Properties { get; private set; } = null!;


        /// <summary>
        /// Create a PacketCapture resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PacketCapture(string name, PacketCaptureArgs args, CustomResourceOptions? options = null)
            : base("azurerm:network/v20170301:PacketCapture", name, args ?? new PacketCaptureArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PacketCapture(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:network/v20170301:PacketCapture", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
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
        public InputList<Inputs.PacketCaptureFilterArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.PacketCaptureFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// The name of the packet capture session.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the network watcher.
        /// </summary>
        [Input("networkWatcherName", required: true)]
        public Input<string> NetworkWatcherName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Describes the storage location for a packet capture session.
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
