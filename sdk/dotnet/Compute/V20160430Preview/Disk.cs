// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Compute.V20160430Preview
{
    /// <summary>
    /// Disk resource.
    /// </summary>
    [AzureNextGenResourceType("azure-nextgen:compute/v20160430preview:Disk")]
    public partial class Disk : Pulumi.CustomResource
    {
        /// <summary>
        /// the storage account type of the disk.
        /// </summary>
        [Output("accountType")]
        public Output<string?> AccountType { get; private set; } = null!;

        /// <summary>
        /// Disk source information. CreationData information cannot be changed after the disk has been created.
        /// </summary>
        [Output("creationData")]
        public Output<Outputs.CreationDataResponse> CreationData { get; private set; } = null!;

        /// <summary>
        /// If creationData.createOption is Empty, this field is mandatory and it indicates the size of the VHD to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
        /// </summary>
        [Output("diskSizeGB")]
        public Output<int?> DiskSizeGB { get; private set; } = null!;

        /// <summary>
        /// Encryption settings for disk or snapshot
        /// </summary>
        [Output("encryptionSettings")]
        public Output<Outputs.EncryptionSettingsResponse?> EncryptionSettings { get; private set; } = null!;

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
        /// The Operating System type.
        /// </summary>
        [Output("osType")]
        public Output<string?> OsType { get; private set; } = null!;

        /// <summary>
        /// A relative URI containing the VM id that has the disk attached.
        /// </summary>
        [Output("ownerId")]
        public Output<string> OwnerId { get; private set; } = null!;

        /// <summary>
        /// The disk provisioning state.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

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
        /// Create a Disk resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Disk(string name, DiskArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:compute/v20160430preview:Disk", name, args ?? new DiskArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Disk(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:compute/v20160430preview:Disk", name, null, MakeResourceOptions(options, id))
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
                    new Pulumi.Alias { Type = "azure-nextgen:compute/v20170330:Disk"},
                    new Pulumi.Alias { Type = "azure-nextgen:compute/v20180401:Disk"},
                    new Pulumi.Alias { Type = "azure-nextgen:compute/v20180601:Disk"},
                    new Pulumi.Alias { Type = "azure-nextgen:compute/v20180930:Disk"},
                    new Pulumi.Alias { Type = "azure-nextgen:compute/v20190301:Disk"},
                    new Pulumi.Alias { Type = "azure-nextgen:compute/v20190701:Disk"},
                    new Pulumi.Alias { Type = "azure-nextgen:compute/v20191101:Disk"},
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
        /// the storage account type of the disk.
        /// </summary>
        [Input("accountType")]
        public Input<Pulumi.AzureNextGen.Compute.V20160430Preview.StorageAccountTypes>? AccountType { get; set; }

        /// <summary>
        /// Disk source information. CreationData information cannot be changed after the disk has been created.
        /// </summary>
        [Input("creationData", required: true)]
        public Input<Inputs.CreationDataArgs> CreationData { get; set; } = null!;

        /// <summary>
        /// The name of the managed disk that is being created. The name can't be changed after the disk is created. Supported characters for the name are a-z, A-Z, 0-9 and _. The maximum name length is 80 characters.
        /// </summary>
        [Input("diskName", required: true)]
        public Input<string> DiskName { get; set; } = null!;

        /// <summary>
        /// If creationData.createOption is Empty, this field is mandatory and it indicates the size of the VHD to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
        /// </summary>
        [Input("diskSizeGB")]
        public Input<int>? DiskSizeGB { get; set; }

        /// <summary>
        /// Encryption settings for disk or snapshot
        /// </summary>
        [Input("encryptionSettings")]
        public Input<Inputs.EncryptionSettingsArgs>? EncryptionSettings { get; set; }

        /// <summary>
        /// Resource location
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The Operating System type.
        /// </summary>
        [Input("osType")]
        public Input<Pulumi.AzureNextGen.Compute.V20160430Preview.OperatingSystemTypes>? OsType { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

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

        public DiskArgs()
        {
        }
    }
}
