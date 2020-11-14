// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getSqlManagedInstance(args: GetSqlManagedInstanceArgs, opts?: pulumi.InvokeOptions): Promise<GetSqlManagedInstanceResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:azuredata/v20200908preview:getSqlManagedInstance", {
        "resourceGroupName": args.resourceGroupName,
        "sqlManagedInstanceName": args.sqlManagedInstanceName,
    }, opts);
}

export interface GetSqlManagedInstanceArgs {
    /**
     * The name of the Azure resource group
     */
    readonly resourceGroupName: string;
    /**
     * Name of SQL Managed Instance
     */
    readonly sqlManagedInstanceName: string;
}

/**
 * A SqlManagedInstance.
 */
export interface GetSqlManagedInstanceResult {
    /**
     * The instance admin user
     */
    readonly admin?: string;
    /**
     * null
     */
    readonly dataControllerId?: string;
    /**
     * The instance end time
     */
    readonly endTime?: string;
    /**
     * The on premise instance endpoint
     */
    readonly instanceEndpoint?: string;
    /**
     * The raw kubernetes information
     */
    readonly k8sRaw?: any;
    /**
     * Last uploaded date from on premise cluster. Defaults to current date time
     */
    readonly lastUploadedDate?: string;
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * The instance start time
     */
    readonly startTime?: string;
    /**
     * Read only system data
     */
    readonly systemData: outputs.azuredata.v20200908preview.SystemDataResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
     */
    readonly type: string;
    /**
     * The instance vCore
     */
    readonly vCore?: string;
}
