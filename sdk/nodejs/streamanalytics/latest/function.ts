// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * A function object, containing all information associated with the named function. All functions are contained under a streaming job.
 *
 * ## Example Usage
 * ### Create a JavaScript function
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const _function = new azurerm.streamanalytics.latest.Function("function", {
 *     functionName: "function8197",
 *     jobName: "sj8653",
 *     properties: {},
 *     resourceGroupName: "sjrg1637",
 * });
 *
 * ```
 * ### Create an Azure ML function
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const _function = new azurerm.streamanalytics.latest.Function("function", {
 *     functionName: "function588",
 *     jobName: "sj9093",
 *     properties: {},
 *     resourceGroupName: "sjrg7",
 * });
 *
 * ```
 */
export class Function extends pulumi.CustomResource {
    /**
     * Get an existing Function resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Function {
        return new Function(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:streamanalytics/latest:Function';

    /**
     * Returns true if the given object is an instance of Function.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Function {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Function.__pulumiType;
    }

    /**
     * Resource name
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * The properties that are associated with a function.
     */
    public readonly properties!: pulumi.Output<outputs.streamanalytics.latest.FunctionPropertiesResponse>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Function resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FunctionArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.functionName === undefined) {
                throw new Error("Missing required property 'functionName'");
            }
            if (!args || args.jobName === undefined) {
                throw new Error("Missing required property 'jobName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["functionName"] = args ? args.functionName : undefined;
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
        const aliasOpts = { aliases: [{ type: "azurerm:streamanalytics/v20160301:Function" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Function.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Function resource.
 */
export interface FunctionArgs {
    /**
     * The name of the function.
     */
    readonly functionName: pulumi.Input<string>;
    /**
     * The name of the streaming job.
     */
    readonly jobName: pulumi.Input<string>;
    /**
     * Resource name
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The properties that are associated with a function.
     */
    readonly properties?: pulumi.Input<inputs.streamanalytics.latest.FunctionProperties>;
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
