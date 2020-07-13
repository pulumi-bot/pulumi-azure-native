// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class JobCollection extends pulumi.CustomResource {
    /**
     * Get an existing JobCollection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): JobCollection {
        return new JobCollection(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:scheduler:JobCollection';

    /**
     * Returns true if the given object is an instance of JobCollection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is JobCollection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === JobCollection.__pulumiType;
    }

    /**
     * Gets or sets the storage account location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Gets or sets the job collection resource name.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * Gets or sets the job collection properties.
     */
    public readonly properties!: pulumi.Output<outputs.scheduler.JobCollectionPropertiesResponse>;
    /**
     * Gets or sets the tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Gets the job collection resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a JobCollection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: JobCollectionArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
            if (!args || args.jobCollectionName === undefined) {
                throw new Error("Missing required property 'jobCollectionName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
        inputs["jobCollectionName"] = args ? args.jobCollectionName : undefined;
        inputs["location"] = args ? args.location : undefined;
        inputs["name"] = args ? args.name : undefined;
        inputs["properties"] = args ? args.properties : undefined;
        inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
        inputs["tags"] = args ? args.tags : undefined;
        inputs["type"] = undefined /*out*/;
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(JobCollection.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a JobCollection resource.
 */
export interface JobCollectionArgs {
    /**
     * The job collection name.
     */
    readonly jobCollectionName: pulumi.Input<string>;
    /**
     * Gets or sets the storage account location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Gets or sets the job collection resource name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Gets or sets the job collection properties.
     */
    readonly properties?: pulumi.Input<inputs.scheduler.JobCollectionProperties>;
    /**
     * The resource group name.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Gets or sets the tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
