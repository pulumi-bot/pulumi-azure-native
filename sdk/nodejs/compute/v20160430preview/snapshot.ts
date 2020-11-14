// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
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
    public static readonly __pulumiType = 'azure-nextgen:compute/v20160430preview:Snapshot';

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
     * the storage account type of the disk.
     */
    public readonly accountType!: pulumi.Output<string | undefined>;
    /**
     * Disk source information. CreationData information cannot be changed after the disk has been created.
     */
    public readonly creationData!: pulumi.Output<outputs.compute.v20160430preview.CreationDataResponse>;
    /**
     * If creationData.createOption is Empty, this field is mandatory and it indicates the size of the VHD to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
     */
    public readonly diskSizeGB!: pulumi.Output<number | undefined>;
    /**
     * Encryption settings for disk or snapshot
     */
    public readonly encryptionSettings!: pulumi.Output<outputs.compute.v20160430preview.EncryptionSettingsResponse | undefined>;
    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The Operating System type.
     */
    public readonly osType!: pulumi.Output<string | undefined>;
    /**
     * A relative URI containing the VM id that has the disk attached.
     */
    public /*out*/ readonly ownerId!: pulumi.Output<string>;
    /**
     * The disk provisioning state.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
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
     * Create a Snapshot resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SnapshotArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.creationData === undefined) {
                throw new Error("Missing required property 'creationData'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.snapshotName === undefined) {
                throw new Error("Missing required property 'snapshotName'");
            }
            inputs["accountType"] = args ? args.accountType : undefined;
            inputs["creationData"] = args ? args.creationData : undefined;
            inputs["diskSizeGB"] = args ? args.diskSizeGB : undefined;
            inputs["encryptionSettings"] = args ? args.encryptionSettings : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["osType"] = args ? args.osType : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["snapshotName"] = args ? args.snapshotName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["ownerId"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["timeCreated"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["accountType"] = undefined /*out*/;
            inputs["creationData"] = undefined /*out*/;
            inputs["diskSizeGB"] = undefined /*out*/;
            inputs["encryptionSettings"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["osType"] = undefined /*out*/;
            inputs["ownerId"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["timeCreated"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:compute/latest:Snapshot" }, { type: "azure-nextgen:compute/v20170330:Snapshot" }, { type: "azure-nextgen:compute/v20180401:Snapshot" }, { type: "azure-nextgen:compute/v20180601:Snapshot" }, { type: "azure-nextgen:compute/v20180930:Snapshot" }, { type: "azure-nextgen:compute/v20190301:Snapshot" }, { type: "azure-nextgen:compute/v20190701:Snapshot" }, { type: "azure-nextgen:compute/v20191101:Snapshot" }, { type: "azure-nextgen:compute/v20200501:Snapshot" }, { type: "azure-nextgen:compute/v20200630:Snapshot" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Snapshot.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Snapshot resource.
 */
export interface SnapshotArgs {
    /**
     * the storage account type of the disk.
     */
    readonly accountType?: pulumi.Input<string>;
    /**
     * Disk source information. CreationData information cannot be changed after the disk has been created.
     */
    readonly creationData: pulumi.Input<inputs.compute.v20160430preview.CreationData>;
    /**
     * If creationData.createOption is Empty, this field is mandatory and it indicates the size of the VHD to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
     */
    readonly diskSizeGB?: pulumi.Input<number>;
    /**
     * Encryption settings for disk or snapshot
     */
    readonly encryptionSettings?: pulumi.Input<inputs.compute.v20160430preview.EncryptionSettings>;
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
     * The name of the snapshot within the given subscription and resource group.
     */
    readonly snapshotName: pulumi.Input<string>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
