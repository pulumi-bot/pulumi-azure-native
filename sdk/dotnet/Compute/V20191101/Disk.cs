// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Compute.V20191101
{
    /// <summary>
    /// Disk resource.
    /// </summary>
    [AzureNextGenResourceType("azure-nextgen:compute/v20191101:Disk")]
    public partial class Disk : Pulumi.CustomResource
    {
        /// <summary>
        /// Disk source information. CreationData information cannot be changed after the disk has been created.
        /// </summary>
        [Output("creationData")]
        public Output<Outputs.CreationDataResponse> CreationData { get; private set; } = null!;

        /// <summary>
        /// The total number of IOPS that will be allowed across all VMs mounting the shared disk as ReadOnly. One operation can transfer between 4k and 256k bytes.
        /// </summary>
        [Output("diskIOPSReadOnly")]
        public Output<double?> DiskIOPSReadOnly { get; private set; } = null!;

        /// <summary>
        /// The number of IOPS allowed for this disk; only settable for UltraSSD disks. One operation can transfer between 4k and 256k bytes.
        /// </summary>
        [Output("diskIOPSReadWrite")]
        public Output<double?> DiskIOPSReadWrite { get; private set; } = null!;

        /// <summary>
        /// The total throughput (MBps) that will be allowed across all VMs mounting the shared disk as ReadOnly. MBps means millions of bytes per second - MB here uses the ISO notation, of powers of 10.
        /// </summary>
        [Output("diskMBpsReadOnly")]
        public Output<double?> DiskMBpsReadOnly { get; private set; } = null!;

        /// <summary>
        /// The bandwidth allowed for this disk; only settable for UltraSSD disks. MBps means millions of bytes per second - MB here uses the ISO notation, of powers of 10.
        /// </summary>
        [Output("diskMBpsReadWrite")]
        public Output<double?> DiskMBpsReadWrite { get; private set; } = null!;

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
        /// The state of the disk.
        /// </summary>
        [Output("diskState")]
        public Output<string> DiskState { get; private set; } = null!;

        /// <summary>
        /// Encryption property can be used to encrypt data at rest with customer managed keys or platform managed keys.
        /// </summary>
        [Output("encryption")]
        public Output<Outputs.EncryptionResponse?> Encryption { get; private set; } = null!;

        /// <summary>
        /// Encryption settings collection used for Azure Disk Encryption, can contain multiple encryption settings per disk or snapshot.
        /// </summary>
        [Output("encryptionSettingsCollection")]
        public Output<Outputs.EncryptionSettingsCollectionResponse?> EncryptionSettingsCollection { get; private set; } = null!;

        /// <summary>
        /// The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
        /// </summary>
        [Output("hyperVGeneration")]
        public Output<string?> HyperVGeneration { get; private set; } = null!;

        /// <summary>
        /// Resource location
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// A relative URI containing the ID of the VM that has the disk attached.
        /// </summary>
        [Output("managedBy")]
        public Output<string> ManagedBy { get; private set; } = null!;

        /// <summary>
        /// List of relative URIs containing the IDs of the VMs that have the disk attached. maxShares should be set to a value greater than one for disks to allow attaching them to multiple VMs.
        /// </summary>
        [Output("managedByExtended")]
        public Output<ImmutableArray<string>> ManagedByExtended { get; private set; } = null!;

        /// <summary>
        /// The maximum number of VMs that can attach to the disk at the same time. Value greater than one indicates a disk that can be mounted on multiple VMs at the same time.
        /// </summary>
        [Output("maxShares")]
        public Output<int?> MaxShares { get; private set; } = null!;

        /// <summary>
        /// Resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

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
        /// Details of the list of all VMs that have the disk attached. maxShares should be set to a value greater than one for disks to allow attaching them to multiple VMs.
        /// </summary>
        [Output("shareInfo")]
        public Output<ImmutableArray<Outputs.ShareInfoElementResponse>> ShareInfo { get; private set; } = null!;

        /// <summary>
        /// The disks sku name. Can be Standard_LRS, Premium_LRS, StandardSSD_LRS, or UltraSSD_LRS.
        /// </summary>
        [Output("sku")]
        public Output<Outputs.DiskSkuResponse?> Sku { get; private set; } = null!;

        /// <summary>
        /// Resource tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The time when the disk was created.
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
        /// The Logical zone list for Disk.
        /// </summary>
        [Output("zones")]
        public Output<ImmutableArray<string>> Zones { get; private set; } = null!;


        /// <summary>
        /// Create a Disk resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Disk(string name, DiskArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:compute/v20191101:Disk", name, args ?? new DiskArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Disk(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:compute/v20191101:Disk", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:compute/latest:Disk"},
                    new Pulumi.Alias { Type = "azure-nextgen:compute/v20160430preview:Disk"},
                    new Pulumi.Alias { Type = "azure-nextgen:compute/v20170330:Disk"},
                    new Pulumi.Alias { Type = "azure-nextgen:compute/v20180401:Disk"},
                    new Pulumi.Alias { Type = "azure-nextgen:compute/v20180601:Disk"},
                    new Pulumi.Alias { Type = "azure-nextgen:compute/v20180930:Disk"},
                    new Pulumi.Alias { Type = "azure-nextgen:compute/v20190301:Disk"},
                    new Pulumi.Alias { Type = "azure-nextgen:compute/v20190701:Disk"},
                    new Pulumi.Alias { Type = "azure-nextgen:compute/v20200501:Disk"},
                    new Pulumi.Alias { Type = "azure-nextgen:compute/v20200630:Disk"},
                    new Pulumi.Alias { Type = "azure-nextgen:compute/v20200930:Disk"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Disk resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Disk Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Disk(name, id, options);
        }
    }

    public sealed class DiskArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Disk source information. CreationData information cannot be changed after the disk has been created.
        /// </summary>
        [Input("creationData", required: true)]
        public Input<Inputs.CreationDataArgs> CreationData { get; set; } = null!;

        /// <summary>
        /// The total number of IOPS that will be allowed across all VMs mounting the shared disk as ReadOnly. One operation can transfer between 4k and 256k bytes.
        /// </summary>
        [Input("diskIOPSReadOnly")]
        public Input<double>? DiskIOPSReadOnly { get; set; }

        /// <summary>
        /// The number of IOPS allowed for this disk; only settable for UltraSSD disks. One operation can transfer between 4k and 256k bytes.
        /// </summary>
        [Input("diskIOPSReadWrite")]
        public Input<double>? DiskIOPSReadWrite { get; set; }

        /// <summary>
        /// The total throughput (MBps) that will be allowed across all VMs mounting the shared disk as ReadOnly. MBps means millions of bytes per second - MB here uses the ISO notation, of powers of 10.
        /// </summary>
        [Input("diskMBpsReadOnly")]
        public Input<double>? DiskMBpsReadOnly { get; set; }

        /// <summary>
        /// The bandwidth allowed for this disk; only settable for UltraSSD disks. MBps means millions of bytes per second - MB here uses the ISO notation, of powers of 10.
        /// </summary>
        [Input("diskMBpsReadWrite")]
        public Input<double>? DiskMBpsReadWrite { get; set; }

        /// <summary>
        /// The name of the managed disk that is being created. The name can't be changed after the disk is created. Supported characters for the name are a-z, A-Z, 0-9 and _. The maximum name length is 80 characters.
        /// </summary>
        [Input("diskName", required: true)]
        public Input<string> DiskName { get; set; } = null!;

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
        /// Encryption settings collection used for Azure Disk Encryption, can contain multiple encryption settings per disk or snapshot.
        /// </summary>
        [Input("encryptionSettingsCollection")]
        public Input<Inputs.EncryptionSettingsCollectionArgs>? EncryptionSettingsCollection { get; set; }

        /// <summary>
        /// The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
        /// </summary>
        [Input("hyperVGeneration")]
        public InputUnion<string, Pulumi.AzureNextGen.Compute.V20191101.HyperVGeneration>? HyperVGeneration { get; set; }

        /// <summary>
        /// Resource location
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The maximum number of VMs that can attach to the disk at the same time. Value greater than one indicates a disk that can be mounted on multiple VMs at the same time.
        /// </summary>
        [Input("maxShares")]
        public Input<int>? MaxShares { get; set; }

        /// <summary>
        /// The Operating System type.
        /// </summary>
        [Input("osType")]
        public Input<Pulumi.AzureNextGen.Compute.V20191101.OperatingSystemTypes>? OsType { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The disks sku name. Can be Standard_LRS, Premium_LRS, StandardSSD_LRS, or UltraSSD_LRS.
        /// </summary>
        [Input("sku")]
        public Input<Inputs.DiskSkuArgs>? Sku { get; set; }

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

        [Input("zones")]
        private InputList<string>? _zones;

        /// <summary>
        /// The Logical zone list for Disk.
        /// </summary>
        public InputList<string> Zones
        {
            get => _zones ?? (_zones = new InputList<string>());
            set => _zones = value;
        }

        public DiskArgs()
        {
        }
    }
}
