// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * The lock information.
 */
export class ManagementLockAtResourceLevel extends pulumi.CustomResource {
    /**
     * Get an existing ManagementLockAtResourceLevel resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ManagementLockAtResourceLevel {
        return new ManagementLockAtResourceLevel(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:authorization/latest:ManagementLockAtResourceLevel';

    /**
     * Returns true if the given object is an instance of ManagementLockAtResourceLevel.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ManagementLockAtResourceLevel {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ManagementLockAtResourceLevel.__pulumiType;
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
     * Create a ManagementLockAtResourceLevel resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ManagementLockAtResourceLevelArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.level === undefined) {
                throw new Error("Missing required property 'level'");
            }
            if (!args || args.lockName === undefined) {
                throw new Error("Missing required property 'lockName'");
            }
            if (!args || args.parentResourcePath === undefined) {
                throw new Error("Missing required property 'parentResourcePath'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.resourceName === undefined) {
                throw new Error("Missing required property 'resourceName'");
            }
            if (!args || args.resourceProviderNamespace === undefined) {
                throw new Error("Missing required property 'resourceProviderNamespace'");
            }
            if (!args || args.resourceType === undefined) {
                throw new Error("Missing required property 'resourceType'");
            }
            inputs["level"] = args ? args.level : undefined;
            inputs["lockName"] = args ? args.lockName : undefined;
            inputs["notes"] = args ? args.notes : undefined;
            inputs["owners"] = args ? args.owners : undefined;
            inputs["parentResourcePath"] = args ? args.parentResourcePath : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceName"] = args ? args.resourceName : undefined;
            inputs["resourceProviderNamespace"] = args ? args.resourceProviderNamespace : undefined;
            inputs["resourceType"] = args ? args.resourceType : undefined;
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
        const aliasOpts = { aliases: [{ type: "azure-nextgen:authorization/v20160901:ManagementLockAtResourceLevel" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ManagementLockAtResourceLevel.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ManagementLockAtResourceLevel resource.
 */
export interface ManagementLockAtResourceLevelArgs {
    /**
     * The level of the lock. Possible values are: NotSpecified, CanNotDelete, ReadOnly. CanNotDelete means authorized users are able to read and modify the resources, but not delete. ReadOnly means authorized users can only read from a resource, but they can't modify or delete it.
     */
    readonly level: pulumi.Input<string>;
    /**
     * The name of lock. The lock name can be a maximum of 260 characters. It cannot contain <, > %, &, :, \, ?, /, or any control characters.
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
     * The parent resource identity.
     */
    readonly parentResourcePath: pulumi.Input<string>;
    /**
     * The name of the resource group containing the resource to lock. 
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the resource to lock.
     */
    readonly resourceName: pulumi.Input<string>;
    /**
     * The resource provider namespace of the resource to lock.
     */
    readonly resourceProviderNamespace: pulumi.Input<string>;
    /**
     * The resource type of the resource to lock.
     */
    readonly resourceType: pulumi.Input<string>;
}
