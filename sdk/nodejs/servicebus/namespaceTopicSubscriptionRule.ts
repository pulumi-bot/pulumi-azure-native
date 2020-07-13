// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Description of Rule Resource.
 */
export class NamespaceTopicSubscriptionRule extends pulumi.CustomResource {
    /**
     * Get an existing NamespaceTopicSubscriptionRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): NamespaceTopicSubscriptionRule {
        return new NamespaceTopicSubscriptionRule(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:servicebus:NamespaceTopicSubscriptionRule';

    /**
     * Returns true if the given object is an instance of NamespaceTopicSubscriptionRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NamespaceTopicSubscriptionRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NamespaceTopicSubscriptionRule.__pulumiType;
    }

    /**
     * Resource name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Properties of Rule resource
     */
    public readonly properties!: pulumi.Output<outputs.servicebus.RulepropertiesResponse>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a NamespaceTopicSubscriptionRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: NamespaceTopicSubscriptionRuleArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.namespaceName === undefined) {
                throw new Error("Missing required property 'namespaceName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.subscriptionName === undefined) {
                throw new Error("Missing required property 'subscriptionName'");
            }
            if (!args || args.topicName === undefined) {
                throw new Error("Missing required property 'topicName'");
            }
        inputs["name"] = args ? args.name : undefined;
        inputs["namespaceName"] = args ? args.namespaceName : undefined;
        inputs["properties"] = args ? args.properties : undefined;
        inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
        inputs["subscriptionName"] = args ? args.subscriptionName : undefined;
        inputs["topicName"] = args ? args.topicName : undefined;
        inputs["type"] = undefined /*out*/;
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(NamespaceTopicSubscriptionRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a NamespaceTopicSubscriptionRule resource.
 */
export interface NamespaceTopicSubscriptionRuleArgs {
    /**
     * The rule name.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The namespace name
     */
    readonly namespaceName: pulumi.Input<string>;
    /**
     * Properties of Rule resource
     */
    readonly properties?: pulumi.Input<inputs.servicebus.Ruleproperties>;
    /**
     * Name of the Resource group within the Azure subscription.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The subscription name.
     */
    readonly subscriptionName: pulumi.Input<string>;
    /**
     * The topic name.
     */
    readonly topicName: pulumi.Input<string>;
}