// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The volume container.
 *
 * ## Example Usage
 * ### VolumeContainersCreateOrUpdate
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const volumeContainer = new azurerm.storsimple.v20170601.VolumeContainer("volumeContainer", {
 *     bandwidthSettingId: "/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/bandwidthSettings/bandwidthSetting1",
 *     deviceName: "Device05ForSDKTest",
 *     encryptionKey: {
 *         encryptionAlgorithm: "RSAES_PKCS1_v_1_5",
 *         encryptionCertThumbprint: "A872A2DF196AC7682EE24791E7DE2E2A360F5926",
 *         value: "R//pyVLx/fn58ia098JiLgZB5RY7fVT+6o8a4fmsvjy+ls2UgJphMf25XVqEQCZnsp/5uxteN1M/9ArPIICdhM7M1+b/Ur7kJ0FH0ktxfk7CrPWWJLI4q20LZoduJGI56lREav1VpuLdqw5F9fRcq7zbfgPQ3B/SD0mfumNRiV+AnwbC6msfavIuWrhVDl9iSzEPE+zU06/kpsexnrS81yYT2QlVVUbvpY4F3zfH8TQPpAROTbv2pld6JO4eGOrZ5O1iOr6XCg2TY2W/jf+Ev4z5tqC9VWXE5kh65gjBfpWN0bDWXKekqEhor2crHAxZi4dybdY8Ok1MDWd1CSU8kw==",
 *     },
 *     managerName: "ManagerForSDKTest1",
 *     resourceGroupName: "ResourceGroupForSDKTest",
 *     storageAccountCredentialId: "/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/storageAccountCredentials/safortestrecording",
 *     volumeContainerName: "VolumeContainerForSDKTest",
 * });
 *
 * ```
 */
export class VolumeContainer extends pulumi.CustomResource {
    /**
     * Get an existing VolumeContainer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): VolumeContainer {
        return new VolumeContainer(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:storsimple/v20170601:VolumeContainer';

    /**
     * Returns true if the given object is an instance of VolumeContainer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VolumeContainer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VolumeContainer.__pulumiType;
    }

    /**
     * The bandwidth-rate set on the volume container.
     */
    public readonly bandWidthRateInMbps!: pulumi.Output<number | undefined>;
    /**
     * The ID of the bandwidth setting associated with the volume container.
     */
    public readonly bandwidthSettingId!: pulumi.Output<string | undefined>;
    /**
     * The key used to encrypt data in the volume container. It is required when property 'EncryptionStatus' is "Enabled".
     */
    public readonly encryptionKey!: pulumi.Output<outputs.storsimple.v20170601.AsymmetricEncryptedSecretResponse | undefined>;
    /**
     * The flag to denote whether encryption is enabled or not.
     */
    public /*out*/ readonly encryptionStatus!: pulumi.Output<string>;
    /**
     * The Kind of the object. Currently only Series8000 is supported
     */
    public readonly kind!: pulumi.Output<string | undefined>;
    /**
     * The name of the object.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The owner ship status of the volume container. Only when the status is "NotOwned", the delete operation on the volume container is permitted.
     */
    public /*out*/ readonly ownerShipStatus!: pulumi.Output<string>;
    /**
     * The path ID of storage account associated with the volume container.
     */
    public readonly storageAccountCredentialId!: pulumi.Output<string>;
    /**
     * The total cloud storage for the volume container.
     */
    public /*out*/ readonly totalCloudStorageUsageInBytes!: pulumi.Output<number>;
    /**
     * The hierarchical type of the object.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The number of volumes in the volume Container.
     */
    public /*out*/ readonly volumeCount!: pulumi.Output<number>;

    /**
     * Create a VolumeContainer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VolumeContainerArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.deviceName === undefined) {
                throw new Error("Missing required property 'deviceName'");
            }
            if (!args || args.managerName === undefined) {
                throw new Error("Missing required property 'managerName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.storageAccountCredentialId === undefined) {
                throw new Error("Missing required property 'storageAccountCredentialId'");
            }
            if (!args || args.volumeContainerName === undefined) {
                throw new Error("Missing required property 'volumeContainerName'");
            }
            inputs["bandWidthRateInMbps"] = args ? args.bandWidthRateInMbps : undefined;
            inputs["bandwidthSettingId"] = args ? args.bandwidthSettingId : undefined;
            inputs["deviceName"] = args ? args.deviceName : undefined;
            inputs["encryptionKey"] = args ? args.encryptionKey : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["managerName"] = args ? args.managerName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["storageAccountCredentialId"] = args ? args.storageAccountCredentialId : undefined;
            inputs["volumeContainerName"] = args ? args.volumeContainerName : undefined;
            inputs["encryptionStatus"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["ownerShipStatus"] = undefined /*out*/;
            inputs["totalCloudStorageUsageInBytes"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["volumeCount"] = undefined /*out*/;
        } else {
            inputs["bandWidthRateInMbps"] = undefined /*out*/;
            inputs["bandwidthSettingId"] = undefined /*out*/;
            inputs["encryptionKey"] = undefined /*out*/;
            inputs["encryptionStatus"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["ownerShipStatus"] = undefined /*out*/;
            inputs["storageAccountCredentialId"] = undefined /*out*/;
            inputs["totalCloudStorageUsageInBytes"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["volumeCount"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:storsimple/latest:VolumeContainer" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(VolumeContainer.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a VolumeContainer resource.
 */
export interface VolumeContainerArgs {
    /**
     * The bandwidth-rate set on the volume container.
     */
    readonly bandWidthRateInMbps?: pulumi.Input<number>;
    /**
     * The ID of the bandwidth setting associated with the volume container.
     */
    readonly bandwidthSettingId?: pulumi.Input<string>;
    /**
     * The device name
     */
    readonly deviceName: pulumi.Input<string>;
    /**
     * The key used to encrypt data in the volume container. It is required when property 'EncryptionStatus' is "Enabled".
     */
    readonly encryptionKey?: pulumi.Input<inputs.storsimple.v20170601.AsymmetricEncryptedSecret>;
    /**
     * The Kind of the object. Currently only Series8000 is supported
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * The manager name
     */
    readonly managerName: pulumi.Input<string>;
    /**
     * The resource group name
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The path ID of storage account associated with the volume container.
     */
    readonly storageAccountCredentialId: pulumi.Input<string>;
    /**
     * The name of the volume container.
     */
    readonly volumeContainerName: pulumi.Input<string>;
}
