// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getTask(args: GetTaskArgs, opts?: pulumi.InvokeOptions): Promise<GetTaskResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:datamigration/latest:getTask", {
        "expand": args.expand,
        "groupName": args.groupName,
        "projectName": args.projectName,
        "serviceName": args.serviceName,
        "taskName": args.taskName,
    }, opts);
}

export interface GetTaskArgs {
    /**
     * Expand the response
     */
    readonly expand?: string;
    /**
     * Name of the resource group
     */
    readonly groupName: string;
    /**
     * Name of the project
     */
    readonly projectName: string;
    /**
     * Name of the service
     */
    readonly serviceName: string;
    /**
     * Name of the Task
     */
    readonly taskName: string;
}

/**
 * A task resource
 */
export interface GetTaskResult {
    /**
     * HTTP strong entity tag value. This is ignored if submitted.
     */
    readonly etag?: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * Custom task properties
     */
    readonly properties: outputs.datamigration.latest.ConnectToMongoDbTaskPropertiesResponse | outputs.datamigration.latest.ConnectToSourceOracleSyncTaskPropertiesResponse | outputs.datamigration.latest.ConnectToSourcePostgreSqlSyncTaskPropertiesResponse | outputs.datamigration.latest.ConnectToSourceSqlServerSyncTaskPropertiesResponse | outputs.datamigration.latest.ConnectToSourceSqlServerTaskPropertiesResponse | outputs.datamigration.latest.ConnectToTargetAzureDbForMySqlTaskPropertiesResponse | outputs.datamigration.latest.ConnectToTargetAzureDbForPostgreSqlSyncTaskPropertiesResponse | outputs.datamigration.latest.ConnectToTargetOracleAzureDbForPostgreSqlSyncTaskPropertiesResponse | outputs.datamigration.latest.ConnectToTargetSqlDbTaskPropertiesResponse | outputs.datamigration.latest.ConnectToTargetSqlMISyncTaskPropertiesResponse | outputs.datamigration.latest.ConnectToTargetSqlMITaskPropertiesResponse | outputs.datamigration.latest.ConnectToTargetSqlSqlDbSyncTaskPropertiesResponse | outputs.datamigration.latest.GetTdeCertificatesSqlTaskPropertiesResponse | outputs.datamigration.latest.GetUserTablesOracleTaskPropertiesResponse | outputs.datamigration.latest.GetUserTablesPostgreSqlTaskPropertiesResponse | outputs.datamigration.latest.GetUserTablesSqlSyncTaskPropertiesResponse | outputs.datamigration.latest.GetUserTablesSqlTaskPropertiesResponse | outputs.datamigration.latest.MigrateMongoDbTaskPropertiesResponse | outputs.datamigration.latest.MigrateMySqlAzureDbForMySqlSyncTaskPropertiesResponse | outputs.datamigration.latest.MigrateOracleAzureDbForPostgreSqlSyncTaskPropertiesResponse | outputs.datamigration.latest.MigratePostgreSqlAzureDbForPostgreSqlSyncTaskPropertiesResponse | outputs.datamigration.latest.MigrateSqlServerSqlDbSyncTaskPropertiesResponse | outputs.datamigration.latest.MigrateSqlServerSqlDbTaskPropertiesResponse | outputs.datamigration.latest.MigrateSqlServerSqlMISyncTaskPropertiesResponse | outputs.datamigration.latest.MigrateSqlServerSqlMITaskPropertiesResponse | outputs.datamigration.latest.MigrateSsisTaskPropertiesResponse | outputs.datamigration.latest.ValidateMigrationInputSqlServerSqlDbSyncTaskPropertiesResponse | outputs.datamigration.latest.ValidateMigrationInputSqlServerSqlMISyncTaskPropertiesResponse | outputs.datamigration.latest.ValidateMigrationInputSqlServerSqlMITaskPropertiesResponse | outputs.datamigration.latest.ValidateMongoDbTaskPropertiesResponse | outputs.datamigration.latest.ValidateOracleAzureDbForPostgreSqlSyncTaskPropertiesResponse;
    /**
     * Resource type.
     */
    readonly type: string;
}
