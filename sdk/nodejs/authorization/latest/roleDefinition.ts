// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Role definition.
 *
 * ## Example Usage
 * ### GetConfigurations
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const roleDefinition = new azurerm.authorization.latest.RoleDefinition("roleDefinition", {
 *     roleDefinitionId: "roleDefinitionId",
 *     scope: "scope",
 * });
 *
 * ```
 */
export class RoleDefinition extends pulumi.CustomResource {
    /**
     * Get an existing RoleDefinition resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RoleDefinition {
        return new RoleDefinition(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:authorization/latest:RoleDefinition';

    /**
     * Returns true if the given object is an instance of RoleDefinition.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RoleDefinition {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RoleDefinition.__pulumiType;
    }

    /**
     * Role definition assignable scopes.
     */
    public readonly assignableScopes!: pulumi.Output<string[] | undefined>;
    /**
     * The role definition description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The role definition name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Role definition permissions.
     */
    public readonly permissions!: pulumi.Output<outputs.authorization.latest.PermissionResponse[] | undefined>;
    /**
     * The role name.
     */
    public readonly roleName!: pulumi.Output<string | undefined>;
    /**
     * The role type.
     */
    public readonly roleType!: pulumi.Output<string | undefined>;
    /**
     * The role definition type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a RoleDefinition resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RoleDefinitionArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.roleDefinitionId === undefined) {
                throw new Error("Missing required property 'roleDefinitionId'");
            }
            if (!args || args.scope === undefined) {
                throw new Error("Missing required property 'scope'");
            }
            inputs["assignableScopes"] = args ? args.assignableScopes : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["permissions"] = args ? args.permissions : undefined;
            inputs["roleDefinitionId"] = args ? args.roleDefinitionId : undefined;
            inputs["roleName"] = args ? args.roleName : undefined;
            inputs["roleType"] = args ? args.roleType : undefined;
            inputs["scope"] = args ? args.scope : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["assignableScopes"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["permissions"] = undefined /*out*/;
            inputs["roleName"] = undefined /*out*/;
            inputs["roleType"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:authorization/v20150701:RoleDefinition" }, { type: "azurerm:authorization/v20180101preview:RoleDefinition" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(RoleDefinition.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a RoleDefinition resource.
 */
export interface RoleDefinitionArgs {
    /**
     * Role definition assignable scopes.
     */
    readonly assignableScopes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The role definition description.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Role definition permissions.
     */
    readonly permissions?: pulumi.Input<pulumi.Input<inputs.authorization.latest.Permission>[]>;
    /**
     * The ID of the role definition.
     */
    readonly roleDefinitionId: pulumi.Input<string>;
    /**
     * The role name.
     */
    readonly roleName?: pulumi.Input<string>;
    /**
     * The role type.
     */
    readonly roleType?: pulumi.Input<string>;
    /**
     * The scope of the role definition.
     */
    readonly scope: pulumi.Input<string>;
}
