// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * A Job resource type. The progress and state can be obtained by polling a Job or subscribing to events using EventGrid.
 */
export class Job extends pulumi.CustomResource {
    /**
     * Get an existing Job resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Job {
        return new Job(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:media/latest:Job';

    /**
     * Returns true if the given object is an instance of Job.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Job {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Job.__pulumiType;
    }

    /**
     * Customer provided key, value pairs that will be returned in Job and JobOutput state events.
     */
    public readonly correlationData!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The UTC date and time when the customer has created the Job, in 'YYYY-MM-DDThh:mm:ssZ' format.
     */
    public /*out*/ readonly created!: pulumi.Output<string>;
    /**
     * Optional customer supplied description of the Job.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The UTC date and time at which this Job finished processing.
     */
    public /*out*/ readonly endTime!: pulumi.Output<string>;
    /**
     * The inputs for the Job.
     */
    public readonly input!: pulumi.Output<outputs.media.latest.JobInputClipResponse | outputs.media.latest.JobInputsResponse>;
    /**
     * The UTC date and time when the customer has last updated the Job, in 'YYYY-MM-DDThh:mm:ssZ' format.
     */
    public /*out*/ readonly lastModified!: pulumi.Output<string>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The outputs for the Job.
     */
    public readonly outputs!: pulumi.Output<outputs.media.latest.JobOutputAssetResponse[]>;
    /**
     * Priority with which the job should be processed. Higher priority jobs are processed before lower priority jobs. If not set, the default is normal.
     */
    public readonly priority!: pulumi.Output<string | undefined>;
    /**
     * The UTC date and time at which this Job began processing.
     */
    public /*out*/ readonly startTime!: pulumi.Output<string>;
    /**
     * The current state of the job.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Job resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: JobArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.accountName === undefined) {
                throw new Error("Missing required property 'accountName'");
            }
            if (!args || args.input === undefined) {
                throw new Error("Missing required property 'input'");
            }
            if (!args || args.jobName === undefined) {
                throw new Error("Missing required property 'jobName'");
            }
            if (!args || args.outputs === undefined) {
                throw new Error("Missing required property 'outputs'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.transformName === undefined) {
                throw new Error("Missing required property 'transformName'");
            }
            inputs["accountName"] = args ? args.accountName : undefined;
            inputs["correlationData"] = args ? args.correlationData : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["input"] = args ? args.input : undefined;
            inputs["jobName"] = args ? args.jobName : undefined;
            inputs["outputs"] = args ? args.outputs : undefined;
            inputs["priority"] = args ? args.priority : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["transformName"] = args ? args.transformName : undefined;
            inputs["created"] = undefined /*out*/;
            inputs["endTime"] = undefined /*out*/;
            inputs["lastModified"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["startTime"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["correlationData"] = undefined /*out*/;
            inputs["created"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["endTime"] = undefined /*out*/;
            inputs["input"] = undefined /*out*/;
            inputs["lastModified"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["outputs"] = undefined /*out*/;
            inputs["priority"] = undefined /*out*/;
            inputs["startTime"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:media/v20180330preview:Job" }, { type: "azure-nextgen:media/v20180601preview:Job" }, { type: "azure-nextgen:media/v20180701:Job" }, { type: "azure-nextgen:media/v20200501:Job" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Job.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Job resource.
 */
export interface JobArgs {
    /**
     * The Media Services account name.
     */
    readonly accountName: pulumi.Input<string>;
    /**
     * Customer provided key, value pairs that will be returned in Job and JobOutput state events.
     */
    readonly correlationData?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Optional customer supplied description of the Job.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The inputs for the Job.
     */
    readonly input: pulumi.Input<inputs.media.latest.JobInputClip | inputs.media.latest.JobInputs>;
    /**
     * The Job name.
     */
    readonly jobName: pulumi.Input<string>;
    /**
     * The outputs for the Job.
     */
    readonly outputs: pulumi.Input<pulumi.Input<inputs.media.latest.JobOutputAsset>[]>;
    /**
     * Priority with which the job should be processed. Higher priority jobs are processed before lower priority jobs. If not set, the default is normal.
     */
    readonly priority?: pulumi.Input<string>;
    /**
     * The name of the resource group within the Azure subscription.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The Transform name.
     */
    readonly transformName: pulumi.Input<string>;
}
