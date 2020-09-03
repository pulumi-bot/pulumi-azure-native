// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * The iSCSI server.
 *
 * ## Example Usage
 * ### IscsiServersCreateOrUpdate
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const iscsiServer = new azurerm.storsimple.v20161001.IscsiServer("iscsiServer", {
 *     backupScheduleGroupId: "/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-WSJQERQW3F/backupScheduleGroups/Default-HSDK-WSJQERQW3F-BackupScheduleGroup",
 *     chapId: "/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-WSJQERQW3F/chapSettings/ChapSettingForSDK",
 *     deviceName: "HSDK-WSJQERQW3F",
 *     iscsiServerName: "HSDK-WSJQERQW3F",
 *     managerName: "hAzureSDKOperations",
 *     resourceGroupName: "ResourceGroupForSDKTest",
 *     reverseChapId: "",
 *     storageDomainId: "/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/storageDomains/Default-HSDK-WSJQERQW3F-StorageDomain",
 * });
 *
 * ```
 */
export class IscsiServer extends pulumi.CustomResource {
    /**
     * Get an existing IscsiServer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): IscsiServer {
        return new IscsiServer(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:storsimple/v20161001:IscsiServer';

    /**
     * Returns true if the given object is an instance of IscsiServer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IscsiServer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IscsiServer.__pulumiType;
    }

    /**
     * The backup policy id.
     */
    public readonly backupScheduleGroupId!: pulumi.Output<string>;
    /**
     * The chap id.
     */
    public readonly chapId!: pulumi.Output<string | undefined>;
    /**
     * The description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The reverse chap id.
     */
    public readonly reverseChapId!: pulumi.Output<string | undefined>;
    /**
     * The storage domain id.
     */
    public readonly storageDomainId!: pulumi.Output<string>;
    /**
     * The type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a IscsiServer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IscsiServerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IscsiServerArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as IscsiServerArgs | undefined;
            if (!args || args.backupScheduleGroupId === undefined) {
                throw new Error("Missing required property 'backupScheduleGroupId'");
            }
            if (!args || args.deviceName === undefined) {
                throw new Error("Missing required property 'deviceName'");
            }
            if (!args || args.iscsiServerName === undefined) {
                throw new Error("Missing required property 'iscsiServerName'");
            }
            if (!args || args.managerName === undefined) {
                throw new Error("Missing required property 'managerName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.storageDomainId === undefined) {
                throw new Error("Missing required property 'storageDomainId'");
            }
            inputs["backupScheduleGroupId"] = args ? args.backupScheduleGroupId : undefined;
            inputs["chapId"] = args ? args.chapId : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["deviceName"] = args ? args.deviceName : undefined;
            inputs["iscsiServerName"] = args ? args.iscsiServerName : undefined;
            inputs["managerName"] = args ? args.managerName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["reverseChapId"] = args ? args.reverseChapId : undefined;
            inputs["storageDomainId"] = args ? args.storageDomainId : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:storsimple/latest:IscsiServer" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(IscsiServer.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a IscsiServer resource.
 */
export interface IscsiServerArgs {
    /**
     * The backup policy id.
     */
    readonly backupScheduleGroupId: pulumi.Input<string>;
    /**
     * The chap id.
     */
    readonly chapId?: pulumi.Input<string>;
    /**
     * The description.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The device name.
     */
    readonly deviceName: pulumi.Input<string>;
    /**
     * The iSCSI server name.
     */
    readonly iscsiServerName: pulumi.Input<string>;
    /**
     * The manager name
     */
    readonly managerName: pulumi.Input<string>;
    /**
     * The resource group name
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The reverse chap id.
     */
    readonly reverseChapId?: pulumi.Input<string>;
    /**
     * The storage domain id.
     */
    readonly storageDomainId: pulumi.Input<string>;
}
