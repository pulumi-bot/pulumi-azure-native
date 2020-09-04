// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The lock information.
 */
export class ManagementLockByScope extends pulumi.CustomResource {
    /**
     * Get an existing ManagementLockByScope resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ManagementLockByScope {
        return new ManagementLockByScope(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:authorization/latest:ManagementLockByScope';

    /**
     * Returns true if the given object is an instance of ManagementLockByScope.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ManagementLockByScope {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ManagementLockByScope.__pulumiType;
    }

    /**
     * The level of the lock. Possible values are: NotSpecified, CanNotDelete, ReadOnly. CanNotDelete means authorized users are able to read and modify the resources, but not delete. ReadOnly means authorized users can only read from a resource, but they can't modify or delete it.
     */
    public readonly level!: pulumi.Output<string>;
    /**
     * The name of the lock.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Notes about the lock. Maximum of 512 characters.
     */
    public readonly notes!: pulumi.Output<string | undefined>;
    /**
     * The owners of the lock.
     */
    public readonly owners!: pulumi.Output<outputs.authorization.latest.ManagementLockOwnerResponse[] | undefined>;
    /**
     * The resource type of the lock - Microsoft.Authorization/locks.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ManagementLockByScope resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ManagementLockByScopeArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.level === undefined) {
                throw new Error("Missing required property 'level'");
            }
            if (!args || args.lockName === undefined) {
                throw new Error("Missing required property 'lockName'");
            }
            if (!args || args.scope === undefined) {
                throw new Error("Missing required property 'scope'");
            }
            inputs["level"] = args ? args.level : undefined;
            inputs["lockName"] = args ? args.lockName : undefined;
            inputs["notes"] = args ? args.notes : undefined;
            inputs["owners"] = args ? args.owners : undefined;
            inputs["scope"] = args ? args.scope : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["level"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["notes"] = undefined /*out*/;
            inputs["owners"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:authorization/v20160901:ManagementLockByScope" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ManagementLockByScope.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ManagementLockByScope resource.
 */
export interface ManagementLockByScopeArgs {
    /**
     * The level of the lock. Possible values are: NotSpecified, CanNotDelete, ReadOnly. CanNotDelete means authorized users are able to read and modify the resources, but not delete. ReadOnly means authorized users can only read from a resource, but they can't modify or delete it.
     */
    readonly level: pulumi.Input<string>;
    /**
     * The name of lock.
     */
    readonly lockName: pulumi.Input<string>;
    /**
     * Notes about the lock. Maximum of 512 characters.
     */
    readonly notes?: pulumi.Input<string>;
    /**
     * The owners of the lock.
     */
    readonly owners?: pulumi.Input<pulumi.Input<inputs.authorization.latest.ManagementLockOwner>[]>;
    /**
     * The scope for the lock. When providing a scope for the assignment, use '/subscriptions/{subscriptionId}' for subscriptions, '/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}' for resource groups, and '/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePathIfPresent}/{resourceType}/{resourceName}' for resources.
     */
    readonly scope: pulumi.Input<string>;
}
