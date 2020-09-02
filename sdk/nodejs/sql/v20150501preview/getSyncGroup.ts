// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getSyncGroup(args: GetSyncGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetSyncGroupResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:sql/v20150501preview:getSyncGroup", {
        "databaseName": args.databaseName,
        "resourceGroupName": args.resourceGroupName,
        "serverName": args.serverName,
        "syncGroupName": args.syncGroupName,
    }, opts);
}

export interface GetSyncGroupArgs {
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
     * The name of the sync group.
     */
    readonly syncGroupName: string;
}

/**
 * An Azure SQL Database sync group.
 */
export interface GetSyncGroupResult {
    /**
     * Conflict resolution policy of the sync group.
     */
    readonly conflictResolutionPolicy?: string;
    /**
     * Password for the sync group hub database credential.
     */
    readonly hubDatabasePassword?: string;
    /**
     * User name for the sync group hub database credential.
     */
    readonly hubDatabaseUserName?: string;
    /**
     * Sync interval of the sync group.
     */
    readonly interval?: number;
    /**
     * Last sync time of the sync group.
     */
    readonly lastSyncTime: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * Sync schema of the sync group.
     */
    readonly schema?: outputs.sql.v20150501preview.SyncGroupSchemaResponse;
    /**
     * ARM resource id of the sync database in the sync group.
     */
    readonly syncDatabaseId?: string;
    /**
     * Sync state of the sync group.
     */
    readonly syncState: string;
    /**
     * Resource type.
     */
    readonly type: string;
}
