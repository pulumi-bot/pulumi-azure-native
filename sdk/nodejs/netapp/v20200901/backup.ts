// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Backup of a Volume
 */
export class Backup extends pulumi.CustomResource {
    /**
     * Get an existing Backup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Backup {
        return new Backup(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:netapp/v20200901:Backup';

    /**
     * Returns true if the given object is an instance of Backup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Backup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Backup.__pulumiType;
    }

    /**
     * UUID v4 used to identify the Backup
     */
    public /*out*/ readonly backupId!: pulumi.Output<string>;
    /**
     * Type of backup adhoc or scheduled
     */
    public /*out*/ readonly backupType!: pulumi.Output<string>;
    /**
     * The creation date of the backup
     */
    public /*out*/ readonly creationDate!: pulumi.Output<string>;
    /**
     * Label for backup
     */
    public readonly label!: pulumi.Output<string | undefined>;
    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Azure lifecycle management
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Size of backup
     */
    public /*out*/ readonly size!: pulumi.Output<number>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Backup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BackupArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.accountName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'accountName'");
            }
            if ((!args || args.backupName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'backupName'");
            }
            if ((!args || args.poolName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'poolName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.volumeName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'volumeName'");
            }
            inputs["accountName"] = args ? args.accountName : undefined;
            inputs["backupName"] = args ? args.backupName : undefined;
            inputs["label"] = args ? args.label : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["poolName"] = args ? args.poolName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["volumeName"] = args ? args.volumeName : undefined;
            inputs["backupId"] = undefined /*out*/;
            inputs["backupType"] = undefined /*out*/;
            inputs["creationDate"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["size"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["backupId"] = undefined /*out*/;
            inputs["backupType"] = undefined /*out*/;
            inputs["creationDate"] = undefined /*out*/;
            inputs["label"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["size"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:netapp/latest:Backup" }, { type: "azure-nextgen:netapp/v20200501:Backup" }, { type: "azure-nextgen:netapp/v20200601:Backup" }, { type: "azure-nextgen:netapp/v20200701:Backup" }, { type: "azure-nextgen:netapp/v20200801:Backup" }, { type: "azure-nextgen:netapp/v20201101:Backup" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Backup.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Backup resource.
 */
export interface BackupArgs {
    /**
     * The name of the NetApp account
     */
    readonly accountName: pulumi.Input<string>;
    /**
     * The name of the backup
     */
    readonly backupName: pulumi.Input<string>;
    /**
     * Label for backup
     */
    readonly label?: pulumi.Input<string>;
    /**
     * Resource location
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the capacity pool
     */
    readonly poolName: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the volume
     */
    readonly volumeName: pulumi.Input<string>;
}
