// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Snapshot resource.
 */
export class Snapshot extends pulumi.CustomResource {
    /**
     * Get an existing Snapshot resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Snapshot {
        return new Snapshot(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:compute/v20200501:Snapshot';

    /**
     * Returns true if the given object is an instance of Snapshot.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Snapshot {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Snapshot.__pulumiType;
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
     * The size of the disk in bytes. This field is read only.
     */
    public /*out*/ readonly diskSizeBytes!: pulumi.Output<number>;
    /**
     * If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
     */
    public readonly diskSizeGB!: pulumi.Output<number | undefined>;
    /**
     * Encryption property can be used to encrypt data at rest with customer managed keys or platform managed keys.
     */
    public readonly encryption!: pulumi.Output<outputs.compute.v20200501.EncryptionResponse | undefined>;
    /**
     * Encryption settings collection used be Azure Disk Encryption, can contain multiple encryption settings per disk or snapshot.
     */
    public readonly encryptionSettingsCollection!: pulumi.Output<outputs.compute.v20200501.EncryptionSettingsCollectionResponse | undefined>;
    /**
     * The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
     */
    public readonly hyperVGeneration!: pulumi.Output<string | undefined>;
    /**
     * Whether a snapshot is incremental. Incremental snapshots on the same disk occupy less space than full snapshots and can be diffed.
     */
    public readonly incremental!: pulumi.Output<boolean | undefined>;
    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Unused. Always Null.
     */
    public /*out*/ readonly managedBy!: pulumi.Output<string>;
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
     * The snapshots sku name. Can be Standard_LRS, Premium_LRS, or Standard_ZRS.
     */
    public readonly sku!: pulumi.Output<outputs.compute.v20200501.SnapshotSkuResponse | undefined>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The time when the snapshot was created.
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
     * Create a Snapshot resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SnapshotArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.creationData === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'creationData'");
            }
            if ((!args || args.resourceGroupName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.snapshotName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'snapshotName'");
            }
            inputs["creationData"] = args ? args.creationData : undefined;
            inputs["diskAccessId"] = args ? args.diskAccessId : undefined;
            inputs["diskSizeGB"] = args ? args.diskSizeGB : undefined;
            inputs["encryption"] = args ? args.encryption : undefined;
            inputs["encryptionSettingsCollection"] = args ? args.encryptionSettingsCollection : undefined;
            inputs["hyperVGeneration"] = args ? args.hyperVGeneration : undefined;
            inputs["incremental"] = args ? args.incremental : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["networkAccessPolicy"] = args ? args.networkAccessPolicy : undefined;
            inputs["osType"] = args ? args.osType : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["snapshotName"] = args ? args.snapshotName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["diskSizeBytes"] = undefined /*out*/;
            inputs["managedBy"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["timeCreated"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["uniqueId"] = undefined /*out*/;
        } else {
            inputs["creationData"] = undefined /*out*/;
            inputs["diskAccessId"] = undefined /*out*/;
            inputs["diskSizeBytes"] = undefined /*out*/;
            inputs["diskSizeGB"] = undefined /*out*/;
            inputs["encryption"] = undefined /*out*/;
            inputs["encryptionSettingsCollection"] = undefined /*out*/;
            inputs["hyperVGeneration"] = undefined /*out*/;
            inputs["incremental"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["managedBy"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["networkAccessPolicy"] = undefined /*out*/;
            inputs["osType"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["sku"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["timeCreated"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["uniqueId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:compute/latest:Snapshot" }, { type: "azure-nextgen:compute/v20160430preview:Snapshot" }, { type: "azure-nextgen:compute/v20170330:Snapshot" }, { type: "azure-nextgen:compute/v20180401:Snapshot" }, { type: "azure-nextgen:compute/v20180601:Snapshot" }, { type: "azure-nextgen:compute/v20180930:Snapshot" }, { type: "azure-nextgen:compute/v20190301:Snapshot" }, { type: "azure-nextgen:compute/v20190701:Snapshot" }, { type: "azure-nextgen:compute/v20191101:Snapshot" }, { type: "azure-nextgen:compute/v20200630:Snapshot" }, { type: "azure-nextgen:compute/v20200930:Snapshot" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Snapshot.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Snapshot resource.
 */
export interface SnapshotArgs {
    /**
     * Disk source information. CreationData information cannot be changed after the disk has been created.
     */
    readonly creationData: pulumi.Input<inputs.compute.v20200501.CreationData>;
    /**
     * ARM id of the DiskAccess resource for using private endpoints on disks.
     */
    readonly diskAccessId?: pulumi.Input<string>;
    /**
     * If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
     */
    readonly diskSizeGB?: pulumi.Input<number>;
    /**
     * Encryption property can be used to encrypt data at rest with customer managed keys or platform managed keys.
     */
    readonly encryption?: pulumi.Input<inputs.compute.v20200501.Encryption>;
    /**
     * Encryption settings collection used be Azure Disk Encryption, can contain multiple encryption settings per disk or snapshot.
     */
    readonly encryptionSettingsCollection?: pulumi.Input<inputs.compute.v20200501.EncryptionSettingsCollection>;
    /**
     * The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
     */
    readonly hyperVGeneration?: pulumi.Input<string | enums.compute.v20200501.HyperVGeneration>;
    /**
     * Whether a snapshot is incremental. Incremental snapshots on the same disk occupy less space than full snapshots and can be diffed.
     */
    readonly incremental?: pulumi.Input<boolean>;
    /**
     * Resource location
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Policy for accessing the disk via network.
     */
    readonly networkAccessPolicy?: pulumi.Input<string | enums.compute.v20200501.NetworkAccessPolicy>;
    /**
     * The Operating System type.
     */
    readonly osType?: pulumi.Input<enums.compute.v20200501.OperatingSystemTypes>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The snapshots sku name. Can be Standard_LRS, Premium_LRS, or Standard_ZRS.
     */
    readonly sku?: pulumi.Input<inputs.compute.v20200501.SnapshotSku>;
    /**
     * The name of the snapshot that is being created. The name can't be changed after the snapshot is created. Supported characters for the name are a-z, A-Z, 0-9 and _. The max name length is 80 characters.
     */
    readonly snapshotName: pulumi.Input<string>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
