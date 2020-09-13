// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Disk resource.
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
    public static readonly __pulumiType = 'azurerm:compute/v20200501:Disk';

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
    public readonly creationData!: pulumi.Output<outputs.compute.v20200501.CreationDataResponse>;
    /**
     * ARM id of the DiskAccess resource for using private endpoints on disks.
     */
    public readonly diskAccessId!: pulumi.Output<string | undefined>;
    /**
     * The total number of IOPS that will be allowed across all VMs mounting the shared disk as ReadOnly. One operation can transfer between 4k and 256k bytes.
     */
    public readonly diskIOPSReadOnly!: pulumi.Output<number | undefined>;
    /**
     * The number of IOPS allowed for this disk; only settable for UltraSSD disks. One operation can transfer between 4k and 256k bytes.
     */
    public readonly diskIOPSReadWrite!: pulumi.Output<number | undefined>;
    /**
     * The total throughput (MBps) that will be allowed across all VMs mounting the shared disk as ReadOnly. MBps means millions of bytes per second - MB here uses the ISO notation, of powers of 10.
     */
    public readonly diskMBpsReadOnly!: pulumi.Output<number | undefined>;
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
    public readonly encryption!: pulumi.Output<outputs.compute.v20200501.EncryptionResponse | undefined>;
    /**
     * Encryption settings collection used for Azure Disk Encryption, can contain multiple encryption settings per disk or snapshot.
     */
    public readonly encryptionSettingsCollection!: pulumi.Output<outputs.compute.v20200501.EncryptionSettingsCollectionResponse | undefined>;
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
     * List of relative URIs containing the IDs of the VMs that have the disk attached. maxShares should be set to a value greater than one for disks to allow attaching them to multiple VMs.
     */
    public /*out*/ readonly managedByExtended!: pulumi.Output<string[]>;
    /**
     * The maximum number of VMs that can attach to the disk at the same time. Value greater than one indicates a disk that can be mounted on multiple VMs at the same time.
     */
    public readonly maxShares!: pulumi.Output<number | undefined>;
    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Policy for accessing the disk via network.
     */
    public readonly networkAccessPolicy!: pulumi.Output<string | undefined>;
    /**
     * The Operating System type.
     */
    public readonly osType!: pulumi.Output<string | undefined>;
    /**
     * The disk provisioning state.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Details of the list of all VMs that have the disk attached. maxShares should be set to a value greater than one for disks to allow attaching them to multiple VMs.
     */
    public /*out*/ readonly shareInfo!: pulumi.Output<outputs.compute.v20200501.ShareInfoElementResponse[]>;
    /**
     * The disks sku name. Can be Standard_LRS, Premium_LRS, StandardSSD_LRS, or UltraSSD_LRS.
     */
    public readonly sku!: pulumi.Output<outputs.compute.v20200501.DiskSkuResponse | undefined>;
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
            inputs["diskAccessId"] = args ? args.diskAccessId : undefined;
            inputs["diskIOPSReadOnly"] = args ? args.diskIOPSReadOnly : undefined;
            inputs["diskIOPSReadWrite"] = args ? args.diskIOPSReadWrite : undefined;
            inputs["diskMBpsReadOnly"] = args ? args.diskMBpsReadOnly : undefined;
            inputs["diskMBpsReadWrite"] = args ? args.diskMBpsReadWrite : undefined;
            inputs["diskName"] = args ? args.diskName : undefined;
            inputs["diskSizeGB"] = args ? args.diskSizeGB : undefined;
            inputs["encryption"] = args ? args.encryption : undefined;
            inputs["encryptionSettingsCollection"] = args ? args.encryptionSettingsCollection : undefined;
            inputs["hyperVGeneration"] = args ? args.hyperVGeneration : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["maxShares"] = args ? args.maxShares : undefined;
            inputs["networkAccessPolicy"] = args ? args.networkAccessPolicy : undefined;
            inputs["osType"] = args ? args.osType : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["zones"] = args ? args.zones : undefined;
            inputs["diskSizeBytes"] = undefined /*out*/;
            inputs["diskState"] = undefined /*out*/;
            inputs["managedBy"] = undefined /*out*/;
            inputs["managedByExtended"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["shareInfo"] = undefined /*out*/;
            inputs["timeCreated"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["uniqueId"] = undefined /*out*/;
        } else {
            inputs["creationData"] = undefined /*out*/;
            inputs["diskAccessId"] = undefined /*out*/;
            inputs["diskIOPSReadOnly"] = undefined /*out*/;
            inputs["diskIOPSReadWrite"] = undefined /*out*/;
            inputs["diskMBpsReadOnly"] = undefined /*out*/;
            inputs["diskMBpsReadWrite"] = undefined /*out*/;
            inputs["diskSizeBytes"] = undefined /*out*/;
            inputs["diskSizeGB"] = undefined /*out*/;
            inputs["diskState"] = undefined /*out*/;
            inputs["encryption"] = undefined /*out*/;
            inputs["encryptionSettingsCollection"] = undefined /*out*/;
            inputs["hyperVGeneration"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["managedBy"] = undefined /*out*/;
            inputs["managedByExtended"] = undefined /*out*/;
            inputs["maxShares"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["networkAccessPolicy"] = undefined /*out*/;
            inputs["osType"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["shareInfo"] = undefined /*out*/;
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
        const aliasOpts = { aliases: [{ type: "azurerm:compute/latest:Disk" }, { type: "azurerm:compute/v20160430preview:Disk" }, { type: "azurerm:compute/v20170330:Disk" }, { type: "azurerm:compute/v20180401:Disk" }, { type: "azurerm:compute/v20180601:Disk" }, { type: "azurerm:compute/v20180930:Disk" }, { type: "azurerm:compute/v20190301:Disk" }, { type: "azurerm:compute/v20190701:Disk" }, { type: "azurerm:compute/v20191101:Disk" }, { type: "azurerm:compute/v20200630:Disk" }] };
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
    readonly creationData: pulumi.Input<inputs.compute.v20200501.CreationData>;
    /**
     * ARM id of the DiskAccess resource for using private endpoints on disks.
     */
    readonly diskAccessId?: pulumi.Input<string>;
    /**
     * The total number of IOPS that will be allowed across all VMs mounting the shared disk as ReadOnly. One operation can transfer between 4k and 256k bytes.
     */
    readonly diskIOPSReadOnly?: pulumi.Input<number>;
    /**
     * The number of IOPS allowed for this disk; only settable for UltraSSD disks. One operation can transfer between 4k and 256k bytes.
     */
    readonly diskIOPSReadWrite?: pulumi.Input<number>;
    /**
     * The total throughput (MBps) that will be allowed across all VMs mounting the shared disk as ReadOnly. MBps means millions of bytes per second - MB here uses the ISO notation, of powers of 10.
     */
    readonly diskMBpsReadOnly?: pulumi.Input<number>;
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
    readonly encryption?: pulumi.Input<inputs.compute.v20200501.Encryption>;
    /**
     * Encryption settings collection used for Azure Disk Encryption, can contain multiple encryption settings per disk or snapshot.
     */
    readonly encryptionSettingsCollection?: pulumi.Input<inputs.compute.v20200501.EncryptionSettingsCollection>;
    /**
     * The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
     */
    readonly hyperVGeneration?: pulumi.Input<string>;
    /**
     * Resource location
     */
    readonly location: pulumi.Input<string>;
    /**
     * The maximum number of VMs that can attach to the disk at the same time. Value greater than one indicates a disk that can be mounted on multiple VMs at the same time.
     */
    readonly maxShares?: pulumi.Input<number>;
    /**
     * Policy for accessing the disk via network.
     */
    readonly networkAccessPolicy?: pulumi.Input<string>;
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
    readonly sku?: pulumi.Input<inputs.compute.v20200501.DiskSku>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The Logical zone list for Disk.
     */
    readonly zones?: pulumi.Input<pulumi.Input<string>[]>;
}
