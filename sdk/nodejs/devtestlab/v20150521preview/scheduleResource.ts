// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * A schedule.
 */
export class ScheduleResource extends pulumi.CustomResource {
    /**
     * Get an existing ScheduleResource resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ScheduleResource {
        return new ScheduleResource(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:devtestlab/v20150521preview:ScheduleResource';

    /**
     * Returns true if the given object is an instance of ScheduleResource.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ScheduleResource {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ScheduleResource.__pulumiType;
    }

    /**
     * The daily recurrence of the schedule.
     */
    public readonly dailyRecurrence!: pulumi.Output<outputs.devtestlab.v20150521preview.DayDetailsResponse | undefined>;
    /**
     * The hourly recurrence of the schedule.
     */
    public readonly hourlyRecurrence!: pulumi.Output<outputs.devtestlab.v20150521preview.HourDetailsResponse | undefined>;
    /**
     * The location of the resource.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * The provisioning status of the resource.
     */
    public readonly provisioningState!: pulumi.Output<string | undefined>;
    /**
     * The status of the schedule.
     */
    public readonly status!: pulumi.Output<string | undefined>;
    /**
     * The tags of the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The task type of the schedule.
     */
    public readonly taskType!: pulumi.Output<string | undefined>;
    /**
     * The time zone id.
     */
    public readonly timeZoneId!: pulumi.Output<string | undefined>;
    /**
     * The type of the resource.
     */
    public readonly type!: pulumi.Output<string | undefined>;
    /**
     * The weekly recurrence of the schedule.
     */
    public readonly weeklyRecurrence!: pulumi.Output<outputs.devtestlab.v20150521preview.WeekDetailsResponse | undefined>;

    /**
     * Create a ScheduleResource resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ScheduleResourceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
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
            inputs["id"] = args ? args.id : undefined;
            inputs["labName"] = args ? args.labName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["provisioningState"] = args ? args.provisioningState : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["taskType"] = args ? args.taskType : undefined;
            inputs["timeZoneId"] = args ? args.timeZoneId : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["weeklyRecurrence"] = args ? args.weeklyRecurrence : undefined;
        } else {
            inputs["dailyRecurrence"] = undefined /*out*/;
            inputs["hourlyRecurrence"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["taskType"] = undefined /*out*/;
            inputs["timeZoneId"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["weeklyRecurrence"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:devtestlab/latest:ScheduleResource" }, { type: "azurerm:devtestlab/v20160515:ScheduleResource" }, { type: "azurerm:devtestlab/v20180915:ScheduleResource" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ScheduleResource.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ScheduleResource resource.
 */
export interface ScheduleResourceArgs {
    /**
     * The daily recurrence of the schedule.
     */
    readonly dailyRecurrence?: pulumi.Input<inputs.devtestlab.v20150521preview.DayDetails>;
    /**
     * The hourly recurrence of the schedule.
     */
    readonly hourlyRecurrence?: pulumi.Input<inputs.devtestlab.v20150521preview.HourDetails>;
    /**
     * The identifier of the resource.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * The name of the lab.
     */
    readonly labName: pulumi.Input<string>;
    /**
     * The location of the resource.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the resource.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The provisioning status of the resource.
     */
    readonly provisioningState?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The status of the schedule.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * The tags of the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The task type of the schedule.
     */
    readonly taskType?: pulumi.Input<string>;
    /**
     * The time zone id.
     */
    readonly timeZoneId?: pulumi.Input<string>;
    /**
     * The type of the resource.
     */
    readonly type?: pulumi.Input<string>;
    /**
     * The weekly recurrence of the schedule.
     */
    readonly weeklyRecurrence?: pulumi.Input<inputs.devtestlab.v20150521preview.WeekDetails>;
}
