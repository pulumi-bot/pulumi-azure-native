// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Type of the Storage Target.
 */
export class StorageTarget extends pulumi.CustomResource {
    /**
     * Get an existing StorageTarget resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): StorageTarget {
        return new StorageTarget(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:storagecache/latest:StorageTarget';

    /**
     * Returns true if the given object is an instance of StorageTarget.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StorageTarget {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StorageTarget.__pulumiType;
    }

    /**
     * Properties when targetType is clfs.
     */
    public readonly clfs!: pulumi.Output<outputs.storagecache.latest.ClfsTargetResponse | undefined>;
    /**
     * List of Cache namespace junctions to target for namespace associations.
     */
    public readonly junctions!: pulumi.Output<outputs.storagecache.latest.NamespaceJunctionResponse[] | undefined>;
    /**
     * Name of the Storage Target.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Properties when targetType is nfs3.
     */
    public readonly nfs3!: pulumi.Output<outputs.storagecache.latest.Nfs3TargetResponse | undefined>;
    /**
     * ARM provisioning state, see https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/Addendum.md#provisioningstate-property
     */
    public readonly provisioningState!: pulumi.Output<string | undefined>;
    /**
     * Type of the Storage Target.
     */
    public readonly targetType!: pulumi.Output<string>;
    /**
     * Type of the Storage Target; Microsoft.StorageCache/Cache/StorageTarget
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Properties when targetType is unknown.
     */
    public readonly unknown!: pulumi.Output<outputs.storagecache.latest.UnknownTargetResponse | undefined>;

    /**
     * Create a StorageTarget resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StorageTargetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StorageTargetArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as StorageTargetArgs | undefined;
            if (!args || args.cacheName === undefined) {
                throw new Error("Missing required property 'cacheName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.storageTargetName === undefined) {
                throw new Error("Missing required property 'storageTargetName'");
            }
            if (!args || args.targetType === undefined) {
                throw new Error("Missing required property 'targetType'");
            }
            inputs["cacheName"] = args ? args.cacheName : undefined;
            inputs["clfs"] = args ? args.clfs : undefined;
            inputs["junctions"] = args ? args.junctions : undefined;
            inputs["nfs3"] = args ? args.nfs3 : undefined;
            inputs["provisioningState"] = args ? args.provisioningState : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["storageTargetName"] = args ? args.storageTargetName : undefined;
            inputs["targetType"] = args ? args.targetType : undefined;
            inputs["unknown"] = args ? args.unknown : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:storagecache/v20190801preview:StorageTarget" }, { type: "azurerm:storagecache/v20191101:StorageTarget" }, { type: "azurerm:storagecache/v20200301:StorageTarget" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(StorageTarget.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a StorageTarget resource.
 */
export interface StorageTargetArgs {
    /**
     * Name of Cache. Length of name must be not greater than 80 and chars must be in list of [-0-9a-zA-Z_] char class.
     */
    readonly cacheName: pulumi.Input<string>;
    /**
     * Properties when targetType is clfs.
     */
    readonly clfs?: pulumi.Input<inputs.storagecache.latest.ClfsTarget>;
    /**
     * List of Cache namespace junctions to target for namespace associations.
     */
    readonly junctions?: pulumi.Input<pulumi.Input<inputs.storagecache.latest.NamespaceJunction>[]>;
    /**
     * Properties when targetType is nfs3.
     */
    readonly nfs3?: pulumi.Input<inputs.storagecache.latest.Nfs3Target>;
    /**
     * ARM provisioning state, see https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/Addendum.md#provisioningstate-property
     */
    readonly provisioningState?: pulumi.Input<string>;
    /**
     * Target resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Name of the Storage Target. Length of name must be not greater than 80 and chars must be in list of [-0-9a-zA-Z_] char class.
     */
    readonly storageTargetName: pulumi.Input<string>;
    /**
     * Type of the Storage Target.
     */
    readonly targetType: pulumi.Input<string>;
    /**
     * Properties when targetType is unknown.
     */
    readonly unknown?: pulumi.Input<inputs.storagecache.latest.UnknownTarget>;
}
