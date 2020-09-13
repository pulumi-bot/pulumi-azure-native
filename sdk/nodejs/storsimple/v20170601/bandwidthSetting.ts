// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The bandwidth setting.
 */
export class BandwidthSetting extends pulumi.CustomResource {
    /**
     * Get an existing BandwidthSetting resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): BandwidthSetting {
        return new BandwidthSetting(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:storsimple/v20170601:BandwidthSetting';

    /**
     * Returns true if the given object is an instance of BandwidthSetting.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BandwidthSetting {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BandwidthSetting.__pulumiType;
    }

    /**
     * The Kind of the object. Currently only Series8000 is supported
     */
    public readonly kind!: pulumi.Output<string | undefined>;
    /**
     * The name of the object.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The schedules.
     */
    public readonly schedules!: pulumi.Output<outputs.storsimple.v20170601.BandwidthScheduleResponse[]>;
    /**
     * The hierarchical type of the object.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The number of volumes that uses the bandwidth setting.
     */
    public /*out*/ readonly volumeCount!: pulumi.Output<number>;

    /**
     * Create a BandwidthSetting resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BandwidthSettingArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.bandwidthSettingName === undefined) {
                throw new Error("Missing required property 'bandwidthSettingName'");
            }
            if (!args || args.managerName === undefined) {
                throw new Error("Missing required property 'managerName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.schedules === undefined) {
                throw new Error("Missing required property 'schedules'");
            }
            inputs["bandwidthSettingName"] = args ? args.bandwidthSettingName : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["managerName"] = args ? args.managerName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["schedules"] = args ? args.schedules : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["volumeCount"] = undefined /*out*/;
        } else {
            inputs["kind"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["schedules"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["volumeCount"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:storsimple/latest:BandwidthSetting" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(BandwidthSetting.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a BandwidthSetting resource.
 */
export interface BandwidthSettingArgs {
    /**
     * The bandwidth setting name.
     */
    readonly bandwidthSettingName: pulumi.Input<string>;
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
     * The schedules.
     */
    readonly schedules: pulumi.Input<pulumi.Input<inputs.storsimple.v20170601.BandwidthSchedule>[]>;
}
