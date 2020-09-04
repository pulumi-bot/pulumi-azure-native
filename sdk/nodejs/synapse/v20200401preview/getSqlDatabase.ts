// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getSqlDatabase(args: GetSqlDatabaseArgs, opts?: pulumi.InvokeOptions): Promise<GetSqlDatabaseResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:synapse/v20200401preview:getSqlDatabase", {
        "resourceGroupName": args.resourceGroupName,
        "sqlDatabaseName": args.sqlDatabaseName,
        "workspaceName": args.workspaceName,
    }, opts);
}

export interface GetSqlDatabaseArgs {
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the sql database.
     */
    readonly sqlDatabaseName: string;
    /**
     * The name of the workspace.
     */
    readonly workspaceName: string;
}

/**
 * A sql database resource.
 */
export interface GetSqlDatabaseResult {
    /**
     * The collation of the database.
     */
    readonly collation?: string;
    /**
     * The Guid of the database.
     */
    readonly databaseGuid: string;
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * The max size of the database expressed in bytes.
     */
    readonly maxSizeBytes?: number;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * SystemData of SqlDatabase.
     */
    readonly systemData: outputs.synapse.v20200401preview.SystemDataResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
     */
    readonly type: string;
}
