// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * The policy assignment.
 */
export class PolicyAssignment extends pulumi.CustomResource {
    /**
     * Get an existing PolicyAssignment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PolicyAssignment {
        return new PolicyAssignment(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:authorization/v20151001preview:PolicyAssignment';

    /**
     * Returns true if the given object is an instance of PolicyAssignment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PolicyAssignment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PolicyAssignment.__pulumiType;
    }

    /**
     * The display name of the policy assignment.
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * The name of the policy assignment.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * The ID of the policy definition.
     */
    public readonly policyDefinitionId!: pulumi.Output<string | undefined>;
    /**
     * The scope for the policy assignment.
     */
    public readonly scope!: pulumi.Output<string | undefined>;
    /**
     * The type of the policy assignment.
     */
    public readonly type!: pulumi.Output<string | undefined>;

    /**
     * Create a PolicyAssignment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PolicyAssignmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PolicyAssignmentArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as PolicyAssignmentArgs | undefined;
            if (!args || args.policyAssignmentName === undefined) {
                throw new Error("Missing required property 'policyAssignmentName'");
            }
            if (!args || args.scope === undefined) {
                throw new Error("Missing required property 'scope'");
            }
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["policyAssignmentName"] = args ? args.policyAssignmentName : undefined;
            inputs["policyDefinitionId"] = args ? args.policyDefinitionId : undefined;
            inputs["scope"] = args ? args.scope : undefined;
            inputs["type"] = args ? args.type : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:authorization/latest:PolicyAssignment" }, { type: "azurerm:authorization/v20151101:PolicyAssignment" }, { type: "azurerm:authorization/v20160401:PolicyAssignment" }, { type: "azurerm:authorization/v20161201:PolicyAssignment" }, { type: "azurerm:authorization/v20170601preview:PolicyAssignment" }, { type: "azurerm:authorization/v20180301:PolicyAssignment" }, { type: "azurerm:authorization/v20180501:PolicyAssignment" }, { type: "azurerm:authorization/v20190101:PolicyAssignment" }, { type: "azurerm:authorization/v20190601:PolicyAssignment" }, { type: "azurerm:authorization/v20190901:PolicyAssignment" }, { type: "azurerm:authorization/v20200301:PolicyAssignment" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(PolicyAssignment.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a PolicyAssignment resource.
 */
export interface PolicyAssignmentArgs {
    /**
     * The display name of the policy assignment.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * The ID of the policy assignment.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * The name of the policy assignment.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the policy assignment.
     */
    readonly policyAssignmentName: pulumi.Input<string>;
    /**
     * The ID of the policy definition.
     */
    readonly policyDefinitionId?: pulumi.Input<string>;
    /**
     * The scope for the policy assignment.
     */
    readonly scope: pulumi.Input<string>;
    /**
     * The type of the policy assignment.
     */
    readonly type?: pulumi.Input<string>;
}
