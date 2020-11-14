// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * An input object, containing all information associated with the named input. All inputs are contained under a streaming job.
 */
export class Input extends pulumi.CustomResource {
    /**
     * Get an existing Input resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Input {
        return new Input(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:streamanalytics/v20160301:Input';

    /**
     * Returns true if the given object is an instance of Input.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Input {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Input.__pulumiType;
    }

    /**
     * Resource name
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * The properties that are associated with an input. Required on PUT (CreateOrReplace) requests.
     */
    public readonly properties!: pulumi.Output<outputs.streamanalytics.v20160301.ReferenceInputPropertiesResponse | outputs.streamanalytics.v20160301.StreamInputPropertiesResponse>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Input resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InputArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.inputName === undefined) {
                throw new Error("Missing required property 'inputName'");
            }
            if (!args || args.jobName === undefined) {
                throw new Error("Missing required property 'jobName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["inputName"] = args ? args.inputName : undefined;
            inputs["jobName"] = args ? args.jobName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
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
        const aliasOpts = { aliases: [{ type: "azure-nextgen:streamanalytics/latest:Input" }, { type: "azure-nextgen:streamanalytics/v20170401preview:Input" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Input.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Input resource.
 */
export interface InputArgs {
    /**
     * The name of the input.
     */
    readonly inputName: pulumi.Input<string>;
    /**
     * The name of the streaming job.
     */
    readonly jobName: pulumi.Input<string>;
    /**
     * Resource name
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The properties that are associated with an input. Required on PUT (CreateOrReplace) requests.
     */
    readonly properties?: pulumi.Input<inputs.streamanalytics.v20160301.ReferenceInputProperties | inputs.streamanalytics.v20160301.StreamInputProperties>;
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
