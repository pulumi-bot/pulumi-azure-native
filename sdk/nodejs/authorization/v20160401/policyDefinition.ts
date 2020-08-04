// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The policy definition.
 */
export class PolicyDefinition extends pulumi.CustomResource {
    /**
     * Get an existing PolicyDefinition resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PolicyDefinition {
        return new PolicyDefinition(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:authorization/v20160401:PolicyDefinition';

    /**
     * Returns true if the given object is an instance of PolicyDefinition.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PolicyDefinition {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PolicyDefinition.__pulumiType;
    }

    /**
     * The name of the policy definition. If you do not specify a value for name, the value is inferred from the name value in the request URI.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * The policy definition properties.
     */
    public readonly properties!: pulumi.Output<outputs.authorization.v20160401.PolicyDefinitionPropertiesResponse>;

    /**
     * Create a PolicyDefinition resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PolicyDefinitionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PolicyDefinitionArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as PolicyDefinitionArgs | undefined;
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            inputs["name"] = args ? args.name : undefined;
            inputs["properties"] = args ? args.properties : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(PolicyDefinition.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a PolicyDefinition resource.
 */
export interface PolicyDefinitionArgs {
    /**
     * The name of the policy definition to create.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The policy definition properties.
     */
    readonly properties?: pulumi.Input<inputs.authorization.v20160401.PolicyDefinitionProperties>;
}
