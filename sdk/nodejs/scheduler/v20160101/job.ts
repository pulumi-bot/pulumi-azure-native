// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

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
    public static readonly __pulumiType = 'azurerm:scheduler/v20160101:Job';

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
     * Gets the job resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Gets or sets the job properties.
     */
    public readonly properties!: pulumi.Output<outputs.scheduler.v20160101.JobPropertiesResponse>;
    /**
     * Gets the job resource type.
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
            if (!args || args.jobCollectionName === undefined) {
                throw new Error("Missing required property 'jobCollectionName'");
            }
            if (!args || args.jobName === undefined) {
                throw new Error("Missing required property 'jobName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["jobCollectionName"] = args ? args.jobCollectionName : undefined;
            inputs["jobName"] = args ? args.jobName : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["name"] = undefined /*out*/;
            inputs["properties"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:scheduler/latest:Job" }, { type: "azurerm:scheduler/v20140801preview:Job" }, { type: "azurerm:scheduler/v20160301:Job" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Job.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Job resource.
 */
export interface JobArgs {
    /**
     * The job collection name.
     */
    readonly jobCollectionName: pulumi.Input<string>;
    /**
     * The job name.
     */
    readonly jobName: pulumi.Input<string>;
    /**
     * Gets or sets the job properties.
     */
    readonly properties?: pulumi.Input<inputs.scheduler.v20160101.JobProperties>;
    /**
     * The resource group name.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
