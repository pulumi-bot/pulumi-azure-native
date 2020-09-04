// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * The policy definition.
 *
 * ## Example Usage
 * ### Create or update a policy definition at management group level
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const policyDefinitionAtManagementGroup = new azurerm.management.v20180301.PolicyDefinitionAtManagementGroup("policyDefinitionAtManagementGroup", {
 *     description: "Force resource names to begin with given 'prefix' and/or end with given 'suffix'",
 *     displayName: "Enforce resource naming convention",
 *     managementGroupId: "MyManagementGroup",
 *     metadata: {
 *         category: "Naming",
 *     },
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
export class PolicyDefinitionAtManagementGroup extends pulumi.CustomResource {
    /**
     * Get an existing PolicyDefinitionAtManagementGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PolicyDefinitionAtManagementGroup {
        return new PolicyDefinitionAtManagementGroup(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:management/v20180301:PolicyDefinitionAtManagementGroup';

    /**
     * Returns true if the given object is an instance of PolicyDefinitionAtManagementGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PolicyDefinitionAtManagementGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PolicyDefinitionAtManagementGroup.__pulumiType;
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
     * The policy definition mode. Possible values are NotSpecified, Indexed, and All.
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
     * Create a PolicyDefinitionAtManagementGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PolicyDefinitionAtManagementGroupArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.managementGroupId === undefined) {
                throw new Error("Missing required property 'managementGroupId'");
            }
            if (!args || args.policyDefinitionName === undefined) {
                throw new Error("Missing required property 'policyDefinitionName'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["managementGroupId"] = args ? args.managementGroupId : undefined;
            inputs["metadata"] = args ? args.metadata : undefined;
            inputs["mode"] = args ? args.mode : undefined;
            inputs["parameters"] = args ? args.parameters : undefined;
            inputs["policyDefinitionName"] = args ? args.policyDefinitionName : undefined;
            inputs["policyRule"] = args ? args.policyRule : undefined;
            inputs["policyType"] = args ? args.policyType : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["description"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["metadata"] = undefined /*out*/;
            inputs["mode"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["parameters"] = undefined /*out*/;
            inputs["policyRule"] = undefined /*out*/;
            inputs["policyType"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:management/latest:PolicyDefinitionAtManagementGroup" }, { type: "azurerm:management/v20161201:PolicyDefinitionAtManagementGroup" }, { type: "azurerm:management/v20180501:PolicyDefinitionAtManagementGroup" }, { type: "azurerm:management/v20190101:PolicyDefinitionAtManagementGroup" }, { type: "azurerm:management/v20190601:PolicyDefinitionAtManagementGroup" }, { type: "azurerm:management/v20190901:PolicyDefinitionAtManagementGroup" }, { type: "azurerm:management/v20200301:PolicyDefinitionAtManagementGroup" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(PolicyDefinitionAtManagementGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a PolicyDefinitionAtManagementGroup resource.
 */
export interface PolicyDefinitionAtManagementGroupArgs {
    /**
     * The policy definition description.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The display name of the policy definition.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * The ID of the management group.
     */
    readonly managementGroupId: pulumi.Input<string>;
    /**
     * The policy definition metadata.
     */
    readonly metadata?: pulumi.Input<{[key: string]: any}>;
    /**
     * The policy definition mode. Possible values are NotSpecified, Indexed, and All.
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
