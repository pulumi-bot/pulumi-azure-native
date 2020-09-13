// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * An Azure SQL Database sync member.
 */
export class SyncMember extends pulumi.CustomResource {
    /**
     * Get an existing SyncMember resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SyncMember {
        return new SyncMember(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:sql/v20150501preview:SyncMember';

    /**
     * Returns true if the given object is an instance of SyncMember.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SyncMember {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SyncMember.__pulumiType;
    }

    /**
     * Database name of the member database in the sync member.
     */
    public readonly databaseName!: pulumi.Output<string | undefined>;
    /**
     * Database type of the sync member.
     */
    public readonly databaseType!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Password of the member database in the sync member.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * Server name of the member database in the sync member
     */
    public readonly serverName!: pulumi.Output<string | undefined>;
    /**
     * SQL Server database id of the sync member.
     */
    public readonly sqlServerDatabaseId!: pulumi.Output<string | undefined>;
    /**
     * ARM resource id of the sync agent in the sync member.
     */
    public readonly syncAgentId!: pulumi.Output<string | undefined>;
    /**
     * Sync direction of the sync member.
     */
    public readonly syncDirection!: pulumi.Output<string | undefined>;
    /**
     * Sync state of the sync member.
     */
    public /*out*/ readonly syncState!: pulumi.Output<string>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * User name of the member database in the sync member.
     */
    public readonly userName!: pulumi.Output<string | undefined>;

    /**
     * Create a SyncMember resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SyncMemberArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.databaseName === undefined) {
                throw new Error("Missing required property 'databaseName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.serverName === undefined) {
                throw new Error("Missing required property 'serverName'");
            }
            if (!args || args.syncGroupName === undefined) {
                throw new Error("Missing required property 'syncGroupName'");
            }
            if (!args || args.syncMemberName === undefined) {
                throw new Error("Missing required property 'syncMemberName'");
            }
            inputs["databaseName"] = args ? args.databaseName : undefined;
            inputs["databaseType"] = args ? args.databaseType : undefined;
            inputs["password"] = args ? args.password : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serverName"] = args ? args.serverName : undefined;
            inputs["sqlServerDatabaseId"] = args ? args.sqlServerDatabaseId : undefined;
            inputs["syncAgentId"] = args ? args.syncAgentId : undefined;
            inputs["syncDirection"] = args ? args.syncDirection : undefined;
            inputs["syncGroupName"] = args ? args.syncGroupName : undefined;
            inputs["syncMemberName"] = args ? args.syncMemberName : undefined;
            inputs["userName"] = args ? args.userName : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["syncState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["databaseName"] = undefined /*out*/;
            inputs["databaseType"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["password"] = undefined /*out*/;
            inputs["serverName"] = undefined /*out*/;
            inputs["sqlServerDatabaseId"] = undefined /*out*/;
            inputs["syncAgentId"] = undefined /*out*/;
            inputs["syncDirection"] = undefined /*out*/;
            inputs["syncState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["userName"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:sql/v20190601preview:SyncMember" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(SyncMember.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a SyncMember resource.
 */
export interface SyncMemberArgs {
    /**
     * Database name of the member database in the sync member.
     */
    readonly databaseName: pulumi.Input<string>;
    /**
     * Database type of the sync member.
     */
    readonly databaseType?: pulumi.Input<string>;
    /**
     * Password of the member database in the sync member.
     */
    readonly password?: pulumi.Input<string>;
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Server name of the member database in the sync member
     */
    readonly serverName: pulumi.Input<string>;
    /**
     * SQL Server database id of the sync member.
     */
    readonly sqlServerDatabaseId?: pulumi.Input<string>;
    /**
     * ARM resource id of the sync agent in the sync member.
     */
    readonly syncAgentId?: pulumi.Input<string>;
    /**
     * Sync direction of the sync member.
     */
    readonly syncDirection?: pulumi.Input<string>;
    /**
     * The name of the sync group on which the sync member is hosted.
     */
    readonly syncGroupName: pulumi.Input<string>;
    /**
     * The name of the sync member.
     */
    readonly syncMemberName: pulumi.Input<string>;
    /**
     * User name of the member database in the sync member.
     */
    readonly userName?: pulumi.Input<string>;
}
