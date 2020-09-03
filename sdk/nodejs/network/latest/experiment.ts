// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Defines the properties of an Experiment
 *
 * ## Example Usage
 * ### Creates an Experiment
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const experiment = new azurerm.network.latest.Experiment("experiment", {
 *     description: "this is my first experiment!",
 *     enabledState: "Enabled",
 *     endpointA: {
 *         name: "endpoint A",
 *     },
 *     endpointB: {
 *         name: "endpoint B",
 *     },
 *     experimentName: "MyExperiment",
 *     profileName: "MyProfile",
 *     resourceGroupName: "MyResourceGroup",
 * });
 *
 * ```
 */
export class Experiment extends pulumi.CustomResource {
    /**
     * Get an existing Experiment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Experiment {
        return new Experiment(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/latest:Experiment';

    /**
     * Returns true if the given object is an instance of Experiment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Experiment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Experiment.__pulumiType;
    }

    /**
     * The description of the details or intents of the Experiment
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The state of the Experiment
     */
    public readonly enabledState!: pulumi.Output<string | undefined>;
    /**
     * The endpoint A of an experiment
     */
    public readonly endpointA!: pulumi.Output<outputs.network.latest.EndpointResponse | undefined>;
    /**
     * The endpoint B of an experiment
     */
    public readonly endpointB!: pulumi.Output<outputs.network.latest.EndpointResponse | undefined>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Resource status.
     */
    public readonly resourceState!: pulumi.Output<string | undefined>;
    /**
     * The uri to the Script used in the Experiment
     */
    public /*out*/ readonly scriptFileUri!: pulumi.Output<string>;
    /**
     * The description of Experiment status from the server side
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Experiment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ExperimentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ExperimentArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as ExperimentArgs | undefined;
            if (!args || args.experimentName === undefined) {
                throw new Error("Missing required property 'experimentName'");
            }
            if (!args || args.profileName === undefined) {
                throw new Error("Missing required property 'profileName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["enabledState"] = args ? args.enabledState : undefined;
            inputs["endpointA"] = args ? args.endpointA : undefined;
            inputs["endpointB"] = args ? args.endpointB : undefined;
            inputs["experimentName"] = args ? args.experimentName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["profileName"] = args ? args.profileName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceState"] = args ? args.resourceState : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["scriptFileUri"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:network/v20191101:Experiment" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Experiment.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Experiment resource.
 */
export interface ExperimentArgs {
    /**
     * The description of the details or intents of the Experiment
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The state of the Experiment
     */
    readonly enabledState?: pulumi.Input<string>;
    /**
     * The endpoint A of an experiment
     */
    readonly endpointA?: pulumi.Input<inputs.network.latest.Endpoint>;
    /**
     * The endpoint B of an experiment
     */
    readonly endpointB?: pulumi.Input<inputs.network.latest.Endpoint>;
    /**
     * The Experiment identifier associated with the Experiment
     */
    readonly experimentName: pulumi.Input<string>;
    /**
     * Resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The Profile identifier associated with the Tenant and Partner
     */
    readonly profileName: pulumi.Input<string>;
    /**
     * Name of the Resource group within the Azure subscription.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Resource status.
     */
    readonly resourceState?: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
