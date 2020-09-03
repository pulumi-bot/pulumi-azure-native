// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Replication protected item.
 *
 * ## Example Usage
 * ### Enables protection.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const replicationProtectedItem = new azurerm.recoveryservices.v20180110.ReplicationProtectedItem("replicationProtectedItem", {
 *     fabricName: "cloud1",
 *     protectionContainerName: "cloud_6d224fc6-f326-5d35-96de-fbf51efb3179",
 *     replicatedProtectedItemName: "f8491e4f-817a-40dd-a90c-af773978c75b",
 *     resourceGroupName: "resourceGroupPS1",
 *     resourceName: "vault1",
 * });
 *
 * ```
 */
export class ReplicationProtectedItem extends pulumi.CustomResource {
    /**
     * Get an existing ReplicationProtectedItem resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ReplicationProtectedItem {
        return new ReplicationProtectedItem(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:recoveryservices/v20180110:ReplicationProtectedItem';

    /**
     * Returns true if the given object is an instance of ReplicationProtectedItem.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ReplicationProtectedItem {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ReplicationProtectedItem.__pulumiType;
    }

    /**
     * Resource Location
     */
    public /*out*/ readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource Name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The custom data.
     */
    public readonly properties!: pulumi.Output<outputs.recoveryservices.v20180110.ReplicationProtectedItemPropertiesResponse>;
    /**
     * Resource Type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ReplicationProtectedItem resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReplicationProtectedItemArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ReplicationProtectedItemArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as ReplicationProtectedItemArgs | undefined;
            if (!args || args.fabricName === undefined) {
                throw new Error("Missing required property 'fabricName'");
            }
            if (!args || args.protectionContainerName === undefined) {
                throw new Error("Missing required property 'protectionContainerName'");
            }
            if (!args || args.replicatedProtectedItemName === undefined) {
                throw new Error("Missing required property 'replicatedProtectedItemName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.resourceName === undefined) {
                throw new Error("Missing required property 'resourceName'");
            }
            inputs["fabricName"] = args ? args.fabricName : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["protectionContainerName"] = args ? args.protectionContainerName : undefined;
            inputs["replicatedProtectedItemName"] = args ? args.replicatedProtectedItemName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceName"] = args ? args.resourceName : undefined;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:recoveryservices/latest:ReplicationProtectedItem" }, { type: "azurerm:recoveryservices/v20160810:ReplicationProtectedItem" }, { type: "azurerm:recoveryservices/v20180710:ReplicationProtectedItem" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ReplicationProtectedItem.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ReplicationProtectedItem resource.
 */
export interface ReplicationProtectedItemArgs {
    /**
     * Name of the fabric.
     */
    readonly fabricName: pulumi.Input<string>;
    /**
     * Enable protection input properties.
     */
    readonly properties?: pulumi.Input<inputs.recoveryservices.v20180110.EnableProtectionInputProperties>;
    /**
     * Protection container name.
     */
    readonly protectionContainerName: pulumi.Input<string>;
    /**
     * A name for the replication protected item.
     */
    readonly replicatedProtectedItemName: pulumi.Input<string>;
    /**
     * The name of the resource group where the recovery services vault is present.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the recovery services vault.
     */
    readonly resourceName: pulumi.Input<string>;
}
