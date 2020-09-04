// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Snapshot policy information
 *
 * ## Example Usage
 * ### SnapshotPolicies_Create
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const snapshotPolicy = new azurerm.netapp.v20200601.SnapshotPolicy("snapshotPolicy", {
 *     accountName: "account1",
 *     dailySchedule: {
 *         hour: 10,
 *         minute: 10,
 *         snapshotsToKeep: 10,
 *         usedBytes: 100000,
 *     },
 *     enabled: true,
 *     hourlySchedule: {
 *         minute: 10,
 *         snapshotsToKeep: 10,
 *         usedBytes: 100000,
 *     },
 *     location: "eastus",
 *     monthlySchedule: {
 *         DaysOfMonth: "1,5,11",
 *         hour: 10,
 *         minute: 10,
 *         snapshotsToKeep: 10,
 *         usedBytes: 100000,
 *     },
 *     resourceGroupName: "myRG",
 *     snapshotPolicyName: "snapshotPolicyName",
 *     weeklySchedule: {
 *         day: "Monday",
 *         hour: 10,
 *         minute: 10,
 *         snapshotsToKeep: 10,
 *         usedBytes: 100000,
 *     },
 * });
 *
 * ```
 */
export class SnapshotPolicy extends pulumi.CustomResource {
    /**
     * Get an existing SnapshotPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SnapshotPolicy {
        return new SnapshotPolicy(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:netapp/v20200601:snapshotPolicy';

    /**
     * Returns true if the given object is an instance of SnapshotPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SnapshotPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SnapshotPolicy.__pulumiType;
    }

    /**
     * Schedule for daily snapshots
     */
    public readonly dailySchedule!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The property to decide policy is enabled or not
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * Schedule for hourly snapshots
     */
    public readonly hourlySchedule!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Schedule for monthly snapshots
     */
    public readonly monthlySchedule!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Schedule for weekly snapshots
     */
    public readonly weeklySchedule!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a SnapshotPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SnapshotPolicyArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.accountName === undefined) {
                throw new Error("Missing required property 'accountName'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.snapshotPolicyName === undefined) {
                throw new Error("Missing required property 'snapshotPolicyName'");
            }
            inputs["accountName"] = args ? args.accountName : undefined;
            inputs["dailySchedule"] = args ? args.dailySchedule : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["hourlySchedule"] = args ? args.hourlySchedule : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["monthlySchedule"] = args ? args.monthlySchedule : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["snapshotPolicyName"] = args ? args.snapshotPolicyName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["weeklySchedule"] = args ? args.weeklySchedule : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["dailySchedule"] = undefined /*out*/;
            inputs["enabled"] = undefined /*out*/;
            inputs["hourlySchedule"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["monthlySchedule"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["weeklySchedule"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:netapp/latest:snapshotPolicy" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(SnapshotPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a SnapshotPolicy resource.
 */
export interface SnapshotPolicyArgs {
    /**
     * The name of the NetApp account
     */
    readonly accountName: pulumi.Input<string>;
    /**
     * Schedule for daily snapshots
     */
    readonly dailySchedule?: pulumi.Input<{[key: string]: any}>;
    /**
     * The property to decide policy is enabled or not
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * Schedule for hourly snapshots
     */
    readonly hourlySchedule?: pulumi.Input<{[key: string]: any}>;
    /**
     * Resource location
     */
    readonly location: pulumi.Input<string>;
    /**
     * Schedule for monthly snapshots
     */
    readonly monthlySchedule?: pulumi.Input<{[key: string]: any}>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the snapshot policy target
     */
    readonly snapshotPolicyName: pulumi.Input<string>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Schedule for weekly snapshots
     */
    readonly weeklySchedule?: pulumi.Input<{[key: string]: any}>;
}
