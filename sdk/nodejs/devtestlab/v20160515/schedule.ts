// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * A schedule.
 */
export class Schedule extends pulumi.CustomResource {
    /**
     * Get an existing Schedule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Schedule {
        return new Schedule(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:devtestlab/v20160515:Schedule';

    /**
     * Returns true if the given object is an instance of Schedule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Schedule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Schedule.__pulumiType;
    }

    /**
     * The location of the resource.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The properties of the resource.
     */
    public /*out*/ readonly properties!: pulumi.Output<outputs.devtestlab.v20160515.SchedulePropertiesResponse>;
    /**
     * The tags of the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Schedule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ScheduleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ScheduleArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as ScheduleArgs | undefined;
            if (!args || args.labName === undefined) {
                throw new Error("Missing required property 'labName'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["dailyRecurrence"] = args ? args.dailyRecurrence : undefined;
            inputs["hourlyRecurrence"] = args ? args.hourlyRecurrence : undefined;
            inputs["labName"] = args ? args.labName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["notificationSettings"] = args ? args.notificationSettings : undefined;
            inputs["provisioningState"] = args ? args.provisioningState : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["targetResourceId"] = args ? args.targetResourceId : undefined;
            inputs["taskType"] = args ? args.taskType : undefined;
            inputs["timeZoneId"] = args ? args.timeZoneId : undefined;
            inputs["uniqueIdentifier"] = args ? args.uniqueIdentifier : undefined;
            inputs["weeklyRecurrence"] = args ? args.weeklyRecurrence : undefined;
            inputs["properties"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Schedule.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Schedule resource.
 */
export interface ScheduleArgs {
    /**
     * If the schedule will occur once each day of the week, specify the daily recurrence.
     */
    readonly dailyRecurrence?: pulumi.Input<inputs.devtestlab.v20160515.DayDetails>;
    /**
     * If the schedule will occur multiple times a day, specify the hourly recurrence.
     */
    readonly hourlyRecurrence?: pulumi.Input<inputs.devtestlab.v20160515.HourDetails>;
    /**
     * The name of the lab.
     */
    readonly labName: pulumi.Input<string>;
    /**
     * The location of the resource.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the schedule.
     */
    readonly name: pulumi.Input<string>;
    /**
     * Notification settings.
     */
    readonly notificationSettings?: pulumi.Input<inputs.devtestlab.v20160515.NotificationSettings>;
    /**
     * The provisioning status of the resource.
     */
    readonly provisioningState?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The status of the schedule (i.e. Enabled, Disabled)
     */
    readonly status?: pulumi.Input<string>;
    /**
     * The tags of the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The resource ID to which the schedule belongs
     */
    readonly targetResourceId?: pulumi.Input<string>;
    /**
     * The task type of the schedule (e.g. LabVmsShutdownTask, LabVmAutoStart).
     */
    readonly taskType?: pulumi.Input<string>;
    /**
     * The time zone ID (e.g. Pacific Standard time).
     */
    readonly timeZoneId?: pulumi.Input<string>;
    /**
     * The unique immutable identifier of a resource (Guid).
     */
    readonly uniqueIdentifier?: pulumi.Input<string>;
    /**
     * If the schedule will occur only some days of the week, specify the weekly recurrence.
     */
    readonly weeklyRecurrence?: pulumi.Input<inputs.devtestlab.v20160515.WeekDetails>;
}
