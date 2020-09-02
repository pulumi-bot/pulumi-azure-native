// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getSyncMember(args: GetSyncMemberArgs, opts?: pulumi.InvokeOptions): Promise<GetSyncMemberResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:sql/v20190601preview:getSyncMember", {
        "databaseName": args.databaseName,
        "resourceGroupName": args.resourceGroupName,
        "serverName": args.serverName,
        "syncGroupName": args.syncGroupName,
        "syncMemberName": args.syncMemberName,
    }, opts);
}

export interface GetSyncMemberArgs {
    /**
     * The name of the database on which the sync group is hosted.
     */
    readonly databaseName: string;
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the server.
     */
    readonly serverName: string;
    /**
     * The name of the sync group on which the sync member is hosted.
     */
    readonly syncGroupName: string;
    /**
     * The name of the sync member.
     */
    readonly syncMemberName: string;
}

/**
 * An Azure SQL Database sync member.
 */
export interface GetSyncMemberResult {
    /**
     * Database name of the member database in the sync member.
     */
    readonly databaseName?: string;
    /**
     * Database type of the sync member.
     */
    readonly databaseType?: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * Password of the member database in the sync member.
     */
    readonly password?: string;
    /**
     * Server name of the member database in the sync member
     */
    readonly serverName?: string;
    /**
     * SQL Server database id of the sync member.
     */
    readonly sqlServerDatabaseId?: string;
    /**
     * ARM resource id of the sync agent in the sync member.
     */
    readonly syncAgentId?: string;
    /**
     * Sync direction of the sync member.
     */
    readonly syncDirection?: string;
    /**
     * ARM resource id of the sync member logical database, for sync members in Azure.
     */
    readonly syncMemberAzureDatabaseResourceId?: string;
    /**
     * Sync state of the sync member.
     */
    readonly syncState: string;
    /**
     * Resource type.
     */
    readonly type: string;
    /**
     * Whether to use private link connection.
     */
    readonly usePrivateLinkConnection?: boolean;
    /**
     * User name of the member database in the sync member.
     */
    readonly userName?: string;
}
