// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Compute.V20200930
{
    /// <summary>
    /// Snapshot resource.
    /// </summary>
    [AzureNextGenResourceType("azure-nextgen:compute/v20200930:Snapshot")]
    public partial class Snapshot : Pulumi.CustomResource
    {
        /// <summary>
        /// Disk source information. CreationData information cannot be changed after the disk has been created.
        /// </summary>
        [Output("creationData")]
        public Output<Outputs.CreationDataResponse> CreationData { get; private set; } = null!;

        /// <summary>
        /// ARM id of the DiskAccess resource for using private endpoints on disks.
        /// </summary>
        [Output("diskAccessId")]
        public Output<string?> DiskAccessId { get; private set; } = null!;

        /// <summary>
        /// The size of the disk in bytes. This field is read only.
        /// </summary>
        [Output("diskSizeBytes")]
        public Output<double> DiskSizeBytes { get; private set; } = null!;

        /// <summary>
        /// If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
        /// </summary>
        [Output("diskSizeGB")]
        public Output<int?> DiskSizeGB { get; private set; } = null!;

        /// <summary>
        /// The state of the snapshot.
        /// </summary>
        [Output("diskState")]
        public Output<string> DiskState { get; private set; } = null!;

        /// <summary>
        /// Encryption property can be used to encrypt data at rest with customer managed keys or platform managed keys.
        /// </summary>
        [Output("encryption")]
        public Output<Outputs.EncryptionResponse?> Encryption { get; private set; } = null!;

        /// <summary>
        /// Encryption settings collection used be Azure Disk Encryption, can contain multiple encryption settings per disk or snapshot.
        /// </summary>
        [Output("encryptionSettingsCollection")]
        public Output<Outputs.EncryptionSettingsCollectionResponse?> EncryptionSettingsCollection { get; private set; } = null!;

        /// <summary>
        /// The extended location where the snapshot will be created. Extended location cannot be changed.
        /// </summary>
        [Output("extendedLocation")]
        public Output<Outputs.ExtendedLocationResponse?> ExtendedLocation { get; private set; } = null!;

        /// <summary>
        /// The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
        /// </summary>
        [Output("hyperVGeneration")]
        public Output<string?> HyperVGeneration { get; private set; } = null!;

        /// <summary>
        /// Whether a snapshot is incremental. Incremental snapshots on the same disk occupy less space than full snapshots and can be diffed.
        /// </summary>
        [Output("incremental")]
        public Output<bool?> Incremental { get; private set; } = null!;

        /// <summary>
        /// Resource location
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Unused. Always Null.
        /// </summary>
        [Output("managedBy")]
        public Output<string> ManagedBy { get; private set; } = null!;

        /// <summary>
        /// Resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Policy for accessing the disk via network.
        /// </summary>
        [Output("networkAccessPolicy")]
        public Output<string?> NetworkAccessPolicy { get; private set; } = null!;

        /// <summary>
        /// The Operating System type.
        /// </summary>
        [Output("osType")]
        public Output<string?> OsType { get; private set; } = null!;

        /// <summary>
        /// The disk provisioning state.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Purchase plan information for the image from which the source disk for the snapshot was originally created.
        /// </summary>
        [Output("purchasePlan")]
        public Output<Outputs.PurchasePlanResponse?> PurchasePlan { get; private set; } = null!;

        /// <summary>
        /// The snapshots sku name. Can be Standard_LRS, Premium_LRS, or Standard_ZRS. This is an optional parameter for incremental snapshot and the default behavior is the SKU will be set to the same sku as the previous snapshot
        /// </summary>
        [Output("sku")]
        public Output<Outputs.SnapshotSkuResponse?> Sku { get; private set; } = null!;

        /// <summary>
        /// Resource tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The time when the snapshot was created.
        /// </summary>
        [Output("timeCreated")]
        public Output<string> TimeCreated { get; private set; } = null!;

        /// <summary>
        /// Resource type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Unique Guid identifying the resource.
        /// </summary>
        [Output("uniqueId")]
        public Output<string> UniqueId { get; private set; } = null!;


        /// <summary>
        /// Create a Snapshot resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Snapshot(string name, SnapshotArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:compute/v20200930:Snapshot", name, args ?? new SnapshotArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Snapshot(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:compute/v20200930:Snapshot", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:compute/latest:Snapshot"},
                    new Pulumi.Alias { Type = "azure-nextgen:compute/v20160430preview:Snapshot"},
                    new Pulumi.Alias { Type = "azure-nextgen:compute/v20170330:Snapshot"},
                    new Pulumi.Alias { Type = "azure-nextgen:compute/v20180401:Snapshot"},
                    new Pulumi.Alias { Type = "azure-nextgen:compute/v20180601:Snapshot"},
                    new Pulumi.Alias { Type = "azure-nextgen:compute/v20180930:Snapshot"},
                    new Pulumi.Alias { Type = "azure-nextgen:compute/v20190301:Snapshot"},
                    new Pulumi.Alias { Type = "azure-nextgen:compute/v20190701:Snapshot"},
                    new Pulumi.Alias { Type = "azure-nextgen:compute/v20191101:Snapshot"},
                    new Pulumi.Alias { Type = "azure-nextgen:compute/v20200501:Snapshot"},
                    new Pulumi.Alias { Type = "azure-nextgen:compute/v20200630:Snapshot"},
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
        /// Disk source information. CreationData information cannot be changed after the disk has been created.
        /// </summary>
        [Input("creationData", required: true)]
        public Input<Inputs.CreationDataArgs> CreationData { get; set; } = null!;

        /// <summary>
        /// ARM id of the DiskAccess resource for using private endpoints on disks.
        /// </summary>
        [Input("diskAccessId")]
        public Input<string>? DiskAccessId { get; set; }

        /// <summary>
        /// If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
        /// </summary>
        [Input("diskSizeGB")]
        public Input<int>? DiskSizeGB { get; set; }

        /// <summary>
        /// Encryption property can be used to encrypt data at rest with customer managed keys or platform managed keys.
        /// </summary>
        [Input("encryption")]
        public Input<Inputs.EncryptionArgs>? Encryption { get; set; }

        /// <summary>
        /// Encryption settings collection used be Azure Disk Encryption, can contain multiple encryption settings per disk or snapshot.
        /// </summary>
        [Input("encryptionSettingsCollection")]
        public Input<Inputs.EncryptionSettingsCollectionArgs>? EncryptionSettingsCollection { get; set; }

        /// <summary>
        /// The extended location where the snapshot will be created. Extended location cannot be changed.
        /// </summary>
        [Input("extendedLocation")]
        public Input<Inputs.ExtendedLocationArgs>? ExtendedLocation { get; set; }

        /// <summary>
        /// The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
        /// </summary>
        [Input("hyperVGeneration")]
        public InputUnion<string, Pulumi.AzureNextGen.Compute.V20200930.HyperVGeneration>? HyperVGeneration { get; set; }

        /// <summary>
        /// Whether a snapshot is incremental. Incremental snapshots on the same disk occupy less space than full snapshots and can be diffed.
        /// </summary>
        [Input("incremental")]
        public Input<bool>? Incremental { get; set; }

        /// <summary>
        /// Resource location
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Policy for accessing the disk via network.
        /// </summary>
        [Input("networkAccessPolicy")]
        public InputUnion<string, Pulumi.AzureNextGen.Compute.V20200930.NetworkAccessPolicy>? NetworkAccessPolicy { get; set; }

        /// <summary>
        /// The Operating System type.
        /// </summary>
        [Input("osType")]
        public Input<Pulumi.AzureNextGen.Compute.V20200930.OperatingSystemTypes>? OsType { get; set; }

        /// <summary>
        /// Purchase plan information for the image from which the source disk for the snapshot was originally created.
        /// </summary>
        [Input("purchasePlan")]
        public Input<Inputs.PurchasePlanArgs>? PurchasePlan { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The snapshots sku name. Can be Standard_LRS, Premium_LRS, or Standard_ZRS. This is an optional parameter for incremental snapshot and the default behavior is the SKU will be set to the same sku as the previous snapshot
        /// </summary>
        [Input("sku")]
        public Input<Inputs.SnapshotSkuArgs>? Sku { get; set; }

        /// <summary>
        /// The name of the snapshot that is being created. The name can't be changed after the snapshot is created. Supported characters for the name are a-z, A-Z, 0-9 and _. The max name length is 80 characters.
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

        public SnapshotArgs()
        {
        }
    }
}
