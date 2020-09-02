// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Pipeline resource type.
 *
 * ## Example Usage
 * ### Pipelines_Create
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const pipeline = new azurerm.datafactory.v20180601.Pipeline("pipeline", {
 *     activities: [{
 *         name: "ExampleForeachActivity",
 *         type: "ForEach",
 *     }],
 *     factoryName: "exampleFactoryName",
 *     parameters: {
 *         JobId: {
 *             type: "String",
 *         },
 *         OutputBlobNameList: {
 *             type: "Array",
 *         },
 *     },
 *     pipelineName: "examplePipeline",
 *     resourceGroupName: "exampleResourceGroup",
 *     runDimensions: {
 *         JobId: {
 *             type: "Expression",
 *             value: "@pipeline().parameters.JobId",
 *         },
 *     },
 *     variables: {
 *         TestVariableArray: {
 *             type: "Array",
 *         },
 *     },
 * });
 *
 * ```
 * ### Pipelines_Update
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const pipeline = new azurerm.datafactory.v20180601.Pipeline("pipeline", {
 *     activities: [{
 *         name: "ExampleForeachActivity",
 *         type: "ForEach",
 *     }],
 *     description: "Example description",
 *     factoryName: "exampleFactoryName",
 *     parameters: {
 *         OutputBlobNameList: {
 *             type: "Array",
 *         },
 *     },
 *     pipelineName: "examplePipeline",
 *     resourceGroupName: "exampleResourceGroup",
 * });
 *
 * ```
 */
export class Pipeline extends pulumi.CustomResource {
    /**
     * Get an existing Pipeline resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Pipeline {
        return new Pipeline(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:datafactory/v20180601:Pipeline';

    /**
     * Returns true if the given object is an instance of Pipeline.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Pipeline {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Pipeline.__pulumiType;
    }

    /**
     * List of activities in pipeline.
     */
    public readonly activities!: pulumi.Output<outputs.datafactory.v20180601.ActivityResponse[] | undefined>;
    /**
     * List of tags that can be used for describing the Pipeline.
     */
    public readonly annotations!: pulumi.Output<{[key: string]: any}[] | undefined>;
    /**
     * The max number of concurrent runs for the pipeline.
     */
    public readonly concurrency!: pulumi.Output<number | undefined>;
    /**
     * The description of the pipeline.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Etag identifies change in the resource.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The folder that this Pipeline is in. If not specified, Pipeline will appear at the root level.
     */
    public readonly folder!: pulumi.Output<outputs.datafactory.v20180601.PipelineResponseFolder | undefined>;
    /**
     * The resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * List of parameters for pipeline.
     */
    public readonly parameters!: pulumi.Output<{[key: string]: outputs.datafactory.v20180601.ParameterSpecificationResponse} | undefined>;
    /**
     * Dimensions emitted by Pipeline.
     */
    public readonly runDimensions!: pulumi.Output<{[key: string]: {[key: string]: any}} | undefined>;
    /**
     * The resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * List of variables for pipeline.
     */
    public readonly variables!: pulumi.Output<{[key: string]: outputs.datafactory.v20180601.VariableSpecificationResponse} | undefined>;

    /**
     * Create a Pipeline resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PipelineArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PipelineArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as PipelineArgs | undefined;
            if (!args || args.factoryName === undefined) {
                throw new Error("Missing required property 'factoryName'");
            }
            if (!args || args.pipelineName === undefined) {
                throw new Error("Missing required property 'pipelineName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["activities"] = args ? args.activities : undefined;
            inputs["annotations"] = args ? args.annotations : undefined;
            inputs["concurrency"] = args ? args.concurrency : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["factoryName"] = args ? args.factoryName : undefined;
            inputs["folder"] = args ? args.folder : undefined;
            inputs["parameters"] = args ? args.parameters : undefined;
            inputs["pipelineName"] = args ? args.pipelineName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["runDimensions"] = args ? args.runDimensions : undefined;
            inputs["variables"] = args ? args.variables : undefined;
            inputs["etag"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:datafactory/latest:Pipeline" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Pipeline.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Pipeline resource.
 */
export interface PipelineArgs {
    /**
     * List of activities in pipeline.
     */
    readonly activities?: pulumi.Input<pulumi.Input<inputs.datafactory.v20180601.Activity>[]>;
    /**
     * List of tags that can be used for describing the Pipeline.
     */
    readonly annotations?: pulumi.Input<pulumi.Input<{[key: string]: any}>[]>;
    /**
     * The max number of concurrent runs for the pipeline.
     */
    readonly concurrency?: pulumi.Input<number>;
    /**
     * The description of the pipeline.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The factory name.
     */
    readonly factoryName: pulumi.Input<string>;
    /**
     * The folder that this Pipeline is in. If not specified, Pipeline will appear at the root level.
     */
    readonly folder?: pulumi.Input<inputs.datafactory.v20180601.PipelineFolder>;
    /**
     * List of parameters for pipeline.
     */
    readonly parameters?: pulumi.Input<{[key: string]: pulumi.Input<inputs.datafactory.v20180601.ParameterSpecification>}>;
    /**
     * The pipeline name.
     */
    readonly pipelineName: pulumi.Input<string>;
    /**
     * The resource group name.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Dimensions emitted by Pipeline.
     */
    readonly runDimensions?: pulumi.Input<{[key: string]: pulumi.Input<{[key: string]: any}>}>;
    /**
     * List of variables for pipeline.
     */
    readonly variables?: pulumi.Input<{[key: string]: pulumi.Input<inputs.datafactory.v20180601.VariableSpecification>}>;
}
