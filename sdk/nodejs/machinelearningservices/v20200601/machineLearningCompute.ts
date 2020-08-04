// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Machine Learning compute object wrapped into ARM resource envelope.
 */
export class MachineLearningCompute extends pulumi.CustomResource {
    /**
     * Get an existing MachineLearningCompute resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): MachineLearningCompute {
        return new MachineLearningCompute(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:machinelearningservices/v20200601:MachineLearningCompute';

    /**
     * Returns true if the given object is an instance of MachineLearningCompute.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MachineLearningCompute {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MachineLearningCompute.__pulumiType;
    }

    /**
     * The identity of the resource.
     */
    public readonly identity!: pulumi.Output<outputs.machinelearningservices.v20200601.IdentityResponse | undefined>;
    /**
     * Specifies the location of the resource.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Specifies the name of the resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Compute properties
     */
    public readonly properties!: pulumi.Output<outputs.machinelearningservices.v20200601.ComputeResponse>;
    /**
     * The sku of the workspace.
     */
    public readonly sku!: pulumi.Output<outputs.machinelearningservices.v20200601.SkuResponse | undefined>;
    /**
     * Contains resource tags defined as key/value pairs.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Specifies the type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a MachineLearningCompute resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MachineLearningComputeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MachineLearningComputeArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as MachineLearningComputeArgs | undefined;
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.workspaceName === undefined) {
                throw new Error("Missing required property 'workspaceName'");
            }
            inputs["identity"] = args ? args.identity : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["workspaceName"] = args ? args.workspaceName : undefined;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(MachineLearningCompute.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a MachineLearningCompute resource.
 */
export interface MachineLearningComputeArgs {
    /**
     * The identity of the resource.
     */
    readonly identity?: pulumi.Input<inputs.machinelearningservices.v20200601.Identity>;
    /**
     * Specifies the location of the resource.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Name of the Azure Machine Learning compute.
     */
    readonly name: pulumi.Input<string>;
    /**
     * Compute properties
     */
    readonly properties?: pulumi.Input<inputs.machinelearningservices.v20200601.Compute>;
    /**
     * Name of the resource group in which workspace is located.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The sku of the workspace.
     */
    readonly sku?: pulumi.Input<inputs.machinelearningservices.v20200601.Sku>;
    /**
     * Contains resource tags defined as key/value pairs.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name of Azure Machine Learning workspace.
     */
    readonly workspaceName: pulumi.Input<string>;
}
