// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Description of a namespace authorization rule.
 */
export class QueueAuthorizationRule extends pulumi.CustomResource {
    /**
     * Get an existing QueueAuthorizationRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): QueueAuthorizationRule {
        return new QueueAuthorizationRule(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:servicebus/v20140901:QueueAuthorizationRule';

    /**
     * Returns true if the given object is an instance of QueueAuthorizationRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is QueueAuthorizationRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === QueueAuthorizationRule.__pulumiType;
    }

    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * AuthorizationRule properties.
     */
    public /*out*/ readonly properties!: pulumi.Output<outputs.servicebus.v20140901.SharedAccessAuthorizationRulePropertiesResponse>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a QueueAuthorizationRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: QueueAuthorizationRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: QueueAuthorizationRuleArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as QueueAuthorizationRuleArgs | undefined;
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.namespaceName === undefined) {
                throw new Error("Missing required property 'namespaceName'");
            }
            if (!args || args.queueName === undefined) {
                throw new Error("Missing required property 'queueName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.rights === undefined) {
                throw new Error("Missing required property 'rights'");
            }
            inputs["claimType"] = args ? args.claimType : undefined;
            inputs["claimValue"] = args ? args.claimValue : undefined;
            inputs["keyName"] = args ? args.keyName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namespaceName"] = args ? args.namespaceName : undefined;
            inputs["primaryKey"] = args ? args.primaryKey : undefined;
            inputs["queueName"] = args ? args.queueName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["rights"] = args ? args.rights : undefined;
            inputs["secondaryKey"] = args ? args.secondaryKey : undefined;
            inputs["properties"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(QueueAuthorizationRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a QueueAuthorizationRule resource.
 */
export interface QueueAuthorizationRuleArgs {
    /**
     * A string that describes Claim Type for authorization rule.
     */
    readonly claimType?: pulumi.Input<string>;
    /**
     * A string that describes Claim Value of authorization rule.
     */
    readonly claimValue?: pulumi.Input<string>;
    /**
     * A string that describes the Key Name of authorization rule.
     */
    readonly keyName?: pulumi.Input<string>;
    /**
     * data center location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The authorization rule name.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The namespace name
     */
    readonly namespaceName: pulumi.Input<string>;
    /**
     * A base64-encoded 256-bit primary key for signing and validating the SAS token.
     */
    readonly primaryKey?: pulumi.Input<string>;
    /**
     * The queue name.
     */
    readonly queueName: pulumi.Input<string>;
    /**
     * Name of the Resource group within the Azure subscription.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The rights associated with the rule.
     */
    readonly rights: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A base64-encoded 256-bit primary key for signing and validating the SAS token.
     */
    readonly secondaryKey?: pulumi.Input<string>;
}
