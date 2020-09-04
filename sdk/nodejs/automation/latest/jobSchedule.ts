// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Definition of the job schedule.
 *
 * ## Example Usage
 * ### Create a job schedule
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const jobSchedule = new azurerm.automation.latest.JobSchedule("jobSchedule", {
 *     automationAccountName: "ContoseAutomationAccount",
 *     jobScheduleId: "0fa462ba-3aa2-4138-83ca-9ebc3bc55cdc",
 *     parameters: {
 *         jobscheduletag01: "jobschedulevalue01",
 *         jobscheduletag02: "jobschedulevalue02",
 *     },
 *     resourceGroupName: "rg",
 *     runbook: {
 *         name: "TestRunbook",
 *     },
 *     schedule: {
 *         name: "ScheduleNameGoesHere332204b5-debe-4348-a5c7-6357457189f2",
 *     },
 * });
 *
 * ```
 */
export class JobSchedule extends pulumi.CustomResource {
    /**
     * Get an existing JobSchedule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): JobSchedule {
        return new JobSchedule(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:automation/latest:JobSchedule';

    /**
     * Returns true if the given object is an instance of JobSchedule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is JobSchedule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === JobSchedule.__pulumiType;
    }

    /**
     * Gets or sets the id of job schedule.
     */
    public readonly jobScheduleId!: pulumi.Output<string | undefined>;
    /**
     * Gets the name of the variable.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Gets or sets the parameters of the job schedule.
     */
    public readonly parameters!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Gets or sets the hybrid worker group that the scheduled job should run on.
     */
    public readonly runOn!: pulumi.Output<string | undefined>;
    /**
     * Gets or sets the runbook.
     */
    public readonly runbook!: pulumi.Output<outputs.automation.latest.RunbookAssociationPropertyResponse | undefined>;
    /**
     * Gets or sets the schedule.
     */
    public readonly schedule!: pulumi.Output<outputs.automation.latest.ScheduleAssociationPropertyResponse | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a JobSchedule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: JobScheduleArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.automationAccountName === undefined) {
                throw new Error("Missing required property 'automationAccountName'");
            }
            if (!args || args.jobScheduleId === undefined) {
                throw new Error("Missing required property 'jobScheduleId'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.runbook === undefined) {
                throw new Error("Missing required property 'runbook'");
            }
            if (!args || args.schedule === undefined) {
                throw new Error("Missing required property 'schedule'");
            }
            inputs["automationAccountName"] = args ? args.automationAccountName : undefined;
            inputs["jobScheduleId"] = args ? args.jobScheduleId : undefined;
            inputs["parameters"] = args ? args.parameters : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["runOn"] = args ? args.runOn : undefined;
            inputs["runbook"] = args ? args.runbook : undefined;
            inputs["schedule"] = args ? args.schedule : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["jobScheduleId"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["parameters"] = undefined /*out*/;
            inputs["runOn"] = undefined /*out*/;
            inputs["runbook"] = undefined /*out*/;
            inputs["schedule"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:automation/v20151031:JobSchedule" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(JobSchedule.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a JobSchedule resource.
 */
export interface JobScheduleArgs {
    /**
     * The name of the automation account.
     */
    readonly automationAccountName: pulumi.Input<string>;
    /**
     * The job schedule name.
     */
    readonly jobScheduleId: pulumi.Input<string>;
    /**
     * Gets or sets a list of job properties.
     */
    readonly parameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name of an Azure Resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Gets or sets the hybrid worker group that the scheduled job should run on.
     */
    readonly runOn?: pulumi.Input<string>;
    /**
     * Gets or sets the runbook.
     */
    readonly runbook: pulumi.Input<inputs.automation.latest.RunbookAssociationProperty>;
    /**
     * Gets or sets the schedule.
     */
    readonly schedule: pulumi.Input<inputs.automation.latest.ScheduleAssociationProperty>;
}
