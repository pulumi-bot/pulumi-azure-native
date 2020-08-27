// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * The access control record.
 */
export class AccessControlRecord extends pulumi.CustomResource {
    /**
     * Get an existing AccessControlRecord resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AccessControlRecord {
        return new AccessControlRecord(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:storsimple/v20170601:AccessControlRecord';

    /**
     * Returns true if the given object is an instance of AccessControlRecord.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AccessControlRecord {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AccessControlRecord.__pulumiType;
    }

    /**
     * The iSCSI initiator name (IQN).
     */
    public readonly initiatorName!: pulumi.Output<string>;
    /**
     * The Kind of the object. Currently only Series8000 is supported
     */
    public readonly kind!: pulumi.Output<Kind | undefined>;
    /**
     * The name of the object.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The hierarchical type of the object.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The number of volumes using the access control record.
     */
    public /*out*/ readonly volumeCount!: pulumi.Output<number>;

    /**
     * Create a AccessControlRecord resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AccessControlRecordArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccessControlRecordArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as AccessControlRecordArgs | undefined;
            if (!args || args.accessControlRecordName === undefined) {
                throw new Error("Missing required property 'accessControlRecordName'");
            }
            if (!args || args.initiatorName === undefined) {
                throw new Error("Missing required property 'initiatorName'");
            }
            if (!args || args.managerName === undefined) {
                throw new Error("Missing required property 'managerName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["accessControlRecordName"] = args ? args.accessControlRecordName : undefined;
            inputs["initiatorName"] = args ? args.initiatorName : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["managerName"] = args ? args.managerName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["volumeCount"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:storsimple/v20161001:AccessControlRecord" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(AccessControlRecord.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a AccessControlRecord resource.
 */
export interface AccessControlRecordArgs {
    /**
     * The name of the access control record.
     */
    readonly accessControlRecordName: pulumi.Input<string>;
    /**
     * The iSCSI initiator name (IQN).
     */
    readonly initiatorName: pulumi.Input<string>;
    /**
     * The Kind of the object. Currently only Series8000 is supported
     */
    readonly kind?: pulumi.Input<Kind>;
    /**
     * The manager name
     */
    readonly managerName: pulumi.Input<string>;
    /**
     * The resource group name
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
