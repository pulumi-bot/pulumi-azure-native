// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
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
        return new ManagementLockAtResourceLevel(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:authorization/v20160901:ManagementLockAtResourceLevel';

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
     * The name of the lock.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The properties of the lock.
     */
    public readonly properties!: pulumi.Output<outputs.authorization.v20160901.ManagementLockPropertiesResponse>;
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
    constructor(name: string, args: ManagementLockAtResourceLevelArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ManagementLockAtResourceLevelArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as ManagementLockAtResourceLevelArgs | undefined;
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.parentResourcePath === undefined) {
                throw new Error("Missing required property 'parentResourcePath'");
            }
            if (!args || args.properties === undefined) {
                throw new Error("Missing required property 'properties'");
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
            inputs["name"] = args ? args.name : undefined;
            inputs["parentResourcePath"] = args ? args.parentResourcePath : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceName"] = args ? args.resourceName : undefined;
            inputs["resourceProviderNamespace"] = args ? args.resourceProviderNamespace : undefined;
            inputs["resourceType"] = args ? args.resourceType : undefined;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ManagementLockAtResourceLevel.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ManagementLockAtResourceLevel resource.
 */
export interface ManagementLockAtResourceLevelArgs {
    /**
     * The name of lock. The lock name can be a maximum of 260 characters. It cannot contain <, > %, &, :, \, ?, /, or any control characters.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The parent resource identity.
     */
    readonly parentResourcePath: pulumi.Input<string>;
    /**
     * The properties of the lock.
     */
    readonly properties: pulumi.Input<inputs.authorization.v20160901.ManagementLockProperties>;
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
