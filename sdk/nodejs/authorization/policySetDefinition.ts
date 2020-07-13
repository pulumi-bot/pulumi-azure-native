// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * The policy set definition.
 */
export class PolicySetDefinition extends pulumi.CustomResource {
    /**
     * Get an existing PolicySetDefinition resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PolicySetDefinition {
        return new PolicySetDefinition(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:authorization:PolicySetDefinition';

    /**
     * Returns true if the given object is an instance of PolicySetDefinition.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PolicySetDefinition {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PolicySetDefinition.__pulumiType;
    }

    /**
     * The name of the policy set definition.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The policy definition properties.
     */
    public readonly properties!: pulumi.Output<outputs.authorization.PolicySetDefinitionPropertiesResponse>;
    /**
     * The type of the resource (Microsoft.Authorization/policySetDefinitions).
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a PolicySetDefinition resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: PolicySetDefinitionArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
            if (!args || args.managementGroupId === undefined) {
                throw new Error("Missing required property 'managementGroupId'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
        inputs["managementGroupId"] = args ? args.managementGroupId : undefined;
        inputs["name"] = args ? args.name : undefined;
        inputs["properties"] = args ? args.properties : undefined;
        inputs["type"] = undefined /*out*/;
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(PolicySetDefinition.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a PolicySetDefinition resource.
 */
export interface PolicySetDefinitionArgs {
    /**
     * The ID of the management group.
     */
    readonly managementGroupId: pulumi.Input<string>;
    /**
     * The name of the policy set definition to create.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The policy definition properties.
     */
    readonly properties?: pulumi.Input<inputs.authorization.PolicySetDefinitionProperties>;
}
