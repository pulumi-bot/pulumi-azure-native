// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Description of a namespace authorization rule.
 */
export class NamespaceHybridConnectionAuthorizationRule extends pulumi.CustomResource {
    /**
     * Get an existing NamespaceHybridConnectionAuthorizationRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): NamespaceHybridConnectionAuthorizationRule {
        return new NamespaceHybridConnectionAuthorizationRule(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:relay:NamespaceHybridConnectionAuthorizationRule';

    /**
     * Returns true if the given object is an instance of NamespaceHybridConnectionAuthorizationRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NamespaceHybridConnectionAuthorizationRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NamespaceHybridConnectionAuthorizationRule.__pulumiType;
    }

    /**
     * Resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Authorization rule properties.
     */
    public readonly properties!: pulumi.Output<outputs.relay.AuthorizationRuleResponseProperties>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a NamespaceHybridConnectionAuthorizationRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: NamespaceHybridConnectionAuthorizationRuleArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
            if (!args || args.hybridConnectionName === undefined) {
                throw new Error("Missing required property 'hybridConnectionName'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.namespaceName === undefined) {
                throw new Error("Missing required property 'namespaceName'");
            }
            if (!args || args.properties === undefined) {
                throw new Error("Missing required property 'properties'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
        inputs["hybridConnectionName"] = args ? args.hybridConnectionName : undefined;
        inputs["name"] = args ? args.name : undefined;
        inputs["namespaceName"] = args ? args.namespaceName : undefined;
        inputs["properties"] = args ? args.properties : undefined;
        inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
        inputs["type"] = undefined /*out*/;
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(NamespaceHybridConnectionAuthorizationRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a NamespaceHybridConnectionAuthorizationRule resource.
 */
export interface NamespaceHybridConnectionAuthorizationRuleArgs {
    /**
     * The hybrid connection name.
     */
    readonly hybridConnectionName: pulumi.Input<string>;
    /**
     * The authorization rule name.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The namespace name
     */
    readonly namespaceName: pulumi.Input<string>;
    /**
     * Authorization rule properties.
     */
    readonly properties: pulumi.Input<inputs.relay.AuthorizationRuleProperties>;
    /**
     * Name of the Resource group within the Azure subscription.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
