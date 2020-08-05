// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The policy set definition.
 */
export class PolicySetDefinitionAtManagementGroup extends pulumi.CustomResource {
    /**
     * Get an existing PolicySetDefinitionAtManagementGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PolicySetDefinitionAtManagementGroup {
        return new PolicySetDefinitionAtManagementGroup(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:management/v20190101:PolicySetDefinitionAtManagementGroup';

    /**
     * Returns true if the given object is an instance of PolicySetDefinitionAtManagementGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PolicySetDefinitionAtManagementGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PolicySetDefinitionAtManagementGroup.__pulumiType;
    }

    /**
     * The name of the policy set definition.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The policy definition properties.
     */
    public /*out*/ readonly properties!: pulumi.Output<outputs.management.v20190101.PolicySetDefinitionPropertiesResponse>;
    /**
     * The type of the resource (Microsoft.Authorization/policySetDefinitions).
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a PolicySetDefinitionAtManagementGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PolicySetDefinitionAtManagementGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PolicySetDefinitionAtManagementGroupArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as PolicySetDefinitionAtManagementGroupArgs | undefined;
            if (!args || args.managementGroupId === undefined) {
                throw new Error("Missing required property 'managementGroupId'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.policyDefinitions === undefined) {
                throw new Error("Missing required property 'policyDefinitions'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["managementGroupId"] = args ? args.managementGroupId : undefined;
            inputs["metadata"] = args ? args.metadata : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["parameters"] = args ? args.parameters : undefined;
            inputs["policyDefinitions"] = args ? args.policyDefinitions : undefined;
            inputs["policyType"] = args ? args.policyType : undefined;
            inputs["properties"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(PolicySetDefinitionAtManagementGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a PolicySetDefinitionAtManagementGroup resource.
 */
export interface PolicySetDefinitionAtManagementGroupArgs {
    /**
     * The policy set definition description.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The display name of the policy set definition.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * The ID of the management group.
     */
    readonly managementGroupId: pulumi.Input<string>;
    /**
     * The policy set definition metadata.
     */
    readonly metadata?: pulumi.Input<{[key: string]: any}>;
    /**
     * The name of the policy set definition to create.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The policy set definition parameters that can be used in policy definition references.
     */
    readonly parameters?: pulumi.Input<{[key: string]: any}>;
    /**
     * An array of policy definition references.
     */
    readonly policyDefinitions: pulumi.Input<pulumi.Input<inputs.management.v20190101.PolicyDefinitionReference>[]>;
    /**
     * The type of policy definition. Possible values are NotSpecified, BuiltIn, and Custom.
     */
    readonly policyType?: pulumi.Input<string>;
}
