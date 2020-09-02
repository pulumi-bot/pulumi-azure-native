// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * The policy definition.
 *
 * ## Example Usage
 * ### Create or update a policy definition
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const policyDefinition = new azurerm.authorization.v20190601.PolicyDefinition("policyDefinition", {
 *     description: "Force resource names to begin with given 'prefix' and/or end with given 'suffix'",
 *     displayName: "Enforce resource naming convention",
 *     metadata: {
 *         category: "Naming",
 *     },
 *     mode: "All",
 *     parameters: {
 *         prefix: {
 *             metadata: {
 *                 description: "Resource name prefix",
 *                 displayName: "Prefix",
 *             },
 *             type: "String",
 *         },
 *         suffix: {
 *             metadata: {
 *                 description: "Resource name suffix",
 *                 displayName: "Suffix",
 *             },
 *             type: "String",
 *         },
 *     },
 *     policyDefinitionName: "ResourceNaming",
 *     policyRule: {
 *         "if": {
 *             not: {
 *                 field: "name",
 *                 like: "[concat(parameters('prefix'), '*', parameters('suffix'))]",
 *             },
 *         },
 *         then: {
 *             effect: "deny",
 *         },
 *     },
 * });
 *
 * ```
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
    public static readonly __pulumiType = 'azurerm:authorization/v20190601:PolicyDefinition';

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
     * The policy definition description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The display name of the policy definition.
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * The policy definition metadata.
     */
    public readonly metadata!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The policy definition mode. Some examples are All, Indexed, Microsoft.KeyVault.Data.
     */
    public readonly mode!: pulumi.Output<string | undefined>;
    /**
     * The name of the policy definition.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Required if a parameter is used in policy rule.
     */
    public readonly parameters!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The policy rule.
     */
    public readonly policyRule!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The type of policy definition. Possible values are NotSpecified, BuiltIn, and Custom.
     */
    public readonly policyType!: pulumi.Output<string | undefined>;
    /**
     * The type of the resource (Microsoft.Authorization/policyDefinitions).
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

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
            if (!args || args.policyDefinitionName === undefined) {
                throw new Error("Missing required property 'policyDefinitionName'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["metadata"] = args ? args.metadata : undefined;
            inputs["mode"] = args ? args.mode : undefined;
            inputs["parameters"] = args ? args.parameters : undefined;
            inputs["policyDefinitionName"] = args ? args.policyDefinitionName : undefined;
            inputs["policyRule"] = args ? args.policyRule : undefined;
            inputs["policyType"] = args ? args.policyType : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:authorization/latest:PolicyDefinition" }, { type: "azurerm:authorization/v20151101:PolicyDefinition" }, { type: "azurerm:authorization/v20160401:PolicyDefinition" }, { type: "azurerm:authorization/v20161201:PolicyDefinition" }, { type: "azurerm:authorization/v20180301:PolicyDefinition" }, { type: "azurerm:authorization/v20180501:PolicyDefinition" }, { type: "azurerm:authorization/v20190101:PolicyDefinition" }, { type: "azurerm:authorization/v20190901:PolicyDefinition" }, { type: "azurerm:authorization/v20200301:PolicyDefinition" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(PolicyDefinition.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a PolicyDefinition resource.
 */
export interface PolicyDefinitionArgs {
    /**
     * The policy definition description.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The display name of the policy definition.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * The policy definition metadata.
     */
    readonly metadata?: pulumi.Input<{[key: string]: any}>;
    /**
     * The policy definition mode. Some examples are All, Indexed, Microsoft.KeyVault.Data.
     */
    readonly mode?: pulumi.Input<string>;
    /**
     * Required if a parameter is used in policy rule.
     */
    readonly parameters?: pulumi.Input<{[key: string]: any}>;
    /**
     * The name of the policy definition to create.
     */
    readonly policyDefinitionName: pulumi.Input<string>;
    /**
     * The policy rule.
     */
    readonly policyRule?: pulumi.Input<{[key: string]: any}>;
    /**
     * The type of policy definition. Possible values are NotSpecified, BuiltIn, and Custom.
     */
    readonly policyType?: pulumi.Input<string>;
}
