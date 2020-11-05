// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.NetApp.V20191001
{
    /// <summary>
    /// Snapshot of a Volume
    /// </summary>
    public partial class Snapshot : Pulumi.CustomResource
    {
        /// <summary>
        /// The creation date of the snapshot
        /// </summary>
        [Output("created")]
        public Output<string> Created { get; private set; } = null!;

        /// <summary>
        /// UUID v4 used to identify the FileSystem
        /// </summary>
        [Output("fileSystemId")]
        public Output<string?> FileSystemId { get; private set; } = null!;

        /// <summary>
        /// Resource location
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Azure lifecycle management
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// UUID v4 used to identify the Snapshot
        /// </summary>
        [Output("snapshotId")]
        public Output<string> SnapshotId { get; private set; } = null!;

        /// <summary>
        /// Resource tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Resource type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Snapshot resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Snapshot(string name, SnapshotArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:netapp/v20191001:Snapshot", name, args ?? new SnapshotArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Snapshot(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:netapp/v20191001:Snapshot", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:netapp/latest:Snapshot"},
                    new Pulumi.Alias { Type = "azure-nextgen:netapp/v20170815:Snapshot"},
                    new Pulumi.Alias { Type = "azure-nextgen:netapp/v20190501:Snapshot"},
                    new Pulumi.Alias { Type = "azure-nextgen:netapp/v20190601:Snapshot"},
                    new Pulumi.Alias { Type = "azure-nextgen:netapp/v20190701:Snapshot"},
                    new Pulumi.Alias { Type = "azure-nextgen:netapp/v20190801:Snapshot"},
                    new Pulumi.Alias { Type = "azure-nextgen:netapp/v20191101:Snapshot"},
                    new Pulumi.Alias { Type = "azure-nextgen:netapp/v20200201:Snapshot"},
                    new Pulumi.Alias { Type = "azure-nextgen:netapp/v20200301:Snapshot"},
                    new Pulumi.Alias { Type = "azure-nextgen:netapp/v20200501:Snapshot"},
                    new Pulumi.Alias { Type = "azure-nextgen:netapp/v20200601:Snapshot"},
                    new Pulumi.Alias { Type = "azure-nextgen:netapp/v20200701:Snapshot"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Snapshot resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Snapshot Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Snapshot(name, id, options);
        }
    }

    public sealed class SnapshotArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the NetApp account
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// UUID v4 used to identify the FileSystem
        /// </summary>
        [Input("fileSystemId")]
        public Input<string>? FileSystemId { get; set; }

        /// <summary>
        /// Resource location
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The name of the capacity pool
        /// </summary>
        [Input("poolName", required: true)]
        public Input<string> PoolName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the mount target
        /// </summary>
        [Input("snapshotName", required: true)]
        public Input<string> SnapshotName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The name of the volume
        /// </summary>
        [Input("volumeName", required: true)]
        public Input<string> VolumeName { get; set; } = null!;

        public SnapshotArgs()
        {
        }
    }
}
