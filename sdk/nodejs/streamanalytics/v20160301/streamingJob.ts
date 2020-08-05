// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * A streaming job object, containing all information associated with the named streaming job.
 */
export class StreamingJob extends pulumi.CustomResource {
    /**
     * Get an existing StreamingJob resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): StreamingJob {
        return new StreamingJob(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:streamanalytics/v20160301:StreamingJob';

    /**
     * Returns true if the given object is an instance of StreamingJob.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StreamingJob {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StreamingJob.__pulumiType;
    }

    /**
     * Resource location. Required on PUT (CreateOrReplace) requests.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The properties that are associated with a streaming job.  Required on PUT (CreateOrReplace) requests.
     */
    public /*out*/ readonly properties!: pulumi.Output<outputs.streamanalytics.v20160301.StreamingJobPropertiesResponse>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a StreamingJob resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StreamingJobArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StreamingJobArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as StreamingJobArgs | undefined;
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["compatibilityLevel"] = args ? args.compatibilityLevel : undefined;
            inputs["dataLocale"] = args ? args.dataLocale : undefined;
            inputs["eventsLateArrivalMaxDelayInSeconds"] = args ? args.eventsLateArrivalMaxDelayInSeconds : undefined;
            inputs["eventsOutOfOrderMaxDelayInSeconds"] = args ? args.eventsOutOfOrderMaxDelayInSeconds : undefined;
            inputs["eventsOutOfOrderPolicy"] = args ? args.eventsOutOfOrderPolicy : undefined;
            inputs["functions"] = args ? args.functions : undefined;
            inputs["inputs"] = args ? args.inputs : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["outputErrorPolicy"] = args ? args.outputErrorPolicy : undefined;
            inputs["outputStartMode"] = args ? args.outputStartMode : undefined;
            inputs["outputStartTime"] = args ? args.outputStartTime : undefined;
            inputs["outputs"] = args ? args.outputs : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["transformation"] = args ? args.transformation : undefined;
            inputs["properties"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(StreamingJob.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a StreamingJob resource.
 */
export interface StreamingJobArgs {
    /**
     * Controls certain runtime behaviors of the streaming job.
     */
    readonly compatibilityLevel?: pulumi.Input<string>;
    /**
     * The data locale of the stream analytics job. Value should be the name of a supported .NET Culture from the set https://msdn.microsoft.com/en-us/library/system.globalization.culturetypes(v=vs.110).aspx. Defaults to 'en-US' if none specified.
     */
    readonly dataLocale?: pulumi.Input<string>;
    /**
     * The maximum tolerable delay in seconds where events arriving late could be included.  Supported range is -1 to 1814399 (20.23:59:59 days) and -1 is used to specify wait indefinitely. If the property is absent, it is interpreted to have a value of -1.
     */
    readonly eventsLateArrivalMaxDelayInSeconds?: pulumi.Input<number>;
    /**
     * The maximum tolerable delay in seconds where out-of-order events can be adjusted to be back in order.
     */
    readonly eventsOutOfOrderMaxDelayInSeconds?: pulumi.Input<number>;
    /**
     * Indicates the policy to apply to events that arrive out of order in the input event stream.
     */
    readonly eventsOutOfOrderPolicy?: pulumi.Input<string>;
    /**
     * A list of one or more functions for the streaming job. The name property for each function is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual transformation.
     */
    readonly functions?: pulumi.Input<pulumi.Input<inputs.streamanalytics.v20160301.Function>[]>;
    /**
     * A list of one or more inputs to the streaming job. The name property for each input is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual input.
     */
    readonly inputs?: pulumi.Input<pulumi.Input<inputs.streamanalytics.v20160301.Input>[]>;
    /**
     * Resource location. Required on PUT (CreateOrReplace) requests.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the streaming job.
     */
    readonly name: pulumi.Input<string>;
    /**
     * Indicates the policy to apply to events that arrive at the output and cannot be written to the external storage due to being malformed (missing column values, column values of wrong type or size).
     */
    readonly outputErrorPolicy?: pulumi.Input<string>;
    /**
     * This property should only be utilized when it is desired that the job be started immediately upon creation. Value may be JobStartTime, CustomTime, or LastOutputEventTime to indicate whether the starting point of the output event stream should start whenever the job is started, start at a custom user time stamp specified via the outputStartTime property, or start from the last event output time.
     */
    readonly outputStartMode?: pulumi.Input<string>;
    /**
     * Value is either an ISO-8601 formatted time stamp that indicates the starting point of the output event stream, or null to indicate that the output event stream will start whenever the streaming job is started. This property must have a value if outputStartMode is set to CustomTime.
     */
    readonly outputStartTime?: pulumi.Input<string>;
    /**
     * A list of one or more outputs for the streaming job. The name property for each output is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual output.
     */
    readonly outputs?: pulumi.Input<pulumi.Input<inputs.streamanalytics.v20160301.Output>[]>;
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Describes the SKU of the streaming job. Required on PUT (CreateOrReplace) requests.
     */
    readonly sku?: pulumi.Input<inputs.streamanalytics.v20160301.Sku>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Indicates the query and the number of streaming units to use for the streaming job. The name property of the transformation is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual transformation.
     */
    readonly transformation?: pulumi.Input<inputs.streamanalytics.v20160301.Transformation>;
}
