// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Function information.
 */
export class AppServiceFunction extends pulumi.CustomResource {
    /**
     * Get an existing AppServiceFunction resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AppServiceFunction {
        return new AppServiceFunction(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:web:AppServiceFunction';

    /**
     * Returns true if the given object is an instance of AppServiceFunction.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AppServiceFunction {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AppServiceFunction.__pulumiType;
    }

    /**
     * Kind of resource.
     */
    public readonly kind!: pulumi.Output<string | undefined>;
    /**
     * Resource Name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * FunctionEnvelope resource specific properties
     */
    public readonly properties!: pulumi.Output<outputs.web.FunctionEnvelopeResponseProperties>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a AppServiceFunction resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AppServiceFunctionArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
            if (!args || args.functionName === undefined) {
                throw new Error("Missing required property 'functionName'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
        inputs["functionName"] = args ? args.functionName : undefined;
        inputs["kind"] = args ? args.kind : undefined;
        inputs["name"] = args ? args.name : undefined;
        inputs["properties"] = args ? args.properties : undefined;
        inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
        inputs["type"] = undefined /*out*/;
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AppServiceFunction.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a AppServiceFunction resource.
 */
export interface AppServiceFunctionArgs {
    /**
     * Function name.
     */
    readonly functionName: pulumi.Input<string>;
    /**
     * Kind of resource.
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * Site name.
     */
    readonly name: pulumi.Input<string>;
    /**
     * FunctionEnvelope resource specific properties
     */
    readonly properties?: pulumi.Input<inputs.web.FunctionEnvelopeProperties>;
    /**
     * Name of the resource group to which the resource belongs.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
