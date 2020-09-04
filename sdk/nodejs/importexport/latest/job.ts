// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Contains the job information.
 *
 * ## Example Usage
 * ### Create export job
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const job = new azurerm.importexport.latest.Job("job", {
 *     jobName: "myExportJob",
 *     location: "West US",
 *     resourceGroupName: "myResourceGroup",
 * });
 *
 * ```
 * ### Create import job
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const job = new azurerm.importexport.latest.Job("job", {
 *     jobName: "myJob",
 *     location: "West US",
 *     resourceGroupName: "myResourceGroup",
 * });
 *
 * ```
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
    public static readonly __pulumiType = 'azurerm:importexport/latest:Job';

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
     * Specifies the job identity details
     */
    public /*out*/ readonly identity!: pulumi.Output<outputs.importexport.latest.IdentityDetailsResponse | undefined>;
    /**
     * Specifies the Azure location where the job is created.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Specifies the name of the job.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Specifies the job properties
     */
    public readonly properties!: pulumi.Output<outputs.importexport.latest.JobDetailsResponse>;
    /**
     * Specifies the tags that are assigned to the job.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Specifies the type of the job resource.
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
            if (!args || args.jobName === undefined) {
                throw new Error("Missing required property 'jobName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["jobName"] = args ? args.jobName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["identity"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["identity"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["properties"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:importexport/v20161101:Job" }, { type: "azurerm:importexport/v20200801:Job" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Job.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Job resource.
 */
export interface JobArgs {
    /**
     * The name of the import/export job.
     */
    readonly jobName: pulumi.Input<string>;
    /**
     * Specifies the supported Azure location where the job should be created
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the job properties
     */
    readonly properties?: pulumi.Input<inputs.importexport.latest.JobDetails>;
    /**
     * The resource group name uniquely identifies the resource group within the user subscription.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Specifies the tags that will be assigned to the job.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
