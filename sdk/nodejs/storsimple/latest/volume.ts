// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * The volume.
 *
 * ## Example Usage
 * ### VolumesCreateOrUpdate
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const volume = new azurerm.storsimple.latest.Volume("volume", {
 *     accessControlRecordIds: ["/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/accessControlRecords/ACR2"],
 *     deviceName: "Device05ForSDKTest",
 *     managerName: "ManagerForSDKTest1",
 *     monitoringStatus: "Enabled",
 *     resourceGroupName: "ResourceGroupForSDKTest",
 *     sizeInBytes: 5368709120,
 *     volumeContainerName: "VolumeContainerForSDKTest",
 *     volumeName: "Volume1ForSDKTest",
 *     volumeStatus: "Offline",
 *     volumeType: "Tiered",
 * });
 *
 * ```
 */
export class Volume extends pulumi.CustomResource {
    /**
     * Get an existing Volume resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Volume {
        return new Volume(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:storsimple/latest:Volume';

    /**
     * Returns true if the given object is an instance of Volume.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Volume {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Volume.__pulumiType;
    }

    /**
     * The IDs of the access control records, associated with the volume.
     */
    public readonly accessControlRecordIds!: pulumi.Output<string[]>;
    /**
     * The IDs of the backup policies, in which this volume is part of.
     */
    public /*out*/ readonly backupPolicyIds!: pulumi.Output<string[]>;
    /**
     * The backup status of the volume.
     */
    public /*out*/ readonly backupStatus!: pulumi.Output<string>;
    /**
     * The Kind of the object. Currently only Series8000 is supported
     */
    public readonly kind!: pulumi.Output<string | undefined>;
    /**
     * The monitoring status of the volume.
     */
    public readonly monitoringStatus!: pulumi.Output<string>;
    /**
     * The name of the object.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The operation status on the volume.
     */
    public /*out*/ readonly operationStatus!: pulumi.Output<string>;
    /**
     * The size of the volume in bytes.
     */
    public readonly sizeInBytes!: pulumi.Output<number>;
    /**
     * The hierarchical type of the object.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The ID of the volume container, in which this volume is created.
     */
    public /*out*/ readonly volumeContainerId!: pulumi.Output<string>;
    /**
     * The volume status.
     */
    public readonly volumeStatus!: pulumi.Output<string>;
    /**
     * The type of the volume.
     */
    public readonly volumeType!: pulumi.Output<string>;

    /**
     * Create a Volume resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VolumeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VolumeArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as VolumeArgs | undefined;
            if (!args || args.accessControlRecordIds === undefined) {
                throw new Error("Missing required property 'accessControlRecordIds'");
            }
            if (!args || args.deviceName === undefined) {
                throw new Error("Missing required property 'deviceName'");
            }
            if (!args || args.managerName === undefined) {
                throw new Error("Missing required property 'managerName'");
            }
            if (!args || args.monitoringStatus === undefined) {
                throw new Error("Missing required property 'monitoringStatus'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.sizeInBytes === undefined) {
                throw new Error("Missing required property 'sizeInBytes'");
            }
            if (!args || args.volumeContainerName === undefined) {
                throw new Error("Missing required property 'volumeContainerName'");
            }
            if (!args || args.volumeName === undefined) {
                throw new Error("Missing required property 'volumeName'");
            }
            if (!args || args.volumeStatus === undefined) {
                throw new Error("Missing required property 'volumeStatus'");
            }
            if (!args || args.volumeType === undefined) {
                throw new Error("Missing required property 'volumeType'");
            }
            inputs["accessControlRecordIds"] = args ? args.accessControlRecordIds : undefined;
            inputs["deviceName"] = args ? args.deviceName : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["managerName"] = args ? args.managerName : undefined;
            inputs["monitoringStatus"] = args ? args.monitoringStatus : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sizeInBytes"] = args ? args.sizeInBytes : undefined;
            inputs["volumeContainerName"] = args ? args.volumeContainerName : undefined;
            inputs["volumeName"] = args ? args.volumeName : undefined;
            inputs["volumeStatus"] = args ? args.volumeStatus : undefined;
            inputs["volumeType"] = args ? args.volumeType : undefined;
            inputs["backupPolicyIds"] = undefined /*out*/;
            inputs["backupStatus"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["operationStatus"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["volumeContainerId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:storsimple/v20170601:Volume" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Volume.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Volume resource.
 */
export interface VolumeArgs {
    /**
     * The IDs of the access control records, associated with the volume.
     */
    readonly accessControlRecordIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The device name
     */
    readonly deviceName: pulumi.Input<string>;
    /**
     * The Kind of the object. Currently only Series8000 is supported
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * The manager name
     */
    readonly managerName: pulumi.Input<string>;
    /**
     * The monitoring status of the volume.
     */
    readonly monitoringStatus: pulumi.Input<string>;
    /**
     * The resource group name
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The size of the volume in bytes.
     */
    readonly sizeInBytes: pulumi.Input<number>;
    /**
     * The volume container name.
     */
    readonly volumeContainerName: pulumi.Input<string>;
    /**
     * The volume name.
     */
    readonly volumeName: pulumi.Input<string>;
    /**
     * The volume status.
     */
    readonly volumeStatus: pulumi.Input<string>;
    /**
     * The type of the volume.
     */
    readonly volumeType: pulumi.Input<string>;
}
