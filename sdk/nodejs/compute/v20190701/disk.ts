// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Disk resource.
 *
 * ## Example Usage
 * ### Create a managed disk by copying a snapshot.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const disk = new azurerm.compute.v20190701.Disk("disk", {
 *     creationData: {
 *         createOption: "Copy",
 *         sourceResourceId: "subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot",
 *     },
 *     diskName: "myDisk",
 *     location: "West US",
 *     resourceGroupName: "myResourceGroup",
 * });
 *
 * ```
 * ### Create a managed disk by importing an unmanaged blob from a different subscription.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const disk = new azurerm.compute.v20190701.Disk("disk", {
 *     creationData: {
 *         createOption: "Import",
 *         sourceUri: "https://mystorageaccount.blob.core.windows.net/osimages/osimage.vhd",
 *         storageAccountId: "subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/myStorageAccount",
 *     },
 *     diskName: "myDisk",
 *     location: "West US",
 *     resourceGroupName: "myResourceGroup",
 * });
 *
 * ```
 * ### Create a managed disk by importing an unmanaged blob from the same subscription.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const disk = new azurerm.compute.v20190701.Disk("disk", {
 *     creationData: {
 *         createOption: "Import",
 *         sourceUri: "https://mystorageaccount.blob.core.windows.net/osimages/osimage.vhd",
 *     },
 *     diskName: "myDisk",
 *     location: "West US",
 *     resourceGroupName: "myResourceGroup",
 * });
 *
 * ```
 * ### Create a managed disk from a platform image.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const disk = new azurerm.compute.v20190701.Disk("disk", {
 *     creationData: {
 *         createOption: "FromImage",
 *         imageReference: {
 *             id: "/Subscriptions/{subscriptionId}/Providers/Microsoft.Compute/Locations/uswest/Publishers/Microsoft/ArtifactTypes/VMImage/Offers/{offer}",
 *         },
 *     },
 *     diskName: "myDisk",
 *     location: "West US",
 *     osType: "Windows",
 *     resourceGroupName: "myResourceGroup",
 * });
 *
 * ```
 * ### Create a managed disk from an existing managed disk in the same or different subscription.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const disk = new azurerm.compute.v20190701.Disk("disk", {
 *     creationData: {
 *         createOption: "Copy",
 *         sourceResourceId: "subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myDisk1",
 *     },
 *     diskName: "myDisk2",
 *     location: "West US",
 *     resourceGroupName: "myResourceGroup",
 * });
 *
 * ```
 * ### Create a managed upload disk.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const disk = new azurerm.compute.v20190701.Disk("disk", {
 *     creationData: {
 *         createOption: "Upload",
 *         uploadSizeBytes: 10737418752,
 *     },
 *     diskName: "myDisk",
 *     location: "West US",
 *     resourceGroupName: "myResourceGroup",
 * });
 *
 * ```
 * ### Create an empty managed disk.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const disk = new azurerm.compute.v20190701.Disk("disk", {
 *     creationData: {
 *         createOption: "Empty",
 *     },
 *     diskName: "myDisk",
 *     diskSizeGB: 200,
 *     location: "West US",
 *     resourceGroupName: "myResourceGroup",
 * });
 *
 * ```
 */
export class Disk extends pulumi.CustomResource {
    /**
     * Get an existing Disk resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Disk {
        return new Disk(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:compute/v20190701:Disk';

    /**
     * Returns true if the given object is an instance of Disk.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Disk {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Disk.__pulumiType;
    }

    /**
     * Disk source information. CreationData information cannot be changed after the disk has been created.
     */
    public readonly creationData!: pulumi.Output<outputs.compute.v20190701.CreationDataResponse>;
    /**
     * The number of IOPS allowed for this disk; only settable for UltraSSD disks. One operation can transfer between 4k and 256k bytes.
     */
    public readonly diskIOPSReadWrite!: pulumi.Output<number | undefined>;
    /**
     * The bandwidth allowed for this disk; only settable for UltraSSD disks. MBps means millions of bytes per second - MB here uses the ISO notation, of powers of 10.
     */
    public readonly diskMBpsReadWrite!: pulumi.Output<number | undefined>;
    /**
     * The size of the disk in bytes. This field is read only.
     */
    public /*out*/ readonly diskSizeBytes!: pulumi.Output<number>;
    /**
     * If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
     */
    public readonly diskSizeGB!: pulumi.Output<number | undefined>;
    /**
     * The state of the disk.
     */
    public /*out*/ readonly diskState!: pulumi.Output<string>;
    /**
     * Encryption property can be used to encrypt data at rest with customer managed keys or platform managed keys.
     */
    public readonly encryption!: pulumi.Output<outputs.compute.v20190701.EncryptionResponse | undefined>;
    /**
     * Encryption settings collection used for Azure Disk Encryption, can contain multiple encryption settings per disk or snapshot.
     */
    public readonly encryptionSettingsCollection!: pulumi.Output<outputs.compute.v20190701.EncryptionSettingsCollectionResponse | undefined>;
    /**
     * The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
     */
    public readonly hyperVGeneration!: pulumi.Output<string | undefined>;
    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * A relative URI containing the ID of the VM that has the disk attached.
     */
    public /*out*/ readonly managedBy!: pulumi.Output<string>;
    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The Operating System type.
     */
    public readonly osType!: pulumi.Output<string | undefined>;
    /**
     * The disk provisioning state.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The disks sku name. Can be Standard_LRS, Premium_LRS, StandardSSD_LRS, or UltraSSD_LRS.
     */
    public readonly sku!: pulumi.Output<outputs.compute.v20190701.DiskSkuResponse | undefined>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The time when the disk was created.
     */
    public /*out*/ readonly timeCreated!: pulumi.Output<string>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Unique Guid identifying the resource.
     */
    public /*out*/ readonly uniqueId!: pulumi.Output<string>;
    /**
     * The Logical zone list for Disk.
     */
    public readonly zones!: pulumi.Output<string[] | undefined>;

    /**
     * Create a Disk resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DiskArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.creationData === undefined) {
                throw new Error("Missing required property 'creationData'");
            }
            if (!args || args.diskName === undefined) {
                throw new Error("Missing required property 'diskName'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["creationData"] = args ? args.creationData : undefined;
            inputs["diskIOPSReadWrite"] = args ? args.diskIOPSReadWrite : undefined;
            inputs["diskMBpsReadWrite"] = args ? args.diskMBpsReadWrite : undefined;
            inputs["diskName"] = args ? args.diskName : undefined;
            inputs["diskSizeGB"] = args ? args.diskSizeGB : undefined;
            inputs["encryption"] = args ? args.encryption : undefined;
            inputs["encryptionSettingsCollection"] = args ? args.encryptionSettingsCollection : undefined;
            inputs["hyperVGeneration"] = args ? args.hyperVGeneration : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["osType"] = args ? args.osType : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["zones"] = args ? args.zones : undefined;
            inputs["diskSizeBytes"] = undefined /*out*/;
            inputs["diskState"] = undefined /*out*/;
            inputs["managedBy"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["timeCreated"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["uniqueId"] = undefined /*out*/;
        } else {
            inputs["creationData"] = undefined /*out*/;
            inputs["diskIOPSReadWrite"] = undefined /*out*/;
            inputs["diskMBpsReadWrite"] = undefined /*out*/;
            inputs["diskSizeBytes"] = undefined /*out*/;
            inputs["diskSizeGB"] = undefined /*out*/;
            inputs["diskState"] = undefined /*out*/;
            inputs["encryption"] = undefined /*out*/;
            inputs["encryptionSettingsCollection"] = undefined /*out*/;
            inputs["hyperVGeneration"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["managedBy"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["osType"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["sku"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["timeCreated"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["uniqueId"] = undefined /*out*/;
            inputs["zones"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:compute/latest:Disk" }, { type: "azurerm:compute/v20160430preview:Disk" }, { type: "azurerm:compute/v20170330:Disk" }, { type: "azurerm:compute/v20180401:Disk" }, { type: "azurerm:compute/v20180601:Disk" }, { type: "azurerm:compute/v20180930:Disk" }, { type: "azurerm:compute/v20190301:Disk" }, { type: "azurerm:compute/v20191101:Disk" }, { type: "azurerm:compute/v20200501:Disk" }, { type: "azurerm:compute/v20200630:Disk" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Disk.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Disk resource.
 */
export interface DiskArgs {
    /**
     * Disk source information. CreationData information cannot be changed after the disk has been created.
     */
    readonly creationData: pulumi.Input<inputs.compute.v20190701.CreationData>;
    /**
     * The number of IOPS allowed for this disk; only settable for UltraSSD disks. One operation can transfer between 4k and 256k bytes.
     */
    readonly diskIOPSReadWrite?: pulumi.Input<number>;
    /**
     * The bandwidth allowed for this disk; only settable for UltraSSD disks. MBps means millions of bytes per second - MB here uses the ISO notation, of powers of 10.
     */
    readonly diskMBpsReadWrite?: pulumi.Input<number>;
    /**
     * The name of the managed disk that is being created. The name can't be changed after the disk is created. Supported characters for the name are a-z, A-Z, 0-9 and _. The maximum name length is 80 characters.
     */
    readonly diskName: pulumi.Input<string>;
    /**
     * If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
     */
    readonly diskSizeGB?: pulumi.Input<number>;
    /**
     * Encryption property can be used to encrypt data at rest with customer managed keys or platform managed keys.
     */
    readonly encryption?: pulumi.Input<inputs.compute.v20190701.Encryption>;
    /**
     * Encryption settings collection used for Azure Disk Encryption, can contain multiple encryption settings per disk or snapshot.
     */
    readonly encryptionSettingsCollection?: pulumi.Input<inputs.compute.v20190701.EncryptionSettingsCollection>;
    /**
     * The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
     */
    readonly hyperVGeneration?: pulumi.Input<string>;
    /**
     * Resource location
     */
    readonly location: pulumi.Input<string>;
    /**
     * The Operating System type.
     */
    readonly osType?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The disks sku name. Can be Standard_LRS, Premium_LRS, StandardSSD_LRS, or UltraSSD_LRS.
     */
    readonly sku?: pulumi.Input<inputs.compute.v20190701.DiskSku>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The Logical zone list for Disk.
     */
    readonly zones?: pulumi.Input<pulumi.Input<string>[]>;
}
