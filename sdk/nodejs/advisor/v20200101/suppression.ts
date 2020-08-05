// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The details of the snoozed or dismissed rule; for example, the duration, name, and GUID associated with the rule.
 */
export class Suppression extends pulumi.CustomResource {
    /**
     * Get an existing Suppression resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Suppression {
        return new Suppression(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:advisor/v20200101:Suppression';

    /**
     * Returns true if the given object is an instance of Suppression.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Suppression {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Suppression.__pulumiType;
    }

    /**
     * The name of the resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The properties of the suppression.
     */
    public /*out*/ readonly properties!: pulumi.Output<outputs.advisor.v20200101.SuppressionPropertiesResponse>;
    /**
     * The type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Suppression resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SuppressionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SuppressionArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as SuppressionArgs | undefined;
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.recommendationId === undefined) {
                throw new Error("Missing required property 'recommendationId'");
            }
            if (!args || args.resourceUri === undefined) {
                throw new Error("Missing required property 'resourceUri'");
            }
            inputs["name"] = args ? args.name : undefined;
            inputs["recommendationId"] = args ? args.recommendationId : undefined;
            inputs["resourceUri"] = args ? args.resourceUri : undefined;
            inputs["suppressionId"] = args ? args.suppressionId : undefined;
            inputs["ttl"] = args ? args.ttl : undefined;
            inputs["properties"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Suppression.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Suppression resource.
 */
export interface SuppressionArgs {
    /**
     * The name of the suppression.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The recommendation ID.
     */
    readonly recommendationId: pulumi.Input<string>;
    /**
     * The fully qualified Azure Resource Manager identifier of the resource to which the recommendation applies.
     */
    readonly resourceUri: pulumi.Input<string>;
    /**
     * The GUID of the suppression.
     */
    readonly suppressionId?: pulumi.Input<string>;
    /**
     * The duration for which the suppression is valid.
     */
    readonly ttl?: pulumi.Input<string>;
}
